package types

import (
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHTLC constructs a new HTLC instance
func NewHTLC(
	sender sdk.AccAddress,
	to sdk.AccAddress,
	receiverOnOtherChain string,
	amount sdk.Coins,
	secret tmbytes.HexBytes,
	timestamp uint64,
	expirationHeight uint64,
	state HTLCState,
) HTLC {
	return HTLC{
		Sender:               sender,
		To:                   to,
		ReceiverOnOtherChain: receiverOnOtherChain,
		Amount:               amount,
		Secret:               secret,
		Timestamp:            timestamp,
		ExpirationHeight:     expirationHeight,
		State:                state,
	}
}

// Validate validates the HTLC
func (h HTLC) Validate() error {
	// TODO
	return nil
}

// GetHashLock calculates the hash lock from the given secret and timestamp
func GetHashLock(secret tmbytes.HexBytes, timestamp uint64) []byte {
	if timestamp > 0 {
		return tmhash.Sum(append(secret, sdk.Uint64ToBigEndian(timestamp)...))
	}

	return tmhash.Sum(secret)
}
