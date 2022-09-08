package apiservice

import (
	"context"
	"math/big"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/rewards"
	"github.com/iotexproject/iotex-analyser-api/common/votings"
	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/pkg/errors"
)

// ChainService is the service to handle chain related requests
type ChainService struct {
	api.UnimplementedChainServiceServer
}

// Chain returns the chain info
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

	exactCirculatingSupply, err := common.GetExactCirculatingSupply(height, totalSupply)
	if err != nil {
		return nil, err
	}
	resp.ExactCirculatingSupply = exactCirculatingSupply

	conn, err := common.NewDefaultGRPCConn(config.Default.RPC)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := iotexapi.NewAPIServiceClient(conn)

	availableRewards, err := rewards.GetAvailableRewards(ctx, client)
	if err != nil {
		return nil, err
	}
	totalRewards, err := rewards.GetTotalRewards(ctx, client)
	if err != nil {
		return nil, err
	}
	resp.Rewards = &api.ChainResponse_Rewards{
		TotalAvailable: availableRewards.String(),
		TotalBalance:   totalRewards.String(),
		TotalUnclaimed: new(big.Int).Sub(totalRewards, availableRewards).String(),
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

// MostRecentTPS gives the most recent TPS
func (s *ChainService) MostRecentTPS(ctx context.Context, req *api.MostRecentTPSRequest) (*api.MostRecentTPSResponse, error) {
	resp := &api.MostRecentTPSResponse{}

	_, height, err := common.GetCurrentEpochAndHeight()
	if err != nil {
		return nil, err
	}

	blockWindow := req.GetBlockWindow()
	if height < blockWindow {
		blockWindow = height
	}

	start := height - blockWindow + 1
	end := height
	db := db.DB()
	query := "select (select timestamp from block where block_height=?) start_time,(select timestamp from block where block_height=?) end_time,sum(num_actions) num_actions from block where block_height>=? and block_height<=?"
	var result struct {
		StartTime  uint64
		EndTime    uint64
		NumActions uint64
	}
	err = db.Raw(query, start, end, start, end).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	resp.MostRecentTPS = float64(result.NumActions) / float64(result.EndTime-result.StartTime)
	return resp, nil
}

// NumberOfActions gives the number of actions within a epoch frame
func (s *ChainService) NumberOfActions(ctx context.Context, req *api.NumberOfActionsRequest) (*api.NumberOfActionsResponse, error) {
	resp := &api.NumberOfActionsResponse{}

	currentEpoch, _, err := common.GetCurrentEpochAndHeight()
	if err != nil {
		return nil, err
	}

	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	if startEpoch > currentEpoch {
		return resp, nil
	}
	endEpoch := startEpoch + epochCount - 1
	db := db.DB()
	query := "select sum(num_actions) num_actions from block b right join block_meta bm on bm.block_height=b.block_height where bm.epoch_num>=? and bm.epoch_num<=?"
	var result struct {
		NumActions uint64
	}
	err = db.Raw(query, startEpoch, endEpoch).Scan(&result).Error
	if err != nil {
		return nil, err
	}

	resp.Exist = true
	resp.Count = result.NumActions
	return resp, nil
}

// TotalTransferredTokens gives the amount of tokens transferred within a time frame
func (s *ChainService) TotalTransferredTokens(ctx context.Context, req *api.TotalTransferredTokensRequest) (*api.TotalTransferredTokensResponse, error) {
	db := db.DB()
	resp := &api.TotalTransferredTokensResponse{}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	endEpoch := startEpoch + epochCount - 1
	startHeight := common.GetEpochHeight(startEpoch)
	endHeight := common.GetEpochLastBlockHeight(endEpoch)
	query := "select SUM(amount) from block_receipt_transactions where block_height>=? and block_height<=?"

	var result string
	if err := db.WithContext(ctx).Raw(query, startHeight, endHeight).Scan(&result).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get total number of holders")
	}

	resp.TotalTransferredTokens = result
	return resp, nil
}

func (s *ChainService) ChartSync(ctx context.Context, req *api.ChartSyncRequest) (*api.ChartSyncResponse, error) {
	resp := &api.ChartSyncResponse{}
	db := db.DB()
	query := "select max(block_height),date(timestamp) from block GROUP BY date(timestamp)"
	rows, err := db.WithContext(ctx).Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var blkNo uint64
	var date string
	for rows.Next() {
		if err := rows.Scan(&blkNo, &date); err != nil {
			return nil, err
		}
		resp.States = append(resp.States, &api.ChartSyncResponse_State{
			BlockNumber:   blkNo,
			Time:          date,
			ServerVersion: getHardForkVersion(blkNo),
		})
	}
	return resp, nil
}

//https://iotexscan.io/hard-fork-history
func getHardForkVersion(blk uint64) string {
	vers := []struct {
		h uint64
		v string
	}{
		{432001, "0.6.2"},
		{864001, "0.7.2"},
		{1512001, "0.8.3"},
		{1641601, "0.9.0"},
		{1816201, "0.10.0"},
		{5165641, "1.0.0"},
		{6544441, "1.1.0"},
		{11267641, "1.2.0"},
		{12289321, "1.3.0"},
		{13685401, "1.4.0"},
		{13816441, "1.5.0"},
		{13979161, "1.6.0"},
		{16509241, "1.7.0"},
		{17662681, "1.8.0"},
	}
	for i := len(vers) - 1; i >= 0; i-- {
		if blk >= vers[i].h {
			return vers[i].v
		}
	}
	return "0.6.0"
}
