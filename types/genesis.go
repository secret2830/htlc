package types

import (
	"encoding/hex"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// GenesisState contains all HTLC state that must be provided at genesis
type GenesisState struct {
	PendingHTLCs map[string]HTLC `json:"pending_htlcs" yaml:"pending_htlcs"` // claimable HTLCs
}

// NewGenesisState constructs a new GenesisState instance
func NewGenesisState(
	pendingHTLCs map[string]HTLC,
) GenesisState {
	return GenesisState{
		PendingHTLCs: pendingHTLCs,
	}
}

// DefaultGenesisState gets the raw genesis message for testing
func DefaultGenesisState() GenesisState {
	return GenesisState{
		PendingHTLCs: map[string]HTLC{},
	}
}

// ValidateGenesis validates the provided HTLC genesis state to ensure the
// expected invariants holds.
func ValidateGenesis(data GenesisState) error {
	for hashLockStr, htlc := range data.PendingHTLCs {
		if _, err := hex.DecodeString(hashLockStr); err != nil {
			return sdkerrors.Wrap(ErrInvalidHashLock, hashLockStr)
		}

		if err := htlc.Validate(); err != nil {
			return err
		}
	}

	return nil
}
