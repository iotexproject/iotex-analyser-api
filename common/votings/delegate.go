package votings

import (
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/db"
)

// GetOperatorAddress returns the operator addresss of the given delegate and epoch
func GetOperatorAddress(delegateName string, epoch uint64) (string, error) {
	blkHeight := common.GetEpochLastBlockHeight(epoch)
	query := "SELECT operator_address FROM candidate WHERE block_height<=? AND name = ? ORDER BY id desc LIMIT 1"
	var result string
	if err := db.DB().Raw(query, blkHeight, delegateName).Scan(&result).Error; err != nil {
		return "", err
	}
	return result, nil
}
