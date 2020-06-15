package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateHTLC{}, "irismod/htlc/MsgCreateHTLC", nil)
	cdc.RegisterConcrete(MsgClaimHTLC{}, "irismod/htlc/MsgClaimHTLC", nil)
	cdc.RegisterConcrete(MsgRefundHTLC{}, "irismod/htlc/MsgRefundHTLC", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
