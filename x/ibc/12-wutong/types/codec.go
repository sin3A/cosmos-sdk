package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

const (
	// SubModuleName for the localhost (loopback) client
	SubModuleName = "wutong"
)

// SubModuleCdc defines the IBC localhost client codec.
var SubModuleCdc *codec.Codec

// RegisterCodec registers the localhost types
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(ClientState{}, "ibc/client/wutongchain/ClientState", nil)
	cdc.RegisterConcrete(MsgCreateClient{}, "ibc/client/wutongchain/MsgCreateClient", nil)
	SetSubModuleCodec(cdc)
}

// SetSubModuleCodec sets the ibc localhost client codec
func SetSubModuleCodec(cdc *codec.Codec) {
	SubModuleCdc = cdc
}
