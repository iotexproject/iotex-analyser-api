package votings

import (
	"testing"

	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestGetProbationList(t *testing.T) {
	require := require.New(t)
	_, err := initDB()
	require.NoError(err)
	probationList, err := GetProbationList(8750)
	require.NoError(err)
	require.Equal(2, len(probationList))
	require.Equal(uint64(8750), probationList[0].EpochNumber)
	require.Equal(uint64(1), probationList[0].Count)
	require.Equal("io1n3llm0zzlles6pvpzugzajyzyjehztztxf2rff", probationList[0].Address)
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
