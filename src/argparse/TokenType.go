package argparse

//go:generate ../../bin/enum -type=tokenType -package=argparse
//go:generate ../../bin/passThroughWidget -type=tokenType

type (
	//gen:enum unknownValue unknownTokenType
	//gen:enum default unknownTokenType
	//gen:passThroughWidget widgetType Base
	//gen:passThroughWidget baseTypeWidget widgets.BuiltinInt
	//gen:passThroughWidget widgetPackage github.com/barbell-math/util/src/widgets
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

var (
	flagTokens = map[tokenType]struct{}{
		shortFlagToken: {},
		longFlagToken:  {},
	}
	argumentTokens = map[tokenType]struct{}{
		valueToken: {},
	}
)
