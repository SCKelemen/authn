package salt

type SaltStrategy interface {
	GenerateSalt() []byte
}


