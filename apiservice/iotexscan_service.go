package apiservice

// IotexscanService implements the etherscan-compatible endpoints that back the
// legacy iotexscan-v3 backend in iotex-kit (modules/modules-iotexscan). These
// were previously served by iotex-kit connecting directly to the analyzer
// Postgres via ANALYZER_DATABASE_URL; the SQL below is ported 1:1 from those
// kit handlers so the response shapes stay identical.
//
// STATUS: SKELETON. Each handler has the ported SQL and scan wiring. Table
// names were reconciled against the live schema used elsewhere in apiservice:
// block_action -> block_action_partition (used everywhere in this repo);
// block_receipt_transactions, block_receipt_logs, erc20/721_transfers,
// erc1155_transfer_singles_v2_2_2, action_execution, method_bytes, store all
// match existing queries (see account_service.go / action_service.go /
// xrc721_service.go / chain_service.go). Remaining TODO markers below cover the
// 0x/hex framing and column-set confirmations that need a live diff vs kit.

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IotexscanService struct {
	api.UnimplementedIotexscanServiceServer
}

// ─────────────────── helpers ───────────────────

// toIo normalizes a 0x/io address to io1... form. Empty stays empty.
func toIo(addr string) (string, error) {
	if addr == "" {
		return "", nil
	}
	if len(addr) >= 2 && (addr[:2] == "0x" || addr[:2] == "0X") {
		a, err := address.FromHex(addr)
		if err != nil {
			return "", err
		}
		return a.String(), nil
	}
	return addr, nil
}

// toHex converts an io1... address to its 0x hex form for etherscan-shaped
// output. Empty stays empty; a non-io value is returned unchanged so callers
// can pass through already-hex or sentinel values.
func toHex(addr string) string {
	if addr == "" {
		return ""
	}
	a, err := address.FromString(addr)
	if err != nil {
		return addr
	}
	return a.Hex()
}

// normSort clamps sort to "asc"/"desc" (default desc, matching kit).
func normSort(s string) string {
	if s == "asc" {
		return "asc"
	}
	return "desc"
}

// blockRangeClause appends optional block-height bounds and returns the SQL
// fragment plus the args to bind, in order. col is the qualified column name
// (e.g. "b.block_height").
func blockRangeClause(col string, start, end uint64) (string, []interface{}) {
	frag := ""
	args := []interface{}{}
	if start > 0 {
		frag += fmt.Sprintf(" and %s >= ?", col)
		args = append(args, start)
	}
	if end > 0 {
		frag += fmt.Sprintf(" and %s <= ?", col)
		args = append(args, end)
	}
	return frag, args
}

// transferInnerUnion builds an index-friendly two-leg UNION over a *_transfers
// table (one leg per sender/recipient), each with optional contract + block
// range filters and an inner LIMIT of skip+first (required for deep-page
// correctness). UNION (not UNION ALL) drops the self-transfer duplicate by id.
// Returns the SQL and its bind args in order. `sort` orders each leg's id.
// The caller wraps this (aliased `et`) and joins block_action/block.
func transferInnerUnion(table, owner, contract string, startBlock, endBlock, skip, first uint64, sort string) (string, []interface{}) {
	// Per-leg filter fragment (unqualified columns; runs inside the leg).
	legFilter := ""
	legArgs := []interface{}{}
	if contract != "" {
		legFilter += " AND contract_address = ?"
		legArgs = append(legArgs, contract)
	}
	if startBlock > 0 {
		legFilter += " AND block_height >= ?"
		legArgs = append(legArgs, startBlock)
	}
	if endBlock > 0 {
		legFilter += " AND block_height <= ?"
		legArgs = append(legArgs, endBlock)
	}
	inner := skip + first
	leg := func(col string) string {
		return "(SELECT * FROM " + table + " WHERE " + col + " = ?" + legFilter +
			" ORDER BY id " + sort + " LIMIT ?)"
	}
	sql := leg("sender") + " UNION " + leg("recipient")
	// Args: sender-leg (owner, legArgs..., inner), recipient-leg (owner, legArgs..., inner).
	args := []interface{}{}
	for i := 0; i < 2; i++ {
		args = append(args, owner)
		args = append(args, legArgs...)
		args = append(args, inner)
	}
	return sql, args
}

// ─────────────────── GetTxListByAddress (account.txlist) ───────────────────

