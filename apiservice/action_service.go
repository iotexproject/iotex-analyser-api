package apiservice

import (
	"context"
	"database/sql"
	"encoding/hex"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/actions"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/pkg/errors"
)

type ActionService struct {
	api.UnimplementedActionServiceServer
}

func (s *ActionService) ActionByVoter(ctx context.Context, req *api.ActionRequest) (*api.ActionResponse, error) {
	resp := &api.ActionResponse{
		Count:      0,
		Exist:      false,
		ActionList: make([]*api.ActionInfo, 0),
	}
	address, err := common.Address(req.GetAddress())
	if err != nil {
		return nil, err
	}
	bucketIDs, err := actions.GetBucketIDsByVoter(*address)
	if err != nil {
		return nil, err
	}
	if len(bucketIDs) == 0 {
		return resp, nil
	}
	count, err := actions.GetBucketActionCountByBuckets(bucketIDs)
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
	actionInfoList, err := actions.GetBucketActionInfoByBuckets(bucketIDs, skip, first)
	if err != nil {
		return nil, err
	}
	for _, actionInfo := range actionInfoList {
		resp.ActionList = append(resp.ActionList, &api.ActionInfo{
			ActHash:   actionInfo.ActHash,
			BlkHash:   actionInfo.BlkHash,
			Timestamp: uint64(actionInfo.Timestamp.Unix()),
			ActType:   actionInfo.ActType,
			Sender:    actionInfo.Sender,
			Recipient: actionInfo.Recipient,
			Amount:    actionInfo.Amount,
			GasFee:    actionInfo.GasFee,
			BlkHeight: actionInfo.BlkHeight,
		})
	}
	return resp, nil
}

func (s *ActionService) GetXrc20ByAddress(ctx context.Context, req *api.ActionRequest) (*api.ActionResponse, error) {
	resp := &api.ActionResponse{
		Count:   0,
		Exist:   false,
		XrcList: make([]*api.XrcInfo, 0),
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
		resp.XrcList = append(resp.XrcList, &api.XrcInfo{
			ActHash:   info.ActionHash,
			BlkHeight: info.BlockHeight,
			Timestamp: uint64(info.Timestamp.Unix()),
			From:      info.Sender,
			To:        info.Recipient,
			Quantity:  info.Amount,
			Contract:  info.ContractAddress,
		})
	}
	return resp, nil
}

// ActionByDates finds actions by dates
func (s *ActionService) ActionByDates(ctx context.Context, req *api.ActionByDatesRequest) (*api.ActionByDatesResponse, error) {
	resp := &api.ActionByDatesResponse{}
	startDate := req.GetStartDate()
	endDate := req.GetEndDate()
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())

	count, err := actions.GetActionCountByDates(startDate, endDate)
	if err != nil {
		return nil, err
	}
	resp.Count = uint64(count)
	if count == 0 {
		return resp, nil
	}
	resp.Exist = true
	actionInfoList, err := actions.GetActionInfoByDates(startDate, endDate, skip, first)
	if err != nil {
		return nil, err
	}
	for _, actionInfo := range actionInfoList {
		resp.Actions = append(resp.Actions, &api.ActionInfo{
			ActHash:   actionInfo.ActHash,
			BlkHash:   actionInfo.BlkHash,
			Timestamp: uint64(actionInfo.Timestamp.Unix()),
			ActType:   actionInfo.ActType,
			Sender:    actionInfo.Sender,
			Recipient: actionInfo.Recipient,
			Amount:    actionInfo.Amount,
			GasFee:    actionInfo.GasFee,
			BlkHeight: actionInfo.BlkHeight,
		})
	}
	return resp, nil
}

