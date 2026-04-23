package apiservice

import (
	"context"
	"database/sql"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/pkg/errors"
)

type ApprovalService struct {
	api.UnimplementedApprovalServiceServer
}

// GetXRC20Approvals returns the latest ERC20 approvals for each (contract, spender) pair
// for the given owner address (amount > 0).
func (s *ApprovalService) GetXRC20Approvals(ctx context.Context, req *api.GetXRC20ApprovalsRequest) (*api.GetXRC20ApprovalsResponse, error) {
	resp := &api.GetXRC20ApprovalsResponse{}
	gormDB := db.DB()

	var rows []struct {
		ActionHash      string
		ContractAddress string
		Owner           string
		Spender         string
		Amount          sql.NullString
		Timestamp       sql.NullString
	}
	q := `SELECT DISTINCT ON (contract_address, spender)
		action_hash, contract_address, owner, spender, amount::text AS amount,
		to_char(timestamp AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS timestamp
	FROM erc20_approvals
	WHERE owner = ? AND amount > 0
	ORDER BY contract_address, spender, timestamp DESC`

	if err := gormDB.WithContext(ctx).Raw(q, req.GetOwnerAddress()).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get xrc20 approvals")
	}
	for _, r := range rows {
		info := &api.XRC20ApprovalInfo{
			ActionHash:      r.ActionHash,
			ContractAddress: r.ContractAddress,
			Owner:           r.Owner,
			Spender:         r.Spender,
		}
		if r.Amount.Valid {
			info.Amount = r.Amount.String
		}
		if r.Timestamp.Valid {
			info.Timestamp = r.Timestamp.String
		}
		resp.Approvals = append(resp.Approvals, info)
	}
	return resp, nil
}

// GetXRC721Approvals returns the latest ERC721 approvals for each (contract, approved) pair
// for the given owner address (approved != zero address).
func (s *ApprovalService) GetXRC721Approvals(ctx context.Context, req *api.GetXRC721ApprovalsRequest) (*api.GetXRC721ApprovalsResponse, error) {
	resp := &api.GetXRC721ApprovalsResponse{}
	gormDB := db.DB()

	var rows []struct {
		ActionHash      string
		ContractAddress string
		Owner           string
		Approved        string
		TokenID         sql.NullString
		Timestamp       sql.NullString
	}
	q := `SELECT DISTINCT ON (contract_address, approved)
		action_hash, contract_address, owner, approved, token_id::text AS token_id,
		to_char(timestamp AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS"Z"') AS timestamp
	FROM erc721_approvals
	WHERE owner = ? AND approved != '0x0000000000000000000000000000000000000000'
	ORDER BY contract_address, approved, timestamp DESC`

	if err := gormDB.WithContext(ctx).Raw(q, req.GetOwnerAddress()).Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get xrc721 approvals")
	}
	for _, r := range rows {
		info := &api.XRC721ApprovalInfo{
			ActionHash:      r.ActionHash,
			ContractAddress: r.ContractAddress,
			Owner:           r.Owner,
			Approved:        r.Approved,
		}
		if r.TokenID.Valid {
			info.TokenId = r.TokenID.String
		}
		if r.Timestamp.Valid {
			info.Timestamp = r.Timestamp.String
		}
		resp.Approvals = append(resp.Approvals, info)
	}
	return resp, nil
}
