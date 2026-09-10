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
	"google.golang.org/grpc/credentials/insecure"
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

// ArchiveChainClient returns a client for the archive endpoint, for reads that
// address a past height. DefaultChainClient points at the light-node pool,
// which keeps only the latest 256 blocks of state and answers any older height
// with "history is pruned". Falls back to the regular endpoint when no archive
// endpoint is configured, so a network without one keeps working (and keeps
// failing the same way it did before) rather than erroring at startup.
func ArchiveChainClient() (iotexapi.APIServiceClient, error) {
	endpoint := config.Default.ArchiveRPC
	if endpoint == "" {
		endpoint = config.Default.RPC
	}
	conn, err := NewDefaultGRPCConn(endpoint)
	if err != nil {
		return nil, err
	}
	return iotexapi.NewAPIServiceClient(conn), nil
}

// NewDefaultGRPCConn creates a default grpc connection, with retry and — unless
// config.RPCInsecure is set — TLS.
func NewDefaultGRPCConn(endpoint string) (*grpc.ClientConn, error) {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(100 * time.Second)),
		grpc_retry.WithMax(3),
	}
	creds := credentials.NewTLS(&tls.Config{})
	if config.Default.RPCInsecure {
		creds = insecure.NewCredentials()
	}
	return grpc.Dial(endpoint,
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(opts...)),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
		grpc.WithTransportCredentials(creds))
}