func (s *IotexscanService) GetTxListByAddress(ctx context.Context, req *api.TransferListRequest) (*api.TxListResponse, error) {
	resp := &api.TxListResponse{}
	addr, err := toIo(req.GetAddress())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address: %v", err)
	}
	if addr == "" {
		return nil, status.Errorf(codes.InvalidArgument, "address is required")
	}
	sort := normSort(req.GetSort())
	limit := req.GetOffset()
	if limit == 0 {
		limit = 25
	}
	skip := (max64(req.GetPage(), 1) - 1) * limit

	rangeFrag, rangeArgs := blockRangeClause("b.block_height", req.GetStartBlock(), req.GetEndBlock())

	// Ported from kit account.txlist (analyzerSlow pool). NOTE: kit used
	// `block_action` / `action_execution` / `block` / `method_bytes`. TODO:
	// confirm partitioned table names in this schema.
	query := `SELECT
			b.block_height,
			ROUND(EXTRACT(EPOCH FROM b.timestamp)) AS time_stamp,
			('0x' || b.action_hash) AS hash,
			b.nonce,
			block.block_hash,
			b.sender,
			b.recipient,
			b.amount,
			b.gas_limit,
			b.gas_price,
			b.status,
			CASE WHEN b.status = 1 THEN 0 ELSE 1 END AS is_error,
			('0x' || encode(COALESCE(a.data, b.payload), 'hex')) AS input,
			b.contract_address,
			b.gas_consumed,
			('0x' || SUBSTRING(encode(a.data, 'hex'), 0, 9)) AS method_id,
			m."methodName" AS function_name
		FROM block_action_partition b
		LEFT JOIN action_execution a ON b.action_hash = a.action_hash
		LEFT JOIN block ON b.block_height = block.block_height
		LEFT JOIN method_bytes m ON SUBSTRING(encode(a.data, 'hex'), 0, 9) = m.bytecode
		WHERE (b.sender = ? OR b.recipient = ?)` + rangeFrag +
		` ORDER BY b.block_height ` + sort + ` LIMIT ? OFFSET ?`

	args := append([]interface{}{addr, addr}, rangeArgs...)
	args = append(args, limit, skip)

	rows, err := db.DB().WithContext(ctx).Raw(query, args...).Rows()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query txlist")
	}
	defer rows.Close()

	for rows.Next() {
		var (
			blockNumber, timeStamp, hash, nonce, blockHash, sender, recipient,
			amount, gas, gasPrice, status_, isError, input, contractAddr,
			gasUsed, methodID sql.NullString
			functionName sql.NullString
		)
		if err := rows.Scan(&blockNumber, &timeStamp, &hash, &nonce, &blockHash,
			&sender, &recipient, &amount, &gas, &gasPrice, &status_, &isError,
			&input, &contractAddr, &gasUsed, &methodID, &functionName); err != nil {
			return nil, errors.Wrap(err, "scan txlist row")
		}
		resp.Results = append(resp.Results, &api.NativeTxInfo{
			BlockNumber:     blockNumber.String,
			TimeStamp:       timeStamp.String,
			Hash:            hash.String,
			Nonce:           nonce.String,
			BlockHash:       blockHash.String,
			From:            toHex(sender.String),
			To:              toHex(recipient.String),
			Value:           amount.String,
			Gas:             gas.String,
			GasPrice:        gasPrice.String,
			IsError:         isError.String,
			TxreceiptStatus: status_.String,
			Input:           input.String,
			ContractAddress: toHex(contractAddr.String),
			GasUsed:         gasUsed.String,
			MethodId:        methodID.String,
			FunctionName:    functionName.String,
		})
	}
	return resp, rows.Err()
}

// ─────────────────── GetTokenTxByAddress (account.tokentx, ERC20) ───────────────────

