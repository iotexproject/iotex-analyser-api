package common

import (
	"crypto/tls"
	"time"

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

// func GetLatestNativeMintTime(height uint64) (time.Time, error) {
// 	db := db.DB()
// 	currentEpoch := GetEpochNum(height)
// 	lastEpochStartHeight := GetEpochHeight(currentEpoch - 1)
// 	getQuery := fmt.Sprintf(selectBlockHistory,
// 		blocks.BlockHistoryTableName, actions.ActionHistoryTableName)
// 	stmt, err := db.Prepare(getQuery)
// 	if err != nil {
// 		return time.Time{}, err
// 	}
// 	defer stmt.Close()
// 	var unixTimeStamp int64
// 	if err := stmt.QueryRow("putPollResult", height, lastEpochStartHeight).Scan(&unixTimeStamp); err != nil {
// 		return time.Time{}, err
// 	}
// 	log.S().Debugf("putpollresult block timestamp before height %d is %d\n", height, unixTimeStamp)
// 	//change unixTimeStamp to be a time.Time
// 	return time.Unix(unixTimeStamp, 0), nil
// }
