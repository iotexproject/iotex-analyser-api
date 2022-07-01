package apiservice

import (
	"context"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/actions"
)

type XRC20Service struct {
	api.UnimplementedXRC20ServiceServer
}

// XRC20ByAddress returns the XRC20 info of the address.
func (s *XRC20Service) XRC20ByAddress(ctx context.Context, req *api.XRC20ByAddressRequest) (*api.XRC20ByAddressResponse, error) {
	resp := &api.XRC20ByAddressResponse{
		Count: 0,
		Exist: false,
		Xrc20: make([]*api.Xrc20Action, 0),
	}
	address, err := common.Address(req.GetAddress())
	if err != nil {
		return nil, err
	}
	count, err := actions.GetXrc20CountByAddress(*address)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return resp, nil
	}
	resp.Exist = true
	resp.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	infos, err := actions.GetXrc20InfoByAddress(*address, skip, first)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		resp.Xrc20 = append(resp.Xrc20, &api.Xrc20Action{
			ActHash:   info.ActionHash,
			Timestamp: uint64(info.Timestamp.Unix()),
			Sender:    info.Sender,
			Recipient: info.Recipient,
			Amount:    info.Amount,
			Contract:  info.ContractAddress,
		})
	}
	return resp, nil
}

// XRC20ByContractAddress returns Xrc20 actions given the Xrc20 contract address
func (s *XRC20Service) XRC20ByContractAddress(ctx context.Context, req *api.XRC20ByContractAddressRequest) (*api.XRC20ByContractAddressResponse, error) {
	resp := &api.XRC20ByContractAddressResponse{
		Count: 0,
		Exist: false,
		Xrc20: make([]*api.Xrc20Action, 0),
	}
	address, err := common.Address(req.GetAddress())
	if err != nil {
		return nil, err
	}
	count, err := actions.GetXrc20CountByContractAddress(*address)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return resp, nil
	}
	resp.Exist = true
	resp.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	infos, err := actions.GetXrc20InfoByContractAddress(*address, skip, first)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		resp.Xrc20 = append(resp.Xrc20, &api.Xrc20Action{
			ActHash:   info.ActionHash,
			Timestamp: uint64(info.Timestamp.Unix()),
			Sender:    info.Sender,
			Recipient: info.Recipient,
			Amount:    info.Amount,
			Contract:  info.ContractAddress,
		})
	}
	return resp, nil
}

// XRC20ByPage returns Xrc20 actions by pagination
func (s *XRC20Service) XRC20ByPage(ctx context.Context, req *api.XRC20ByPageRequest) (*api.XRC20ByPageResponse, error) {
	resp := &api.XRC20ByPageResponse{
		Count: 0,
		Exist: false,
		Xrc20: make([]*api.Xrc20Action, 0),
	}
	count, err := actions.GetXrc20Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return resp, nil
	}
	resp.Exist = true
	resp.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	infos, err := actions.GetXrc20InfoByPage(skip, first)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		resp.Xrc20 = append(resp.Xrc20, &api.Xrc20Action{
			ActHash:   info.ActionHash,
			Timestamp: uint64(info.Timestamp.Unix()),
			Sender:    info.Sender,
			Recipient: info.Recipient,
			Amount:    info.Amount,
			Contract:  info.ContractAddress,
		})
	}
	return resp, nil
}

// XRC20Addresses returns Xrc20 contract addresses
func (s *XRC20Service) XRC20Addresses(ctx context.Context, req *api.XRC20AddressesRequest) (*api.XRC20AddressesResponse, error) {
	resp := &api.XRC20AddressesResponse{
		Count:     0,
		Exist:     false,
		Addresses: make([]string, 0),
	}
	count, err := actions.GetXrc20ContractAddressCount()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return resp, nil
	}
	resp.Exist = true
	resp.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	resp.Addresses, err = actions.GetXrc20ContractAddress(skip, first)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// XRC20TokenHolderAddresses returns Xrc20 token holder addresses given a Xrc20 contract address
func (s *XRC20Service) XRC20TokenHolderAddresses(ctx context.Context, req *api.XRC20TokenHolderAddressesRequest) (*api.XRC20TokenHolderAddressesResponse, error) {
	resp := &api.XRC20TokenHolderAddressesResponse{
		Count:     0,
		Addresses: make([]string, 0),
	}
	tokenAddress := req.GetTokenAddress()
	count, err := actions.GetXrc20TokenHolderCountByTokenAddress(tokenAddress)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return resp, nil
	}
	resp.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	resp.Addresses, err = actions.GetXrc20TokenHolder(tokenAddress, skip, first)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