func (s *IotexscanService) GetTokenTxByAddress(ctx context.Context, req *api.TransferListRequest) (*api.TokenTxResponse, error) {
	resp := &api.TokenTxResponse{}
	addr, contract, limit, skip, sort, _, _, err := parseTransferReq(req, "et.block_height")
	if err != nil {
		return nil, err
	}

	// Ported from kit account.tokentx (erc20_transfers + block_action + block).
	// A plain `(sender=? OR recipient=?) ORDER BY block_height` on erc20_transfers
	// (220M+ rows) makes the planner do a backward scan (~120s for late matches).
	// Split into two index-friendly legs (each hits its own (sender|recipient, id)
	// index) UNION'd together, then join block_action/block on the small result
	// set. Same optimization as common/actions/xrc20.go GetXrc20InfoByAddress.
	inner, innerArgs := transferInnerUnion("erc20_transfers", addr, contract, req.GetStartBlock(), req.GetEndBlock(), skip, limit, sort)
	query := `SELECT
			et.block_height,
			ROUND(EXTRACT(EPOCH FROM ba.timestamp)) AS time_stamp,
			('0x' || et.action_hash) AS hash,
			ba.nonce,
			b.block_hash,
			et.sender,
			et.recipient,
			et.contract_address,
			et.amount,
			ba.gas_limit,
			ba.gas_price,
			ba.gas_consumed
		FROM (` + inner + `) et
		LEFT JOIN block_action_partition ba ON et.action_hash = ba.action_hash
		LEFT JOIN block b ON et.block_height = b.block_height
		ORDER BY et.id ` + sort + ` LIMIT ? OFFSET ?`

	args := append(innerArgs, limit, skip)

	rows, err := db.DB().WithContext(ctx).Raw(query, args...).Rows()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query tokentx")
	}
	defer rows.Close()
	for rows.Next() {
		var (
			blockNumber, timeStamp, hash, nonce, blockHash, sender, recipient,
			contractAddr, value, gas, gasPrice, gasUsed sql.NullString
		)
		if err := rows.Scan(&blockNumber, &timeStamp, &hash, &nonce, &blockHash,
			&sender, &recipient, &contractAddr, &value, &gas, &gasPrice, &gasUsed); err != nil {
			return nil, errors.Wrap(err, "scan tokentx row")
		}
		resp.Results = append(resp.Results, &api.Erc20TransferInfo{
			BlockNumber:     blockNumber.String,
			TimeStamp:       timeStamp.String,
			Hash:            hash.String,
			Nonce:           nonce.String,
			BlockHash:       blockHash.String,
			From:            toHex(sender.String),
			To:              toHex(recipient.String),
			ContractAddress: toHex(contractAddr.String),
			Value:           value.String,
			Gas:             gas.String,
			GasPrice:        gasPrice.String,
			GasUsed:         gasUsed.String,
		})
	}
	return resp, rows.Err()
}

// ─────────────────── GetTokenNftTxByAddress (account.tokennfttx, ERC721) ───────────────────

func (s *IotexscanService) GetTokenNftTxByAddress(ctx context.Context, req *api.TransferListRequest) (*api.TokenNftTxResponse, error) {
	resp := &api.TokenNftTxResponse{}
	addr, contract, limit, skip, sort, _, _, err := parseTransferReq(req, "et.block_height")
	if err != nil {
		return nil, err
	}

	// Ported from kit account.tokennfttx. kit's SQL said `erc721_transfers`, but
	// this schema's real table is erc721_transfers_v2_2_3 (see xrc721_service.go).
	// Two-leg UNION for index-friendly sender/recipient scans (see tokentx).
	inner, innerArgs := transferInnerUnion("erc721_transfers_v2_2_3", addr, contract, req.GetStartBlock(), req.GetEndBlock(), skip, limit, sort)
	query := `SELECT
			et.block_height,
			ROUND(EXTRACT(EPOCH FROM ba.timestamp)) AS time_stamp,
			('0x' || et.action_hash) AS hash,
			ba.nonce,
			b.block_hash,
			et.sender,
			et.recipient,
			et.contract_address,
			et.token_id,
			ba.gas_limit,
			ba.gas_price,
			ba.gas_consumed
		FROM (` + inner + `) et
		LEFT JOIN block_action_partition ba ON et.action_hash = ba.action_hash
		LEFT JOIN block b ON et.block_height = b.block_height
		ORDER BY et.id ` + sort + ` LIMIT ? OFFSET ?`

	args := append(innerArgs, limit, skip)

	rows, err := db.DB().WithContext(ctx).Raw(query, args...).Rows()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query tokennfttx")
	}
	defer rows.Close()
	for rows.Next() {
		var (
			blockNumber, timeStamp, hash, nonce, blockHash, sender, recipient,
			contractAddr, tokenID, gas, gasPrice, gasUsed sql.NullString
		)
		if err := rows.Scan(&blockNumber, &timeStamp, &hash, &nonce, &blockHash,
			&sender, &recipient, &contractAddr, &tokenID, &gas, &gasPrice, &gasUsed); err != nil {
			return nil, errors.Wrap(err, "scan tokennfttx row")
		}
		resp.Results = append(resp.Results, &api.Erc721TransferInfo{
			BlockNumber:     blockNumber.String,
			TimeStamp:       timeStamp.String,
			Hash:            hash.String,
			Nonce:           nonce.String,
			BlockHash:       blockHash.String,
			From:            toHex(sender.String),
			To:              toHex(recipient.String),
			ContractAddress: toHex(contractAddr.String),
			TokenId:         tokenID.String,
			Gas:             gas.String,
			GasPrice:        gasPrice.String,
			GasUsed:         gasUsed.String,
		})
	}
	return resp, rows.Err()
}

