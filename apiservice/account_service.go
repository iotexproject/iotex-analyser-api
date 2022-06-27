package apiservice

import (
	"context"
	"database/sql"
	"math/big"
	"sort"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/rewards"
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
	db := db.DB()
	height := req.GetHeight()
	contractAddress := req.GetContractAddress()
	for _, addr := range req.GetAddress() {
		if len(addr) > 2 && (addr[:2] == "0x" || addr[:2] == "0X") {
			add, err := address.FromHex(addr)
			if err != nil {
				return nil, err
			}

			addr = add.String()
		}

		if len(contractAddress) > 2 && (contractAddress[:2] == "0x" || contractAddress[:2] == "0X") {
			add, err := address.FromHex(contractAddress)
			if err != nil {
				return nil, err
			}

			contractAddress = add.String()
		}
		//get receive amount
		var toAmount sql.NullString
		query := "SELECT SUM(amount) FROM token_erc20 t WHERE block_height<=? AND t.recipient=? AND t.contract_address=?"
		err := db.Raw(query, height, addr, contractAddress).Scan(&toAmount).Error
		if err != nil {
			return nil, err
		}

		//get cost amount
		var fromAmount sql.NullString
		query = "SELECT SUM(amount) FROM token_erc20 t WHERE block_height<=? AND t.sender=? AND contract_address=?"
		err = db.Raw(query, height, addr, contractAddress).Scan(&fromAmount).Error
		if err != nil {
			return nil, err
		}

		to, ok := big.NewInt(0).SetString(toAmount.String, 10)
		if !ok {
			to = big.NewInt(0)
		}
		from, ok := big.NewInt(0).SetString(fromAmount.String, 10)
		if !ok {
			from = big.NewInt(0)
		}
		balance := new(big.Int).Sub(to, from)
		resp.Balance = append(resp.Balance, util.RauToString(balance, 6))
	}

	return resp, nil
}

//grpcurl -plaintext -d '{"startEpoch": 22416, "epochCount": 1, "rewardAddress": "io12mgttmfa2ffn9uqvn0yn37f4nz43d248l2ga85"}' 127.0.0.1:8888 api.AccountService.Hermes
func (s *AccountService) Hermes(ctx context.Context, req *api.HermesRequest) (*api.HermesResponse, error) {
	resp := &api.HermesResponse{}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	rewardAddress := req.GetRewardAddress()
	endEpoch := startEpoch + epochCount - 1
	waiverThreshold := 100

	distributePlanMap, err := distributionPlanByRewardAddress(startEpoch, endEpoch, rewardAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get reward distribution plan")
	}

	delegateMap := make(map[uint64][]string)
	for delegateName, planMap := range distributePlanMap {
		for epochNumber := range planMap {
			if _, ok := delegateMap[epochNumber]; !ok {
				delegateMap[epochNumber] = make([]string, 0)
			}
			delegateMap[epochNumber] = append(delegateMap[epochNumber], delegateName)
		}
	}

	accountRewardsMap, err := accountRewards(delegateMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account rewards")
	}

	voterVotesMap, err := rewards.WeightedVotesBySearchPairs(delegateMap)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get voter votes")
	}
	hermesDistributions := make([]*DelegateHermesDistribution, 0, len(accountRewardsMap))
	for delegate, rewardsMap := range accountRewardsMap {
		planMap := distributePlanMap[delegate]
		epochVoterMap := voterVotesMap[delegate]

		voterAddrToReward := make(map[string]*big.Int)
		balanceAfterDistribution := big.NewInt(0)
		voterCountMap := make(map[string]bool)
		feeWaiver := true
		var stakingAddress string

		for epoch, rewards := range rewardsMap {
			distributePlan := planMap[epoch]
			voterMap := epochVoterMap[epoch]

			if stakingAddress == "" {
				stakingAddress = distributePlan.StakingAddress
			}

			totalRewards := new(big.Int).Set(rewards.BlockReward)
			totalRewards.Add(totalRewards, rewards.EpochReward).Add(totalRewards, rewards.FoundationBonus)
			balanceAfterDistribution.Add(balanceAfterDistribution, totalRewards)
			waiverThresholdF := float64(waiverThreshold)
			if distributePlan.BlockRewardPercentage < waiverThresholdF || distributePlan.EpochRewardPercentage < waiverThresholdF || distributePlan.FoundationBonusPercentage < waiverThresholdF {
				feeWaiver = false
			}
			distrReward, err := calculatedDistributedReward(distributePlan, rewards)
			if err != nil {
				return nil, errors.Wrap(err, "failed to calculate distributed reward")
			}
			for voterAddr, weightedVotes := range voterMap {
				amount := new(big.Int).Set(distrReward)
				amount = amount.Mul(amount, weightedVotes).Div(amount, distributePlan.TotalWeightedVotes)
				if _, ok := voterAddrToReward[voterAddr]; !ok {
					voterAddrToReward[voterAddr] = big.NewInt(0)
				}
				voterAddrToReward[voterAddr].Add(voterAddrToReward[voterAddr], amount)
				balanceAfterDistribution.Sub(balanceAfterDistribution, amount)
				voterCountMap[voterAddr] = true
			}
		}
		rewardDistribution, err := convertVoterDistributionMapToList(voterAddrToReward)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert voter distribution map to list")
		}
		hermesDistributions = append(hermesDistributions, &DelegateHermesDistribution{
			DelegateName:        delegate,
			Distributions:       rewardDistribution,
			StakingIotexAddress: stakingAddress,
			VoterCount:          uint64(len(voterCountMap)),
			WaiveServiceFee:     feeWaiver,
			Refund:              balanceAfterDistribution.String(),
		})
	}

	hermesDistribution := make([]*api.HermesDistribution, 0, len(hermesDistributions))
	for _, ret := range hermesDistributions {
		rds := make([]*api.RewardDistribution, 0)
		for _, distribution := range ret.Distributions {
			v := &api.RewardDistribution{
				VoterEthAddress:   distribution.VoterEthAddress,
				VoterIotexAddress: distribution.VoterIotexAddress,
				Amount:            distribution.Amount,
			}
			rds = append(rds, v)
		}
		sort.Slice(rds, func(i, j int) bool { return rds[i].VoterEthAddress < rds[j].VoterEthAddress })

		hermesDistribution = append(hermesDistribution, &api.HermesDistribution{
			DelegateName:        ret.DelegateName,
			RewardDistribution:  rds,
			StakingIotexAddress: ret.StakingIotexAddress,
			VoterCount:          ret.VoterCount,
			WaiveServiceFee:     ret.WaiveServiceFee,
			Refund:              ret.Refund,
		})
	}
	sort.Slice(hermesDistribution, func(i, j int) bool { return hermesDistribution[i].DelegateName < hermesDistribution[j].DelegateName })
	resp.HermesDistribution = hermesDistribution
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

	query := "SELECT count(DISTINCT(address)) FROM account_income_count"

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
