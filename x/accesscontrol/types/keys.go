package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ResourceDependencyMappingKey is the key used for the keeper store
var (
	ResourceDependencyMappingKey = 0x01
	WasmMappingKey               = 0x02
)

const (
	// ModuleName defines the module name
	ModuleName = "accesscontrol"

	QuerierRoute = ModuleName

	// Append "acl" to prevent prefix collision with "acc" module
	StoreKey = "acl" + ModuleName

	RouterKey               = ModuleName
	ProposalTypeText string = "Text"

	MaxDescriptionLength int = 5000
	MaxTitleLength       int = 140
)

var (
	ErrInvalidTitle       = sdkerrors.Register(ModuleName, 1, "title is invalid")
	ErrTitleLenth         = sdkerrors.Register(ModuleName, 2, "title is longer than max length")
	ErrInvalidDescription = sdkerrors.Register(ModuleName, 3, "proposal description cannot be blank")
	ErrDescriptioneLenth  = sdkerrors.Register(ModuleName, 4, "proposal description is longer than max length")
)

func GetResourceDependencyMappingKey() []byte {
	return []byte{byte(ResourceDependencyMappingKey)}
}

func GetResourceDependencyKey(messageKey MessageKey) []byte {
	return append(GetResourceDependencyMappingKey(), []byte(messageKey)...)
}
