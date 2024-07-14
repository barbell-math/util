package lexer

import (
	"fmt"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/customerr"
)

type (
	Regex string
)

func (r *Regex) Compile() (DFA, error) {
	tokens := r.toTokenStream()
	nfa, _, err := r.buildNFA(0, tokens, 0)
	err = customerr.AppendError(err, tokens.Stop())
	if err != nil {
		return DFA{}, err
	}
	fmt.Println(nfa)
	// build NFA, catch errors in process
	// transition NFA to DFA
	return DFA{}, nil
}

func (r *Regex) toTokenStream() iter.Iter[byte] {
	slashFound := false
	return iter.StrElems(string(*r)).Next(
		func(
			index int,
			val byte,
			status iter.IteratorFeedback,
		) (iter.IteratorFeedback, byte, error) {
			if val < validCharLowerBound || val > validCharUpperBound {
				return iter.Break, 0, customerr.AppendError(
					RegexSyntaxError,
					customerr.Wrap(
						RegexInvalidCharacter,
						"Character: %c (%d)",
						val, val,
					),
				)
			}

			if slashFound {
				slashFound = false
				if _, ok := escapeChars[val]; ok {
					return iter.Continue, val, nil
				} else {
					return iter.Break, 0, customerr.AppendError(
						RegexSyntaxError,
						customerr.Wrap(
							RegexInvalidEscapeSequence,
							"Escape sequence: \\%c",
							val,
						),
					)
				}
			}

			if val == '\\' && !slashFound {
				slashFound = true
				return iter.Iterate, 0, nil
			}
			if v, ok := specialCharEncoding[val]; ok {
				return iter.Continue, v, nil
			}
			return iter.Continue, val, nil
		},
	)
}

func (r *Regex) buildNFA(
	lastChar byte,
	tokens iter.Iter[byte],
	curId nfaID,
) (NFA, byte, error) {
	curNFA := NewNFA()
	val, err, cont := tokens.PullOne()
	for ; err == nil && cont; val, err, cont = tokens.PullOne() {
		if val == rParenChar {
			if lastChar != lParenChar {
				return NFA{}, 0, customerr.AppendError(
					RegexSyntaxError,
					RegexInbalancedParens,
				)
			}
			return curNFA, rParenChar, nil
		} else if val == starChar {
			curNFA.ApplyKleene()
		} else if val == barChar {
			subNFA, _, err := r.buildNFA(val, tokens, curId)
			if err != nil {
				return NFA{}, 0, err
			}
			curNFA.AddBranch(subNFA)
		} else if val == lParenChar {
			subNFA, lastChar, err := r.buildNFA(val, tokens, curId)
			if err != nil {
				return NFA{}, 0, err
			}
			if lastChar != rParenChar {
				return NFA{}, 0, customerr.AppendError(
					RegexSyntaxError,
					RegexInbalancedParens,
				)
			}
			curNFA.AppendNFA(subNFA)
		} else {
			curNFA.AppendTransition(val)
		}
	}
	return curNFA, 0, err
}
