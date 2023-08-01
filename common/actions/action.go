package actions

import (
	"context"
	"fmt"
	"time"

	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/db"
)

// GetBlockTimes returns the block times
func GetBlockTimes() (time.Time, time.Time, error) {
	db := db.DB()
	var result struct {
		MinTime time.Time
		MaxTime time.Time
	}
	query := "select (select timestamp min_time from block order by block_height asc limit 1), (select timestamp max_time from block order by block_height desc limit 1)"
	if err := db.Raw(query).Scan(&result).Error; err != nil {
		return time.Time{}, time.Time{}, err
	}
	return result.MinTime, result.MaxTime, nil
}

func GetBlockStatsByDate(unixtime int64) (minBlkHeight uint64, maxBlkHeight uint64, totalSize uint64, err error) {
	db := db.DB()
	date := time.Unix(unixtime, 0).Format("2006-01-02")
	var result struct {
		MinBlkHeight uint64
		MaxBlkHeight uint64
		TotalSize    uint64
	}
	//use timestamp index to get the block height
	query := fmt.Sprintf("select min(b.block_height) min_blk_height,max(b.block_height) max_blk_height,sum(bm.block_size) total_size from block b inner join block_meta bm using(block_height) where timestamp::date='%s'::date", date)
	if err := db.Raw(query).Scan(&result).Error; err != nil {
		return 0, 0, 0, err
	}
	return result.MinBlkHeight, result.MaxBlkHeight, result.TotalSize, nil
}

func fromUnixTime(unixtime uint64) string {
	return time.Unix(int64(unixtime), 0).UTC().Format("2006-01-02 15:04:05")
}

func getHeightByDate(unixtime uint64) (minBlkHeight uint64, maxBlkHeight uint64, err error) {
	db := db.DB()
	date := time.Unix(int64(unixtime), 0).Format("2006-01-02")
	var result struct {
		MinBlkHeight uint64
		MaxBlkHeight uint64
	}
	//use timestamp index to get the block height
	query := fmt.Sprintf("select min(block_height) min_blk_height,max(block_height) max_blk_height from block_action where timestamp::date='%s'::date group by timestamp::date", date)
	if err := db.Raw(query).Scan(&result).Error; err != nil {
		return 0, 0, err
	}
	return result.MinBlkHeight, result.MaxBlkHeight, nil
}

