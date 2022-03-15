package gmssl_test

import (
	// stded25519 "crypto/ed25519"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"

	tmgmssl "github.com/tendermint/tendermint/crypto/gmssl"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"

	gmssl "github.com/cosmos/cosmos-sdk/crypto/keys/gmssl"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sm2 "github.com/cosmos/cosmos-sdk/crypto/keys/sm2"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"bytes"
	"fmt"
)
func TestSm2GmSSLMatched(t *testing.T) {
	sm2sk := sm2.GenPrivKey()
	gmsslsk := gmssl.Bytes2PrivKey(sm2sk.Key)
	// fmt.Println(sm2sk.Key)
	// fmt.Println(gmsslsk.Bytes())
	if bytes.Equal(sm2sk.Bytes()[:], gmsslsk.Bytes()[:]) {
		fmt.Println("Private Key Assign Succeed!")
	} else {
		t.Fatalf("Private Key Assign Failed")
	}

	sm2skMarshal, _ := sm2sk.Marshal()
	gmsslskMarshal, _ := gmsslsk.Marshal()
	if bytes.Equal(sm2skMarshal, gmsslskMarshal) {
		fmt.Println("Private Key Marshal Matched!")
	} else {
		t.Fatalf("Private Key Marshal Failed")
	}

	var sm2skUnmarshal sm2.PrivKey
	if sm2skUnmarshal.Unmarshal(sm2skMarshal) != nil {
		t.Fatalf("sm2sk Unmarshal Failed")
	}
	var gmsslskUnmarshal gmssl.PrivKey
	if gmsslskUnmarshal.Unmarshal((gmsslskMarshal)) != nil {
		t.Fatalf("gmssl Unmarshal Failed")
	}

	if bytes.Equal(sm2skUnmarshal.Bytes()[:], gmsslskUnmarshal.Bytes()[:]) && 
	bytes.Equal(sm2sk.Bytes()[:], gmsslskUnmarshal.Bytes()[:]) {
		fmt.Println("Private Key Unmarshal Matched!")
	} else {
		t.Fatalf("Private Key Unmarshal Failed")
	}

	sm2data := make([]byte, sm2sk.Size())
	gmssldata := make([]byte, gmsslsk.Size())
	// var err error

	fmt.Println(sm2sk.MarshalTo(sm2data))
	fmt.Println(gmsslsk.MarshalTo(gmssldata))
	if bytes.Equal(sm2skUnmarshal.Bytes()[:], gmsslskUnmarshal.Bytes()[:]) && 
	bytes.Equal(sm2data, gmssldata) {
		fmt.Println("Private Key Unmarshal Matched!")
	} else {
		t.Fatalf("Private Key Unmarshal Failed")
	}

	sm2pk := sm2sk.PubKey().(*sm2.PubKey)
	gmsslpk1 := gmssl.Bytes2PubKey(sm2pk.Key)
	gmsslpk2 := gmsslsk.PubKey().(*gmssl.PubKey)
	if bytes.Equal(sm2pk.Bytes()[:], gmsslpk1.Bytes()[:]) &&
	bytes.Equal(sm2pk.Bytes()[:], gmsslpk2.Bytes()[:]) {
		fmt.Println("Public Key Assign Succeed!")
	} else {
		t.Fatalf("Public Key Assign Failed")
	}

	sm2pkMarshal, _ := sm2pk.Marshal()
	gmsslpkMarshal, _ := gmsslpk1.Marshal()
	if bytes.Equal(sm2pkMarshal, gmsslpkMarshal) {
		fmt.Println("Public Key Marshal Matched!")
	} else {
		t.Fatalf("Public Key Marshal Failed")
	}

	var sm2pkUnmarshal sm2.PubKey
	if sm2pkUnmarshal.Unmarshal(sm2pkMarshal) != nil {
		t.Fatalf("sm2pk Unmarshal Failed")
	}
	var gmsslpkUnmarshal gmssl.PubKey
	if gmsslpkUnmarshal.Unmarshal((gmsslpkMarshal)) != nil {
		t.Fatalf("gmsslpk Unmarshal Failed")
	}

	if bytes.Equal(sm2pkUnmarshal.Bytes()[:], gmsslpkUnmarshal.Bytes()[:]) && 
	bytes.Equal(sm2pk.Bytes()[:], gmsslpkUnmarshal.Bytes()[:]) {
		fmt.Println("Public Key Unmarshal Matched!")
	} else {
		t.Fatalf("Public Key Unmarshal Failed")
	}

	sm2pkdata := make([]byte, sm2pk.Size())
	gmsslpkdata := make([]byte, gmsslpk1.Size())
	// var err error

	fmt.Println(sm2pk.MarshalTo(sm2pkdata))
	fmt.Println(gmsslpk1.MarshalTo(gmsslpkdata))
	if bytes.Equal(sm2pkdata, gmsslpkdata) {
		fmt.Println("Public Key Unmarshal Matched!")
	} else {
		t.Fatalf("Public Key Unmarshal Failed")
	}
}

