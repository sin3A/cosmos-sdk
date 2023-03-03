package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(
		&MsgUpdateResourceDependencyMappingProposal{},
		"cosmos-sdk/MsgUpdateResourceDependencyMappingProposal",
		nil,
	)
	cdc.RegisterConcrete(
		&MsgUpdateWasmDependencyMappingProposal{},
		"cosmos-sdk/MsgUpdateWasmDependencyMappingProposal",
		nil,
	)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*govtypes.Content)(nil),
		&MsgUpdateResourceDependencyMappingProposal{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterWasmDependency{},
		&govtypes.MsgSubmitProposal{},
		&govtypes.MsgVote{},
		&govtypes.MsgVoteWeighted{},
		&govtypes.MsgDeposit{},
	)

	registry.RegisterInterface(
		"cosmos.gov.v1beta1.Content",
		(*govtypes.Content)(nil),
		&govtypes.TextProposal{},
	)
}

var ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
