package apiservice

import (
	"context"
	"log"
	"math/big"
	"sort"

	"github.com/iotexproject/iotex-analyser-api/api"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-analyser-api/common/rewards"
	"github.com/iotexproject/iotex-analyser-api/db"
	"github.com/iotexproject/iotex-analyser-api/model"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// VoterRewardService serves the IIP-59 on-chain voter reward tables written by
// the analyser's voter_reward plugin.
type VoterRewardService struct {
	api.UnimplementedVoterRewardServiceServer
}

const (
	// maxUnifiedEpochSpan bounds the epoch window UnifiedVoterRewards will
	// scan. IoTeX runs 24 epochs a day, so this is roughly eleven years —
	// beyond any real query, and small enough that start_epoch+epoch_count
	// cannot overflow.
	maxUnifiedEpochSpan = 100000

	// maxUnifiedFetch bounds what one UnifiedVoterRewards call pulls from each
	// leg when the caller asks for an unbounded page. common.PageSize returns 0
	// for a pagination message that omits `first`, and gorm renders Limit(0) as
	// no LIMIT at all, so without this a single request would materialise the
	// voter's entire history from both sources.
	maxUnifiedFetch = 10000
)

// unifiedFetchLimit is how many rows one leg must supply for the requested page
// to be correct.
//
// The merge is a sort, not a filter, so a row that sits beyond position
// offset+limit within its own source cannot land inside the page once the two
// sources are interleaved — anything ahead of it in its own leg is also ahead
// of it after the merge. Fetching that many from each leg is therefore exact,
// not an approximation.
func unifiedFetchLimit(offset, limit int) int {
	if limit <= 0 || offset < 0 || offset > maxUnifiedFetch-limit {
		return maxUnifiedFetch
	}
	if n := offset + limit; n < maxUnifiedFetch {
		return n
	}
	return maxUnifiedFetch
}

// unifiedOnchainTotals aggregates the on-chain leg over the whole epoch window,
// independent of the page.
func unifiedOnchainTotals(
	ctx context.Context, voter string, startEpoch, endEpoch uint64,
) (uint64, string, error) {
	base := filterQuery(ctx).Where(
		"d.voter_address = ? AND d.epoch_number >= ? AND d.epoch_number <= ?",
		voter, startEpoch, endEpoch)

	var count int64
	if err := base.Session(&gorm.Session{}).Count(&count).Error; err != nil {
		return 0, "", errors.Wrap(err, "count on-chain voter rewards")
	}
	total, err := sumAmounts(base)
	if err != nil {
		return 0, "", errors.Wrap(err, "sum on-chain voter rewards")
	}
	return uint64(count), total, nil
}

// paymentRow is the join shape shared by the by-voter and by-delegate queries:
// the distribution row plus the timestamp of the block that paid it.
type paymentRow struct {
	model.VoterRewardDistribution
	Timestamp int64
}

// filterQuery is the WHERE-only half of a payment query. Aggregates run
// against it directly; only the row fetch adds the block join and the
// projection. Reusing one pre-projected query for both is what silently
// corrupted the totals here first time round — Pluck and Count each rewrite
// the SELECT list, and a query that already carries one does not survive that
// cleanly.
func filterQuery(ctx context.Context) *gorm.DB {
	return db.DB().WithContext(ctx).Table("voter_reward_distribution AS d")
}

// withPaymentRows adds the block join and projection needed to materialise
// paymentRow values.
func withPaymentRows(q *gorm.DB) *gorm.DB {
	return q.
		// block.timestamp is a SQL timestamp, not an integer, so it has to be
		// converted before it can be defaulted or returned as a unix second.
		Select("d.*, COALESCE(EXTRACT(EPOCH FROM b.timestamp), 0)::bigint AS timestamp").
		Joins("LEFT JOIN block AS b ON b.block_height = d.block_height")
}

func applyEraRange(q *gorm.DB, startEra, endEra uint64) *gorm.DB {
	if startEra > 0 {
		q = q.Where("d.era >= ?", startEra)
	}
	if endEra > 0 {
		q = q.Where("d.era <= ?", endEra)
	}
	return q
}

func toPayments(rows []paymentRow) []*api.VoterRewardPayment {
	out := make([]*api.VoterRewardPayment, 0, len(rows))
	for _, r := range rows {
		out = append(out, &api.VoterRewardPayment{
			Era:              r.Era,
			BlockHeight:      r.BlockHeight,
			EpochNumber:      r.EpochNumber,
			DelegateId:       r.DelegateID,
			DelegateName:     r.DelegateName,
			VoterAddress:     r.VoterAddress,
			RecipientAddress: r.RecipientAddress,
			Amount:           r.Amount,
			Compounded:       r.Compounded,
			CompoundBucketId: r.CompoundBucketID,
			ActionHash:       r.ActionHash,
		})
	}
	return out
}

// sumAmounts adds decimal strings without going through float. Amounts are
// Rau, so they routinely exceed 2^53 and a float64 sum would quietly lose the
// low digits of every total this API reports.
func sumAmounts(q *gorm.DB) (string, error) {
	var total *string
	// SUM over numeric(60,0) stays exact in Postgres; casting to text keeps it
	// exact on the way into Go. Routing it through float64 would silently drop
	// the low digits of every Rau total this API reports.
	if err := q.Session(&gorm.Session{}).
		Select("CAST(COALESCE(SUM(d.amount), 0) AS TEXT)").
		Scan(&total).Error; err != nil {
		return "0", err
	}
	if total == nil || *total == "" {
		return "0", nil
	}
	if _, ok := new(big.Int).SetString(*total, 10); !ok {
		return "0", errors.Errorf("unparseable amount total %q", *total)
	}
	return *total, nil
}

func (s *VoterRewardService) VoterRewardsByVoter(
	ctx context.Context, req *api.VoterRewardsByVoterRequest,
) (*api.VoterRewardsByVoterResponse, error) {
	voter := req.GetVoterAddress()
	if voter == "" {
		return nil, status.Error(codes.InvalidArgument, "voter_address is required")
	}
	base := applyEraRange(filterQuery(ctx).Where("d.voter_address = ?", voter),
		req.GetStartEra(), req.GetEndEra())

	var count int64
	if err := base.Session(&gorm.Session{}).Count(&count).Error; err != nil {
		return nil, errors.Wrap(err, "count voter reward payments")
	}
	resp := &api.VoterRewardsByVoterResponse{Count: uint64(count), TotalAmount: "0"}
	if count == 0 {
		return resp, nil
	}
	resp.Exist = true

	total, err := sumAmounts(base)
	if err != nil {
		return nil, errors.Wrap(err, "sum voter reward payments")
	}
	resp.TotalAmount = total

	var rows []paymentRow
	if err := withPaymentRows(base.Session(&gorm.Session{})).
		Order("d.block_height DESC, d.log_index DESC, d.row_index DESC").
		Offset(int(common.PageOffset(req.GetPagination()))).
		Limit(int(common.PageSize(req.GetPagination()))).
		Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "query voter reward payments")
	}
	resp.Payments = toPayments(rows)
	return resp, nil
}

