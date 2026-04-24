package apiservice

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/actions"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/pkg/errors"
)

type XRC721Service struct {
	api.UnimplementedXRC721ServiceServer
}

// XRC721ByAddress returns the XRC721 info of the address.
func (s *XRC721Service) XRC721ByAddress(ctx context.Context, req *api.XRC721ByAddressRequest) (*api.XRC721ByAddressResponse, error) {
	resp := &api.XRC721ByAddressResponse{
		Count:  0,
		Exist:  false,
		Xrc721: make([]*api.Xrc721Action, 0),
	}
	address, err := common.Address(req.GetAddress())
	if err != nil {
		return nil, err
	}
	count, err := actions.GetXrc721CountByAddress(*address)
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
	infos, err := actions.GetXrc721InfoByAddress(*address, skip, first)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		resp.Xrc721 = append(resp.Xrc721, &api.Xrc721Action{
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

// XRC721ByContractAddress returns Xrc721 actions given the Xrc721 contract address
func (s *XRC721Service) XRC721ByContractAddress(ctx context.Context, req *api.XRC721ByContractAddressRequest) (*api.XRC721ByContractAddressResponse, error) {
	resp := &api.XRC721ByContractAddressResponse{
		Count:  0,
		Exist:  false,
		Xrc721: make([]*api.Xrc721Action, 0),
	}
	address, err := common.Address(req.GetAddress())
	if err != nil {
		return nil, err
	}
	count, err := actions.GetXrc721CountByContractAddress(*address)
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
	infos, err := actions.GetXrc721InfoByContractAddress(*address, skip, first)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		resp.Xrc721 = append(resp.Xrc721, &api.Xrc721Action{
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

// XRC721ByPage returns Xrc721 actions by pagination
func (s *XRC721Service) XRC721ByPage(ctx context.Context, req *api.XRC721ByPageRequest) (*api.XRC721ByPageResponse, error) {
	resp := &api.XRC721ByPageResponse{
		Count:  0,
		Exist:  false,
		Xrc721: make([]*api.Xrc721Action, 0),
	}
	count, err := actions.GetXrc721Count()
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
	infos, err := actions.GetXrc721InfoByPage(skip, first)
	if err != nil {
		return nil, err
	}
	for _, info := range infos {
		resp.Xrc721 = append(resp.Xrc721, &api.Xrc721Action{
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

// XRC721Addresses returns Xrc721 contract addresses
func (s *XRC721Service) XRC721Addresses(ctx context.Context, req *api.XRC721AddressesRequest) (*api.XRC721AddressesResponse, error) {
	resp := &api.XRC721AddressesResponse{
		Count:     0,
		Exist:     false,
		Addresses: make([]string, 0),
	}
	count, err := actions.GetXrc721ContractAddressCount()
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
	resp.Addresses, err = actions.GetXrc721ContractAddress(skip, first)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// XRC721TokenHolderAddresses returns Xrc721 token holder addresses given a Xrc721 contract address
func (s *XRC721Service) XRC721TokenHolderAddresses(ctx context.Context, req *api.XRC721TokenHolderAddressesRequest) (*api.XRC721TokenHolderAddressesResponse, error) {
	resp := &api.XRC721TokenHolderAddressesResponse{
		Count:     0,
		Addresses: make([]string, 0),
	}
	tokenAddress := req.GetTokenAddress()
	count, err := actions.GetXrc721TokenHolderCountByTokenAddress(tokenAddress)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return resp, nil
	}
	resp.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	resp.Addresses, err = actions.GetXrc721TokenHolder(tokenAddress, skip, first)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetNFTTransferList returns NFT transfers (xrc721 + xrc1155) with optional contract/address filters.
func (s *XRC721Service) GetNFTTransferList(ctx context.Context, req *api.GetNFTTransferListRequest) (*api.GetNFTTransferListResponse, error) {
	resp := &api.GetNFTTransferListResponse{}
	gormDB := db.DB()
	contract := req.GetContractAddress()
	address := req.GetAddress()

	// Determine address type: isValidAddress ↔ sender/recipient; otherwise token_id filter
	isAddr := len(address) > 0 && (len(address) == 41 || (len(address) > 2 && address[:2] == "0x"))

	tsExpr := `to_char(timestamp AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"')`

	buildUnion := func(limit uint64) string {
		var contractFilter721, contractFilter1155, addrFilter721, addrFilter1155, addrFilterBatch string
		if contract != "" {
			contractFilter721 = fmt.Sprintf(" AND contract_address = '%s'", contract)
			contractFilter1155 = contractFilter721
		}
		if address != "" {
			if isAddr {
				addrFilter721 = fmt.Sprintf(" AND (sender = '%s' OR recipient = '%s')", address, address)
				addrFilter1155 = addrFilter721
				addrFilterBatch = addrFilter721
			} else {
				addrFilter721 = fmt.Sprintf(" AND token_id::text = '%s'", address)
				addrFilter1155 = fmt.Sprintf(" AND _id::text = '%s'", address)
				addrFilterBatch = fmt.Sprintf(" AND convert_from(ids, 'UTF8') = '%s'", address)
			}
		}
		return fmt.Sprintf(`
			(SELECT id, 'xrc721' AS type, block_height, action_hash, contract_address,
			        token_id::text AS token_id, '1' AS value, sender, recipient, %s AS timestamp
			 FROM erc721_transfers_v2_2_3 WHERE 1=1%s%s ORDER BY id DESC LIMIT %d)
			UNION ALL
			(SELECT id, 'xrc1155' AS type, block_height, action_hash, contract_address,
			        _id::text AS token_id, value::text AS value, sender, recipient, %s AS timestamp
			 FROM erc1155_transfer_singles_v2_2_2 WHERE 1=1%s%s ORDER BY id DESC LIMIT %d)
			UNION ALL
			(SELECT id, 'xrc1155' AS type, block_height, action_hash, contract_address,
			        convert_from(ids, 'UTF8') AS token_id, convert_from(values, 'UTF8') AS value,
			        sender, recipient, %s AS timestamp
			 FROM erc1155_transfer_batchs_v2_2_2 WHERE 1=1%s%s ORDER BY id DESC LIMIT %d)`,
			tsExpr, contractFilter721, addrFilter721, limit,
			tsExpr, contractFilter1155, addrFilter1155, limit,
			tsExpr, contractFilter1155, addrFilterBatch, limit,
		)
	}

	// count
	var count int64
	countQ := fmt.Sprintf("SELECT COUNT(*) FROM (%s) t", buildUnion(1000000))
	if err := gormDB.WithContext(ctx).Raw(countQ).Scan(&count).Error; err != nil {
		return nil, errors.Wrap(err, "failed to count NFT transfers")
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
		Type            string
		BlockHeight     uint64
		ActionHash      string
		ContractAddress string
		TokenId         sql.NullString
		Value           sql.NullString
		Sender          string
		Recipient       string
		Timestamp       sql.NullString
	}
	q := fmt.Sprintf("SELECT * FROM (%s) t ORDER BY timestamp DESC LIMIT ? OFFSET ?", buildUnion(queryLimit))
	if err := gormDB.WithContext(ctx).Raw(q, first, skip).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get NFT transfers")
	}
	for _, r := range rows {
		info := &api.NFTTransferInfo{
			Id:              r.Id,
			Type:            r.Type,
			BlockHeight:     r.BlockHeight,
			ActionHash:      r.ActionHash,
			ContractAddress: r.ContractAddress,
			Sender:          r.Sender,
			Recipient:       r.Recipient,
		}
		if r.TokenId.Valid {
			info.TokenId = r.TokenId.String
		}
		if r.Value.Valid {
			info.Value = r.Value.String
		}
		if r.Timestamp.Valid {
			info.Timestamp = r.Timestamp.String
		}
		resp.Transfers = append(resp.Transfers, info)
	}
	return resp, nil
}

// GetNFTHoldersByContract returns NFT holders for a contract from erc1155_721_holders_v2_1_0.
func (s *XRC721Service) GetNFTHoldersByContract(ctx context.Context, req *api.GetNFTHoldersByContractRequest) (*api.GetNFTHoldersByContractResponse, error) {
	resp := &api.GetNFTHoldersByContractResponse{}
	gormDB := db.DB()
	contract := req.GetContractAddress()

	var count int64
	if err := gormDB.WithContext(ctx).Raw(
		`SELECT COUNT(DISTINCT holder) FROM erc1155_721_holders_v2_1_0 WHERE contract_address = ?`, contract,
	).Scan(&count).Error; err != nil {
		return nil, errors.Wrap(err, "failed to count NFT holders")
	}
	resp.Count = count

	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	var rows []struct {
		Address string
		Balance string
	}
	q := `SELECT holder AS address,
               SUM(CASE WHEN erc_type = '721' THEN 1 ELSE token_value END)::text AS balance
          FROM erc1155_721_holders_v2_1_0
          WHERE contract_address = ?
          GROUP BY holder
          ORDER BY balance DESC
          LIMIT ? OFFSET ?`
	if err := gormDB.WithContext(ctx).Raw(q, contract, first, skip).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get NFT holders")
	}
	for _, r := range rows {
		resp.Holders = append(resp.Holders, &api.NFTHolderInfo{
			Address: r.Address,
			Balance: r.Balance,
		})
	}
	return resp, nil
}
