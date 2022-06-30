package actions

import (
	"context"

	"github.com/iotexproject/iotex-analyser-api/db"
)

func GetEvmTransferCount(ctx context.Context, address string) (int64, error) {
	var count int64
	db := db.DB()
	query := "select count(id) filter (where type='execution') from block_receipt_transactions where  (sender=? or recipient=?)"
	if err := db.WithContext(ctx).Raw(query, address, address).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetEvmTransferInfoByAddress(ctx context.Context, address string, skip, first uint64) ([]*ActionInfo, error) {
	var actionInfos []*ActionInfo
	db := db.DB()
	query := "select t1.block_height blk_height, t1.action_hash act_hash, t2.block_hash blk_hash,t1.type act_type,sender,recipient,t1.amount,t2.timestamp from block_receipt_transactions t1 left join block t2 on t2.block_height=t1.block_height where (sender=? or recipient=?) and type='execution' order by t2.block_height desc limit ? offset ?"
	if err := db.Raw(query, address, address, first, skip).Scan(&actionInfos).Error; err != nil {
		return nil, err
	}
	return actionInfos, nil
}
