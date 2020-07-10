package merkle

import (
	"hash"
)

func leafHash(leaf []byte) []byte {
	return genHash().Sum(leaf)
}

func innerHash(left []byte, right []byte) []byte {
	return genHash().Sum(append(left, right...))
}

func genHash() hash.Hash {
	hasher.Reset()
	return hasher
}