// ActionByHash finds actions by hash
func (s *ActionService) ActionByHash(ctx context.Context, req *api.ActionByHashRequest) (*api.ActionByHashResponse, error) {
	resp := &api.ActionByHashResponse{}
	actHash := req.GetActHash()

	actionInfo, err := actions.GetActionInfoByHash(actHash)
	if err != nil {
		return nil, err
	}
	resp.ActionInfo = &api.ActionInfo{
		ActHash:            actionInfo.ActHash,
		BlkHash:            actionInfo.BlkHash,
		Timestamp:          uint64(actionInfo.Timestamp.Unix()),
		ActType:            actionInfo.ActType,
		Sender:             actionInfo.Sender,
		Recipient:          actionInfo.Recipient,
		Amount:             actionInfo.Amount,
		GasFee:             actionInfo.GasFee,
		BlkHeight:          actionInfo.BlkHeight,
		GasPrice:           actionInfo.GasPrice,
		GasLimit:           actionInfo.GasLimit,
		GasConsumed:        actionInfo.GasConsumed,
		Nonce:              actionInfo.Nonce,
		Status:             actionInfo.Status,
		ContractAddress:    actionInfo.ContractAddress,
		ExecutionRevertMsg: actionInfo.ExecutionRevertMsg,
		ChainId:            actionInfo.ChainId,
		MethodName:         actionInfo.MethodName,
	}
	resp.Exist = actionInfo.ActHash != ""

	brt, err := actions.GetBlockReceiptTransactionByHash(actHash)
	if err != nil {
		return nil, err
	}
	for _, receipt := range brt {
		resp.EvmTransfers = append(resp.EvmTransfers, &api.ActionByHashResponse_EvmTransfers{
			Sender:    receipt.Sender,
			Recipient: receipt.Recipient,
			Amount:    receipt.Amount,
		})
	}

	gormDB := db.DB()

	// EIP-1559 fields from action_type table
	var actionTypeRow struct {
		Type        sql.NullString
		AccessList  sql.NullString
		GasTipCap   sql.NullString
		GasFeeCap   sql.NullString
		BlobGas     sql.NullString
		BlobFeeCap  sql.NullString
		BlobHashes  sql.NullString
		BlobGasPrice sql.NullString
	}
	if err := gormDB.WithContext(ctx).Raw(
		"SELECT type, access_list, gas_tip_cap, gas_fee_cap, blob_gas, blob_fee_cap, blob_hashes, blob_gas_price FROM action_type WHERE hash = ?",
		actHash,
	).Scan(&actionTypeRow).Error; err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrap(err, "failed to get action type")
	}
	if actionTypeRow.Type.Valid || actionTypeRow.GasTipCap.Valid {
		resp.ActionTypeInfo = &api.ActionByHashResponse_ActionTypeInfo{}
		if actionTypeRow.Type.Valid {
			resp.ActionTypeInfo.Type = actionTypeRow.Type.String
		}
		if actionTypeRow.AccessList.Valid {
			resp.ActionTypeInfo.AccessList = actionTypeRow.AccessList.String
		}
		if actionTypeRow.GasTipCap.Valid {
			resp.ActionTypeInfo.GasTipCap = actionTypeRow.GasTipCap.String
		}
		if actionTypeRow.GasFeeCap.Valid {
			resp.ActionTypeInfo.GasFeeCap = actionTypeRow.GasFeeCap.String
		}
		if actionTypeRow.BlobGas.Valid {
			resp.ActionTypeInfo.BlobGas = actionTypeRow.BlobGas.String
		}
		if actionTypeRow.BlobFeeCap.Valid {
			resp.ActionTypeInfo.BlobFeeCap = actionTypeRow.BlobFeeCap.String
		}
		if actionTypeRow.BlobHashes.Valid {
			resp.ActionTypeInfo.BlobHashes = actionTypeRow.BlobHashes.String
		}
		if actionTypeRow.BlobGasPrice.Valid {
			resp.ActionTypeInfo.BlobGasPrice = actionTypeRow.BlobGasPrice.String
		}
	}

	// Input data from action_execution or block_action
	var inputDataRow struct {
		Data []byte
	}
	if err := gormDB.WithContext(ctx).Raw(
		`SELECT COALESCE(ae.data, ba.payload) as data
		FROM action_execution ae
		RIGHT JOIN block_action ba ON ae.action_hash = ba.action_hash
		WHERE ba.action_hash = ? LIMIT 1`,
		actHash,
	).Scan(&inputDataRow).Error; err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrap(err, "failed to get input data")
	}
	if len(inputDataRow.Data) > 0 {
		resp.InputData = hex.EncodeToString(inputDataRow.Data)
	}

	// Logs from block_receipt_logs
	var logRows []struct {
		BlockHeight uint64
		Address     string
		Topic0      sql.NullString
		Topic1      sql.NullString
		Topic2      sql.NullString
		Topic3      sql.NullString
		Data        []byte
		ActionHash  string
		Index       int64
	}
	if err := gormDB.WithContext(ctx).Raw(
		"SELECT block_height, address, topic0, topic1, topic2, topic3, data, action_hash, index FROM block_receipt_logs WHERE action_hash = ?",
		actHash,
	).Scan(&logRows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get logs")
	}
	for _, r := range logRows {
		logEntry := &api.ActionByHashResponse_ActionLog{
			BlockHeight: r.BlockHeight,
			Address:     r.Address,
			ActionHash:  r.ActionHash,
			Index:       r.Index,
			Data:        r.Data,
		}
		if r.Topic0.Valid {
			logEntry.Topic0 = r.Topic0.String
		}
		if r.Topic1.Valid {
			logEntry.Topic1 = r.Topic1.String
		}
		if r.Topic2.Valid {
			logEntry.Topic2 = r.Topic2.String
		}
		if r.Topic3.Valid {
			logEntry.Topic3 = r.Topic3.String
		}
		resp.Logs = append(resp.Logs, logEntry)
	}

	// Token transfers: erc20 + erc721 union
	var transferRows []struct {
		ID              int64
		ContractAddress string
		Sender          string
		Recipient       string
		Amount          string
		Type            string
	}
	if err := gormDB.WithContext(ctx).Raw(
		`SELECT id, contract_address, sender, recipient, amount, 'erc20' as type FROM erc20_transfers WHERE action_hash = ?
		UNION ALL
		SELECT id, contract_address, sender, recipient, token_id as amount, 'nft' as type FROM erc721_transfers_v2_2_3 WHERE action_hash = ?`,
		actHash, actHash,
	).Scan(&transferRows).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get token transfers")
	}
	for _, r := range transferRows {
		resp.TokenTransfers = append(resp.TokenTransfers, &api.ActionByHashResponse_TokenTransfer{
			Id:              r.ID,
			ContractAddress: r.ContractAddress,
			Sender:          r.Sender,
			Recipient:       r.Recipient,
			Amount:          r.Amount,
			Type:            r.Type,
		})
	}

	// Block base fee from block_meta
	var baseFeeRow struct {
		BlockBaseFee sql.NullString
	}
	if err := gormDB.WithContext(ctx).Raw(
		`SELECT base_fee as block_base_fee FROM block_meta
		WHERE block_height = (SELECT block_height FROM block_action WHERE action_hash = ? LIMIT 1)`,
		actHash,
	).Scan(&baseFeeRow).Error; err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrap(err, "failed to get block base fee")
	}
	if baseFeeRow.BlockBaseFee.Valid {
		resp.BlockBaseFee = baseFeeRow.BlockBaseFee.String
	}

	// Stake action from staking_buckets
	var stakeRow struct {
		BucketID     sql.NullInt64
		Amount       sql.NullString
		StakedAmount sql.NullString
		Duration     sql.NullString
		AutoStake    sql.NullBool
		Candidate    sql.NullString
		ActType      sql.NullString
		OwnerAddress sql.NullString
	}
	if err := gormDB.WithContext(ctx).Raw(
		"SELECT bucket_id, amount, staked_amount, duration, auto_stake, candidate, act_type, owner_address FROM staking_buckets WHERE action_hash = ? ORDER BY id DESC LIMIT 1",
		actHash,
	).Scan(&stakeRow).Error; err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrap(err, "failed to get stake action")
	}
	if stakeRow.ActType.Valid || stakeRow.BucketID.Valid {
		resp.StakeAction = &api.ActionByHashResponse_StakeAction{}
		if stakeRow.BucketID.Valid {
			resp.StakeAction.BucketId = stakeRow.BucketID.Int64
		}
		if stakeRow.Amount.Valid {
			resp.StakeAction.Amount = stakeRow.Amount.String
		}
		if stakeRow.StakedAmount.Valid {
			resp.StakeAction.StakedAmount = stakeRow.StakedAmount.String
		}
		if stakeRow.Duration.Valid {
			resp.StakeAction.Duration = stakeRow.Duration.String
		}
		if stakeRow.AutoStake.Valid {
			resp.StakeAction.AutoStake = stakeRow.AutoStake.Bool
		}
		if stakeRow.Candidate.Valid {
			resp.StakeAction.Candidate = stakeRow.Candidate.String
		}
		if stakeRow.ActType.Valid {
			resp.StakeAction.ActType = stakeRow.ActType.String
		}
		if stakeRow.OwnerAddress.Valid {
			resp.StakeAction.OwnerAddress = stakeRow.OwnerAddress.String
		}
	}

	return resp, nil
}

