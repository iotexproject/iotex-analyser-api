package chainrpc

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"
	"google.golang.org/grpc"
)

// fakeClient embeds the interface so we only need to implement the method we
// actually call. Any unintended call to another method nil-derefs in the test,
// which is a feature: it surfaces accidental new RPC calls.
type fakeClient struct {
	iotexapi.APIServiceClient
	calls atomic.Int32
	fn    func(context.Context, *iotexapi.GetActionsRequest) (*iotexapi.GetActionsResponse, error)
}

func (f *fakeClient) GetActions(ctx context.Context, in *iotexapi.GetActionsRequest, _ ...grpc.CallOption) (*iotexapi.GetActionsResponse, error) {
	f.calls.Add(1)
	return f.fn(ctx, in)
}

func setupTest(t *testing.T, fn func(context.Context, *iotexapi.GetActionsRequest) (*iotexapi.GetActionsResponse, error)) *fakeClient {
	t.Helper()
	lc, err := lru.New[string, []byte](16)
	if err != nil {
		t.Fatalf("lru.New: %v", err)
	}
	fc := &fakeClient{fn: fn}
	mu.Lock()
	client = fc
	cache = lc
	mu.Unlock()
	t.Cleanup(func() {
		mu.Lock()
		client = nil
		cache = nil
		mu.Unlock()
	})
	return fc
}

func TestGetActionData_NotInitialized(t *testing.T) {
	mu.Lock()
	client = nil
	cache = nil
	mu.Unlock()
	_, err := GetActionData(context.Background(), "deadbeef")
	if !errors.Is(err, ErrNotInitialized) {
		t.Fatalf("expected ErrNotInitialized, got %v", err)
	}
}

func TestGetActionData_CacheMissThenHit(t *testing.T) {
	want := []byte{0xa9, 0x05, 0x9c, 0xbb, 0x01, 0x02}
	fc := setupTest(t, func(_ context.Context, _ *iotexapi.GetActionsRequest) (*iotexapi.GetActionsResponse, error) {
		return &iotexapi.GetActionsResponse{
			ActionInfo: []*iotexapi.ActionInfo{
				{Action: &iotextypes.Action{Core: &iotextypes.ActionCore{
					Action: &iotextypes.ActionCore_Execution{Execution: &iotextypes.Execution{Data: want}},
				}}},
			},
		}, nil
	})

	got, err := GetActionData(context.Background(), "hash1")
	if err != nil || string(got) != string(want) {
		t.Fatalf("first call: got=%x err=%v", got, err)
	}
	got2, err := GetActionData(context.Background(), "hash1")
	if err != nil || string(got2) != string(want) {
		t.Fatalf("second call: got=%x err=%v", got2, err)
	}
	if c := fc.calls.Load(); c != 1 {
		t.Fatalf("expected 1 RPC call, got %d", c)
	}
}

func TestGetActionData_NonExecutionCached(t *testing.T) {
	fc := setupTest(t, func(_ context.Context, _ *iotexapi.GetActionsRequest) (*iotexapi.GetActionsResponse, error) {
		// Non-Execution action: Core.Action is something else (or nil here).
		return &iotexapi.GetActionsResponse{
			ActionInfo: []*iotexapi.ActionInfo{
				{Action: &iotextypes.Action{Core: &iotextypes.ActionCore{}}},
			},
		}, nil
	})
	got, err := GetActionData(context.Background(), "transferHash")
	if err != nil {
		t.Fatalf("err=%v", err)
	}
	if len(got) != 0 {
		t.Fatalf("expected empty data, got %x", got)
	}
	// Second call should hit cache.
	_, _ = GetActionData(context.Background(), "transferHash")
	if c := fc.calls.Load(); c != 1 {
		t.Fatalf("expected empty result to be cached; got %d RPC calls", c)
	}
}

func TestGetActionData_RPCErrorNotCached(t *testing.T) {
	var stage atomic.Int32
	fc := setupTest(t, func(_ context.Context, _ *iotexapi.GetActionsRequest) (*iotexapi.GetActionsResponse, error) {
		if stage.Load() == 0 {
			return nil, errors.New("upstream down")
		}
		return &iotexapi.GetActionsResponse{
			ActionInfo: []*iotexapi.ActionInfo{
				{Action: &iotextypes.Action{Core: &iotextypes.ActionCore{
					Action: &iotextypes.ActionCore_Execution{Execution: &iotextypes.Execution{Data: []byte{0xde, 0xad}}},
				}}},
			},
		}, nil
	})

	if _, err := GetActionData(context.Background(), "flaky"); err == nil {
		t.Fatalf("expected error on first call")
	}
	stage.Store(1)
	got, err := GetActionData(context.Background(), "flaky")
	if err != nil || string(got) != string([]byte{0xde, 0xad}) {
		t.Fatalf("second call after recovery: got=%x err=%v", got, err)
	}
	if c := fc.calls.Load(); c != 2 {
		t.Fatalf("expected 2 RPC calls (error not cached), got %d", c)
	}
}

func TestGetActionData_EmptyResponseInfoCached(t *testing.T) {
	fc := setupTest(t, func(_ context.Context, _ *iotexapi.GetActionsRequest) (*iotexapi.GetActionsResponse, error) {
		return &iotexapi.GetActionsResponse{ActionInfo: nil}, nil
	})
	_, _ = GetActionData(context.Background(), "missing")
	_, _ = GetActionData(context.Background(), "missing")
	if c := fc.calls.Load(); c != 1 {
		t.Fatalf("expected 1 RPC call for missing hash, got %d", c)
	}
}
