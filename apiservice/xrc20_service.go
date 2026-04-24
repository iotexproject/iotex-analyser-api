package apiservice

import (
	"context"
	"database/sql"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/actions"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/pkg/errors"
)

type XRC20Service struct {
	api.UnimplementedXRC20ServiceServer
}

// XRC20ByAddress returns the XRC20 info of the address.
func (s *XRC20Service) XRC20ByAddress(ctx context.Context, req *api.XRC20ByAddressRequest) (*api.XRC20ByAddressResponse, error) {
	resp := &api.XRC20ByAddressResponse{
		Count: 0,
		Exist: false,
		Xrc20: make([]*api.Xrc20Action, 0),
	}
	address, err := common.Address(req.GetAddress())
	if err != nil {
		return nil, err
	}
	count, err := actions.GetXrc20CountByAddress(*address)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return resp, nil
	}
	resp.Exist = true
	resp.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	infos, err := actions.GetXrc20InfoByAddress(*address, skip, first)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		resp.Xrc20 = append(resp.Xrc20, &api.Xrc20Action{
			ActHash:   info.ActionHash,
			Timestamp: uint64(info.Timestamp.Unix()),
			Sender:    info.Sender,
			Recipient: info.Recipient,
			Amount:    info.Amount,
			Contract:  info.ContractAddress,
		})
	}
	return resp, nil
}

// XRC20ByContractAddress returns Xrc20 actions given the Xrc20 contract address
func (s *XRC20Service) XRC20ByContractAddress(ctx context.Context, req *api.XRC20ByContractAddressRequest) (*api.XRC20ByContractAddressResponse, error) {
	resp := &api.XRC20ByContractAddressResponse{
		Count: 0,
		Exist: false,
		Xrc20: make([]*api.Xrc20Action, 0),
	}
	address, err := common.Address(req.GetAddress())
	if err != nil {
		return nil, err
	}
	count, err := actions.GetXrc20CountByContractAddress(*address)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return resp, nil
	}
	resp.Exist = true
	resp.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	infos, err := actions.GetXrc20InfoByContractAddress(*address, skip, first)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		resp.Xrc20 = append(resp.Xrc20, &api.Xrc20Action{
			ActHash:   info.ActionHash,
			Timestamp: uint64(info.Timestamp.Unix()),
			Sender:    info.Sender,
			Recipient: info.Recipient,
			Amount:    info.Amount,
			Contract:  info.ContractAddress,
		})
	}
	return resp, nil
}

// XRC20ByPage returns Xrc20 actions by pagination
func (s *XRC20Service) XRC20ByPage(ctx context.Context, req *api.XRC20ByPageRequest) (*api.XRC20ByPageResponse, error) {
	resp := &api.XRC20ByPageResponse{
		Count: 0,
		Exist: false,
		Xrc20: make([]*api.Xrc20Action, 0),
	}
	count, err := actions.GetXrc20Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return resp, nil
	}
	resp.Exist = true
	resp.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	infos, err := actions.GetXrc20InfoByPage(skip, first)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		resp.Xrc20 = append(resp.Xrc20, &api.Xrc20Action{
			ActHash:   info.ActionHash,
			Timestamp: uint64(info.Timestamp.Unix()),
			Sender:    info.Sender,
			Recipient: info.Recipient,
			Amount:    info.Amount,
			Contract:  info.ContractAddress,
		})
	}
	return resp, nil
}

// XRC20Addresses returns Xrc20 contract addresses
func (s *XRC20Service) XRC20Addresses(ctx context.Context, req *api.XRC20AddressesRequest) (*api.XRC20AddressesResponse, error) {
	resp := &api.XRC20AddressesResponse{
		Count:     0,
		Exist:     false,
		Addresses: make([]string, 0),
	}
	count, err := actions.GetXrc20ContractAddressCount()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return resp, nil
	}
	resp.Exist = true
	resp.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	resp.Addresses, err = actions.GetXrc20ContractAddress(skip, first)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// XRC20TokenHolderAddresses returns Xrc20 token holder addresses given a Xrc20 contract address
func (s *XRC20Service) XRC20TokenHolderAddresses(ctx context.Context, req *api.XRC20TokenHolderAddressesRequest) (*api.XRC20TokenHolderAddressesResponse, error) {
	resp := &api.XRC20TokenHolderAddressesResponse{
		Count:     0,
		Addresses: make([]string, 0),
	}
	tokenAddress := req.GetTokenAddress()
	count, err := actions.GetXrc20TokenHolderCountByTokenAddress(tokenAddress)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return resp, nil
	}
	resp.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	resp.Addresses, err = actions.GetXrc20TokenHolder(tokenAddress, skip, first)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetXRC20TransfersByContract returns ERC20 transfers filtered by contract and optional sender/recipient.
