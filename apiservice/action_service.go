package apiservice

import (
	"context"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/actions"
)

type ActionService struct {
	api.UnimplementedActionServiceServer
}

func (s *ActionService) ActionByVoter(ctx context.Context, req *api.ActionRequest) (*api.ActionResponse, error) {
	resp := &api.ActionResponse{
		Count:      0,
		Exist:      false,
		ActionList: make([]*api.ActionInfo, 0),
	}
	address, err := common.Address(req.GetAddress())
	if err != nil {
		return nil, err
	}
	bucketIDs, err := actions.GetBucketIDsByVoter(*address)
	if err != nil {
		return nil, err
	}
	if len(bucketIDs) == 0 {
		return resp, nil
	}
	count, err := actions.GetBucketActionCountByBuckets(bucketIDs)
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
	actionInfoList, err := actions.GetBucketActionInfoByBuckets(bucketIDs, skip, first)
	if err != nil {
		return nil, err
	}
	for _, actionInfo := range actionInfoList {
		resp.ActionList = append(resp.ActionList, &api.ActionInfo{
			ActHash:   actionInfo.ActHash,
			BlkHash:   actionInfo.BlkHash,
			Timestamp: actionInfo.Timestamp,
			ActType:   actionInfo.ActType,
			Sender:    actionInfo.Sender,
			Recipient: actionInfo.Recipient,
			Amount:    actionInfo.Amount,
			GasFee:    actionInfo.GasFee,
			BlkHeight: actionInfo.BlkHeight,
		})
	}
	return resp, nil
}

func (s *ActionService) GetXrc20ByAddress(ctx context.Context, req *api.ActionRequest) (*api.ActionResponse, error) {
	resp := &api.ActionResponse{
		Count:   0,
		Exist:   false,
		XrcList: make([]*api.XrcInfo, 0),
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
		resp.XrcList = append(resp.XrcList, &api.XrcInfo{
			ActHash:   info.ActionHash,
			BlkHeight: info.BlockHeight,
			Timestamp: uint64(info.Timestamp.Unix()),
			From:      info.Sender,
			To:        info.Recipient,
			Quantity:  info.Amount,
			Contract:  info.ContractAddress,
		})
	}
	return resp, nil
}

//ActionByDates finds actions by dates
func (s *ActionService) ActionByDates(ctx context.Context, req *api.ActionByDatesRequest) (*api.ActionByDatesResponse, error) {
	resp := &api.ActionByDatesResponse{}
	startDate := req.GetStartDate()
	endDate := req.GetEndDate()
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())

	count, err := actions.GetActionCountByDates(startDate, endDate)
	if err != nil {
		return nil, err
	}
	resp.Count = uint64(count)
	if count == 0 {
		return resp, nil
	}
	resp.Exist = true
	actionInfoList, err := actions.GetActionInfoByDates(startDate, endDate, skip, first)
	if err != nil {
		return nil, err
	}
	for _, actionInfo := range actionInfoList {
		resp.Actions = append(resp.Actions, &api.ActionInfo{
			ActHash:   actionInfo.ActHash,
			BlkHash:   actionInfo.BlkHash,
			Timestamp: actionInfo.Timestamp,
			ActType:   actionInfo.ActType,
			Sender:    actionInfo.Sender,
			Recipient: actionInfo.Recipient,
			Amount:    actionInfo.Amount,
			GasFee:    actionInfo.GasFee,
			BlkHeight: actionInfo.BlkHeight,
		})
	}
	return resp, nil
}

//ActionByHash finds actions by hash
func (s *ActionService) ActionByHash(ctx context.Context, req *api.ActionByHashRequest) (*api.ActionByHashResponse, error) {
	resp := &api.ActionByHashResponse{}
	actHash := req.GetActHash()

	actionInfo, err := actions.GetActionInfoByHash(actHash)
	if err != nil {
		return nil, err
	}
	resp.ActionInfo = &api.ActionInfo{
		ActHash:   actionInfo.ActHash,
		BlkHash:   actionInfo.BlkHash,
		Timestamp: actionInfo.Timestamp,
		ActType:   actionInfo.ActType,
		Sender:    actionInfo.Sender,
		Recipient: actionInfo.Recipient,
		Amount:    actionInfo.Amount,
		GasFee:    actionInfo.GasFee,
		BlkHeight: actionInfo.BlkHeight,
	}
	brt, err := actions.GetBlockReceiptTransactionByHash(actHash)
	if err != nil {
		return nil, err
	}
	for _, receipt := range brt {
		resp.EvmTransfers = append(resp.EvmTransfers, &api.ActionByHashResponse_EvmTransfers{
			Sender:    receipt.Sender,
			Recipient: receipt.Recipient,
			Amount:    receipt.Amount,
		})
	}
	return resp, nil
}

