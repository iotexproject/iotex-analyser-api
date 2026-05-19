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

## API Token
API token is required to access the API. You can generate a JWT token using the `ioctl` command-line tool — see [ioctl JWT Auth Tokens](https://docs.iotex.io/depin/builders/reference-docs/ioctl-client/jwt-auth-tokens) for details.

```shell
ioctl jwt sign --with-arguments '{"exp":"1767024000","sub":"AnalyserAPI","scope":"Read"}' -s user
```

use generated token to access the API

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.Chain \
  --header 'Content-Type: application/json' \
  --data '{}' \
  --header 'Authorization: Bearer <TOKEN>'
```

<a name="pagination-Pagination"></a>

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
query{
  Chain{
    mostRecentEpoch
    mostRecentBlockHeight
    totalSupply
    totalCirculatingSupply
    exactCirculatingSupply
    totalCirculatingSupplyNoRewardPool
    votingResultMeta{
      totalCandidates
      totalWeightedVotes
      votedTokens
    }
    rewards{
      totalBalance
      totalAvailable
      totalUnclaimed
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "Chain": {
      "exactCirculatingSupply": "9252829326969999922855499899",
      "mostRecentBlockHeight": 18980166,
      "mostRecentEpoch": 28884,
      "rewards": {
        "totalAvailable": "440833151247362249902675121",
        "totalBalance": "498742289434892391257888603",
        "totalUnclaimed": "57909138187530141355213482"
      },
      "totalCirculatingSupply": "9457829322133866787170249906",
      "totalCirculatingSupplyNoRewardPool": "9016996170886504537267574785",
      "totalSupply": "9457829326969999922855499899",
      "votingResultMeta": {
        "totalCandidates": 75,
        "totalWeightedVotes": "3079386720450377171762829252",
        "votedTokens": "3743652944403254009755212860"
      }
    }
  }
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
| exactCirculatingSupply | [string](#string) |  | exact circulating supply |
| rewards | [ChainResponse.Rewards](#api-ChainResponse-Rewards) |  | rewards |






<a name="api-ChainResponse-Rewards"></a>

### ChainResponse.Rewards



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| totalBalance | [string](#string) |  | total balance |
| totalUnclaimed | [string](#string) |  | total unclaimed |
| totalAvailable | [string](#string) |  | total available |

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
	"startEpoch": 20000,
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


## BlockSizeByHeightRequest

BlockSizeByHeightRequest gives the block size and server version by block height

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.BlockSizeByHeight \
  --header 'Content-Type: application/json' \
  --data '{
	"height": 3233333
}'
```

```graphql
query{
  BlockSizeByHeight(height:3233333){
    serverVersion
    blockSize
  }
}

```

> Example response:

```json
{
  "data": {
    "BlockSizeByHeight": {
      "blockSize": 1707.578,
      "serverVersion": "0.10.0"
    }
  }
}
```

### HTTP Request

`POST /api.ChainService.BlockSizeByHeight`

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


## GetLatestBlockHeight

GetLatestBlockHeight returns the latest block height.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.GetLatestBlockHeight \
  --header 'Content-Type: application/json' \
  --data '{}'
```

```graphql
query {
  GetLatestBlockHeight {
    height
  }
}
```

> Example response:

```json
{
  "data": {
    "GetLatestBlockHeight": {
      "height": 19000000
    }
  }
}
```

### HTTP Request

`POST /api.ChainService.GetLatestBlockHeight`

<a name="api-GetLatestBlockHeightRequest"></a>

### GetLatestBlockHeightRequest

(no fields)

<a name="api-GetLatestBlockHeightResponse"></a>

### GetLatestBlockHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  | latest block height |

## GetBlocks

GetBlocks returns a paginated list of blocks. Use `before_height` as a cursor for client-controlled pagination.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.GetBlocks \
  --header 'Content-Type: application/json' \
  --data '{
  "page": 1,
  "limit": 10
}'
```

```graphql
query {
  GetBlocks(page: 1, limit: 10) {
    blocks {
      block_height
      block_hash
      producer_address
      num_actions
      timestamp
      gas_consumed
      producer_name
      block_reward
      epoch_num
      priority_bonus
      base_fee
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "GetBlocks": {
      "blocks": [
        {
          "base_fee": "1000000000",
          "block_hash": "abc123...",
          "block_height": 19000000,
          "block_reward": "900000000000000000",
          "epoch_num": 28900,
          "gas_consumed": 21000,
          "num_actions": 3,
          "priority_bonus": "0",
          "producer_address": "io1...",
          "producer_name": "iotexlab",
          "timestamp": 1714000000
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.ChainService.GetBlocks`

<a name="api-GetBlocksRequest"></a>

### GetBlocksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [uint64](#uint64) |  | page number (starts from 1); ignored when before_height > 0 |
| limit | [uint64](#uint64) |  | number of blocks per page |
| before_height | [uint64](#uint64) |  | cursor: return blocks with height <= before_height (0 = use page) |

<a name="api-GetBlocksResponse"></a>

### GetBlocksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| blocks | [BlockInfo](#api-BlockInfo) | repeated | list of blocks |

<a name="api-BlockInfo"></a>

### BlockInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| block_height | [uint64](#uint64) |  | block height |
| block_hash | [string](#string) |  | block hash |
| producer_address | [string](#string) |  | producer address |
| num_actions | [uint64](#uint64) |  | number of actions |
| timestamp | [int64](#int64) |  | block timestamp (unix) |
| gas_consumed | [uint64](#uint64) |  | gas consumed |
| producer_name | [string](#string) |  | producer name |
| block_reward | [string](#string) |  | block reward |
| epoch_num | [uint64](#uint64) |  | epoch number |
| priority_bonus | [string](#string) |  | priority bonus |
| base_fee | [string](#string) |  | base fee |

## GetBlockByHeight

GetBlockByHeight returns a single block by its height.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.GetBlockByHeight \
  --header 'Content-Type: application/json' \
  --data '{
  "height": 19000000
}'
```

```graphql
query {
  GetBlockByHeight(height: 19000000) {
    exist
    block {
      block_height
      block_hash
      producer_name
      num_actions
      timestamp
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "GetBlockByHeight": {
      "exist": true,
      "block": {
        "block_height": 19000000,
        "block_hash": "abc123...",
        "producer_name": "iotexlab",
        "num_actions": 3,
        "timestamp": 1714000000
      }
    }
  }
}
```

### HTTP Request

`POST /api.ChainService.GetBlockByHeight`

<a name="api-GetBlockByHeightRequest"></a>

### GetBlockByHeightRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  | block height |

<a name="api-GetBlockByHeightResponse"></a>

### GetBlockByHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether the block exists |
| block | [BlockInfo](#api-BlockInfo) |  | block information |

## GetEpochInfo

GetEpochInfo returns the current epoch number and its starting block height.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.GetEpochInfo \
  --header 'Content-Type: application/json' \
  --data '{}'
```

> Example response:

```json
{
  "epoch_height": 18980001,
  "epoch_num": 28884
}
```

### HTTP Request

`POST /api.ChainService.GetEpochInfo`

<a name="api-GetEpochInfoRequest"></a>

### GetEpochInfoRequest

(no fields)

<a name="api-GetEpochInfoResponse"></a>

### GetEpochInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epoch_height | [uint64](#uint64) |  | first block height of the current epoch |
| epoch_num | [uint64](#uint64) |  | current epoch number |

## GetLatestStakingRecord

GetLatestStakingRecord returns the most recent staking statistics.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.GetLatestStakingRecord \
  --header 'Content-Type: application/json' \
  --data '{}'
```

> Example response:

```json
{
  "total_supply": "9457829326969999922855499899",
  "all_staking": "3743652944403254009755212860",
  "staking_ratio": "0.3959"
}
```

### HTTP Request

`POST /api.ChainService.GetLatestStakingRecord`

<a name="api-GetLatestStakingRecordRequest"></a>

### GetLatestStakingRecordRequest

(no fields)

<a name="api-GetLatestStakingRecordResponse"></a>

### GetLatestStakingRecordResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total_supply | [string](#string) |  | total supply |
| all_staking | [string](#string) |  | total staked IOTX (in rau) |
| staking_ratio | [string](#string) |  | staking ratio (decimal string) |

## GetPeakTps

GetPeakTps returns the all-time peak TPS (max block action count / 5-second block time). Supports incremental cursor-based scanning.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.GetPeakTps \
  --header 'Content-Type: application/json' \
  --data '{
  "start_block_height": 0
}'
```

> Example response:

```json
{
  "num_actions": "42.60",
  "block_height": 19000000
}
```

### HTTP Request

`POST /api.ChainService.GetPeakTps`

<a name="api-GetPeakTpsRequest"></a>

### GetPeakTpsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start_block_height | [uint64](#uint64) |  | 0 = scan all blocks; > 0 = only look at blocks after this height (cursor for incremental updates) |

<a name="api-GetPeakTpsResponse"></a>

### GetPeakTpsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| num_actions | [string](#string) |  | peak TPS (max block actions / 5), rounded to 2 decimal places |
| block_height | [uint64](#uint64) |  | current max block height (cursor for next call) |

## GetActionHistory

GetActionHistory returns aggregated action counts bucketed by time interval.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.GetActionHistory \
  --header 'Content-Type: application/json' \
  --data '{
  "start_time": "2024-01-01 00:00:00",
  "end_time": "2024-01-02 00:00:00",
  "interval": "hour"
}'
```

> Example response:

```json
{
  "data": [
    { "date": "2024-01-01 00:00:00", "sum": 1234 },
    { "date": "2024-01-01 01:00:00", "sum": 987 }
  ]
}
```

### HTTP Request

`POST /api.ChainService.GetActionHistory`

<a name="api-GetActionHistoryRequest"></a>

### GetActionHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start_time | [string](#string) |  | UTC datetime string (e.g. "2024-01-01 00:00:00") |
| end_time | [string](#string) |  | UTC datetime string |
| interval | [string](#string) |  | bucket size: "minute", "hour", or "day" |

<a name="api-GetActionHistoryResponse"></a>

### GetActionHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [ActionHistoryPoint](#api-ActionHistoryPoint) | repeated | list of time-bucketed action counts |

<a name="api-ActionHistoryPoint"></a>

### ActionHistoryPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| date | [string](#string) |  | UTC datetime string |
| sum | [uint64](#uint64) |  | total actions in bucket |

## GetStakingRatioHistory

GetStakingRatioHistory returns the daily staking ratio over a time range.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.GetStakingRatioHistory \
  --header 'Content-Type: application/json' \
  --data '{
  "start_time": "2024-01-01",
  "end_time": "2024-01-31"
}'
```

> Example response:

```json
{
  "data": [
    { "date_time": "2024-01-01", "ratio": "0.3959" },
    { "date_time": "2024-01-02", "ratio": "0.3961" }
  ]
}
```

### HTTP Request

`POST /api.ChainService.GetStakingRatioHistory`

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
| data | [StakingRatioPoint](#api-StakingRatioPoint) | repeated | list of daily staking ratio data points |

<a name="api-StakingRatioPoint"></a>

### StakingRatioPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| date_time | [string](#string) |  | date |
| ratio | [string](#string) |  | staking ratio (decimal string) |

## GetChainStats

GetChainStats returns total action count + total/circulating supply (IOTX units). Values come from the pre-computed iotexscanv3_kv table (refreshed ~every 60 s by the iotex-statistics Windmill job).

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.GetChainStats \
  --header 'Content-Type: application/json' \
  --data '{}'
```

> Example response:

```json
{
  "actions_num": "213194125",
  "total_supply": "9441368502",
  "circulating_supply": "9441368498"
}
```

### HTTP Request

`POST /api.ChainService.GetChainStats`

<a name="api-GetChainStatsRequest"></a>

### GetChainStatsRequest

(empty)

<a name="api-GetChainStatsResponse"></a>

### GetChainStatsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actions_num | [uint64](#uint64) |  | total number of actions on chain |
| total_supply | [string](#string) |  | total supply in IOTX (rau / 1e18, integer) |
| circulating_supply | [string](#string) |  | circulating supply in IOTX (rau / 1e18, integer) |

## GetTpsHistory

GetTpsHistory returns daily avg/max TPS over a date range. Values are `AVG(num_actions)/2.5` and `MAX(num_actions)/2.5` aggregated from the block table; 2.5 s is the post-Wake block interval. Date range queries crossing the Wake hard fork (block 36893881) will overestimate TPS by 2× for pre-Wake days.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.GetTpsHistory \
  --header 'Content-Type: application/json' \
  --data '{
  "start": "2026-05-15",
  "end": "2026-05-18"
}'
```

> Example response:

```json
{
  "data": [
    { "date": "2026-05-15", "avg_tps": 1.19, "max_tps": 10.00 },
    { "date": "2026-05-16", "avg_tps": 1.23, "max_tps": 14.40 },
    { "date": "2026-05-17", "avg_tps": 1.29, "max_tps": 14.00 },
    { "date": "2026-05-18", "avg_tps": 1.17, "max_tps": 18.00 }
  ]
}
```

### HTTP Request

`POST /api.ChainService.GetTpsHistory`

<a name="api-GetTpsHistoryRequest"></a>

### GetTpsHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [string](#string) |  | start date, YYYY-MM-DD (UTC, inclusive) |
| end | [string](#string) |  | end date, YYYY-MM-DD (UTC, inclusive) |

<a name="api-GetTpsHistoryResponse"></a>

### GetTpsHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [TpsHistoryPoint](#api-TpsHistoryPoint) | repeated | daily TPS points |

<a name="api-TpsHistoryPoint"></a>

### TpsHistoryPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| date | [string](#string) |  | YYYY-MM-DD (UTC) |
| avg_tps | [double](#double) |  | average TPS, 2 decimal places |
| max_tps | [double](#double) |  | peak TPS observed in a single block that day, 2 decimal places |

## GetGasHistory

GetGasHistory returns daily gas-price stats and total gas fee. Aggregated live from block_action joined with block. Rows where gas_price = 0 (system actions) are excluded. Date ranges longer than ~30 days will be slow; keep ranges short.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.GetGasHistory \
  --header 'Content-Type: application/json' \
  --data '{
  "start": "2026-05-17",
  "end": "2026-05-18"
}'
```

> Example response (values are in rau; 1 IOTX = 10^18 rau):

```json
{
  "data": [
    {
      "date": "2026-05-17",
      "max_gas_price": "17897676565538",
      "min_gas_price": "1000000000000",
      "avg_gas_price": "1021883138757",
      "total_gas_fee": "7110675516295301992332"
    },
    {
      "date": "2026-05-18",
      "max_gas_price": "2029734718659213",
      "min_gas_price": "1000000000000",
      "avg_gas_price": "1160216242181",
      "total_gas_fee": "7597570410715163176791"
    }
  ]
}
```

### HTTP Request

`POST /api.ChainService.GetGasHistory`

<a name="api-GetGasHistoryRequest"></a>

### GetGasHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [string](#string) |  | start date, YYYY-MM-DD (UTC, inclusive) |
| end | [string](#string) |  | end date, YYYY-MM-DD (UTC, inclusive) |

<a name="api-GetGasHistoryResponse"></a>

### GetGasHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [GasHistoryPoint](#api-GasHistoryPoint) | repeated | daily gas-fee points |

<a name="api-GasHistoryPoint"></a>

### GasHistoryPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| date | [string](#string) |  | YYYY-MM-DD (UTC) |
| max_gas_price | [string](#string) |  | MAX(gas_price) for the day, in rau |
| min_gas_price | [string](#string) |  | MIN(gas_price) for the day, in rau |
| avg_gas_price | [string](#string) |  | ROUND(AVG(gas_price)), simple arithmetic mean, in rau |
| total_gas_fee | [string](#string) |  | SUM(gas_price * gas_consumed) for the day, in rau |

## GetSupplyHistory

GetSupplyHistory returns daily total/circulating supply and daily burn/issue over a date range. Burn/issue are derived from the day-over-day supply deltas (zero address only receives → ΔtotalSupply = −burn; lock address only sends → Δcirculating = −burn + issue). The first day in the range and any day after a chain halt or missing-block gap will have empty burn/issue.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ChainService.GetSupplyHistory \
  --header 'Content-Type: application/json' \
  --data '{
  "start": "2023-11-01",
  "end": "2023-11-03"
}'
```

> Example response (amounts in IOTX, 2 decimals):

```json
{
  "data": [
    {
      "date": "2023-11-01",
      "total_supply": "9443041459.47",
      "circulating_supply": "9443041454.63",
      "burn": "28125.00",
      "issue": "0.00"
    },
    {
      "date": "2023-11-02",
      "total_supply": "9443003959.47",
      "circulating_supply": "9443003954.63",
      "burn": "37500.00",
      "issue": "0.00"
    },
    {
      "date": "2023-11-03",
      "total_supply": "9442985209.47",
      "circulating_supply": "9442985204.63",
      "burn": "18750.00",
      "issue": "0.00"
    }
  ]
}
```

### HTTP Request

`POST /api.ChainService.GetSupplyHistory`

<a name="api-GetSupplyHistoryRequest"></a>

### GetSupplyHistoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [string](#string) |  | start date, YYYY-MM-DD (UTC, inclusive) |
| end | [string](#string) |  | end date, YYYY-MM-DD (UTC, inclusive) |

<a name="api-GetSupplyHistoryResponse"></a>

### GetSupplyHistoryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [SupplyHistoryPoint](#api-SupplyHistoryPoint) | repeated | daily supply points |

<a name="api-SupplyHistoryPoint"></a>

### SupplyHistoryPoint



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| date | [string](#string) |  | YYYY-MM-DD (UTC) |
| total_supply | [string](#string) |  | end-of-day total supply, IOTX, 2 decimals |
| circulating_supply | [string](#string) |  | end-of-day circulating supply, IOTX, 2 decimals |
| burn | [string](#string) |  | day-over-day burn, IOTX, 2 decimals; empty on first day / after gap |
| issue | [string](#string) |  | day-over-day issuance from lock address, IOTX, 2 decimals; empty on first day / after gap |

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
	"epochRewardPerc": 90,
	"blockRewardPerc": 100,
	"foundationBonusPerc": 0
}'
```

```graphql
query {
  BookKeeping(
    startEpoch: 23328
    epochCount: 10
    delegateName: "iotexlab"
    epochRewardPerc: 90
    foundationBonusPerc: 0
    blockRewardPerc: 100
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




## PaidToDelegates

PaidToDelegates returns the total reward amounts paid to each delegate for a given schedule (daily or monthly) and date.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.DelegateService.PaidToDelegates \
  --header 'Content-Type: application/json' \
  --data '{
  "schedule": "DAILY",
  "date": "2024-01-15"
}'
```

```graphql
query {
  PaidToDelegates(schedule: DAILY, date: "2024-01-15") {
    delegateInfo {
      delegateName
      amount
      blockReward
      epochReward
      foundationBonus
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "PaidToDelegates": {
      "delegateInfo": [
        {
          "delegateName": "iotexlab",
          "amount": "12000000000000000000",
          "blockReward": "4000000000000000000",
          "epochReward": "6000000000000000000",
          "foundationBonus": "2000000000000000000"
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.DelegateService.PaidToDelegates`

<a name="api-PaidToDelegatesRequest"></a>

### PaidToDelegatesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| schedule | [PaidToDelegatesRequest.Schedule](#api-PaidToDelegatesRequest-Schedule) |  | MONTHLY (0) or DAILY (1) |
| date | [string](#string) |  | date string (YYYY-MM-DD) |

<a name="api-PaidToDelegatesResponse"></a>

### PaidToDelegatesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegateInfo | [PaidToDelegatesResponse.DelegateInfo](#api-PaidToDelegatesResponse-DelegateInfo) | repeated | list of delegate reward info |

<a name="api-PaidToDelegatesResponse-DelegateInfo"></a>

### PaidToDelegatesResponse.DelegateInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegateName | [string](#string) |  | delegate name |
| amount | [string](#string) |  | total reward amount |
| blockReward | [string](#string) |  | block reward portion |
| epochReward | [string](#string) |  | epoch reward portion |
| foundationBonus | [string](#string) |  | foundation bonus portion |

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

## ContractInfo

ContractInfo returns contract info

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.AccountService.ContractInfo \
  --header 'Content-Type: application/json' \
  --data '{
    "contractAddress": ["io1yf0rvr34yxwjcx70t0x5rzz0skzurccy8wgpwe","io1mcy7wn2g3z7yue04385vzw7wnacamax8aaahn6"]
}'
```

```graphql
query {
  ContractInfo(
    contractAddress: [
      "io1yf0rvr34yxwjcx70t0x5rzz0skzurccy8wgpwe"
      "io1mcy7wn2g3z7yue04385vzw7wnacamax8aaahn6"
    ]
  ) {
    contracts {
			contractAddress
      deployer
      createTime
      exist
      callTimes
      accumulatedGas
    }
  }
}

```

> Example response:

```json
{
  "data": {
    "ContractInfo": {
      "contracts": [
        {
          "accumulatedGas": "",
          "callTimes": 0,
          "contractAddress": "io1yf0rvr34yxwjcx70t0x5rzz0skzurccy8wgpwe",
          "createTime": "",
          "deployer": "",
          "exist": false
        },
        {
          "accumulatedGas": "0.174381",
          "callTimes": 3,
          "contractAddress": "io1mcy7wn2g3z7yue04385vzw7wnacamax8aaahn6",
          "createTime": "2019-04-25 01:45:20 +0000 UTC",
          "deployer": "io10e0525sfrf53yh2aljmm3sn9jq5njk7l6jfauj",
          "exist": true
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.AccountService.ContractInfo`

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




## Erc20TokenBalanceByHeight

Erc20TokenBalanceByHeight returns the ERC20 token balance of given addresses at a specific block height.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.AccountService.Erc20TokenBalanceByHeight \
  --header 'Content-Type: application/json' \
  --data '{
  "address": ["io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4"],
  "height": 19000000,
  "contract_address": "io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw"
}'
```

```graphql
query {
  Erc20TokenBalanceByHeight(
    address: ["io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4"]
    height: 19000000
    contract_address: "io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw"
  ) {
    height
    contract_address
    balance
    decimals
  }
}
```

> Example response:

```json
{
  "data": {
    "Erc20TokenBalanceByHeight": {
      "height": 19000000,
      "contract_address": "io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw",
      "balance": ["1000000000000000000"],
      "decimals": 18
    }
  }
}
```

### HTTP Request

`POST /api.AccountService.Erc20TokenBalanceByHeight`

<a name="api-Erc20TokenBalanceByHeightRequest"></a>

### Erc20TokenBalanceByHeightRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) | repeated | list of addresses |
| height | [uint64](#uint64) |  | block height |
| contract_address | [string](#string) |  | ERC20 contract address |

<a name="api-Erc20TokenBalanceByHeightResponse"></a>

### Erc20TokenBalanceByHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  | block height |
| contract_address | [string](#string) |  | ERC20 contract address |
| balance | [string](#string) | repeated | balance list (in token units) |
| decimals | [uint64](#uint64) |  | token decimals |

## GetAccountMeta

GetAccountMeta returns account metadata including whether the address is a contract and its bytecode hash.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.AccountService.GetAccountMeta \
  --header 'Content-Type: application/json' \
  --data '{
  "addresses": ["io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4"]
}'
```

> Example response:

```json
{
  "accounts": [
    {
      "address": "io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4",
      "is_contract": false,
      "block_height": 0,
      "contract_bytecode_hash": ""
    }
  ]
}
```

### HTTP Request

`POST /api.AccountService.GetAccountMeta`

<a name="api-GetAccountMetaRequest"></a>

### GetAccountMetaRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| addresses | [string](#string) | repeated | list of addresses to query |

<a name="api-GetAccountMetaResponse"></a>

### GetAccountMetaResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accounts | [AccountMetaInfo](#api-AccountMetaInfo) | repeated | list of account metadata |

<a name="api-AccountMetaInfo"></a>

### AccountMetaInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | account address |
| is_contract | [bool](#bool) |  | whether the address is a contract |
| block_height | [uint64](#uint64) |  | block height at contract creation (0 for non-contracts) |
| contract_bytecode_hash | [string](#string) |  | SHA256 hash of contract bytecode |

## GetContractCreateInfo

GetContractCreateInfo returns the creation action hash and creator address for a given contract.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.AccountService.GetContractCreateInfo \
  --header 'Content-Type: application/json' \
  --data '{
  "address": "io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw"
}'
```

> Example response:

```json
{
  "action_hash": "abc123...",
  "creator": "io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4"
}
```

### HTTP Request

`POST /api.AccountService.GetContractCreateInfo`

<a name="api-GetContractCreateInfoRequest"></a>

### GetContractCreateInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | contract address |

<a name="api-GetContractCreateInfoResponse"></a>

### GetContractCreateInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action_hash | [string](#string) |  | creation action hash |
| creator | [string](#string) |  | creator address |

## GetAddressNFTBalances

GetAddressNFTBalances returns the NFT (XRC721/XRC1155) token balances held by an address.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.AccountService.GetAddressNFTBalances \
  --header 'Content-Type: application/json' \
  --data '{
  "address": "io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4"
}'
```

> Example response:

```json
{
  "balances": [
    {
      "contract_address": "io1nft...",
      "type": "xrc721",
      "balance": "3"
    }
  ]
}
```

### HTTP Request

`POST /api.AccountService.GetAddressNFTBalances`

<a name="api-GetAddressNFTBalancesRequest"></a>

### GetAddressNFTBalancesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | wallet address |

<a name="api-GetAddressNFTBalancesResponse"></a>

### GetAddressNFTBalancesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balances | [NFTBalanceInfo](#api-NFTBalanceInfo) | repeated | list of NFT balances per contract |

<a name="api-NFTBalanceInfo"></a>

### NFTBalanceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  | NFT contract address |
| type | [string](#string) |  | token standard: "xrc721" or "xrc1155" |
| balance | [string](#string) |  | balance |

## GetAddressTokenBalances

GetAddressTokenBalances returns the ERC20 token balances held by an address.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.AccountService.GetAddressTokenBalances \
  --header 'Content-Type: application/json' \
  --data '{
  "address": "io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4"
}'
```

> Example response:

```json
{
  "balances": [
    {
      "contract_address": "io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw",
      "balance": "1000000000000000000"
    }
  ]
}
```

### HTTP Request

`POST /api.AccountService.GetAddressTokenBalances`

<a name="api-GetAddressTokenBalancesRequest"></a>

### GetAddressTokenBalancesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | wallet address |

<a name="api-GetAddressTokenBalancesResponse"></a>

### GetAddressTokenBalancesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balances | [TokenBalanceInfo](#api-TokenBalanceInfo) | repeated | list of ERC20 token balances |

<a name="api-TokenBalanceInfo"></a>

### TokenBalanceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  | ERC20 contract address |
| balance | [string](#string) |  | balance (in raw units) |

## GetTopAccounts

GetTopAccounts returns top stakers filtered by stake amount, duration, and mf (momentum factor).

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.AccountService.GetTopAccounts \
  --header 'Content-Type: application/json' \
  --data '{
  "pagination": { "skip": 0, "first": 10 },
  "stake_amount": "more",
  "stake_duration": "more"
}'
```

> Example response:

```json
{
  "count": 1234,
  "accounts": [
    {
      "owner_address": "io1...",
      "bucket_id": 1001,
      "staked_amount": "10000000000000000000000",
      "duration": "91 days",
      "mf": "1.00",
      "last_update": "2024-01-15",
      "balance": "50000000000000000000000"
    }
  ]
}
```

### HTTP Request

`POST /api.AccountService.GetTopAccounts`

<a name="api-GetTopAccountsRequest"></a>

### GetTopAccountsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [pagination.Pagination](#pagination-Pagination) |  | pagination info |
| stake_amount | [string](#string) |  | "more" (>= 10000 IOTX) or "less" |
| stake_duration | [string](#string) |  | "more" (>= 91 days) or "less" |
| mf | [string](#string) |  | "hold" (mf > 0) or "" (mf = 0) |

<a name="api-GetTopAccountsResponse"></a>

### GetTopAccountsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [int64](#int64) |  | total number of matching accounts |
| accounts | [TopAccountRow](#api-TopAccountRow) | repeated | list of top accounts |

<a name="api-TopAccountRow"></a>

### TopAccountRow



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner_address | [string](#string) |  | account address |
| bucket_id | [uint64](#uint64) |  | staking bucket ID |
| staked_amount | [string](#string) |  | staked amount |
| duration | [string](#string) |  | stake duration |
| mf | [string](#string) |  | momentum factor |
| last_update | [string](#string) |  | last update time |
| balance | [string](#string) |  | IOTX balance |

## GetTopAccountsByBalance

GetTopAccountsByBalance returns the top accounts ranked by IOTX net balance (inflow minus outflow).

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.AccountService.GetTopAccountsByBalance \
  --header 'Content-Type: application/json' \
  --data '{
  "limit": 10,
  "offset": 0
}'
```

> Example response:

```json
{
  "count": 5000,
  "accounts": [
    {
      "address": "io1...",
      "balance": "123456789000000000000000",
      "total_actions": 500
    }
  ]
}
```

### HTTP Request

`POST /api.AccountService.GetTopAccountsByBalance`

<a name="api-GetTopAccountsByBalanceRequest"></a>

### GetTopAccountsByBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int64](#int64) |  | number of records to return |
| offset | [int64](#int64) |  | starting offset |

<a name="api-GetTopAccountsByBalanceResponse"></a>

### GetTopAccountsByBalanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [int64](#int64) |  | total number of accounts |
| accounts | [AccountBalanceRow](#api-AccountBalanceRow) | repeated | list of accounts |

<a name="api-AccountBalanceRow"></a>

### AccountBalanceRow



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | account address |
| balance | [string](#string) |  | net balance (inflow - outflow) in rau |
| total_actions | [int64](#int64) |  | total number of actions |

# Voting Service API

## CandidateInfo

CandidateInfo provides candidate information

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.VotingService.CandidateInfo \
  --header 'Content-Type: application/json' \
  --data '{
    "startEpoch": 20000,
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
    "startEpoch": 20000,
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

## VotingMeta

VotingMeta provides metadata of voting results

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.VotingService.VotingMeta \
  --header 'Content-Type: application/json' \
  --data '{
    "startEpoch": 20000,
    "epochCount": 2
}'
```

```graphql
query {
  VotingMeta(startEpoch: 20000, epochCount: 2) {
    exist
    candidateMeta {
      epochNumber
      totalCandidates
      consensusDelegates
      totalWeightedVotes
      votedTokens
    }
  }
}

```

> Example response:

```json
{
  "data": {
    "VotingMeta": {
      "candidateMeta": [
        {
          "consensusDelegates": 36,
          "epochNumber": 20000,
          "totalCandidates": 58,
          "totalWeightedVotes": "3497033939381331462899534869",
          "votedTokens": "2828696111178324496444652661"
        },
        {
          "consensusDelegates": 36,
          "epochNumber": 20001,
          "totalCandidates": 58,
          "totalWeightedVotes": "3497019706023954289836152135",
          "votedTokens": "2828703742673077328930935452"
        }
      ],
      "exist": true
    }
  }
}
```

### HTTP Request

`POST /api.VotingService.VotingMeta`

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



## GetCurrentDelegates

GetCurrentDelegates returns the current delegate list ordered by vote weight.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.VotingService.GetCurrentDelegates \
  --header 'Content-Type: application/json' \
  --data '{}'
```

```graphql
query {
  GetCurrentDelegates {
    exist
    delegates {
      id
      name
      vote_weight
      productivity
      candidate
      operator_address
      active
      block_height
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "GetCurrentDelegates": {
      "exist": true,
      "delegates": [
        {
          "id": 1,
          "name": "iotexlab",
          "vote_weight": "3079386720450377171762829252",
          "productivity": 0.9998,
          "candidate": "io1...",
          "operator_address": "io1...",
          "active": true,
          "block_height": 19000000
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.VotingService.GetCurrentDelegates`

<a name="api-GetCurrentDelegatesRequest"></a>

### GetCurrentDelegatesRequest

(no fields)

<a name="api-GetCurrentDelegatesResponse"></a>

### GetCurrentDelegatesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether delegates exist |
| delegates | [CurrentDelegateInfo](#api-CurrentDelegateInfo) | repeated | list of current delegates |

<a name="api-CurrentDelegateInfo"></a>

### CurrentDelegateInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | internal id |
| name | [string](#string) |  | delegate name |
| vote_weight | [string](#string) |  | total vote weight |
| productivity | [double](#double) |  | block production productivity |
| candidate | [string](#string) |  | candidate address |
| operator_address | [string](#string) |  | operator address |
| active | [bool](#bool) |  | whether currently active |
| block_height | [uint64](#uint64) |  | block height of last update |

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
| gasPrice | [string](#string) |  | gas price |
| gasLimit | [uint64](#uint64) |  | gas limit |
| gasConsumed | [uint64](#uint64) |  | gas consumed |
| nonce | [uint64](#uint64) |  | nonce |
| status | [uint64](#uint64) |  | execution status |
| contractAddress | [string](#string) |  | contract address |
| executionRevertMsg | [string](#string) |  | execution revert message |
| chainId | [uint64](#uint64) |  | chain id |
| methodName | [string](#string) |  | method name |

## ActionByHash

ActionByHash finds actions by hash. Use the optional `include_fields` parameter to request additional data (each field triggers a separate DB query).

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ActionService.ActionByHash \
  --header 'Content-Type: application/json' \
  --data '{
    "actHash": "160a75d845c5ef35e6b2e697dc752066ee7a0dacf750c8c1a6a187090dd3df9f",
    "include_fields": ["action_type", "logs", "token_transfers"]
}'
```

```graphql
query {
  ActionByHash(
    actHash: "160a75d845c5ef35e6b2e697dc752066ee7a0dacf750c8c1a6a187090dd3df9f"
    include_fields: ["action_type", "logs", "token_transfers"]
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
      methodName
    }
    evmTransfers {
      sender
      recipient
      amount
    }
    action_type_info {
      type
      gas_tip_cap
      gas_fee_cap
    }
    logs {
      block_height
      address
      topic0
      data
    }
    token_transfers {
      contract_address
      sender
      recipient
      amount
      type
    }
    block_base_fee
    stake_action {
      bucket_id
      amount
      candidate
      act_type
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
| include_fields | [string](#string) | repeated | optional subset of extra fields to fetch: `action_type`, `input_data`, `logs`, `token_transfers`, `base_fee`, `stake_action`; empty = none |






<a name="api-ActionByHashResponse"></a>

### ActionByHashResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether action exists |
| actionInfo | [ActionInfo](#api-ActionInfo) |  |  |
| evmTransfers | [ActionByHashResponse.EvmTransfers](#api-ActionByHashResponse-EvmTransfers) | repeated |  |
| action_type_info | [ActionByHashResponse.ActionTypeInfo](#api-ActionByHashResponse-ActionTypeInfo) |  | EIP-2930/1559/4844 type info (requires `action_type` in include_fields) |
| input_data | [string](#string) |  | hex-encoded execution input data (requires `input_data` in include_fields) |
| logs | [ActionByHashResponse.ActionLog](#api-ActionByHashResponse-ActionLog) | repeated | transaction logs (requires `logs` in include_fields) |
| token_transfers | [ActionByHashResponse.TokenTransfer](#api-ActionByHashResponse-TokenTransfer) | repeated | ERC20/NFT token transfers (requires `token_transfers` in include_fields) |
| block_base_fee | [string](#string) |  | block base fee in rau (requires `base_fee` in include_fields) |
| stake_action | [ActionByHashResponse.StakeAction](#api-ActionByHashResponse-StakeAction) |  | staking action details (requires `stake_action` in include_fields) |






<a name="api-ActionByHashResponse-EvmTransfers"></a>

### ActionByHashResponse.EvmTransfers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender | [string](#string) |  | sender address |
| recipient | [string](#string) |  | recipient address |
| amount | [string](#string) |  | amount transferred |

<a name="api-ActionByHashResponse-ActionTypeInfo"></a>

### ActionByHashResponse.ActionTypeInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | transaction type |
| access_list | [string](#string) |  | EIP-2930 access list (JSON) |
| gas_tip_cap | [string](#string) |  | EIP-1559 gas tip cap |
| gas_fee_cap | [string](#string) |  | EIP-1559 gas fee cap |
| blob_gas | [string](#string) |  | EIP-4844 blob gas |
| blob_fee_cap | [string](#string) |  | EIP-4844 blob fee cap |
| blob_hashes | [string](#string) |  | EIP-4844 blob hashes |
| blob_gas_price | [string](#string) |  | EIP-4844 blob gas price |

<a name="api-ActionByHashResponse-ActionLog"></a>

### ActionByHashResponse.ActionLog



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| block_height | [uint64](#uint64) |  | block height |
| address | [string](#string) |  | contract address emitting the log |
| topic0 | [string](#string) |  | first log topic |
| topic1 | [string](#string) |  | second log topic |
| topic2 | [string](#string) |  | third log topic |
| topic3 | [string](#string) |  | fourth log topic |
| data | [bytes](#bytes) |  | raw log data (base64-encoded in JSON) |
| action_hash | [string](#string) |  | action hash |
| index | [int64](#int64) |  | log index in block |

<a name="api-ActionByHashResponse-TokenTransfer"></a>

### ActionByHashResponse.TokenTransfer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | record id |
| contract_address | [string](#string) |  | token contract address |
| sender | [string](#string) |  | sender address |
| recipient | [string](#string) |  | recipient address |
| amount | [string](#string) |  | token amount |
| type | [string](#string) |  | transfer type: "erc20" or "nft" |

<a name="api-ActionByHashResponse-StakeAction"></a>

### ActionByHashResponse.StakeAction



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bucket_id | [int64](#int64) |  | staking bucket ID |
| amount | [string](#string) |  | action amount |
| staked_amount | [string](#string) |  | staked amount in bucket |
| duration | [string](#string) |  | stake duration |
| auto_stake | [bool](#bool) |  | whether auto-stake is enabled |
| candidate | [string](#string) |  | candidate name |
| act_type | [string](#string) |  | stake action type |
| owner_address | [string](#string) |  | bucket owner address |

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
| sender | [string](#string) |  | optional: filter by sender address |
| recipient | [string](#string) |  | optional: filter by recipient address |
| actionType | [string](#string) |  | optional: filter by action type |
| startTime | [string](#string) |  | optional: start time filter (ISO 8601) |
| endTime | [string](#string) |  | optional: end time filter (ISO 8601) |






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


## ActionByVoter

ActionByVoter returns actions by voter address.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ActionService.ActionByVoter \
  --header 'Content-Type: application/json' \
  --data '{
  "address": "io19msajm9hv4u793jvnwcy23plkwzffywjh257sz",
  "pagination": { "skip": 0, "first": 10 }
}'
```

```graphql
query {
  ActionByVoter(
    address: "io19msajm9hv4u793jvnwcy23plkwzffywjh257sz"
    pagination: { skip: 0, first: 10 }
  ) {
    exist
    count
    actionList {
      actHash
      actType
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
    "ActionByVoter": {
      "exist": true,
      "count": 42,
      "actionList": [
        {
          "actHash": "abc123...",
          "actType": "stakeCreate",
          "amount": "100000000000000000000",
          "timestamp": 1714000000
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.ActionService.ActionByVoter`

<a name="api-ActionByVoterRequest"></a>

### ActionRequest (ActionByVoter)



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | voter address |
| actHash | [string](#string) |  | action hash (optional) |
| pagination | [pagination.Pagination](#pagination-Pagination) |  | pagination info |

<a name="api-ActionByVoterResponse"></a>

### ActionResponse (ActionByVoter)



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether actions exist |
| count | [uint64](#uint64) |  | total number of actions |
| actionList | [ActionInfo](#api-ActionInfo) | repeated | list of actions |
| evmTransferList | [EvmTransferInfo](#api-EvmTransferInfo) | repeated | list of EVM transfers |
| xrcList | [XrcInfo](#api-XrcInfo) | repeated | list of XRC20 transfers |

## GetXrc20ByAddress

GetXrc20ByAddress returns XRC20 transfers by address (ActionService version).

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ActionService.GetXrc20ByAddress \
  --header 'Content-Type: application/json' \
  --data '{
  "address": "io19msajm9hv4u793jvnwcy23plkwzffywjh257sz",
  "pagination": { "skip": 0, "first": 10 }
}'
```

```graphql
query {
  GetXrc20ByAddress(
    address: "io19msajm9hv4u793jvnwcy23plkwzffywjh257sz"
    pagination: { skip: 0, first: 10 }
  ) {
    exist
    count
    xrcList {
      actHash
      from
      to
      quantity
      timestamp
      contract
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "GetXrc20ByAddress": {
      "exist": true,
      "count": 25,
      "xrcList": [
        {
          "actHash": "abc123...",
          "from": "io1...",
          "to": "io1...",
          "quantity": "1000000000000000000",
          "timestamp": 1714000000,
          "contract": "io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw"
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.ActionService.GetXrc20ByAddress`

<a name="api-XrcInfo"></a>

### XrcInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actHash | [string](#string) |  | action hash |
| from | [string](#string) |  | sender address |
| to | [string](#string) |  | recipient address |
| quantity | [string](#string) |  | token amount |
| blkHeight | [uint64](#uint64) |  | block height |
| timestamp | [uint64](#uint64) |  | unix timestamp |
| contract | [string](#string) |  | token contract address |

## ActionList

ActionList returns the latest actions with pagination. Use `start_block_height` to enable PostgreSQL partition pruning for faster queries.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ActionService.ActionList \
  --header 'Content-Type: application/json' \
  --data '{
  "pagination": { "skip": 0, "first": 20 },
  "start_block_height": 18000000
}'
```

```graphql
query {
  ActionList(
    pagination: { skip: 0, first: 20 }
    start_block_height: 18000000
  ) {
    exist
    count
    actions {
      actHash
      actType
      sender
      recipient
      amount
      blkHeight
      timestamp
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "ActionList": {
      "exist": true,
      "count": 1000000,
      "actions": [
        {
          "actHash": "abc123...",
          "actType": "transfer",
          "sender": "io1...",
          "recipient": "io1...",
          "amount": "1000000000000000000",
          "blkHeight": 19000000,
          "timestamp": 1714000000
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.ActionService.ActionList`

<a name="api-ActionListRequest"></a>

### ActionListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [pagination.Pagination](#pagination-Pagination) |  | pagination info |
| start_block_height | [uint64](#uint64) |  | optional: enables partition pruning for queries starting at this block height |

<a name="api-ActionListResponse"></a>

### ActionListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether actions exist |
| count | [uint64](#uint64) |  | total number of actions |
| actions | [ActionInfo](#api-ActionInfo) | repeated | list of actions |

## ActionByHeight

ActionByHeight returns actions at a specific block height.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ActionService.ActionByHeight \
  --header 'Content-Type: application/json' \
  --data '{
  "height": 19000000,
  "pagination": { "skip": 0, "first": 50 }
}'
```

```graphql
query {
  ActionByHeight(height: 19000000, pagination: { skip: 0, first: 50 }) {
    exist
    count
    actions {
      actHash
      actType
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
    "ActionByHeight": {
      "exist": true,
      "count": 3,
      "actions": [
        {
          "actHash": "abc123...",
          "actType": "transfer",
          "sender": "io1...",
          "recipient": "io1...",
          "amount": "1000000000000000000"
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.ActionService.ActionByHeight`

<a name="api-ActionByHeightRequest"></a>

### ActionByHeightRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  | block height |
| pagination | [pagination.Pagination](#pagination-Pagination) |  | pagination info |

<a name="api-ActionByHeightResponse"></a>

### ActionByHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether actions exist at this height |
| count | [uint64](#uint64) |  | total number of actions |
| actions | [ActionInfo](#api-ActionInfo) | repeated | list of actions |

## ContractInteractors

ContractInteractors returns the distinct sender addresses that have interacted with a given contract, optionally filtered by start time.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ActionService.ContractInteractors \
  --header 'Content-Type: application/json' \
  --data '{
  "address": "io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw",
  "startTime": "2024-01-01T00:00:00Z"
}'
```

```graphql
query {
  ContractInteractors(
    address: "io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw"
    startTime: "2024-01-01T00:00:00Z"
  ) {
    senders
  }
}
```

> Example response:

```json
{
  "data": {
    "ContractInteractors": {
      "senders": [
        "io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4",
        "io1..."
      ]
    }
  }
}
```

### HTTP Request

`POST /api.ActionService.ContractInteractors`

<a name="api-ContractInteractorsRequest"></a>

### ContractInteractorsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | contract address |
| startTime | [string](#string) |  | optional: filter interactions after this time (ISO 8601) |

<a name="api-ContractInteractorsResponse"></a>

### ContractInteractorsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| senders | [string](#string) | repeated | list of distinct sender addresses |

## GetInternalTxns

GetInternalTxns returns a paginated list of EVM internal transactions.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ActionService.GetInternalTxns \
  --header 'Content-Type: application/json' \
  --data '{
  "pagination": { "skip": 0, "first": 20 }
}'
```

> Example response:

```json
{
  "txns": [
    {
      "id": 1001,
      "block_height": 19000000,
      "action_hash": "abc123...",
      "type": "execution",
      "amount": "1000000000000000000",
      "sender": "io1...",
      "recipient": "io1...",
      "timestamp": "2024-01-15T10:00:00Z"
    }
  ],
  "count": 5000000
}
```

### HTTP Request

`POST /api.ActionService.GetInternalTxns`

<a name="api-GetInternalTxnsRequest"></a>

### GetInternalTxnsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [pagination.Pagination](#pagination-Pagination) |  | pagination info |

<a name="api-GetInternalTxnsResponse"></a>

### GetInternalTxnsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| txns | [InternalTxnInfo](#api-InternalTxnInfo) | repeated | list of internal transactions |
| count | [uint64](#uint64) |  | total number of internal transactions |

<a name="api-InternalTxnInfo"></a>

### InternalTxnInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | record id |
| block_height | [uint64](#uint64) |  | block height |
| action_hash | [string](#string) |  | action hash |
| type | [string](#string) |  | transaction type |
| amount | [string](#string) |  | transfer amount |
| sender | [string](#string) |  | sender address |
| recipient | [string](#string) |  | recipient address |
| timestamp | [string](#string) |  | timestamp |

## GetStakingActionsByAddress

GetStakingActionsByAddress returns paginated staking actions for a given owner address.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ActionService.GetStakingActionsByAddress \
  --header 'Content-Type: application/json' \
  --data '{
  "owner_address": "io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4",
  "pagination": { "skip": 0, "first": 20 }
}'
```

> Example response:

```json
{
  "actions": [
    {
      "id": 1,
      "block_height": 19000000,
      "action_hash": "abc123...",
      "sender": "io1...",
      "amount": "10000000000000000000000",
      "action_type": "stakeCreate",
      "timestamp": "2024-01-15T10:00:00Z"
    }
  ],
  "count": 50
}
```

### HTTP Request

`POST /api.ActionService.GetStakingActionsByAddress`

<a name="api-GetStakingActionsByAddressRequest"></a>

### GetStakingActionsByAddressRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner_address | [string](#string) |  | bucket owner address |
| pagination | [pagination.Pagination](#pagination-Pagination) |  | pagination info |

<a name="api-GetStakingActionsByAddressResponse"></a>

### GetStakingActionsByAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actions | [StakingActionInfo](#api-StakingActionInfo) | repeated | list of staking actions |
| count | [uint64](#uint64) |  | total number of staking actions |

<a name="api-StakingActionInfo"></a>

### StakingActionInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | record id |
| block_height | [uint64](#uint64) |  | block height |
| action_hash | [string](#string) |  | action hash |
| sender | [string](#string) |  | sender address |
| amount | [string](#string) |  | action amount |
| action_type | [string](#string) |  | staking action type |
| timestamp | [string](#string) |  | timestamp |

# Staking Service API

## VoteByHeight

Get the stake amount and voting weight of the voter's specified height

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.StakingService.VoteByHeight \
  --header 'Content-Type: application/json' \
  --data '{
    "address": ["io1k0w5vtlglnm742jacv2xjczg5l3s44gyy6h536","io1p8dwgkkgzcpwz6snwzh5v2nhf658gw8aen47pw"],
    "height": 20924943
}'
```

```graphql
query {
  VoteByHeight(
    address: [
      "io1k0w5vtlglnm742jacv2xjczg5l3s44gyy6h536"
      "io1p8dwgkkgzcpwz6snwzh5v2nhf658gw8aen47pw"
    ]
    height: 20924943
  ) {
    stakeAmount
    height
    voteWeight
  }
}
```

> Example response:

```json
{
  "data": {
    "VoteByHeight": {
      "height": 20924943,
      "stakeAmount": [
        "495.994311518986506235",
        "51127.627429245086002541"
      ],
      "voteWeight": [
        "631.859744093707405579",
        "60543.734008399268885716"
      ]
    }
  }
}
```

### HTTP Request

`POST /api.StakingService.VoteByHeight`

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

## BucketByID

Get staking bucket details by bucket IDs at a given block height. Set `includeSystem` to `true` to also return system staking buckets (v1/v2/v3). All `duration` values are returned in seconds.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.StakingService.BucketByID \
  --header 'Content-Type: application/json' \
  --data '{
    "bucketId": [1, 2],
    "height": 0,
    "includeSystem": false
}'
```

```graphql
query {
  BucketByID(
    bucketId: [1, 2]
    height: 0
    includeSystem: false
  ) {
    height
    nativeBuckets {
      bucketId
      ownerAddress
      candidate
      stakedAmount
      votingPower
      duration
      autoStake
      createTime
      stakeStartTime
      unstakeStartTime
      blockHeight
    }
  }
}
```

> Example response:

```json
{
  "height": 37000000,
  "nativeBuckets": [
    {
      "bucketId": "1",
      "ownerAddress": "io1jh0ekmccywfkmj7e8qsuzsupnlk3w5337hjjg",
      "candidate": "io1xpq62aw85uqzrccg9y5hnryv8ld2nkpycc3gza",
      "stakedAmount": "1200000000000000000000000",
      "votingPower": "2012011090252587925504000",
      "duration": 7862400,
      "autoStake": true,
      "createTime": 1553591580,
      "stakeStartTime": 1553591580,
      "unstakeStartTime": 0,
      "blockHeight": 36999980
    }
  ]
}
```

### HTTP Request

`POST /api.StakingService.BucketByID`

<a name="api-BucketByIDRequest"></a>

### BucketByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bucketId | [uint64](#uint64) | repeated | list of bucket IDs to query |
| height | [uint64](#uint64) |  | block height (0 = latest indexed height) |
| includeSystem | [bool](#bool) |  | whether to include system buckets in response (default false) |

<a name="api-BucketByIDResponse"></a>

### BucketByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  | block height of the query |
| nativeBuckets | [StakingBucketInfo](#api-StakingBucketInfo) | repeated | native staking buckets |
| systemBuckets | [StakingBucketInfo](#api-StakingBucketInfo) | repeated | system staking buckets v1 (only when includeSystem=true) |
| systemV2Buckets | [StakingBucketInfo](#api-StakingBucketInfo) | repeated | system staking buckets v2 (only when includeSystem=true) |
| systemV3Buckets | [StakingBucketInfo](#api-StakingBucketInfo) | repeated | system staking buckets v3 (only when includeSystem=true) |

<a name="api-StakingBucketInfo"></a>

### StakingBucketInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bucketId | [uint64](#uint64) |  | unique bucket ID |
| ownerAddress | [string](#string) |  | bucket owner IoTeX address |
| candidate | [string](#string) |  | delegate address receiving the vote |
| stakedAmount | [string](#string) |  | total staked amount in Rau (decimal string) |
| votingPower | [string](#string) |  | voting power in Rau (decimal string) |
| duration | [uint32](#uint32) |  | lock duration in seconds |
| autoStake | [bool](#bool) |  | whether auto-stake is enabled |
| createTime | [uint32](#uint32) |  | bucket creation time (unix timestamp) |
| stakeStartTime | [uint32](#uint32) |  | staking start time (unix timestamp) |
| unstakeStartTime | [uint32](#uint32) |  | unstaking start time (unix timestamp, 0 if not unstaking) |
| blockHeight | [uint64](#uint64) |  | block height of this bucket state |


## CandidateVoteByHeight

CandidateVoteByHeight returns the stake amount and voting weight for candidate addresses at a given block height.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.StakingService.CandidateVoteByHeight \
  --header 'Content-Type: application/json' \
  --data '{
  "address": ["io1..."],
  "height": 19000000
}'
```

```graphql
query {
  CandidateVoteByHeight(address: ["io1..."], height: 19000000) {
    height
    stakeAmount
    voteWeight
    address
  }
}
```

> Example response:

```json
{
  "data": {
    "CandidateVoteByHeight": {
      "height": 19000000,
      "stakeAmount": ["10000000000000000000000"],
      "voteWeight": ["12000000000000000000000"],
      "address": ["io1..."]
    }
  }
}
```

### HTTP Request

`POST /api.StakingService.CandidateVoteByHeight`

<a name="api-CandidateVoteByHeightRequest"></a>

### CandidateVoteByHeightRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) | repeated | candidate address list |
| height | [uint64](#uint64) |  | block height |

<a name="api-CandidateVoteByHeightResponse"></a>

### CandidateVoteByHeightResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| height | [uint64](#uint64) |  | block height |
| stakeAmount | [string](#string) | repeated | stake amount list |
| voteWeight | [string](#string) | repeated | vote weight list |
| address | [string](#string) | repeated | address list |

## GetBucketList

GetBucketList returns a paginated list of staking buckets with sorting and interval filters. Supports multiple bucket versions.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.StakingService.GetBucketList \
  --header 'Content-Type: application/json' \
  --data '{
  "limit": 20,
  "offset": 0,
  "sort": "timestamp:desc",
  "interval": "7D",
  "version": "native"
}'
```

```graphql
query {
  GetBucketList(limit: 20, offset: 0, sort: "timestamp:desc", interval: "7D", version: "native") {
    buckets {
      bucket_id
      owner_address
      candidate
      staked_amount
      duration
      auto_stake
      timestamp
    }
    count
    group_count
  }
}
```

> Example response:

```json
{
  "data": {
    "GetBucketList": {
      "buckets": [
        {
          "bucket_id": 1001,
          "owner_address": "io1...",
          "candidate": "iotexlab",
          "staked_amount": "10000000000000000000000",
          "duration": "91 days",
          "auto_stake": true,
          "timestamp": "2024-01-15T10:00:00Z"
        }
      ],
      "count": 5000,
      "group_count": 5000
    }
  }
}
```

### HTTP Request

`POST /api.StakingService.GetBucketList`

<a name="api-GetBucketListRequest"></a>

### GetBucketListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int64](#int64) |  | number of records per page |
| offset | [int64](#int64) |  | starting offset |
| sort | [string](#string) |  | sort field and direction (e.g. "timestamp:desc") |
| interval | [string](#string) |  | time interval filter: "1D", "7D", "30D", "1Y", "ALL" |
| version | [string](#string) |  | bucket version: "native", "nft_v1", "nft_v2", "nft_v3" |

<a name="api-GetBucketListResponse"></a>

### GetBucketListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buckets | [BucketInfoEx](#api-BucketInfoEx) | repeated | list of buckets |
| count | [int64](#int64) |  | total number of matching buckets |
| group_count | [int64](#int64) |  | total group count |

<a name="api-BucketInfoEx"></a>

### BucketInfoEx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bucket_id | [int64](#int64) |  | bucket ID |
| action_hash | [string](#string) |  | creation action hash |
| timestamp | [string](#string) |  | creation timestamp |
| create_time | [string](#string) |  | create time |
| stake_start_time | [string](#string) |  | stake start time |
| unstake_start_time | [string](#string) |  | unstake start time |
| amount | [string](#string) |  | action amount |
| staked_amount | [string](#string) |  | staked amount |
| act_type | [string](#string) |  | action type |
| sender | [string](#string) |  | sender address |
| owner_address | [string](#string) |  | bucket owner address |
| candidate | [string](#string) |  | candidate name |
| auto_stake | [bool](#bool) |  | whether auto-stake is enabled |
| duration | [string](#string) |  | stake duration |
| gas_price | [string](#string) |  | gas price |
| gas_limit | [string](#string) |  | gas limit |
| recipient | [string](#string) |  | recipient address |
| delegate_name | [string](#string) |  | delegate name |

## GetBucketsByBucketId

GetBucketsByBucketId returns the history of actions for a specific staking bucket ID.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.StakingService.GetBucketsByBucketId \
  --header 'Content-Type: application/json' \
  --data '{
  "bucket_id": 1001,
  "limit": 20,
  "offset": 0,
  "version": "native"
}'
```

```graphql
query {
  GetBucketsByBucketId(bucket_id: 1001, limit: 20, offset: 0, version: "native") {
    buckets {
      bucket_id
      act_type
      staked_amount
      timestamp
    }
    count
  }
}
```

> Example response:

```json
{
  "data": {
    "GetBucketsByBucketId": {
      "buckets": [
        {
          "bucket_id": 1001,
          "act_type": "stakeCreate",
          "staked_amount": "10000000000000000000000",
          "timestamp": "2024-01-15T10:00:00Z"
        }
      ],
      "count": 5
    }
  }
}
```

### HTTP Request

`POST /api.StakingService.GetBucketsByBucketId`

<a name="api-GetBucketsByBucketIdRequest"></a>

### GetBucketsByBucketIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bucket_id | [int64](#int64) |  | staking bucket ID |
| limit | [int64](#int64) |  | number of records per page |
| offset | [int64](#int64) |  | starting offset |
| version | [string](#string) |  | bucket version: "native", "nft_v1", "nft_v2", "nft_v3" |

<a name="api-GetBucketsByBucketIdResponse"></a>

### GetBucketsByBucketIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buckets | [BucketInfoEx](#api-BucketInfoEx) | repeated | list of bucket history records |
| count | [int64](#int64) |  | total number of records |

## GetBucketByBucketId

GetBucketByBucketId returns the details of a single staking bucket.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.StakingService.GetBucketByBucketId \
  --header 'Content-Type: application/json' \
  --data '{
  "bucket_id": 1001,
  "version": "native"
}'
```

```graphql
query {
  GetBucketByBucketId(bucket_id: 1001, version: "native") {
    exist
    bucket {
      bucket_id
      owner_address
      candidate
      staked_amount
      duration
      auto_stake
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "GetBucketByBucketId": {
      "exist": true,
      "bucket": {
        "bucket_id": 1001,
        "owner_address": "io1...",
        "candidate": "iotexlab",
        "staked_amount": "10000000000000000000000",
        "duration": "91 days",
        "auto_stake": true
      }
    }
  }
}
```

### HTTP Request

`POST /api.StakingService.GetBucketByBucketId`

<a name="api-GetBucketByBucketIdRequest"></a>

### GetBucketByBucketIdRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bucket_id | [int64](#int64) |  | staking bucket ID |
| version | [string](#string) |  | bucket version: "native", "nft_v1", "nft_v2", "nft_v3" |

<a name="api-GetBucketByBucketIdResponse"></a>

### GetBucketByBucketIdResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether the bucket exists |
| bucket | [BucketInfoEx](#api-BucketInfoEx) |  | bucket details |

## GetNativeBuckets

GetNativeBuckets returns a paginated list of native staking buckets.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.StakingService.GetNativeBuckets \
  --header 'Content-Type: application/json' \
  --data '{
  "limit": 20,
  "offset": 0
}'
```

```graphql
query {
  GetNativeBuckets(limit: 20, offset: 0) {
    buckets {
      bucket_id
      owner_address
      candidate
      staked_amount
      duration
      auto_stake
    }
    count
  }
}
```

> Example response:

```json
{
  "data": {
    "GetNativeBuckets": {
      "buckets": [
        {
          "bucket_id": 1001,
          "owner_address": "io1...",
          "candidate": "iotexlab",
          "staked_amount": "10000000000000000000000",
          "duration": "91 days",
          "auto_stake": true
        }
      ],
      "count": 50000
    }
  }
}
```

### HTTP Request

`POST /api.StakingService.GetNativeBuckets`

<a name="api-GetNativeBucketsRequest"></a>

### GetNativeBucketsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int64](#int64) |  | number of records per page |
| offset | [int64](#int64) |  | starting offset |

<a name="api-GetNativeBucketsResponse"></a>

### GetNativeBucketsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| buckets | [BucketInfoEx](#api-BucketInfoEx) | repeated | list of native buckets |
| count | [int64](#int64) |  | total number of native buckets |

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


## GetXRC20TransfersByContract

GetXRC20TransfersByContract returns ERC20 transfers for a given contract, with optional sender/recipient filter.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.XRC20Service.GetXRC20TransfersByContract \
  --header 'Content-Type: application/json' \
  --data '{
  "contract_address": "io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw",
  "pagination": { "skip": 0, "first": 20 }
}'
```

> Example response:

```json
{
  "exist": true,
  "count": 1000,
  "transfers": [
    {
      "id": 1,
      "block_height": 19000000,
      "action_hash": "abc123...",
      "contract_address": "io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw",
      "amount": "1000000000000000000",
      "sender": "io1...",
      "recipient": "io1...",
      "timestamp": "2024-01-15T10:00:00Z"
    }
  ]
}
```

### HTTP Request

`POST /api.XRC20Service.GetXRC20TransfersByContract`

<a name="api-GetXRC20TransfersByContractRequest"></a>

### GetXRC20TransfersByContractRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  | ERC20 contract address |
| address | [string](#string) |  | optional: filter by sender or recipient address |
| pagination | [pagination.Pagination](#pagination-Pagination) |  | pagination info |

<a name="api-GetXRC20TransfersByContractResponse"></a>

### GetXRC20TransfersByContractResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether transfers exist |
| count | [int64](#int64) |  | total number of transfers |
| transfers | [XRC20TransferInfo](#api-XRC20TransferInfo) | repeated | list of transfers |

<a name="api-XRC20TransferInfo"></a>

### XRC20TransferInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | record id |
| block_height | [uint64](#uint64) |  | block height |
| action_hash | [string](#string) |  | action hash |
| contract_address | [string](#string) |  | token contract address |
| amount | [string](#string) |  | transfer amount |
| sender | [string](#string) |  | sender address |
| recipient | [string](#string) |  | recipient address |
| timestamp | [string](#string) |  | timestamp |

## GetXRC20HoldersByContract

GetXRC20HoldersByContract returns all holders of a given ERC20 token with their balances.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.XRC20Service.GetXRC20HoldersByContract \
  --header 'Content-Type: application/json' \
  --data '{
  "contract_address": "io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw",
  "pagination": { "skip": 0, "first": 20 }
}'
```

> Example response:

```json
{
  "count": 5000,
  "holders": [
    {
      "address": "io1...",
      "balance": "1000000000000000000"
    }
  ]
}
```

### HTTP Request

`POST /api.XRC20Service.GetXRC20HoldersByContract`

<a name="api-GetXRC20HoldersByContractRequest"></a>

### GetXRC20HoldersByContractRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  | ERC20 contract address |
| pagination | [pagination.Pagination](#pagination-Pagination) |  | pagination info |

<a name="api-GetXRC20HoldersByContractResponse"></a>

### GetXRC20HoldersByContractResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [int64](#int64) |  | total number of holders |
| holders | [XRC20HolderInfo](#api-XRC20HolderInfo) | repeated | list of holders with balances |

<a name="api-XRC20HolderInfo"></a>

### XRC20HolderInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | holder address |
| balance | [string](#string) |  | token balance |

## GetXRC20TokenBalance

GetXRC20TokenBalance returns the ERC20 token balance for a specific address and contract.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.XRC20Service.GetXRC20TokenBalance \
  --header 'Content-Type: application/json' \
  --data '{
  "contract_address": "io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw",
  "address": "io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4"
}'
```

> Example response:

```json
{
  "balance": "1000000000000000000"
}
```

### HTTP Request

`POST /api.XRC20Service.GetXRC20TokenBalance`

<a name="api-GetXRC20TokenBalanceRequest"></a>

### GetXRC20TokenBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  | ERC20 contract address |
| address | [string](#string) |  | holder address |

<a name="api-GetXRC20TokenBalanceResponse"></a>

### GetXRC20TokenBalanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balance | [string](#string) |  | token balance (raw units) |

## GetXRC20Stats

GetXRC20Stats returns per-token holders / transfer / daily_transfer counts, ordered by holders DESC. Holders come from the pre-aggregated erc20_holder_agg view (balance > 0). Transfer counts are aggregated live from erc20_transfers. Page size is hard-capped at 50 — first=0 or any value above 50 is treated as 50.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.XRC20Service.GetXRC20Stats \
  --header 'Content-Type: application/json' \
  --data '{
  "pagination": { "first": 10, "skip": 0 }
}'
```

> Example response:

```json
{
  "count": "2068",
  "items": [
    { "address": "io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw", "holders": "25379", "transfer": "901907", "daily_transfer": "0" },
    { "address": "io109mf3ua2tfkm2lkzq432hu03ys7vuv4d5m40gk", "holders": "20033", "transfer": "39277",  "daily_transfer": "0" },
    { "address": "io1aq4hq4z8r5l4ejp9c5p4pk3mefj000jwuyrlk2", "holders": "13310", "transfer": "27341",  "daily_transfer": "0" }
  ]
}
```

### HTTP Request

`POST /api.XRC20Service.GetXRC20Stats`

<a name="api-GetXRC20StatsRequest"></a>

### GetXRC20StatsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [Pagination](#pagination-Pagination) |  | first capped at 50; first=0 → 50 |

<a name="api-GetXRC20StatsResponse"></a>

### GetXRC20StatsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint64](#uint64) |  | total number of XRC20 contracts with at least one holder |
| items | [XRC20StatsItem](#api-XRC20StatsItem) | repeated | top-K page of stats, sorted by holders DESC |

<a name="api-XRC20StatsItem"></a>

### XRC20StatsItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | contract address (io1...) |
| holders | [uint64](#uint64) |  | unique holder count (balance > 0) |
| transfer | [uint64](#uint64) |  | all-time transfer count |
| daily_transfer | [uint64](#uint64) |  | transfer count in the prior UTC day |

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


## GetNFTTransferList

GetNFTTransferList returns NFT transfers (XRC721 + XRC1155) with optional contract and address filters.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.XRC721Service.GetNFTTransferList \
  --header 'Content-Type: application/json' \
  --data '{
  "pagination": { "skip": 0, "first": 20 },
  "contract_address": "io1nft..."
}'
```

> Example response:

```json
{
  "exist": true,
  "count": 10000,
  "transfers": [
    {
      "id": 1,
      "type": "xrc721",
      "block_height": 19000000,
      "action_hash": "abc123...",
      "contract_address": "io1nft...",
      "token_id": "42",
      "value": "1",
      "sender": "io1...",
      "recipient": "io1...",
      "timestamp": "2024-01-15T10:00:00Z"
    }
  ]
}
```

### HTTP Request

`POST /api.XRC721Service.GetNFTTransferList`

<a name="api-GetNFTTransferListRequest"></a>

### GetNFTTransferListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | [pagination.Pagination](#pagination-Pagination) |  | pagination info |
| contract_address | [string](#string) |  | optional: filter by contract address |
| address | [string](#string) |  | optional: filter by sender, recipient, or token ID |

<a name="api-GetNFTTransferListResponse"></a>

### GetNFTTransferListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exist | [bool](#bool) |  | whether transfers exist |
| count | [int64](#int64) |  | total number of transfers |
| transfers | [NFTTransferInfo](#api-NFTTransferInfo) | repeated | list of NFT transfers |

<a name="api-NFTTransferInfo"></a>

### NFTTransferInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [uint64](#uint64) |  | record id |
| type | [string](#string) |  | token standard: "xrc721" or "xrc1155" |
| block_height | [uint64](#uint64) |  | block height |
| action_hash | [string](#string) |  | action hash |
| contract_address | [string](#string) |  | NFT contract address |
| token_id | [string](#string) |  | token ID |
| value | [string](#string) |  | transfer value (amount for xrc1155) |
| sender | [string](#string) |  | sender address |
| recipient | [string](#string) |  | recipient address |
| timestamp | [string](#string) |  | timestamp |

## GetNFTHoldersByContract

GetNFTHoldersByContract returns NFT holders for a given contract with their balances.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.XRC721Service.GetNFTHoldersByContract \
  --header 'Content-Type: application/json' \
  --data '{
  "contract_address": "io1nft...",
  "pagination": { "skip": 0, "first": 20 }
}'
```

> Example response:

```json
{
  "count": 500,
  "holders": [
    {
      "address": "io1...",
      "balance": "3"
    }
  ]
}
```

### HTTP Request

`POST /api.XRC721Service.GetNFTHoldersByContract`

<a name="api-GetNFTHoldersByContractRequest"></a>

### GetNFTHoldersByContractRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | [string](#string) |  | NFT contract address |
| pagination | [pagination.Pagination](#pagination-Pagination) |  | pagination info |

<a name="api-GetNFTHoldersByContractResponse"></a>

### GetNFTHoldersByContractResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [int64](#int64) |  | total number of holders |
| holders | [NFTHolderInfo](#api-NFTHolderInfo) | repeated | list of holders |

<a name="api-NFTHolderInfo"></a>

### NFTHolderInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | holder address |
| balance | [string](#string) |  | number of NFTs held |

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

## HermesBucket

HermesBucket returns bucket-level reward distributions for Hermes delegates, showing per-bucket breakdown.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.HermesService.HermesBucket \
  --header 'Content-Type: application/json' \
  --data '{
  "startEpoch": 20000,
  "epochCount": 24,
  "rewardAddress": ["io12mgttmfa2ffn9uqvn0yn37f4nz43d248l2ga85"]
}'
```

```graphql
query {
  HermesBucket(
    startEpoch: 20000
    epochCount: 24
    rewardAddress: ["io12mgttmfa2ffn9uqvn0yn37f4nz43d248l2ga85"]
  ) {
    hermesBucketDistribution {
      delegateName
      stakingIotexAddress
      voterCount
      waiveServiceFee
      refund
      bucketRewardDistribution {
        voterEthAddress
        voterIotexAddress
        bucketID
        amount
      }
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "HermesBucket": {
      "hermesBucketDistribution": [
        {
          "delegateName": "iotexlab",
          "stakingIotexAddress": "io1...",
          "voterCount": 150,
          "waiveServiceFee": true,
          "refund": "0",
          "bucketRewardDistribution": [
            {
              "voterEthAddress": "0x...",
              "voterIotexAddress": "io1...",
              "bucketID": 1001,
              "amount": "1000000000000000000"
            }
          ]
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.HermesService.HermesBucket`

<a name="api-HermesBucketResponse"></a>

### HermesBucketResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| hermesBucketDistribution | [HermesBucketDistribution](#api-HermesBucketDistribution) | repeated | list of bucket reward distributions by delegate |

<a name="api-HermesBucketDistribution"></a>

### HermesBucketDistribution



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| delegateName | [string](#string) |  | delegate name |
| bucketRewardDistribution | [BucketRewardDistribution](#api-BucketRewardDistribution) | repeated | list of per-bucket reward distributions |
| stakingIotexAddress | [string](#string) |  | delegate IoTeX staking address |
| voterCount | [uint64](#uint64) |  | number of voters |
| waiveServiceFee | [bool](#bool) |  | whether the delegate waives the service fee |
| refund | [string](#string) |  | refund amount |

<a name="api-BucketRewardDistribution"></a>

### BucketRewardDistribution



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| voterEthAddress | [string](#string) |  | voter's ERC20 address |
| voterIotexAddress | [string](#string) |  | voter's IoTeX address |
| bucketID | [uint64](#uint64) |  | staking bucket ID |
| amount | [string](#string) |  | reward amount |

## HermesDropRecords

HermesDropRecords inserts Hermes drop records for a delegate and epoch.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.HermesService.HermesDropRecords \
  --header 'Content-Type: application/json' \
  --data '{
  "epochNumber": 28884,
  "delegateName": "iotexlab",
  "voterAddress": "io1...",
  "actHash": "abc123...",
  "bucketID": 1001,
  "amount": "1000000000000000000"
}'
```

> Example response:

```json
{
  "success": true
}
```

### HTTP Request

`POST /api.HermesService.HermesDropRecords`

<a name="api-HermesDropRecordsRequest"></a>

### HermesDropRecordsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| epochNumber | [uint64](#uint64) |  | end epoch number |
| delegateName | [string](#string) |  | delegate name |
| voterAddress | [string](#string) |  | voter address |
| actHash | [string](#string) |  | action hash |
| bucketID | [uint64](#uint64) |  | bucket ID |
| amount | [string](#string) |  | reward amount |

<a name="api-HermesDropRecordsResponse"></a>

### HermesDropRecordsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | whether drop records were successfully inserted |

# Approval Service API

## GetXRC20Approvals

GetXRC20Approvals returns the ERC20 token approvals made by a given owner address.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ApprovalService.GetXRC20Approvals \
  --header 'Content-Type: application/json' \
  --data '{
  "owner_address": "io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4"
}'
```

```graphql
query {
  GetXRC20Approvals(owner_address: "io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4") {
    approvals {
      action_hash
      contract_address
      owner
      spender
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
    "GetXRC20Approvals": {
      "approvals": [
        {
          "action_hash": "abc123...",
          "contract_address": "io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw",
          "owner": "io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4",
          "spender": "io1...",
          "amount": "115792089237316195423570985008687907853269984665640564039457584007913129639935",
          "timestamp": "2024-01-15T10:00:00Z"
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.ApprovalService.GetXRC20Approvals`

<a name="api-GetXRC20ApprovalsRequest"></a>

### GetXRC20ApprovalsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner_address | [string](#string) |  | owner address |

<a name="api-GetXRC20ApprovalsResponse"></a>

### GetXRC20ApprovalsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| approvals | [XRC20ApprovalInfo](#api-XRC20ApprovalInfo) | repeated | list of ERC20 approvals |

<a name="api-XRC20ApprovalInfo"></a>

### XRC20ApprovalInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action_hash | [string](#string) |  | approval action hash |
| contract_address | [string](#string) |  | ERC20 contract address |
| owner | [string](#string) |  | token owner address |
| spender | [string](#string) |  | approved spender address |
| amount | [string](#string) |  | approved amount |
| timestamp | [string](#string) |  | approval timestamp |

## GetXRC721Approvals

GetXRC721Approvals returns the NFT approvals made by a given owner address.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ApprovalService.GetXRC721Approvals \
  --header 'Content-Type: application/json' \
  --data '{
  "owner_address": "io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4"
}'
```

```graphql
query {
  GetXRC721Approvals(owner_address: "io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4") {
    approvals {
      action_hash
      contract_address
      owner
      approved
      token_id
      timestamp
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "GetXRC721Approvals": {
      "approvals": [
        {
          "action_hash": "abc123...",
          "contract_address": "io1nft...",
          "owner": "io1x58dug5237g40hrtme7qx4nva9x98ehk4wchz4",
          "approved": "io1...",
          "token_id": "42",
          "timestamp": "2024-01-15T10:00:00Z"
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.ApprovalService.GetXRC721Approvals`

<a name="api-GetXRC721ApprovalsRequest"></a>

### GetXRC721ApprovalsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| owner_address | [string](#string) |  | owner address |

<a name="api-GetXRC721ApprovalsResponse"></a>

### GetXRC721ApprovalsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| approvals | [XRC721ApprovalInfo](#api-XRC721ApprovalInfo) | repeated | list of NFT approvals |

<a name="api-XRC721ApprovalInfo"></a>

### XRC721ApprovalInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| action_hash | [string](#string) |  | approval action hash |
| contract_address | [string](#string) |  | NFT contract address |
| owner | [string](#string) |  | token owner address |
| approved | [string](#string) |  | approved operator address |
| token_id | [string](#string) |  | token ID |
| timestamp | [string](#string) |  | approval timestamp |

# Actions Service API

## GetEvmTransferDetailListByAddress

GetEvmTransferDetailListByAddress returns a paginated list of EVM transfer details for an address.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ActionsService.GetEvmTransferDetailListByAddress \
  --header 'Content-Type: application/json' \
  --data '{
  "address": "io14u5d66rt465ykm7t2847qllj0reml27q30kr75",
  "offset": 0,
  "size": 10
}'
```

```graphql
query {
  GetEvmTransferDetailListByAddress(
    address: "io14u5d66rt465ykm7t2847qllj0reml27q30kr75"
    offset: 0
    size: 10
  ) {
    count
    results {
      actHash
      blkHeight
      sender
      recipient
      blkHash
      amount
      timeStamp
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "GetEvmTransferDetailListByAddress": {
      "count": 250,
      "results": [
        {
          "actHash": "abc123...",
          "blkHeight": 19000000,
          "sender": "io1...",
          "recipient": "io1...",
          "blkHash": "def456...",
          "amount": "1000000000000000000",
          "timeStamp": 1714000000
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.ActionsService.GetEvmTransferDetailListByAddress`

<a name="api-ActionsRequest"></a>

### ActionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [string](#string) |  | wallet address |
| height | [uint64](#uint64) |  | optional: filter by block height |
| offset | [uint64](#uint64) |  | starting offset |
| size | [uint64](#uint64) |  | number of records to return |
| sort | [string](#string) |  | sort direction (e.g. "desc") |

<a name="api-EvmTransferDetailListByAddressResponse"></a>

### EvmTransferDetailListByAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint64](#uint64) |  | total number of EVM transfers |
| results | [EvmTransferDetailResult](#api-EvmTransferDetailResult) | repeated | list of EVM transfers |

<a name="api-EvmTransferDetailResult"></a>

### EvmTransferDetailResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actHash | [string](#string) |  | action hash |
| blkHeight | [uint64](#uint64) |  | block height |
| sender | [string](#string) |  | sender address |
| recipient | [string](#string) |  | recipient address |
| blkHash | [string](#string) |  | block hash |
| amount | [string](#string) |  | transfer amount |
| timeStamp | [uint64](#uint64) |  | unix timestamp |

## GetAllActionsByAddress

GetAllActionsByAddress returns a combined paginated list of all action types (native, XRC20, XRC721, EVM transfer) for an address.

```shell
curl --request POST \
  --url https://analyser-api.iotex.io/api.ActionsService.GetAllActionsByAddress \
  --header 'Content-Type: application/json' \
  --data '{
  "address": "io14u5d66rt465ykm7t2847qllj0reml27q30kr75",
  "offset": 0,
  "size": 10,
  "sort": "desc"
}'
```

```graphql
query {
  GetAllActionsByAddress(
    address: "io14u5d66rt465ykm7t2847qllj0reml27q30kr75"
    offset: 0
    size: 10
    sort: "desc"
  ) {
    count
    results {
      actHash
      blkHeight
      sender
      recipient
      actType
      amount
      timeStamp
      recordType
    }
  }
}
```

> Example response:

```json
{
  "data": {
    "GetAllActionsByAddress": {
      "count": 5000,
      "results": [
        {
          "actHash": "abc123...",
          "blkHeight": 19000000,
          "sender": "io1...",
          "recipient": "io1...",
          "actType": "transfer",
          "amount": "1000000000000000000",
          "timeStamp": 1714000000,
          "recordType": 0
        }
      ]
    }
  }
}
```

### HTTP Request

`POST /api.ActionsService.GetAllActionsByAddress`

<a name="api-AllActionsByAddressResponse"></a>

### AllActionsByAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [uint64](#uint64) |  | total number of actions |
| results | [AllActionsByAddressResult](#api-AllActionsByAddressResult) | repeated | list of all actions |

<a name="api-AllActionsByAddressResult"></a>

### AllActionsByAddressResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actHash | [string](#string) |  | action hash |
| blkHeight | [uint64](#uint64) |  | block height |
| sender | [string](#string) |  | sender address |
| recipient | [string](#string) |  | recipient address |
| actType | [string](#string) |  | action type |
| amount | [string](#string) |  | amount |
| timeStamp | [uint64](#uint64) |  | unix timestamp |
| recordType | [AllActionsByAddressResult.RecordType](#api-AllActionsByAddressResult-RecordType) |  | record type: NATIVE(0), XRC20(1), XRC721(2), EVMTRANSFER(3) |
