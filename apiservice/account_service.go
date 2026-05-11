package apiservice

import (
	"context"
	"database/sql"
	"math/big"
	"strings"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/accounts"
	"github.com/iotexproject/iotex-analyser-api/common/actions"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-core/v2/ioctl/util"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// maxAuthorizationListLimit caps GetAuthorizationsByAuthority's first parameter
// to keep the result set bounded.
const maxAuthorizationListLimit = 100

type AccountService struct {
	api.UnimplementedAccountServiceServer
}

// IotexBalanceByHeight returns the balance of the given address at the given height.
// curl -d '{"address": ["io1ryztljunahyml9s7atfwtsx7s8wvr5maufa6zp", "io1j4mn2ga590z6es2fs07fy2wjn3yf09f4rkfljc"], "height":8927781 }' http://127.0.0.1:7778/api.AccountService.IotexBalanceByHeight
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

// grpcurl -plaintext -d '{"address": "io1ryztljunahyml9s7atfwtsx7s8wvr5maufa6zp", "height":8927781 }' 127.0.0.1:7777 api.AccountService.GetErc20TokenBalanceByHeight
// curl -d '{"address": ["io1ryztljunahyml9s7atfwtsx7s8wvr5maufa6zp", "io1j4mn2ga590z6es2fs07fy2wjn3yf09f4rkfljc"], "contract_address": "io1w97pslyg7qdayp8mfnffxkjkpapaf83wmmll2l", "height":8927781 }}' http://127.0.0.1:7778/api.AccountService.GetErc20TokenBalanceByHeight
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
	//fetch decimals

	decimals, err := accounts.ReadERC20DecimalsWithCache(contractAddress)
	if err != nil {
		return nil, err
	}
	resp.Decimals = uint64(decimals)

	balance, err := accounts.Erc20TokenBalanceByHeight(req.GetHeight(), addres, contractAddress)
	if err != nil {
		return nil, err
	}
	balances := make([]string, len(balance))
	for i := 0; i < len(balance); i++ {
		balances[i] = util.RauToString(balance[i], decimals)
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
	query := "SELECT DISTINCT sender, block_height FROM block_action_partition ORDER BY block_height desc limit ?"

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

// ContractInfo return contract info
func (s *AccountService) ContractInfo(ctx context.Context, req *api.ContractInfoRequest) (*api.ContractInfoResponse, error) {
	resp := &api.ContractInfoResponse{}

	contractAddresses := req.GetContractAddress()
	for _, contractAddress := range contractAddresses {
		contractsRes := &api.ContractInfoResponse_Contract{
			ContractAddress: contractAddress,
		}
		contractExist, blockHeight, err := accounts.ContractIsExist(contractAddress)
		if err != nil {
			return nil, err
		}
		contractsRes.Exist = contractExist
		if contractExist {
			actionInfo, err := actions.GetActionInfoByBlockHeightAndContractAddress(blockHeight, contractAddress)
			if err != nil {
				return nil, err
			}
			contractsRes.Deployer = actionInfo.Sender
			contractsRes.CreateTime = actionInfo.Timestamp.String()
			callTimes, gas, err := accounts.GetContractCallTimesAndAccumulatedGas(contractAddress)
			if err != nil {
				return nil, err
			}
			contractsRes.CallTimes = callTimes
			contractsRes.AccumulatedGas = gas
		}
		resp.Contracts = append(resp.Contracts, contractsRes)
	}
	return resp, nil
}

// GetAccountMeta returns account metadata (is_contract, block_height, bytecode hash) for given addresses
func (s *AccountService) GetAccountMeta(ctx context.Context, req *api.GetAccountMetaRequest) (*api.GetAccountMetaResponse, error) {
	resp := &api.GetAccountMetaResponse{}
	addrs := req.GetAddresses()
	if len(addrs) == 0 {
		return resp, nil
	}
	gormDB := db.DB()
	var rows []struct {
		Address              string
		IsContract           bool
		BlockHeight          uint64
		ContractBytecodeHash sql.NullString
	}
	if err := gormDB.WithContext(ctx).Raw(
		`SELECT address, is_contract, block_height,
			CASE WHEN contract_byte_code IS NOT NULL THEN ENCODE(sha256(contract_byte_code), 'hex') ELSE NULL END AS contract_bytecode_hash
		FROM account_meta WHERE address IN ?`,
		addrs,
	).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get account meta")
	}
	for _, r := range rows {
		info := &api.AccountMetaInfo{
			Address:     r.Address,
			IsContract:  r.IsContract,
			BlockHeight: r.BlockHeight,
		}
		if r.ContractBytecodeHash.Valid {
			info.ContractBytecodeHash = r.ContractBytecodeHash.String
		}
		resp.Accounts = append(resp.Accounts, info)
	}
	return resp, nil
}

// GetAddressNFTBalances returns NFT token balances grouped by (contract, type) for an address.
func (s *AccountService) GetAddressNFTBalances(ctx context.Context, req *api.GetAddressNFTBalancesRequest) (*api.GetAddressNFTBalancesResponse, error) {
	resp := &api.GetAddressNFTBalancesResponse{}
	gormDB := db.DB()
	var rows []struct {
		ContractAddress string
		ErcType         string
		Balance         string
	}
	q := `SELECT contract_address, ('xrc' || erc_type) AS erc_type,
               SUM(CASE WHEN erc_type = '721' THEN 1 ELSE token_value END)::text AS balance
          FROM erc1155_721_holders_v2_1_0
          WHERE holder = ? AND token_value > 0
          GROUP BY contract_address, erc_type`
	if err := gormDB.WithContext(ctx).Raw(q, req.GetAddress()).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get NFT balances")
	}
	for _, r := range rows {
		resp.Balances = append(resp.Balances, &api.NFTBalanceInfo{
			ContractAddress: r.ContractAddress,
			Type:            r.ErcType,
			Balance:         r.Balance,
		})
	}
	return resp, nil
}

