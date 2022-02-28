package votings

import (
	"os"
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
	_, err := config.New(os.Getenv("ConfigPath"))
	if err != nil {
		return nil, err
	}
	return db.Connect()
}
