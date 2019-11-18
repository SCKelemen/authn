package bcryptstrategy

import (
	"golang.org/x/crypto/bcrypt"
)

// BCryptStrategy implements the HashAlgorithm Interface for SHA512. This is considered legacy.
type BCryptStrategy struct{}

// Hash implements HashAlgorithm.Hash (interface) for Sha512. The Salt is not used.
func (strategy BCryptStrategy) Hash(password []byte, salt []byte) (bool, []byte) {
	// takes password bytes, and cost factor
	// not really sure how this salt is used. Perhaps I need to append it to the password beforehand
	hash, err := bcrypt.GenerateFromPassword(password, 14)
	if err != nil {
		return false, nil
	}
	return true, hash[:]
}

// Verify implement HashAlgorithm.Verify for Sha512. The Salt is not used.
func (strategy BCryptStrategy) Verify(known []byte, given []byte, salt []byte) bool {
	// this also performs a constant time compare, but only returns an error.
	return bcrypt.CompareHashAndPassword(known, given) == nil
}
