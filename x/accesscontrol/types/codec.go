package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(
		&MsgUpdateResourceDependencyMappingProposal{},
		"cosmos-sdk/MsgUpdateResourceDependencyMappingProposal",
		nil,
	)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgUpdateResourceDependencyMappingProposal{},
	)
}

var ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
