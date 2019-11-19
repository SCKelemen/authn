package random

import "crypto/rand"

type RandomStrategy struct{}

func (rs RandomStrategy) GenerateSalt() []byte {
	length := 32
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil
	}

	return bytes
}
