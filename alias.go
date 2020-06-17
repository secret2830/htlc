package htlc

// nolint

import (
	"github.com/irismod/htlc/keeper"
	"github.com/irismod/htlc/types"
)

const (
	ModuleName                       = types.ModuleName
	StoreKey                         = types.StoreKey
	QuerierRoute                     = types.QuerierRoute
	RouterKey                        = types.RouterKey
	HTLCAccName                      = types.HTLCAccName
	QueryHTLC                        = types.QueryHTLC
	EventTypeCreateHTLC              = types.EventTypeCreateHTLC
	EventTypeClaimHTLC               = types.EventTypeClaimHTLC
	EventTypeRefundHTLC              = types.EventTypeRefundHTLC
	EventTypeHTLCExpired             = types.EventTypeHTLCExpired
	AttributeValueCategory           = types.AttributeValueCategory
	AttributeKeySender               = types.AttributeKeySender
	AttributeKeyReceiver             = types.AttributeKeyReceiver
	AttributeKeyReceiverOnOtherChain = types.AttributeKeyReceiverOnOtherChain
	AttributeKeyAmount               = types.AttributeKeyAmount
	AttributeKeyHashLock             = types.AttributeKeyHashLock
	AttributeKeyTimeLock             = types.AttributeKeyTimeLock
	AttributeKeySecret               = types.AttributeKeySecret
	Open                             = types.Open
	Completed                        = types.Completed
	Refunded                         = types.Refunded
	Expired                          = types.Expired
)

var (
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	ModuleCdc           = types.ModuleCdc
	RegisterCodec       = types.RegisterCodec
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	NewGenesisState     = types.NewGenesisState
	NewHTLC             = types.NewHTLC
)

type (
	Keeper          = keeper.Keeper
	HTLC            = types.HTLC
	GenesisState    = types.GenesisState
	MsgCreateHTLC   = types.MsgCreateHTLC
	MsgClaimHTLC    = types.MsgClaimHTLC
	MsgRefundHTLC   = types.MsgRefundHTLC
	QueryHTLCParams = types.QueryHTLCParams
)
