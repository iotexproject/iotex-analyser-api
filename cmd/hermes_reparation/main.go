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
	_, err = db.Connect()
	if err != nil {
		log.Fatalf("failed to connect DB, %v", err)
	}
	rewardsOri, err := apiservice.GetHermeOrigin(epochStart, epochCount)
	if err != nil {
		log.Fatalf("failed to get hermes origin, %v", err)
	}
	rewardsFix, err := apiservice.GetHermeFixed(epochStart, epochCount)
	if err != nil {
		log.Fatalf("failed to get hermes fixed, %v", err)
	}
	//rewards := make(map[string]map[string]*big.Int)
	for delegate, rewardsMap := range rewardsFix {
		rewardMapOrigin := rewardsOri[delegate]
		for voter, reward := range rewardsMap {
			if _, ok := rewardMapOrigin[voter]; !ok {
				continue
			}
			rewardOrigin := rewardMapOrigin[voter]
			if rewardOrigin.Cmp(reward) == 0 {
				continue
			}
			diffAmount := new(big.Int).Sub(reward, rewardOrigin)
			log.Printf("%s,%s,%s,%s,%s\n", delegate, voter, reward.String(), rewardOrigin.String(), diffAmount.String())
		}
	}
}
