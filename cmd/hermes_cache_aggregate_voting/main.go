package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
)

var (
	epochStart uint64
	epochCount uint64
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configPath", "config.yaml", "configPath")
	flag.Uint64Var(&epochStart, "epochStart", 37021, "epochStart")
	flag.Uint64Var(&epochCount, "epochCount", 0, "epochCount")
}

type AggregateVoting struct {
	EpochNumber    uint64
	CandidateName  string
	VoterAddress   string
	NativeFlag     bool
	AggregateVotes string
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
	_, err = db.Connect()
	if err != nil {
		log.Fatalf("failed to connect DB, %v", err)
	}
	for epochNumber := epochStart; epochNumber < epochStart+epochCount; epochNumber++ {
		fileName := fmt.Sprintf("epoch_fix_%d.csv", epochNumber)
		if _, err := os.Stat(fileName); err == nil {
			continue
		}
		db := db.DB()
		rows, err := db.Raw("SELECT candidate_name,voter_address,aggregate_votes FROM hermes_aggregate_votings_fix WHERE epoch_number = ?", epochNumber).Rows()
		if err != nil {
			log.Fatalf("failed to get aggregate votings, %v", err)
		}
		defer rows.Close()
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("failed to open file, %v", err)
		}
		writer := csv.NewWriter(file)
		for rows.Next() {
			var vote AggregateVoting
			if err := rows.Scan(&vote.CandidateName, &vote.VoterAddress, &vote.AggregateVotes); err != nil {
				log.Fatalf("failed to scan aggregate votings, %v", err)
			}
			if err := writer.Write([]string{vote.CandidateName, vote.VoterAddress, vote.AggregateVotes}); err != nil {
				log.Fatalf("failed to write aggregate votings, %v", err)
			}
		}
		writer.Flush()
		file.Close()
		log.Printf("epoch %d done\n", epochNumber)
	}
}