// ─────────────────── GetToken1155TxByAddress (account.token1155tx) ───────────────────

func (s *IotexscanService) GetToken1155TxByAddress(ctx context.Context, req *api.TransferListRequest) (*api.Token1155TxResponse, error) {
	resp := &api.Token1155TxResponse{}
	addr, contract, limit, skip, sort, _, _, err := parseTransferReq(req, "et.block_height")
	if err != nil {
		return nil, err
	}

	// Ported from kit account.token1155tx (erc1155_transfer_singles_v2_2_2 + block_action + block).
	// Two-leg UNION for index-friendly sender/recipient scans (see tokentx).
	inner, innerArgs := transferInnerUnion("erc1155_transfer_singles_v2_2_2", addr, contract, req.GetStartBlock(), req.GetEndBlock(), skip, limit, sort)
	query := `SELECT
			et.block_height,
			ROUND(EXTRACT(EPOCH FROM ba.timestamp)) AS time_stamp,
			('0x' || et.action_hash) AS hash,
			ba.nonce,
			b.block_hash,
			et.sender,
			et.recipient,
			et.contract_address,
			et._id,
			et.value,
			ba.gas_limit,
			ba.gas_price,
			ba.gas_consumed
		FROM (` + inner + `) et
		LEFT JOIN block_action_partition ba ON et.action_hash = ba.action_hash
		LEFT JOIN block b ON et.block_height = b.block_height
		ORDER BY et.id ` + sort + ` LIMIT ? OFFSET ?`

	args := append(innerArgs, limit, skip)

	rows, err := db.DB().WithContext(ctx).Raw(query, args...).Rows()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query token1155tx")
	}
	defer rows.Close()
	for rows.Next() {
		var (
			blockNumber, timeStamp, hash, nonce, blockHash, sender, recipient,
			contractAddr, tokenID, tokenValue, gas, gasPrice, gasUsed sql.NullString
		)
		if err := rows.Scan(&blockNumber, &timeStamp, &hash, &nonce, &blockHash,
			&sender, &recipient, &contractAddr, &tokenID, &tokenValue, &gas, &gasPrice, &gasUsed); err != nil {
			return nil, errors.Wrap(err, "scan token1155tx row")
		}
		resp.Results = append(resp.Results, &api.Erc1155TransferInfo{
			BlockNumber:     blockNumber.String,
			TimeStamp:       timeStamp.String,
			Hash:            hash.String,
			Nonce:           nonce.String,
			BlockHash:       blockHash.String,
			From:            toHex(sender.String),
			To:              toHex(recipient.String),
			ContractAddress: toHex(contractAddr.String),
			TokenId:         tokenID.String,
			TokenValue:      tokenValue.String,
			Gas:             gas.String,
			GasPrice:        gasPrice.String,
			GasUsed:         gasUsed.String,
		})
	}
	return resp, rows.Err()
}

// ─────────────────── GetTxListInternal (account.txlistinternal) ───────────────────