func (s *XRC20Service) GetXRC20TransfersByContract(ctx context.Context, req *api.GetXRC20TransfersByContractRequest) (*api.GetXRC20TransfersByContractResponse, error) {
	resp := &api.GetXRC20TransfersByContractResponse{}
	gormDB := db.DB()
	contract := req.GetContractAddress()
	address := req.GetAddress()

	tsExpr := `to_char(timestamp AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"')`

	var count int64
	if address != "" {
		if err := gormDB.WithContext(ctx).Raw(
			`SELECT ((SELECT COUNT(1) FROM erc20_transfers WHERE contract_address = ? AND sender = ?)
			        +(SELECT COUNT(1) FROM erc20_transfers WHERE contract_address = ? AND recipient = ?)) AS count`,
			contract, address, contract, address,
		).Scan(&count).Error; err != nil {
			return nil, errors.Wrap(err, "failed to count xrc20 transfers")
		}
	} else {
		if err := gormDB.WithContext(ctx).Raw(
			`SELECT COUNT(1) FROM erc20_transfers WHERE contract_address = ?`, contract,
		).Scan(&count).Error; err != nil {
			return nil, errors.Wrap(err, "failed to count xrc20 transfers")
		}
	}
	resp.Count = count
	if count == 0 {
		return resp, nil
	}
	resp.Exist = true

	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	queryLimit := skip + first

	var rows []struct {
		Id              uint64
		BlockHeight     uint64
		ActionHash      string
		ContractAddress string
		Amount          string
		Sender          string
		Recipient       string
		Timestamp       sql.NullString
	}

	var q string
	if address != "" {
		q = `SELECT id, block_height, action_hash, contract_address, amount, sender, recipient,
		          ` + tsExpr + ` AS timestamp
		     FROM erc20_transfers
		     WHERE id = ANY(ARRAY(
		       (SELECT id FROM erc20_transfers WHERE contract_address = ? AND sender = ? ORDER BY id DESC LIMIT ?)
		       UNION
		       (SELECT id FROM erc20_transfers WHERE contract_address = ? AND recipient = ? ORDER BY id DESC LIMIT ?)
		     ))
		     ORDER BY id DESC LIMIT ? OFFSET ?`
		if err := gormDB.WithContext(ctx).Raw(q,
			contract, address, queryLimit,
			contract, address, queryLimit,
			first, skip,
		).Scan(&rows).Error; err != nil {
			return nil, errors.Wrap(err, "failed to get xrc20 transfers")
		}
	} else {
		q = `SELECT id, block_height, action_hash, contract_address, amount, sender, recipient,
		          ` + tsExpr + ` AS timestamp
		     FROM erc20_transfers
		     WHERE contract_address = ?
		     ORDER BY id DESC LIMIT ? OFFSET ?`
		if err := gormDB.WithContext(ctx).Raw(q, contract, first, skip).Scan(&rows).Error; err != nil {
			return nil, errors.Wrap(err, "failed to get xrc20 transfers")
		}
	}
	for _, r := range rows {
		info := &api.XRC20TransferInfo{
			Id:              r.Id,
			BlockHeight:     r.BlockHeight,
			ActionHash:      r.ActionHash,
			ContractAddress: r.ContractAddress,
			Amount:          r.Amount,
			Sender:          r.Sender,
			Recipient:       r.Recipient,
		}
		if r.Timestamp.Valid {
			info.Timestamp = r.Timestamp.String
		}
		resp.Transfers = append(resp.Transfers, info)
	}
	return resp, nil
}

// GetXRC20HoldersByContract returns all holders of a given ERC20 token with positive balances.
func (s *XRC20Service) GetXRC20HoldersByContract(ctx context.Context, req *api.GetXRC20HoldersByContractRequest) (*api.GetXRC20HoldersByContractResponse, error) {
	resp := &api.GetXRC20HoldersByContractResponse{}
	gormDB := db.DB()
	contract := req.GetContractAddress()

	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	var rows []struct {
		Address string
		Balance string
	}
	q := `SELECT holder AS address, balance::text AS balance
          FROM (
              SELECT holder, SUM(balance_change) AS balance
              FROM (
                  SELECT recipient AS holder, amount::numeric AS balance_change
                  FROM erc20_transfers WHERE contract_address = ?
                  UNION ALL
                  SELECT sender AS holder, -amount::numeric AS balance_change
                  FROM erc20_transfers WHERE contract_address = ?
              ) t
              GROUP BY holder
              HAVING SUM(balance_change) > 0
          ) h
          ORDER BY balance DESC
          LIMIT ? OFFSET ?`
	if err := gormDB.WithContext(ctx).Raw(q, contract, contract, first, skip).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get xrc20 holders")
	}
	resp.Count = int64(len(rows))
	for _, r := range rows {
		resp.Holders = append(resp.Holders, &api.XRC20HolderInfo{
			Address: r.Address,
			Balance: r.Balance,
		})
	}
	return resp, nil
}

// GetXRC20TokenBalance returns the ERC20 token balance for a specific address.
func (s *XRC20Service) GetXRC20TokenBalance(ctx context.Context, req *api.GetXRC20TokenBalanceRequest) (*api.GetXRC20TokenBalanceResponse, error) {
	resp := &api.GetXRC20TokenBalanceResponse{}
	gormDB := db.DB()
	contract := req.GetContractAddress()
	address := req.GetAddress()

	var row struct {
		Balance sql.NullString
	}
	q := `SELECT (
              COALESCE(SUM(amount::numeric) FILTER (WHERE recipient = ?), 0)
              - COALESCE(SUM(amount::numeric) FILTER (WHERE sender = ?), 0)
          )::text AS balance
          FROM erc20_transfers
          WHERE contract_address = ? AND (recipient = ? OR sender = ?)`
	if err := gormDB.WithContext(ctx).Raw(q, address, address, contract, address, address).Scan(&row).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get xrc20 token balance")
	}
	if row.Balance.Valid {
		resp.Balance = row.Balance.String
	} else {
		resp.Balance = "0"
	}
	return resp, nil
}
