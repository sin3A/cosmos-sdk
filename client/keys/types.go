package keys

import (
	"errors"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// used for outputting keys.Info over REST

// AddNewKey request a new key
type AddNewKey struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Mnemonic string `json:"mnemonic"`
	Account  int    `json:"account,string,omitempty"`
	Index    int    `json:"index,string,omitempty"`
}

// NewAddNewKey constructs a new AddNewKey request structure.
func NewAddNewKey(name, password, mnemonic string, account, index int) AddNewKey {
	return AddNewKey{
		Name:     name,
		Password: password,
		Mnemonic: mnemonic,
		Account:  account,
		Index:    index,
	}
}

// RecoverKeyBody recovers a key
type RecoverKey struct {
	Password string `json:"password"`
	Mnemonic string `json:"mnemonic"`
	Account  int    `json:"account,string,omitempty"`
	Index    int    `json:"index,string,omitempty"`
}

// NewRecoverKey constructs a new RecoverKey request structure.
func NewRecoverKey(password, mnemonic string, account, index int) RecoverKey {
	return RecoverKey{Password: password, Mnemonic: mnemonic, Account: account, Index: index}
}

// UpdateKeyReq requests updating a key
type UpdateKeyReq struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// NewUpdateKeyReq constructs a new UpdateKeyReq structure.
func NewUpdateKeyReq(old, new string) UpdateKeyReq {
	return UpdateKeyReq{OldPassword: old, NewPassword: new}
}

// DeleteKeyReq requests deleting a key
type DeleteKeyReq struct {
	Password string `json:"password"`
}

// NewDeleteKeyReq constructs a new DeleteKeyReq structure.
func NewDeleteKeyReq(password string) DeleteKeyReq { return DeleteKeyReq{Password: password} }

type signedMsg struct {
	ChainID string `json:"chain_id"`
	Type    string `json:"type"`
	Data    []byte `json:"data"`
	Sig     []byte `json:"sig"`
}

func newSignedMsg(chainID, typ string, data, sig []byte) signedMsg {
	return signedMsg{ChainID: chainID, Type: typ, Data: data, Sig: sig}
}

func (m signedMsg) Bytes() []byte {
	m.Sig = nil
	bz, err := cdc.MarshalJSON(m)
	if err != nil {
		panic(err)
	}

	return sdk.MustSortJSON(bz)
}

func (m signedMsg) Validate() error {
	if len(strings.TrimSpace(m.ChainID)) == 0 {
		return errors.New("chain_id must not be empty")
	}
	if len(m.Data) == 0 {
		return errors.New("payload must not be empty")
	}

	switch m.Type {
	case "string":
		if !printableASCIIRegex.MatchString(string(m.Data)) {
			return errors.New("string payload must contain only ASCII printable characters")
		}
		if len(string(m.Data)) == 0 {
			return errors.New("string payload must not be empty")
		}
		if len(string(m.Data)) > stringMsgMaxLength {
			return errors.New("string payload size must be smaller than 256 characters")
		}
	case "object":
		if len(string(m.Data)) > objectMsgMaxLength {
			return errors.New("string payload size must be smaller than 256 characters")
		}
	default:
		return errors.New("message type must be either string or object")
	}

	return nil
}
