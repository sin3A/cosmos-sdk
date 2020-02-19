package types

import (
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/sm2"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// nolint:deadcode,unused
var (
	pk1      = sm2.GenPrivKey().PubKey()
	pk2      = sm2.GenPrivKey().PubKey()
	pk3      = sm2.GenPrivKey().PubKey()
	addr1    = pk1.Address()
	addr2    = pk2.Address()
	addr3    = pk3.Address()
	valAddr1 = sdk.ValAddress(addr1)
	valAddr2 = sdk.ValAddress(addr2)
	valAddr3 = sdk.ValAddress(addr3)

	emptyAddr   sdk.ValAddress
	emptyPubkey crypto.PubKey
)
