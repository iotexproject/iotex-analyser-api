# 首页统计 / 图表 API 使用说明

`iotex-analyser-api >= v1.1.1` 新增 5 个端点，用于驱动 iotexscan 首页统计面板、`/charts/*` 图表页、Token 列表页。

本文档对接原始需求里的 5 个接口，但 URL 命名 / 响应字段都换成了项目自身的 gRPC + grpc-gateway 风格，需对齐才能调通。

## 1. 基础信息

| 项目 | 值 |
|---|---|
| Mainnet base URL | `https://analyser-api.iotex.io` |
| Testnet base URL | `https://analyser-api.testnet.iotex.io` |
| HTTP 方法 | **全部用 `POST`**（不是 GET — 由 grpc-gateway 生成）|
| Content-Type | `application/json` |
| 鉴权 | JWT，HTTP header `Authorization: Bearer <token>` |
| 时区 | 所有日期参数/返回值都按 **UTC** 解释 |

> 注：HTTP 端点都是 `POST /api.{Service}.{Method}`，**查询参数走 JSON body 而不是 query string**。原始需求里 `?start=...&end=...` 的写法不适用。

## 2. 5 个端点的命名对照

| # | 用途 | 原始需求 URL | 实际 URL |
|---|---|---|---|
| 1 | 首页统计 | `GET /api/v2/chain/stats` | `POST /api.ChainService.GetChainStats` |
| 2 | TPS 时序 | `GET /api/v2/chain/tps-history` | `POST /api.ChainService.GetTpsHistory` |
| 3 | Gas 费时序 | `GET /api/v2/chain/gas-history` | `POST /api.ChainService.GetGasHistory` |
| 4 | 供应量时序 | `GET /api/v2/chain/supply-history` | `POST /api.ChainService.GetSupplyHistory` |
| 5 | Token 统计 | `GET /api/v2/tokens/xrc20-stats` | `POST /api.XRC20Service.GetXRC20Stats` |

## 3. 通用约定

### 3.1 JSON 字段命名

Proto 内部用 snake_case，grpc-gateway 默认转 lowerCamelCase 输出到 JSON：
- `actions_num` → `actionsNum`
- `total_supply` → `totalSupply`
- `max_gas_price` → `maxGasPrice`
- `daily_transfer` → `dailyTransfer`

### 3.2 大整数 / `uint64` 都是字符串

proto3 标准 JSON 编码把 `uint64` 编码成 string，避免 JS 浮点精度丢失。所以 `actionsNum: "213194125"`，**前端要按字符串解析**。

`maxGasPrice` / `totalGasFee` 这些 rau 单位的大值也是 string。

### 3.3 日期参数

`start` / `end` 形如 `"YYYY-MM-DD"`，UTC，**两端都包含**（inclusive）。

错误格式 / 空串 / `start > end` 会返回 `InvalidArgument` 而不是 500：

```json
{
  "code": 3,
  "message": "invalid start date \"bad\": expected YYYY-MM-DD"
}
```

### 3.4 响应都是对象不是数组

原始需求里有些接口期望响应直接是数组（`[{...}, {...}]`）；实际响应都是对象，list 在 `data` 字段下：

```json
{ "data": [ {...}, {...} ] }
```

这是 grpc-gateway 的约束。

---

## 4. 端点详解

### 4.1 GetChainStats — 首页 3 个数字

**用途**：首页统计面板里的 Total Actions / Total Supply / Circulating Supply。

```shell
curl -X POST https://analyser-api.iotex.io/api.ChainService.GetChainStats \
  -H 'Authorization: Bearer <token>' \
  -H 'Content-Type: application/json' \
  -d '{}'
```

**请求体**：空对象 `{}`。

**响应（mainnet 实测）**：

```json
{
  "actionsNum": "213194125",
  "totalSupply": "9441368502",
  "circulatingSupply": "9441368498"
}
```

| 字段 | 类型 | 单位 | 说明 |
|---|---|---|---|
| `actionsNum` | uint64（JSON string） | — | 链上累计 action 总数 |
| `totalSupply` | string | **IOTX**（已 /1e18，整数，无小数点） | 当前总供应量 |
| `circulatingSupply` | string | IOTX（整数） | 当前流通供应量 |

