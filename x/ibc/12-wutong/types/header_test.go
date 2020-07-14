package types

import (
	"encoding/hex"
	"github.com/stretchr/testify/require"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"testing"
	"time"
)

func TestHeaderValidateBasic(t *testing.T) {

	blockHash, err := hex.DecodeString("2b60353e89141d5afb5ad1a33bb3876fa808c73e145cea0b34a7e5e18dc8b217")
	require.NoError(t, err)

	prevBlockHash, err := hex.DecodeString("3e4687b43e01e941e25daa67114a230c1fdb7f4f1d93c3d168ce6229ae93659a")
	require.NoError(t, err)

	stateRoot, err := hex.DecodeString("1f2e3e5f32bbaf42ec5d7b5e2539622406fc8bbe7182a5fc2e100df6fed9dea14242424242424242424242424242424242424242424242424242424242424242")
	require.NoError(t, err)

	txsRoot, err := hex.DecodeString("41f7fda2dc0d9e518d4525cb4430b92b771a7d57bed5e03b66d5d0b4f83a08ba")
	require.NoError(t, err)

	tx, err := hex.DecodeString("697566e5c32219c6568cb694c6d61d6d50ace004f402f0cc5a724eeb35d26ea8")
	require.NoError(t, err)

	header := Header{
		Version: 0,
		BlockID: BlockID{
			Hash:   blockHash,
			Height: 3,
		},
		PrevBlockID: BlockID{
			Hash:   prevBlockHash,
			Height: 2,
		},
		Time:      time.Unix(1593411367299, 0),
		StateRoot: stateRoot,
		TxsRoot:   txsRoot,
		Txs:       []tmbytes.HexBytes{tx},
	}

	err = header.ValidateBasic()
	require.NoError(t, err)
}
