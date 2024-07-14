package lexer

type MatchStatus byte

const (
	NoMatch MatchStatus = 0
	Match   MatchStatus = 1 << iota
	PossibleMatch
)

const (
	validCharLowerBound byte = 0
	validCharUpperBound byte = 127
	invalid             byte = iota + 127
	lambdaChar
	lParenChar
	rParenChar
	backSlashChar
	starChar
	barChar
)

var (
	escapeChars map[byte]struct{} = map[byte]struct{}{
		'\\': {},
		'(':  {},
		')':  {},
		'*':  {},
		'|':  {},
		'_':  {},
	}

	specialCharEncoding map[byte]byte = map[byte]byte{
		'(': lParenChar,
		')': rParenChar,
		'*': starChar,
		'|': barChar,
		'_': lambdaChar,
	}
)
