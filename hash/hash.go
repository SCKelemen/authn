package hash

type HashAlg uint8

const (
	INVALID HashAlg = iota

	_beginValid
	SHA512
	BCRYPT
	SCRYPT
	PBKDF2
	_endValid
)

type HashAlgorithm interface {
	Hash(password []byte, salt []byte) (bool, []byte)
	Verify(known []byte, given []byte, salt []byte) bool
}
