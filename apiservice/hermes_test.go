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

type hermesDistribution []struct {
	DelegateName       graphql.String
	RewardDistribution []struct {
		VoterIotexAddress graphql.String
		Amount            graphql.String
	}
	StakingIotexAddress graphql.String
	VoterCount          graphql.Int
	WaiveServiceFee     graphql.Boolean
	Refund              graphql.String
}

var (
	analyticsEndpoint string = "http://35.237.19.13:8080/query"
	// curl 'http://35.237.19.13:8080/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://35.237.19.13:8080' --data-binary '{"query":"\nquery {\n  hermes(startEpoch: 22420, epochCount: 2, \n    rewardAddress: \"io12mgttmfa2ffn9uqvn0yn37f4nz43d248l2ga85\", waiverThreshold: 100) {\n    hermesDistribution {\n      delegateName,\n      rewardDistribution{\n        voterEthAddress,\n        voterIotexAddress,\n        amount\n      },\n      stakingIotexAddress,\n      voterCount,\n      waiveServiceFee,\n      refund\n    }\n  }\n}"}' --compressed
)

func TestHermes(t *testing.T) {
	require := require.New(t)
	var startEpoch, epochCount uint64
	var rewardAddress string
	startEpoch = 22420
	epochCount = 2
	rewardAddress = "io12mgttmfa2ffn9uqvn0yn37f4nz43d248l2ga85"

	start := time.Now()
	dist, err := getHermesV1(startEpoch, epochCount, rewardAddress)
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", "getHermesV1", elapsed)
	require.NoError(err)
	start = time.Now()
	dist2, err := getHermesV2(startEpoch, epochCount, rewardAddress)
	elapsed = time.Since(start)
	fmt.Printf("%s took %s\n", "getHermesV2", elapsed)
	require.NoError(err)
	require.Equal(len(dist), len(dist2))
	for _, h1 := range dist {
		for _, h2 := range dist2 {
			if h2.DelegateName == h1.DelegateName {
				if string(h2.DelegateName) == "hackster" {
					continue
				}
				require.Equal(h2.Refund, h1.Refund, h2.DelegateName)
				require.Equal(h2.StakingIotexAddress, h1.StakingIotexAddress)
				require.Equal(h2.VoterCount, h1.VoterCount)
				require.Equal(h2.WaiveServiceFee, h1.WaiveServiceFee)
				require.Equal(h2.RewardDistribution, h1.RewardDistribution)
			}
		}
	}
	//sort.Slice(dist2, func(i, j int) bool { return dist2[i].DelegateName < dist2[j].DelegateName })
	//require.Equal(dist, dist2)
}

func getHermesV1(startEpoch uint64, epochCount uint64, rewardAddress string) (hermesDistribution, error) {
	waiverThreshold := 100
	gqlClient := graphql.NewClient(analyticsEndpoint, nil)
	variables := map[string]interface{}{
		"startEpoch":      graphql.Int(startEpoch),
		"epochCount":      graphql.Int(epochCount),
		"rewardAddress":   graphql.String(rewardAddress),
		"waiverThreshold": graphql.Int(waiverThreshold),
	}
	type query struct {
		Hermes struct {
			Exist              graphql.Boolean
			HermesDistribution hermesDistribution
		} `graphql:"hermes(startEpoch: $startEpoch, epochCount: $epochCount, rewardAddress: $rewardAddress, waiverThreshold: $waiverThreshold)"`
	}
	var output query
	err := gqlClient.Query(context.Background(), &output, variables)
	if err != nil {
		return nil, err
	}
	if !output.Hermes.Exist {
		return nil, errors.New("hermes info doesn't exist within the epoch range")
	}
	return output.Hermes.HermesDistribution, nil
}

func getHermesV2(startEpoch uint64, epochCount uint64, rewardAddress string) (hermesDistribution, error) {
	type query struct {
		Hermes struct {
			HermesDistribution hermesDistribution
		} `graphql:"Hermes(startEpoch: $startEpoch, epochCount: $epochCount, rewardAddress: $rewardAddress)"`
	}
	variables := map[string]interface{}{
		"startEpoch":    graphql.Int(startEpoch),
		"epochCount":    graphql.Int(epochCount),
		"rewardAddress": graphql.String(rewardAddress),
	}
	gqlClient := graphql.NewClient("http://127.0.0.1:8889/graphql", nil)
	var output query
	err := gqlClient.Query(context.Background(), &output, variables)
	if err != nil {
		return nil, err
	}
	return output.Hermes.HermesDistribution, nil
}
