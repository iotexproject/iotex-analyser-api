# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api_account.proto](#api_account-proto)
    - [Erc20TokenBalanceByHeightRequest](#api-Erc20TokenBalanceByHeightRequest)
    - [Erc20TokenBalanceByHeightResponse](#api-Erc20TokenBalanceByHeightResponse)
    - [HermesDistribution](#api-HermesDistribution)
    - [HermesRequest](#api-HermesRequest)
    - [HermesResponse](#api-HermesResponse)
    - [IotexBalanceByHeightRequest](#api-IotexBalanceByHeightRequest)
    - [IotexBalanceByHeightResponse](#api-IotexBalanceByHeightResponse)
    - [RewardDistribution](#api-RewardDistribution)
  
    - [AccountService](#api-AccountService)
  
- [api_action.proto](#api_action-proto)
    - [ActionInfo](#api-ActionInfo)
    - [ActionRequest](#api-ActionRequest)
    - [ActionResponse](#api-ActionResponse)
    - [EvmTransferInfo](#api-EvmTransferInfo)
    - [XrcInfo](#api-XrcInfo)
  
    - [ActionService](#api-ActionService)
  
- [api_actions.proto](#api_actions-proto)
    - [ActionsByAddressResponse](#api-ActionsByAddressResponse)
    - [ActionsByAddressResult](#api-ActionsByAddressResult)
    - [ActionsRequest](#api-ActionsRequest)
    - [AllActionsByAddressResponse](#api-AllActionsByAddressResponse)
    - [AllActionsByAddressResult](#api-AllActionsByAddressResult)
    - [EvmTransferDetailListByAddressResponse](#api-EvmTransferDetailListByAddressResponse)
    - [EvmTransferDetailResult](#api-EvmTransferDetailResult)
    - [Xrc20ByAddressResponse](#api-Xrc20ByAddressResponse)
    - [Xrc20ByAddressResult](#api-Xrc20ByAddressResult)
  
    - [AllActionsByAddressResult.RecordType](#api-AllActionsByAddressResult-RecordType)
  
    - [ActionsService](#api-ActionsService)
  
- [api_chain.proto](#api_chain-proto)
    - [ChainRequest](#api-ChainRequest)
    - [ChainResponse](#api-ChainResponse)
  
    - [ChainService](#api-ChainService)
  
- [api_delegate.proto](#api_delegate-proto)
    - [BookKeepingRequest](#api-BookKeepingRequest)
    - [BookKeepingResponse](#api-BookKeepingResponse)
    - [BucketInfo](#api-BucketInfo)
    - [BucketInfoList](#api-BucketInfoList)
    - [BucketInfoRequest](#api-BucketInfoRequest)
    - [BucketInfoResponse](#api-BucketInfoResponse)
    - [DelegateRewardDistribution](#api-DelegateRewardDistribution)
    - [HermesByDelegateDistributionRatio](#api-HermesByDelegateDistributionRatio)
    - [HermesByDelegateRequest](#api-HermesByDelegateRequest)
    - [HermesByDelegateResponse](#api-HermesByDelegateResponse)
    - [HermesByDelegateVoterInfo](#api-HermesByDelegateVoterInfo)
    - [Productivity](#api-Productivity)
    - [ProductivityRequest](#api-ProductivityRequest)
    - [ProductivityResponse](#api-ProductivityResponse)
    - [Reward](#api-Reward)
    - [RewardRequest](#api-RewardRequest)
    - [RewardResponse](#api-RewardResponse)
  
    - [DelegateService](#api-DelegateService)
  
- [api_staking.proto](#api_staking-proto)
    - [CandidateVoteByHeightRequest](#api-CandidateVoteByHeightRequest)
    - [CandidateVoteByHeightResponse](#api-CandidateVoteByHeightResponse)
    - [VoteByHeightRequest](#api-VoteByHeightRequest)
    - [VoteByHeightResponse](#api-VoteByHeightResponse)
  
    - [StakingService](#api-StakingService)
  
- [include/pagination.proto](#include_pagination-proto)
    - [Pagination](#pagination-Pagination)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_account-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_account.proto



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






<a name="api-HermesRequest"></a>

### HermesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| startEpoch | [uint64](#uint64) |  | Start epoch number |
| epochCount | [uint64](#uint64) |  | Number of epochs to query |
| rewardAddress | [string](#string) |  | Name of reward address |






<a name="api-HermesResponse"></a>

### HermesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hermesDistribution | [HermesDistribution](#api-HermesDistribution) | repeated |  |






<a name="api-IotexBalanceByHeightRequest"></a>

### IotexBalanceByHeightRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) | repeated |  |
| height | [uint64](#uint64) |  |  |






<a name="api-IotexBalanceByHeightResponse"></a>

### IotexBalanceByHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  |  |
| balance | [string](#string) | repeated |  |






<a name="api-RewardDistribution"></a>

### RewardDistribution



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| voterEthAddress | [string](#string) |  | voter’s ERC20 address |
| voterIotexAddress | [string](#string) |  | voter’s IoTeX address |
| amount | [string](#string) |  | amount of reward distribution |





 

 

 


<a name="api-AccountService"></a>

### AccountService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| IotexBalanceByHeight | [IotexBalanceByHeightRequest](#api-IotexBalanceByHeightRequest) | [IotexBalanceByHeightResponse](#api-IotexBalanceByHeightResponse) |  |
| Erc20TokenBalanceByHeight | [Erc20TokenBalanceByHeightRequest](#api-Erc20TokenBalanceByHeightRequest) | [Erc20TokenBalanceByHeightResponse](#api-Erc20TokenBalanceByHeightResponse) |  |
| Hermes | [HermesRequest](#api-HermesRequest) | [HermesResponse](#api-HermesResponse) | Hermes gives delegates who register the service of automatic reward distribution an overview of the reward distributions to their voters within a range of epochs |

 



<a name="api_action-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_action.proto



<a name="api-ActionInfo"></a>

### ActionInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actHash | [string](#string) |  |  |
| blkHash | [string](#string) |  |  |
| actType | [string](#string) |  |  |
| sender | [string](#string) |  |  |
| recipient | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| timestamp | [uint64](#uint64) |  |  |
| gasFee | [string](#string) |  |  |
| blkHeight | [uint64](#uint64) |  |  |






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
| ActionByAddress | [ActionRequest](#api-ActionRequest) | [ActionResponse](#api-ActionResponse) |  |
| EvmTransfersByAddress | [ActionRequest](#api-ActionRequest) | [ActionResponse](#api-ActionResponse) |  |
| GetXrc20ByAddress | [ActionRequest](#api-ActionRequest) | [ActionResponse](#api-ActionResponse) |  |

 



<a name="api_actions-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_actions.proto



<a name="api-ActionsByAddressResponse"></a>

### ActionsByAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint64](#uint64) |  |  |
| results | [ActionsByAddressResult](#api-ActionsByAddressResult) | repeated |  |






<a name="api-ActionsByAddressResult"></a>

### ActionsByAddressResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actHash | [string](#string) |  |  |
| blkHash | [string](#string) |  |  |
| timeStamp | [uint64](#uint64) |  |  |
| actType | [string](#string) |  |  |
| sender | [string](#string) |  |  |
| recipient | [string](#string) |  |  |
| amount | [string](#string) |  |  |
| gasFee | [string](#string) |  |  |






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






<a name="api-Xrc20ByAddressResponse"></a>

### Xrc20ByAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint64](#uint64) |  |  |
| results | [Xrc20ByAddressResult](#api-Xrc20ByAddressResult) | repeated |  |






<a name="api-Xrc20ByAddressResult"></a>

### Xrc20ByAddressResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actHash | [string](#string) |  |  |
| blkHeight | [uint64](#uint64) |  |  |
| from | [string](#string) |  |  |
| to | [string](#string) |  |  |
| contractAddress | [string](#string) |  |  |
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
| GetActionsByAddress | [ActionsRequest](#api-ActionsRequest) | [ActionsByAddressResponse](#api-ActionsByAddressResponse) |  |
| GetXrc20ByAddress | [ActionsRequest](#api-ActionsRequest) | [Xrc20ByAddressResponse](#api-Xrc20ByAddressResponse) |  |
| GetXrc721ByAddress | [ActionsRequest](#api-ActionsRequest) | [Xrc20ByAddressResponse](#api-Xrc20ByAddressResponse) |  |
| GetEvmTransferDetailListByAddress | [ActionsRequest](#api-ActionsRequest) | [EvmTransferDetailListByAddressResponse](#api-EvmTransferDetailListByAddressResponse) |  |
| GetAllActionsByAddress | [ActionsRequest](#api-ActionsRequest) | [AllActionsByAddressResponse](#api-AllActionsByAddressResponse) |  |

 



<a name="api_chain-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_chain.proto



<a name="api-ChainRequest"></a>

### ChainRequest







<a name="api-ChainResponse"></a>

### ChainResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mostRecentEpoch | [uint64](#uint64) |  | most recent epoch |
| mostRecentBlockHeight | [uint64](#uint64) |  | most recent block height |





 

 

 


<a name="api-ChainService"></a>

### ChainService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Chain | [ChainRequest](#api-ChainRequest) | [ChainResponse](#api-ChainResponse) |  |

 



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
| percentage | [uint64](#uint64) |  | percentage of the reward to be paid to the delegate |
| includeBlockReward | [bool](#bool) |  | whether to include block reward |
| includeFoundationBonus | [bool](#bool) |  | whether to include foundation bonus |






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
| actionHash | [string](#string) |  | action hash |
| timestamp | [string](#string) |  | timestamp |






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





 

 

 


<a name="api-DelegateService"></a>

### DelegateService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| BucketInfo | [BucketInfoRequest](#api-BucketInfoRequest) | [BucketInfoResponse](#api-BucketInfoResponse) | BucketInfo provides voting bucket detail information for candidates within a range of epochs |
| BookKeeping | [BookKeepingRequest](#api-BookKeepingRequest) | [BookKeepingResponse](#api-BookKeepingResponse) | BookKeeping gives delegates an overview of the reward distributions to their voters within a range of epochs |
| Productivity | [ProductivityRequest](#api-ProductivityRequest) | [ProductivityResponse](#api-ProductivityResponse) | Productivity gives block productivity of producers within a range of epochs |
| Reward | [RewardRequest](#api-RewardRequest) | [RewardResponse](#api-RewardResponse) | Rewards provides reward detail information for candidates within a range of epochs |
| HermesByDelegate | [HermesByDelegateRequest](#api-HermesByDelegateRequest) | [HermesByDelegateResponse](#api-HermesByDelegateResponse) | HermesByDelegate returns Hermes delegates&#39; distribution history |

 



<a name="api_staking-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api_staking.proto



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






<a name="api-VoteByHeightRequest"></a>

### VoteByHeightRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) | repeated |  |
| height | [uint64](#uint64) |  |  |






<a name="api-VoteByHeightResponse"></a>

### VoteByHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  |  |
| stakeAmount | [string](#string) | repeated |  |
| voteWeight | [string](#string) | repeated |  |





 

 

 


<a name="api-StakingService"></a>

### StakingService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| VoteByHeight | [VoteByHeightRequest](#api-VoteByHeightRequest) | [VoteByHeightResponse](#api-VoteByHeightResponse) |  |
| CandidateVoteByHeight | [CandidateVoteByHeightRequest](#api-CandidateVoteByHeightRequest) | [CandidateVoteByHeightResponse](#api-CandidateVoteByHeightResponse) |  |

 



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

