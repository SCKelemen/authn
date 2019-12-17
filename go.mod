module github.com/SCKelemen/authn

go 1.13

require (
	exclaim v0.0.0
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413
	hash v0.0.0
	password v0.0.0
	salt v0.0.0
	yell v0.0.0
)

replace exclaim v0.0.0 => ./exclaim

replace yell v0.0.0 => ./yell

replace password v0.0.0 => ./password

replace salt v0.0.0 => ./salt

replace hash v0.0.0 => ./hash