func (s *VoterRewardService) VoterRewardsByDelegate(
	ctx context.Context, req *api.VoterRewardsByDelegateRequest,
) (*api.VoterRewardsByDelegateResponse, error) {
	base := filterQuery(ctx)
	switch {
	case req.GetDelegateId() != "":
		base = base.Where("d.delegate_id = ?", req.GetDelegateId())
	case req.GetDelegateName() != "":
		base = base.Where("d.delegate_name = ?", req.GetDelegateName())
	default:
		return nil, status.Error(codes.InvalidArgument, "delegate_id or delegate_name is required")
	}
	base = applyEraRange(base, req.GetStartEra(), req.GetEndEra())

	var count int64
	if err := base.Session(&gorm.Session{}).Count(&count).Error; err != nil {
		return nil, errors.Wrap(err, "count delegate reward payments")
	}
	resp := &api.VoterRewardsByDelegateResponse{Count: uint64(count), TotalAmount: "0"}
	if count == 0 {
		return resp, nil
	}
	resp.Exist = true

	total, err := sumAmounts(base)
	if err != nil {
		return nil, errors.Wrap(err, "sum delegate reward payments")
	}
	resp.TotalAmount = total

	var voters int64
	if err := base.Session(&gorm.Session{}).
		Distinct("d.voter_address").Count(&voters).Error; err != nil {
		return nil, errors.Wrap(err, "count distinct voters")
	}
	resp.VoterCount = uint64(voters)

	var rows []paymentRow
	if err := withPaymentRows(base.Session(&gorm.Session{})).
		Order("d.block_height DESC, d.log_index DESC, d.row_index DESC").
		Offset(int(common.PageOffset(req.GetPagination()))).
		Limit(int(common.PageSize(req.GetPagination()))).
		Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "query delegate reward payments")
	}
	resp.Payments = toPayments(rows)
	return resp, nil
}

