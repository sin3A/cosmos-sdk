package types

import (
	"time"

	clientexported "github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
)

var _ clientexported.Header = Header{}

// Header defines the wutongchain consensus Header
type Header struct {
	Version         uint64    `json:"signed_header" yaml:"signed_header"`
	Height          uint64    `json:"validator_set" yaml:"validator_set"`
	Time            time.Time `json:"time" yaml:"time"`
	BlockID         string    `json:"block_id" yaml:"block_id"`
	LastBlockID     string    `json:"last_block_id" yaml:"last_block_id"`
	WorldStateRoot  []byte    `json:"world_state_root" yaml:"world_state_root"`
	TransactionRoot []byte    `json:"transaction_root" yaml:"transaction_root"`
}

func (h Header) ClientType() clientexported.ClientType {
	return clientexported.WuTong
}

func (h Header) GetHeight() uint64 {
	return h.Height
}

// ValidateBasic calls the SignedHeader ValidateBasic function
// and checks that validatorsets are not nil
func (h Header) ValidateBasic() error {
	//TODO
	return nil
}
