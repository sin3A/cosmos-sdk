package accesscontrol

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/accesscontrol/keeper"
	"github.com/cosmos/cosmos-sdk/x/accesscontrol/types"
)

func HandleMsgUpdateResourceDependencyMappingProposal(ctx sdk.Context, k *keeper.Keeper, p *types.MsgUpdateResourceDependencyMappingProposal) (*sdk.Result, error) {
	for _, resourceDepMapping := range p.MessageDependencyMapping {
		k.SetResourceDependencyMapping(ctx, resourceDepMapping)
	}
	return nil, nil
}

func NewProposalHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch c := msg.(type) {
		case *types.MsgUpdateResourceDependencyMappingProposal:
			return HandleMsgUpdateResourceDependencyMappingProposal(ctx, &k, c)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized accesscontrol proposal content type: %T", c)
		}
	}
}

// NewHandler returns a handler for accesscontrol messages.
func NewHandler(k keeper.Keeper) sdk.Handler {
	return nil
}
