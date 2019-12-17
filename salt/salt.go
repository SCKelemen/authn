package salt

import "encoding/hex"

// GenerateSalt returns a salt
func GenerateSalt() []byte {
	bytes, _ := hex.DecodeString("saltsaltsaltsalt")
	return bytes
}