func (s *VoterRewardService) VoterRewardEra(
	ctx context.Context, req *api.VoterRewardEraRequest,
) (*api.VoterRewardEraResponse, error) {
	q := db.DB().WithContext(ctx).Model(&model.VoterRewardEra{})
	if era := req.GetEra(); era > 0 {
		q = q.Where("era = ?", era)
	} else {
		q = q.Order("era DESC")
	}
	var row model.VoterRewardEra
	if err := q.Take(&row).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &api.VoterRewardEraResponse{}, nil
		}
		return nil, errors.Wrap(err, "query voter reward era")
	}

	resp := &api.VoterRewardEraResponse{
		Exist:            true,
		Era:              row.Era,
		FreezeHeight:     row.FreezeHeight,
		FirstChunkAt:     row.FirstChunkAt,
		CompletedHeight:  row.CompletedHeight,
		Completed:        row.CompletedHeight > 0,
		ScanPhase:        row.ScanPhase,
		ResumeVoter:      row.ResumeVoter,
		DelegateCount:    uint64(row.DelegateCount),
		VoterCount:       row.VoterCount,
		TotalFrozen:      row.TotalFrozen,
		TotalDistributed: row.TotalDistributed,
		OverrunResidue:   row.OverrunResidue,
		LastChunkAt:      row.LastChunkAt,
	}

	// Era numbers step by epochsPerRewardEra, so neighbours have to be looked
	// up rather than derived by ±1.
	var prev, next *uint64
	if err := db.DB().WithContext(ctx).Model(&model.VoterRewardEra{}).
		Select("MAX(era)").Where("era < ?", row.Era).Scan(&prev).Error; err != nil {
		return nil, errors.Wrap(err, "query previous era")
	}
	if err := db.DB().WithContext(ctx).Model(&model.VoterRewardEra{}).
		Select("MIN(era)").Where("era > ?", row.Era).Scan(&next).Error; err != nil {
		return nil, errors.Wrap(err, "query next era")
	}
	if prev != nil {
		resp.PrevEra = *prev
	}
	if next != nil {
		resp.NextEra = *next
	}

	// Per-delegate allocation: what each delegate froze against what it has
	// actually paid so far in this era.
	var configs []model.DelegateRewardConfig
	if err := db.DB().WithContext(ctx).
		Where("era = ?", row.Era).Find(&configs).Error; err != nil {
		return nil, errors.Wrap(err, "query delegate reward config")
	}
	paid, err := distributedByDelegate(ctx, row.Era)
	if err != nil {
		return nil, err
	}
	for _, c := range configs {
		if !c.OptedIn {
			// A legacy delegate has no allocation in this settlement at all;
			// listing it with zeros would read as "paid nothing this era"
			// rather than "not part of this pipeline".
			continue
		}
		resp.Allocations = append(resp.Allocations, &api.EraDelegateAllocation{
			DelegateId:           c.DelegateID,
			DelegateName:         c.DelegateName,
			AmountFrozen:         amountOrZero(c.VoterAmountFrozen),
			AmountDistributed:    amountOrZero(paid[c.DelegateID]),
			TotalWeight:          c.TotalWeight,
			OptedIn:              c.OptedIn,
			CommissionConfigured: c.CommissionConfigured,
			BlockCommissionBps:   c.BlockCommissionBps,
			EpochCommissionBps:   c.EpochCommissionBps,
		})
	}
	sort.Slice(resp.Allocations, func(i, j int) bool {
		return resp.Allocations[i].DelegateName < resp.Allocations[j].DelegateName
	})
	return resp, nil
}

