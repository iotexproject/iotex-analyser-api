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

func (s *ActionService) GetActionByVoter(ctx context.Context, req *api.ActionRequest) (*api.ActionResponse, error) {
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

func (s *ActionService) GetEvmTransfersByAddress(ctx context.Context, req *api.ActionRequest) (*api.ActionResponse, error) {
	resp := &api.ActionResponse{
		Count:           0,
		Exist:           false,
		EvmTransferList: make([]*api.EvmTransferInfo, 0),
	}
	address, err := common.Address(req.GetAddress())
	if err != nil {
		return nil, err
	}
	count, err := actions.GetEvmTransferCount(*address)
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
	actionInfos, err := actions.GetEvmTransferInfoByAddress(*address, skip, first)
	if err != nil {
		return nil, err
	}
	for _, info := range actionInfos {
		resp.EvmTransferList = append(resp.EvmTransferList, &api.EvmTransferInfo{
			ActHash:   info.ActHash,
			BlkHash:   info.BlkHash,
			Timestamp: info.Timestamp,
			From:      info.Sender,
			To:        info.Recipient,
			Quantity:  info.Amount,
			BlkHeight: info.BlkHeight,
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
