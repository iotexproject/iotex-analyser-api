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
