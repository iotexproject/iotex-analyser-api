package tasks

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/iotexproject/iotex-analyser-api/common/actions"
	"github.com/iotexproject/iotex-analyser-api/internal/sync/errgroup"
	"github.com/iotexproject/iotex-core/pkg/log"
	"github.com/millken/gocache"
)

func ChainSyncWorker() {
	ticker := time.NewTicker(time.Hour * 6)
	defer ticker.Stop()
	firstTick := false

	tickerChan := func() <-chan time.Time {
		if !firstTick {
			firstTick = true
			c := make(chan time.Time, 1)
			c <- time.Now()
			return c
		}

		return ticker.C
	}
	for {
		select {
		case <-tickerChan():
			if err := chainSync(); err != nil {
				log.S().Errorf("failed to run worker %s", err)
				return
			}
		}
	}
}

type SyncBlock struct {
	MinHeight uint64
	MaxHeight uint64
	TotalSize uint64
}

func chainSync() error {
	minTime, maxTime, err := actions.GetBlockTimes()
	if err != nil {
		return err
	}
	var statsMap sync.Map
	g := new(errgroup.Group)
	for i := minTime.Unix(); i <= maxTime.Unix()-86400; i += 86400 {
		i := i
		g.Go(func(ctx context.Context) error {
			minBlkHeight, maxBlkHeight, blkSize, err := actions.GetBlockStatsByDate(i)
			if err != nil {
				return err
			}
			statsMap.Store(i, &SyncBlock{MinHeight: minBlkHeight, MaxHeight: maxBlkHeight, TotalSize: blkSize})
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}
	syncBlks := []*SyncBlock{}
	statsMap.Range(func(key, value interface{}) bool {
		syncBlks = append(syncBlks, value.(*SyncBlock))
		return true
	})
	sort.Slice(syncBlks, func(i, j int) bool {
		return syncBlks[i].MinHeight < syncBlks[j].MinHeight
	})
	for i := 1; i < len(syncBlks)-1; i++ {
		syncBlks[i].TotalSize = syncBlks[i].TotalSize + syncBlks[i-1].TotalSize
	}
	gocache.Set("syncBlks", syncBlks, 0)
	return nil
}
