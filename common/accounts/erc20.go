package accounts

import (
	"context"
	"encoding/hex"
	"math/big"
	"strconv"
	"strings"
	"sync"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/config"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-core/action"
	"github.com/iotexproject/iotex-core/test/identityset"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
)

func Erc20TokenBalanceByHeight(height uint64, addresses []string, contractAddress string) ([]*big.Int, error) {
	result := make([]*big.Int, 0)
	db := db.DB()

	var res []struct {
		Address string
		Balance string
	}
	//select address, ramount-COALESCE(samount, 0) balance from (SELECT t.recipient address,COALESCE(SUM(amount),0) ramount FROM erc20_transfers t WHERE block_height<=8947000 AND t.recipient='io1slh2qa2q0zd0skmyj3aw5gdr7qzqpdvh32zzpc' AND t.contract_address='io1qppu9nz834xqrenzllr4h57hzfpqefd0xnsu3d' group by address) t1 left join (SELECT t.sender address,COALESCE(SUM(amount),0) samount FROM erc20_transfers t WHERE block_height<=8947000 AND t.sender='io1slh2qa2q0zd0skmyj3aw5gdr7qzqpdvh32zzpd' AND t.contract_address='io1qppu9nz834xqrenzllr4h57hzfpqefd0xnsu3d' group by address) t2 using(address)

	err := db.Raw("select address, ramount-COALESCE(samount, 0) balance from (SELECT t.recipient address,COALESCE(SUM(amount),0) ramount FROM erc20_transfers t WHERE block_height<=? AND t.recipient in ? AND t.contract_address=? group by address) t1 left join (SELECT t.sender address,COALESCE(SUM(amount),0) samount FROM erc20_transfers t WHERE block_height<=? AND t.sender in ? AND t.contract_address=? group by address) t2 using(address)", height, addresses, contractAddress, height, addresses, contractAddress).Scan(&res).Error
	if err != nil {
		return nil, err
	}
	results := make(map[string]*big.Int)
	for _, v := range res {
		balance, ok := new(big.Int).SetString(v.Balance, 10)
		if !ok {
			balance = big.NewInt(0)
		}
		results[v.Address] = balance
	}

	for _, addr := range addresses {
		if addr[:2] == "0x" || addr[:2] == "0X" {
			add, err := address.FromHex(addr)
			if err != nil {
				return nil, err
			}
			addr = add.String()

		}
		addr = strings.ToLower(addr)
		if results[addr] == nil {
			result = append(result, big.NewInt(0))
		} else {
			result = append(result, results[addr])
		}
	}

	return result, nil
}

var (
	_ERC20ContractDecimals = sync.Map{}
)

func ReadERC20DecimalsWithCache(contractAddress string) (int, error) {
	val, ok := _ERC20ContractDecimals.Load(contractAddress)
	if ok {
		return val.(int), nil
	}
	conn, err := common.NewDefaultGRPCConn(config.Default.RPC)
	if err != nil {
		return 6, nil
	}
	defer conn.Close()
	client := iotexapi.NewAPIServiceClient(conn)
	decimals, err := ReadERC20Decimals(client, contractAddress)
	if err != nil {
		return 6, nil
	}
	_ERC20ContractDecimals.Store(contractAddress, decimals)
	return decimals, nil

}

// ReadERC20Decimals read ERC20 decimals
func ReadERC20Decimals(client iotexapi.APIServiceClient, contractAddr string) (int, error) {
	decimals := 6 //default decimal

	nonce := uint64(1)
	transferAmount := big.NewInt(0)
	gasLimit := uint64(100000)
	gasPrice := big.NewInt(10000000)
	callerAddress := identityset.Address(30).String()
	callData, _ := hex.DecodeString("313ce567")
	execution, err := action.NewExecution(contractAddr, nonce, transferAmount, gasLimit, gasPrice, callData)
	if err != nil {
		return decimals, nil
	}
	request := &iotexapi.ReadContractRequest{
		Execution:     execution.Proto(),
		CallerAddress: callerAddress,
	}

	res, err := client.ReadContract(context.Background(), request)
	if err != nil {
		return decimals, nil
	}
	if res.Data != "" {
		tmp, _ := strconv.ParseInt(res.Data, 16, 64)
		decimals = int(tmp)
	}
	return decimals, nil
}
