package common

import (
	"database/sql"
	"math/big"
	"strings"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/pkg/errors"
)

const (
	lockAddresses = "io1uqhmnttmv0pg8prugxxn7d8ex9angrvfjfthxa" // Separate multiple addresses with ","
	totalBalance  = "12700000000000000000000000000"             // 10B + 2.7B (due to Postmortem 1)
	nsv1Balance   = "262281303940000000000000000"
	bnfxBalance   = "3414253030000000000000000"
)

func AccountBalanceByHeight(height uint64, addresses []string) ([]*big.Int, error) {
	result := make([]*big.Int, 0)
	db := db.DB()
	for _, addr := range addresses {
		if addr[:2] == "0x" || addr[:2] == "0X" {
			add, err := address.FromHex(addr)
			if err != nil {
				return nil, err
			}

			addr = add.String()
		}

		var amount sql.NullString
		query := "SELECT sum(in_flow)-sum(out_flow) from account_income WHERE block_height<=? and address=?"
		err := db.Raw(query, height, addr).Scan(&amount).Error
		if err != nil {
			return nil, err
		}
		balance, ok := big.NewInt(0).SetString(amount.String, 10)
		if !ok {
			balance = big.NewInt(0)
		}
		result = append(result, balance)
	}
	return result, nil
}

func GetTotalSupply(height uint64) (string, error) {
	// get zero address balance.
	zeroAddressBalance, err := AccountBalanceByHeight(height, []string{address.ZeroAddress})
	if err != nil {
		return "", err
	}

	// Convert string format to big.Int format
	totalBalanceInt, _ := new(big.Int).SetString(totalBalance, 10)
	nsv1BalanceInt, _ := new(big.Int).SetString(nsv1Balance, 10)
	bnfxBalanceInt, _ := new(big.Int).SetString(bnfxBalance, 10)

	// Compute 10B + 2.7B (due to Postmortem 1) - Balance(all zero address) - Balance(nsv1) - Balance(bnfx)
	return new(big.Int).Sub(new(big.Int).Sub(new(big.Int).Sub(totalBalanceInt, zeroAddressBalance[0]), nsv1BalanceInt), bnfxBalanceInt).String(), nil
}

func GetTotalCirculatingSupply(height uint64, totalSupply string) (string, error) {
	locked, err := AccountBalanceByHeight(height, strings.Split(lockAddresses, ","))
	if err != nil {
		return "", err
	}

	lockedBig := new(big.Int)
	for _, v := range locked {
		lockedBig = new(big.Int).Add(lockedBig, v)
	}
	totalSupplyBig, ok := new(big.Int).SetString(totalSupply, 10)
	if !ok {
		return "", errors.New("failed to format to big int:" + totalSupply)
	}

	return new(big.Int).Sub(totalSupplyBig, lockedBig).String(), nil
}

func GetTotalCirculatingSupplyNoRewardPool(availableRewards, totalCirculatingSupply string) (string, error) {
	availableRewardsBig, ok := new(big.Int).SetString(availableRewards, 10)
	if !ok {
		return "", errors.New("failed to format to big int:" + availableRewards)
	}

	totalCirculatingSupplyBig, ok := new(big.Int).SetString(totalCirculatingSupply, 10)
	if !ok {
		return "", errors.New("failed to format to big int:" + totalCirculatingSupply)
	}

	return new(big.Int).Sub(totalCirculatingSupplyBig, availableRewardsBig).String(), nil
}
