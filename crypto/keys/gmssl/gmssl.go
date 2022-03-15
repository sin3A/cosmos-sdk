package gmssl

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"math/big"

	cryptobyte "golang.org/x/crypto/cryptobyte"
	cbasn1 "golang.org/x/crypto/cryptobyte/asn1"


	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	gmssl "github.com/tendermint/tendermint/crypto/gmssl"
)

const (
	PrivKeyName = "cosmos/PrivKeyGmSSL"
	PubKeyName  = "cosmos/PubKeyGmSSL"

	PrivKeySize   = 32
	PubKeySize    = 33
	SignatureSize = 64

	keyType = "gmssl"
)

var (
	_ cryptotypes.PrivKey  = &PrivKey{}
	_ codec.AminoMarshaler = &PrivKey{}
)

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func (privKey *PrivKey) Type() string {
	return keyType
}

func (privKey PrivKey) MarshalAmino() ([]byte, error) {
	return privKey.Key.GetKeyBuffer()
}

func (privKey *PrivKey) UnmarshalAmino(bz []byte) error {
	if len(bz) != PrivKeySize {
		return fmt.Errorf("invalid privkey size")
	}
	sm2keygenargs := [][2]string{
		{"ec_paramgen_curve", "sm2p256v1"},
		{"ec_param_enc", "named_curve"},
	}
	sm2sk, err := gmssl.GeneratePrivateKeyByBuffer("EC", sm2keygenargs, nil, bz[:])
	PanicError(err)
	privKey.Key = sm2sk

	return nil
}

// MarshalAminoJSON overrides Amino JSON marshalling.
func (privKey PrivKey) MarshalAminoJSON() ([]byte, error) {
	// When we marshal to Amino JSON, we don't marshal the "key" field itself,
	// just its contents (i.e. the key bytes).
	return privKey.MarshalAmino()
}

// UnmarshalAminoJSON overrides Amino JSON marshalling.
func (privKey *PrivKey) UnmarshalAminoJSON(bz []byte) error {
	return privKey.UnmarshalAmino(bz)
}

func (privKey *PrivKey) Bytes() []byte {
	ret, err := privKey.Key.GetKeyBuffer()
	PanicError(err)
	return ret
}

func (privKey *PrivKey) Sign(msg []byte) ([]byte, error) {
	sm3ctx, err := gmssl.NewDigestContext(gmssl.SM3)
	PanicError(err)
	err = sm3ctx.Reset()
	PanicError(err)
	
	default_uid := "1234567812345678"
	sm2_zid, err:= privKey.Key.ComputeSM2IDDigest(default_uid)
	PanicError(err)

	err = sm3ctx.Update(sm2_zid)
	PanicError(err)

	err = sm3ctx.Update(msg)
	PanicError(err)

	digest, err := sm3ctx.Final()
	PanicError(err)

	sig, err := privKey.Key.Sign("sm2sign", digest, nil)
	if err != nil {
		return nil, err
	}
	var (
		r, s  = &big.Int{}, &big.Int{}
		inner cryptobyte.String
	)
	input := cryptobyte.String(sig)
	if !input.ReadASN1(&inner, cbasn1.SEQUENCE) ||
		!input.Empty() ||
		!inner.ReadASN1Integer(r) ||
		!inner.ReadASN1Integer(s) ||
		!inner.Empty() {
		return nil, gmssl.GetErrors()
	}

	ret := make([]byte, 64)
	R := r.Bytes()
	S := s.Bytes()
	copy(ret[32-len(R):32], R[:])
	copy(ret[64-len(S):], S[:])

	return ret[:], nil
}

func (privKey *PrivKey) Equals(other cryptotypes.LedgerPrivKey) bool {
	if privKey.Type() != other.Type() {
		return false
	}

	return subtle.ConstantTimeCompare(privKey.Bytes(), other.Bytes()) == 1
}

func Bytes2PrivKey(bz []byte) *PrivKey {
	if len(bz) != PrivKeySize {
		return nil
	}
	sm2keygenargs := [][2]string{
		{"ec_paramgen_curve", "sm2p256v1"},
		{"ec_param_enc", "named_curve"},
	}
	sk, err := gmssl.GeneratePrivateKeyByBuffer("EC", sm2keygenargs, nil, bz[:])
	PanicError(err)

	return &PrivKey{Key: sk}
}

func (privKey *PrivKey) PubKey() cryptotypes.PubKey {
	pk, err := privKey.Key.GetPublicKey()
	PanicError(err)

	return &PubKey{Key: pk}
}