func TestSignAndValidateGmSSL(t *testing.T) {
	privKey := gmssl.GenPrivKey()
	pubKey := privKey.PubKey()

	msg := crypto.CRandBytes(1000)
	sig, err := privKey.Sign(msg)
	require.Nil(t, err)

	// Test the signature
	assert.True(t, pubKey.VerifySignature(msg, sig))

	// ----
	// Test cross packages verification

	// ----
	// Mutate the signature, just one bit.
	// TODO: Replace this with a much better fuzzer, tendermint/ed25519/issues/10
	sig[7] ^= byte(0x01)
	assert.False(t, pubKey.VerifySignature(msg, sig))
}

func TestPubKeyEquals(t *testing.T) {
	gmsslPubKey := gmssl.GenPrivKey().PubKey().(*gmssl.PubKey)

	testCases := []struct {
		msg      string
		pubKey   cryptotypes.PubKey
		other    cryptotypes.PubKey
		expectEq bool
	}{
		{
			"different bytes",
			gmsslPubKey,
			gmssl.GenPrivKey().PubKey(),
			false,
		},
		{
			"equals",
			gmsslPubKey,
			gmssl.Bytes2PubKey(gmsslPubKey.Bytes()),
			true,
		},
		{
			"different types",
			gmsslPubKey,
			secp256k1.GenPrivKey().PubKey(),
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.msg, func(t *testing.T) {
			eq := tc.pubKey.Equals(tc.other)
			require.Equal(t, eq, tc.expectEq)
		})
	}
}

func TestAddressGmSSL(t *testing.T) {
	pk := gmssl.Bytes2PubKey([]byte{125, 80, 29, 208, 159, 53, 119, 198, 73, 53, 187, 33, 199, 144, 62, 255, 1, 235, 117, 96, 128, 211, 17, 45, 34, 64, 189, 165, 33, 182, 54, 206})
	addr := pk.Address()
	require.Len(t, addr, 20, "Address must be 20 bytes long")
}

func TestPrivKeyEquals(t *testing.T) {
	gmsslPrivKey := gmssl.GenPrivKey()

	testCases := []struct {
		msg      string
		privKey  cryptotypes.PrivKey
		other    cryptotypes.PrivKey
		expectEq bool
	}{
		{
			"different bytes",
			gmsslPrivKey,
			gmssl.GenPrivKey(),
			false,
		},
		{
			"equals",
			gmsslPrivKey,
			gmssl.Bytes2PrivKey(gmsslPrivKey.Bytes()),
			true,
		},
		{
			"different types",
			gmsslPrivKey,
			secp256k1.GenPrivKey(),
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.msg, func(t *testing.T) {
			eq := tc.privKey.Equals(tc.other)
			require.Equal(t, eq, tc.expectEq)
		})
	}
}

