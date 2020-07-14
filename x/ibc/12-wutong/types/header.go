package types

import (
	"bytes"
	"fmt"
	"time"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	"github.com/cosmos/cosmos-sdk/x/ibc/12-wutong/types/merkle"

	"github.com/tjfoc/gmsm/sm3"

	clientexported "github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
)

var _ clientexported.Header = Header{}

// Header defines the wutongchain consensus Header
type Header struct {
	Version     uint64             `json:"signed_header" yaml:"signed_header"`
	BlockID     BlockID            `json:"block_id" yaml:"block_id"`
	PrevBlockID BlockID            `json:"prev_block_id" yaml:"prev_block_id"`
	Time        time.Time          `json:"time" yaml:"time"`
	StateRoot   tmbytes.HexBytes   `json:"state_root" yaml:"state_root"`
	TxsRoot     tmbytes.HexBytes   `json:"txs_root" yaml:"txs_root"`
	Txs         []tmbytes.HexBytes `json:"txs" yaml:"txs"`
}

type BlockID struct {
	Hash   tmbytes.HexBytes `json:"hash" yaml:"hash"`
	Height uint64           `json:"height" yaml:"height"`
}

func (h Header) ClientType() clientexported.ClientType {
	return clientexported.WuTong
}

func (h Header) GetHeight() uint64 {
	return h.BlockID.Height
}

func (h Header) ValidateBasic() error {
	if h.BlockID.Height != h.PrevBlockID.Height+1 {
		return fmt.Errorf("invalid block:%d", h.BlockID.Height)
	}

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d", h.Version))
	buf.WriteString(fmt.Sprintf("%d", h.BlockID.Height))
	buf.WriteString(h.Time.String())
	buf.Write(h.PrevBlockID.Hash)
	buf.Write(h.StateRoot)
	buf.Write(h.TxsRoot)
	buf.WriteString(fmt.Sprintf("%d", len(h.Txs)))

	var data [][]byte
	for _, hash := range h.Txs {
		buf.Write(hash)
		data = append(data, hash)
	}

	//TODO
	//blockHash := sm3.Sm3Sum(buf.Bytes())
	//if !bytes.Equal(blockHash, h.BlockID.Hash) {
	//	return fmt.Errorf("invalid block:%d,expect blockHash:%x,actual:%x", h.BlockID.Height, blockHash, h.BlockID.Hash)
	//}

	txRoot := merkle.HashFromByteSlices(data)
	if !bytes.Equal(txRoot, h.TxsRoot) {
		return fmt.Errorf("invalid block:%d,expect txsRoot:%x, actual:%x", h.BlockID.Height, txRoot, h.TxsRoot)
	}

	return nil
}

func (h Header) VerifyTx(tx []byte) error {
	txHash := sm3.Sm3Sum(tx)
	for _, hash := range h.Txs {
		if !bytes.Equal(hash, txHash) {
			return nil
		}
	}
	return fmt.Errorf("invalid tx:%s", txHash)
}
