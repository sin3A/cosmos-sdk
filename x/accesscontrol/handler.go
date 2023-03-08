package accesscontrol

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/accesscontrol/keeper"
	"github.com/cosmos/cosmos-sdk/x/accesscontrol/types"
)

func NewProposalHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		//ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch c := msg.(type) {
		case *types.MsgUpdateResourceDependencyMappingProposal:
			res, err := msgServer.UpdateResourceDependencyMappingProposal(sdk.WrapSDKContext(ctx), c)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized accesscontrol proposal content type: %T", c)
		}
	}
}