// ActionByAddress finds actions by address
func (s *ActionService) ActionByAddress(ctx context.Context, req *api.ActionByAddressRequest) (*api.ActionByAddressResponse, error) {
	resp := &api.ActionByAddressResponse{
		Count:   0,
		Exist:   false,
		Actions: make([]*api.ActionInfo, 0),
	}
	address, err := common.Address(req.GetAddress())
	if err != nil {
		return nil, err
	}

	count, err := actions.GetActionCountByAddress(ctx, *address)
	if err != nil {
		return nil, err
	}
	resp.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	actionInfoList, err := actions.GetActionInfoByAddress(ctx, *address, skip, first)
	if err != nil {
		return nil, err
	}
	for _, actionInfo := range actionInfoList {
		resp.Actions = append(resp.Actions, &api.ActionInfo{
			ActHash:   actionInfo.ActHash,
			BlkHash:   actionInfo.BlkHash,
			Timestamp: actionInfo.Timestamp,
			ActType:   actionInfo.ActType,
			Sender:    actionInfo.Sender,
			Recipient: actionInfo.Recipient,
			Amount:    actionInfo.Amount,
			GasFee:    actionInfo.GasFee,
			BlkHeight: actionInfo.BlkHeight,
		})
	}
	return resp, nil
}

// ActionByType finds actions by type
func (s *ActionService) ActionByType(ctx context.Context, req *api.ActionByTypeRequest) (*api.ActionByTypeResponse, error) {
	resp := &api.ActionByTypeResponse{
		Count:   0,
		Exist:   false,
		Actions: make([]*api.ActionInfo, 0),
	}
	typ := req.GetType()

	count, err := actions.GetActionCountByType(ctx, typ)
	if err != nil {
		return nil, err
	}
	resp.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	actionInfoList, err := actions.GetActionInfoByType(ctx, typ, skip, first)
	if err != nil {
		return nil, err
	}
	for _, actionInfo := range actionInfoList {
		resp.Actions = append(resp.Actions, &api.ActionInfo{
			ActHash:   actionInfo.ActHash,
			BlkHash:   actionInfo.BlkHash,
			Timestamp: actionInfo.Timestamp,
			ActType:   actionInfo.ActType,
			Sender:    actionInfo.Sender,
			Recipient: actionInfo.Recipient,
			Amount:    actionInfo.Amount,
			GasFee:    actionInfo.GasFee,
			BlkHeight: actionInfo.BlkHeight,
		})
	}
	return resp, nil
}

//EvmTransfersByAddress finds EVM transfers by address
func (s *ActionService) EvmTransfersByAddress(ctx context.Context, req *api.EvmTransfersByAddressRequest) (*api.EvmTransfersByAddressResponse, error) {
	resp := &api.EvmTransfersByAddressResponse{
		Count:        0,
		Exist:        false,
		EvmTransfers: make([]*api.EvmTransfersByAddressResponse_EvmTransfer, 0),
	}
	address, err := common.Address(req.GetAddress())
	if err != nil {
		return nil, err
	}
	count, err := actions.GetEvmTransferCount(ctx, *address)
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
	actionInfos, err := actions.GetEvmTransferInfoByAddress(ctx, *address, skip, first)
	if err != nil {
		return nil, err
	}
	for _, info := range actionInfos {
		resp.EvmTransfers = append(resp.EvmTransfers, &api.EvmTransfersByAddressResponse_EvmTransfer{
			ActHash:   info.ActHash,
			BlkHash:   info.BlkHash,
			Timestamp: info.Timestamp,
			Sender:    info.Sender,
			Recipient: info.Recipient,
			Amount:    info.Amount,
			BlkHeight: info.BlkHeight,
		})
	}
	return resp, nil
}
