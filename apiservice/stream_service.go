package apiservice

import (
	"math/big"
	"sort"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
)

// StreamService is the service to handle stream related requests
type StreamService struct {
	api.UnimplementedStreamServiceServer
}

// Supply returns the supply info, including total supply, circulating supply and exact circulating supply
// the API is streaming, so the client can get the supply info in real time
func (s *StreamService) Supply(req *api.SupplyRequest, res api.StreamService_SupplyServer) error {
	height := req.GetStartHeight()

	var totalCirculatingSupply string

	zeroAddressFlow, err := common.GetAddressFlow(address.ZeroAddress)
	if err != nil {
		return err
	}
	lockAddr := "io1uqhmnttmv0pg8prugxxn7d8ex9angrvfjfthxa"
	lockAddressFlow, err := common.GetAddressFlow(lockAddr)
	if err != nil {
		return err
	}
	zeroAddressFlowMap := zeroAddressFlow.ToMap()
	lockAddressFlowMap := lockAddressFlow.ToMap()
	type addrFlow struct {
		Height          uint64
		ZeroAddressFlow string
		LockAddressFlow string
	}
	var addrFlows []addrFlow
	for k, v := range zeroAddressFlowMap {
		if vv, ok := lockAddressFlowMap[k]; ok {
			addrFlows = append(addrFlows, addrFlow{
				Height:          k,
				ZeroAddressFlow: v,
				LockAddressFlow: vv,
			})
			delete(lockAddressFlowMap, k)
		} else {
			addrFlows = append(addrFlows, addrFlow{
				Height:          k,
				ZeroAddressFlow: v,
				LockAddressFlow: "0",
			})
		}
	}
	for k, v := range lockAddressFlowMap {
		addrFlows = append(addrFlows, addrFlow{
			Height:          k,
			ZeroAddressFlow: "0",
			LockAddressFlow: v,
		})
	}
	sort.Slice(addrFlows, func(i, j int) bool {
		return addrFlows[i].Height < addrFlows[j].Height
	})
	zeroAddressBalanceTotal := new(big.Int)
	lockAddressBalanceTotal := new(big.Int)

	type supplyTy struct {
		Height                 uint64
		TotalSupply            *big.Int
		TotalCirculatingSupply *big.Int
	}
	var supplyBalances []supplyTy
	for _, v := range addrFlows {
		zeroAddressBalance, _ := new(big.Int).SetString(v.ZeroAddressFlow, 10)
		zeroAddressBalanceTotal.Add(zeroAddressBalanceTotal, zeroAddressBalance)
		lockAddressBalance, _ := new(big.Int).SetString(v.LockAddressFlow, 10)
		lockAddressBalanceTotal.Add(lockAddressBalanceTotal, lockAddressBalance)
		totalSupply := new(big.Int).Sub(new(big.Int).Sub(new(big.Int).Sub(common.TotalBalanceInt, zeroAddressBalanceTotal), common.Nsv1BalanceInt), common.BnfxBalanceInt)
		totalCirculatingSupply := new(big.Int).Sub(totalSupply, lockAddressBalanceTotal)
		supplyBalances = append(supplyBalances, supplyTy{
			Height:                 v.Height,
			TotalSupply:            totalSupply,
			TotalCirculatingSupply: totalCirculatingSupply,
		})
	}
	getTotalSupply := func(height uint64) (*big.Int, *big.Int) {
		var totalSupply, circulatingSupply *big.Int
		for _, v := range supplyBalances {
			if v.Height > height {
				break
			}
			totalSupply = v.TotalSupply
			circulatingSupply = v.TotalCirculatingSupply
		}
		return totalSupply, circulatingSupply
	}
	for height <= req.GetEndHeight() {
		totalSupply, err := common.GetTotalSupply(height)
		if err != nil {
			return err
		}
		totalSupplyBig, totalCirculatingSupplyBig := getTotalSupply(height)
		totalCirculatingSupply, err = common.GetTotalCirculatingSupply(height, totalSupply)
		if err != nil {
			return err
		}
		if err := res.Send(&api.SupplyResponse{
			Height:            height,
			TotalSupply:       totalSupply,
			CirculatingSupply: totalCirculatingSupply,
		}); err != nil {
			return err
		}
		if err := res.Send(&api.SupplyResponse{
			Height:            height,
			TotalSupply:       totalSupplyBig.String(),
			CirculatingSupply: totalCirculatingSupplyBig.String(),
		}); err != nil {
			return err
		}
		height++
	}
	return nil
}
