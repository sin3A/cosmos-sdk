package merkle

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	tmmerkle "github.com/tendermint/tendermint/crypto/merkle"
	"testing"
)

func getTestData(t *testing.T) (data [][]byte) {
	data1, err := hex.DecodeString("7b245a1d3eb1c27f5b0270e0716327e9325e84feab1ea2134e5f529ab88397f5")
	require.NoError(t, err)
	data = append(data, data1)

	data2, err := hex.DecodeString("8c318b09ad2c63c527286fc6b8d99c8f73b253529c3a956e0f91b51e720eed7e")
	require.NoError(t, err)
	data = append(data, data2)

	data3, err := hex.DecodeString("66d04cf478349ebddbb0b59c223f4fbf51de8e5ff8e0011629ae80b994be81e3")
	require.NoError(t, err)
	data = append(data, data3)

	data4, err := hex.DecodeString("8accdd94e130c81867f487e694cbe08cef78a7cfc3dff03301555d5d9775fe77")
	require.NoError(t, err)
	data = append(data, data4)

	data5, err := hex.DecodeString("494204064cd68938d872e2d0c20f0b636b8503eba71ec9d86bd4b3195c039aad")
	require.NoError(t, err)
	data = append(data, data5)

	data6, err := hex.DecodeString("958fac8ad97f003ce03ecfdb5b379bd00b52ed497bc4dd58ac380229dbef48e6")
	require.NoError(t, err)
	data = append(data, data6)

	data7, err := hex.DecodeString("20be72b01f7571aaa78426fb50484fa953557abf80f58a6245da46c14f7ce55e")
	require.NoError(t, err)
	data = append(data, data7)

	data8, err := hex.DecodeString("ee9943eba1fc7fa50cba34a9a055ba1b4a072d31c2a8e4198573403e31fcbe06")
	require.NoError(t, err)
	data = append(data, data8)

	data9, err := hex.DecodeString("abac167cd66d014e969e25f6788e297a0f02c5c944f032aed2ee0e45924e2238")
	require.NoError(t, err)
	data = append(data, data9)

	data10, err := hex.DecodeString("8888e7fe141d8d8a3489ee12ab3b778ce45dd7b4b56fb33884123b63101376ec")
	require.NoError(t, err)
	data = append(data, data10)

	data11, err := hex.DecodeString("15d1e8525f76b9d59e4c1cfa192a14df79843a677bd3d20b6db0ba7d688ef57b")
	require.NoError(t, err)
	data = append(data, data11)

	data12, err := hex.DecodeString("e6d05c744ed05106db38e850ce76321b8a47b59bcbf4141a5ff497ea4c8ce796")
	require.NoError(t, err)
	data = append(data, data12)
	return data
}

func TestVerify(t *testing.T) {
	data := getTestData(t)
	var root, proofs = ProofsFromByteSlices(data)

	fmt.Printf("Root = %x\n", root)
	fmt.Printf("Verify = %v\n", proofs[0].Verify(root, data[0]))

	proof := proofs[0]
	bz, err := json.Marshal(proof.TmProof())
	require.NoError(t, err)

	var tmProof tmmerkle.Proof
	err = json.Unmarshal(bz, &tmProof)
	require.NoError(t, err)

	pf, err := ProofOf(tmProof)
	require.NoError(t, err)
	fmt.Printf("Verify = %v\n", pf.Verify(root, data[0]))

}
