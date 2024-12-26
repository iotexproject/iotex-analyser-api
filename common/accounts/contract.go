package accounts

import (
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-core/v2/ioctl/util"
)

func ContractIsExist(contractAddress string) (bool, uint64, error) {
	var exist struct {
		IsContract  bool
		BlockHeight uint64
	}
	db := db.DB()
	query := "select is_contract,block_height from account_meta where address=?"
	if err := db.Raw(query, contractAddress).Scan(&exist).Error; err != nil {
		return false, 0, err
	}
	return exist.IsContract, exist.BlockHeight, nil
}

func GetContractCallTimesAndAccumulatedGas(contractAddress string) (uint64, string, error) {
	var result struct {
		Count int
		Sum   string
	}
	query := "select count(*),sum(ba.gas_price*ba.gas_consumed) from action_execution ae left join block_action ba using(action_hash) where ae.contract=?"
	db := db.DB()
	if err := db.Raw(query, contractAddress).Scan(&result).Error; err != nil {
		return 0, "", err
	}
	gas, _ := util.StringToIOTX(result.Sum)
	return uint64(result.Count), gas, nil
}