// GetAddressTokenBalances returns ERC20 token balances (positive) for an address.
func (s *AccountService) GetAddressTokenBalances(ctx context.Context, req *api.GetAddressTokenBalancesRequest) (*api.GetAddressTokenBalancesResponse, error) {
	resp := &api.GetAddressTokenBalancesResponse{}
	gormDB := db.DB()
	addr := req.GetAddress()
	var rows []struct {
		ContractAddress string
		Balance         string
	}
	q := `SELECT contract_address, balance::text
          FROM (
              SELECT contract_address,
                  COALESCE(SUM(amount::numeric) FILTER (WHERE recipient = ?), 0) -
                  COALESCE(SUM(amount::numeric) FILTER (WHERE sender = ?), 0) AS balance
              FROM erc20_transfers
              WHERE recipient = ? OR sender = ?
              GROUP BY contract_address
          ) a
          WHERE balance > 0`
	if err := gormDB.WithContext(ctx).Raw(q, addr, addr, addr, addr).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get token balances")
	}
	for _, r := range rows {
		resp.Balances = append(resp.Balances, &api.TokenBalanceInfo{
			ContractAddress: r.ContractAddress,
			Balance:         r.Balance,
		})
	}
	return resp, nil
}

// GetTopAccounts returns top stakers from stats_top_list_view with stake/duration/mf filters.
func (s *AccountService) GetTopAccounts(ctx context.Context, req *api.GetTopAccountsRequest) (*api.GetTopAccountsResponse, error) {
	resp := &api.GetTopAccountsResponse{}
	gormDB := db.DB()

	var stakeAmountCond, stakeDurationCond, mfCond string
	if req.GetStakeAmount() == "less" {
		stakeAmountCond = "staked_amount < 10000000000000000000000"
	} else {
		stakeAmountCond = "staked_amount >= 10000000000000000000000"
	}
	if req.GetStakeDuration() == "less" {
		stakeDurationCond = "duration < 91"
	} else {
		stakeDurationCond = "duration >= 91"
	}
	if req.GetMf() == "" {
		mfCond = "mf = 0"
	} else {
		mfCond = "mf > 0"
	}
	wherePart := stakeAmountCond + " AND " + stakeDurationCond + " AND " + mfCond

	var count int64
	if err := gormDB.WithContext(ctx).Raw("SELECT COUNT(1) FROM stats_top_list_view WHERE " + wherePart).Scan(&count).Error; err != nil {
		return nil, errors.Wrap(err, "failed to count top accounts")
	}
	resp.Count = count

	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	var rows []struct {
		OwnerAddress sql.NullString
		BucketId     uint64
		StakedAmount sql.NullString
		Duration     sql.NullString
		Mf           sql.NullString
		LastUpdate   sql.NullString
		Balance      sql.NullString
	}
	q := `SELECT owner_address, bucket_id,
               ROUND(staked_amount / 1e18, 2)::text AS staked_amount,
               duration::text, mf::text,
               to_char(last_update AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS last_update,
               balance::text
          FROM stats_top_list_view
          WHERE ` + wherePart + ` ORDER BY staked_amount DESC LIMIT ? OFFSET ?`
	if err := gormDB.WithContext(ctx).Raw(q, first, skip).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get top accounts")
	}
	for _, r := range rows {
		row := &api.TopAccountRow{BucketId: r.BucketId}
		if r.OwnerAddress.Valid {
			row.OwnerAddress = r.OwnerAddress.String
		}
		if r.StakedAmount.Valid {
			row.StakedAmount = r.StakedAmount.String
		}
		if r.Duration.Valid {
			row.Duration = r.Duration.String
		}
		if r.Mf.Valid {
			row.Mf = r.Mf.String
		}
		if r.LastUpdate.Valid {
			row.LastUpdate = r.LastUpdate.String
		}
		if r.Balance.Valid {
			row.Balance = r.Balance.String
		}
		resp.Accounts = append(resp.Accounts, row)
	}
	return resp, nil
}