func GenPrivKey() *PrivKey {
	sm2keygenargs := [][2]string{
		{"ec_paramgen_curve", "sm2p256v1"},
		{"ec_param_enc", "named_curve"},
	}
	sm2sk, err := gmssl.GeneratePrivateKey("EC", sm2keygenargs, nil)
	PanicError(err)
	return &PrivKey{Key: sm2sk}
}

func GenPrivKeyFromSecret(secret []byte) *PrivKey {
	secHash := sha256.Sum256(secret);
	sm2keygenargs := [][2]string{
		{"ec_paramgen_curve", "sm2p256v1"},
		{"ec_param_enc", "named_curve"},
	}
	sm2sk, err := gmssl.GeneratePrivateKeyBySecret("EC", sm2keygenargs, nil, secHash[:])
	PanicError(err)
	return &PrivKey{Key :sm2sk}
}




var _ cryptotypes.PubKey = &PubKey{}
var _ codec.AminoMarshaler = &PubKey{}

func (pubKey *PubKey) Address() cryptotypes.Address {
	if len(pubKey.Bytes()) != PubKeySize {
		panic("pubkey is incorrect size")
	}
	// For ADR-28 compatible address we would need to
	// return address.Hash(proto.MessageName(pubKey), pubKey.Key)
	return crypto.Address(tmhash.SumTruncated(pubKey.Bytes()))
}

func (pubKey *PubKey) Bytes() []byte {
	ret, err := pubKey.Key.GetKeyBuffer()
	PanicError(err)
	return ret[:]
}

func (pubKey *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	// make sure we use the same algorithm to sign
	if len(sig) != SignatureSize {
		return false
	}

	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:])

	var b cryptobyte.Builder
	b.AddASN1(cbasn1.SEQUENCE, func(b *cryptobyte.Builder) {
		b.AddASN1BigInt(r)
		b.AddASN1BigInt(s)
	})

	sm3ctx, err := gmssl.NewDigestContext(gmssl.SM3)
	PanicError(err)

	err = sm3ctx.Reset()
	PanicError(err)
	
	default_uid := "1234567812345678"
	sm2_zid, err := pubKey.Key.ComputeSM2IDDigest(default_uid)
	PanicError(err)

	err = sm3ctx.Update(sm2_zid)
	PanicError(err)

	err = sm3ctx.Update(msg)
	PanicError(err)

	digest, err := sm3ctx.Final()
	PanicError(err)

	ret, err := b.Bytes()
	PanicError(err)

	return pubKey.Key.Verify("sm2sign", digest, ret, nil) == nil
}

func (pubKey *PubKey) String() string {
	return fmt.Sprintf("PubKeyEd25519{%X}", pubKey.Bytes())
}

func (pubKey *PubKey) Type() string {
	return keyType
}

func (pubKey *PubKey) Equals(other cryptotypes.PubKey) bool {
	if pubKey.Type() != other.Type() {
		return false
	}

	return subtle.ConstantTimeCompare(pubKey.Bytes(), other.Bytes()) == 1
}

func Bytes2PubKey(bz []byte) *PubKey {
	if len(bz) != PubKeySize {
		return nil
	}

	sm2keygenargs := [][2]string{
		{"ec_paramgen_curve", "sm2p256v1"},
		{"ec_param_enc", "named_curve"},
	}
	pubkey, err := gmssl.GeneratePublicKeyByBuffer("EC", sm2keygenargs, nil, bz[:])
	PanicError(err)

	// var pubKey *PubKey
	// pubKey.Key = pubkey
	return &PubKey{Key: pubkey}
}

func (pubKey PubKey) MarshalAmino() ([]byte, error) {
	return (&pubKey).Bytes(), nil
}

func (pubKey *PubKey) UnmarshalAmino(bz []byte) error {
	if len(bz) != PubKeySize {
		return errors.Wrap(errors.ErrInvalidPubKey, "invalid pubkey size")
	}

	sm2keygenargs := [][2]string{
		{"ec_paramgen_curve", "sm2p256v1"},
		{"ec_param_enc", "named_curve"},
	}
	pubkey, err := gmssl.GeneratePublicKeyByBuffer("EC", sm2keygenargs, nil, bz[:])
	PanicError(err)

	pubKey.Key = pubkey
	return nil
}

func (pubKey PubKey) MarshalAminoJSON() ([]byte, error) {
	// When we marshal to Amino JSON, we don't marshal the "key" field itself,
	// just its contents (i.e. the key bytes).
	return pubKey.MarshalAmino()
}

func (pubKey *PubKey) UnmarshalAminoJSON(bz []byte) error {
	return pubKey.UnmarshalAmino(bz)
}