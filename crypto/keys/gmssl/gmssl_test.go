package gmssl_test

import (
	// stded25519 "crypto/ed25519"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"
	// tmed25519 "github.com/tendermint/tendermint/crypto/ed25519"
	tmgmssl "github.com/tendermint/tendermint/crypto/gmssl"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	// ed25519 "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	gmssl "github.com/cosmos/cosmos-sdk/crypto/keys/gmssl"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"fmt"
)

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
			"ed25519 private key",
			privKey,
			&gmssl.PrivKey{},
			append([]byte{64}, privKey.Bytes()...), // Length-prefixed.
			"\"" + base64.StdEncoding.EncodeToString(privKey.Bytes()) + "\"",
		},
		{
			"ed25519 public key",
			pubKey,
			&gmssl.PubKey{},
			append([]byte{32}, pubKey.Bytes()...), // Length-prefixed.
			"\"" + base64.StdEncoding.EncodeToString(pubKey.Bytes()) + "\"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// Do a round trip of encoding/decoding binary.
			fmt.Println(tc.desc)
			// fmt.Println(tc.msg)
			fmt.Println(tc.expBinary)
			bz, err := aminoCdc.Marshal(tc.msg)
			require.NoError(t, err)
			require.Equal(t, tc.expBinary, bz)
			fmt.Println(bz)

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
	tmPubKey := tmPrivKey.PubKey()
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
			tmPrivKey,
			privKey,
			aminoCdc.Marshal,
		},
		{
			"gmssl private key, JSON",
			tmPrivKey,
			privKey,
			aminoCdc.MarshalJSON,
		},
		{
			"gmssl public key, binary",
			tmPubKey,
			pubKey,
			aminoCdc.Marshal,
		},
		{
			"gmssl public key, JSON",
			tmPubKey,
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
