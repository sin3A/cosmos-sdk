package wutong

import (
	"bytes"
	"time"

	"github.com/cosmos/cosmos-sdk/x/ibc/12-wutong/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clientexported "github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/02-client/types"
)

// CheckValidityAndUpdateState checks if the provided header is valid and updates
// the consensus state if appropriate. It returns an error if:
// - the client or header provided are not parseable to tendermint types
// - the header is invalid
// - header height is lower than the latest client height
//
func CheckValidityAndUpdateState(
	clientState clientexported.ClientState, header clientexported.Header,
	currentTimestamp time.Time,
) (clientexported.ClientState, clientexported.ConsensusState, error) {
	cs, ok := clientState.(types.ClientState)
	if !ok {
		return nil, nil, sdkerrors.Wrap(
			clienttypes.ErrInvalidClientType, "light client is not from Tendermint",
		)
	}

	h, ok := header.(types.Header)
	if !ok {
		return nil, nil, sdkerrors.Wrap(
			clienttypes.ErrInvalidHeader, "header is not from Tendermint",
		)
	}

	if err := checkValidity(cs, h, currentTimestamp); err != nil {
		return nil, nil, err
	}

	tmClientState, consensusState := update(cs, h)
	return tmClientState, consensusState, nil
}

// checkValidity checks if the WuTong header is valid.
//
func checkValidity(
	clientState types.ClientState, header types.Header, currentTimestamp time.Time,
) error {
	// assert header timestamp is past latest clientstate timestamp
	if header.Time.Unix() <= clientState.GetLatestTimestamp().Unix() {
		return sdkerrors.Wrapf(
			clienttypes.ErrInvalidHeader,
			"header blocktime ≤ latest client state block time (%s ≤ %s)",
			header.Time.String(), clientState.GetLatestTimestamp().String(),
		)
	}

	// assert header height is equal last height + 1
	if header.GetHeight() != clientState.GetLatestHeight()+1 {
		return sdkerrors.Wrapf(
			clienttypes.ErrInvalidHeader,
			"header height != latest + 1 client state height (%d != %d)", header.GetHeight(), clientState.GetLatestHeight()+1,
		)
	}

	// assert blockHash is equal last blockHash
	if !bytes.Equal(header.PrevBlockID.Hash, clientState.LastHeader.BlockID.Hash) {
		return sdkerrors.Wrapf(
			clienttypes.ErrInvalidHeader,
			"blockHash != prevBlockHash (%d != %d)", header.GetHeight(), clientState.GetLatestHeight()+1,
		)
	}

	return nil
}

// update the consensus state from a new header
func update(clientState types.ClientState, header types.Header) (types.ClientState, types.ConsensusState) {
	clientState.LastHeader = header
	consensusState := types.ConsensusState{
		Height:    header.BlockID.Height,
		Timestamp: header.Time,
	}
	return clientState, consensusState
}
