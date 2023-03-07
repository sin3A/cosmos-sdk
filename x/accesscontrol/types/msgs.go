package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	acltypes "github.com/cosmos/cosmos-sdk/types/accesscontrol"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
)

func NewMsgUpdateResourceDependencyMappingProposal(title, description, from string, messageDependencyMapping []acltypes.MessageDependencyMapping) *MsgUpdateResourceDependencyMappingProposal {
	return &MsgUpdateResourceDependencyMappingProposal{title, description, messageDependencyMapping, from}
}

// Route implements Msg
func (p *MsgUpdateResourceDependencyMappingProposal) Route() string {
	return RouterKey
}

// Type implements Msg
func (p *MsgUpdateResourceDependencyMappingProposal) Type() string {
	return "update_params"
}

// ValidateBasic implements Msg
func (p *MsgUpdateResourceDependencyMappingProposal) ValidateBasic() error {
	title := p.Title
	if len(strings.TrimSpace(title)) == 0 {
		return sdkerrors.Wrap(ErrInvalidTitle, "proposal title cannot be blank")
	}
	if len(title) > MaxTitleLength {
		return sdkerrors.Wrapf(ErrTitleLenth, "proposal title is longer than max length of %d", MaxTitleLength)
	}

	description := p.Description
	if len(description) == 0 {
		return sdkerrors.Wrap(ErrInvalidDescription, "proposal description cannot be blank")
	}
	if len(description) > MaxDescriptionLength {
		return sdkerrors.Wrapf(ErrDescriptioneLenth, "proposal description is longer than max length of %d", MaxDescriptionLength)
	}

	return nil
}

// GetSignBytes implements Msg
func (p *MsgUpdateResourceDependencyMappingProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(p)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (p *MsgUpdateResourceDependencyMappingProposal) GetSigners() []sdk.AccAddress {
	singer, err := sdk.AccAddressFromBech32(p.Operator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{singer}
}
