package tests

//go:generate ../../bin/enum -type=tokenType -package=tests

type (
	//gen:enum unknownValue unknownTokenType
	//gen:enum default unknownTokenType
	tokenType int
)

const (
	//gen:enum string unknownTokenType
	unknownTokenType tokenType = iota
	// Represents a short flag token value
	//gen:enum string shortFlagToken
	shortFlagToken
	// Represents a long flag token value
	//gen:enum string longFlagToken
	longFlagToken
	// Represents a argument value that would be attached to a token
	//gen:enum string valueToken
	valueToken
)
