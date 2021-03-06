package fuzz

import (
	"github.com/prysmaticlabs/go-ssz"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/blocks"
	stateTrie "github.com/prysmaticlabs/prysm/beacon-chain/state"
	prylabs_testing "github.com/prysmaticlabs/prysm/fuzz/testing"
	"github.com/prysmaticlabs/prysm/shared/params"
)

// BeaconFuzzBlockHeader using the corpora from sigp/beacon-fuzz.
func BeaconFuzzBlockHeader(b []byte) ([]byte, bool) {
	params.UseMainnetConfig()
	input := &InputBlockHeader{}
	if err := ssz.Unmarshal(b, input); err != nil {
		return fail(err)
	}
	s, err := prylabs_testing.GetBeaconFuzzState(input.StateID)
	if err != nil || s == nil {
		return nil, false
	}
	st, err := stateTrie.InitializeFromProto(s)
	if err != nil {
		return fail(err)
	}
	post, err := blocks.ProcessBlockHeaderNoVerify(st, input.Block)
	if err != nil {
		return fail(err)
	}
	return success(post)
}