// ActionByAddress finds actions by address
func (s *ActionService) ActionByAddress(ctx context.Context, req *api.ActionByAddressRequest) (*api.ActionByAddressResponse, error) {
	resp := &api.ActionByAddressResponse{
		Count:   0,
		Exist:   false,
		Actions: make([]*api.ActionInfo, 0),
	}
	address, err := common.Address(req.GetAddress())
	if err != nil {
		return nil, err
	}
	sender := req.GetSender()
	recipient := req.GetRecipient()
	actionType := req.GetActionType()
	startTime := req.GetStartTime()
	endTime := req.GetEndTime()

	count, err := actions.GetActionCountByAddress(ctx, *address, sender, recipient, actionType, startTime, endTime)
	if err != nil {
		return nil, err
	}
	resp.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	actionInfoList, err := actions.GetActionInfoByAddress(ctx, *address, skip, first, sender, recipient, actionType, startTime, endTime)
	if err != nil {
		return nil, err
	}
	for _, actionInfo := range actionInfoList {
		resp.Actions = append(resp.Actions, &api.ActionInfo{
			ActHash:         actionInfo.ActHash,
			BlkHash:         actionInfo.BlkHash,
			Timestamp:       uint64(actionInfo.Timestamp.Unix()),
			ActType:         actionInfo.ActType,
			Sender:          actionInfo.Sender,
			Recipient:       actionInfo.Recipient,
			Amount:          actionInfo.Amount,
			GasFee:          actionInfo.GasFee,
			BlkHeight:       actionInfo.BlkHeight,
			GasPrice:        actionInfo.GasPrice,
			GasConsumed:     actionInfo.GasConsumed,
			Nonce:           actionInfo.Nonce,
			Status:          actionInfo.Status,
			ContractAddress: actionInfo.ContractAddress,
			MethodName:      actionInfo.MethodName,
		})
	}
	return resp, nil
}

