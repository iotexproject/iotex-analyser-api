package actions

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddressBlockRanges(t *testing.T) {
	require := require.New(t)
	require.Equal([]blockRange{{}}, addressBlockRanges(0, 1000000))
	require.Equal([]blockRange{{}}, addressBlockRanges(999999, 1000000))
	require.Equal([]blockRange{{lo: 1000000}, {lo: 0, hi: 1000000}}, addressBlockRanges(1000000, 1000000))
	require.Equal([]blockRange{
		{lo: 53000000},
		{lo: 52000000, hi: 53000000},
		{lo: 50000000, hi: 52000000},
		{lo: 46000000, hi: 50000000},
		{lo: 38000000, hi: 46000000},
		{lo: 22000000, hi: 38000000},
		{lo: 0, hi: 22000000},
	}, addressBlockRanges(53816905, 1000000))
}

func TestBlockRangeConditions(t *testing.T) {
	require := require.New(t)
	filterArgs := []interface{}{"execution"}

	sql, args := blockRange{}.conditions(" AND a.action_type=?", filterArgs)
	require.Equal(" AND a.action_type=?", sql)
	require.Equal([]interface{}{"execution"}, args)

	sql, args = blockRange{lo: 52000000, hi: 53000000}.conditions(" AND a.action_type=?", filterArgs)
	require.Equal(" AND a.action_type=? AND a.block_height>=? AND a.block_height<?", sql)
	require.Equal([]interface{}{"execution", uint64(52000000), uint64(53000000)}, args)
	require.Equal([]interface{}{"execution"}, filterArgs, "caller's args must not be modified")
}

// fakeRanges serves count/fetch from in-memory rows, mimicking the SQL: rows inside the
// range, newest first, with offset/limit.
type fakeRanges struct {
	rows           []*ActionInfo // newest first
	counts, fetchs []blockRange
}

func (f *fakeRanges) inRange(r blockRange) []*ActionInfo {
	var out []*ActionInfo
	for _, row := range f.rows {
		if row.BlkHeight >= r.lo && (r.hi == 0 || row.BlkHeight < r.hi) {
			out = append(out, row)
		}
	}
	return out
}

func (f *fakeRanges) count(r blockRange) (uint64, error) {
	f.counts = append(f.counts, r)
	return uint64(len(f.inRange(r))), nil
}

func (f *fakeRanges) fetch(r blockRange, skip, first uint64) ([]*ActionInfo, error) {
	f.fetchs = append(f.fetchs, r)
	rows := f.inRange(r)
	if skip >= uint64(len(rows)) {
		return nil, nil
	}
	rows = rows[skip:]
	if first < uint64(len(rows)) {
		rows = rows[:first]
	}
	return rows, nil
}

// newFakeRanges builds rows at the given heights; ids (ActHash) grow with height.
func newFakeRanges(heights []uint64) *fakeRanges {
	sort.Slice(heights, func(i, j int) bool { return heights[i] > heights[j] })
	f := &fakeRanges{}
	for i, h := range heights {
		f.rows = append(f.rows, &ActionInfo{BlkHeight: h, ActHash: fmt.Sprintf("%08d", len(heights)-i)})
	}
	return f
}

func expectedPage(rows []*ActionInfo, skip, first uint64) []*ActionInfo {
	if skip >= uint64(len(rows)) || first == 0 {
		return nil
	}
	rows = rows[skip:]
	if first < uint64(len(rows)) {
		rows = rows[:first]
	}
	return rows
}

func TestPageAcrossRangesMatchesSingleQuery(t *testing.T) {
	const step = 100
	rnd := rand.New(rand.NewSource(1))
	for _, tc := range []struct {
		name    string
		tip     uint64
		heights func() []uint64
	}{
		{"empty", 5400, func() []uint64 { return nil }},
		{"busy every block", 5400, func() []uint64 {
			var hs []uint64
			for h := uint64(0); h <= 5400; h++ {
				hs = append(hs, h, h) // also a duplicate row per block, like self-transfers
			}
			return hs
		}},
		{"dormant since early history", 5400, func() []uint64 {
			var hs []uint64
			for i := 0; i < 300; i++ {
				hs = append(hs, uint64(rnd.Intn(700)))
			}
			return hs
		}},
		{"sparse random", 5400, func() []uint64 {
			var hs []uint64
			for i := 0; i < 50; i++ {
				hs = append(hs, uint64(rnd.Intn(5401)))
			}
			return hs
		}},
		{"rows above tip", 5400, func() []uint64 { return []uint64{5450, 5401, 5399, 5300, 10} }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newFakeRanges(tc.heights())
			ranges := addressBlockRanges(tc.tip, step)
			for _, skip := range []uint64{0, 1, 3, 19, 20, 21, 99, 100, 250, 1000, 10799, 10800, 20000} {
				for _, first := range []uint64{0, 1, 20, 1000} {
					got, err := pageAcrossRanges(ranges, skip, first, f.count, f.fetch)
					require.NoError(t, err)
					require.Equal(t, expectedPage(f.rows, skip, first), got, "skip=%d first=%d", skip, first)
				}
			}
		})
	}
}

func TestPageAcrossRangesStopsEarly(t *testing.T) {
	require := require.New(t)
	var heights []uint64
	for h := uint64(0); h <= 5450; h++ {
		heights = append(heights, h)
	}
	ranges := addressBlockRanges(5450, 100) // newest range [5400, ∞) holds 51 rows

	// First page of a busy address: one fetch on the newest range, no count.
	f := newFakeRanges(heights)
	got, err := pageAcrossRanges(ranges, 0, 20, f.count, f.fetch)
	require.NoError(err)
	require.Len(got, 20)
	require.Empty(f.counts)
	require.Equal([]blockRange{{lo: 5400}}, f.fetchs)

	// A skip past the newest range counts it but does not fetch it.
	f = newFakeRanges(heights)
	got, err = pageAcrossRanges(ranges, 60, 20, f.count, f.fetch)
	require.NoError(err)
	require.Len(got, 20)
	require.Equal([]blockRange{{lo: 5400}, {lo: 5300, hi: 5400}}, f.counts)
	require.Equal([]blockRange{{lo: 5300, hi: 5400}}, f.fetchs)
}

func TestPageAcrossRangesPropagatesErrors(t *testing.T) {
	boom := fmt.Errorf("boom")
	ranges := addressBlockRanges(5400, 100)
	ok := func(blockRange) (uint64, error) { return 0, nil }
	_, err := pageAcrossRanges(ranges, 1, 20, func(blockRange) (uint64, error) { return 0, boom }, nil)
	require.ErrorIs(t, err, boom)
	_, err = pageAcrossRanges(ranges, 0, 20, ok, func(blockRange, uint64, uint64) ([]*ActionInfo, error) { return nil, boom })
	require.ErrorIs(t, err, boom)
}
