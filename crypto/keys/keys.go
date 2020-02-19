package keys

// SigningAlgo defines an algorithm to derive key-pairs which can be used for cryptographic signing.
type SigningAlgo string

const (
	// MultiAlgo implies that a pubkey is a multisignature
	MultiAlgo = SigningAlgo("multi")
	// Sm2 represents the Sm2 signature system.
	Sm2 = SigningAlgo("sm2")
	// Secp256k1 uses the Bitcoin secp256k1 ECDSA parameters.
	Secp256k1 = SigningAlgo("secp256k1")
	// Sr25519 represents the Sr25519 signature system.
	Sr25519 = SigningAlgo("sr25519")
)
