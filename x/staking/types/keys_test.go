package types

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/crypto/sm2"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	keysPK1   = sm2.GenPrivKeySm2FromSecret([]byte{1}).PubKey()
	keysPK2   = sm2.GenPrivKeySm2FromSecret([]byte{2}).PubKey()
	keysPK3   = sm2.GenPrivKeySm2FromSecret([]byte{3}).PubKey()
	keysAddr1 = keysPK1.Address()
	keysAddr2 = keysPK2.Address()
	keysAddr3 = keysPK3.Address()
)

func TestGetValidatorPowerRank(t *testing.T) {
	valAddr1 := sdk.ValAddress(keysAddr1)
	emptyDesc := Description{}
	val1 := NewValidator(valAddr1, keysPK1, emptyDesc)
	val1.Tokens = sdk.ZeroInt()
	val2, val3, val4 := val1, val1, val1
	val2.Tokens = sdk.TokensFromConsensusPower(1)
	val3.Tokens = sdk.TokensFromConsensusPower(10)
	x := new(big.Int).Exp(big.NewInt(2), big.NewInt(40), big.NewInt(0))
	val4.Tokens = sdk.TokensFromConsensusPower(x.Int64())

	tests := []struct {
		validator Validator
		wantHex   string
	}{
		{val1, "23000000000000000088f7334fe55129cd834bd11578d4affab3549224"},
		{val2, "23000000000000000188f7334fe55129cd834bd11578d4affab3549224"},
		{val3, "23000000000000000a88f7334fe55129cd834bd11578d4affab3549224"},
		{val4, "23000001000000000088f7334fe55129cd834bd11578d4affab3549224"},
	}
	for i, tt := range tests {
		got := hex.EncodeToString(getValidatorPowerRank(tt.validator))

		assert.Equal(t, tt.wantHex, got, "Keys did not match on test case %d", i)
	}
}

func TestGetREDByValDstIndexKey(t *testing.T) {
	tests := []struct {
		delAddr    sdk.AccAddress
		valSrcAddr sdk.ValAddress
		valDstAddr sdk.ValAddress
		wantHex    string
	}{
		{sdk.AccAddress(keysAddr1), sdk.ValAddress(keysAddr1), sdk.ValAddress(keysAddr1),
			"367708ccb01aaed6327cb42eea872b50054cab6ddb7708ccb01aaed6327cb42eea872b50054cab6ddb7708ccb01aaed6327cb42eea872b50054cab6ddb"},
		{sdk.AccAddress(keysAddr1), sdk.ValAddress(keysAddr2), sdk.ValAddress(keysAddr3),
			"363e82b09ddae3c553ac4463ebf5a66b33b0960e8c7708ccb01aaed6327cb42eea872b50054cab6ddb1704abd06ea45315ec7a33d3c70f18317712f4f1"},
		{sdk.AccAddress(keysAddr2), sdk.ValAddress(keysAddr1), sdk.ValAddress(keysAddr3),
			"363e82b09ddae3c553ac4463ebf5a66b33b0960e8c1704abd06ea45315ec7a33d3c70f18317712f4f17708ccb01aaed6327cb42eea872b50054cab6ddb"},
	}
	for i, tt := range tests {
		got := hex.EncodeToString(GetREDByValDstIndexKey(tt.delAddr, tt.valSrcAddr, tt.valDstAddr))

		assert.Equal(t, tt.wantHex, got, "Keys did not match on test case %d", i)
	}
}

func TestGetREDByValSrcIndexKey(t *testing.T) {
	tests := []struct {
		delAddr    sdk.AccAddress
		valSrcAddr sdk.ValAddress
		valDstAddr sdk.ValAddress
		wantHex    string
	}{
		{sdk.AccAddress(keysAddr1), sdk.ValAddress(keysAddr1), sdk.ValAddress(keysAddr1),
			"357708ccb01aaed6327cb42eea872b50054cab6ddb7708ccb01aaed6327cb42eea872b50054cab6ddb7708ccb01aaed6327cb42eea872b50054cab6ddb"},
		{sdk.AccAddress(keysAddr1), sdk.ValAddress(keysAddr2), sdk.ValAddress(keysAddr3),
			"351704abd06ea45315ec7a33d3c70f18317712f4f17708ccb01aaed6327cb42eea872b50054cab6ddb3e82b09ddae3c553ac4463ebf5a66b33b0960e8c"},
		{sdk.AccAddress(keysAddr2), sdk.ValAddress(keysAddr1), sdk.ValAddress(keysAddr3),
			"357708ccb01aaed6327cb42eea872b50054cab6ddb1704abd06ea45315ec7a33d3c70f18317712f4f13e82b09ddae3c553ac4463ebf5a66b33b0960e8c"},
	}
	for i, tt := range tests {
		got := hex.EncodeToString(GetREDByValSrcIndexKey(tt.delAddr, tt.valSrcAddr, tt.valDstAddr))

		assert.Equal(t, tt.wantHex, got, "Keys did not match on test case %d", i)
	}
}
