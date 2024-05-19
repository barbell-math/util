package lexer

import (
	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/customerr"
)

type (
    Regex string
    MatchStatus byte
)

const (
	invalid byte=iota
    lambdaChar
	lParenChar
	rParenChar
	backSlashChar
	starChar
	barChar
)

const (
    NoMatch MatchStatus=0
    Match MatchStatus=1<<iota
    PossibleMatch
)

var (
	escapeChars map[byte]struct{}=map[byte]struct{}{
		'\\': {},
		'(': {},
		')': {},
		'*': {},
		'|': {},
		'_': {},
	}

	specialCharEncoding map[byte]byte=map[byte]byte{
		'(': lParenChar,
		')': rParenChar,
		'*': starChar,
		'|': barChar,
		'_': lambdaChar,
	}
)

func (r *Regex)Compile() (DFA,error) {
	r.buildNFA(r.toTokenStream(),0)
    // build NFA, catch errors in process
    // transition NFA to DFA
	return DFA{},nil
}

// Valid characters: ASCII table 32-126
// \_ = _, \(=(, \)=), \\=\, \*=*, \|=|
func (r *Regex)toTokenStream() iter.Iter[byte] {
	slashFound:=false
	return iter.StrElems(string(*r)).Next(
		func(
			index int,
			val byte,
			status iter.IteratorFeedback,
		) (iter.IteratorFeedback, byte, error) {
			if val<32 && val>126 {
				return iter.Break,0,customerr.AppendError(
					RegexSyntaxError,
					customerr.Wrap(
						RegexInvalidCharacter,
						"Character: %c (%d)",
						val,val,
					),
				)
			}

			if slashFound {
				slashFound=false
				if _,ok:=escapeChars[val]; ok {
					return iter.Continue,val,nil
				} else {
					return iter.Break,0,customerr.AppendError(
						RegexSyntaxError,
						customerr.Wrap(
							RegexInvalidEscapeSequence,
							"Escape sequence: \\%c",
							val,
						),
					)
				}
			}

			if val=='\\' && !slashFound {
				slashFound=true
				return iter.Iterate,0,nil
			}
			if v,ok:=specialCharEncoding[val]; ok {
				return iter.Continue,v,nil
			}
			return iter.Continue,val,nil
		},
	)
}

func (r *Regex)buildNFA(tokens iter.Iter[byte], curId nfaID) (NFA, byte, error) {
	// prevChar:=invalid
	// curNFA:=NewNFA()
	// for val,err,cont:=tokens.PullOne(); err==nil && cont; val,err,cont=tokens.PullOne() {
	// 	if val==starChar {
	// 		if prevChar!=rParenChar {
	// 			// curNFA.ApplyKleene()
	// 		} else {
	// 			// curNFA.ApplyKleeneToLastChar()
	// 		}
	// 	} else if val==barChar {
	// 		subNFA,_,err:=r.buildNFA(tokens,curId)
	// 		if err!=nil {
	// 			return NFA{},0,err
	// 		}
	// 		// curNFA.AddBranch(subNFA)
	// 	} else if val==lParenChar {
	// 		subNFA,lastChar,err:=r.buildNFA(tokens,curId)
	// 		if err!=nil {
	// 			return NFA{},0,err
	// 		}
	// 		if lastChar!=rParenChar {
	// 			return NFA{},0,customerr.AppendError(
	// 				RegexSyntaxError,
	// 				RegexInbalancedParens,
	// 			)
	// 		}
	// 		// curNFA.Append(subNFA)
	// 		// absorb subNFA
	// 	} else {
	// 		// curNFA.AppendTransition(char)
	// 	}
	// }
	return NFA{},0,nil
}
