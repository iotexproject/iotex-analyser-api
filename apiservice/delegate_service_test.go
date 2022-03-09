package apiservice

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/shurcooL/graphql"
	"github.com/stretchr/testify/require"
)

type BucketInfo struct {
	VoterEthAddress   string `json:"voterEthAddress"`
	VoterIotexAddress string `json:"voterIotexAddress"`
	IsNative          bool   `json:"isNative"`
	Votes             string `json:"votes"`
	WeightedVotes     string `json:"weightedVotes"`
	RemainingDuration string `json:"remainingDuration"`
	StartTime         string `json:"startTime"`
	Decay             bool   `json:"decay"`
}

type BucketInfoList struct {
	EpochNumber int           `json:"epochNumber"`
	BucketInfo  []*BucketInfo `json:"bucketInfo"`
	Count       int           `json:"count"`
}

func TestDelegateGetBucketInfo(t *testing.T) {
	require := require.New(t)
	var startEpoch, epochCount uint64
	var delegateName string
	startEpoch = 24738
	epochCount = 1
	delegateName = "metanyx"

	start := time.Now()
	dist, err := getBucketInfoV1(startEpoch, epochCount, delegateName)
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", "getBucketInfoV1", elapsed)
	require.NoError(err)
	start = time.Now()
	dist2, err := getBucketInfoV2(startEpoch, epochCount, delegateName)
	elapsed = time.Since(start)
	fmt.Printf("%s took %s\n", "getBucketInfoV2", elapsed)
	require.NoError(err)
	require.Equal(len(dist), len(dist2))
	for _, d2 := range dist2[0].BucketInfo {
		found := false
		for _, d1 := range dist[0].BucketInfo {
			if d2.VoterIotexAddress == d1.VoterIotexAddress &&
				d2.VoterEthAddress == d1.VoterEthAddress &&
				d2.IsNative == d1.IsNative &&
				d2.Votes == d1.Votes &&
				d2.WeightedVotes == d1.WeightedVotes &&
				d2.RemainingDuration == d1.RemainingDuration &&
				d2.StartTime == d1.StartTime &&
				d2.Decay == d1.Decay {
				found = true
			}
		}
		require.True(found, d2.VoterIotexAddress)
	}
}

func TestDelegateBucketInfoV2(t *testing.T) {
	require := require.New(t)
	var startEpoch, epochCount uint64
	var delegateName string
	startEpoch = 24738
	epochCount = 1
	delegateName = "metanyx"

	start := time.Now()
	_, err := getBucketInfoV2(startEpoch, epochCount, delegateName)
	elapsed := time.Since(start)
	fmt.Printf("%s took  %s\n", "getBucketInfoV2", elapsed)
	require.NoError(err)
}

func getBucketInfoV1(startEpoch uint64, epochCount uint64, delegateName string) ([]*BucketInfoList, error) {
	gqlClient := graphql.NewClient("http://127.0.0.1:8089/query", nil)
	variables := map[string]interface{}{
		"startEpoch":   graphql.Int(startEpoch),
		"epochCount":   graphql.Int(epochCount),
		"delegateName": graphql.String(delegateName),
	}

	type BucketInfoOutput struct {
		Exist          bool              `json:"exist"`
		BucketInfoList []*BucketInfoList `json:"bucketInfoList"`
	}
	type query struct {
		Delegate struct {
			BucketInfo *BucketInfoOutput `json:"bucketInfo"`
		} `graphql:"delegate(startEpoch: $startEpoch, epochCount: $epochCount, delegateName: $delegateName)"`
	}
	var output query
	err := gqlClient.Query(context.Background(), &output, variables)
	if err != nil {
		return nil, err
	}
	if !output.Delegate.BucketInfo.Exist {
		return nil, errors.New("delegate info doesn't exist within the epoch range")
	}
	return output.Delegate.BucketInfo.BucketInfoList, nil
}

func getBucketInfoV2(startEpoch uint64, epochCount uint64, delegateName string) ([]*BucketInfoList, error) {
	gqlClient := graphql.NewClient("http://204.236.138.172:8889/graphql", nil)
	variables := map[string]interface{}{
		"startEpoch":   graphql.Int(startEpoch),
		"epochCount":   graphql.Int(epochCount),
		"delegateName": graphql.String(delegateName),
	}
	type query struct {
		GetBucketInfo struct {
			BucketInfo struct {
				Exist          bool `json:"exist"`
				BucketInfoList []*BucketInfoList
			}
		} `graphql:"GetBucketInfo(startEpoch: $startEpoch, epochCount: $epochCount, delegateName: $delegateName,pagination: { skip: 0, first: 8000 })"`
	}
	var output query
	err := gqlClient.Query(context.Background(), &output, variables)
	if err != nil {
		return nil, err
	}
	if !output.GetBucketInfo.BucketInfo.Exist {
		return nil, errors.New("delegate info doesn't exist within the epoch range")
	}
	return output.GetBucketInfo.BucketInfo.BucketInfoList, nil
}
