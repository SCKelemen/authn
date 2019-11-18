package sha512strategy

import (
	"crypto/sha512"
	"crypto/subtle"
)

// Sha512Strategy implements the HashAlgorithm Interface for SHA512. This is considered legacy.
type Sha512Strategy struct{}

// Hash implements HashAlgorithm.Hash (interface) for Sha512. The Salt is not used.
func (strategy Sha512Strategy) Hash(password []byte, salt []byte) (bool, []byte) {
	hash := sha512.Sum512(password)
	return true, hash[:]
}

// Verify implement HashAlgorithm.Verify for Sha512. The Salt is not used.
func (strategy Sha512Strategy) Verify(known []byte, given []byte, salt []byte) bool {
	hash := sha512.Sum512(given)
	if subtle.ConstantTimeCompare(hash[:], known) == 1 {
		return true
	}
	return false
}
