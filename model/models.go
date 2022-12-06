package model

type StakingBucket struct {
	ID               uint64
	BucketID         uint64
	BlockHeight      uint64
	CreateTime       int64
	StakeStartTime   int64
	UnstakeStartTime int64
	StakedAmount     string
	VotingPower      string
	OwnerAddress     string
	Candidate        string
	Amount           string
	ActType          string
	Sender           string
	ActionHash       string
	Timestamp        int64
	AutoStake        bool
	Duration         uint32
}
