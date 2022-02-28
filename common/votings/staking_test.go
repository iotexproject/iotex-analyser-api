package votings

import (
	"testing"

	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/stretchr/testify/require"
)

func TestCandidate(t *testing.T) {
	require := require.New(t)
	_, err := initDB()
	require.NoError(err)
	epochNum := uint64(24738)
	res1, err := GetCandidateList(epochNum)
	require.NoError(err)
	chainClient := common.ChainClient("api.iotex.one:80")
	epochHeight := common.GetEpochHeight(epochNum)
	res2, err := GetAllStakingCandidates(chainClient, epochHeight)
	require.NoError(err)
	require.Equal(res1.Candidates, res2.Candidates)
}

func TestVoteBucketList(t *testing.T) {
	require := require.New(t)
	_, err := initDB()
	require.NoError(err)
	epochNum := uint64(24738)
	res1, err := GetVoteBucketList(epochNum)
	require.NoError(err)
	chainClient := common.ChainClient("api.iotex.one:80")
	epochHeight := common.GetEpochHeight(epochNum)
	res2, err := GetAllStakingBuckets(chainClient, epochHeight)
	require.NoError(err)
	require.Equal(res1.Buckets, res2.Buckets)
}
