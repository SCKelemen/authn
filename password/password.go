package password

import (
	"encoding/hex"
	"salt"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) []byte {
	salts := make(chan []byte)
	go func() { salts <- salt.GenerateSalt() }()
	bytes, _ := hex.DecodeString(password)
	return argon2.IDKey(bytes, <-salts, 3, 32*1024, 4, 32) //hmac.New(argon2.New) //hhmac(password, <-salts)
}

// HMAC follows the HMAC construction for combining Keys with data
// Hash(( K XOR opad ) Hash(( K XOR ipad) M ))
// Here, the key, or K, is the password, and the data/message, M, is the salt
func hhmac(password, iv string) []byte {
	return []byte(password)
}

// OPAD generates the out pad.
// 5c5c5c repeated for the length of the block
func opad(blocksize uint16) []byte {
	return []byte("password")
}

// IPAD generates the inner padding.
// 363636 repeated for the length of the block
// and terminated with 00
func ipad(blocksize uint16) []byte {
	return []byte("password")
}
