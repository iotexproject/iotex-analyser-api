package model

import "github.com/shopspring/decimal"

type HermesDropRecords struct {
	ID           uint64          `gorm:"primary_key"`
	EpochNumber  uint64          `gorm:"unsigned;index; uniqueIndex:idx_edv,priority:1;"`
	DelegateName string          `gorm:"size:42;index;not null;uniqueIndex:idx_edv,priority:1;"`
	VoterAddress string          `gorm:"size:42;not null;default:'';index:,type:hash;uniqueIndex:idx_edv,priority:1;"`
	Amount       decimal.Decimal `gorm:"type:decimal(42,0);not null;default:0;"`
	ActHash      string          `gorm:"size:64;not null;index:,length:9"`
	BucketID     uint64
}

func (HermesDropRecords) TableName() string {
	return "hermes_drop_records"
}
