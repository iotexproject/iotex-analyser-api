package actions

import (
	"github.com/iotexproject/iotex-analyser-api/db"
)

func GetXrc721CountByAddress(address string) (int64, error) {
	var count int64
	db := db.DB()
	query := "select count(id) from erc721_transfers where  (sender=? or recipient=?)"
	if err := db.Raw(query, address, address).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetXrc721InfoByAddress(address string, skip, first uint64) ([]*XrcInfo, error) {
	var lists []*XrcInfo
	db := db.DB()
	query := "select t1.action_hash,t1.sender,t1.recipient,t1.token_id amount,t1.block_height,t1.timestamp,t1.contract_address from erc721_transfers t1 where (sender=? or recipient=?) order by t1.id desc limit ? offset ?"
	if err := db.Raw(query, address, address, first, skip).Scan(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func GetXrc721CountByContractAddress(address string) (int64, error) {
	var count int64
	db := db.DB()
	query := "select count(id) from erc721_transfers where  contract_address=?"
	if err := db.Raw(query, address, address).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetXrc721InfoByContractAddress(address string, skip, first uint64) ([]*XrcInfo, error) {
	var lists []*XrcInfo
	db := db.DB()
	query := "select t1.action_hash,t1.sender,t1.recipient,t1.token_id amount,t1.block_height,t1.timestamp,t1.contract_address from erc721_transfers t1 where contract_address=? order by t1.id desc limit ? offset ?"
	if err := db.Raw(query, address, first, skip).Scan(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func GetXrc721Count() (int64, error) {
	var count int64
	db := db.DB()
	//improve count performance
	query := "select reltuples::bigint as estimate_rows from pg_class where relname = 'erc721_transfers'"
	if err := db.Raw(query).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetXrc721InfoByPage(skip, first uint64) ([]*XrcInfo, error) {
	var lists []*XrcInfo
	db := db.DB()
	query := "select t1.action_hash,t1.sender,t1.recipient,t1.token_id amount,t1.block_height,t1.timestamp,t1.contract_address from erc721_transfers t1 order by t1.id desc limit ? offset ?"
	if err := db.Raw(query, first, skip).Scan(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func GetXrc721ContractAddressCount() (int64, error) {
	var count int64
	db := db.DB()
	query := "SELECT COUNT(DISTINCT contract_address) FROM erc721_holders"
	if err := db.Raw(query).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetXrc721ContractAddress(skip, first uint64) ([]string, error) {
	var lists []string
	db := db.DB()
	query := "SELECT DISTINCT contract_address FROM erc721_holders limit ? offset ?"
	if err := db.Raw(query, first, skip).Scan(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func GetXrc721TokenHolderCountByTokenAddress(contractAddress string) (int64, error) {
	var count int64
	db := db.DB()
	query := "SELECT COUNT(1) FROM erc721_holders WHERE contract_address=?"
	if err := db.Raw(query, contractAddress).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetXrc721TokenHolder(contractAddress string, skip, first uint64) ([]string, error) {
	var lists []string
	db := db.DB()
	query := "SELECT holder FROM erc721_holders WHERE contract_address=? limit ? offset ?"
	if err := db.Raw(query, contractAddress, first, skip).Scan(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}