func (s *IotexscanService) GetTxListInternal(ctx context.Context, req *api.InternalTransferRequest) (*api.TxListInternalResponse, error) {
	resp := &api.TxListInternalResponse{}
	addr, err := toIo(req.GetAddress())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address: %v", err)
	}
	sort := normSort(req.GetSort())
	limit := req.GetOffset()
	if limit == 0 {
		limit = 25
	}
	skip := (max64(req.GetPage(), 1) - 1) * limit

	// Ported from kit account.txlistinternal
	// (block_receipt_transactions brt JOIN account_meta on contract sender).
	// TODO: kit `SELECT *`'d brt; enumerate the exact columns the kit response
	// consumers expect and map them here instead of *.
	where := "am.is_contract = true"
	args := []interface{}{}
	if addr != "" {
		where += " AND (brt.sender = ? OR brt.recipient = ?)"
		args = append(args, addr, addr)
	}
	if req.GetTxHash() != "" {
		where += " AND brt.action_hash = ?"
		args = append(args, req.GetTxHash())
	}
	rangeFrag, rangeArgs := blockRangeClause("brt.block_height", req.GetStartBlock(), req.GetEndBlock())
	where += rangeFrag
	args = append(args, rangeArgs...)

	query := `SELECT brt.block_height, brt.action_hash, brt.sender, brt.recipient, brt.amount,
			ROUND(EXTRACT(EPOCH FROM b.timestamp)) AS time_stamp
		FROM block_receipt_transactions brt
		INNER JOIN account_meta am ON brt.sender = am.address
		LEFT JOIN block b ON brt.block_height = b.block_height
		WHERE ` + where +
		` ORDER BY brt.block_height ` + sort + ` LIMIT ? OFFSET ?`
	args = append(args, limit, skip)

	rows, err := db.DB().WithContext(ctx).Raw(query, args...).Rows()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query txlistinternal")
	}
	defer rows.Close()
	for rows.Next() {
		var blockHeight, actionHash, sender, recipient, amount, timeStamp sql.NullString
		if err := rows.Scan(&blockHeight, &actionHash, &sender, &recipient, &amount, &timeStamp); err != nil {
			return nil, errors.Wrap(err, "scan txlistinternal row")
		}
		resp.Results = append(resp.Results, &api.InternalTransferInfo{
			BlockHeight: blockHeight.String,
			ActionHash:  actionHash.String,
			Sender:      toHex(sender.String),
			Recipient:   toHex(recipient.String),
			Amount:      amount.String,
			TimeStamp:   timeStamp.String,
		})
	}
	return resp, rows.Err()
}

// ─────────────────── GetContractLogs (log.getLogs) ───────────────────

func (s *IotexscanService) GetContractLogs(ctx context.Context, req *api.GetContractLogsRequest) (*api.GetContractLogsResponse, error) {
	resp := &api.GetContractLogsResponse{}
	addr, err := toIo(req.GetAddress())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid address: %v", err)
	}
	if addr == "" {
		return nil, status.Errorf(codes.InvalidArgument, "address is required")
	}
	limit := req.GetOffset()
	if limit == 0 {
		limit = 25
	}
	skip := (max64(req.GetPage(), 1) - 1) * limit

	// Ported from kit log.getLogs (block_receipt_logs brl LEFT JOIN block_action ba).
	query := `SELECT
			brl.block_height, brl.address, brl.topic0, brl.topic1, brl.topic2, brl.topic3,
			encode(brl.data, 'hex') AS data, brl.action_hash, brl.index,
			ba.gas_consumed, EXTRACT(EPOCH FROM ba.timestamp) AS timestamp, ba.gas_price
		FROM block_receipt_logs brl
		LEFT JOIN block_action_partition ba ON brl.action_hash = ba.action_hash
		WHERE brl.block_height >= ? AND brl.block_height <= ? AND brl.address = ?
		LIMIT ? OFFSET ?`

	rows, err := db.DB().WithContext(ctx).Raw(query,
		req.GetFromBlock(), req.GetToBlock(), addr, limit, skip).Rows()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query logs")
	}
	defer rows.Close()
	for rows.Next() {
		var (
			blockHeight, index, gasConsumed sql.NullInt64
			timestamp                       sql.NullFloat64
			logAddr, t0, t1, t2, t3, data, actionHash, gasPrice sql.NullString
		)
		if err := rows.Scan(&blockHeight, &logAddr, &t0, &t1, &t2, &t3, &data,
			&actionHash, &index, &gasConsumed, &timestamp, &gasPrice); err != nil {
			return nil, errors.Wrap(err, "scan log row")
		}
		resp.Logs = append(resp.Logs, &api.ContractLogInfo{
			BlockHeight: uint64(blockHeight.Int64),
			Address:     logAddr.String,
			Topic0:      t0.String,
			Topic1:      t1.String,
			Topic2:      t2.String,
			Topic3:      t3.String,
			Data:        data.String,
			ActionHash:  actionHash.String,
			Index:       uint64(index.Int64),
			GasConsumed: uint64(gasConsumed.Int64),
			Timestamp:   int64(timestamp.Float64),
			GasPrice:    gasPrice.String,
		})
	}
	return resp, rows.Err()
}

