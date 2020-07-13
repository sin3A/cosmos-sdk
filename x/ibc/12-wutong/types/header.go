package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/x/ibc/12-wutong/types/merkle"

	"github.com/tjfoc/gmsm/sm3"

	clientexported "github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
)

var _ clientexported.Header = Header{}

// Header defines the wutongchain consensus Header
type Header struct {
	Version         uint64    `json:"signed_header" yaml:"signed_header"`
	Time            time.Time `json:"time" yaml:"time"`
	BlockID         BlockID   `json:"block_id" yaml:"block_id"`
	PrevBlockID     BlockID   `json:"prev_block_id" yaml:"prev_block_id"`
	WorldStateRoot  []byte    `json:"world_state_root" yaml:"world_state_root"`
	TransactionRoot []byte    `json:"transaction_root" yaml:"transaction_root"`
	Txs             []string  `json:"txs" yaml:"txs"`
}

type BlockID struct {
	Hash   string `json:"hash" yaml:"hash"`
	Height uint64 `json:"height" yaml:"height"`
}

func (h Header) ClientType() clientexported.ClientType {
	return clientexported.WuTong
}

func (h Header) GetHeight() uint64 {
	return h.BlockID.Height
}

func (h Header) ValidateBasic() error {
	if h.BlockID.Height < h.PrevBlockID.Height {
		return fmt.Errorf("invalid bolck:%d", h.BlockID.Height)
	}

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d", h.Version))
	buf.WriteString(fmt.Sprintf("%d", h.BlockID.Height))
	buf.WriteString(h.Time.String())
	buf.WriteString(h.PrevBlockID.Hash)
	buf.Write(h.WorldStateRoot)
	buf.WriteString(fmt.Sprintf("%d", len(h.Txs)))

	var data [][]byte
	for _, hash := range h.Txs {
		buf.WriteString(hash)
		bz, err := hex.DecodeString(hash)
		if err != nil {
			return err
		}
		data = append(data, bz)
	}

	blockHash := hex.EncodeToString(sm3.Sm3Sum(buf.Bytes()))
	if blockHash == h.BlockID.Hash {
		return nil
	}

	txRoot := merkle.HashFromByteSlices(data)
	if bytes.Equal(txRoot, h.TransactionRoot) {
		return nil
	}
	return fmt.Errorf("invalid bolck:%d", h.BlockID.Height)
}

func (h Header) VerifyTx(tx []byte) error {
	txHash := hex.EncodeToString(sm3.Sm3Sum(tx))
	for _, hash := range h.Txs {
		if hash == txHash {
			return nil
		}
	}
	return fmt.Errorf("invalid tx:%s", txHash)
}
