package actions

import "github.com/iotexproject/iotex-analyser-api/db"

func GetActionCountByAddress(addr string) (int64, error) {
	var count int64
	db := db.DB()
	if err := db.Table("block_action a").Where("a.sender=? or a.recipient=?", addr, addr).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetActionInfoByAddress(addr string, skip, first uint64) ([]*ActionInfo, error) {
	var actionInfos []*ActionInfo
	db := db.DB()
	query := "SELECT a.action_hash act_hash,a.action_type act_type,a.sender,a.recipient,a.amount,a.gas_price*r.gas_consumed as gas_fee,a.block_height blk_height,b.block_hash blk_hash,b.timestamp FROM block_action a inner join block b on b.block_height=a.block_height inner join block_receipt r on r.action_hash=a.action_hash where a.sender=? or a.recipient=? order by a.id asc limit ? offset ?"
	if err := db.Raw(query, addr, addr, first, skip).Scan(&actionInfos).Error; err != nil {
		return nil, err
	}
	return actionInfos, nil
}