**数据来源 & 注意**：

- 3 个值都读自 `iotexscanv3_kv` 表，由 Windmill 的 `sync-iotex-statistics.deno.ts` 定时刷新，**约 65 秒延迟**。
- ⚠️ Testnet 上这张表里**没有数据**（Windmill 没在 testnet 配置同步），调 testnet 会返回 `{}`（所有字段都是 zero value）。如果 testnet 也要展示，请联系 DevOps 把 Windmill job 加上 testnet 实例。
- 原始需求里 `total_supply` 是 string 形如 `"9876543210"`（整数 IOTX，无小数）—— 我们一致。

---

### 4.2 GetTpsHistory — 每日 TPS 时序

**用途**：首页 TPS 折线图（Avg / Peak 切换）。

```shell
curl -X POST https://analyser-api.iotex.io/api.ChainService.GetTpsHistory \
  -H 'Authorization: Bearer <token>' \
  -H 'Content-Type: application/json' \
  -d '{"start":"2026-05-15","end":"2026-05-18"}'
```

**请求体**：

```json
{ "start": "2026-05-15", "end": "2026-05-18" }
```

**响应（mainnet 实测）**：

```json
{
  "data": [
    { "date": "2026-05-15", "avgTps": 1.19, "maxTps": 10.0 },
    { "date": "2026-05-16", "avgTps": 1.23, "maxTps": 14.4 },
    { "date": "2026-05-17", "avgTps": 1.29, "maxTps": 14.0 },
    { "date": "2026-05-18", "avgTps": 1.17, "maxTps": 18.0 }
  ]
}
```

| 字段 | 类型 | 说明 |
|---|---|---|
| `date` | string `YYYY-MM-DD` | UTC 日 |
| `avgTps` | double | `AVG(num_actions) / 2.5` 四舍五入到 2 位小数 |
| `maxTps` | double | `MAX(num_actions) / 2.5` 四舍五入到 2 位小数 |

**数据来源 & 注意**：

- 直接从 `block` 表实时按日 `GROUP BY` 聚合。
- 出块间隔常量 **`2.5` 秒**（Wake hard fork 后值）。**查询 Wake (block 36893881, 2025-08) 之前的日期会把 TPS 高估 2 倍**（实际间隔是 5 s）。首页只展示近期数据时无影响。

---

### 4.3 GetGasHistory — 每日 Gas 费时序

**用途**：`/charts/gasprice` 与 `/charts/gasused`。

```shell
curl -X POST https://analyser-api.iotex.io/api.ChainService.GetGasHistory \
  -H 'Authorization: Bearer <token>' \
  -H 'Content-Type: application/json' \
  -d '{"start":"2026-05-17","end":"2026-05-18"}'
```

**响应（mainnet 实测）**：

```json
{
  "data": [
    {
      "date": "2026-05-17",
      "maxGasPrice": "17897676565538",
      "minGasPrice": "1000000000000",
      "avgGasPrice": "1021883138757",
      "totalGasFee": "7110675516295301992332"
    },
    {
      "date": "2026-05-18",
      "maxGasPrice": "2029734718659213",
      "minGasPrice": "1000000000000",
      "avgGasPrice": "1160216242181",
      "totalGasFee": "7597570410715163176791"
    }
  ]
}
```

| 字段 | 类型 | 单位 | 说明 |
|---|---|---|---|
| `date` | string `YYYY-MM-DD` | — | UTC 日 |
| `maxGasPrice` | string | rau | 当日 `MAX(gas_price)`，排除 `gas_price = 0` |
| `minGasPrice` | string | rau | 当日 `MIN(gas_price)`，排除 `gas_price = 0` |
| `avgGasPrice` | string | rau | 当日 `ROUND(AVG(gas_price))`，简单算术平均（非加权） |
| `totalGasFee` | string | rau | 当日 `SUM(gas_price × gas_consumed)` |

**数据来源 & 注意**：

