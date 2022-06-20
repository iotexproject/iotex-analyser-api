package apiservice

import (
	"context"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/rewards"
	"github.com/iotexproject/iotex-analyser-api/common/votings"
	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
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

	totalSupply, err := common.GetTotalSupply(height)
	if err != nil {
		return nil, err
	}
	resp.TotalSupply = totalSupply

	totalCirculatingSupply, err := common.GetTotalCirculatingSupply(height, totalSupply)
	if err != nil {
		return nil, err
	}
	resp.TotalCirculatingSupply = totalCirculatingSupply

	conn, err := iotex.NewDefaultGRPCConn(config.Default.RPC)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := iotexapi.NewAPIServiceClient(conn)

	availableRewards, err := rewards.GetAvailableRewards(ctx, client)
	if err != nil {
		return nil, err
	}
	totalCirculatingSupplyNoRewardPool, err := common.GetTotalCirculatingSupplyNoRewardPool(availableRewards.String(), totalCirculatingSupply)
	if err != nil {
		return nil, err
	}
	resp.TotalCirculatingSupplyNoRewardPool = totalCirculatingSupplyNoRewardPool

	meta, err := votings.GetVotingMeta()
	if err != nil {
		return nil, err
	}

	resp.VotingResultMeta = &api.VotingResultMeta{
		TotalCandidates:    meta.TotalCandidates,
		TotalWeightedVotes: meta.TotalWeightedVotes,
		VotedTokens:        meta.VotedTokens,
	}

	return resp, nil
}