// ActionByType finds actions by type
func (s *ActionService) ActionByType(ctx context.Context, req *api.ActionByTypeRequest) (*api.ActionByTypeResponse, error) {
	resp := &api.ActionByTypeResponse{
		Count:   0,
		Exist:   false,
		Actions: make([]*api.ActionInfo, 0),
	}
	typ := req.GetType()

	count, err := actions.GetActionCountByType(ctx, typ)
	if err != nil {
		return nil, err
	}
	resp.Count = uint64(count)
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	actionInfoList, err := actions.GetActionInfoByType(ctx, typ, skip, first)
	if err != nil {
		return nil, err
	}
	for _, actionInfo := range actionInfoList {
		resp.Actions = append(resp.Actions, &api.ActionInfo{
			ActHash:   actionInfo.ActHash,
			BlkHash:   actionInfo.BlkHash,
			Timestamp: uint64(actionInfo.Timestamp.Unix()),
			ActType:   actionInfo.ActType,
			Sender:    actionInfo.Sender,
			Recipient: actionInfo.Recipient,
			Amount:    actionInfo.Amount,
			GasFee:    actionInfo.GasFee,
			BlkHeight: actionInfo.BlkHeight,
		})
	}
	return resp, nil
}

// EvmTransfersByAddress finds EVM transfers by address
func (s *ActionService) EvmTransfersByAddress(ctx context.Context, req *api.EvmTransfersByAddressRequest) (*api.EvmTransfersByAddressResponse, error) {
	resp := &api.EvmTransfersByAddressResponse{
		Count:        0,
		Exist:        false,
		EvmTransfers: make([]*api.EvmTransfersByAddressResponse_EvmTransfer, 0),
	}
	address, err := common.Address(req.GetAddress())
	if err != nil {
		return nil, err
	}
	count, err := actions.GetEvmTransferCount(ctx, *address)
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
	actionInfos, err := actions.GetEvmTransferInfoByAddress(ctx, *address, skip, first)
	if err != nil {
		return nil, err
	}
	for _, info := range actionInfos {
		resp.EvmTransfers = append(resp.EvmTransfers, &api.EvmTransfersByAddressResponse_EvmTransfer{
			ActHash:   info.ActHash,
			BlkHash:   info.BlkHash,
			Timestamp: uint64(info.Timestamp.Unix()),
			Sender:    info.Sender,
			Recipient: info.Recipient,
			Amount:    info.Amount,
			BlkHeight: info.BlkHeight,
		})
	}
	return resp, nil
}

