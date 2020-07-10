package types

import (
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clientexported "github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/02-client/types"
	commitmentexported "github.com/cosmos/cosmos-sdk/x/ibc/23-commitment/exported"
)

// ConsensusState defines a Tendermint consensus state
type ConsensusState struct {
	Timestamp    time.Time               `json:"timestamp" yaml:"timestamp"`
	Height       uint64                  `json:"height" yaml:"height"`
}

// NewConsensusState creates a new ConsensusState instance.
func NewConsensusState(
	timestamp time.Time, height uint64,
) ConsensusState {
	return ConsensusState{
		Timestamp:    timestamp,
		Height:       height,
	}
}

// ClientType returns Tendermint
func (ConsensusState) ClientType() clientexported.ClientType {
	return clientexported.WuTong
}

// GetRoot returns the commitment Root for the specific
func (cs ConsensusState) GetRoot() commitmentexported.Root {
	return nil
}

// GetHeight returns the height for the specific consensus state
func (cs ConsensusState) GetHeight() uint64 {
	return cs.Height
}

// GetTimestamp returns block time in nanoseconds at which the consensus state was stored
func (cs ConsensusState) GetTimestamp() uint64 {
	return uint64(cs.Timestamp.UnixNano())
}

// ValidateBasic defines a basic validation for the tendermint consensus state.
func (cs ConsensusState) ValidateBasic() error {
	if cs.Height == 0 {
		return sdkerrors.Wrap(clienttypes.ErrInvalidConsensus, "height cannot be 0")
	}
	if cs.Timestamp.IsZero() {
		return sdkerrors.Wrap(clienttypes.ErrInvalidConsensus, "timestamp cannot be zero Unix time")
	}
	if cs.Timestamp.UnixNano() < 0 {
		return sdkerrors.Wrap(clienttypes.ErrInvalidConsensus, "timestamp cannot be negative Unix time")
	}
	return nil
}