func distributedByDelegate(ctx context.Context, era uint64) (map[string]string, error) {
	var rows []struct {
		DelegateID string
		Amount     string
	}
	// SUM over a decimal column stays exact in Postgres; casting to text keeps
	// it that way on the way out.
	if err := db.DB().WithContext(ctx).
		Table("voter_reward_distribution").
		Select("delegate_id, CAST(SUM(amount) AS TEXT) AS amount").
		Where("era = ?", era).
		Group("delegate_id").
		Scan(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "sum distributed by delegate")
	}
	out := make(map[string]string, len(rows))
	for _, r := range rows {
		out[r.DelegateID] = r.Amount
	}
	return out, nil
}

func amountOrZero(s string) string {
	if s == "" {
		return "0"
	}
	return s
}

func (s *VoterRewardService) DelegateRewardConfig(
	ctx context.Context, req *api.DelegateRewardConfigRequest,
) (*api.DelegateRewardConfigResponse, error) {
	era := req.GetEra()
	if era == 0 {
		if err := db.DB().WithContext(ctx).
			Model(&model.DelegateRewardConfig{}).
			Select("COALESCE(MAX(era), 0)").Scan(&era).Error; err != nil {
			return nil, errors.Wrap(err, "resolve latest config era")
		}
	}
	q := db.DB().WithContext(ctx).Where("era = ?", era)
	if ids := req.GetDelegateIds(); len(ids) > 0 {
		q = q.Where("delegate_id IN ?", ids)
	}
	var rows []model.DelegateRewardConfig
	if err := q.Order("delegate_name").Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "query delegate reward config")
	}
	resp := &api.DelegateRewardConfigResponse{}
	for _, r := range rows {
		resp.Delegates = append(resp.Delegates, &api.DelegateRewardConfigEntry{
			DelegateId:           r.DelegateID,
			DelegateName:         r.DelegateName,
			Era:                  r.Era,
			OptedIn:              r.OptedIn,
			OptInSource:          r.OptInSource,
			OptInHeight:          r.OptInHeight,
			FreezeHeight:         r.FreezeHeight,
			BlockCommissionBps:   r.BlockCommissionBps,
			EpochCommissionBps:   r.EpochCommissionBps,
			CommissionConfigured: r.CommissionConfigured,
			TotalWeight:          r.TotalWeight,
			PayoutAddress:        r.PayoutAddress,
			VoterAmountFrozen:    amountOrZero(r.VoterAmountFrozen),
		})
	}
	return resp, nil
}

func (s *VoterRewardService) VoterRewardDestination(
	ctx context.Context, req *api.VoterRewardDestinationRequest,
) (*api.VoterRewardDestinationResponse, error) {
	voters := req.GetVoterAddresses()
	if len(voters) == 0 {
		return nil, status.Error(codes.InvalidArgument, "voter_addresses is required")
	}
	var rows []model.VoterRewardDestination
	if err := db.DB().WithContext(ctx).
		Where("voter_address IN ?", voters).
		Order("block_height ASC, id ASC").
		Find(&rows).Error; err != nil {
		return nil, errors.Wrap(err, "query voter reward destinations")
	}
	// Last change wins; the table keeps history so a payout can be read
	// against the destination in force at its height.
	latest := make(map[string]model.VoterRewardDestination, len(rows))
	for _, r := range rows {
		latest[r.VoterAddress] = r
	}

	resp := &api.VoterRewardDestinationResponse{}
	for _, v := range voters {
		entry := &api.VoterRewardDestinationEntry{
			VoterAddress: v,
			// With no override the protocol pays the voter directly, so
			// reporting the voter's own address is the true answer rather
			// than a placeholder.
			Recipient: v,
		}
		if r, ok := latest[v]; ok && r.NewRecipient != "" {
			entry.Recipient = r.NewRecipient
			entry.ExplicitlySet = true
			entry.UpdatedHeight = r.BlockHeight
		}
		resp.Destinations = append(resp.Destinations, entry)
	}
	return resp, nil
}