// ActionList returns paginated list of latest actions
func (s *ActionService) ActionList(ctx context.Context, req *api.ActionListRequest) (*api.ActionListResponse, error) {
	resp := &api.ActionListResponse{
		Count:   0,
		Exist:   false,
		Actions: make([]*api.ActionInfo, 0),
	}

	startBh := req.GetStartBlockHeight()
	var count int64
	var err error
	if startBh > 0 {
		count, err = actions.GetActionCountFromHeight(startBh)
	} else {
		count, err = actions.GetActionCount()
	}
	if err != nil {
		return nil, err
	}
	resp.Count = uint64(count)
	if count == 0 {
		return resp, nil
	}
	resp.Exist = true
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	var actionInfoList []*actions.ActionInfo
	if startBh > 0 {
		actionInfoList, err = actions.GetActionInfoListFromHeight(startBh, skip, first)
	} else {
		actionInfoList, err = actions.GetActionInfoList(skip, first)
	}
	if err != nil {
		return nil, err
	}
	for _, actionInfo := range actionInfoList {
		resp.Actions = append(resp.Actions, &api.ActionInfo{
			ActHash:         actionInfo.ActHash,
			BlkHash:         actionInfo.BlkHash,
			Timestamp:       uint64(actionInfo.Timestamp.Unix()),
			ActType:         actionInfo.ActType,
			Sender:          actionInfo.Sender,
			Recipient:       actionInfo.Recipient,
			Amount:          actionInfo.Amount,
			GasFee:          actionInfo.GasFee,
			BlkHeight:       actionInfo.BlkHeight,
			GasPrice:        actionInfo.GasPrice,
			GasLimit:        actionInfo.GasLimit,
			GasConsumed:     actionInfo.GasConsumed,
			Nonce:           actionInfo.Nonce,
			Status:          actionInfo.Status,
			ContractAddress: actionInfo.ContractAddress,
			MethodName:      actionInfo.MethodName,
		})
	}
	return resp, nil
}

// ActionByHeight finds actions by block height
func (s *ActionService) ActionByHeight(ctx context.Context, req *api.ActionByHeightRequest) (*api.ActionByHeightResponse, error) {
	resp := &api.ActionByHeightResponse{
		Count:   0,
		Exist:   false,
		Actions: make([]*api.ActionInfo, 0),
	}
	height := req.GetHeight()
	count, err := actions.GetActionCountByHeight(height)
	if err != nil {
		return nil, err
	}
	resp.Count = uint64(count)
	if count == 0 {
		return resp, nil
	}
	resp.Exist = true
	skip := common.PageOffset(req.GetPagination())
	first := common.PageSize(req.GetPagination())
	actionInfoList, err := actions.GetActionInfoByHeight(height, skip, first)
	if err != nil {
		return nil, err
	}
	for _, actionInfo := range actionInfoList {
		resp.Actions = append(resp.Actions, &api.ActionInfo{
			ActHash:         actionInfo.ActHash,
			BlkHash:         actionInfo.BlkHash,
			Timestamp:       uint64(actionInfo.Timestamp.Unix()),
			ActType:         actionInfo.ActType,
			Sender:          actionInfo.Sender,
			Recipient:       actionInfo.Recipient,
			Amount:          actionInfo.Amount,
			GasFee:          actionInfo.GasFee,
			BlkHeight:       actionInfo.BlkHeight,
			GasPrice:        actionInfo.GasPrice,
			GasLimit:        actionInfo.GasLimit,
			GasConsumed:     actionInfo.GasConsumed,
			Nonce:           actionInfo.Nonce,
			Status:          actionInfo.Status,
			ContractAddress: actionInfo.ContractAddress,
			MethodName:      actionInfo.MethodName,
		})
	}
	return resp, nil
}

// ContractInteractors returns distinct senders who interacted with a contract
func (s *ActionService) ContractInteractors(ctx context.Context, req *api.ContractInteractorsRequest) (*api.ContractInteractorsResponse, error) {
	resp := &api.ContractInteractorsResponse{
		Senders: make([]string, 0),
	}
	address := req.GetAddress()
	startTime := req.GetStartTime()
	senders, err := actions.GetContractInteractors(address, startTime)
	if err != nil {
		return nil, err
	}
	resp.Senders = senders
	return resp, nil
}