func GetActionCountByAddress(ctx context.Context, addr string) (int64, error) {
	var count int64
	db := db.DB()
	query := "select (SELECT count(*) FROM block_action a WHERE a.sender=?)+(SELECT count(*) FROM block_action a WHERE a.recipient=?)"
	if err := db.WithContext(ctx).Raw(query, addr, addr).Scan(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetActionInfoByAddress(ctx context.Context, addr string, skip, first uint64) ([]*ActionInfo, error) {
	var actionInfos []*ActionInfo
	db := db.DB()
	query := "SELECT a.action_hash act_hash,a.action_type act_type,a.sender,a.recipient,a.amount,a.gas_price*r.gas_consumed as gas_fee,a.block_height blk_height,b.block_hash blk_hash,b.timestamp FROM (select * from (SELECT a.id,a.action_hash,a.action_type,a.sender,a.recipient,a.amount,a.block_height,a.gas_price FROM block_action a where a.sender=? union all SELECT a.id,a.action_hash,a.action_type,a.sender,a.recipient,a.amount,a.block_height,a.gas_price FROM block_action a where a.recipient=?)tmp order by id desc limit ? offset ?) a left join block b on b.block_height=a.block_height left join block_receipts r on r.action_hash=a.action_hash"
	if err := db.WithContext(ctx).Raw(query, addr, addr, first, skip).Scan(&actionInfos).Error; err != nil {
		return nil, err
	}
	return actionInfos, nil
}

func GetActionCountByDates(startDate, endDate uint64) (int64, error) {
	var count int64
	db := db.DB()
	startHeight, _, err := getHeightByDate(startDate)
	if err != nil {
		return 0, err
	}
	_, endHeight, err := getHeightByDate(endDate)
	if err != nil {
		return 0, err
	}
	query := "SELECT count(*) FROM block_action a left join block b on b.block_height=a.block_height where b.block_height>=? and b.block_height<=? and b.timestamp>=? and b.timestamp<=?"
	if err := db.Raw(query, startHeight, endHeight, fromUnixTime(startDate), fromUnixTime(endDate)).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetActionInfoByDates(startDate, endDate uint64, skip, first uint64) ([]*ActionInfo, error) {
	var actionInfos []*ActionInfo
	db := db.DB()
	startHeight, _, err := getHeightByDate(startDate)
	if err != nil {
		return nil, err
	}
	_, endHeight, err := getHeightByDate(endDate)
	if err != nil {
		return nil, err
	}
	query := "SELECT a.action_hash act_hash,a.action_type act_type,a.sender,a.recipient,a.amount,a.gas_price*r.gas_consumed as gas_fee,a.block_height blk_height,b.block_hash blk_hash,b.timestamp FROM block_action a left join block b on b.block_height=a.block_height left join block_receipts r on r.action_hash=a.action_hash where b.block_height>=? and b.block_height<=? and b.timestamp>=? and b.timestamp<=? order by a.id asc limit ? offset ?"
	if err := db.Raw(query, startHeight, endHeight, fromUnixTime(startDate), fromUnixTime(endDate), first, skip).Scan(&actionInfos).Error; err != nil {
		return nil, err
	}
	return actionInfos, nil
}

func GetActionInfoByHash(actHash string) (*ActionInfo, error) {
	var actionInfo *ActionInfo
	db := db.DB()

	query := "SELECT a.action_hash act_hash,a.action_type act_type,a.sender,a.recipient,a.amount,a.gas_price*r.gas_consumed as gas_fee,a.block_height blk_height,b.block_hash blk_hash,b.timestamp FROM block_action a left join block b on b.block_height=a.block_height left join block_receipts r on r.action_hash=a.action_hash where a.action_hash=?"
	if err := db.Raw(query, actHash).Scan(&actionInfo).Error; err != nil {
		return nil, err
	}
	if actionInfo == nil {
		return nil, common.ErrActionNotExist
	}
	return actionInfo, nil
}

func GetActionInfoByBlockHeightAndContractAddress(blockHeight uint64, contractAddress string) (*ActionInfo, error) {
	var actionInfo *ActionInfo
	db := db.DB()

	query := "SELECT a.action_hash act_hash,a.action_type act_type,a.sender,a.recipient,a.amount,a.gas_price*r.gas_consumed as gas_fee,a.block_height blk_height,b.block_hash blk_hash,b.timestamp FROM block_action a left join block b on b.block_height=a.block_height left join block_receipts r on r.action_hash=a.action_hash where a.block_height=? and (a.contract_address=? or a.recipient=?)"
	if err := db.Raw(query, blockHeight, contractAddress, contractAddress).Scan(&actionInfo).Error; err != nil {
		return nil, err
	}
	if actionInfo == nil {
		return nil, common.ErrActionNotExist
	}
	return actionInfo, nil
}

func GetBlockReceiptTransactionByHash(actHash string) ([]*BlockReceiptTransaction, error) {
	var blkReceiptTransactions []*BlockReceiptTransaction
	db := db.DB()

	query := "SELECT * FROM block_receipt_transactions WHERE action_hash=?"
	if err := db.Raw(query, actHash).Scan(&blkReceiptTransactions).Error; err != nil {
		return nil, err
	}
	return blkReceiptTransactions, nil
}

func GetActionCountByType(ctx context.Context, typ string) (int64, error) {
	var count int64
	db := db.DB()
	//TODO: fix the slow query to get the count
	if err := db.Table("block_action a").WithContext(ctx).Where("a.action_type=?", typ).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetActionInfoByType(ctx context.Context, typ string, skip, first uint64) ([]*ActionInfo, error) {
	var actionInfos []*ActionInfo
	db := db.DB()
	query := "SELECT a.action_hash act_hash,a.action_type act_type,a.sender,a.recipient,a.amount,a.gas_price*r.gas_consumed as gas_fee,a.block_height blk_height,b.block_hash blk_hash,b.timestamp FROM (SELECT a.action_hash,a.action_type,a.sender,a.recipient,a.amount,a.block_height,a.gas_price FROM block_action a where a.action_type=? order by id desc limit ? offset ?) a left join block b on b.block_height=a.block_height left join block_receipts r on r.action_hash=a.action_hash"
	if err := db.WithContext(ctx).Raw(query, typ, first, skip).Scan(&actionInfos).Error; err != nil {
		return nil, err
	}
	return actionInfos, nil
}