// UnifiedVoterRewards merges the on-chain and Hermes pipelines behind one
// shape.
//
// This exists so the branch does not get written twice in the frontends. A
// delegate that opts in stops producing Hermes bookkeeping entirely, so a UI
// that queries only Hermes shows a voter an empty table and implies their
// rewards stopped, when in fact only the pipeline changed. Deciding which
// source to read per delegate is a property of chain state, not of the
// caller, so it belongs here.
func (s *VoterRewardService) UnifiedVoterRewards(
	ctx context.Context, req *api.UnifiedVoterRewardsRequest,
) (*api.UnifiedVoterRewardsResponse, error) {
	voter := req.GetVoterAddress()
	if voter == "" {
		return nil, status.Error(codes.InvalidArgument, "voter_address is required")
	}
	startEpoch := req.GetStartEpoch()
	epochCount := req.GetEpochCount()
	if epochCount == 0 {
		epochCount = 1
	}
	// Reject rather than clamp. startEpoch+epochCount-1 wraps for a large
	// epochCount, and a wrapped endEpoch turns into `epoch >= start AND
	// epoch <= something-smaller` — an empty result that reads to the caller as
	// "you earned nothing" rather than "your range was nonsense".
	if epochCount > maxUnifiedEpochSpan {
		return nil, status.Errorf(codes.InvalidArgument,
			"epoch_count %d exceeds the maximum span of %d", epochCount, maxUnifiedEpochSpan)
	}
	endEpoch := startEpoch + epochCount - 1
	if endEpoch < startEpoch {
		return nil, status.Error(codes.InvalidArgument, "start_epoch + epoch_count overflows")
	}

	// Both legs are fetched, merged, then sliced in memory, because there is no
	// single ordered index across the two sources. That makes the fetch bound
	// the caller's cost, so it has to be a bound: fetching everything and
	// throwing away all but one page would also defeat the LIMIT that lets
	// Postgres stop early on the Hermes join's `ORDER BY t1.id DESC`.
	//
	// Each leg is capped at offset+limit, since the merge cannot promote a row
	// past position offset+limit in the combined ordering if it was already
	// beyond that position within its own source.
	offset := int(common.PageOffset(req.GetPagination()))
	limit := int(common.PageSize(req.GetPagination()))
	fetchLimit := unifiedFetchLimit(offset, limit)

	rowsOut := make([]*api.UnifiedRewardRow, 0, fetchLimit)
	seenSources := map[api.RewardSource]bool{}

	// On-chain leg. Filtered by the epoch that paid, which is the one unit
	// both pipelines share — Hermes has no notion of an era.
	var onchain []paymentRow
	if err := withPaymentRows(filterQuery(ctx)).
		Where("d.voter_address = ? AND d.epoch_number >= ? AND d.epoch_number <= ?",
			voter, startEpoch, endEpoch).
		Order("d.block_height DESC").
		Limit(fetchLimit).
		Scan(&onchain).Error; err != nil {
		return nil, errors.Wrap(err, "query on-chain voter rewards")
	}
	for _, r := range onchain {
		rowsOut = append(rowsOut, &api.UnifiedRewardRow{
			Epoch:        r.EpochNumber,
			Era:          r.Era,
			DelegateName: r.DelegateName,
			DelegateId:   r.DelegateID,
			Amount:       r.Amount,
			Source:       api.RewardSource_ONCHAIN_IIP59,
			ActionHash:   r.ActionHash,
			Timestamp:    uint64(r.Timestamp),
			Compounded:   r.Compounded,
			BucketId:     r.CompoundBucketID,
		})
		seenSources[api.RewardSource_ONCHAIN_IIP59] = true
	}

	// Hermes leg, reusing the existing bookkeeping query so the two views can
	// never disagree about what Hermes paid.
	//
	// A failure here degrades rather than propagates. The whole reason this
	// endpoint exists is that a voter must never be shown an empty table when
	// one pipeline is simply unavailable — failing the request would reproduce
	// that same "your rewards vanished" impression, just with an error page.
	var hermesUnavailable bool
	hermesRows, err := rewards.GetHermesByVoter(ctx, startEpoch, endEpoch, voter, 0, uint64(fetchLimit))
	if err != nil {
		hermesUnavailable = true
		hermesRows = nil
		log.Printf("UnifiedVoterRewards: hermes leg unavailable for %s: %v", voter, err)
	}
	for _, h := range hermesRows {
		rowsOut = append(rowsOut, &api.UnifiedRewardRow{
			Epoch:        h.EndEpoch,
			DelegateName: h.DelegateName,
			Amount:       h.Amount,
			Source:       api.RewardSource_HERMES_OFFCHAIN,
			ActionHash:   h.ActionHash,
			Timestamp:    uint64(h.Timestamp.Unix()),
		})
		seenSources[api.RewardSource_HERMES_OFFCHAIN] = true
	}

	sort.SliceStable(rowsOut, func(i, j int) bool {
		if rowsOut[i].Epoch != rowsOut[j].Epoch {
			return rowsOut[i].Epoch > rowsOut[j].Epoch
		}
		return rowsOut[i].Timestamp > rowsOut[j].Timestamp
	})

	// count and total_amount describe the whole match, not the page and not the
	// fetch window — the same contract VoterRewardsByVoter offers. They have to
	// be aggregated in SQL now that the row fetch is capped; deriving them from
	// rowsOut would silently report the cap as the voter's lifetime earnings.
	onchainCount, onchainTotal, err := unifiedOnchainTotals(ctx, voter, startEpoch, endEpoch)
	if err != nil {
		return nil, err
	}
	total, ok := new(big.Int).SetString(onchainTotal, 10)
	if !ok {
		return nil, errors.Errorf("unparseable on-chain amount total %q", onchainTotal)
	}
	count := onchainCount

	if !hermesUnavailable {
		hCount, hTotal, err := rewards.GetTotalHermesByVoter(ctx, startEpoch, endEpoch, voter)
		if err != nil {
			// The row fetch already succeeded, so the leg is reachable; only the
			// aggregate failed. Fall back to marking it unavailable rather than
			// reporting a total that omits Hermes without saying so.
			hermesUnavailable = true
			log.Printf("UnifiedVoterRewards: hermes totals unavailable for %s: %v", voter, err)
		} else {
			count += uint64(hCount)
			if v, ok := new(big.Int).SetString(hTotal, 10); ok {
				total.Add(total, v)
			} else if hTotal != "" {
				return nil, errors.Errorf("unparseable hermes amount total %q", hTotal)
			}
		}
	}

	resp := &api.UnifiedVoterRewardsResponse{
		Exist:             count > 0,
		Count:             count,
		TotalAmount:       total.String(),
		HermesUnavailable: hermesUnavailable,
	}
	// Paginate after the merge: the two legs are separate queries and there is
	// no single ordered index across both. offset/limit were resolved before
	// the fetch so each leg could be capped to what this page can possibly need.
	if offset > len(rowsOut) {
		offset = len(rowsOut)
	}
	end := offset + limit
	if limit <= 0 || end > len(rowsOut) {
		end = len(rowsOut)
	}
	resp.Rewards = rowsOut[offset:end]
	for src := range seenSources {
		resp.Sources = append(resp.Sources, src)
	}
	sort.Slice(resp.Sources, func(i, j int) bool { return resp.Sources[i] < resp.Sources[j] })
	return resp, nil
}
