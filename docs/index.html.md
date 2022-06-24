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
	"epochCount": 5,
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

## HermesByDelegate

HermesByDelegate returns Hermes delegates' distribution history

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

`POST /api.DelegateService.HermesByDelegate`

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

# Account Service API

## Hermes

Hermes gives delegates who register the service of automatic reward distribution an overview of the reward distributions to their voters within a range of epochs

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.AccountService.Hermes \
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

`POST /api.AccountService.Hermes`

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