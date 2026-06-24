package actions

import (
	"time"

	"github.com/iotexproject/iotex-analyser-api/db"
)

type XrcInfo struct {
	ActionHash      string
	Sender          string
	Recipient       string
	Amount          string
	BlockHeight     uint64
	Timestamp       time.Time
	ContractAddress string
}

func GetXrc20CountByAddress(address string) (int64, error) {
	var count int64
	db := db.DB()
	query := "select count(id) from erc20_transfers where  (sender=? or recipient=?)"
	if err := db.Raw(query, address, address).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// `OR ... ORDER BY id DESC LIMIT N` on erc20_transfers makes the planner pick
// a backward pkey scan (~120s on a 220M-row table for addresses that match
// late). Split into two index scans + UNION so each leg uses its own
// (sender|recipient, id DESC) composite index. Inner LIMIT must be skip+first
// for deep pages to be correct; UNION (not UNION ALL) drops the self-transfer
// duplicate by id.
func GetXrc20InfoByAddress(address string, skip, first uint64) ([]*XrcInfo, error) {
	var lists []*XrcInfo
	db := db.DB()
	inner := skip + first
	query := `SELECT t.action_hash, t.sender, t.recipient, t.amount, t.block_height, t.timestamp, t.contract_address FROM (
		(SELECT id, action_hash, sender, recipient, amount, block_height, timestamp, contract_address
		 FROM erc20_transfers WHERE sender=? ORDER BY id DESC LIMIT ?)
		UNION
		(SELECT id, action_hash, sender, recipient, amount, block_height, timestamp, contract_address
		 FROM erc20_transfers WHERE recipient=? ORDER BY id DESC LIMIT ?)
	) t ORDER BY t.id DESC LIMIT ? OFFSET ?`
	if err := db.Raw(query, address, inner, address, inner, first, skip).Scan(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func GetXrc20CountByContractAddress(address string) (int64, error) {
	var count int64
	db := db.DB()
	query := "select count(*) from erc20_transfers where contract_address=?"
	if err := db.Raw(query, address, address).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetXrc20InfoByContractAddress(address string, skip, first uint64) ([]*XrcInfo, error) {
	var lists []*XrcInfo
	db := db.DB()
	query := "select t1.action_hash,t1.sender,t1.recipient,t1.amount,t1.block_height,t1.timestamp,t1.contract_address from erc20_transfers t1 where contract_address=? order by t1.id desc limit ? offset ?"
	if err := db.Raw(query, address, first, skip).Scan(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func GetXrc20Count() (int64, error) {
	var count int64
	db := db.DB()
	//improve count performance
	query := "select reltuples::bigint as estimate_rows from pg_class where relname = 'erc20_transfers'"
	if err := db.Raw(query).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetXrc20InfoByPage(skip, first uint64) ([]*XrcInfo, error) {
	var lists []*XrcInfo
	db := db.DB()
	query := "select t1.action_hash,t1.sender,t1.recipient,t1.amount,t1.block_height,t1.timestamp,t1.contract_address from erc20_transfers t1 order by t1.id desc limit ? offset ?"
	if err := db.Raw(query, first, skip).Scan(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func GetXrc20ContractAddressCount() (int64, error) {
	var count int64
	db := db.DB()
	query := "SELECT COUNT(*) FROM (SELECT DISTINCT contract_address FROM erc20_holders) AS temp"
	if err := db.Raw(query).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetXrc20ContractAddress(skip, first uint64) ([]string, error) {
	var lists []string
	db := db.DB()
	query := "SELECT DISTINCT contract_address FROM erc20_holders limit ? offset ?"
	if err := db.Raw(query, first, skip).Scan(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func GetXrc20TokenHolderCountByTokenAddress(contractAddress string) (int64, error) {
	var count int64
	db := db.DB()
	query := "SELECT COUNT(1) FROM erc20_holders WHERE contract_address=?"
	if err := db.Raw(query, contractAddress).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetXrc20TokenHolder(contractAddress string, skip, first uint64) ([]string, error) {
	var lists []string
	db := db.DB()
	query := "SELECT holder FROM erc20_holders WHERE contract_address=? limit ? offset ?"
	if err := db.Raw(query, contractAddress, first, skip).Scan(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}
