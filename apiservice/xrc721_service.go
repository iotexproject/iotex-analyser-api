package apiservice

import (
	"context"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/actions"
)

type XRC721Service struct {
	api.UnimplementedXRC721ServiceServer
}

// XRC721ByAddress returns the XRC721 info of the address.
func (s *XRC721Service) XRC721ByAddress(ctx context.Context, req *api.XRC721ByAddressRequest) (*api.XRC721ByAddressResponse, error) {
	resp := &api.XRC721ByAddressResponse{
		Count:  0,
		Exist:  false,
		Xrc721: make([]*api.Xrc721Action, 0),
	}
	address, err := common.Address(req.GetAddress())
	if err != nil {
		return nil, err
	}
	count, err := actions.GetXrc721CountByAddress(*address)
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
	infos, err := actions.GetXrc721InfoByAddress(*address, skip, first)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		resp.Xrc721 = append(resp.Xrc721, &api.Xrc721Action{
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

// XRC721ByContractAddress returns Xrc721 actions given the Xrc721 contract address
func (s *XRC721Service) XRC721ByContractAddress(ctx context.Context, req *api.XRC721ByContractAddressRequest) (*api.XRC721ByContractAddressResponse, error) {
	resp := &api.XRC721ByContractAddressResponse{
		Count:  0,
		Exist:  false,
		Xrc721: make([]*api.Xrc721Action, 0),
	}
	address, err := common.Address(req.GetAddress())
	if err != nil {
		return nil, err
	}
	count, err := actions.GetXrc721CountByContractAddress(*address)
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
	infos, err := actions.GetXrc721InfoByContractAddress(*address, skip, first)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		resp.Xrc721 = append(resp.Xrc721, &api.Xrc721Action{
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

// XRC721ByPage returns Xrc721 actions by pagination
func (s *XRC721Service) XRC721ByPage(ctx context.Context, req *api.XRC721ByPageRequest) (*api.XRC721ByPageResponse, error) {
	resp := &api.XRC721ByPageResponse{
		Count:  0,
		Exist:  false,
		Xrc721: make([]*api.Xrc721Action, 0),
	}
	count, err := actions.GetXrc721Count()
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
	infos, err := actions.GetXrc721InfoByPage(skip, first)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		resp.Xrc721 = append(resp.Xrc721, &api.Xrc721Action{
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

// XRC721Addresses returns Xrc721 contract addresses
func (s *XRC721Service) XRC721Addresses(ctx context.Context, req *api.XRC721AddressesRequest) (*api.XRC721AddressesResponse, error) {
	resp := &api.XRC721AddressesResponse{
		Count:     0,
		Exist:     false,
		Addresses: make([]string, 0),
	}
	count, err := actions.GetXrc721ContractAddressCount()
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
	resp.Addresses, err = actions.GetXrc721ContractAddress(skip, first)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// XRC721TokenHolderAddresses returns Xrc721 token holder addresses given a Xrc721 contract address
func (s *XRC721Service) XRC721TokenHolderAddresses(ctx context.Context, req *api.XRC721TokenHolderAddressesRequest) (*api.XRC721TokenHolderAddressesResponse, error) {
	resp := &api.XRC721TokenHolderAddressesResponse{
		Count:     0,
		Addresses: make([]string, 0),
	}
	tokenAddress := req.GetTokenAddress()
	count, err := actions.GetXrc721TokenHolderCountByTokenAddress(tokenAddress)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return resp, nil
	}
	resp.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	resp.Addresses, err = actions.GetXrc721TokenHolder(tokenAddress, skip, first)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
