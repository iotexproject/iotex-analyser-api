package common

import "github.com/iotexproject/iotex-analyser-api/db"

type AddressFlow []struct {
	BlockHeight uint64
	Balance     string
}

func (af AddressFlow) ToMap() map[uint64]string {
	m := make(map[uint64]string)
	for _, v := range af {
		m[v.BlockHeight] = v.Balance
	}
	return m
}

func GetAddressFlow(addr string) (*AddressFlow, error) {
	db := db.DB()
	flow := &AddressFlow{}
	err := db.Select("block_height,sum(in_flow)-sum(out_flow) balance").Where("address = ?", addr).Table("account_income").Group("block_height").Order("block_height asc").Scan(&flow).Error
	if err != nil {
		return nil, err
	}
	return flow, nil
}
