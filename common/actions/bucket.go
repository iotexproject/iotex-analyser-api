package actions

import (
	"math/big"

	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-analyser-api/model"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func GetBucketIDsByVoter(address string) ([]uint64, error) {
	var buckets []uint64
	db := db.DB()
	query := "select bucket_id from staking_actions t1 where owner_address=? and id in (select max(id) from staking_actions t2 where t1.bucket_id=t2.bucket_id)"
	if err := db.Raw(query, address).Find(&buckets).Error; err != nil {
		return nil, err
	}
	return buckets, nil
}

func GetBucketActionCountByBuckets(bucketIDs []uint64) (int64, error) {
	var count int64
	db := db.DB()
	query := "select count(distinct t1.act_hash) from staking_actions t1 where t1.bucket_id in (?)"
	if err := db.Raw(query, bucketIDs).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetBucketActionInfoByBuckets(bucketIDs []uint64, skip, first uint64) ([]*ActionInfo, error) {
	var actionInfos []*ActionInfo
	db := db.DB()
	query := "select distinct action_hash act_hash,t2.action_type act_type,t2.sender,t2.recipient,t2.amount,t2.gas_price*t2.gas_consumed as gas_fee,t2.block_height blk_height,t1.block_hash blk_hash,t1.timestamp from  block_action t2 left join block t1 on t1.block_height=t2.block_height  where t2.action_hash in (select distinct t1.act_hash from staking_actions t1 where t1.bucket_id in (?)) order by t1.timestamp desc limit ? offset ?"
	if err := db.Raw(query, bucketIDs, first, skip).Scan(&actionInfos).Error; err != nil {
		return nil, err
	}
	return actionInfos, nil
}

type VoteBucketList struct {
	EpochNumber uint64
	BucketList  []byte
}

func GetVoteBucketList(epochNum uint64) (*iotextypes.VoteBucketList, error) {
	voteBucketListAll := &iotextypes.VoteBucketList{}
	var vbl VoteBucketList
	if err := db.DB().Table("vote_bucketlist").Where("epoch_number = ?", epochNum).First(&vbl).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to get vote bucket list in epoch %d", epochNum)
	}
	if err := proto.Unmarshal(vbl.BucketList, voteBucketListAll); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal vote bucket list in epoch %d", epochNum)
	}
	return voteBucketListAll, nil
}

func getLatestStakingBucketOwnerWithHeight(bucketID, height uint64) (*model.StakingBucket, error) {
	var stakingBucket model.StakingBucket
	db := db.DB()
	if err := db.Table("staking_buckets").Where("block_height<=? and bucket_id=?", height, bucketID).Order("id desc").Limit(1).Scan(&stakingBucket).Error; err != nil {
		return nil, err
	}
	return &stakingBucket, nil
}

func GetBucketIDsByVoterAndHeight(addr string, height uint64) ([]uint64, error) {
	db := db.DB()
	var ids []struct {
		BucketID uint64
	}
	if err := db.Table("staking_buckets").Distinct("bucket_id").Where("block_height<=? and owner_address=?", height, addr).Find(&ids).Error; err != nil {
		return nil, err
	}
	bucketID := []uint64{}
	for _, id := range ids {
		bucketID = append(bucketID, id.BucketID)
	}
	return bucketID, nil
}

func GetStakedBucketByVoterAndHeight(addr string, height uint64) (*big.Int, *big.Int, error) {
	db := db.DB()

	bucketIDs, err := GetBucketIDsByVoterAndHeight(addr, height)
	if err != nil {
		return nil, nil, err
	}
	var stakingBuckets []*model.StakingBucket
	query := "select t1.* from staking_buckets t1 INNER JOIN (select MAX(id)AS max_id  from staking_buckets t4 where block_height<=? and bucket_id in ? GROUP BY bucket_id) as t2 on t2.max_id=t1.id"
	if err := db.Raw(query, height, bucketIDs).Scan(&stakingBuckets).Error; err != nil {
		return nil, nil, err
	}
	totalStakeAmount := big.NewInt(0)
	totalVotingPower := big.NewInt(0)
	for _, stakingBucket := range stakingBuckets {
		if addr != stakingBucket.OwnerAddress {
			continue
		}
		stakeAmount, _ := big.NewInt(0).SetString(stakingBucket.StakedAmount, 0)
		totalStakeAmount.Add(totalStakeAmount, stakeAmount)
		votingPower, _ := big.NewInt(0).SetString(stakingBucket.VotingPower, 0)
		totalVotingPower.Add(totalVotingPower, votingPower)
	}
	// if err := db.Table("staking_buckets").Distinct("bucket_id").Where("block_height<=? and owner_address=?", height, addr).Find(&ids).Error; err != nil {
	// 	return nil, nil, err
	// }
	// totalStakeAmount := big.NewInt(0)
	// totalVotingPower := big.NewInt(0)
	// for _, id := range ids {
	// 	stakingBucket, err := getLatestStakingBucketOwnerWithHeight(id.BucketID, height)
	// 	if err != nil {
	// 		return nil, nil, err
	// 	}
	// 	if addr != stakingBucket.OwnerAddress {
	// 		continue
	// 	}
	// 	stakeAmount, _ := big.NewInt(0).SetString(stakingBucket.StakedAmount, 0)
	// 	totalStakeAmount.Add(totalStakeAmount, stakeAmount)
	// 	votingPower, _ := big.NewInt(0).SetString(stakingBucket.VotingPower, 0)
	// 	totalVotingPower.Add(totalVotingPower, votingPower)
	// }
	return totalStakeAmount, totalVotingPower, nil
}
