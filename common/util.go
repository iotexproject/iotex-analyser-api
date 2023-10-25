package common

import (
	"crypto/tls"
	"time"

	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// GetCurrentEpochAndHeight returns current epoch and blockHeight
func GetCurrentEpochAndHeight() (uint64, uint64, error) {
	var ret struct {
		BlockHeight uint64
		EpochNum    uint64
	}
	db := db.DB()
	if err := db.Table("block_meta").Select("block_height,epoch_num").Last(&ret).Error; err != nil {
		return 0, 0, err
	}
	return ret.EpochNum, ret.BlockHeight, nil
}

func ChainClient(endpoint string) iotexapi.APIServiceClient {
	opt := grpc.WithInsecure()
	conn, err := grpc.Dial(endpoint, opt)
	if err != nil {
		panic(err)
	}

	return iotexapi.NewAPIServiceClient(conn)
}

func DefaultChainClient() (iotexapi.APIServiceClient, error) {
	conn, err := NewDefaultGRPCConn(config.Default.RPC)
	if err != nil {
		return nil, err
	}
	return iotexapi.NewAPIServiceClient(conn), nil
}

// NewDefaultGRPCConn creates a default grpc connection. With tls and retry.
func NewDefaultGRPCConn(endpoint string) (*grpc.ClientConn, error) {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(100 * time.Second)),
		grpc_retry.WithMax(3),
	}
	return grpc.Dial(endpoint,
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(opts...)),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
}
