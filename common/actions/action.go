package actions

import (
	"context"
	"fmt"
	"time"

	"github.com/iotexproject/iotex-analyser-api/db"
)

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
	query := "SELECT a.action_hash act_hash,a.action_type act_type,a.sender,a.recipient,a.amount,a.gas_price*r.gas_consumed as gas_fee,a.block_height blk_height,b.block_hash blk_hash,b.timestamp FROM (SELECT a.action_hash,a.action_type,a.sender,a.recipient,a.amount,a.block_height,a.gas_price FROM block_action a where a.sender=? or a.recipient=? order by id desc limit ? offset ?) a left join block b on b.block_height=a.block_height left join block_receipt r on r.action_hash=a.action_hash"
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
	if err := db.Raw(query, startHeight, endHeight, startDate, endDate).Count(&count).Error; err != nil {
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
	query := "SELECT a.action_hash act_hash,a.action_type act_type,a.sender,a.recipient,a.amount,a.gas_price*r.gas_consumed as gas_fee,a.block_height blk_height,b.block_hash blk_hash,b.timestamp FROM block_action a left join block b on b.block_height=a.block_height left join block_receipt r on r.action_hash=a.action_hash where b.block_height>=? and b.block_height<=? and b.timestamp>=? and b.timestamp<=? order by a.id asc limit ? offset ?"
	if err := db.Raw(query, startHeight, endHeight, startDate, endDate, first, skip).Scan(&actionInfos).Error; err != nil {
		return nil, err
	}
	return actionInfos, nil
}

func GetActionInfoByHash(actHash string) (*ActionInfo, error) {
	var actionInfo *ActionInfo
	db := db.DB()

	query := "SELECT a.action_hash act_hash,a.action_type act_type,a.sender,a.recipient,a.amount,a.gas_price*r.gas_consumed as gas_fee,a.block_height blk_height,b.block_hash blk_hash,b.timestamp FROM block_action a left join block b on b.block_height=a.block_height left join block_receipt r on r.action_hash=a.action_hash where a.action_hash=?"
	if err := db.Raw(query, actHash).Scan(&actionInfo).Error; err != nil {
		return nil, err
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
	query := "SELECT a.action_hash act_hash,a.action_type act_type,a.sender,a.recipient,a.amount,a.gas_price*r.gas_consumed as gas_fee,a.block_height blk_height,b.block_hash blk_hash,b.timestamp FROM (SELECT a.action_hash,a.action_type,a.sender,a.recipient,a.amount,a.block_height,a.gas_price FROM block_action a where a.action_type=? order by id desc limit ? offset ?) a left join block b on b.block_height=a.block_height left join block_receipt r on r.action_hash=a.action_hash"
	if err := db.WithContext(ctx).Raw(query, typ, first, skip).Scan(&actionInfos).Error; err != nil {
		return nil, err
	}
	return actionInfos, nil
}
