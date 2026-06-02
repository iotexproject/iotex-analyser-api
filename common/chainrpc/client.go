// Package chainrpc proxies a small set of iotex-core gRPC calls used by the
// analyser-api when the underlying value is no longer kept in the database.
//
// The first consumer is action_execution.data: once the analyser only stores
// the first 4 bytes of calldata, ActionByHash's input_data field must be
// fetched from the chain. A process-wide LRU absorbs duplicate lookups from
// repeated page loads of the same tx in iotexscan.
package chainrpc

import (
	"context"
	"errors"
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/iotexproject/iotex-analyser-api/common"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"google.golang.org/grpc"
)

const (
	cacheSize  = 10_000
	rpcTimeout = 5 * time.Second
)

// ErrNotInitialized is returned by GetActionData when the package was never
// wired up — e.g. in unit tests or local dev with no CHAIN_GRPC_ENDPOINT set.
// Callers should degrade gracefully (return empty input_data) rather than fail
// the whole RPC.
var ErrNotInitialized = errors.New("chainrpc: not initialized")

var (
	mu     sync.RWMutex
	conn   *grpc.ClientConn
	client iotexapi.APIServiceClient
	cache  *lru.Cache[string, []byte]
)

// Init dials the iotex-core gRPC endpoint and prepares the LRU. Safe to call
// once at startup. An empty endpoint disables the package; later GetActionData
// calls will return ErrNotInitialized so the caller can degrade.
func Init(endpoint string) error {
	mu.Lock()
	defer mu.Unlock()
	if endpoint == "" {
		return nil
	}
	c, err := common.NewDefaultGRPCConn(endpoint)
	if err != nil {
		return err
	}
	lc, err := lru.New[string, []byte](cacheSize)
	if err != nil {
		_ = c.Close()
		return err
	}
	conn = c
	client = iotexapi.NewAPIServiceClient(c)
	cache = lc
	return nil
}

// Close releases the underlying gRPC connection. Safe to call at shutdown.
func Close() error {
	mu.Lock()
	defer mu.Unlock()
	if conn == nil {
		return nil
	}
	c := conn
	conn = nil
	client = nil
	cache = nil
	return c.Close()
}

// GetActionData returns the raw calldata bytes of an Execution action by its
// hex hash (no 0x prefix, matching the format stored in action_execution).
//
// Non-Execution actions and unknown hashes return (nil, nil) — both are
// cached so a request storm against a non-Execution hash doesn't repeatedly
// hit the chain.
func GetActionData(ctx context.Context, actionHash string) ([]byte, error) {
	mu.RLock()
	c := client
	lc := cache
	mu.RUnlock()
	if c == nil || lc == nil {
		return nil, ErrNotInitialized
	}
	if v, ok := lc.Get(actionHash); ok {
		return v, nil
	}

	rpcCtx, cancel := context.WithTimeout(ctx, rpcTimeout)
	defer cancel()
	resp, err := c.GetActions(rpcCtx, &iotexapi.GetActionsRequest{
		Lookup: &iotexapi.GetActionsRequest_ByHash{
			ByHash: &iotexapi.GetActionByHashRequest{
				ActionHash:   actionHash,
				CheckPending: false,
			},
		},
	})
	if err != nil {
		// errors are not cached: a transient outage on api.iotex.one
		// shouldn't poison the entry for the LRU's lifetime.
		return nil, err
	}

	var data []byte
	if len(resp.GetActionInfo()) > 0 {
		data = resp.GetActionInfo()[0].GetAction().GetCore().GetExecution().GetData()
	}
	lc.Add(actionHash, data)
	return data, nil
}
