package apiservice

import (
	"context"
	"math/big"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/accounts"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-core/ioctl/util"
	"github.com/pkg/errors"
)

type AccountService struct {
	api.UnimplementedAccountServiceServer
}

// IotexBalanceByHeight returns the balance of the given address at the given height.
//curl -d '{"address": ["io1ryztljunahyml9s7atfwtsx7s8wvr5maufa6zp", "io1j4mn2ga590z6es2fs07fy2wjn3yf09f4rkfljc"], "height":8927781 }' http://127.0.0.1:7778/api.AccountService.IotexBalanceByHeight
func (s *AccountService) IotexBalanceByHeight(ctx context.Context, req *api.IotexBalanceByHeightRequest) (*api.IotexBalanceByHeightResponse, error) {
	resp := &api.IotexBalanceByHeightResponse{
		Height: req.GetHeight(),
	}
	balance, err := common.AccountBalanceByHeight(req.GetHeight(), req.Address)
	if err != nil {
		return nil, err
	}
	balances := make([]string, len(balance))
	for i := 0; i < len(balance); i++ {
		balances[i] = util.RauToString(balance[i], util.IotxDecimalNum)
	}

	resp.Balance = balances

	return resp, nil
}

//grpcurl -plaintext -d '{"address": "io1ryztljunahyml9s7atfwtsx7s8wvr5maufa6zp", "height":8927781 }' 127.0.0.1:7777 api.AccountService.GetErc20TokenBalanceByHeight
//curl -d '{"address": ["io1ryztljunahyml9s7atfwtsx7s8wvr5maufa6zp", "io1j4mn2ga590z6es2fs07fy2wjn3yf09f4rkfljc"], "contract_address": "io1w97pslyg7qdayp8mfnffxkjkpapaf83wmmll2l", "height":8927781 }}' http://127.0.0.1:7778/api.AccountService.GetErc20TokenBalanceByHeight
func (s *AccountService) Erc20TokenBalanceByHeight(ctx context.Context, req *api.Erc20TokenBalanceByHeightRequest) (*api.Erc20TokenBalanceByHeightResponse, error) {
	resp := &api.Erc20TokenBalanceByHeightResponse{
		Height:          req.GetHeight(),
		ContractAddress: req.GetContractAddress(),
	}
	contractAddress := req.GetContractAddress()
	if len(contractAddress) > 2 && (contractAddress[:2] == "0x" || contractAddress[:2] == "0X") {
		add, err := address.FromHex(contractAddress)
		if err != nil {
			return nil, err
		}

		contractAddress = add.String()
	}
	addres := make([]string, 0)
	for _, addr := range req.GetAddress() {
		if len(addr) > 2 && addr[:2] == "0x" || addr[:2] == "0X" {
			add, err := address.FromHex(addr)
			if err != nil {
				return nil, err
			}
			addr = add.String()

		}
		addres = append(addres, addr)
	}

	balance, err := accounts.Erc20TokenBalanceByHeight(req.GetHeight(), addres, contractAddress)
	if err != nil {
		return nil, err
	}
	balances := make([]string, len(balance))
	for i := 0; i < len(balance); i++ {
		balances[i] = util.RauToString(balance[i], 6)
	}

	resp.Balance = balances

	return resp, nil
}

// ActiveAccounts returns the active accounts.
func (s *AccountService) ActiveAccounts(ctx context.Context, req *api.ActiveAccountsRequest) (*api.ActiveAccountsResponse, error) {
	db := db.DB()
	resp := &api.ActiveAccountsResponse{}

	count := req.GetCount()
	if count == 0 {
		count = 100
	}
	query := "SELECT DISTINCT sender, block_height FROM block_action ORDER BY block_height desc limit ?"

	var results []struct {
		Sender      string
		BlockHeight uint64
	}
	if err := db.Raw(query, count).Scan(&results).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get active accounts")
	}
	resp.ActiveAccounts = make([]string, 0, len(results))
	for _, result := range results {
		resp.ActiveAccounts = append(resp.ActiveAccounts, result.Sender)
	}
	return resp, nil
}

// OperatorAddress finds the delegate's operator address given the delegate's alias name
func (s *AccountService) OperatorAddress(ctx context.Context, req *api.OperatorAddressRequest) (*api.OperatorAddressResponse, error) {
	db := db.DB()
	resp := &api.OperatorAddressResponse{}

	aliasName := req.GetAliasName()

	query := "SELECT operator_address FROM delegate where name=? ORDER BY id desc"

	var result string
	if err := db.WithContext(ctx).Raw(query, aliasName).Scan(&result).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get operator address")
	}
	resp.Exist = result != ""
	resp.OperatorAddress = result
	return resp, nil
}

// Alias returns the alias name given the delegate's operator address
func (s *AccountService) Alias(ctx context.Context, req *api.AliasRequest) (*api.AliasResponse, error) {
	db := db.DB()
	resp := &api.AliasResponse{}

	operatorAddress := req.GetOperatorAddress()

	query := "SELECT name FROM delegate where operator_address=? ORDER BY id desc"

	var result string
	if err := db.WithContext(ctx).Raw(query, operatorAddress).Scan(&result).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get alias")
	}
	resp.Exist = result != ""
	resp.AliasName = result
	return resp, nil
}

// TotalNumberOfHolders returns the total number of holders
func (s *AccountService) TotalNumberOfHolders(ctx context.Context, req *api.TotalNumberOfHoldersRequest) (*api.TotalNumberOfHoldersResponse, error) {
	db := db.DB()
	resp := &api.TotalNumberOfHoldersResponse{}

	query := "SELECT COUNT(*) FROM (SELECT DISTINCT address FROM account_income_count) AS temp"

	var result uint64
	if err := db.WithContext(ctx).Raw(query).Scan(&result).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get total number of holders")
	}
	resp.TotalNumberOfHolders = result
	return resp, nil
}

// TotalAccountSupply returns total amount of tokens held by IoTeX accounts
func (s *AccountService) TotalAccountSupply(ctx context.Context, req *api.TotalAccountSupplyRequest) (*api.TotalAccountSupplyResponse, error) {
	db := db.DB()
	resp := &api.TotalAccountSupplyResponse{}

	query := "SELECT sum(in_flow)-sum(out_flow) FROM account_income_count"

	var result string
	if err := db.WithContext(ctx).Raw(query).Scan(&result).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get total number of holders")
	}
	ret, ok := new(big.Int).SetString(result, 10)
	if !ok {
		return nil, errors.New("failed to format to big int:" + result)
	}
	if ret.Sign() < 0 {
		result = ret.Abs(ret).String()
	}
	resp.TotalAccountSupply = result
	return resp, nil
}
