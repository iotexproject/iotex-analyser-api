package actions

import "time"

type ActionInfo struct {
	BlkHeight          uint64
	ActHash            string
	BlkHash            string
	ActType            string
	Sender             string
	Recipient          string
	Amount             string
	GasFee             string
	GasPrice           string
	GasLimit           uint64
	GasConsumed        uint64
	Nonce              uint64
	Status             uint64
	ContractAddress    string
	ExecutionRevertMsg string
	ChainId            uint64
	Timestamp          time.Time
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
