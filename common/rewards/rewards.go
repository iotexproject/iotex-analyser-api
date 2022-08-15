package rewards

import (
	"context"
	"math/big"

	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/pkg/errors"
)

func GetAvailableRewards(ctx context.Context, client iotexapi.APIServiceClient) (*big.Int, error) {
	request := &iotexapi.ReadStateRequest{
		ProtocolID: []byte("rewarding"),
		MethodName: []byte("AvailableBalance"),
	}

	response, err := client.ReadState(ctx, request)
	if err != nil {
		return nil, err
	}
	availableRewardInt, ok := new(big.Int).SetString(string(response.Data), 10)
	if !ok {
		err = errors.New("failed to format to big int:" + string(response.Data))
		return nil, err
	}
	return availableRewardInt, nil
}

func GetTotalRewards(ctx context.Context, client iotexapi.APIServiceClient) (*big.Int, error) {
	request := &iotexapi.ReadStateRequest{
		ProtocolID: []byte("rewarding"),
		MethodName: []byte("TotalBalance"),
	}

	response, err := client.ReadState(ctx, request)
	if err != nil {
		return nil, err
	}
	totalRewardRau, ok := new(big.Int).SetString(string(response.Data), 10)
	if !ok {
		err = errors.New("failed to format to big int:" + string(response.Data))
		return nil, err
	}
	return totalRewardRau, nil
}
