package actions

import (
	"context"

	"github.com/iotexproject/iotex-analyser-api/db"
)

// actionPartitionStep is the block_height span of one block_action_partition
// partition. It mirrors the analyser's block_action_partition `partitionStep`
// (1000000 in every deployment); a mismatch only makes the ranges below prune less
// precisely, it never changes the result.
const actionPartitionStep = 1000000

// blockRange is the block_height range [lo, hi). hi == 0 leaves it unbounded above and
// lo == 0 unbounded below, so the zero value covers everything.
type blockRange struct {
	lo, hi uint64
}

// conditions appends the range to an address filter built by buildAddressFilterConditions.
func (r blockRange) conditions(filterSQL string, filterArgs []interface{}) (string, []interface{}) {
	args := append([]interface{}{}, filterArgs...)
	if r.lo > 0 {
		filterSQL += " AND a.block_height>=?"
		args = append(args, r.lo)
	}
	if r.hi > 0 {
		filterSQL += " AND a.block_height<?"
		args = append(args, r.hi)
	}
	return filterSQL, args
}

// addressBlockRanges splits [0, ∞) into ranges ordered newest first, aligned to
// partition boundaries. The first range is the partition holding tip and everything
// above it (so rows indexed after tip was read are not missed); each older range spans
// twice as many partitions as the previous one, which bounds the number of queries for
// an address whose history is old or sparse.
func addressBlockRanges(tip, step uint64) []blockRange {
	top := tip / step * step
	ranges := []blockRange{{lo: top}}
	for hi, span := top, step; hi > 0; span *= 2 {
		var lo uint64
		if hi > span {
			lo = hi - span
		}
		ranges = append(ranges, blockRange{lo: lo, hi: hi})
		hi = lo
	}
	return ranges
}

// pageAcrossRanges returns rows [skip, skip+first) of the concatenation of the ranges,
// where fetch returns a range's rows newest first. This equals the single-query page
// because ids grow with block_height, so every row of a range is newer than every row
// of the ranges after it. count is only called while rows still need to be skipped,
// letting whole ranges be skipped without fetching them.
func pageAcrossRanges(
	ranges []blockRange,
	skip, first uint64,
	count func(blockRange) (uint64, error),
	fetch func(r blockRange, skip, first uint64) ([]*ActionInfo, error),
) ([]*ActionInfo, error) {
	var out []*ActionInfo
	for _, r := range ranges {
		if first == 0 {
			break
		}
		if skip > 0 {
			n, err := count(r)
			if err != nil {
				return nil, err
			}
			if n <= skip {
				skip -= n
				continue
			}
		}
		rows, err := fetch(r, skip, first)
		if err != nil {
			return nil, err
		}
		out = append(out, rows...)
		skip = 0
		if uint64(len(rows)) >= first {
			break
		}
		first -= uint64(len(rows))
	}
	return out, nil
}

func getLatestBlockHeight(ctx context.Context) (uint64, error) {
	var height uint64
	if err := db.DB().WithContext(ctx).Raw("SELECT COALESCE(max(block_height), 0) FROM block").Scan(&height).Error; err != nil {
		return 0, err
	}
	return height, nil
}
