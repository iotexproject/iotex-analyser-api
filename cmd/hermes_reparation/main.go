package main

import (
	"flag"
	"io"
	"log"
	"math/big"
	"os"

	"github.com/iotexproject/iotex-analyser-api/apiservice"
	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-core/ioctl/util"
)

const (
	ConfigPath = "ConfigPath"
)

var (
	epochStart uint64
	epochCount uint64
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configPath", "config.yaml", "configPath")
	flag.Uint64Var(&epochStart, "epochStart", 37121, "epochStart")
	flag.Uint64Var(&epochCount, "epochCount", 0, "epochCount")
}

func main() {
	flag.Parse()
	if epochStart == 0 || epochCount == 0 {
		log.Fatalf("epochStart and epochCount must be set")
	}

	_, err := config.New(configPath)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	if config.Default.LogPath != "" {
		f, err := os.OpenFile(config.Default.LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		log.SetOutput(io.MultiWriter(f, os.Stdout))
	}
	log.Println("db connect")
	_, err = db.Connect()
	if err != nil {
		log.Fatalf("failed to connect DB, %v", err)
	}
	delegateMap, distributePlanMap, accountRewardsMap, err := apiservice.GetCommonData(epochStart, epochCount)
	if err != nil {
		log.Fatalf("failed to get common data, %v", err)
	}
	rewardsOri, err := apiservice.GetHermeOrigin(delegateMap, distributePlanMap, accountRewardsMap)
	if err != nil {
		log.Fatalf("failed to get hermes origin, %v", err)
	}
	rewardsFix, err := apiservice.GetHermeFixed(delegateMap, distributePlanMap, accountRewardsMap)
	if err != nil {
		log.Fatalf("failed to get hermes fixed, %v", err)
	}
	tmpFile := "hermes1.csv"
	f, err := os.OpenFile(tmpFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	f.Truncate(0)
	n := 0

	voterRewardsOri := getVoterReward(rewardsOri)
	voterRewardsFix := getVoterReward(rewardsFix)
	for voter, reward := range voterRewardsFix {
		rewardOrigin, ok := voterRewardsOri[voter]
		if !ok {
			continue
		}
		if rewardOrigin.Cmp(reward) == 0 {
			continue
		}
		diffAmount := new(big.Int).Sub(reward, rewardOrigin)
		if diffAmount.Cmp(big.NewInt(0)) < 0 {
			log.Printf("voter %s, reward %s, rewardOrigin %s, diffAmount %s", voter, reward.String(), rewardOrigin.String(), diffAmount.String())
		}
		diffPrice := util.RauToString(diffAmount, util.IotxDecimalNum)

		f.WriteString(voter + "," + reward.String() + "," + rewardOrigin.String() + "," + diffAmount.String() + "," + diffPrice + "\n")
		n++
	}
	// for delegate, rewardsMap := range rewardsFix {
	// 	rewardMapOrigin := rewardsOri[delegate]
	// 	for voter, reward := range rewardsMap {
	// 		if _, ok := rewardMapOrigin[voter]; !ok {
	// 			continue
	// 		}
	// 		rewardOrigin := rewardMapOrigin[voter]
	// 		if rewardOrigin.Cmp(reward) == 0 {
	// 			continue
	// 		}
	// 		diffAmount := new(big.Int).Sub(reward, rewardOrigin)
	// 		if diffAmount.Cmp(big.NewInt(0)) < 0 {
	// 			log.Printf("delegate %s, voter %s, reward %s, rewardOrigin %s, diffAmount %s", delegate, voter, reward.String(), rewardOrigin.String(), diffAmount.String())
	// 		}
	// 		f.WriteString(delegate + "," + voter + "," + reward.String() + "," + rewardOrigin.String() + "," + diffAmount.String() + "\n")
	// 		n++
	// 	}
	// }
	f.Close()
	log.Printf("total %d\n", n)
}

func getVoterReward(rewards apiservice.HermesDistributionReward) map[string]*big.Int {
	voterReward := make(map[string]*big.Int)
	for _, rewardsMap := range rewards {
		for voter, reward := range rewardsMap {
			if _, ok := voterReward[voter]; !ok {
				voterReward[voter] = big.NewInt(0)
			}
			voterReward[voter].Add(voterReward[voter], reward)
		}
	}
	return voterReward
}
