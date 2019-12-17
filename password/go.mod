module github.com/SCKelemen/authn/password

go 1.13

require (
	salt v0.0.0
	hash v0.0.0

)

replace salt v0.0.0 => ./salt

replace hash v0.0.0 => ./hash
