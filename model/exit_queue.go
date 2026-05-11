package model

type CandidateExitQueue struct {
	ID                uint64
	CandidateName     string
	CandidateIdentity string
	Status            string
	RequestHeight     uint64
	RequestHash       string
	ScheduleHeight    uint64
	ScheduleHash      string
	ConfirmHeight     uint64
	ConfirmHash       string
	ScheduledAt       uint64
}

func (CandidateExitQueue) TableName() string {
	return "candidate_exit_queue"
}
