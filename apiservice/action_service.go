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
	resp := &api.ActionResponse{}
	resp.ActionList = &api.ActionList{
		Count:   0,
		Exist:   false,
		Actions: make([]*api.ActionInfo, 0),
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
	resp.ActionList.Exist = true
	resp.ActionList.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	actionInfoList, err := actions.GetBucketActionInfoByBuckets(bucketIDs, skip, first)
	if err != nil {
		return nil, err
	}
	for _, actionInfo := range actionInfoList {
		resp.ActionList.Actions = append(resp.ActionList.Actions, &api.ActionInfo{
			ActHash:   actionInfo.ActHash,
			BlkHash:   actionInfo.BlkHash,
			TimeStamp: actionInfo.TimeStamp,
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
	resp := &api.ActionResponse{}

	return resp, nil
}
