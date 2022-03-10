package apiservice

import (
	"context"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
)

type ChainService struct {
	api.UnimplementedChainServiceServer
}

func (s *ChainService) Chain(ctx context.Context, req *api.ChainRequest) (*api.ChainResponse, error) {
	resp := &api.ChainResponse{}

	epoch, height, err := common.GetCurrentEpochAndHeight()
	if err != nil {
		return nil, err
	}
	resp.MostRecentEpoch = epoch
	resp.MostRecentBlockHeight = height

	return resp, nil
}
