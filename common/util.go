package common

import (
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"google.golang.org/grpc"
)

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