// ─────────────────── GetGasOracle (gastracker.gasoracle) ───────────────────

func (s *IotexscanService) GetGasOracle(ctx context.Context, req *api.GetGasOracleRequest) (*api.GetGasOracleResponse, error) {
	resp := &api.GetGasOracleResponse{}
	var row struct {
		Value sql.NullString
	}
	// Ported from kit gastracker.gasoracle: SELECT value FROM store WHERE key='iotx_gas_oracle'.
	// The value is a JSON blob; return it raw so kit keeps parsing/scaling it.
	if err := db.DB().WithContext(ctx).Raw(
		`SELECT value FROM store WHERE key = 'iotx_gas_oracle' LIMIT 1`,
	).Scan(&row).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query gas oracle")
	}
	if row.Value.Valid {
		resp.Exist = true
		resp.ValueJson = row.Value.String
	}
	return resp, nil
}

// ─────────────────── GetDailyNewAddresses (stats.dailynewaddress) ───────────────────

func (s *IotexscanService) GetDailyNewAddresses(ctx context.Context, req *api.GetDailyNewAddressesRequest) (*api.GetDailyNewAddressesResponse, error) {
	resp := &api.GetDailyNewAddressesResponse{}
	if req.GetStartDate() == "" || req.GetEndDate() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "start_date and end_date are required")
	}
	sort := normSort(req.GetSort())

	// Ported from kit stats.dailynewaddress, but with the date range parameterized
	// (kit hard-coded 2024-01-01..2024-08-01 — a bug/placeholder we fix here).
	query := `SELECT count(1) AS cnt, (b.timestamp::date AT TIME ZONE 'UTC') AS date
		FROM account_meta am
		LEFT JOIN block b ON am.block_height = b.block_height
		WHERE b.timestamp::date >= ? AND b.timestamp::date <= ?
		GROUP BY b.timestamp::date
		ORDER BY b.timestamp::date ` + sort

	rows, err := db.DB().WithContext(ctx).Raw(query, req.GetStartDate(), req.GetEndDate()).Rows()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query daily new addresses")
	}
	defer rows.Close()
	for rows.Next() {
		var cnt sql.NullInt64
		var date sql.NullString
		if err := rows.Scan(&cnt, &date); err != nil {
			return nil, errors.Wrap(err, "scan daily new address row")
		}
		resp.Data = append(resp.Data, &api.DailyNewAddressPoint{
			Date:            date.String,
			NewAddressCount: uint64(cnt.Int64),
		})
	}
	return resp, rows.Err()
}

// ─────────────────── GetContractCreationBatch (contract.getcontractcreation) ───────────────────

func (s *IotexscanService) GetContractCreationBatch(ctx context.Context, req *api.GetContractCreationBatchRequest) (*api.GetContractCreationBatchResponse, error) {
	resp := &api.GetContractCreationBatchResponse{}
	if len(req.GetAddresses()) == 0 {
		return resp, nil
	}
	ioAddrs := make([]string, 0, len(req.GetAddresses()))
	for _, a := range req.GetAddresses() {
		io, err := toIo(a)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid address %q: %v", a, err)
		}
		ioAddrs = append(ioAddrs, io)
	}

	// Ported from kit contract.getcontractcreationLoader
	// (account_meta am JOIN block_action a on contract creation).
	// Use IN (?) with a slice: GORM expands it to IN ($1,$2,...). ANY(?) does
	// NOT work here — GORM/pgx won't render a Go slice as a PG array literal.
	query := `SELECT a.action_hash, a.sender AS contract_creator, am.address AS contract_address
		FROM account_meta am
		INNER JOIN block_action_partition a ON am.block_height = a.block_height AND am.address = a.contract_address
		WHERE am.address IN (?)`

	rows, err := db.DB().WithContext(ctx).Raw(query, ioAddrs).Rows()
	if err != nil {
		return nil, errors.Wrap(err, "failed to query contract creation batch")
	}
	defer rows.Close()
	for rows.Next() {
		var actionHash, creator, contractAddr sql.NullString
		if err := rows.Scan(&actionHash, &creator, &contractAddr); err != nil {
			return nil, errors.Wrap(err, "scan contract creation row")
		}
		resp.Items = append(resp.Items, &api.ContractCreationInfo{
			ContractAddress: toHex(contractAddr.String),
			ContractCreator: toHex(creator.String),
			ActionHash:      actionHash.String,
		})
	}
	return resp, rows.Err()
}