// GetTopAccountsByBalance returns top accounts sorted by IOTX balance (in_flow - out_flow) from account_income_count.
func (s *AccountService) GetTopAccountsByBalance(ctx context.Context, req *api.GetTopAccountsByBalanceRequest) (*api.GetTopAccountsByBalanceResponse, error) {
	resp := &api.GetTopAccountsByBalanceResponse{}
	gormDB := db.DB()

	limit := req.GetLimit()
	if limit <= 0 {
		limit = 20
	}
	offset := req.GetOffset()

	var count int64
	if err := gormDB.WithContext(ctx).Raw("SELECT COUNT(1) FROM account_income_count").Scan(&count).Error; err != nil {
		return nil, errors.Wrap(err, "failed to count accounts")
	}
	resp.Count = count

	var rows []struct {
		Address      string
		Balance      sql.NullString
		TotalActions int64
	}
	q := `SELECT address,
		(in_flow - out_flow)::text AS balance,
		(in_num_actions + out_num_actions) AS total_actions
	FROM account_income_count
	ORDER BY (in_flow - out_flow) DESC
	LIMIT ? OFFSET ?`
	if err := gormDB.WithContext(ctx).Raw(q, limit, offset).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get top accounts by balance")
	}
	for _, r := range rows {
		row := &api.AccountBalanceRow{
			Address:      r.Address,
			TotalActions: r.TotalActions,
		}
		if r.Balance.Valid {
			row.Balance = r.Balance.String
		}
		resp.Accounts = append(resp.Accounts, row)
	}
	return resp, nil
}

// GetContractCreateInfo returns the action hash and creator for a contract
func (s *AccountService) GetContractCreateInfo(ctx context.Context, req *api.GetContractCreateInfoRequest) (*api.GetContractCreateInfoResponse, error) {
	resp := &api.GetContractCreateInfoResponse{}
	gormDB := db.DB()
	var row struct {
		ActionHash sql.NullString
		Creator    sql.NullString
	}
	if err := gormDB.WithContext(ctx).Raw(
		`SELECT a.action_hash, a.sender AS creator
		FROM account_meta am
		LEFT JOIN block_action_partition a ON am.block_height = a.block_height AND am.address = a.contract_address
		WHERE am.address = ? LIMIT 1`,
		req.GetAddress(),
	).Scan(&row).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get contract create info")
	}
	if row.ActionHash.Valid {
		resp.ActionHash = row.ActionHash.String
	}
	if row.Creator.Valid {
		resp.Creator = row.Creator.String
	}
	return resp, nil
}

func (s *AccountService) GetAuthorizationsByAuthority(ctx context.Context, req *api.GetAuthorizationsByAuthorityRequest) (*api.GetAuthorizationsByAuthorityResponse, error) {
	authority := strings.ToLower(req.GetAuthority())
	if authority == "" {
		return nil, status.Error(codes.InvalidArgument, "authority is required")
	}
	if req.GetSkip() < 0 {
		return nil, status.Error(codes.InvalidArgument, "skip must be >= 0")
	}
	if req.GetFirst() < 0 {
		return nil, status.Error(codes.InvalidArgument, "first must be >= 0")
	}

	skip := int(req.GetSkip())
	first := int(req.GetFirst())
	if first <= 0 {
		first = 20
	}
	if first > maxAuthorizationListLimit {
		first = maxAuthorizationListLimit
	}

	// Use the package-level DB handle (initialized once at startup) rather
	// than calling db.Connect() per request, which would re-init the global
	// pool and race with other handlers.
	gormDB := db.DB()

	var count int64
	if err := gormDB.WithContext(ctx).Raw(
		`SELECT COUNT(*) FROM "authorization" WHERE authority = ?`, authority,
	).Scan(&count).Error; err != nil {
		return nil, errors.Wrap(err, "failed to count authorizations")
	}

	type authRow struct {
		ActionHash  string
		BlockHeight uint64
		ChainID     string
		Address     string
		Nonce       string
		YParity     string
		Authority   string
		Valid       *bool
	}
	var rows []authRow
	if err := gormDB.WithContext(ctx).Raw(
		`SELECT action_hash, block_height, chain_id, address, nonce, y_parity, authority, valid
		 FROM "authorization" WHERE authority = ?
		 ORDER BY block_height DESC, "index" DESC
		 LIMIT ? OFFSET ?`,
		authority, first, skip,
	).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query authorizations")
	}

	entries := make([]*api.AuthorizationHistoryEntry, 0, len(rows))
	for _, r := range rows {
		entry := &api.AuthorizationHistoryEntry{
			ActionHash:  r.ActionHash,
			BlockHeight: r.BlockHeight,
			ChainId:     r.ChainID,
			Address:     r.Address,
			Nonce:       r.Nonce,
			YParity:     r.YParity,
			Authority:   r.Authority,
		}
		if r.Valid != nil {
			entry.Valid = *r.Valid
		}
		entries = append(entries, entry)
	}

	return &api.GetAuthorizationsByAuthorityResponse{
		Authorizations: entries,
		Count:          count,
	}, nil
}
