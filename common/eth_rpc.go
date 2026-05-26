package common

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/iotexproject/iotex-address/address"
	"github.com/pkg/errors"
)

var ethRPCClient = &http.Client{Timeout: 10 * time.Second}

type ethRPCRequest struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	ID      int    `json:"id"`
}

type ethRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ethRPCResponse struct {
	Result string       `json:"result"`
	Error  *ethRPCError `json:"error,omitempty"`
}

// EthGetBalanceAtHeight returns the IOTX balance (rau) of an io1.../0x...
// address at the given block height, by calling eth_getBalance against the
// configured eth-archive endpoint. height == 0 means latest. Returns an
// error if the endpoint is empty or the call fails.
func EthGetBalanceAtHeight(ctx context.Context, endpoint, addr string, height uint64) (*big.Int, error) {
	if endpoint == "" {
		return nil, errors.New("eth archive endpoint not configured")
	}
	// The chain gRPC config uses bare host:port (e.g. archive-api.mainnet.iotex.one:443);
	// it's easy to copy that style here, but net/http needs an explicit scheme.
	// Fail loudly with a clear message rather than the cryptic
	// "missing protocol scheme" from http.NewRequest.
	if u, err := url.Parse(endpoint); err != nil || (u.Scheme != "http" && u.Scheme != "https") {
		return nil, errors.Errorf("eth archive endpoint %q must include http:// or https:// scheme", endpoint)
	}
	hexAddr, err := ioToHexAddress(addr)
	if err != nil {
		return nil, err
	}
	blockTag := "latest"
	if height > 0 {
		blockTag = fmt.Sprintf("0x%x", height)
	}
	body, err := json.Marshal(ethRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_getBalance",
		Params:  []any{hexAddr, blockTag},
		ID:      1,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := ethRPCClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "eth_getBalance request failed")
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("eth_getBalance http %d: %s", resp.StatusCode, string(raw))
	}
	var parsed ethRPCResponse
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return nil, errors.Wrap(err, "eth_getBalance decode")
	}
	if parsed.Error != nil {
		return nil, errors.Errorf("eth_getBalance rpc error: %s", parsed.Error.Message)
	}
	hex := strings.TrimPrefix(parsed.Result, "0x")
	if hex == "" {
		return big.NewInt(0), nil
	}
	out, ok := new(big.Int).SetString(hex, 16)
	if !ok {
		return nil, errors.Errorf("eth_getBalance: cannot parse %q", parsed.Result)
	}
	return out, nil
}

// ioToHexAddress accepts io1... or 0x... and returns the 0x... form.
// Special protocol pseudo-addresses (io0000…rewardingprotocol /
// io000…stakingprotocol) are mapped to the hex of their underlying 20-byte
// protocol hash — the chain holds their actual balance under that hex.
func ioToHexAddress(addr string) (string, error) {
	if len(addr) < 3 {
		return "", address.ErrInvalidAddr
	}
	if addr[:2] == "0x" || addr[:2] == "0X" {
		return addr, nil
	}
	switch addr {
	case address.RewardingPoolAddr:
		a, err := address.FromBytes(address.RewardingProtocolAddrHash[:])
		if err != nil {
			return "", err
		}
		return a.Hex(), nil
	case address.StakingBucketPoolAddr:
		a, err := address.FromBytes(address.StakingProtocolAddrHash[:])
		if err != nil {
			return "", err
		}
		return a.Hex(), nil
	}
	a, err := address.FromString(addr)
	if err != nil {
		return "", err
	}
	return a.Hex(), nil
}
