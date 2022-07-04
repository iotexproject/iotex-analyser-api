package actions

import (
	"math/big"
	"os"
	"testing"

	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	_, err := config.New(os.Getenv("ConfigPath"))
	if err != nil {
		return nil, err
	}
	return db.Connect()
}

func TestBucket(t *testing.T) {
	require := require.New(t)
	_, err := initDB()
	require.NoError(err)
	buckets, err := GetVoteBucketList(20000)
	require.NoError(err)
	totalVoted := big.NewInt(0)
	for _, bucket := range buckets.GetBuckets() {
		amount, _ := big.NewInt(0).SetString(bucket.StakedAmount, 10)
		totalVoted.Add(totalVoted, amount)
	}
	require.Equal(totalVoted.String(), "1000000000000000000")
}
