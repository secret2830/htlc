package htlc

import (
	"fmt"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker handles block beginning logic for HTLC
func BeginBlocker(ctx sdk.Context, k Keeper) {
	ctx = ctx.WithLogger(ctx.Logger().With("handler", "beginBlock").With("module", "irismod/htlc"))

	currentBlockHeight := uint64(ctx.BlockHeight())

	k.IterateHTLCExpiredQueueByHeight(
		ctx,
		currentBlockHeight,
		func(hlock tmbytes.HexBytes, h HTLC) (stop bool) {
			// update the state
			h.State = Expired
			k.SetHTLC(ctx, h, hlock)

			// delete from the expiration queue
			k.DeleteHTLCFromExpiredQueue(ctx, currentBlockHeight, hlock)

			ctx.EventManager().EmitEvents(sdk.Events{
				sdk.NewEvent(
					EventTypeHTLCExpired,
					sdk.NewAttribute(AttributeKeyHashLock, hlock.String()),
				),
			})

			ctx.Logger().Info(fmt.Sprintf("HTLC [%s] is expired", hlock.String()))

			return false
		},
	)
}
