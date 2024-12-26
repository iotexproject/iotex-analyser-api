package common

import (
	"math/big"
	"strings"
	"time"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-core/v2/ioctl/util"
	"github.com/pkg/errors"
)

const (
	lockAddresses = "io1uqhmnttmv0pg8prugxxn7d8ex9angrvfjfthxa" // Separate multiple addresses with ","
	totalBalance  = "12700000000000000000000000000"             // 10B + 2.7B (due to Postmortem 1)
	nsv1Balance   = "262281303940000000000000000"
	bnfxBalance   = "3414253030000000000000000"
)

var (
	TotalBalanceInt, _ = new(big.Int).SetString(totalBalance, 10)
	Nsv1BalanceInt, _  = new(big.Int).SetString(nsv1Balance, 10)
	BnfxBalanceInt, _  = new(big.Int).SetString(bnfxBalance, 10)
)

func AccountBalanceByHeight(height uint64, addresses []string) ([]*big.Int, error) {
	result := make([]*big.Int, 0)
	db := db.DB()
	addres := make([]string, 0)
	for _, addr := range addresses {
		if addr[:2] == "0x" || addr[:2] == "0X" {
			add, err := address.FromHex(addr)
			if err != nil {
				return nil, err
			}
			addr = add.String()

		}
		addres = append(addres, addr)
	}
	var res []struct {
		Address string
		Balance string
	}
	err := db.Select("address,sum(in_flow)-sum(out_flow) balance").Where("block_height<=? and address in ?", height, addres).Table("account_income").Group("address").Scan(&res).Error
	if err != nil {
		return nil, err
	}
	results := make(map[string]*big.Int)
	for _, v := range res {
		balance, ok := new(big.Int).SetString(v.Balance, 10)
		if !ok {
			balance = big.NewInt(0)
		}
		results[v.Address] = balance
	}

	for _, addr := range addresses {
		if addr[:2] == "0x" || addr[:2] == "0X" {
			add, err := address.FromHex(addr)
			if err != nil {
				return nil, err
			}
			addr = add.String()

		}
		addr = strings.ToLower(addr)
		if results[addr] == nil {
			result = append(result, big.NewInt(0))
		} else {
			result = append(result, results[addr])
		}
	}

	return result, nil
}

func GetTotalSupply(height uint64) (string, error) {
	// get zero address balance.
	zeroAddressBalance, err := AccountBalanceByHeight(height, []string{address.ZeroAddress})
	if err != nil {
		return "", err
	}

	// Compute 10B + 2.7B (due to Postmortem 1) - Balance(all zero address) - Balance(nsv1) - Balance(bnfx)
	return new(big.Int).Sub(new(big.Int).Sub(new(big.Int).Sub(TotalBalanceInt, zeroAddressBalance[0]), Nsv1BalanceInt), BnfxBalanceInt).String(), nil
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

func GetExactCirculatingSupply(height uint64, totalCirculatingSupply string) (string, error) {
	totalCirculatingSupplyBig, ok := new(big.Int).SetString(totalCirculatingSupply, 10)
	if !ok {
		return "", errors.New("failed to format to big int:" + totalCirculatingSupply)
	}
	scheduledSupply := []struct {
		Date   string
		Supply string
	}{
		{"2023-03-31", "10000000000"},
		{"2023-02-28", "9490000000"},
		{"2023-01-31", "9475000000"},
		{"2022-12-31", "9460000000"},
		{"2022-11-30", "9445000000"},
		{"2022-10-31", "9430000000"},
		{"2022-09-30", "9415000000"},
		{"2022-08-31", "9400000000"},
		{"2022-07-31", "9285000000"},
	}
	exactSupply := new(big.Int)
	totalAmount, err := util.StringToRau("9490000000", util.IotxDecimalNum)
	if err != nil {
		return "", err
	}
	for _, v := range scheduledSupply {
		t, err := time.Parse("2006-01-02", v.Date)
		if err != nil {
			return "", err
		}
		if time.Now().UTC().After(t.UTC()) {
			amount, err := util.StringToRau(v.Supply, util.IotxDecimalNum)
			if err != nil {
				return "", err
			}
			burned := new(big.Int).Sub(totalAmount, totalCirculatingSupplyBig)
			exactSupply = new(big.Int).Sub(amount, burned)
			break
		}
	}

	return exactSupply.String(), nil
}
