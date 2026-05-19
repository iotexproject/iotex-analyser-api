# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api_account.proto](#api_account-proto)
    - [AccountBalanceRow](#api-AccountBalanceRow)
    - [AccountMetaInfo](#api-AccountMetaInfo)
    - [ActiveAccountsRequest](#api-ActiveAccountsRequest)
    - [ActiveAccountsResponse](#api-ActiveAccountsResponse)
    - [AliasRequest](#api-AliasRequest)
    - [AliasResponse](#api-AliasResponse)
    - [AuthorizationHistoryEntry](#api-AuthorizationHistoryEntry)
    - [ContractInfoRequest](#api-ContractInfoRequest)
    - [ContractInfoResponse](#api-ContractInfoResponse)
    - [ContractInfoResponse.Contract](#api-ContractInfoResponse-Contract)
    - [Erc20TokenBalanceByHeightRequest](#api-Erc20TokenBalanceByHeightRequest)
    - [Erc20TokenBalanceByHeightResponse](#api-Erc20TokenBalanceByHeightResponse)
    - [GetAccountMetaRequest](#api-GetAccountMetaRequest)
    - [GetAccountMetaResponse](#api-GetAccountMetaResponse)
    - [GetAddressNFTBalancesRequest](#api-GetAddressNFTBalancesRequest)
    - [GetAddressNFTBalancesResponse](#api-GetAddressNFTBalancesResponse)
    - [GetAddressTokenBalancesRequest](#api-GetAddressTokenBalancesRequest)
    - [GetAddressTokenBalancesResponse](#api-GetAddressTokenBalancesResponse)
    - [GetAuthorizationsByAuthorityRequest](#api-GetAuthorizationsByAuthorityRequest)
    - [GetAuthorizationsByAuthorityResponse](#api-GetAuthorizationsByAuthorityResponse)
    - [GetContractCreateInfoRequest](#api-GetContractCreateInfoRequest)
    - [GetContractCreateInfoResponse](#api-GetContractCreateInfoResponse)
    - [GetTopAccountsByBalanceRequest](#api-GetTopAccountsByBalanceRequest)
    - [GetTopAccountsByBalanceResponse](#api-GetTopAccountsByBalanceResponse)
    - [GetTopAccountsRequest](#api-GetTopAccountsRequest)
    - [GetTopAccountsResponse](#api-GetTopAccountsResponse)
    - [IotexBalanceByHeightRequest](#api-IotexBalanceByHeightRequest)
    - [IotexBalanceByHeightResponse](#api-IotexBalanceByHeightResponse)
    - [NFTBalanceInfo](#api-NFTBalanceInfo)
    - [OperatorAddressRequest](#api-OperatorAddressRequest)
    - [OperatorAddressResponse](#api-OperatorAddressResponse)
    - [TokenBalanceInfo](#api-TokenBalanceInfo)
    - [TopAccountRow](#api-TopAccountRow)
    - [TotalAccountSupplyRequest](#api-TotalAccountSupplyRequest)
    - [TotalAccountSupplyResponse](#api-TotalAccountSupplyResponse)
    - [TotalNumberOfHoldersRequest](#api-TotalNumberOfHoldersRequest)
    - [TotalNumberOfHoldersResponse](#api-TotalNumberOfHoldersResponse)
  
    - [AccountService](#api-AccountService)
  
- [api_action.proto](#api_action-proto)
    - [ActionByAddressRequest](#api-ActionByAddressRequest)
    - [ActionByAddressResponse](#api-ActionByAddressResponse)
    - [ActionByDatesRequest](#api-ActionByDatesRequest)
    - [ActionByDatesResponse](#api-ActionByDatesResponse)
    - [ActionByHashRequest](#api-ActionByHashRequest)
    - [ActionByHashResponse](#api-ActionByHashResponse)
    - [ActionByHashResponse.ActionLog](#api-ActionByHashResponse-ActionLog)
    - [ActionByHashResponse.ActionTypeInfo](#api-ActionByHashResponse-ActionTypeInfo)
    - [ActionByHashResponse.AuthorizationEntry](#api-ActionByHashResponse-AuthorizationEntry)
    - [ActionByHashResponse.EvmTransfers](#api-ActionByHashResponse-EvmTransfers)
    - [ActionByHashResponse.StakeAction](#api-ActionByHashResponse-StakeAction)
    - [ActionByHashResponse.TokenTransfer](#api-ActionByHashResponse-TokenTransfer)
    - [ActionByHeightRequest](#api-ActionByHeightRequest)
    - [ActionByHeightResponse](#api-ActionByHeightResponse)
    - [ActionByTypeRequest](#api-ActionByTypeRequest)
    - [ActionByTypeResponse](#api-ActionByTypeResponse)
    - [ActionInfo](#api-ActionInfo)
    - [ActionListRequest](#api-ActionListRequest)
    - [ActionListResponse](#api-ActionListResponse)
    - [ActionRequest](#api-ActionRequest)
    - [ActionResponse](#api-ActionResponse)
    - [ContractInteractorsRequest](#api-ContractInteractorsRequest)
    - [ContractInteractorsResponse](#api-ContractInteractorsResponse)
    - [EvmTransferInfo](#api-EvmTransferInfo)
    - [EvmTransfersByAddressRequest](#api-EvmTransfersByAddressRequest)
    - [EvmTransfersByAddressResponse](#api-EvmTransfersByAddressResponse)
    - [EvmTransfersByAddressResponse.EvmTransfer](#api-EvmTransfersByAddressResponse-EvmTransfer)
    - [GetInternalTxnsRequest](#api-GetInternalTxnsRequest)
    - [GetInternalTxnsResponse](#api-GetInternalTxnsResponse)
    - [GetStakingActionsByAddressRequest](#api-GetStakingActionsByAddressRequest)
    - [GetStakingActionsByAddressResponse](#api-GetStakingActionsByAddressResponse)
    - [InternalTxnInfo](#api-InternalTxnInfo)
    - [StakingActionInfo](#api-StakingActionInfo)
    - [XrcInfo](#api-XrcInfo)
  
    - [ActionService](#api-ActionService)
  
- [api_actions.proto](#api_actions-proto)
    - [ActionsRequest](#api-ActionsRequest)
    - [AllActionsByAddressResponse](#api-AllActionsByAddressResponse)
    - [AllActionsByAddressResult](#api-AllActionsByAddressResult)
    - [EvmTransferDetailListByAddressResponse](#api-EvmTransferDetailListByAddressResponse)
    - [EvmTransferDetailResult](#api-EvmTransferDetailResult)
  
    - [AllActionsByAddressResult.RecordType](#api-AllActionsByAddressResult-RecordType)
  
    - [ActionsService](#api-ActionsService)
  
- [api_approval.proto](#api_approval-proto)
    - [GetXRC20ApprovalsRequest](#api-GetXRC20ApprovalsRequest)
    - [GetXRC20ApprovalsResponse](#api-GetXRC20ApprovalsResponse)
    - [GetXRC721ApprovalsRequest](#api-GetXRC721ApprovalsRequest)
    - [GetXRC721ApprovalsResponse](#api-GetXRC721ApprovalsResponse)
    - [XRC20ApprovalInfo](#api-XRC20ApprovalInfo)
    - [XRC721ApprovalInfo](#api-XRC721ApprovalInfo)
  
    - [ApprovalService](#api-ApprovalService)
  
- [api_chain.proto](#api_chain-proto)
    - [ActionHistoryPoint](#api-ActionHistoryPoint)
    - [BlockInfo](#api-BlockInfo)
    - [BlockSizeByHeightRequest](#api-BlockSizeByHeightRequest)
    - [BlockSizeByHeightResponse](#api-BlockSizeByHeightResponse)
    - [ChainRequest](#api-ChainRequest)
    - [ChainResponse](#api-ChainResponse)
    - [ChainResponse.Rewards](#api-ChainResponse-Rewards)
    - [GasHistoryPoint](#api-GasHistoryPoint)
    - [GetActionHistoryRequest](#api-GetActionHistoryRequest)
    - [GetActionHistoryResponse](#api-GetActionHistoryResponse)
    - [GetBlockByHeightRequest](#api-GetBlockByHeightRequest)
    - [GetBlockByHeightResponse](#api-GetBlockByHeightResponse)
    - [GetBlocksRequest](#api-GetBlocksRequest)
    - [GetBlocksResponse](#api-GetBlocksResponse)
    - [GetChainStatsRequest](#api-GetChainStatsRequest)
    - [GetChainStatsResponse](#api-GetChainStatsResponse)
    - [GetEpochInfoRequest](#api-GetEpochInfoRequest)
    - [GetEpochInfoResponse](#api-GetEpochInfoResponse)
    - [GetGasHistoryRequest](#api-GetGasHistoryRequest)
    - [GetGasHistoryResponse](#api-GetGasHistoryResponse)
    - [GetLatestBlockHeightRequest](#api-GetLatestBlockHeightRequest)
    - [GetLatestBlockHeightResponse](#api-GetLatestBlockHeightResponse)
    - [GetLatestStakingRecordRequest](#api-GetLatestStakingRecordRequest)
    - [GetLatestStakingRecordResponse](#api-GetLatestStakingRecordResponse)
    - [GetPeakTpsRequest](#api-GetPeakTpsRequest)
    - [GetPeakTpsResponse](#api-GetPeakTpsResponse)
    - [GetStakingRatioHistoryRequest](#api-GetStakingRatioHistoryRequest)
    - [GetStakingRatioHistoryResponse](#api-GetStakingRatioHistoryResponse)
    - [GetSupplyHistoryRequest](#api-GetSupplyHistoryRequest)
    - [GetSupplyHistoryResponse](#api-GetSupplyHistoryResponse)
    - [GetTpsHistoryRequest](#api-GetTpsHistoryRequest)
    - [GetTpsHistoryResponse](#api-GetTpsHistoryResponse)
    - [MostRecentTPSRequest](#api-MostRecentTPSRequest)
    - [MostRecentTPSResponse](#api-MostRecentTPSResponse)
    - [NumberOfActionsRequest](#api-NumberOfActionsRequest)
    - [NumberOfActionsResponse](#api-NumberOfActionsResponse)
    - [StakingRatioPoint](#api-StakingRatioPoint)
    - [SupplyHistoryPoint](#api-SupplyHistoryPoint)
    - [TotalTransferredTokensRequest](#api-TotalTransferredTokensRequest)
    - [TotalTransferredTokensResponse](#api-TotalTransferredTokensResponse)
    - [TpsHistoryPoint](#api-TpsHistoryPoint)
    - [VotingResultMeta](#api-VotingResultMeta)
  
    - [ChainService](#api-ChainService)
  
- [api_delegate.proto](#api_delegate-proto)
    - [BookKeepingRequest](#api-BookKeepingRequest)
    - [BookKeepingResponse](#api-BookKeepingResponse)
    - [BucketInfo](#api-BucketInfo)
    - [BucketInfoList](#api-BucketInfoList)
    - [BucketInfoRequest](#api-BucketInfoRequest)
    - [BucketInfoResponse](#api-BucketInfoResponse)
    - [DelegateRewardDistribution](#api-DelegateRewardDistribution)
    - [PaidToDelegatesRequest](#api-PaidToDelegatesRequest)
    - [PaidToDelegatesResponse](#api-PaidToDelegatesResponse)
    - [PaidToDelegatesResponse.DelegateInfo](#api-PaidToDelegatesResponse-DelegateInfo)
    - [ProbationHistoricalRateRequest](#api-ProbationHistoricalRateRequest)
    - [ProbationHistoricalRateResponse](#api-ProbationHistoricalRateResponse)
    - [Productivity](#api-Productivity)
    - [ProductivityRequest](#api-ProductivityRequest)
    - [ProductivityResponse](#api-ProductivityResponse)
    - [Reward](#api-Reward)
    - [RewardRequest](#api-RewardRequest)
    - [RewardResponse](#api-RewardResponse)
    - [StakingRequest](#api-StakingRequest)
    - [StakingResponse](#api-StakingResponse)
    - [StakingResponse.StakingInfo](#api-StakingResponse-StakingInfo)
  
    - [PaidToDelegatesRequest.Schedule](#api-PaidToDelegatesRequest-Schedule)
  
    - [DelegateService](#api-DelegateService)
  
- [api_exit_queue.proto](#api_exit_queue-proto)
    - [ExitQueueEntry](#api-ExitQueueEntry)
    - [GetExitQueueRequest](#api-GetExitQueueRequest)
    - [GetExitQueueResponse](#api-GetExitQueueResponse)
  
    - [ExitQueueService](#api-ExitQueueService)
  
- [api_hermes.proto](#api_hermes-proto)
    - [BucketRewardDistribution](#api-BucketRewardDistribution)
    - [HermesAverageStatsRequest](#api-HermesAverageStatsRequest)
    - [HermesAverageStatsResponse](#api-HermesAverageStatsResponse)
    - [HermesAverageStatsResponse.AveragePerEpoch](#api-HermesAverageStatsResponse-AveragePerEpoch)
    - [HermesBucketDistribution](#api-HermesBucketDistribution)
    - [HermesBucketResponse](#api-HermesBucketResponse)
    - [HermesByDelegateDistributionRatio](#api-HermesByDelegateDistributionRatio)
    - [HermesByDelegateRequest](#api-HermesByDelegateRequest)
    - [HermesByDelegateResponse](#api-HermesByDelegateResponse)
    - [HermesByDelegateVoterInfo](#api-HermesByDelegateVoterInfo)
    - [HermesByVoterRequest](#api-HermesByVoterRequest)
    - [HermesByVoterResponse](#api-HermesByVoterResponse)
    - [HermesByVoterResponse.Delegate](#api-HermesByVoterResponse-Delegate)
    - [HermesDistribution](#api-HermesDistribution)
    - [HermesDropRecordsRequest](#api-HermesDropRecordsRequest)
    - [HermesDropRecordsResponse](#api-HermesDropRecordsResponse)
    - [HermesMetaRequest](#api-HermesMetaRequest)
    - [HermesMetaResponse](#api-HermesMetaResponse)
    - [HermesRequest](#api-HermesRequest)
    - [HermesResponse](#api-HermesResponse)
    - [RewardDistribution](#api-RewardDistribution)
  
    - [HermesService](#api-HermesService)
  
- [api_staking.proto](#api_staking-proto)
    - [BucketByIDRequest](#api-BucketByIDRequest)
    - [BucketByIDResponse](#api-BucketByIDResponse)
    - [BucketInfoEx](#api-BucketInfoEx)
    - [CandidateVoteByHeightRequest](#api-CandidateVoteByHeightRequest)
    - [CandidateVoteByHeightResponse](#api-CandidateVoteByHeightResponse)
    - [GetBucketByBucketIdRequest](#api-GetBucketByBucketIdRequest)
    - [GetBucketByBucketIdResponse](#api-GetBucketByBucketIdResponse)
    - [GetBucketListRequest](#api-GetBucketListRequest)
    - [GetBucketListResponse](#api-GetBucketListResponse)
    - [GetBucketsByBucketIdRequest](#api-GetBucketsByBucketIdRequest)
    - [GetBucketsByBucketIdResponse](#api-GetBucketsByBucketIdResponse)
    - [GetNativeBucketsRequest](#api-GetNativeBucketsRequest)
    - [GetNativeBucketsResponse](#api-GetNativeBucketsResponse)
    - [StakingBucketInfo](#api-StakingBucketInfo)
    - [VoteByHeightRequest](#api-VoteByHeightRequest)
    - [VoteByHeightResponse](#api-VoteByHeightResponse)
  
    - [StakingService](#api-StakingService)
  
- [api_stream.proto](#api_stream-proto)
    - [SupplyRequest](#api-SupplyRequest)
    - [SupplyResponse](#api-SupplyResponse)
  
    - [StreamService](#api-StreamService)
  
- [api_voting.proto](#api_voting-proto)
    - [CandidateInfoRequest](#api-CandidateInfoRequest)
    - [CandidateInfoResponse](#api-CandidateInfoResponse)
    - [CandidateInfoResponse.CandidateInfo](#api-CandidateInfoResponse-CandidateInfo)
    - [CandidateInfoResponse.Candidates](#api-CandidateInfoResponse-Candidates)
    - [CurrentDelegateInfo](#api-CurrentDelegateInfo)
    - [GetCurrentDelegatesRequest](#api-GetCurrentDelegatesRequest)
    - [GetCurrentDelegatesResponse](#api-GetCurrentDelegatesResponse)
    - [RewardSourcesRequest](#api-RewardSourcesRequest)
    - [RewardSourcesResponse](#api-RewardSourcesResponse)
    - [RewardSourcesResponse.DelegateDistributions](#api-RewardSourcesResponse-DelegateDistributions)
    - [VotingMetaRequest](#api-VotingMetaRequest)
    - [VotingMetaResponse](#api-VotingMetaResponse)
    - [VotingMetaResponse.CandidateMeta](#api-VotingMetaResponse-CandidateMeta)
  
    - [VotingService](#api-VotingService)
  
- [api_xrc20.proto](#api_xrc20-proto)
    - [GetXRC20HoldersByContractRequest](#api-GetXRC20HoldersByContractRequest)
    - [GetXRC20HoldersByContractResponse](#api-GetXRC20HoldersByContractResponse)
    - [GetXRC20StatsRequest](#api-GetXRC20StatsRequest)
    - [GetXRC20StatsResponse](#api-GetXRC20StatsResponse)
    - [GetXRC20TokenBalanceRequest](#api-GetXRC20TokenBalanceRequest)
    - [GetXRC20TokenBalanceResponse](#api-GetXRC20TokenBalanceResponse)
    - [GetXRC20TransfersByContractRequest](#api-GetXRC20TransfersByContractRequest)
    - [GetXRC20TransfersByContractResponse](#api-GetXRC20TransfersByContractResponse)
    - [XRC20AddressesRequest](#api-XRC20AddressesRequest)
    - [XRC20AddressesResponse](#api-XRC20AddressesResponse)
    - [XRC20ByAddressRequest](#api-XRC20ByAddressRequest)
    - [XRC20ByAddressResponse](#api-XRC20ByAddressResponse)
    - [XRC20ByContractAddressRequest](#api-XRC20ByContractAddressRequest)
    - [XRC20ByContractAddressResponse](#api-XRC20ByContractAddressResponse)
    - [XRC20ByPageRequest](#api-XRC20ByPageRequest)
    - [XRC20ByPageResponse](#api-XRC20ByPageResponse)
    - [XRC20HolderInfo](#api-XRC20HolderInfo)
    - [XRC20StatsItem](#api-XRC20StatsItem)
    - [XRC20TokenHolderAddressesRequest](#api-XRC20TokenHolderAddressesRequest)
    - [XRC20TokenHolderAddressesResponse](#api-XRC20TokenHolderAddressesResponse)
    - [XRC20TransferInfo](#api-XRC20TransferInfo)
    - [Xrc20Action](#api-Xrc20Action)
  
    - [XRC20Service](#api-XRC20Service)
  
- [api_xrc721.proto](#api_xrc721-proto)
    - [GetNFTHoldersByContractRequest](#api-GetNFTHoldersByContractRequest)
    - [GetNFTHoldersByContractResponse](#api-GetNFTHoldersByContractResponse)
    - [GetNFTTransferListRequest](#api-GetNFTTransferListRequest)
    - [GetNFTTransferListResponse](#api-GetNFTTransferListResponse)
    - [NFTHolderInfo](#api-NFTHolderInfo)
    - [NFTTransferInfo](#api-NFTTransferInfo)
    - [XRC721AddressesRequest](#api-XRC721AddressesRequest)
    - [XRC721AddressesResponse](#api-XRC721AddressesResponse)
    - [XRC721ByAddressRequest](#api-XRC721ByAddressRequest)
    - [XRC721ByAddressResponse](#api-XRC721ByAddressResponse)
    - [XRC721ByContractAddressRequest](#api-XRC721ByContractAddressRequest)
    - [XRC721ByContractAddressResponse](#api-XRC721ByContractAddressResponse)
    - [XRC721ByPageRequest](#api-XRC721ByPageRequest)
    - [XRC721ByPageResponse](#api-XRC721ByPageResponse)
    - [XRC721TokenHolderAddressesRequest](#api-XRC721TokenHolderAddressesRequest)
    - [XRC721TokenHolderAddressesResponse](#api-XRC721TokenHolderAddressesResponse)
    - [Xrc721Action](#api-Xrc721Action)
  
    - [XRC721Service](#api-XRC721Service)
  
- [include/pagination.proto](#include_pagination-proto)
    - [Pagination](#pagination-Pagination)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_account-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_account.proto



<a name="api-AccountBalanceRow"></a>

### AccountBalanceRow



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| balance | [string](#string) |  | in_flow - out_flow as decimal string (rau) |
| total_actions | [int64](#int64) |  | in_num_actions &#43; out_num_actions |






<a name="api-AccountMetaInfo"></a>

### AccountMetaInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| is_contract | [bool](#bool) |  |  |
| block_height | [uint64](#uint64) |  |  |
| contract_bytecode_hash | [string](#string) |  |  |






<a name="api-ActiveAccountsRequest"></a>

### ActiveAccountsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint64](#uint64) |  | number of account addresses to be queried for active accounts |






<a name="api-ActiveAccountsResponse"></a>

### ActiveAccountsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| activeAccounts | [string](#string) | repeated | list of account addresses |






<a name="api-AliasRequest"></a>

### AliasRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operatorAddress | [string](#string) |  | delegate&#39;s operator address |






<a name="api-AliasResponse"></a>

### AliasResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether the operator address exists |
| aliasName | [string](#string) |  | delegate&#39;s alias name |






<a name="api-AuthorizationHistoryEntry"></a>

### AuthorizationHistoryEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action_hash | [string](#string) |  |  |
| block_height | [uint64](#uint64) |  |  |
| chain_id | [string](#string) |  |  |
| address | [string](#string) |  | delegated contract |
| nonce | [string](#string) |  |  |
| y_parity | [string](#string) |  |  |
| authority | [string](#string) |  | recovered signer (= the account queried) |
| valid | [bool](#bool) |  | whether this authorization was accepted by the chain at inclusion time |






<a name="api-ContractInfoRequest"></a>

### ContractInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contractAddress | [string](#string) | repeated | contract address |






<a name="api-ContractInfoResponse"></a>

### ContractInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contracts | [ContractInfoResponse.Contract](#api-ContractInfoResponse-Contract) | repeated |  |






<a name="api-ContractInfoResponse-Contract"></a>

### ContractInfoResponse.Contract



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether the contract address exists |
| deployer | [string](#string) |  | contract creator |
| createTime | [string](#string) |  | contract create time |
| callTimes | [uint64](#uint64) |  | contract call times |
| accumulatedGas | [string](#string) |  | accumulated transaction fee |
| contractAddress | [string](#string) |  | contract address |






<a name="api-Erc20TokenBalanceByHeightRequest"></a>

### Erc20TokenBalanceByHeightRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) | repeated |  |
| height | [uint64](#uint64) |  |  |
| contract_address | [string](#string) |  |  |






<a name="api-Erc20TokenBalanceByHeightResponse"></a>

### Erc20TokenBalanceByHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  |  |
| contract_address | [string](#string) |  |  |
| balance | [string](#string) | repeated |  |
| decimals | [uint64](#uint64) |  |  |






<a name="api-GetAccountMetaRequest"></a>

### GetAccountMetaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| addresses | [string](#string) | repeated |  |






<a name="api-GetAccountMetaResponse"></a>

### GetAccountMetaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accounts | [AccountMetaInfo](#api-AccountMetaInfo) | repeated |  |






<a name="api-GetAddressNFTBalancesRequest"></a>

### GetAddressNFTBalancesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |






<a name="api-GetAddressNFTBalancesResponse"></a>

### GetAddressNFTBalancesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balances | [NFTBalanceInfo](#api-NFTBalanceInfo) | repeated |  |






<a name="api-GetAddressTokenBalancesRequest"></a>

### GetAddressTokenBalancesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |






<a name="api-GetAddressTokenBalancesResponse"></a>

### GetAddressTokenBalancesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balances | [TokenBalanceInfo](#api-TokenBalanceInfo) | repeated |  |






<a name="api-GetAuthorizationsByAuthorityRequest"></a>

### GetAuthorizationsByAuthorityRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authority | [string](#string) |  |  |
| skip | [int32](#int32) |  |  |
| first | [int32](#int32) |  |  |






<a name="api-GetAuthorizationsByAuthorityResponse"></a>

### GetAuthorizationsByAuthorityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authorizations | [AuthorizationHistoryEntry](#api-AuthorizationHistoryEntry) | repeated |  |
| count | [int64](#int64) |  |  |






<a name="api-GetContractCreateInfoRequest"></a>

### GetContractCreateInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |






<a name="api-GetContractCreateInfoResponse"></a>

### GetContractCreateInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action_hash | [string](#string) |  |  |
| creator | [string](#string) |  |  |






<a name="api-GetTopAccountsByBalanceRequest"></a>

### GetTopAccountsByBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int64](#int64) |  |  |
| offset | [int64](#int64) |  |  |






<a name="api-GetTopAccountsByBalanceResponse"></a>

### GetTopAccountsByBalanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [int64](#int64) |  |  |
| accounts | [AccountBalanceRow](#api-AccountBalanceRow) | repeated |  |






<a name="api-GetTopAccountsRequest"></a>

### GetTopAccountsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |
| stake_amount | [string](#string) |  | &#34;more&#34; (&gt;= 10000 IOTX) or &#34;less&#34; |
| stake_duration | [string](#string) |  | &#34;more&#34; (&gt;= 91 days) or &#34;less&#34; |
| mf | [string](#string) |  | &#34;hold&#34; (mf &gt; 0) or &#34;&#34; (mf = 0) |






<a name="api-GetTopAccountsResponse"></a>

### GetTopAccountsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [int64](#int64) |  |  |
| accounts | [TopAccountRow](#api-TopAccountRow) | repeated |  |






<a name="api-IotexBalanceByHeightRequest"></a>

### IotexBalanceByHeightRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) | repeated | address lists |
| height | [uint64](#uint64) |  | block height |






<a name="api-IotexBalanceByHeightResponse"></a>

### IotexBalanceByHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  | block height |
| balance | [string](#string) | repeated | balance at the given height. |






<a name="api-NFTBalanceInfo"></a>

### NFTBalanceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  |  |
| type | [string](#string) |  | &#34;xrc721&#34; or &#34;xrc1155&#34; |
| balance | [string](#string) |  |  |






<a name="api-OperatorAddressRequest"></a>

### OperatorAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| aliasName | [string](#string) |  | delegate&#39;s alias name |






<a name="api-OperatorAddressResponse"></a>

### OperatorAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether the alias name exists |
| operatorAddress | [string](#string) |  | operator address associated with the given alias name |






<a name="api-TokenBalanceInfo"></a>

### TokenBalanceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  |  |
| balance | [string](#string) |  |  |






<a name="api-TopAccountRow"></a>

### TopAccountRow



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner_address | [string](#string) |  |  |
| bucket_id | [uint64](#uint64) |  |  |
| staked_amount | [string](#string) |  |  |
| duration | [string](#string) |  |  |
| mf | [string](#string) |  |  |
| last_update | [string](#string) |  |  |
| balance | [string](#string) |  |  |






<a name="api-TotalAccountSupplyRequest"></a>

### TotalAccountSupplyRequest







<a name="api-TotalAccountSupplyResponse"></a>

### TotalAccountSupplyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| totalAccountSupply | [string](#string) |  | total amount of tokens held by IoTeX accounts |






<a name="api-TotalNumberOfHoldersRequest"></a>

### TotalNumberOfHoldersRequest







<a name="api-TotalNumberOfHoldersResponse"></a>

### TotalNumberOfHoldersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| totalNumberOfHolders | [uint64](#uint64) |  | total number of IOTX holders so far |





 

 

 


<a name="api-AccountService"></a>

### AccountService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| IotexBalanceByHeight | [IotexBalanceByHeightRequest](#api-IotexBalanceByHeightRequest) | [IotexBalanceByHeightResponse](#api-IotexBalanceByHeightResponse) | IotexBalanceByHeight returns the balance of the given address at the given height. |
| Erc20TokenBalanceByHeight | [Erc20TokenBalanceByHeightRequest](#api-Erc20TokenBalanceByHeightRequest) | [Erc20TokenBalanceByHeightResponse](#api-Erc20TokenBalanceByHeightResponse) |  |
| ActiveAccounts | [ActiveAccountsRequest](#api-ActiveAccountsRequest) | [ActiveAccountsResponse](#api-ActiveAccountsResponse) | ActiveAccounts lists most recently active accounts |
| OperatorAddress | [OperatorAddressRequest](#api-OperatorAddressRequest) | [OperatorAddressResponse](#api-OperatorAddressResponse) | OperatorAddress finds the delegate&#39;s operator address given the delegate&#39;s alias name |
| Alias | [AliasRequest](#api-AliasRequest) | [AliasResponse](#api-AliasResponse) | Alias finds the delegate&#39;s alias name given the delegate&#39;s operator address |
| TotalNumberOfHolders | [TotalNumberOfHoldersRequest](#api-TotalNumberOfHoldersRequest) | [TotalNumberOfHoldersResponse](#api-TotalNumberOfHoldersResponse) | TotalNumberOfHolders returns total number of IOTX holders so far |
| TotalAccountSupply | [TotalAccountSupplyRequest](#api-TotalAccountSupplyRequest) | [TotalAccountSupplyResponse](#api-TotalAccountSupplyResponse) | TotalAccountSupply returns total amount of tokens held by IoTeX accounts |
| ContractInfo | [ContractInfoRequest](#api-ContractInfoRequest) | [ContractInfoResponse](#api-ContractInfoResponse) | ContractInfo returns contract info by address, include contract creator, contract create time, contract call times, accumulated transaction fee |
| GetAccountMeta | [GetAccountMetaRequest](#api-GetAccountMetaRequest) | [GetAccountMetaResponse](#api-GetAccountMetaResponse) | GetAccountMeta returns account metadata (is_contract, block_height, bytecode hash) |
| GetContractCreateInfo | [GetContractCreateInfoRequest](#api-GetContractCreateInfoRequest) | [GetContractCreateInfoResponse](#api-GetContractCreateInfoResponse) | GetContractCreateInfo returns the action hash and creator for a contract |
| GetAddressNFTBalances | [GetAddressNFTBalancesRequest](#api-GetAddressNFTBalancesRequest) | [GetAddressNFTBalancesResponse](#api-GetAddressNFTBalancesResponse) | GetAddressNFTBalances returns NFT token balances for an address |
| GetAddressTokenBalances | [GetAddressTokenBalancesRequest](#api-GetAddressTokenBalancesRequest) | [GetAddressTokenBalancesResponse](#api-GetAddressTokenBalancesResponse) | GetAddressTokenBalances returns ERC20 token balances for an address |
| GetTopAccounts | [GetTopAccountsRequest](#api-GetTopAccountsRequest) | [GetTopAccountsResponse](#api-GetTopAccountsResponse) | GetTopAccounts returns top stakers from stats_top_list_view with filters |
| GetTopAccountsByBalance | [GetTopAccountsByBalanceRequest](#api-GetTopAccountsByBalanceRequest) | [GetTopAccountsByBalanceResponse](#api-GetTopAccountsByBalanceResponse) | GetTopAccountsByBalance returns top accounts by IOTX balance from account_income_count |
| GetAuthorizationsByAuthority | [GetAuthorizationsByAuthorityRequest](#api-GetAuthorizationsByAuthorityRequest) | [GetAuthorizationsByAuthorityResponse](#api-GetAuthorizationsByAuthorityResponse) | GetAuthorizationsByAuthority returns EIP-7702 authorization history for an authority address |

 



<a name="api_action-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_action.proto



<a name="api-ActionByAddressRequest"></a>

### ActionByAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | sender address or recipient address |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |
| sender | [string](#string) |  | optional filter by sender |
| recipient | [string](#string) |  | optional filter by recipient |
| actionType | [string](#string) |  | optional filter by action type |
| startTime | [string](#string) |  | optional start time filter (ISO 8601) |
| endTime | [string](#string) |  | optional end time filter (ISO 8601) |






<a name="api-ActionByAddressResponse"></a>

### ActionByAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether actions exist for the given address |
| actions | [ActionInfo](#api-ActionInfo) | repeated |  |
| count | [uint64](#uint64) |  | total number of actions for the given address |






<a name="api-ActionByDatesRequest"></a>

### ActionByDatesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startDate | [uint64](#uint64) |  | start date in unix epoch time |
| endDate | [uint64](#uint64) |  | end date in unix epoch time |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-ActionByDatesResponse"></a>

### ActionByDatesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether actions exist within the time frame |
| actions | [ActionInfo](#api-ActionInfo) | repeated |  |
| count | [uint64](#uint64) |  | total number of actions within the time frame |






<a name="api-ActionByHashRequest"></a>

### ActionByHashRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actHash | [string](#string) |  | action hash |
| include_fields | [string](#string) | repeated | optional subset: action_type, input_data, logs, token_transfers, base_fee, stake_action; empty = none |






<a name="api-ActionByHashResponse"></a>

### ActionByHashResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether actions exist within the time frame |
| actionInfo | [ActionInfo](#api-ActionInfo) |  |  |
| evmTransfers | [ActionByHashResponse.EvmTransfers](#api-ActionByHashResponse-EvmTransfers) | repeated |  |
| action_type_info | [ActionByHashResponse.ActionTypeInfo](#api-ActionByHashResponse-ActionTypeInfo) |  |  |
| input_data | [string](#string) |  | hex-encoded execution data |
| logs | [ActionByHashResponse.ActionLog](#api-ActionByHashResponse-ActionLog) | repeated |  |
| token_transfers | [ActionByHashResponse.TokenTransfer](#api-ActionByHashResponse-TokenTransfer) | repeated |  |
| block_base_fee | [string](#string) |  |  |
| stake_action | [ActionByHashResponse.StakeAction](#api-ActionByHashResponse-StakeAction) |  |  |






<a name="api-ActionByHashResponse-ActionLog"></a>

### ActionByHashResponse.ActionLog



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| block_height | [uint64](#uint64) |  |  |
| address | [string](#string) |  |  |
| topic0 | [string](#string) |  |  |
| topic1 | [string](#string) |  |  |
| topic2 | [string](#string) |  |  |
| topic3 | [string](#string) |  |  |
| data | [bytes](#bytes) |  | raw log data (base64 in JSON) |
| action_hash | [string](#string) |  |  |
| index | [int64](#int64) |  |  |






<a name="api-ActionByHashResponse-ActionTypeInfo"></a>

### ActionByHashResponse.ActionTypeInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  |  |
| access_list | [string](#string) |  |  |
| gas_tip_cap | [string](#string) |  |  |
| gas_fee_cap | [string](#string) |  |  |
| blob_gas | [string](#string) |  |  |
| blob_fee_cap | [string](#string) |  |  |
| blob_hashes | [string](#string) |  |  |
| blob_gas_price | [string](#string) |  |  |
| authorization_list | [ActionByHashResponse.AuthorizationEntry](#api-ActionByHashResponse-AuthorizationEntry) | repeated | EIP-7702 |






<a name="api-ActionByHashResponse-AuthorizationEntry"></a>

### ActionByHashResponse.AuthorizationEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| chain_id | [string](#string) |  |  |
| address | [string](#string) |  | delegate contract |
| nonce | [string](#string) |  |  |
| y_parity | [string](#string) |  |  |
| r | [string](#string) |  |  |
| s | [string](#string) |  |  |
| authority | [string](#string) |  | recovered signer |
| valid | [bool](#bool) |  | whether this authorization was accepted by the chain at inclusion time |






<a name="api-ActionByHashResponse-EvmTransfers"></a>

### ActionByHashResponse.EvmTransfers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender address |
| recipient | [string](#string) |  | recipient address |
| amount | [string](#string) |  | amount transferred |






<a name="api-ActionByHashResponse-StakeAction"></a>

### ActionByHashResponse.StakeAction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bucket_id | [int64](#int64) |  |  |
| amount | [string](#string) |  |  |
| staked_amount | [string](#string) |  |  |
| duration | [string](#string) |  |  |
| auto_stake | [bool](#bool) |  |  |
| candidate | [string](#string) |  |  |
| act_type | [string](#string) |  |  |
| owner_address | [string](#string) |  |  |






<a name="api-ActionByHashResponse-TokenTransfer"></a>

### ActionByHashResponse.TokenTransfer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| contract_address | [string](#string) |  |  |
| sender | [string](#string) |  |  |
| recipient | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| type | [string](#string) |  | &#34;erc20&#34; or &#34;nft&#34; |






<a name="api-ActionByHeightRequest"></a>

### ActionByHeightRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  |  |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-ActionByHeightResponse"></a>

### ActionByHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  |  |
| count | [uint64](#uint64) |  |  |
| actions | [ActionInfo](#api-ActionInfo) | repeated |  |






<a name="api-ActionByTypeRequest"></a>

### ActionByTypeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | action type |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-ActionByTypeResponse"></a>

### ActionByTypeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether actions exist for the given type |
| actions | [ActionInfo](#api-ActionInfo) | repeated |  |
| count | [uint64](#uint64) |  | total number of actions for the given type |






<a name="api-ActionInfo"></a>

### ActionInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actHash | [string](#string) |  | action hash |
| blkHash | [string](#string) |  | block hash |
| actType | [string](#string) |  | action type |
| sender | [string](#string) |  | sender address |
| recipient | [string](#string) |  | recipient address |
| amount | [string](#string) |  | amount transferred |
| timestamp | [uint64](#uint64) |  | unix timestamp |
| gasFee | [string](#string) |  | gas fee |
| blkHeight | [uint64](#uint64) |  | block height |
| gasPrice | [string](#string) |  | gas price |
| gasLimit | [uint64](#uint64) |  | gas limit |
| gasConsumed | [uint64](#uint64) |  | gas consumed |
| nonce | [uint64](#uint64) |  | nonce |
| status | [uint64](#uint64) |  | execution status |
| contractAddress | [string](#string) |  | contract address |
| executionRevertMsg | [string](#string) |  | execution revert message |
| chainId | [uint64](#uint64) |  | chain id |
| methodName | [string](#string) |  | method name from action_execution &#43; method_bytes |






<a name="api-ActionListRequest"></a>

### ActionListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |
| start_block_height | [uint64](#uint64) |  |  |






<a name="api-ActionListResponse"></a>

### ActionListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  |  |
| count | [uint64](#uint64) |  |  |
| actions | [ActionInfo](#api-ActionInfo) | repeated |  |






<a name="api-ActionRequest"></a>

### ActionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| actHash | [string](#string) |  |  |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-ActionResponse"></a>

### ActionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  |  |
| count | [uint64](#uint64) |  |  |
| actionList | [ActionInfo](#api-ActionInfo) | repeated |  |
| evmTransferList | [EvmTransferInfo](#api-EvmTransferInfo) | repeated |  |
| xrcList | [XrcInfo](#api-XrcInfo) | repeated |  |






<a name="api-ContractInteractorsRequest"></a>

### ContractInteractorsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | contract address |
| startTime | [string](#string) |  | optional time filter start |






<a name="api-ContractInteractorsResponse"></a>

### ContractInteractorsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| senders | [string](#string) | repeated |  |






<a name="api-EvmTransferInfo"></a>

### EvmTransferInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actHash | [string](#string) |  |  |
| blkHash | [string](#string) |  |  |
| from | [string](#string) |  |  |
| to | [string](#string) |  |  |
| quantity | [string](#string) |  |  |
| blkHeight | [uint64](#uint64) |  |  |
| timestamp | [uint64](#uint64) |  |  |






<a name="api-EvmTransfersByAddressRequest"></a>

### EvmTransfersByAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | sender address or recipient address |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-EvmTransfersByAddressResponse"></a>

### EvmTransfersByAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether EVM transfers exist for the given address |
| count | [uint64](#uint64) |  | total number of EVM transfers for the given address |
| evmTransfers | [EvmTransfersByAddressResponse.EvmTransfer](#api-EvmTransfersByAddressResponse-EvmTransfer) | repeated |  |






<a name="api-EvmTransfersByAddressResponse-EvmTransfer"></a>

### EvmTransfersByAddressResponse.EvmTransfer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actHash | [string](#string) |  | action hash |
| blkHash | [string](#string) |  | block hash |
| sender | [string](#string) |  | sender address |
| recipient | [string](#string) |  | recipient address |
| amount | [string](#string) |  | amount transferred |
| blkHeight | [uint64](#uint64) |  | block height |
| timestamp | [uint64](#uint64) |  | unix timestamp |






<a name="api-GetInternalTxnsRequest"></a>

### GetInternalTxnsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-GetInternalTxnsResponse"></a>

### GetInternalTxnsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| txns | [InternalTxnInfo](#api-InternalTxnInfo) | repeated |  |
| count | [uint64](#uint64) |  |  |






<a name="api-GetStakingActionsByAddressRequest"></a>

### GetStakingActionsByAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner_address | [string](#string) |  |  |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-GetStakingActionsByAddressResponse"></a>

### GetStakingActionsByAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actions | [StakingActionInfo](#api-StakingActionInfo) | repeated |  |
| count | [uint64](#uint64) |  |  |






<a name="api-InternalTxnInfo"></a>

### InternalTxnInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| block_height | [uint64](#uint64) |  |  |
| action_hash | [string](#string) |  |  |
| type | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| sender | [string](#string) |  |  |
| recipient | [string](#string) |  |  |
| timestamp | [string](#string) |  |  |






<a name="api-StakingActionInfo"></a>

### StakingActionInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| block_height | [uint64](#uint64) |  |  |
| action_hash | [string](#string) |  |  |
| sender | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| action_type | [string](#string) |  |  |
| timestamp | [string](#string) |  |  |






<a name="api-XrcInfo"></a>

### XrcInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actHash | [string](#string) |  |  |
| from | [string](#string) |  |  |
| to | [string](#string) |  |  |
| quantity | [string](#string) |  |  |
| blkHeight | [uint64](#uint64) |  |  |
| timestamp | [uint64](#uint64) |  |  |
| contract | [string](#string) |  |  |





 

 

 


<a name="api-ActionService"></a>

### ActionService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ActionByVoter | [ActionRequest](#api-ActionRequest) | [ActionResponse](#api-ActionResponse) |  |
| GetXrc20ByAddress | [ActionRequest](#api-ActionRequest) | [ActionResponse](#api-ActionResponse) |  |
| ActionByDates | [ActionByDatesRequest](#api-ActionByDatesRequest) | [ActionByDatesResponse](#api-ActionByDatesResponse) | ActionByDates finds actions by dates |
| ActionByHash | [ActionByHashRequest](#api-ActionByHashRequest) | [ActionByHashResponse](#api-ActionByHashResponse) | ActionByHash finds actions by hash |
| ActionByAddress | [ActionByAddressRequest](#api-ActionByAddressRequest) | [ActionByAddressResponse](#api-ActionByAddressResponse) | ActionByAddress finds actions by address |
| ActionByType | [ActionByTypeRequest](#api-ActionByTypeRequest) | [ActionByTypeResponse](#api-ActionByTypeResponse) | ActionByType finds actions by action type |
| EvmTransfersByAddress | [EvmTransfersByAddressRequest](#api-EvmTransfersByAddressRequest) | [EvmTransfersByAddressResponse](#api-EvmTransfersByAddressResponse) | EvmTransfersByAddress finds EVM transfers by address |
| ActionList | [ActionListRequest](#api-ActionListRequest) | [ActionListResponse](#api-ActionListResponse) | ActionList returns paginated list of latest actions |
| ActionByHeight | [ActionByHeightRequest](#api-ActionByHeightRequest) | [ActionByHeightResponse](#api-ActionByHeightResponse) | ActionByHeight finds actions by block height |
| ContractInteractors | [ContractInteractorsRequest](#api-ContractInteractorsRequest) | [ContractInteractorsResponse](#api-ContractInteractorsResponse) | ContractInteractors returns distinct senders who interacted with a contract |
| GetInternalTxns | [GetInternalTxnsRequest](#api-GetInternalTxnsRequest) | [GetInternalTxnsResponse](#api-GetInternalTxnsResponse) | GetInternalTxns returns paginated EVM internal transactions (block_receipt_transactions type=execution) |
| GetStakingActionsByAddress | [GetStakingActionsByAddressRequest](#api-GetStakingActionsByAddressRequest) | [GetStakingActionsByAddressResponse](#api-GetStakingActionsByAddressResponse) | GetStakingActionsByAddress returns paginated staking actions for an owner address |

 



<a name="api_actions-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_actions.proto



<a name="api-ActionsRequest"></a>

### ActionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| height | [uint64](#uint64) |  |  |
| offset | [uint64](#uint64) |  |  |
| size | [uint64](#uint64) |  |  |
| sort | [string](#string) |  |  |






<a name="api-AllActionsByAddressResponse"></a>

### AllActionsByAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint64](#uint64) |  |  |
| results | [AllActionsByAddressResult](#api-AllActionsByAddressResult) | repeated |  |






<a name="api-AllActionsByAddressResult"></a>

### AllActionsByAddressResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actHash | [string](#string) |  |  |
| blkHeight | [uint64](#uint64) |  |  |
| sender | [string](#string) |  |  |
| recipient | [string](#string) |  |  |
| actType | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| timeStamp | [uint64](#uint64) |  |  |
| recordType | [AllActionsByAddressResult.RecordType](#api-AllActionsByAddressResult-RecordType) |  |  |






<a name="api-EvmTransferDetailListByAddressResponse"></a>

### EvmTransferDetailListByAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint64](#uint64) |  |  |
| results | [EvmTransferDetailResult](#api-EvmTransferDetailResult) | repeated |  |






<a name="api-EvmTransferDetailResult"></a>

### EvmTransferDetailResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actHash | [string](#string) |  |  |
| blkHeight | [uint64](#uint64) |  |  |
| sender | [string](#string) |  |  |
| recipient | [string](#string) |  |  |
| blkHash | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| timeStamp | [uint64](#uint64) |  |  |





 


<a name="api-AllActionsByAddressResult-RecordType"></a>

### AllActionsByAddressResult.RecordType


| Name | Number | Description |
| ---- | ------ | ----------- |
| NATIVE | 0 |  |
| XRC20 | 1 |  |
| XRC721 | 2 |  |
| EVMTRANSFER | 3 |  |


 

 


<a name="api-ActionsService"></a>

### ActionsService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetEvmTransferDetailListByAddress | [ActionsRequest](#api-ActionsRequest) | [EvmTransferDetailListByAddressResponse](#api-EvmTransferDetailListByAddressResponse) |  |
| GetAllActionsByAddress | [ActionsRequest](#api-ActionsRequest) | [AllActionsByAddressResponse](#api-AllActionsByAddressResponse) |  |

 



<a name="api_approval-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_approval.proto



<a name="api-GetXRC20ApprovalsRequest"></a>

### GetXRC20ApprovalsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner_address | [string](#string) |  |  |






<a name="api-GetXRC20ApprovalsResponse"></a>

### GetXRC20ApprovalsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| approvals | [XRC20ApprovalInfo](#api-XRC20ApprovalInfo) | repeated |  |






<a name="api-GetXRC721ApprovalsRequest"></a>

### GetXRC721ApprovalsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner_address | [string](#string) |  |  |






<a name="api-GetXRC721ApprovalsResponse"></a>

### GetXRC721ApprovalsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| approvals | [XRC721ApprovalInfo](#api-XRC721ApprovalInfo) | repeated |  |






<a name="api-XRC20ApprovalInfo"></a>

### XRC20ApprovalInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action_hash | [string](#string) |  |  |
| contract_address | [string](#string) |  |  |
| owner | [string](#string) |  |  |
| spender | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| timestamp | [string](#string) |  |  |






<a name="api-XRC721ApprovalInfo"></a>

### XRC721ApprovalInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action_hash | [string](#string) |  |  |
| contract_address | [string](#string) |  |  |
| owner | [string](#string) |  |  |
| approved | [string](#string) |  |  |
| token_id | [string](#string) |  |  |
| timestamp | [string](#string) |  |  |





 

 

 


<a name="api-ApprovalService"></a>

### ApprovalService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetXRC20Approvals | [GetXRC20ApprovalsRequest](#api-GetXRC20ApprovalsRequest) | [GetXRC20ApprovalsResponse](#api-GetXRC20ApprovalsResponse) |  |
| GetXRC721Approvals | [GetXRC721ApprovalsRequest](#api-GetXRC721ApprovalsRequest) | [GetXRC721ApprovalsResponse](#api-GetXRC721ApprovalsResponse) |  |

 



<a name="api_chain-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_chain.proto



<a name="api-ActionHistoryPoint"></a>

### ActionHistoryPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| date | [string](#string) |  | UTC datetime string |
| sum | [uint64](#uint64) |  | total actions in bucket |






<a name="api-BlockInfo"></a>

### BlockInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| block_height | [uint64](#uint64) |  | block height |
| block_hash | [string](#string) |  | block hash |
| producer_address | [string](#string) |  | producer address |
| num_actions | [uint64](#uint64) |  | number of actions |
| timestamp | [int64](#int64) |  | block timestamp (unix timestamp) |
| gas_consumed | [uint64](#uint64) |  | gas consumed |
| producer_name | [string](#string) |  | producer name |
| block_reward | [string](#string) |  | block reward |
| epoch_num | [uint64](#uint64) |  | epoch number |
| priority_bonus | [string](#string) |  | priority bonus |
| base_fee | [string](#string) |  | base fee |






<a name="api-BlockSizeByHeightRequest"></a>

### BlockSizeByHeightRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  | block height |






<a name="api-BlockSizeByHeightResponse"></a>

### BlockSizeByHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| blockSize | [double](#double) |  | size |
| serverVersion | [string](#string) |  | version |






<a name="api-ChainRequest"></a>

### ChainRequest







<a name="api-ChainResponse"></a>

### ChainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mostRecentEpoch | [uint64](#uint64) |  | most recent epoch |
| mostRecentBlockHeight | [uint64](#uint64) |  | most recent block height |
| totalSupply | [string](#string) |  | total supply |
| totalCirculatingSupply | [string](#string) |  | total circulating supply |
| totalCirculatingSupplyNoRewardPool | [string](#string) |  | total circulating supply no reward pool |
| votingResultMeta | [VotingResultMeta](#api-VotingResultMeta) |  | voting result meta |
| exactCirculatingSupply | [string](#string) |  | exact circulating supply |
| rewards | [ChainResponse.Rewards](#api-ChainResponse-Rewards) |  | rewards |






<a name="api-ChainResponse-Rewards"></a>

### ChainResponse.Rewards



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| totalBalance | [string](#string) |  | total balance |
| totalUnclaimed | [string](#string) |  | total unclaimed |
| totalAvailable | [string](#string) |  | total available |






<a name="api-GasHistoryPoint"></a>

### GasHistoryPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| date | [string](#string) |  |  |
| max_gas_price | [string](#string) |  |  |
| min_gas_price | [string](#string) |  |  |
| avg_gas_price | [string](#string) |  |  |
| total_gas_fee | [string](#string) |  |  |






<a name="api-GetActionHistoryRequest"></a>

### GetActionHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start_time | [string](#string) |  | UTC datetime string (e.g. &#34;2024-01-01 00:00:00&#34;) |
| end_time | [string](#string) |  | UTC datetime string |
| interval | [string](#string) |  | &#34;minute&#34;, &#34;hour&#34;, or &#34;day&#34; |






<a name="api-GetActionHistoryResponse"></a>

### GetActionHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [ActionHistoryPoint](#api-ActionHistoryPoint) | repeated |  |






<a name="api-GetBlockByHeightRequest"></a>

### GetBlockByHeightRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  | block height |






<a name="api-GetBlockByHeightResponse"></a>

### GetBlockByHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  |  |
| block | [BlockInfo](#api-BlockInfo) |  |  |






<a name="api-GetBlocksRequest"></a>

### GetBlocksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [uint64](#uint64) |  | page number (starts from 1); ignored when before_height &gt; 0 |
| limit | [uint64](#uint64) |  | number of blocks per page |
| before_height | [uint64](#uint64) |  | cursor: return blocks with height &lt;= before_height (0 = use page) |






<a name="api-GetBlocksResponse"></a>

### GetBlocksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| blocks | [BlockInfo](#api-BlockInfo) | repeated | list of blocks |






<a name="api-GetChainStatsRequest"></a>

### GetChainStatsRequest







<a name="api-GetChainStatsResponse"></a>

### GetChainStatsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actions_num | [uint64](#uint64) |  | total number of actions, from kv table |
| total_supply | [string](#string) |  | IOTX (rau / 1e18), decimal string |
| circulating_supply | [string](#string) |  | IOTX (rau / 1e18), decimal string |






<a name="api-GetEpochInfoRequest"></a>

### GetEpochInfoRequest







<a name="api-GetEpochInfoResponse"></a>

### GetEpochInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epoch_height | [uint64](#uint64) |  | first block height of the current epoch |
| epoch_num | [uint64](#uint64) |  | current epoch number |






<a name="api-GetGasHistoryRequest"></a>

### GetGasHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [string](#string) |  | YYYY-MM-DD |
| end | [string](#string) |  | YYYY-MM-DD (inclusive) |






<a name="api-GetGasHistoryResponse"></a>

### GetGasHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [GasHistoryPoint](#api-GasHistoryPoint) | repeated |  |






<a name="api-GetLatestBlockHeightRequest"></a>

### GetLatestBlockHeightRequest







<a name="api-GetLatestBlockHeightResponse"></a>

### GetLatestBlockHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  | latest block height |






<a name="api-GetLatestStakingRecordRequest"></a>

### GetLatestStakingRecordRequest







<a name="api-GetLatestStakingRecordResponse"></a>

### GetLatestStakingRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total_supply | [string](#string) |  | total supply |
| all_staking | [string](#string) |  | total staked IOTX (in rau) |
| staking_ratio | [string](#string) |  | staking ratio (decimal string) |






<a name="api-GetPeakTpsRequest"></a>

### GetPeakTpsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start_block_height | [uint64](#uint64) |  | 0 = scan all blocks; &gt; 0 = only look at blocks after this height |






<a name="api-GetPeakTpsResponse"></a>

### GetPeakTpsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| num_actions | [string](#string) |  | max(num_actions)/5 rounded to 2 decimal places |
| block_height | [uint64](#uint64) |  | current max block height (cursor for next call) |






<a name="api-GetStakingRatioHistoryRequest"></a>

### GetStakingRatioHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start_time | [string](#string) |  | optional start date (YYYY-MM-DD) |
| end_time | [string](#string) |  | optional end date (YYYY-MM-DD) |






<a name="api-GetStakingRatioHistoryResponse"></a>

### GetStakingRatioHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [StakingRatioPoint](#api-StakingRatioPoint) | repeated |  |






<a name="api-GetSupplyHistoryRequest"></a>

### GetSupplyHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [string](#string) |  | YYYY-MM-DD |
| end | [string](#string) |  | YYYY-MM-DD (inclusive) |






<a name="api-GetSupplyHistoryResponse"></a>

### GetSupplyHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [SupplyHistoryPoint](#api-SupplyHistoryPoint) | repeated |  |






<a name="api-GetTpsHistoryRequest"></a>

### GetTpsHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [string](#string) |  | start date, YYYY-MM-DD (UTC) |
| end | [string](#string) |  | end date, YYYY-MM-DD (UTC, inclusive) |






<a name="api-GetTpsHistoryResponse"></a>

### GetTpsHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [TpsHistoryPoint](#api-TpsHistoryPoint) | repeated |  |






<a name="api-MostRecentTPSRequest"></a>

### MostRecentTPSRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| blockWindow | [uint64](#uint64) |  | number of last blocks that are backtracked to compute TPS |






<a name="api-MostRecentTPSResponse"></a>

### MostRecentTPSResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mostRecentTPS | [double](#double) |  | latest transactions per second |






<a name="api-NumberOfActionsRequest"></a>

### NumberOfActionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | starting epoch number |
| epochCount | [uint64](#uint64) |  | epoch count |






<a name="api-NumberOfActionsResponse"></a>

### NumberOfActionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether the starting epoch number is less than the most recent epoch number |
| count | [uint64](#uint64) |  | number of actions |






<a name="api-StakingRatioPoint"></a>

### StakingRatioPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| date_time | [string](#string) |  |  |
| ratio | [string](#string) |  |  |






<a name="api-SupplyHistoryPoint"></a>

### SupplyHistoryPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| date | [string](#string) |  |  |
| total_supply | [string](#string) |  | IOTX |
| circulating_supply | [string](#string) |  | IOTX |
| burn | [string](#string) |  | IOTX, daily |
| issue | [string](#string) |  | IOTX, daily |






<a name="api-TotalTransferredTokensRequest"></a>

### TotalTransferredTokensRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | starting epoch number |
| epochCount | [uint64](#uint64) |  | epoch count |






<a name="api-TotalTransferredTokensResponse"></a>

### TotalTransferredTokensResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| totalTransferredTokens | [string](#string) |  | total tranferred tokens |






<a name="api-TpsHistoryPoint"></a>

### TpsHistoryPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| date | [string](#string) |  | YYYY-MM-DD |
| avg_tps | [double](#double) |  | avg(num_actions)/2.5 |
| max_tps | [double](#double) |  | max(num_actions)/2.5 |






<a name="api-VotingResultMeta"></a>

### VotingResultMeta



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| totalCandidates | [uint64](#uint64) |  | total candidates |
| totalWeightedVotes | [string](#string) |  | total weighted votes |
| votedTokens | [string](#string) |  | voted tokens |





 

 

 


<a name="api-ChainService"></a>

### ChainService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Chain | [ChainRequest](#api-ChainRequest) | [ChainResponse](#api-ChainResponse) |  |
| MostRecentTPS | [MostRecentTPSRequest](#api-MostRecentTPSRequest) | [MostRecentTPSResponse](#api-MostRecentTPSResponse) | MostRecentTPS gives the latest transactions per second |
| NumberOfActions | [NumberOfActionsRequest](#api-NumberOfActionsRequest) | [NumberOfActionsResponse](#api-NumberOfActionsResponse) | NumberOfActions gives the number of actions |
| TotalTransferredTokens | [TotalTransferredTokensRequest](#api-TotalTransferredTokensRequest) | [TotalTransferredTokensResponse](#api-TotalTransferredTokensResponse) | TotalTransferredTokens gives the amount of tokens transferred within a time frame |
| BlockSizeByHeight | [BlockSizeByHeightRequest](#api-BlockSizeByHeightRequest) | [BlockSizeByHeightResponse](#api-BlockSizeByHeightResponse) | BlockSizeByHeight gives the block size by height |
| GetLatestBlockHeight | [GetLatestBlockHeightRequest](#api-GetLatestBlockHeightRequest) | [GetLatestBlockHeightResponse](#api-GetLatestBlockHeightResponse) | GetLatestBlockHeight gives the latest block height |
| GetBlocks | [GetBlocksRequest](#api-GetBlocksRequest) | [GetBlocksResponse](#api-GetBlocksResponse) | GetBlocks gives a list of blocks with pagination |
| GetBlockByHeight | [GetBlockByHeightRequest](#api-GetBlockByHeightRequest) | [GetBlockByHeightResponse](#api-GetBlockByHeightResponse) | GetBlockByHeight returns a single block by its height |
| GetEpochInfo | [GetEpochInfoRequest](#api-GetEpochInfoRequest) | [GetEpochInfoResponse](#api-GetEpochInfoResponse) | GetEpochInfo returns the current epoch number and its starting block height |
| GetLatestStakingRecord | [GetLatestStakingRecordRequest](#api-GetLatestStakingRecordRequest) | [GetLatestStakingRecordResponse](#api-GetLatestStakingRecordResponse) | GetLatestStakingRecord returns the most recent staking statistics |
| GetPeakTps | [GetPeakTpsRequest](#api-GetPeakTpsRequest) | [GetPeakTpsResponse](#api-GetPeakTpsResponse) | GetPeakTps returns the all-time peak TPS (max block actions / 5-second block time) |
| GetActionHistory | [GetActionHistoryRequest](#api-GetActionHistoryRequest) | [GetActionHistoryResponse](#api-GetActionHistoryResponse) | GetActionHistory returns aggregated action counts over a time range |
| GetStakingRatioHistory | [GetStakingRatioHistoryRequest](#api-GetStakingRatioHistoryRequest) | [GetStakingRatioHistoryResponse](#api-GetStakingRatioHistoryResponse) | GetStakingRatioHistory returns staking ratio history over a time range |
| GetChainStats | [GetChainStatsRequest](#api-GetChainStatsRequest) | [GetChainStatsResponse](#api-GetChainStatsResponse) | GetChainStats returns total action count &#43; total/circulating supply (IOTX units) |
| GetTpsHistory | [GetTpsHistoryRequest](#api-GetTpsHistoryRequest) | [GetTpsHistoryResponse](#api-GetTpsHistoryResponse) | GetTpsHistory returns daily avg/max TPS over a date range |
| GetGasHistory | [GetGasHistoryRequest](#api-GetGasHistoryRequest) | [GetGasHistoryResponse](#api-GetGasHistoryResponse) | GetGasHistory returns daily gas price stats and total gas fee over a date range |
| GetSupplyHistory | [GetSupplyHistoryRequest](#api-GetSupplyHistoryRequest) | [GetSupplyHistoryResponse](#api-GetSupplyHistoryResponse) | GetSupplyHistory returns daily total/circulating supply (IOTX) and daily burn/issue over a date range |

 



<a name="api_delegate-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_delegate.proto



<a name="api-BookKeepingRequest"></a>

### BookKeepingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | epoch number to start from |
| epochCount | [uint64](#uint64) |  | number of epochs to query |
| delegateName | [string](#string) |  | name of the delegate |
| pagination | [pagination.Pagination](#pagination-Pagination) |  | Pagination info |
| epochRewardPerc | [uint64](#uint64) |  | percentage of the epoch reward to be paid to the delegate |
| blockRewardPerc | [uint64](#uint64) |  | percentage of the block reward to be paid to the delegate |
| foundationBonusPerc | [uint64](#uint64) |  | percentage of the foundation bonus to be paid to the delegate |






<a name="api-BookKeepingResponse"></a>

### BookKeepingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether the delegate has bookkeeping information within the specified epoch range |
| count | [uint64](#uint64) |  | total number of reward distributions |
| rewardDistribution | [DelegateRewardDistribution](#api-DelegateRewardDistribution) | repeated |  |






<a name="api-BucketInfo"></a>

### BucketInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| voterEthAddress | [string](#string) |  | voter’s ERC20 address |
| voterIotexAddress | [string](#string) |  | voter&#39;s IoTeX address |
| isNative | [bool](#bool) |  | whether the bucket is native |
| votes | [string](#string) |  | voter&#39;s votes |
| weightedVotes | [string](#string) |  | voter’s weighted votes |
| remainingDuration | [string](#string) |  | bucket remaining duration |
| startTime | [string](#string) |  | bucket start time |
| decay | [bool](#bool) |  | whether the vote weight decays |
| bucketID | [uint64](#uint64) |  | bucket id |






<a name="api-BucketInfoList"></a>

### BucketInfoList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epochNumber | [uint64](#uint64) |  | epoch number |
| count | [uint64](#uint64) |  | total number of buckets in the given epoch for the given delegate |
| bucketInfo | [BucketInfo](#api-BucketInfo) | repeated |  |






<a name="api-BucketInfoRequest"></a>

### BucketInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | Epoch number to start from |
| epochCount | [uint64](#uint64) |  | Number of epochs to query |
| delegateName | [string](#string) |  | Name of the delegate |
| pagination | [pagination.Pagination](#pagination-Pagination) |  | Pagination info |






<a name="api-BucketInfoResponse"></a>

### BucketInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether the delegate has voting bucket information within the specified epoch range |
| count | [uint64](#uint64) |  | total number of buckets in the given epoch for the given delegate |
| bucketInfoList | [BucketInfoList](#api-BucketInfoList) | repeated |  |






<a name="api-DelegateRewardDistribution"></a>

### DelegateRewardDistribution



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| voterEthAddress | [string](#string) |  | voter’s ERC20 address |
| voterIotexAddress | [string](#string) |  | voter’s IoTeX address |
| amount | [string](#string) |  | amount of reward distribution |






<a name="api-PaidToDelegatesRequest"></a>

### PaidToDelegatesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schedule | [PaidToDelegatesRequest.Schedule](#api-PaidToDelegatesRequest-Schedule) |  |  |
| date | [string](#string) |  |  |






<a name="api-PaidToDelegatesResponse"></a>

### PaidToDelegatesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegateInfo | [PaidToDelegatesResponse.DelegateInfo](#api-PaidToDelegatesResponse-DelegateInfo) | repeated |  |






<a name="api-PaidToDelegatesResponse-DelegateInfo"></a>

### PaidToDelegatesResponse.DelegateInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegateName | [string](#string) |  | delegate name |
| amount | [string](#string) |  | amount of reward distribution |
| blockReward | [string](#string) |  | amount of block rewards |
| epochReward | [string](#string) |  | amount of epoch rewards |
| foundationBonus | [string](#string) |  | amount of foundation bonus |






<a name="api-ProbationHistoricalRateRequest"></a>

### ProbationHistoricalRateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | starting epoch number |
| epochCount | [uint64](#uint64) |  | epoch count |
| delegateName | [string](#string) |  | candidate name |






<a name="api-ProbationHistoricalRateResponse"></a>

### ProbationHistoricalRateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| probationHistoricalRate | [string](#string) |  | probation historical rate |






<a name="api-Productivity"></a>

### Productivity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether the delegate has productivity information within the specified epoch range |
| production | [uint64](#uint64) |  | number of block productions |
| expectedProduction | [uint64](#uint64) |  | number of expected block productions |






<a name="api-ProductivityRequest"></a>

### ProductivityRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | starting epoch number |
| epochCount | [uint64](#uint64) |  | epoch count |
| delegateName | [string](#string) |  | producer name |






<a name="api-ProductivityResponse"></a>

### ProductivityResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| productivity | [Productivity](#api-Productivity) |  |  |






<a name="api-Reward"></a>

### Reward



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| blockReward | [string](#string) |  | amount of block rewards |
| epochReward | [string](#string) |  | amount of epoch rewards |
| foundationBonus | [string](#string) |  | amount of foundation bonus |
| exist | [bool](#bool) |  | whether the delegate has reward information within the specified epoch range |






<a name="api-RewardRequest"></a>

### RewardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | Epoch number to start from |
| epochCount | [uint64](#uint64) |  | Number of epochs to query |
| delegateName | [string](#string) |  | Name of the delegate |






<a name="api-RewardResponse"></a>

### RewardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reward | [Reward](#api-Reward) |  |  |






<a name="api-StakingRequest"></a>

### StakingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | starting epoch number |
| epochCount | [uint64](#uint64) |  | epoch count |
| delegateName | [string](#string) |  | candidate name |






<a name="api-StakingResponse"></a>

### StakingResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether the delegate has staking information within the specified epoch range |
| stakingInfo | [StakingResponse.StakingInfo](#api-StakingResponse-StakingInfo) | repeated |  |






<a name="api-StakingResponse-StakingInfo"></a>

### StakingResponse.StakingInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epochNumber | [uint64](#uint64) |  | epoch number |
| totalStaking | [string](#string) |  | total staking amount |
| selfStaking | [string](#string) |  | candidate’s self-staking amount |





 


<a name="api-PaidToDelegatesRequest-Schedule"></a>

### PaidToDelegatesRequest.Schedule


| Name | Number | Description |
| ---- | ------ | ----------- |
| MONTHLY | 0 |  |
| DAILY | 1 |  |


 

 


<a name="api-DelegateService"></a>

### DelegateService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| BucketInfo | [BucketInfoRequest](#api-BucketInfoRequest) | [BucketInfoResponse](#api-BucketInfoResponse) | BucketInfo provides voting bucket detail information for candidates within a range of epochs |
| BookKeeping | [BookKeepingRequest](#api-BookKeepingRequest) | [BookKeepingResponse](#api-BookKeepingResponse) | BookKeeping gives delegates an overview of the reward distributions to their voters within a range of epochs |
| Productivity | [ProductivityRequest](#api-ProductivityRequest) | [ProductivityResponse](#api-ProductivityResponse) | Productivity gives block productivity of producers within a range of epochs |
| Reward | [RewardRequest](#api-RewardRequest) | [RewardResponse](#api-RewardResponse) | Rewards provides reward detail information for candidates within a range of epochs |
| Staking | [StakingRequest](#api-StakingRequest) | [StakingResponse](#api-StakingResponse) | Staking provides staking information for candidates within a range of epochs |
| ProbationHistoricalRate | [ProbationHistoricalRateRequest](#api-ProbationHistoricalRateRequest) | [ProbationHistoricalRateResponse](#api-ProbationHistoricalRateResponse) | ProbationHistoricalRate provides the rate of probation for a given delegate |
| PaidToDelegates | [PaidToDelegatesRequest](#api-PaidToDelegatesRequest) | [PaidToDelegatesResponse](#api-PaidToDelegatesResponse) | PaidToDelegates provides the amount of rewards paid to delegates |

 



<a name="api_exit_queue-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_exit_queue.proto



<a name="api-ExitQueueEntry"></a>

### ExitQueueEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| candidate_name | [string](#string) |  |  |
| candidate_identity | [string](#string) |  |  |
| status | [string](#string) |  | &#34;requested&#34;, &#34;scheduled&#34;, &#34;confirmed&#34; |
| request_height | [uint64](#uint64) |  |  |
| request_hash | [string](#string) |  |  |
| schedule_height | [uint64](#uint64) |  |  |
| schedule_hash | [string](#string) |  |  |
| confirm_height | [uint64](#uint64) |  |  |
| confirm_hash | [string](#string) |  |  |
| scheduled_at | [uint64](#uint64) |  | block height when exit will execute |






<a name="api-GetExitQueueRequest"></a>

### GetExitQueueRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| skip | [int64](#int64) |  |  |
| first | [int64](#int64) |  |  |
| statuses | [string](#string) | repeated | optional filter; if empty, returns all statuses. Each value must be one of &#34;requested&#34;, &#34;scheduled&#34;, or &#34;confirmed&#34;. Multiple values are OR-ed. |
| candidate_identity | [string](#string) |  | optional filter; if empty, returns entries for all candidates. When set, matches the candidate_identity column exactly (0x... lowercased hex, matching how the indexer writes it). |






<a name="api-GetExitQueueResponse"></a>

### GetExitQueueResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exits | [ExitQueueEntry](#api-ExitQueueEntry) | repeated |  |
| count | [int64](#int64) |  | total matching rows |





 

 

 


<a name="api-ExitQueueService"></a>

### ExitQueueService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetExitQueue | [GetExitQueueRequest](#api-GetExitQueueRequest) | [GetExitQueueResponse](#api-GetExitQueueResponse) |  |

 



<a name="api_hermes-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_hermes.proto



<a name="api-BucketRewardDistribution"></a>

### BucketRewardDistribution



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| voterEthAddress | [string](#string) |  | voter’s ERC20 address |
| voterIotexAddress | [string](#string) |  | voter’s IoTeX address |
| bucketID | [uint64](#uint64) |  | voter&#39;s bucketID |
| amount | [string](#string) |  | amount of reward distribution |






<a name="api-HermesAverageStatsRequest"></a>

### HermesAverageStatsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | starting epoch number |
| epochCount | [uint64](#uint64) |  | epoch count |
| rewardAddress | [string](#string) |  | Name of reward address |






<a name="api-HermesAverageStatsResponse"></a>

### HermesAverageStatsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether Hermes has bookkeeping information within the specified epoch range |
| averagePerEpoch | [HermesAverageStatsResponse.AveragePerEpoch](#api-HermesAverageStatsResponse-AveragePerEpoch) | repeated |  |






<a name="api-HermesAverageStatsResponse-AveragePerEpoch"></a>

### HermesAverageStatsResponse.AveragePerEpoch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegateName | [string](#string) |  | delegate name |
| rewardDistribution | [string](#string) |  | reward distribution amount on average |
| totalWeightedVotes | [string](#string) |  | total weighted votes on average |






<a name="api-HermesBucketDistribution"></a>

### HermesBucketDistribution



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegateName | [string](#string) |  | delegate name |
| bucketRewardDistribution | [BucketRewardDistribution](#api-BucketRewardDistribution) | repeated |  |
| stakingIotexAddress | [string](#string) |  | delegate IoTeX staking address |
| voterCount | [uint64](#uint64) |  | number of voters |
| waiveServiceFee | [bool](#bool) |  | whether the delegate is qualified for waiving the service fee |
| refund | [string](#string) |  | amount of refund |






<a name="api-HermesBucketResponse"></a>

### HermesBucketResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hermesBucketDistribution | [HermesBucketDistribution](#api-HermesBucketDistribution) | repeated |  |






<a name="api-HermesByDelegateDistributionRatio"></a>

### HermesByDelegateDistributionRatio



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epochNumber | [uint64](#uint64) |  | epoch number |
| blockRewardRatio | [double](#double) |  | ratio of block reward being distributed |
| epochRewardRatio | [double](#double) |  | ratio of epoch reward being distributed |
| foundationBonusRatio | [double](#double) |  | ratio of foundation bonus being distributed |






<a name="api-HermesByDelegateRequest"></a>

### HermesByDelegateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | Epoch number to start from |
| epochCount | [uint64](#uint64) |  | Number of epochs to query |
| delegateName | [string](#string) |  | Name of the delegate |
| pagination | [pagination.Pagination](#pagination-Pagination) |  | Pagination info |






<a name="api-HermesByDelegateResponse"></a>

### HermesByDelegateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether the delegate has hermes information within the specified epoch range |
| count | [uint64](#uint64) |  | total number of reward distributions |
| voterInfoList | [HermesByDelegateVoterInfo](#api-HermesByDelegateVoterInfo) | repeated |  |
| totalRewardsDistributed | [string](#string) |  | total reward amount distributed |
| distributionRatio | [HermesByDelegateDistributionRatio](#api-HermesByDelegateDistributionRatio) | repeated |  |






<a name="api-HermesByDelegateVoterInfo"></a>

### HermesByDelegateVoterInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| voterAddress | [string](#string) |  | voter address |
| fromEpoch | [uint64](#uint64) |  | starting epoch |
| toEpoch | [uint64](#uint64) |  | ending epoch |
| amount | [string](#string) |  | distributino amount |
| actHash | [string](#string) |  | action hash |
| timestamp | [uint64](#uint64) |  | unix timestamp |






<a name="api-HermesByVoterRequest"></a>

### HermesByVoterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | Start epoch number |
| epochCount | [uint64](#uint64) |  | Number of epochs to query |
| voterAddress | [string](#string) |  | voter address |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-HermesByVoterResponse"></a>

### HermesByVoterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether the voter uses Hermes within the specified epoch range |
| delegates | [HermesByVoterResponse.Delegate](#api-HermesByVoterResponse-Delegate) | repeated |  |
| count | [uint64](#uint64) |  | total number of reward receivings |
| totalRewardReceived | [string](#string) |  | total reward amount received |






<a name="api-HermesByVoterResponse-Delegate"></a>

### HermesByVoterResponse.Delegate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegateName | [string](#string) |  | delegate name |
| fromEpoch | [uint64](#uint64) |  | starting epoch of bookkeeping |
| toEpoch | [uint64](#uint64) |  | ending epoch of bookkeeping |
| amount | [string](#string) |  | receiving amount |
| actHash | [string](#string) |  | action hash |
| timestamp | [uint64](#uint64) |  | unix timestamp |






<a name="api-HermesDistribution"></a>

### HermesDistribution



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegateName | [string](#string) |  | delegate name |
| rewardDistribution | [RewardDistribution](#api-RewardDistribution) | repeated |  |
| stakingIotexAddress | [string](#string) |  | delegate IoTeX staking address |
| voterCount | [uint64](#uint64) |  | number of voters |
| waiveServiceFee | [bool](#bool) |  | whether the delegate is qualified for waiving the service fee |
| refund | [string](#string) |  | amount of refund |






<a name="api-HermesDropRecordsRequest"></a>

### HermesDropRecordsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epochNumber | [uint64](#uint64) |  | end epoch number |
| delegateName | [string](#string) |  | delegate name |
| voterAddress | [string](#string) |  | Name of voter address |
| actHash | [string](#string) |  | Name of actHash |
| bucketID | [uint64](#uint64) |  | bucket ID |
| amount | [string](#string) |  | Name of amount |






<a name="api-HermesDropRecordsResponse"></a>

### HermesDropRecordsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | whether the drop records are successfully generated |






<a name="api-HermesMetaRequest"></a>

### HermesMetaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | starting epoch number |
| epochCount | [uint64](#uint64) |  | epoch count |






<a name="api-HermesMetaResponse"></a>

### HermesMetaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether Hermes has bookkeeping information within the specified epoch range |
| numberOfDelegates | [uint64](#uint64) |  | number of Hermes delegates within the epoch range |
| numberOfRecipients | [uint64](#uint64) |  | number of voters who vote for Hermes delegates within the epoch range |
| totalRewardDistributed | [string](#string) |  | total reward amount distributed within the epoch range |






<a name="api-HermesRequest"></a>

### HermesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | Start epoch number |
| epochCount | [uint64](#uint64) |  | Number of epochs to query |
| rewardAddress | [string](#string) | repeated | Name of reward address |






<a name="api-HermesResponse"></a>

### HermesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hermesDistribution | [HermesDistribution](#api-HermesDistribution) | repeated |  |






<a name="api-RewardDistribution"></a>

### RewardDistribution



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| voterEthAddress | [string](#string) |  | voter’s ERC20 address |
| voterIotexAddress | [string](#string) |  | voter’s IoTeX address |
| amount | [string](#string) |  | amount of reward distribution |





 

 

 


<a name="api-HermesService"></a>

### HermesService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Hermes | [HermesRequest](#api-HermesRequest) | [HermesResponse](#api-HermesResponse) | Hermes gives delegates who register the service of automatic reward distribution an overview of the reward distributions to their voters within a range of epochs |
| HermesBucket | [HermesRequest](#api-HermesRequest) | [HermesBucketResponse](#api-HermesBucketResponse) | Hermes gives delegates who register the service of automatic reward distribution an overview of the bucket reward distributions to their voters within a range of epochs |
| HermesByVoter | [HermesByVoterRequest](#api-HermesByVoterRequest) | [HermesByVoterResponse](#api-HermesByVoterResponse) | HermesByVoter returns Hermes voters&#39; receiving history |
| HermesByDelegate | [HermesByDelegateRequest](#api-HermesByDelegateRequest) | [HermesByDelegateResponse](#api-HermesByDelegateResponse) | HermesByDelegate returns Hermes delegates&#39; distribution history |
| HermesMeta | [HermesMetaRequest](#api-HermesMetaRequest) | [HermesMetaResponse](#api-HermesMetaResponse) | HermesMeta provides Hermes platform metadata |
| HermesAverageStats | [HermesAverageStatsRequest](#api-HermesAverageStatsRequest) | [HermesAverageStatsResponse](#api-HermesAverageStatsResponse) | HermesAverageStats returns the Hermes average statistics |
| HermesDropRecords | [HermesDropRecordsRequest](#api-HermesDropRecordsRequest) | [HermesDropRecordsResponse](#api-HermesDropRecordsResponse) | HermesDropRecords inserts the Hermes drop records |

 



<a name="api_staking-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_staking.proto



<a name="api-BucketByIDRequest"></a>

### BucketByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bucket_id | [uint64](#uint64) | repeated |  |
| height | [uint64](#uint64) |  | 0 = latest indexed height |
| include_system | [bool](#bool) |  | whether to query system buckets, default false |






<a name="api-BucketByIDResponse"></a>

### BucketByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  |  |
| native_buckets | [StakingBucketInfo](#api-StakingBucketInfo) | repeated |  |
| system_buckets | [StakingBucketInfo](#api-StakingBucketInfo) | repeated |  |
| system_v2_buckets | [StakingBucketInfo](#api-StakingBucketInfo) | repeated |  |
| system_v3_buckets | [StakingBucketInfo](#api-StakingBucketInfo) | repeated |  |






<a name="api-BucketInfoEx"></a>

### BucketInfoEx
BucketInfoEx carries all fields for the bucket list/detail pages.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bucket_id | [int64](#int64) |  |  |
| action_hash | [string](#string) |  |  |
| timestamp | [string](#string) |  |  |
| create_time | [string](#string) |  |  |
| stake_start_time | [string](#string) |  |  |
| unstake_start_time | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| staked_amount | [string](#string) |  |  |
| act_type | [string](#string) |  |  |
| sender | [string](#string) |  |  |
| owner_address | [string](#string) |  |  |
| candidate | [string](#string) |  |  |
| auto_stake | [bool](#bool) |  |  |
| duration | [string](#string) |  |  |
| gas_price | [string](#string) |  |  |
| gas_limit | [string](#string) |  |  |
| recipient | [string](#string) |  |  |
| delegate_name | [string](#string) |  |  |






<a name="api-CandidateVoteByHeightRequest"></a>

### CandidateVoteByHeightRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) | repeated |  |
| height | [uint64](#uint64) |  |  |






<a name="api-CandidateVoteByHeightResponse"></a>

### CandidateVoteByHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  |  |
| stakeAmount | [string](#string) | repeated |  |
| voteWeight | [string](#string) | repeated |  |
| address | [string](#string) | repeated |  |






<a name="api-GetBucketByBucketIdRequest"></a>

### GetBucketByBucketIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bucket_id | [int64](#int64) |  |  |
| version | [string](#string) |  |  |






<a name="api-GetBucketByBucketIdResponse"></a>

### GetBucketByBucketIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  |  |
| bucket | [BucketInfoEx](#api-BucketInfoEx) |  |  |






<a name="api-GetBucketListRequest"></a>

### GetBucketListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int64](#int64) |  |  |
| offset | [int64](#int64) |  |  |
| sort | [string](#string) |  | e.g. &#34;timestamp:desc&#34; |
| interval | [string](#string) |  | &#34;1D&#34;, &#34;7D&#34;, &#34;30D&#34;, &#34;1Y&#34;, &#34;ALL&#34; |
| version | [string](#string) |  | &#34;native&#34;, &#34;nft_v1&#34;, &#34;nft_v2&#34;, &#34;nft_v3&#34; |






<a name="api-GetBucketListResponse"></a>

### GetBucketListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buckets | [BucketInfoEx](#api-BucketInfoEx) | repeated |  |
| count | [int64](#int64) |  |  |
| group_count | [int64](#int64) |  |  |






<a name="api-GetBucketsByBucketIdRequest"></a>

### GetBucketsByBucketIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bucket_id | [int64](#int64) |  |  |
| limit | [int64](#int64) |  |  |
| offset | [int64](#int64) |  |  |
| version | [string](#string) |  |  |






<a name="api-GetBucketsByBucketIdResponse"></a>

### GetBucketsByBucketIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buckets | [BucketInfoEx](#api-BucketInfoEx) | repeated |  |
| count | [int64](#int64) |  |  |






<a name="api-GetNativeBucketsRequest"></a>

### GetNativeBucketsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int64](#int64) |  |  |
| offset | [int64](#int64) |  |  |






<a name="api-GetNativeBucketsResponse"></a>

### GetNativeBucketsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buckets | [BucketInfoEx](#api-BucketInfoEx) | repeated |  |
| count | [int64](#int64) |  |  |






<a name="api-StakingBucketInfo"></a>

### StakingBucketInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bucket_id | [uint64](#uint64) |  |  |
| owner_address | [string](#string) |  |  |
| candidate | [string](#string) |  |  |
| staked_amount | [string](#string) |  |  |
| voting_power | [string](#string) |  |  |
| duration | [uint32](#uint32) |  | seconds |
| auto_stake | [bool](#bool) |  |  |
| create_time | [uint32](#uint32) |  | unix timestamp |
| stake_start_time | [uint32](#uint32) |  | unix timestamp |
| unstake_start_time | [uint32](#uint32) |  | unix timestamp, 0 if not unstaking |
| block_height | [uint64](#uint64) |  |  |






<a name="api-VoteByHeightRequest"></a>

### VoteByHeightRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) | repeated | voter address list |
| height | [uint64](#uint64) |  | block height |






<a name="api-VoteByHeightResponse"></a>

### VoteByHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  | block height |
| stakeAmount | [string](#string) | repeated | stake amount list |
| voteWeight | [string](#string) | repeated | vote weight list |





 

 

 


<a name="api-StakingService"></a>

### StakingService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| VoteByHeight | [VoteByHeightRequest](#api-VoteByHeightRequest) | [VoteByHeightResponse](#api-VoteByHeightResponse) | Get the stake amount and voting weight of the voter&#39;s specified height |
| CandidateVoteByHeight | [CandidateVoteByHeightRequest](#api-CandidateVoteByHeightRequest) | [CandidateVoteByHeightResponse](#api-CandidateVoteByHeightResponse) |  |
| BucketByID | [BucketByIDRequest](#api-BucketByIDRequest) | [BucketByIDResponse](#api-BucketByIDResponse) |  |
| GetBucketList | [GetBucketListRequest](#api-GetBucketListRequest) | [GetBucketListResponse](#api-GetBucketListResponse) |  |
| GetBucketsByBucketId | [GetBucketsByBucketIdRequest](#api-GetBucketsByBucketIdRequest) | [GetBucketsByBucketIdResponse](#api-GetBucketsByBucketIdResponse) |  |
| GetBucketByBucketId | [GetBucketByBucketIdRequest](#api-GetBucketByBucketIdRequest) | [GetBucketByBucketIdResponse](#api-GetBucketByBucketIdResponse) |  |
| GetNativeBuckets | [GetNativeBucketsRequest](#api-GetNativeBucketsRequest) | [GetNativeBucketsResponse](#api-GetNativeBucketsResponse) |  |

 



<a name="api_stream-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_stream.proto



<a name="api-SupplyRequest"></a>

### SupplyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startHeight | [uint64](#uint64) |  | start block height |
| endHeight | [uint64](#uint64) |  | end block height |






<a name="api-SupplyResponse"></a>

### SupplyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  | block height |
| totalSupply | [string](#string) |  | total supply |
| circulatingSupply | [string](#string) |  | circulating supply |





 

 

 


<a name="api-StreamService"></a>

### StreamService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Supply | [SupplyRequest](#api-SupplyRequest) | [SupplyResponse](#api-SupplyResponse) stream |  |

 



<a name="api_voting-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_voting.proto



<a name="api-CandidateInfoRequest"></a>

### CandidateInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | starting epoch number |
| epochCount | [uint64](#uint64) |  | epoch count |






<a name="api-CandidateInfoResponse"></a>

### CandidateInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| candidateInfo | [CandidateInfoResponse.CandidateInfo](#api-CandidateInfoResponse-CandidateInfo) | repeated |  |






<a name="api-CandidateInfoResponse-CandidateInfo"></a>

### CandidateInfoResponse.CandidateInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epochNumber | [uint64](#uint64) |  | epoch number |
| candidates | [CandidateInfoResponse.Candidates](#api-CandidateInfoResponse-Candidates) | repeated |  |






<a name="api-CandidateInfoResponse-Candidates"></a>

### CandidateInfoResponse.Candidates



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | candidate name |
| address | [string](#string) |  | canddiate address |
| totalWeightedVotes | [string](#string) |  | total weighted votes |
| selfStakingTokens | [string](#string) |  | candidate self-staking tokens |
| operatorAddress | [string](#string) |  | candidate operator address |
| rewardAddress | [string](#string) |  | candidate reward address |






<a name="api-CurrentDelegateInfo"></a>

### CurrentDelegateInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| name | [string](#string) |  |  |
| vote_weight | [string](#string) |  |  |
| productivity | [double](#double) |  |  |
| candidate | [string](#string) |  |  |
| operator_address | [string](#string) |  |  |
| active | [bool](#bool) |  |  |
| block_height | [uint64](#uint64) |  |  |






<a name="api-GetCurrentDelegatesRequest"></a>

### GetCurrentDelegatesRequest







<a name="api-GetCurrentDelegatesResponse"></a>

### GetCurrentDelegatesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  |  |
| delegates | [CurrentDelegateInfo](#api-CurrentDelegateInfo) | repeated |  |






<a name="api-RewardSourcesRequest"></a>

### RewardSourcesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | starting epoch number |
| epochCount | [uint64](#uint64) |  | epoch count |
| voterIotxAddress | [string](#string) |  | voter IoTeX address |






<a name="api-RewardSourcesResponse"></a>

### RewardSourcesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether the voter has reward information within the specified epoch range |
| delegateDistributions | [RewardSourcesResponse.DelegateDistributions](#api-RewardSourcesResponse-DelegateDistributions) | repeated |  |






<a name="api-RewardSourcesResponse-DelegateDistributions"></a>

### RewardSourcesResponse.DelegateDistributions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegateName | [string](#string) |  | delegate name |
| amount | [string](#string) |  | amount of reward distribution |






<a name="api-VotingMetaRequest"></a>

### VotingMetaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | starting epoch number |
| epochCount | [uint64](#uint64) |  | epoch count |






<a name="api-VotingMetaResponse"></a>

### VotingMetaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether the starting epoch number is less than the most recent epoch number |
| candidateMeta | [VotingMetaResponse.CandidateMeta](#api-VotingMetaResponse-CandidateMeta) | repeated |  |






<a name="api-VotingMetaResponse-CandidateMeta"></a>

### VotingMetaResponse.CandidateMeta



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epochNumber | [uint64](#uint64) |  | epoch number |
| consensusDelegates | [uint64](#uint64) |  | number of consensus delegates in the epoch |
| totalCandidates | [uint64](#uint64) |  | number of total delegates in the epoch |
| totalWeightedVotes | [string](#string) |  | candidate total weighted votes in the epoch |
| votedTokens | [string](#string) |  | total voted tokens in the epoch |





 

 

 


<a name="api-VotingService"></a>

### VotingService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CandidateInfo | [CandidateInfoRequest](#api-CandidateInfoRequest) | [CandidateInfoResponse](#api-CandidateInfoResponse) |  |
| RewardSources | [RewardSourcesRequest](#api-RewardSourcesRequest) | [RewardSourcesResponse](#api-RewardSourcesResponse) | RewardSources provides reward sources for voters |
| VotingMeta | [VotingMetaRequest](#api-VotingMetaRequest) | [VotingMetaResponse](#api-VotingMetaResponse) | VotingMeta provides metadata of voting results |
| GetCurrentDelegates | [GetCurrentDelegatesRequest](#api-GetCurrentDelegatesRequest) | [GetCurrentDelegatesResponse](#api-GetCurrentDelegatesResponse) | GetCurrentDelegates returns the current delegate list ordered by vote weight |

 



<a name="api_xrc20-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_xrc20.proto



<a name="api-GetXRC20HoldersByContractRequest"></a>

### GetXRC20HoldersByContractRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  |  |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-GetXRC20HoldersByContractResponse"></a>

### GetXRC20HoldersByContractResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [int64](#int64) |  |  |
| holders | [XRC20HolderInfo](#api-XRC20HolderInfo) | repeated |  |






<a name="api-GetXRC20StatsRequest"></a>

### GetXRC20StatsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-GetXRC20StatsResponse"></a>

### GetXRC20StatsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint64](#uint64) |  | total number of xrc20 tokens |
| items | [XRC20StatsItem](#api-XRC20StatsItem) | repeated |  |






<a name="api-GetXRC20TokenBalanceRequest"></a>

### GetXRC20TokenBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  |  |
| address | [string](#string) |  |  |






<a name="api-GetXRC20TokenBalanceResponse"></a>

### GetXRC20TokenBalanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balance | [string](#string) |  |  |






<a name="api-GetXRC20TransfersByContractRequest"></a>

### GetXRC20TransfersByContractRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  |  |
| address | [string](#string) |  | optional sender/recipient filter |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-GetXRC20TransfersByContractResponse"></a>

### GetXRC20TransfersByContractResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  |  |
| count | [int64](#int64) |  |  |
| transfers | [XRC20TransferInfo](#api-XRC20TransferInfo) | repeated |  |






<a name="api-XRC20AddressesRequest"></a>

### XRC20AddressesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-XRC20AddressesResponse"></a>

### XRC20AddressesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether Xrc20 contract addresses exist |
| count | [uint64](#uint64) |  | total number of Xrc20 contract addresses |
| addresses | [string](#string) | repeated |  |






<a name="api-XRC20ByAddressRequest"></a>

### XRC20ByAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | sender address or recipient address |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-XRC20ByAddressResponse"></a>

### XRC20ByAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether Xrc20 actions exist for the given sender address or recipient address |
| count | [uint64](#uint64) |  | total number of Xrc20 actions |
| xrc20 | [Xrc20Action](#api-Xrc20Action) | repeated |  |






<a name="api-XRC20ByContractAddressRequest"></a>

### XRC20ByContractAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | contract address |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-XRC20ByContractAddressResponse"></a>

### XRC20ByContractAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether Xrc20 actions exist for the given contract address |
| count | [uint64](#uint64) |  | total number of Xrc20 actions |
| xrc20 | [Xrc20Action](#api-Xrc20Action) | repeated |  |






<a name="api-XRC20ByPageRequest"></a>

### XRC20ByPageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-XRC20ByPageResponse"></a>

### XRC20ByPageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether Xrc20 actions exist for the given contract address |
| count | [uint64](#uint64) |  | total number of Xrc20 actions |
| xrc20 | [Xrc20Action](#api-Xrc20Action) | repeated |  |






<a name="api-XRC20HolderInfo"></a>

### XRC20HolderInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| balance | [string](#string) |  |  |






<a name="api-XRC20StatsItem"></a>

### XRC20StatsItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| holders | [uint64](#uint64) |  |  |
| transfer | [uint64](#uint64) |  |  |
| daily_transfer | [uint64](#uint64) |  |  |






<a name="api-XRC20TokenHolderAddressesRequest"></a>

### XRC20TokenHolderAddressesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tokenAddress | [string](#string) |  | token contract address |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-XRC20TokenHolderAddressesResponse"></a>

### XRC20TokenHolderAddressesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint64](#uint64) |  | total number of token holder addresses |
| addresses | [string](#string) | repeated |  |






<a name="api-XRC20TransferInfo"></a>

### XRC20TransferInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| block_height | [uint64](#uint64) |  |  |
| action_hash | [string](#string) |  |  |
| contract_address | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| sender | [string](#string) |  |  |
| recipient | [string](#string) |  |  |
| timestamp | [string](#string) |  |  |






<a name="api-Xrc20Action"></a>

### Xrc20Action



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actHash | [string](#string) |  | action hash |
| sender | [string](#string) |  | sender address |
| recipient | [string](#string) |  | recipient address |
| amount | [string](#string) |  | amount transferred |
| timestamp | [uint64](#uint64) |  | unix timestamp |
| contract | [string](#string) |  | contract address |





 

 

 


<a name="api-XRC20Service"></a>

### XRC20Service


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| XRC20ByAddress | [XRC20ByAddressRequest](#api-XRC20ByAddressRequest) | [XRC20ByAddressResponse](#api-XRC20ByAddressResponse) | XRC20ByAddress returns Xrc20 actions given the sender address or recipient address |
| XRC20ByContractAddress | [XRC20ByContractAddressRequest](#api-XRC20ByContractAddressRequest) | [XRC20ByContractAddressResponse](#api-XRC20ByContractAddressResponse) | XRC20ByContractAddress returns Xrc20 actions given the Xrc20 contract address |
| XRC20ByPage | [XRC20ByPageRequest](#api-XRC20ByPageRequest) | [XRC20ByPageResponse](#api-XRC20ByPageResponse) | XRC20ByPage returns Xrc20 actions by pagination |
| XRC20Addresses | [XRC20AddressesRequest](#api-XRC20AddressesRequest) | [XRC20AddressesResponse](#api-XRC20AddressesResponse) | XRC20Addresses returns Xrc20 contract addresses |
| XRC20TokenHolderAddresses | [XRC20TokenHolderAddressesRequest](#api-XRC20TokenHolderAddressesRequest) | [XRC20TokenHolderAddressesResponse](#api-XRC20TokenHolderAddressesResponse) | XRC20TokenHolderAddresses returns Xrc20 token holder addresses given a Xrc20 contract address |
| GetXRC20TransfersByContract | [GetXRC20TransfersByContractRequest](#api-GetXRC20TransfersByContractRequest) | [GetXRC20TransfersByContractResponse](#api-GetXRC20TransfersByContractResponse) | GetXRC20TransfersByContract returns ERC20 transfers filtered by contract and optional sender/recipient |
| GetXRC20HoldersByContract | [GetXRC20HoldersByContractRequest](#api-GetXRC20HoldersByContractRequest) | [GetXRC20HoldersByContractResponse](#api-GetXRC20HoldersByContractResponse) | GetXRC20HoldersByContract returns all holders of a given ERC20 token with balances |
| GetXRC20TokenBalance | [GetXRC20TokenBalanceRequest](#api-GetXRC20TokenBalanceRequest) | [GetXRC20TokenBalanceResponse](#api-GetXRC20TokenBalanceResponse) | GetXRC20TokenBalance returns the ERC20 token balance for a specific address |
| GetXRC20Stats | [GetXRC20StatsRequest](#api-GetXRC20StatsRequest) | [GetXRC20StatsResponse](#api-GetXRC20StatsResponse) | GetXRC20Stats returns per-token holder/transfer counts, ordered by holders DESC |

 



<a name="api_xrc721-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_xrc721.proto



<a name="api-GetNFTHoldersByContractRequest"></a>

### GetNFTHoldersByContractRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  |  |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-GetNFTHoldersByContractResponse"></a>

### GetNFTHoldersByContractResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [int64](#int64) |  |  |
| holders | [NFTHolderInfo](#api-NFTHolderInfo) | repeated |  |






<a name="api-GetNFTTransferListRequest"></a>

### GetNFTTransferListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |
| contract_address | [string](#string) |  | optional; leave empty to query all contracts |
| address | [string](#string) |  | optional; sender/recipient or token_id filter |






<a name="api-GetNFTTransferListResponse"></a>

### GetNFTTransferListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  |  |
| count | [int64](#int64) |  |  |
| transfers | [NFTTransferInfo](#api-NFTTransferInfo) | repeated |  |






<a name="api-NFTHolderInfo"></a>

### NFTHolderInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  |  |
| balance | [string](#string) |  |  |






<a name="api-NFTTransferInfo"></a>

### NFTTransferInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  |  |
| type | [string](#string) |  | &#34;xrc721&#34; or &#34;xrc1155&#34; |
| block_height | [uint64](#uint64) |  |  |
| action_hash | [string](#string) |  |  |
| contract_address | [string](#string) |  |  |
| token_id | [string](#string) |  |  |
| value | [string](#string) |  |  |
| sender | [string](#string) |  |  |
| recipient | [string](#string) |  |  |
| timestamp | [string](#string) |  |  |






<a name="api-XRC721AddressesRequest"></a>

### XRC721AddressesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-XRC721AddressesResponse"></a>

### XRC721AddressesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether Xrc721 contract addresses exist |
| count | [uint64](#uint64) |  | total number of Xrc721 contract addresses |
| addresses | [string](#string) | repeated |  |






<a name="api-XRC721ByAddressRequest"></a>

### XRC721ByAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | sender address or recipient address |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-XRC721ByAddressResponse"></a>

### XRC721ByAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether Xrc721 actions exist for the given sender address or recipient address |
| count | [uint64](#uint64) |  | total number of Xrc721 actions |
| xrc721 | [Xrc721Action](#api-Xrc721Action) | repeated |  |






<a name="api-XRC721ByContractAddressRequest"></a>

### XRC721ByContractAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | contract address |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-XRC721ByContractAddressResponse"></a>

### XRC721ByContractAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether Xrc721 actions exist for the given contract address |
| count | [uint64](#uint64) |  | total number of Xrc721 actions |
| xrc721 | [Xrc721Action](#api-Xrc721Action) | repeated |  |






<a name="api-XRC721ByPageRequest"></a>

### XRC721ByPageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-XRC721ByPageResponse"></a>

### XRC721ByPageResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether Xrc721 actions exist for the given contract address |
| count | [uint64](#uint64) |  | total number of Xrc721 actions |
| xrc721 | [Xrc721Action](#api-Xrc721Action) | repeated |  |






<a name="api-XRC721TokenHolderAddressesRequest"></a>

### XRC721TokenHolderAddressesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tokenAddress | [string](#string) |  | token contract address |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-XRC721TokenHolderAddressesResponse"></a>

### XRC721TokenHolderAddressesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint64](#uint64) |  | total number of token holder addresses |
| addresses | [string](#string) | repeated |  |






<a name="api-Xrc721Action"></a>

### Xrc721Action



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actHash | [string](#string) |  | action hash |
| sender | [string](#string) |  | sender address |
| recipient | [string](#string) |  | recipient address |
| amount | [string](#string) |  | amount transferred |
| timestamp | [uint64](#uint64) |  | unix timestamp |
| contract | [string](#string) |  | contract address |





 

 

 


<a name="api-XRC721Service"></a>

### XRC721Service


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| XRC721ByAddress | [XRC721ByAddressRequest](#api-XRC721ByAddressRequest) | [XRC721ByAddressResponse](#api-XRC721ByAddressResponse) | XRC721ByAddress returns Xrc721 actions given the sender address or recipient address |
| XRC721ByContractAddress | [XRC721ByContractAddressRequest](#api-XRC721ByContractAddressRequest) | [XRC721ByContractAddressResponse](#api-XRC721ByContractAddressResponse) | XRC721ByContractAddress returns Xrc721 actions given the Xrc721 contract address |
| XRC721ByPage | [XRC721ByPageRequest](#api-XRC721ByPageRequest) | [XRC721ByPageResponse](#api-XRC721ByPageResponse) | XRC721ByPage returns Xrc721 actions by pagination |
| XRC721Addresses | [XRC721AddressesRequest](#api-XRC721AddressesRequest) | [XRC721AddressesResponse](#api-XRC721AddressesResponse) | XRC721Addresses returns Xrc721 contract addresses |
| XRC721TokenHolderAddresses | [XRC721TokenHolderAddressesRequest](#api-XRC721TokenHolderAddressesRequest) | [XRC721TokenHolderAddressesResponse](#api-XRC721TokenHolderAddressesResponse) | XRC721TokenHolderAddresses returns Xrc721 token holder addresses given a Xrc721 contract address |
| GetNFTTransferList | [GetNFTTransferListRequest](#api-GetNFTTransferListRequest) | [GetNFTTransferListResponse](#api-GetNFTTransferListResponse) | GetNFTTransferList returns NFT transfers (xrc721 &#43; xrc1155) with optional contract/address filters |
| GetNFTHoldersByContract | [GetNFTHoldersByContractRequest](#api-GetNFTHoldersByContractRequest) | [GetNFTHoldersByContractResponse](#api-GetNFTHoldersByContractResponse) | GetNFTHoldersByContract returns NFT holders for a contract |

 



<a name="include_pagination-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## include/pagination.proto



<a name="pagination-Pagination"></a>

### Pagination



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| skip | [uint64](#uint64) |  | starting index of results |
| first | [uint64](#uint64) |  | number of records per page |
| order | [string](#string) |  |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