func TestMarshalAmino(t *testing.T) {
	aminoCdc := codec.NewLegacyAmino()
	privKey := gmssl.GenPrivKey()
	pubKey := privKey.PubKey().(*gmssl.PubKey)

	testCases := []struct {
		desc      string
		msg       codec.AminoMarshaler
		typ       interface{}
		expBinary []byte
		expJSON   string
	}{
		{
			"gmssl private key",
			privKey,
			&gmssl.PrivKey{},
			append([]byte{32}, privKey.Bytes()...), // Length-prefixed.
			"\"" + base64.StdEncoding.EncodeToString(privKey.Bytes()) + "\"",
		},
		{
			"gmssl public key",
			pubKey,
			&gmssl.PubKey{},
			append([]byte{33}, pubKey.Bytes()...), // Length-prefixed.
			"\"" + base64.StdEncoding.EncodeToString(pubKey.Bytes()) + "\"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// Do a round trip of encoding/decoding binary.
			bz, err := aminoCdc.Marshal(tc.msg)

			require.NoError(t, err)

			require.Equal(t, tc.expBinary, bz)

			err = aminoCdc.Unmarshal(bz, tc.typ)
			require.NoError(t, err)

			require.Equal(t, tc.msg, tc.typ)

			// Do a round trip of encoding/decoding JSON.
			bz, err = aminoCdc.MarshalJSON(tc.msg)

			require.NoError(t, err)
			require.Equal(t, tc.expJSON, string(bz))
			fmt.Println(bz)

			err = aminoCdc.UnmarshalJSON(bz, tc.typ)
			require.NoError(t, err)

			require.Equal(t, tc.msg, tc.typ)
		})
	}
}

func TestMarshalAmino_BackwardsCompatibility(t *testing.T) {
	aminoCdc := codec.NewLegacyAmino()
	// Create Tendermint keys.
	tmPrivKey := tmgmssl.GenPrivKey()
	tmPrivKeyBytes := tmPrivKey.Bytes()
	tmPubKey := tmPrivKey.PubKey().(*tmgmssl.PubKeySm2)
	tmPubKeyBytes := tmPubKey.Bytes()
	// Create our own keys, with the same private key as Tendermint's.
	privKey := gmssl.Bytes2PrivKey(tmPrivKey.Bytes())
	pubKey := privKey.PubKey().(*gmssl.PubKey)

	testCases := []struct {
		desc      string
		tmKey     interface{}
		ourKey    interface{}
		marshalFn func(o interface{}) ([]byte, error)
	}{
		{
			"gmssl private key, binary",
			tmPrivKeyBytes,
			privKey,
			aminoCdc.Marshal,
		},
		{
			"gmssl private key, JSON",
			tmPrivKeyBytes,
			privKey,
			aminoCdc.MarshalJSON,
		},
		{
			"gmssl public key, binary",
			tmPubKeyBytes,
			pubKey,
			aminoCdc.Marshal,
		},
		{
			"gmssl public key, JSON",
			tmPubKeyBytes,
			pubKey,
			aminoCdc.MarshalJSON,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// Make sure Amino encoding override is not breaking backwards compatibility.
			bz1, err := tc.marshalFn(tc.tmKey)
			require.NoError(t, err)
			bz2, err := tc.marshalFn(tc.ourKey)
			require.NoError(t, err)
			require.Equal(t, bz1, bz2)
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	require := require.New(t)
	privKey := gmssl.GenPrivKey()
	pk := privKey.PubKey()

	registry := types.NewInterfaceRegistry()
	cryptocodec.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)

	bz, err := cdc.MarshalInterfaceJSON(pk)
	require.NoError(err)

	var pk2 cryptotypes.PubKey
	err = cdc.UnmarshalInterfaceJSON(bz, &pk2)
	require.NoError(err)
	require.True(pk2.Equals(pk))
}
