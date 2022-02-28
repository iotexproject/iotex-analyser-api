package votings

type (
	VotingInfo struct {
		EpochNumber       uint64
		VoterAddress      string
		IsNative          bool
		Votes             string
		WeightedVotes     string
		RemainingDuration string
		StartTime         string
		Decay             bool
	}

	VoteBucketList struct {
		EpochNumber uint64
		BucketList  []byte
	}

	CandidateList struct {
		EpochNumber   uint64
		CandidateList []byte
	}

	ProbationList struct {
		EpochNumber   uint64
		IntensityRate uint64
		Address       string
		Count         uint64
	}
)
