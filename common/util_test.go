package common

import (
	"testing"

	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestGetCurrentEpochAndHeight(t *testing.T) {
	require := require.New(t)
	_, err := initDB()
	require.NoError(err)
	epoch, height, err := GetCurrentEpochAndHeight()
	require.NoError(err)
	require.Equal(uint64(236), epoch)
	require.Equal(uint64(84601), height)
}

func initDB() (*gorm.DB, error) {
	config.Default.Database = config.Database{
		Driver:   "postgres",
		Name:     "mainnet",
		Host:     "127.0.0.1",
		Port:     "5435",
		User:     "postgres",
		Password: "admin",
		Debug:    true,
	}
	return db.Connect()
}
