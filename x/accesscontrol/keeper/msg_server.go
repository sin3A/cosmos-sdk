package keeper

import (
	context "context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/accesscontrol/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) UpdateResourceDependencyMappingProposal(ctx context.Context, proposal *types.MsgUpdateResourceDependencyMappingProposal) (*types.MsgUpdateResourceDependencyMappingProposalResponse, error) {
	ctxMsg := sdk.UnwrapSDKContext(ctx)
	_, err := HandleMsgUpdateResourceDependencyMappingProposal(ctxMsg, m.Keeper, proposal)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func HandleMsgUpdateResourceDependencyMappingProposal(ctx sdk.Context, k Keeper, p *types.MsgUpdateResourceDependencyMappingProposal) (*sdk.Result, error) {
	for _, resourceDepMapping := range p.MessageDependencyMapping {
		k.SetResourceDependencyMapping(ctx, resourceDepMapping)
	}
	return nil, nil
}
