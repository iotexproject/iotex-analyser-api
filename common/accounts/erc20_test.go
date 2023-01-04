package accounts

import (
	"testing"

	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/stretchr/testify/require"
)

func TestReadERC20Decimals(t *testing.T) {
	require := require.New(t)
	conn, err := common.NewDefaultGRPCConn("api.mainnet.iotex.one:443")
	require.NoError(err)
	defer conn.Close()
	client := iotexapi.NewAPIServiceClient(conn)
	decimals, err := ReadERC20Decimals(client, "io1f4acssp65t6s90egjkzpvrdsrjjyysnvxgqjrh")
	require.NoError(err)
	require.Equal(18, decimals)
}
