package apiservice

import "testing"

// The two vaults iotex-core lists in genesis. Verified against mainnet chain
// data 2026-08-31: 44 and 15 candidates respectively use them as their reward
// address, and every other candidate has one of its own.
var testVaults = map[string]bool{
	"io19604a05s2p3mecam2zz7d27hcr6ndyw80wvkmh": true,
	"io12mgttmfa2ffn9uqvn0yn37f4nz43d248l2ga85": true,
}

// TestClassifyRewardRouting pins the three-way split.
//
// The bug this guards against is collapsing it to two: before this endpoint the
// explorer rendered only "on-chain" and "Hermes", so every delegate that
// distributes rewards on its own terms was labelled Hermes -- a claim about who
// pays their voters that is simply false.
func TestClassifyRewardRouting(t *testing.T) {
	hermesVault := "io19604a05s2p3mecam2zz7d27hcr6ndyw80wvkmh"
	ownAddr := "io1k7rqg4ksjx93f467mdfjzm6un6ezskwcppls80"

	for _, tc := range []struct {
		name   string
		optIn  bool
		reward string
		want   string
	}{
		{"opted in, own reward address", true, ownAddr, rewardRoutingOnchain},
		{"not opted in, reward address is a Hermes vault", false, hermesVault, rewardRoutingHermes},
		{"not opted in, own reward address", false, ownAddr, rewardRoutingSelf},
		// The auto-migration case: iotex-core opts a delegate in at IIP-59
		// activation *because* its reward address was a Hermes vault, so both
		// signals are set and only the bit is current. Reading the address
		// first would report Hermes for a delegate the protocol now pays.
		{"opted in while still pointing at a vault", true, hermesVault, rewardRoutingOnchain},
		// An empty reward address is not a vault, so it must not fall into the
		// Hermes bucket by accident.
		{"not opted in, empty reward address", false, "", rewardRoutingSelf},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := classifyRewardRouting(tc.optIn, tc.reward, testVaults); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

// A nil vault set must degrade to "self" rather than panic: it is what a
// deployment with the config key unset looks like.
func TestClassifyRewardRoutingWithoutVaults(t *testing.T) {
	if got := classifyRewardRouting(false, "io1anything", nil); got != rewardRoutingSelf {
		t.Fatalf("got %q, want %q", got, rewardRoutingSelf)
	}
	if got := classifyRewardRouting(true, "io1anything", nil); got != rewardRoutingOnchain {
		t.Fatalf("got %q, want %q", got, rewardRoutingOnchain)
	}
}
