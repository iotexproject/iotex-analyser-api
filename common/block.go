package common

import (
	"time"

	"github.com/iotexproject/iotex-analyser-api/db"
)

func GetBlockHeightRangeByMonth(t time.Time) (uint64, uint64, error) {
	db := db.DB()

	var result struct {
		MinHeight uint64
		MaxHeight uint64
	}
	firstDayInMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.UTC)
	lastDayInMonth := firstDayInMonth.AddDate(0, 1, -1)
	query := "select min(block_height) min_height,max(block_height) max_height from block where timestamp::date between ? and ?"
	if err := db.Raw(query, firstDayInMonth.Format("2006-01-02"), lastDayInMonth.Format("2006-01-02")).Scan(&result).Error; err != nil {
		return 0, 0, err
	}
	return result.MinHeight, result.MaxHeight, nil
}

func GetBlockHeightRangeByDay(t time.Time) (uint64, uint64, error) {
	db := db.DB()

	var result struct {
		MinHeight uint64
		MaxHeight uint64
	}
	query := "select min(block_height) min_height,max(block_height) max_height from block where timestamp::date =?"
	if err := db.Raw(query, t.Format("2006-01-02")).Scan(&result).Error; err != nil {
		return 0, 0, err
	}
	return result.MinHeight, result.MaxHeight, nil
}
