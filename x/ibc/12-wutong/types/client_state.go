package types

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	clientexported "github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
	connectionexported "github.com/cosmos/cosmos-sdk/x/ibc/03-connection/exported"
	channelexported "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/exported"
	commitmentexported "github.com/cosmos/cosmos-sdk/x/ibc/23-commitment/exported"
	host "github.com/cosmos/cosmos-sdk/x/ibc/24-host"
)

var _ clientexported.ClientState = ClientState{}

// ClientState requires (read-only) access to keys outside the client prefix.
type ClientState struct {
	ID string `json:"id" yaml:"id"`
	// Last Header that was stored by client
	LastHeader Header `json:"last_header" yaml:"last_header"`
}

// NewClientState creates a new ClientState instance
func NewClientState(clientID string, header Header) ClientState {
	return ClientState{
		ID:         clientID,
		LastHeader: header,
	}
}

// GetID returns the loop-back client state identifier.
func (cs ClientState) GetID() string {
	return cs.ID
}

// GetChainID returns an empty string
func (cs ClientState) GetChainID() string {
	return clientexported.ClientTypeWuTong
}

// ClientType is localhost.
func (cs ClientState) ClientType() clientexported.ClientType {
	return clientexported.WuTong
}

// GetLatestHeight returns the latest height stored.
func (cs ClientState) GetLatestHeight() uint64 {
	return cs.LastHeader.BlockID.Height
} // GetLatestHeight returns the latest height stored.

func (cs ClientState) GetLatestTimestamp() time.Time {
	return cs.LastHeader.Time
}

// IsFrozen returns false.
func (cs ClientState) IsFrozen() bool {
	return false
}

// Validate performs a basic validation of the client state fields.
func (cs ClientState) Validate() error {
	if err := host.DefaultClientIdentifierValidator(cs.ID); err != nil {
		return err
	}
	if err := cs.LastHeader.ValidateBasic(); err != nil {
		return err
	}
	return nil
}

// VerifyClientConsensusState verifies a proof of the consensus
// state of the loop-back client.
// VerifyClientConsensusState verifies a proof of the consensus state of the
// Tendermint client stored on the target machine.
func (cs ClientState) VerifyClientConsensusState(
	cdc *codec.Codec,
	_ commitmentexported.Root,
	height uint64,
	_ string,
	consensusHeight uint64,
	prefix commitmentexported.Prefix,
	_ commitmentexported.Proof,
	consensusState clientexported.ConsensusState,
) error {
	//nothing
	return nil
}

// VerifyConnectionState verifies a proof of the connection state of the
// specified connection end stored locally.
func (cs ClientState) VerifyConnectionState(
	cdc codec.Marshaler,
	_ uint64,
	prefix commitmentexported.Prefix,
	_ commitmentexported.Proof,
	connectionID string,
	connectionEnd connectionexported.ConnectionI,
	_ clientexported.ConsensusState,
) error {
	//nothing
	return nil
}

// VerifyChannelState verifies a proof of the channel state of the specified
// channel end, under the specified port, stored on the local machine.
func (cs ClientState) VerifyChannelState(
	cdc codec.Marshaler,
	_ uint64,
	prefix commitmentexported.Prefix,
	_ commitmentexported.Proof,
	portID,
	channelID string,
	channel channelexported.ChannelI,
	_ clientexported.ConsensusState,
) error {
	//nothing
	return nil
}

// VerifyPacketCommitment verifies a proof of an outgoing packet commitment at
// the specified port, specified channel, and specified sequence.
func (cs ClientState) VerifyPacketCommitment(
	_ uint64,
	_ commitmentexported.Prefix,
	_ commitmentexported.Proof,
	_,
	_ string,
	_ uint64,
	txRawData []byte,
	_ clientexported.ConsensusState,
) error {
	return cs.LastHeader.VerifyTx(txRawData)
}

// VerifyPacketAcknowledgement verifies a proof of an incoming packet
// acknowledgement at the specified port, specified channel, and specified sequence.
func (cs ClientState) VerifyPacketAcknowledgement(
	_ uint64,
	prefix commitmentexported.Prefix,
	_ commitmentexported.Proof,
	portID,
	channelID string,
	sequence uint64,
	acknowledgement []byte,
	_ clientexported.ConsensusState,
) error {
	//nothing
	return nil
}

// VerifyPacketAcknowledgementAbsence verifies a proof of the absence of an
// incoming packet acknowledgement at the specified port, specified channel, and
// specified sequence.
func (cs ClientState) VerifyPacketAcknowledgementAbsence(
	_ uint64,
	prefix commitmentexported.Prefix,
	_ commitmentexported.Proof,
	portID,
	channelID string,
	sequence uint64,
	_ clientexported.ConsensusState,
) error {
	//nothing
	return nil
}

// VerifyNextSequenceRecv verifies a proof of the next sequence number to be
// received of the specified channel at the specified port.
func (cs ClientState) VerifyNextSequenceRecv(
	_ uint64,
	prefix commitmentexported.Prefix,
	_ commitmentexported.Proof,
	portID,
	channelID string,
	nextSequenceRecv uint64,
	_ clientexported.ConsensusState,
) error {
	//nothing
	return nil
}

// consensusStatePath takes an Identifier and returns a Path under which to
// store the consensus state of a client.
func consensusStatePath(clientID string) string {
	return fmt.Sprintf("consensusState/%s", clientID)
}