- 实时聚合 `block_action × block`，使用 `block_height` 范围 + 分区裁剪。
- ⚠️ **大范围慢**：7 天约 **3 s**；超过 30 天会到 10+ s；不建议一次拉 1 年的数据。如果有这种需求要做 cron 预聚合，请先告知后端。
- 单位是 rau（`1 IOTX = 10¹⁸ rau`）。前端展示时除 1e18 → IOTX，再除 1e9 → Gwei，再除 1e3 → Twei。
- 排除 `gas_price = 0`（系统 action）。

---

### 4.4 GetSupplyHistory — 每日供应量时序

**用途**：`/charts/supply` 与 `/charts/circulating-supply`，可以一次拿到 supply / burn / issue 4 个字段。

```shell
curl -X POST https://analyser-api.iotex.io/api.ChainService.GetSupplyHistory \
  -H 'Authorization: Bearer <token>' \
  -H 'Content-Type: application/json' \
  -d '{"start":"2023-11-01","end":"2023-11-07"}'
```

**响应（mainnet 实测，2023-11 有真实 burn 活动）**：

```json
{
  "data": [
    {
      "date": "2023-11-01",
      "totalSupply": "9443041459.47",
      "circulatingSupply": "9443041454.63",
      "burn":  "28125.00",
      "issue": "0.00"
    },
    {
      "date": "2023-11-02",
      "totalSupply": "9443003959.47",
      "circulatingSupply": "9443003954.63",
      "burn":  "37500.00",
      "issue": "0.00"
    },
    ...
  ]
}
```

| 字段 | 类型 | 单位 | 说明 |
|---|---|---|---|
| `date` | string `YYYY-MM-DD` | — | UTC 日 |
| `totalSupply` | string | **IOTX**（小数 2 位） | 当日**收盘时**的总供应量 |
| `circulatingSupply` | string | IOTX（小数 2 位） | 当日收盘时的流通供应量 |
| `burn` | string | IOTX（小数 2 位） | 当日新增销毁量 = `prev.totalSupply − today.totalSupply` |
| `issue` | string | IOTX（小数 2 位） | 当日新增释放量 = `today.circ − prev.circ + burn` |

**数据来源 & 注意**：

- 数据源：`block_supply × block`。每日取该日最后一个 block 的 `total_supply` / `total_circulating_supply` 当作"收盘值"。
- `burn` / `issue` 不是数据库里独立存的字段，而是从 supply 时序**逆推**出来的：
  - 零地址只收不发 → `Δtotal_supply = −burn`
  - 锁仓地址只发不收 → `Δcirculating ≈ −burn + issue`
- 单位都是 **IOTX**（rau 已经除以 1e18），保留 2 位小数。和原始需求文档一致。
- ⚠️ **首日 `burn` / `issue` 为空字符串**（没有前一日基线，无法做差）。前端展示需容错。
- ⚠️ **日间断点**：如果中间某一天链没出块（halt / 缺数据），那一天会从 `data` 里直接缺失，**且断点后第一天的 `burn` / `issue` 也会为空**（避免误归属多日变化到单一天）。前端应该按 `date` 而不是数组下标对齐。

---

### 4.5 GetXRC20Stats — Token 列表统计

**用途**：Token 列表页（按 holders 降序）。

```shell
curl -X POST https://analyser-api.iotex.io/api.XRC20Service.GetXRC20Stats \
  -H 'Authorization: Bearer <token>' \
  -H 'Content-Type: application/json' \
  -d '{"pagination":{"first":10,"skip":0}}'
```

**请求体**：

```json
{ "pagination": { "first": 10, "skip": 0 } }
```

| 字段 | 默认 / 上限 | 说明 |
|---|---|---|
| `pagination.first` | 默认 `50`、**硬上限 `50`** | 每页条数；传 `0` 或不传 → 50；超过 50 → 自动 cap 到 50 |
| `pagination.skip` | 默认 `0` | 偏移 |

**响应（mainnet 实测，top 3）**：

