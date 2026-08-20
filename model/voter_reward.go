package model

// Mirrors of the four tables the analyser's voter_reward plugin writes. Only
// the columns this API reads are declared; GORM maps by name so the rest are
// simply not selected.

type VoterRewardDistribution struct {
	ID               uint64
	BlockHeight      uint64
	EpochNumber      uint64
	Era              uint64
	ActionHash       string
	LogIndex         uint32
	RowIndex         uint32
	DelegateID       string `gorm:"column:delegate_id"`
	DelegateName     string
	VoterAddress     string
	RecipientAddress string
	Amount           string
	Compounded       bool
	CompoundBucketID uint64 `gorm:"column:compound_bucket_id"`
}

func (VoterRewardDistribution) TableName() string { return "voter_reward_distribution" }

type VoterRewardEra struct {
	Era              uint64
	FreezeHeight     uint64
	FirstChunkAt     uint64
	CompletedHeight  uint64
	ScanPhase        uint32
	ResumeVoter      string
	DelegateCount    uint32
	VoterCount       uint64
	TotalFrozen      string
	TotalDistributed string
	OverrunResidue   string
	LastChunkAt      uint64
}

func (VoterRewardEra) TableName() string { return "voter_reward_era" }

type DelegateRewardConfig struct {
	ID                   uint64
	DelegateID           string `gorm:"column:delegate_id"`
	Era                  uint64
	BlockHeight          uint64
	DelegateName         string
	OptedIn              bool
	OptInSource          string
	OptInHeight          uint64
	FreezeHeight         uint64
	BlockCommissionBps   uint64
	EpochCommissionBps   uint64
	CommissionConfigured bool
	TotalWeight          string
	VoterAmountFrozen    string
	PayoutAddress        string
}

func (DelegateRewardConfig) TableName() string { return "delegate_reward_config" }

type VoterRewardDestination struct {
	ID           uint64
	BlockHeight  uint64
	ActionHash   string
	VoterAddress string
	OldRecipient string
	NewRecipient string
}

func (VoterRewardDestination) TableName() string { return "voter_reward_destination" }
