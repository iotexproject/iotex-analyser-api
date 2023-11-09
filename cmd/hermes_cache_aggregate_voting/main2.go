package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type AggregateVoting struct {
	EpochNumber    uint64
	CandidateName  string
	VoterAddress   string
	NativeFlag     bool
	AggregateVotes string
}

func main() {
	epochNum := uint64(37021)
	fileName := fmt.Sprintf("epoch_%d.csv", epochNum)
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	votes := make([]AggregateVoting, 0)
	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1
	votesArr, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, row := range votesArr {
		votes = append(votes, AggregateVoting{
			EpochNumber:    epochNum,
			CandidateName:  row[0],
			VoterAddress:   row[1],
			AggregateVotes: row[2],
		})
		log.Printf("%+v", votes)
		break
	}
}