// ─────────────────── GetBlockNumberByTime (block.getblocknobytime) ───────────────────

func (s *IotexscanService) GetBlockNumberByTime(ctx context.Context, req *api.GetBlockNumberByTimeRequest) (*api.GetBlockNumberByTimeResponse, error) {
	resp := &api.GetBlockNumberByTimeResponse{}
	// kit builds ISO from unix seconds: dayjs(ts*1000).toISOString().
	iso := time.Unix(req.GetTimestamp(), 0).UTC().Format(time.RFC3339)

	// Ported from kit block.getblocknobytime. Strict < / > with ORDER BY to pick
	// the nearest block before/after the instant.
	var query string
	if req.GetClosest() == "after" {
		query = `SELECT block_height FROM block WHERE timestamp > ? ORDER BY timestamp ASC LIMIT 1`
	} else {
		query = `SELECT block_height FROM block WHERE timestamp < ? ORDER BY timestamp DESC LIMIT 1`
	}
	var row struct {
		BlockHeight sql.NullInt64
	}
	if err := db.DB().WithContext(ctx).Raw(query, iso).Scan(&row).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query block number by time")
	}
	if row.BlockHeight.Valid {
		resp.Exist = true
		resp.BlockHeight = uint64(row.BlockHeight.Int64)
	}
	return resp, nil
}

// ─────────────────── GetActionStatusByHash (transaction.getstatus/gettxreceiptstatus) ───────────────────

func (s *IotexscanService) GetActionStatusByHash(ctx context.Context, req *api.GetActionStatusByHashRequest) (*api.GetActionStatusByHashResponse, error) {
	resp := &api.GetActionStatusByHashResponse{}
	hash := req.GetActionHash()
	if hash == "" {
		return nil, status.Errorf(codes.InvalidArgument, "action_hash is required")
	}
	// kit stored hashes without 0x; strip a leading 0x/0X if present.
	if len(hash) >= 2 && (hash[:2] == "0x" || hash[:2] == "0X") {
		hash = hash[2:]
	}

	// Ported from kit transaction.getstatus tx-part (SELECT * + isError).
	query := `SELECT a.action_hash, a.block_height, a.action_type, a.sender, a.recipient,
			a.amount, a.gas_price, a.gas_limit, r.gas_consumed, a.nonce, a.status,
			CASE WHEN a.status = 1 THEN 0 ELSE 1 END AS is_error,
			a.contract_address, ROUND(EXTRACT(EPOCH FROM b.timestamp)) AS timestamp
		FROM block_action_partition a
		LEFT JOIN block b ON a.block_height = b.block_height
		LEFT JOIN block_receipts r ON a.action_hash = r.action_hash
		WHERE a.action_hash = ? LIMIT 1`

	var row struct {
		ActionHash      sql.NullString
		BlockHeight     sql.NullInt64
		ActionType      sql.NullString
		Sender          sql.NullString
		Recipient       sql.NullString
		Amount          sql.NullString
		GasPrice        sql.NullString
		GasLimit        sql.NullInt64
		GasConsumed     sql.NullInt64
		Nonce           sql.NullInt64
		Status          sql.NullInt64
		IsError         sql.NullInt64
		ContractAddress sql.NullString
		Timestamp       sql.NullInt64
	}
	if err := db.DB().WithContext(ctx).Raw(query, hash).Scan(&row).Error; err != nil {
		return nil, errors.Wrap(err, "failed to query action status")
	}
	if !row.ActionHash.Valid {
		return resp, nil
	}
	resp.Exist = true
	resp.Action = &api.ActionStatusInfo{
		ActionHash:      row.ActionHash.String,
		BlockHeight:     uint64(row.BlockHeight.Int64),
		ActionType:      row.ActionType.String,
		Sender:          row.Sender.String,
		Recipient:       row.Recipient.String,
		Amount:          row.Amount.String,
		GasPrice:        row.GasPrice.String,
		GasLimit:        uint64(row.GasLimit.Int64),
		GasConsumed:     uint64(row.GasConsumed.Int64),
		Nonce:           uint64(row.Nonce.Int64),
		Status:          uint64(row.Status.Int64),
		IsError:         uint64(row.IsError.Int64),
		ContractAddress: row.ContractAddress.String,
		Timestamp:       row.Timestamp.Int64,
	}

	// Attached staking bucket: first match wins in kit's precedence order
	// (nft_staking_v1 -> nft_staking_v2 -> native_staking).
	if bucket := s.lookupStatusBucket(ctx, hash); bucket != nil {
		resp.BucketExist = true
		resp.Bucket = bucket
	}
	return resp, nil
}

