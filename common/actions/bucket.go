package actions

import (
	"github.com/iotexproject/iotex-analyser-api/db"
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
