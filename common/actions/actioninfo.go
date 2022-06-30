package actions

type ActionInfo struct {
	BlkHeight uint64
	ActHash   string
	BlkHash   string
	ActType   string
	Sender    string
	Recipient string
	Amount    string
	GasFee    string
	Timestamp uint64
}

type BlockReceiptTransaction struct {
	ID          uint64
	BlockHeight uint64
	ActionHash  string
	Type        string
	Amount      string
	Sender      string
	Recipient   string
}
