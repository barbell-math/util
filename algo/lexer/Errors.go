package lexer

import "errors"

var RegexSyntaxError=errors.New("The supplied regular expression has a syntax error.")
var RegexInvalidEscapeSequence=errors.New(
	"The character following the '\\' character was not valid.\n"+
	"Valid escape sequences are: \\l (lambda char), \\( (open paren), "+
	"\\) (close paren), \\\\ (back slash), \\* (Kleene operator), and "+
	"\\| (or operator)",
)
var RegexInvalidCharacter=errors.New("A character value was supplied that is not supported.")
var RegexInbalancedParens=errors.New("There were not enough closing parens to match with the opening parens.")
