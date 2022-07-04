---
title: IoTeX Analytics API Reference

language_tabs: # must be one of https://git.io/vQNgJ
  - shell
  - graphql

toc_footers:
  - <a href='https://iotex.io'>IoTeX</a>

includes:
#   - errors

search: true

code_clipboard: true

meta:
  - name: description
    content: Documentation for IoTex Analytics API
---

# Introduction

Analytics is an application built upon IoTeX core API which extracts data from IoTeX blockchain and reindexes them for applications to use via a GraphQL web interface. You can use the [playground here](https://analyser-api.iotex.io/graphql).

## Common Attributes

<a name="pagination-Pagination"></a>

### Pagination



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| skip | [uint64](#uint64) |  | starting index of results |
| first | [uint64](#uint64) |  | number of records per page |


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

# Chain Service API

## Chain

gives the latest epoch number / block height.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.Chain \
  --header 'Content-Type: application/json' \
  --data '{}'
```

```graphql
query {
	Chain {
		mostRecentEpoch
		mostRecentBlockHeight
	}
}

```

> Example response:

```json
{
	"mostRecentEpoch": "25935",
	"mostRecentBlockHeight": "16856649"
}
```

### HTTP Request

`POST /api.ChainService.Chain`

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

<a name="api-VotingResultMeta"></a>

### VotingResultMeta



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| totalCandidates | [uint64](#uint64) |  | total candidates |
| totalWeightedVotes | [string](#string) |  | total weighted votes |
| votedTokens | [string](#string) |  | voted tokens |

## MostRecentTPS

gives the latest transactions per second

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.MostRecentTPS \
  --header 'Content-Type: application/json' \
  --data '{
	"blockWindow": 5
}'
```

```graphql
query {
	MostRecentTPS(blockWindow: 5) {
		mostRecentTPS
	}
}


```

> Example response:

```json
{
	"data": {
		"MostRecentTPS": {
			"mostRecentTPS": 0.8421052631578947
		}
	}
}
```

### HTTP Request

`POST /api.ChainService.MostRecentTPS`

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

## NumberOfActions

gives the number of actions

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.NumberOfActions \
  --header 'Content-Type: application/json' \
  --data '{
	"startEpoch": 20000,
	"epochCount": 5
}'
```

```graphql
query {
	NumberOfActions(startEpoch: 20000, epochCount: 5) {
		exist
		count
	}
}


```

> Example response:

```json
{
	"data": {
		"NumberOfActions": {
			"count": 12744,
			"exist": true
		}
	}
}
```

### HTTP Request

`POST /api.ChainService.NumberOfActions`

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

## TotalTransferredTokens

TotalTransferredTokens gives the amount of tokens transferred within a time frame

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.TotalTransferredTokens \
  --header 'Content-Type: application/json' \
  --data '{
	"epochStart": 20000,
  "epochCount": 2
}'
```

```graphql
query {
  TotalTransferredTokens(startEpoch: 20000, epochCount: 2) {
    totalTransferredTokens
  }
}

```

> Example response:

```json
{
  "data": {
    "TotalTransferredTokens": {
      "totalTransferredTokens": "2689365862730085490074594"
    }
  }
}
```

### HTTP Request

`POST /api.ChainService.TotalTransferredTokens`

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


# Delegate Service API

## BucketInfo

BucketInfo provides voting bucket detail information for candidates within a range of epochs

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.DelegateService.BucketInfo \
  --header 'Content-Type: application/json' \
  --data '{
	"startEpoch": 24738,
	"epochCount": 2,
	"delegateName": "metanyx",
	"pagination": {
		"skip": 0,
		"first": 5
	}
}'
```

```graphql
query {
	BucketInfo(
		startEpoch: 24738
		epochCount: 2
		delegateName: "metanyx"
		pagination: { skip: 0, first: 5 }
	) {
		bucketInfoList {
			epochNumber
			count
			bucketInfo {
				bucketID
				voterIotexAddress
				votes
				weightedVotes
				remainingDuration
				isNative
				startTime
				
			}
		}
	}
}
```

> Example response:

```json
{
	"exist": true,
	"count": "2",
	"bucketInfoList": [
		{
			"epochNumber": "24738",
			"count": "7136",
			"bucketInfo": [
				{
					"voterEthAddress": "0x2e0491b4925ebc82af97def12b72ead940613293",
					"voterIotexAddress": "io19czfrdyjt67g9tuhmmcjkuh2m9qxzv5nqyve9p",
					"isNative": true,
					"votes": "1283019032866474771223667",
					"weightedVotes": "1848668139159798560008586",
					"remainingDuration": "8400h0m0s",
					"startTime": "2020-06-01 18:14:45 +0000 UTC",
					"decay": false,
					"bucketID": "39"
				},
				...
			]
		},
        ...
	]
}
```

### HTTP Request

`POST /api.DelegateService.BucketInfo`

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


## BookKeeping

BookKeeping gives delegates an overview of the reward distributions to their voters within a range of epochs

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.DelegateService.BookKeeping \
  --header 'Content-Type: application/json' \
  --data '{
	"startEpoch": 23328,
	"epochCount": 10,
	"delegateName": "iotexlab",
	"pagination": {
		"skip": 0,
		"first": 2
	},
	"percentage": 90,
	"includeBlockReward":true,
	"includeFoundationBonus":false
}'
```

```graphql
query {
  BookKeeping(
    startEpoch: 23328
    epochCount: 10
    delegateName: "iotexlab"
    percentage: 90
    includeFoundationBonus: false
    includeBlockReward: false
    pagination: { skip: 0, first: 2 }
  ) {
    count
    rewardDistribution {
      voterEthAddress
      amount
    }
  }
}

```

> Example response:

```json
{
	"exist": true,
	"count": "5567",
	"rewardDistribution": [
		{
			"voterEthAddress": "0x0002d2d9945709b50cfbac675d7e6bdac34575f4",
			"voterIotexAddress": "io1qqpd9kv52uym2r8m43n46lntmtp52a05d7s8gj",
			"amount": "77071218081884295"
		},
		...
	]
}
```

### HTTP Request

`POST /api.DelegateService.BookKeeping`

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

<a name="api-DelegateRewardDistribution"></a>

### DelegateRewardDistribution

| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| voterEthAddress | [string](#string) |  | voter’s ERC20 address |
| voterIotexAddress | [string](#string) |  | voter’s IoTeX address |
| amount | [string](#string) |  | amount of reward distribution |

## Productivity

Productivity gives block productivity of producers within a range of epochs

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.DelegateService.Productivity \
  --header 'Content-Type: application/json' \
  --data '{
	"startEpoch": 25020,
	"epochCount": 10,
	"delegateName": "iotexlab"
}'
```

```graphql
query {
  Productivity(
    startEpoch: 25020
    epochCount: 10
    delegateName: "iotexlab"
  ) {
    productivity {
      exist
      production
      expectedProduction
    }
  }
}


```

> Example response:

```json
{
	"productivity": {
		"exist": true,
		"production": "180",
		"expectedProduction": "180"
	}
}
```

### HTTP Request

`POST /api.DelegateService.Productivity`

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

### Productivity



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether the delegate has productivity information within the specified epoch range |
| production | [uint64](#uint64) |  | number of block productions |
| expectedProduction | [uint64](#uint64) |  | number of expected block productions |


## Reward

Reward provides reward detail information for candidates within a range of epochs

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.DelegateService.Reward \
  --header 'Content-Type: application/json' \
  --data '{
	"startEpoch": 23000,
	"epochCount": 1,
	"delegateName": "iotexlab"
}'
```

```graphql
query {
  Reward(startEpoch: 23000, epochCount: 1, delegateName: "iotexlab") {
    reward {
      exist
      blockReward
      foundationBonus
      epochReward
    }
  }
}

```

> Example response:

```json
{
	"reward": {
		"blockReward": "240000000000000000000",
		"epochReward": "984040630606589747896",
		"foundationBonus": "80000000000000000000",
		"exist": true
	}
}
```

### HTTP Request

`POST /api.DelegateService.Reward`

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

<a name="api-Reward"></a>

### Reward



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| blockReward | [string](#string) |  | amount of block rewards |
| epochReward | [string](#string) |  | amount of epoch rewards |
| foundationBonus | [string](#string) |  | amount of foundation bonus |
| exist | [bool](#bool) |  | whether the delegate has reward information within the specified epoch range |


## Staking

Staking provides staking information for candidates within a range of epochs

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.DelegateService.Staking \
  --header 'Content-Type: application/json' \
  --data '{
	"startEpoch": 20000,
	"epochCount": 2,
	"delegateName": "metanyx"
}'
```

```graphql
query {
  Staking(startEpoch: 20000, epochCount: 2, delegateName: "metanyx") {
    exist
    stakingInfo {
      epochNumber
      totalStaking
      selfStaking
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "Staking": {
      "exist": true,
      "stakingInfo": [
        {
          "epochNumber": 20000,
          "selfStaking": "1266890287625445522068595",
          "totalStaking": "219516310335741609989431119"
        },
        {
          "epochNumber": 20001,
          "selfStaking": "1266890287625445522068595",
          "totalStaking": "219518846815421318945180493"
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.DelegateService.Staking`

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

## ProbationHistoricalRate

ProbationHistoricalRate provides the rate of probation for a given delegate

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.DelegateService.ProbationHistoricalRate \
  --header 'Content-Type: application/json' \
  --data '{
	"startEpoch": 27650,
	"epochCount": 5,
	"delegateName": "chainshield"
}'
```

```graphql
query {
  ProbationHistoricalRate(
    startEpoch: 27650
    epochCount: 5
    delegateName: "chainshield"
  ) {
    probationHistoricalRate
  }
}
```

> Example response:

```json
{
  "data": {
    "ProbationHistoricalRate": {
      "probationHistoricalRate": "0.80"
    }
  }
}
```

### HTTP Request

`POST /api.DelegateService.ProbationHistoricalRate`

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



# Account Service API

## IotexBalanceByHeight

IotexBalanceByHeight returns the balance of the given address at the given height.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.AccountService.IotexBalanceByHeight \
  --header 'Content-Type: application/json' \
  --data '{
	"address": ["io1qnpz47hx5q6r3w876axtrn6yz95d70cjl35r53"], 
  	"height":8927781 
}'
```

```graphql
query {
  IotexBalanceByHeight(
    address: ["io1qnpz47hx5q6r3w876axtrn6yz95d70cjl35r53"]
    height: 8927781
  ) {
    balance
    height
  }
}


```

> Example response:

```json
{
  "data": {
    "IotexBalanceByHeight": {
      "balance": [
        "957.111886573698936216"
      ],
      "height": 8927781
    }
  }
}
```

### HTTP Request

`POST /api.AccountService.IotexBalanceByHeight`

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

## ActiveAccounts

ActiveAccounts lists most recently active accounts

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.AccountService.ActiveAccounts \
  --header 'Content-Type: application/json' \
  --data '{
	"count": 5
}'
```

```graphql
query {
  ActiveAccounts(count:5){
    activeAccounts
  }
}



```

> Example response:

```json
{
  "data": {
    "ActiveAccounts": {
      "activeAccounts": [
        "io1aqf30kqz5rqh6zn82c00j684p2h2t5cg30wm8t",
        "io1lhukp867ume3qn2g7cxn4e47pj0ugfxeqj7nm8",
        "io12mgttmfa2ffn9uqvn0yn37f4nz43d248l2ga85",
        "io12p2td5p5tmaqqztdejl0dqdqalmajylw57x3e8",
        "io17cmrextyfeu4gddwd89g5qncedsnc553dhz7xa"
      ]
    }
  }
}
```

### HTTP Request

`POST /api.AccountService.ActiveAccounts`

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

## OperatorAddress

OperatorAddress finds the delegate's operator address given the delegate's alias name

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.AccountService.OperatorAddress \
  --header 'Content-Type: application/json' \
  --data '{
	"aliasName": "metanyxa"
}'
```

```graphql
query {
  OperatorAddress(aliasName:"metanyx") {
    exist
    operatorAddress
  }
}
```

> Example response:

```json
{
  "data": {
    "OperatorAddress": {
      "exist": true,
      "operatorAddress": "io10reczcaelglh5xmkay65h9vw3e5dp82e8vw0rz"
    }
  }
}
```

### HTTP Request

`POST /api.AccountService.OperatorAddress`

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

## Alias

Alias finds the delegate's alias name given the delegate's operator address

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.AccountService.Alias \
  --header 'Content-Type: application/json' \
  --data '{
	"operatorAddress": "io10reczcaelglh5xmkay65h9vw3e5dp82e8vw0rz"
}'
```

```graphql
query {
  Alias(operatorAddress:"io10reczcaelglh5xmkay65h9vw3e5dp82e8vw0rz") {
    exist
    aliasName
  }
}

```

> Example response:

```json
{
  "data": {
    "Alias": {
      "aliasName": "metanyx",
      "exist": true
    }
  }
}
```

### HTTP Request

`POST /api.AccountService.Alias`

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

## TotalNumberOfHolders

TotalNumberOfHolders returns total number of IOTX holders so far

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.AccountService.TotalNumberOfHolders \
  --header 'Content-Type: application/json' \
  --data '{
}'
```

```graphql
query {
  TotalNumberOfHolders{
    totalNumberOfHolders
  }
}
```

> Example response:

```json
{
  "data": {
    "TotalNumberOfHolders": {
      "totalNumberOfHolders": 511692
    }
  }
}
```

### HTTP Request

`POST /api.AccountService.TotalNumberOfHolders`

<a name="api-TotalNumberOfHoldersRequest"></a>

### TotalNumberOfHoldersRequest







<a name="api-TotalNumberOfHoldersResponse"></a>

### TotalNumberOfHoldersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| totalNumberOfHolders | [uint64](#uint64) |  | total number of IOTX holders so far |

## TotalAccountSupply

TotalAccountSupply returns total amount of tokens held by IoTeX accounts

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.AccountService.TotalAccountSupply \
  --header 'Content-Type: application/json' \
  --data '{
}'
```

```graphql
query {
  TotalAccountSupply{
    totalAccountSupply
  }
}
```

> Example response:

```json
{
  "data": {
    "TotalAccountSupply": {
      "totalAccountSupply": "12496299023824745920503427462"
    }
  }
}
```

### HTTP Request

`POST /api.AccountService.TotalAccountSupply`

<a name="api-TotalAccountSupplyRequest"></a>

### TotalAccountSupplyRequest







<a name="api-TotalAccountSupplyResponse"></a>

### TotalAccountSupplyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| totalAccountSupply | [string](#string) |  | total amount of tokens held by IoTeX accounts |

# Voting Service API

## CandidateInfo

CandidateInfo provides candidate information

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.VotingService.CandidateInfo \
  --header 'Content-Type: application/json' \
  --data '{
    "epochStart": 20000,
    "epochCount": 2
}'
```

```graphql
query {
  CandidateInfo(startEpoch: 20000, epochCount: 2) {
    candidateInfo {
      epochNumber
      candidates {
        name
        address
        totalWeightedVotes
        selfStakingTokens
        operatorAddress
        rewardAddress
      }
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "CandidateInfo": {
      "candidateInfo": [
        {
          "candidates": [
            {
              "address": "io1x3us2fktfq6tnftjzwtvgxh3ymvcfwy9fts7td",
              "name": "binancevote",
              "operatorAddress": "io1jzteq7gc5sh8tfp5auz8wwj97kvdapr9y8wzne",
              "rewardAddress": "io1x3us2fktfq6tnftjzwtvgxh3ymvcfwy9fts7td",
              "selfStakingTokens": "1230047466749539291090944",
              "totalWeightedVotes": "431578548518498595882908724"
            },
            ...
          ],
          "epochNumber": 20000
        },
        {
          "candidates": [
            {
              "address": "io1x3us2fktfq6tnftjzwtvgxh3ymvcfwy9fts7td",
              "name": "binancevote",
              "operatorAddress": "io1jzteq7gc5sh8tfp5auz8wwj97kvdapr9y8wzne",
              "rewardAddress": "io1x3us2fktfq6tnftjzwtvgxh3ymvcfwy9fts7td",
              "selfStakingTokens": "1230047466749539291090944",
              "totalWeightedVotes": "431556053359847416499574130"
            },
            ...
          ],
          "epochNumber": 20001
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.VotingService.CandidateInfo`

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

## RewardSources

RewardSources provides reward sources for voters 

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.VotingService.RewardSources \
  --header 'Content-Type: application/json' \
  --data '{
    "epochStart": 20000,
    "epochCount": 2,
    "voterIotxAddress": "io1rl62pepun2g7sed2tpv4tx7ujynye34fqjv40t"
}'
```

```graphql
query {
  RewardSources(
    startEpoch: 20000
    epochCount: 2
    voterIotxAddress: "io1rl62pepun2g7sed2tpv4tx7ujynye34fqjv40t"
  ) {
    exist
    delegateDistributions {
      delegateName
      amount
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "RewardSources": {
      "delegateDistributions": [
        {
          "amount": "940751518955182",
          "delegateName": "a4x"
        }
      ],
      "exist": true
    }
  }
}
```

### HTTP Request

`POST /api.VotingService.RewardSources`

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

# Action Service API

## ActionByDates

ActionByDates finds actions by dates

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ActionService.ActionByDates \
  --header 'Content-Type: application/json' \
  --data '{
    "startDate": 1624503172,
    "endDate": 1624503182,
    "pagination": {
		"skip": 0,
		"first": 2
	}
}'
```

```graphql
query {
  ActionByDates(
    startDate: 1624503172
    endDate: 1624503182
    pagination: { skip: 0, first: 1 }
  ) {
    exist
    count
    actions {
      actHash
      blkHash
      timestamp
      actType
      sender
      recipient
      amount
      gasFee
      blkHeight
    }
  }
}

```

> Example response:

```json
{
  "data": {
    "ActionByDates": {
      "actions": [
        {
          "actHash": "55cd01e165839ce6c9047bc5b2997808a6177ee23d7194ebe3c2bad419356a02",
          "actType": "grantReward",
          "amount": "0",
          "blkHash": "6ef7a5d37d15bf71d8a9bb9dac87177d6594214ec18f8ce929327382a8b5a54f",
          "blkHeight": 11792456,
          "gasFee": "0",
          "recipient": "",
          "sender": "io1ha87fd54jmgmes5eswsyd52gwm0qjxnnsqlyl0",
          "timestamp": 1624503175
        }
      ],
      "count": 4,
      "exist": true
    }
  }
}
```

### HTTP Request

`POST /api.ActionService.ActionByDates`

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

## ActionByHash

ActionByHash finds actions by hash

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ActionService.ActionByHash \
  --header 'Content-Type: application/json' \
  --data '{
    "actHash": "160a75d845c5ef35e6b2e697dc752066ee7a0dacf750c8c1a6a187090dd3df9f"
}'
```

```graphql
query {
  ActionByHash(
    actHash: "160a75d845c5ef35e6b2e697dc752066ee7a0dacf750c8c1a6a187090dd3df9f"
  ) {
    actionInfo {
      actHash
      blkHash
      timestamp
      actType
      sender
      recipient
      amount
      gasFee
      blkHeight
    }
    evmTransfers {
      sender
      recipient
      amount
    }
  }
}

```

> Example response:

```json
{
  "data": {
    "ActionByHash": {
      "actionInfo": {
        "actHash": "160a75d845c5ef35e6b2e697dc752066ee7a0dacf750c8c1a6a187090dd3df9f",
        "actType": "depositToStake",
        "amount": "173450647808345216",
        "blkHash": "c1504d3181b3065a780f196d601358c0017546d659c7c3324931f16c27e3f135",
        "blkHeight": 17667450,
        "gasFee": "10000000000000000",
        "recipient": "",
        "sender": "io1unvkgm98ma3r2fnfrhep24arjxf6kc8stx0nuc",
        "timestamp": 1653981420
      },
      "evmTransfers": [
        {
          "amount": "10000000000000000",
          "recipient": "io0000000000000000000000rewardingprotocol",
          "sender": "io1unvkgm98ma3r2fnfrhep24arjxf6kc8stx0nuc"
        },
        {
          "amount": "173450647808345216",
          "recipient": "io000000000000000000000000stakingprotocol",
          "sender": "io1unvkgm98ma3r2fnfrhep24arjxf6kc8stx0nuc"
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.ActionService.ActionByHash`

<a name="api-ActionByHashRequest"></a>

### ActionByHashRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actHash | [string](#string) |  | action hash |






<a name="api-ActionByHashResponse"></a>

### ActionByHashResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether actions exist within the time frame |
| actionInfo | [ActionInfo](#api-ActionInfo) |  |  |
| evmTransfers | [ActionByHashResponse.EvmTransfers](#api-ActionByHashResponse-EvmTransfers) | repeated |  |






<a name="api-ActionByHashResponse-EvmTransfers"></a>

### ActionByHashResponse.EvmTransfers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender address |
| recipient | [string](#string) |  | recipient address |
| amount | [string](#string) |  | amount transferred |

## ActionByAddress

ActionByAddress finds actions by address

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ActionService.ActionByAddress \
  --header 'Content-Type: application/json' \
  --data '{
    "address": "io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4",
    "pagination": {
      "skip": 0,
      "first": 5
    }
}'
```

```graphql
query {
  ActionByAddress(
    address: "io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4"
    pagination:{skip:0, first:1}
  ) {
    count
    actions {
      actHash
      blkHash
      timestamp
      actType
      sender
      recipient
      amount
      gasFee
      blkHeight
    }
  }
}

```

> Example response:

```json
{
  "data": {
    "ActionByAddress": {
      "actions": [
        {
          "actHash": "a9eb718db8ec8832badc6bc930e6e1f01717fc9ca2126693c7457b10340f3b73",
          "actType": "transfer",
          "amount": "1000000000000000000000",
          "blkHash": "e6d90aac3af1277ebaafc8e56945037c3c9500732a67472651970caf7dc2da14",
          "blkHeight": 16306635,
          "gasFee": "10000000000000000",
          "recipient": "io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4",
          "sender": "io1z0r07tl77tvphmd8rluuh8h2sa2xqdkzpsuvrh",
          "timestamp": 1647148170
        }
      ],
      "count": 13693
    }
  }
}
```

### HTTP Request

`POST /api.ActionService.ActionByAddress`

<a name="api-ActionByAddressRequest"></a>

### ActionByAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | sender address or recipient address |
| pagination | [pagination.Pagination](#pagination-Pagination) |  |  |






<a name="api-ActionByAddressResponse"></a>

### ActionByAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether actions exist for the given address |
| actions | [ActionInfo](#api-ActionInfo) | repeated |  |
| count | [uint64](#uint64) |  | total number of actions for the given address |

## ActionByType

ActionByType finds actions by action type

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ActionService.ActionByType \
  --header 'Content-Type: application/json' \
  --data '{
    "type": "transfer",
    "pagination": {
      "skip": 0,
      "first": 5
    }
}'
```

```graphql
query {
  ActionByType(
    type: "execution"
    pagination:{skip:0, first:1}
  ) {
    count
    actions {
      actHash
      blkHash
      timestamp
      actType
      sender
      recipient
      amount
      gasFee
      blkHeight
    }
  }
}

```

> Example response:

```json
{
  "data": {
    "ActionByType": {
      "actions": [
        {
          "actHash": "edf65e7ccbfb05e4fbd394db1acc276029c309994879e3a8c07023a753ea8886",
          "actType": "execution",
          "amount": "0",
          "blkHash": "ea06e52306ddcc02404427adcea7628a76a301c4a8f5f08b902a2ac672814292",
          "blkHeight": 5008,
          "gasFee": "1357294000000000000",
          "recipient": "",
          "sender": "io17ch0jth3dxqa7w9vu05yu86mqh0n6502d92lmp",
          "timestamp": 1555949160
        }
      ],
      "count": 15279705
    }
  }
}
```

### HTTP Request

`POST /api.ActionService.ActionByType`

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

## EvmTransfersByAddress

EvmTransfersByAddress finds EVM transfers by address

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ActionService.EvmTransfersByAddress \
  --header 'Content-Type: application/json' \
  --data '{
    "address": "io14yqd25kr6k4zss59u7sq9hme4r862yfpezf9dx",
    "pagination": {
      "skip": 0,
      "first": 5
    }
}'
```

```graphql
query {
  EvmTransfersByAddress(
    address: "io14yqd25kr6k4zss59u7sq9hme4r862yfpezf9dx"
    pagination: { skip: 0, first: 5 }
  ) {
    exist
    count
    evmTransfers {
      sender
      recipient
      actHash
      blkHash
      blkHeight
      amount
      timestamp
    }
  }
}

```

> Example response:

```json
{
  "data": {
    "EvmTransfersByAddress": {
      "count": 303,
      "evmTransfers": [
        {
          "actHash": "fdc21e798a151b866de7adf02b9b2d5e479b96a1865af4aeb108cd4caa528e9a",
          "amount": "20180406453780587",
          "blkHash": "589324fef069e1ad9899989e0587eb1fb0fc6d20311dcec63850ecb41923360a",
          "blkHeight": 18183325,
          "recipient": "io14yqd25kr6k4zss59u7sq9hme4r862yfpezf9dx",
          "sender": "io16y9wk2xnwurvtgmd2mds2gcdfe2lmzad6dcw29",
          "timestamp": 1656565175
        },
        ...
      ],
      "exist": true
    }
  }
}
```

### HTTP Request

`POST /api.ActionService.EvmTransfersByAddress`

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

# XRC20 Service API

## XRC20ByAddress

XRC20ByAddress returns Xrc20 actions given the sender address or recipient address

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.XRC20Service.XRC20ByAddress \
  --header 'Content-Type: application/json' \
  --data '{
    "address": "io1mhvlzj7y2t9y2dtzauyvyzzrvle6l7sekcf245",
    "pagination": {
		"skip": 0,
		"first": 2
	}
}'
```

```graphql
query {
  XRC20ByAddress(
    address: "io1mhvlzj7y2t9y2dtzauyvyzzrvle6l7sekcf245"
    pagination: { skip: 0, first: 1 }
  ) {
    exist
    count
    xrc20 {
      actHash
      contract
      sender
      amount
      recipient
      timestamp
    }
  }
}

```

> Example response:

```json
{
  "data": {
    "XRC20ByAddress": {
      "count": 4027,
      "exist": true,
      "xrc20": [
        {
          "actHash": "67b01c199e69164679b3fd5c34eb2e73d0d59840e28479321ad6ff15e4aa5c6d",
          "amount": "177600000000000000",
          "contract": "io1zl0el07pek4sly8dmscccnm0etd8xr8j02t4y7",
          "recipient": "io12w7agqdgwx7slp8fgcv7mnqvy3yf6j4tz0fnms",
          "sender": "io1mhvlzj7y2t9y2dtzauyvyzzrvle6l7sekcf245",
          "timestamp": 1656639570
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.XRC20Service.XRC20ByAddress`

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

## XRC20ByContractAddress

XRC20ByContractAddress returns Xrc20 actions given the Xrc20 contract address

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.XRC20Service.XRC20ByContractAddress \
  --header 'Content-Type: application/json' \
  --data '{
    "address": "io1gafy2msqmmmqyhrhk4dg3ghc59cplyhekyyu26",
    "pagination": {
		"skip": 0,
		"first": 2
	}
}'
```

```graphql
query {
  XRC20ByContractAddress(
    address: "io1gafy2msqmmmqyhrhk4dg3ghc59cplyhekyyu26"
    pagination: { skip: 0, first: 1 }
  ) {
    exist
    count
    xrc20 {
      actHash
      contract
      sender
      amount
      recipient
      timestamp
    }
  }
}

```

> Example response:

```json
{
  "data": {
    "XRC20ByContractAddress": {
      "count": 1531013,
      "exist": true,
      "xrc20": [
        {
          "actHash": "40336444117c08caa48c9171a566069dfe374512b6d44bd260911a04a8d7424d",
          "amount": "312959963682432085397",
          "contract": "io1gafy2msqmmmqyhrhk4dg3ghc59cplyhekyyu26",
          "recipient": "io1h9kmk9x0e6mzhwmtq5eljqnj80agphyvcheky0",
          "sender": "io19kuwwxhtmtfdk9fnsfn0qs8je38svn7dwe93s4",
          "timestamp": 1656635825
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.XRC20Service.XRC20ByContractAddress`

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

## XRC20ByPage

XRC20ByPage returns Xrc20 actions by pagination

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.XRC20Service.XRC20ByPage \
  --header 'Content-Type: application/json' \
  --data '{
    "pagination": {
		"skip": 0,
		"first": 2
	}
}'
```

```graphql
query {
  XRC20ByPage(
    pagination: { skip: 0, first: 1 }
  ) {
    exist
    count
    xrc20 {
      actHash
      contract
      sender
      amount
      recipient
      timestamp
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "XRC20ByPage": {
      "count": 18238044,
      "exist": true,
      "xrc20": [
        {
          "actHash": "078895bc0ce32304203ecaa1dc1c7294226b356f9ea79d64c2b75e2817e4dbce",
          "amount": "10000000000000000",
          "contract": "io1zl0el07pek4sly8dmscccnm0etd8xr8j02t4y7",
          "recipient": "io1t7pkdvadx2mrnfukzvhsr0xhc2nsjuq9sren7p",
          "sender": "io14xf3pqwydy9vpzxflpqcry75ne5f47654f2rn9",
          "timestamp": 1656642185
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.XRC20Service.XRC20ByPage`

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

## XRC20Addresses

Xrc20Addresses returns Xrc20 contract addresses

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.XRC20Service.XRC20Addresses \
  --header 'Content-Type: application/json' \
  --data '{
    "pagination": {
		"skip": 0,
		"first": 2
	}
}'
```

```graphql
query {
  XRC20Addresses(
    pagination: { skip: 0, first: 2 }
  ) {
    exist
    count
    addresses
  }
}

```

> Example response:

```json
{
  "data": {
    "XRC20Addresses": {
      "addresses": [
        "io10098dx8ntlqy7sshqkdl67a8xp39vqsyjh08pv",
        "io1009dgua7q8x63dpk95wncnp79fx9mz0ak65a6s"
      ],
      "count": 2174,
      "exist": true
    }
  }
}
```

### HTTP Request

`POST /api.XRC20Service.XRC20Addresses`

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

## XRC20TokenHolderAddresses

XRC20TokenHolderAddresses returns Xrc20 token holder addresses given a Xrc20 contract address

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.XRC20Service.XRC20TokenHolderAddresses \
  --header 'Content-Type: application/json' \
  --data '{
    "tokenAddress": "io1gafy2msqmmmqyhrhk4dg3ghc59cplyhekyyu26",
    "pagination": {
		"skip": 0,
		"first": 2
	}
}'
```

```graphql
query {
  XRC20TokenHolderAddresses(
    tokenAddress: "io1gafy2msqmmmqyhrhk4dg3ghc59cplyhekyyu26"
    pagination: { skip: 0, first: 5 }
  ) {
    count
    addresses
  }
}
```

> Example response:

```json
{
  "data": {
    "XRC20TokenHolderAddresses": {
      "addresses": [
        "io1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqd39ym7",
        "io19czfrdyjt67g9tuhmmcjkuh2m9qxzv5nqyve9p",
        "io15lchlhad6dya59fncqtcfm44njwxm6m0j29aq9",
        "io10mk2dqu0e86x7urmadc34v7m3alqqy5l5t822r",
        "io1w4686k0r3fkjqghk694j43csgp8w073ge3s0f0"
      ],
      "count": 11245
    }
  }
}
```

### HTTP Request

`POST /api.XRC20Service.XRC20TokenHolderAddresses`

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

# XRC721 Service API

## XRC721ByAddress

XRC721ByAddress returns Xrc721 actions given the sender address or recipient address

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.XRC721Service.XRC721ByAddress \
  --header 'Content-Type: application/json' \
  --data '{
    "address": "io1lutyka7aw7u872kzsujuz8pwn9qsrcjvs6e7jw",
    "pagination": {
		"skip": 0,
		"first": 2
	}
}'
```

```graphql
query {
  XRC721ByAddress(
    address: "io1lutyka7aw7u872kzsujuz8pwn9qsrcjvs6e7jw"
    pagination: { skip: 0, first: 6 }
  ) {
    exist
    count
    xrc721 {
      actHash
      contract
      sender
      amount
      recipient
      timestamp
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "XRC721ByAddress": {
      "count": 7,
      "exist": true,
      "xrc721": [
        {
          "actHash": "06de09b214d6f58f0af426388eed400b05e87078ef36b4dd723a6342f16afe61",
          "amount": "2273",
          "contract": "io1052s604n44atw5klykwff29tnrtsplqqkdajxf",
          "recipient": "io1lutyka7aw7u872kzsujuz8pwn9qsrcjvs6e7jw",
          "sender": "io1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqd39ym7",
          "timestamp": 1656631180
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.XRC721Service.XRC721ByAddress`

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

## XRC721ByContractAddress

XRC721ByContractAddress returns Xrc721 actions given the Xrc721 contract address

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.XRC721Service.XRC721ByContractAddress \
  --header 'Content-Type: application/json' \
  --data '{
    "address": "io1052s604n44atw5klykwff29tnrtsplqqkdajxf",
    "pagination": {
		"skip": 0,
		"first": 2
	}
}'
```

```graphql
query {
  XRC721ByContractAddress(
    address: "io1052s604n44atw5klykwff29tnrtsplqqkdajxf"
    pagination: { skip: 0, first: 1 }
  ) {
    exist
    count
    xrc721 {
      actHash
      contract
      sender
      amount
      recipient
      timestamp
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "XRC721ByContractAddress": {
      "count": 2279,
      "exist": true,
      "xrc721": [
        {
          "actHash": "06de09b214d6f58f0af426388eed400b05e87078ef36b4dd723a6342f16afe61",
          "amount": "2273",
          "contract": "io1052s604n44atw5klykwff29tnrtsplqqkdajxf",
          "recipient": "io1lutyka7aw7u872kzsujuz8pwn9qsrcjvs6e7jw",
          "sender": "io1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqd39ym7",
          "timestamp": 1656631180
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.XRC721Service.XRC721ByContractAddress`

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

## XRC721ByPage

XRC721ByPage returns Xrc721 actions by pagination

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.XRC721Service.XRC721ByPage \
  --header 'Content-Type: application/json' \
  --data '{
    "pagination": {
		"skip": 0,
		"first": 2
	}
}'
```

```graphql
query {
  XRC721ByPage(
    pagination: { skip: 0, first: 1 }
  ) {
    exist
    count
    xrc721 {
      actHash
      contract
      sender
      amount
      recipient
      timestamp
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "XRC721ByPage": {
      "count": 7251522,
      "exist": true,
      "xrc721": [
        {
          "actHash": "57353c167e2a3fda74b16949f02e8b339f573a0039af1d5164ff6bb6819ea5fe",
          "amount": "2063879",
          "contract": "io1asxdtswkr9p6r9du57ecrhrql865tf2qxue6hw",
          "recipient": "io1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqd39ym7",
          "sender": "io1vygjfdnvv2jrg5szw82fusf9ea4nzz2s4fpy6e",
          "timestamp": 1656660620
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.XRC721Service.XRC721ByPage`

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

## XRC721Addresses

Xrc20Addresses returns Xrc721 contract addresses

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.XRC721Service.XRC721Addresses \
  --header 'Content-Type: application/json' \
  --data '{
    "pagination": {
		"skip": 0,
		"first": 2
	}
}'
```

```graphql
query {
  XRC721Addresses(
    pagination: { skip: 0, first: 2 }
  ) {
    exist
    count
    addresses
  }
}
```

> Example response:

```json
{
  "data": {
    "XRC721Addresses": {
      "addresses": [
        "io104z744srvrvxdtk0kzp7hkquqyxkvyt9uqz6u7",
        "io1052s604n44atw5klykwff29tnrtsplqqkdajxf"
      ],
      "count": 253,
      "exist": true
    }
  }
}
```

### HTTP Request

`POST /api.XRC721Service.XRC721Addresses`

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

## XRC721TokenHolderAddresses

XRC721TokenHolderAddresses returns Xrc721 token holder addresses given a Xrc721 contract address

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.XRC721Service.XRC721TokenHolderAddresses \
  --header 'Content-Type: application/json' \
  --data '{
    "tokenAddress": "io1asxdtswkr9p6r9du57ecrhrql865tf2qxue6hw",
    "pagination": {
		"skip": 0,
		"first": 2
	}
}'
```

```graphql
query {
  XRC721TokenHolderAddresses(
    tokenAddress: "io1asxdtswkr9p6r9du57ecrhrql865tf2qxue6hw"
    pagination: { skip: 0, first: 5 }
  ) {
    count
    addresses
  }
}
```

> Example response:

```json
{
  "data": {
    "XRC721TokenHolderAddresses": {
      "addresses": [
        "io1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqd39ym7",
        "io1asxdtswkr9p6r9du57ecrhrql865tf2qxue6hw",
        "io19nja6v0jfxyxxt70xgrp248k429y5nkq6cstpq",
        "io1mt5u8vzxzuwq7hut99t9tvduhykzgkfclwlsgj",
        "io14qqv4xz8jzk3hexhmp9qdtdghdpamvyzuqajrd"
      ],
      "count": 6348
    }
  }
}
```

### HTTP Request

`POST /api.XRC721Service.XRC721TokenHolderAddresses`

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

# Hermes Service API

## Hermes

Hermes gives delegates who register the service of automatic reward distribution an overview of the reward distributions to their voters within a range of epochs

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.HermesService.Hermes \
  --header 'Content-Type: application/json' \
  --data '{
	"startEpoch": 22420,
	"epochCount": 1,
	"rewardAddress": "io12mgttmfa2ffn9uqvn0yn37f4nz43d248l2ga85"
}'
```

```graphql
query {
	Hermes(
		startEpoch: 22420
		epochCount: 1
		rewardAddress: "io12mgttmfa2ffn9uqvn0yn37f4nz43d248l2ga85"
	) {
		hermesDistribution {
			delegateName
			rewardDistribution {
				voterEthAddress
				voterIotexAddress
				amount
			}
			stakingIotexAddress
			voterCount
			waiveServiceFee
			refund
		}
	}
}

```

> Example response:

```json
{
	"hermesDistribution": [
		{
			"delegateName": "a4x",
			"rewardDistribution": [
				{
					"voterEthAddress": "0x009faf509551ea0784b27f14f00c79d972393302",
					"voterIotexAddress": "io1qz0675y4284q0p9j0u20qrrem9erjvczut23g2",
					"amount": "810850817586367"
				},
                ...
			],
			"stakingIotexAddress": "io1c2cacn26mawwg0vpx2ptnegg600q5kpmv75np0",
			"voterCount": "260",
			"waiveServiceFee": false,
			"refund": "5160457356723700049"
		},
        ...
}
```

### HTTP Request

`POST /api.HermesService.Hermes`

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

<a name="api-RewardDistribution"></a>

### RewardDistribution



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| voterEthAddress | [string](#string) |  | voter’s ERC20 address |
| voterIotexAddress | [string](#string) |  | voter’s IoTeX address |
| amount | [string](#string) |  | amount of reward distribution |

## HermesByVoter

HermesByVoter returns Hermes voters' receiving history

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.HermesService.HermesByVoter \
  --header 'Content-Type: application/json' \
  --data '{
	"startEpoch": 13752,
	"epochCount": 24,
	"voterAddress": "io13hlj049e96gpdxfr0atkhq3d6mhgzxx7mrmg00",
  "pagination": {
		"skip": 0,
		"first": 5
	}  
}'
```

```graphql
query {
  HermesByVoter(
    startEpoch: 13752
    epochCount: 24
    voterAddress: "io13hlj049e96gpdxfr0atkhq3d6mhgzxx7mrmg00",
    pagination:{skip: 0, first: 5}
  ){
    count
    exist
    delegates{
      delegateName
      fromEpoch
      toEpoch
      amount
      actHash
      timestamp
    }
    totalRewardReceived
  }
}

```

> Example response:

```json
{
  "data": {
    "HermesByVoter": {
      "count": 5,
      "delegates": [
        {
          "actHash": "4b9977a739c659967bc04729c71f5d10a0eb4368ad9c24d74f76251cbc174a0c",
          "amount": "188268667946402699",
          "delegateName": "iotexteam",
          "fromEpoch": 13728,
          "timestamp": 1605696600,
          "toEpoch": 13751
        },
        ....
      ],
      "exist": true,
      "totalRewardReceived": "946193665895943769"
    }
  }
}
```

### HTTP Request

`POST /api.HermesService.HermesByVoter`

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

## HermesByDelegate

HermesByDelegate returns Hermes delegates' distribution history

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.HermesService.HermesByDelegate \
  --header 'Content-Type: application/json' \
  --data '{
	"startEpoch": 20000,
	"epochCount": 1,
	"delegateName": "a4x"
}'
```

```graphql
query {
  HermesByDelegate(startEpoch: 20000, epochCount: 24, 
    delegateName: "a4x", pagination:{skip:0, first: 5}) {
    count
    exist
    distributionRatio{
      blockRewardRatio
      epochNumber
      epochRewardRatio
      foundationBonusRatio
    }
    totalRewardsDistributed
    voterInfoList{
      actionHash
      amount
      fromEpoch
      toEpoch
      timestamp
      voterAddress
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "HermesByDelegate": {
      "count": 13,
      "distributionRatio": [
        {
          "blockRewardRatio": 70,
          "epochNumber": 20000,
          "epochRewardRatio": 70,
          "foundationBonusRatio": 70
        },
        ...
      ],
      "exist": true,
      "totalRewardsDistributed": "374486361904633758261",
      "voterInfoList": [
        {
          "actHash": "911dd38de79541b3eb066ee7a03e35121004255323d62040b97f536c9906cdda",
          "amount": "105463180053251824",
          "fromEpoch": 19992,
          "timestamp": 1628538750,
          "toEpoch": 20015,
          "voterAddress": "io13hlj049e96gpdxfr0atkhq3d6mhgzxx7mrmg00"
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.HermesService.HermesByDelegate`

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


<a name="api-HermesByDelegateDistributionRatio"></a>

### HermesByDelegateDistributionRatio



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epochNumber | [uint64](#uint64) |  | epoch number |
| blockRewardRatio | [double](#double) |  | ratio of block reward being distributed |
| epochRewardRatio | [double](#double) |  | ratio of epoch reward being distributed |
| foundationBonusRatio | [double](#double) |  | ratio of foundation bonus being distributed |

## HermesMeta

HermesMeta provides Hermes platform metadata

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.HermesService.HermesMeta \
  --header 'Content-Type: application/json' \
  --data '{
	"startEpoch": 20000,
	"epochCount": 24
}'
```

```graphql
query {
  HermesMeta(startEpoch: 20000, epochCount: 24) {
    exist
    numberOfDelegates
    numberOfRecipients
    totalRewardDistributed
  }
}


```

> Example response:

```json
{
  "data": {
    "HermesMeta": {
      "exist": true,
      "numberOfDelegates": 41,
      "numberOfRecipients": 3871,
      "totalRewardDistributed": "342227806235546638245097"
    }
  }
}
```

### HTTP Request

`POST /api.HermesService.HermesMeta`

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

## HermesAverageStats

HermesAverageStats returns the Hermes average statistics

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.HermesService.HermesAverageStats \
  --header 'Content-Type: application/json' \
  --data '{
	"startEpoch": 20000,
	"epochCount": 24,
  "rewardAddress": "io12mgttmfa2ffn9uqvn0yn37f4nz43d248l2ga85"
}'
```

```graphql
query {
  HermesAverageStats(
    startEpoch: 20000
    epochCount: 24
    rewardAddress: "io12mgttmfa2ffn9uqvn0yn37f4nz43d248l2ga85"
  ) {
    exist
    averagePerEpoch {
      delegateName
      rewardDistribution
      totalWeightedVotes
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "HermesAverageStats": {
      "averagePerEpoch": [
        {
          "delegateName": "a4x",
          "rewardDistribution": "10939144171268859553",
          "totalWeightedVotes": "2897825174867095605694136"
        },
        ...
      ],
      "exist": true
    }
  }
}
```

### HTTP Request

`POST /api.HermesService.HermesAverageStats`

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