```json
{
  "count": "2068",
  "items": [
    {
      "address": "io1hp6y4eqr90j7tmul4w2wa8pm7wx462hq0mg4tw",
      "holders": "25379",
      "transfer": "901907",
      "dailyTransfer": "0"
    },
    {
      "address": "io109mf3ua2tfkm2lkzq432hu03ys7vuv4d5m40gk",
      "holders": "20033",
      "transfer": "39277",
      "dailyTransfer": "0"
    },
    {
      "address": "io1aq4hq4z8r5l4ejp9c5p4pk3mefj000jwuyrlk2",
      "holders": "13310",
      "transfer": "27341",
      "dailyTransfer": "0"
    }
  ]
}
```

| 字段 | 类型 | 说明 |
|---|---|---|
| `count` | uint64（string） | 全网 XRC20 合约总数（balance > 0 的去重 contract 数）|
| `items[].address` | string | 合约 io 地址 |
| `items[].holders` | uint64（string） | 当前持有人数（`erc20_holder_agg` 中 `balance > 0` 的行数）|
| `items[].transfer` | uint64（string） | 历史累计 transfer 数 |
| `items[].dailyTransfer` | uint64（string） | **昨日**（UTC 时间）的 transfer 数 |

**数据来源 & 注意**：

- `holders` 走预聚合表 `erc20_holder_agg`（由 Windmill `f/analyzer/xrc20_holder_agg.deno.ts` 维护，**不要去查 `erc20_holders` 表**——那是 append-only 事件日志，一个合约几百万行）。
- `transfer` / `dailyTransfer` 走 `erc20_transfers` 表的 `COUNT(*)`，按 contract 索引。
- ⚠️ **page size 硬上限 50**：每个 contract 的 `transfer` count 子查询都要 ~0.4 s，top-200 会到 60 s。需要全榜请走后续的预聚合方案。
- ⚠️ Testnet 上 `erc20_holder_agg` 表是空的（Windmill 没在 testnet 跑这个 job），调 testnet 返回 `{}`。

---

## 5. 错误码

所有错误遵循 gRPC status code：

| HTTP status | gRPC code | 触发场景 |
|---|---|---|
| 400 | `INVALID_ARGUMENT` (3) | 日期格式不合法、`start > end`、空串 |
| 401 | — | JWT 缺失 / 无效 / 权限不足 |
| 500 | `UNKNOWN` (2) / `INTERNAL` (13) | 上游 DB 错误、未预期异常 |

错误体例（grpc-gateway 转换后）：

```json
{
  "code": 3,
  "message": "start date \"2026-05-15\" is after end date \"2026-05-01\""
}
```

## 6. 上游数据健康度参考

下面 3 张表是 Windmill 在维护，**不由本项目控制**。如果接口长期返回空 / 滞后明显，先看上游：

| 表 | 哪个 Windmill job 写入 | 当前健康度（截至 2026-05-19） |
|---|---|---|
| `iotexscanv3_kv` | `f/iotexscan/sync-iotex-statistics.deno.ts` | mainnet ✅ ~65 s 延迟；testnet ❌ 表空 |
| `erc20_holder_agg` | `f/analyzer/xrc20_holder_agg.deno.ts` | mainnet ✅ ；testnet ❌ 表空 |
| `block_supply` | analyser 自己写 | ✅ 实时（每 block）|

如果你需要 testnet 上 #1 #5 也展示数据，需要 DevOps 把上述 2 个 Windmill job 在 testnet 实例上也跑起来。

## 7. 旧需求文档差异速查

| 项 | 原始需求 | 实际 |
|---|---|---|
| HTTP 方法 | GET | **POST** |
| 参数位置 | query string `?start=...` | **JSON body** `{"start":"..."}` |
| URL 风格 | REST `/chain/stats` | gRPC `/api.ChainService.GetChainStats` |
| JSON 字段命名 | snake_case (`actions_num`) | **camelCase** (`actionsNum`) |
| uint64 编码 | number | **string**（避免 JS 精度丢失）|
| 列表响应 | `[...]` | **`{ "data": [...] }`** 或 **`{ "items": [...] }`** |
| #4 supply 单位 | IOTX (小数) | 一致 |
| #3 gas price 单位 | rau (string) | 一致 |

> 如果前端已经按原始需求写了一版 fetch / response parser，对照本表批量改一遍就行。