// lookupStatusBucket mirrors kit transaction.getstatus: query the three staking
// sources in precedence order and return the first hit tagged with its type.
func (s *IotexscanService) lookupStatusBucket(ctx context.Context, hash string) *api.ActionStatusBucket {
	type bucketRow struct {
		BucketId         sql.NullString
		Amount           sql.NullString
		AutoStake        sql.NullBool
		Duration         sql.NullString
		CreateTime       sql.NullString
		StakeStartTime   sql.NullString
		UnstakeStartTime sql.NullString
		StakedAmount     sql.NullString
		OwnerAddress     sql.NullString
		EventType        sql.NullString
	}
	sources := []struct {
		typ   string
		query string
	}{
		{"nft_staking_v1", `SELECT bucket_id, amount, auto_stake, duration, create_time, stake_start_time, unstake_start_time, staked_amount, owner_address, event_type FROM system_staking_buckets WHERE act_hash = ? ORDER BY id DESC LIMIT 1`},
		{"nft_staking_v2", `SELECT bucket_id, amount, auto_stake, duration, create_time, stake_start_time, unstake_start_time, staked_amount, owner_address, event_type FROM system_staking_buckets_v2 WHERE act_hash = ? ORDER BY id DESC LIMIT 1`},
		{"native_staking", `SELECT bucket_id, amount, auto_stake, duration, create_time, stake_start_time, unstake_start_time, staked_amount, owner_address, act_type AS event_type FROM staking_buckets WHERE action_hash = ? ORDER BY id DESC LIMIT 1`},
	}
	for _, src := range sources {
		var r bucketRow
		if err := db.DB().WithContext(ctx).Raw(src.query, hash).Scan(&r).Error; err != nil {
			continue
		}
		if !r.BucketId.Valid {
			continue
		}
		return &api.ActionStatusBucket{
			BucketId:         r.BucketId.String,
			Amount:           r.Amount.String,
			AutoStake:        r.AutoStake.Bool,
			Duration:         r.Duration.String,
			CreateTime:       r.CreateTime.String,
			StakeStartTime:   r.StakeStartTime.String,
			UnstakeStartTime: r.UnstakeStartTime.String,
			StakedAmount:     r.StakedAmount.String,
			OwnerAddress:     r.OwnerAddress.String,
			EventType:        r.EventType.String,
			Type:             src.typ,
		}
	}
	return nil
}

// ─────────────────── shared request parsing ───────────────────

// parseTransferReq normalizes the common TransferListRequest fields and builds
// the optional block-range clause. blockCol is the qualified height column used
// for the range bounds (e.g. "et.block_height").
func parseTransferReq(req *api.TransferListRequest, blockCol string) (
	addr, contract string, limit, skip uint64, sort, rangeFrag string, rangeArgs []interface{}, err error,
) {
	addr, err = toIo(req.GetAddress())
	if err != nil {
		return "", "", 0, 0, "", "", nil, status.Errorf(codes.InvalidArgument, "invalid address: %v", err)
	}
	if addr == "" {
		return "", "", 0, 0, "", "", nil, status.Errorf(codes.InvalidArgument, "address is required")
	}
	contract, err = toIo(req.GetContractAddress())
	if err != nil {
		return "", "", 0, 0, "", "", nil, status.Errorf(codes.InvalidArgument, "invalid contract_address: %v", err)
	}
	limit = req.GetOffset()
	if limit == 0 {
		limit = 25
	}
	skip = (max64(req.GetPage(), 1) - 1) * limit
	sort = normSort(req.GetSort())
	rangeFrag, rangeArgs = blockRangeClause(blockCol, req.GetStartBlock(), req.GetEndBlock())
	return addr, contract, limit, skip, sort, rangeFrag, rangeArgs, nil
}

func max64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
