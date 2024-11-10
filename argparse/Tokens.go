package argparse

import (
	"strings"

	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/iter"
)

//go:generate ../bin/structBaseWidget -type=token

type (
	ArgvIter  iter.Iter[string]
	tokenIter iter.Iter[token]
	argValPairs iter.Iter[basic.Pair[*Arg, string]]

	token struct {
		//gen:structBaseWidget identity
		//gen:structBaseWidget baseTypeWidget widgets.BuiltinString
		//gen:structBaseWidget widgetPackage github.com/barbell-math/util/widgets
		value string
		//gen:structBaseWidget identity
		//gen:structBaseWidget baseTypeWidget *tokenType
		//gen:structBaseWidget widgetPackage .
		_type tokenType
	}
)

func ArgvIterFromSlice(argv []string) ArgvIter {
	return ArgvIter(iter.SliceElems[string](argv))
}

func (a ArgvIter) ToIter() iter.Iter[string] {
	return iter.Iter[string](a)
}

func (t tokenIter) ToIter() iter.Iter[token] {
	return iter.Iter[token](t)
}

func (a argValPairs) ToIter() iter.Iter[basic.Pair[*Arg, string]] {
	return iter.Iter[basic.Pair[*Arg, string]](a)
}

func (a ArgvIter) ToTokens() tokenIter {
	tokens := []token{}

	return func(f iter.IteratorFeedback) (token, error, bool) {
		if f == iter.Break {
			tokens = []token{}
			return token{}, nil, false
		}

		if len(tokens) > 0 {
			rv := tokens[len(tokens)-1]
			tokens = tokens[:len(tokens)-1]
			return rv, nil, true
		}

		s, err, cont := a(f)
		if err != nil || !cont {
			return token{}, err, cont
		}

		if regexes[longEqualsFlag].MatchString(s) {
			parts := strings.Split(s, "=")
			tokens = append(tokens, token{
				value: strings.TrimSpace(parts[1]),
				_type: valueToken,
			})
			return token{
				value: strings.TrimSpace(parts[0][2:]),
				_type: longFlagToken,
			}, nil, true
		} else if regexes[longSpaceFlag].MatchString(s) {
			return token{
				value: s[2:],
				_type: longFlagToken,
			}, nil, true
		} else if regexes[shortEqualsFlag].MatchString(s) {
			// short equals flag tokens can never be combined
			parts := strings.Split(s, "=")
			tokens = append(tokens, token{
				value: strings.TrimSpace(parts[1]),
				_type: valueToken,
			})
			return token{
				value: strings.TrimSpace(parts[0][1:]),
				_type: shortFlagToken,
			}, nil, true
		} else if regexes[shortSpaceFlag].MatchString(s) {
			// short tokens can be combined sometimes so add them all
			// individually
			for i:=len(s)-1; i>=2; i-- {
				iterChar:=string(s[i])
				tokens = append(tokens, token{
					value: iterChar,
					_type: shortFlagToken,
				})
			}
			return token{
				value: s[1:2],
				_type: shortFlagToken,
			}, nil, true
		} else {
			return token{
				value: s,
				_type: valueToken,
			}, nil, true
		}
	}
}

func (t tokenIter) toArgValPairs(p *Parser) argValPairs {
	return func(f iter.IteratorFeedback) (basic.Pair[*Arg, string], error, bool) {
		if f==iter.Break {
			return basic.Pair[*Arg, string]{}, nil, false
		}
		
		rv:=basic.Pair[*Arg, string]{}
		iterToken, err, cont:=token{}, error(nil), true

		iterToken, err, cont=t(f)
		if err!=nil || !cont {
			return basic.Pair[*Arg, string]{}, err, cont
		}

		switch iterToken._type {
		case shortFlagToken:
			if rv.A, err=p.getShortArg(iterToken.value[0]); err!=nil {
				return rv, err, false
			}
		case longFlagToken:
			if rv.A, err=p.getLongArg(iterToken.value); err!=nil {
				return rv, err, false
			}
		case valueToken:
			return rv, customerr.Wrap(
				ExpectedArgumentErr,
				"Got: '%s' (%s)", iterToken.value, iterToken._type,
			), false
		default:
			return rv, customerr.Wrap(
				InvalidTokenType, "'%s' (%s)", iterToken.value, iterToken._type,
			), false
		}

		switch rv.A.argType {
		case ValueArgType:
			iterToken, err, cont=t(f)
			if err!=nil {
				return rv, err, cont
			}
			if !cont {
				return rv, customerr.AppendError(
					EndOfTokenStreamErr,
					ExpectedValueErr,
				), false
			}
			if iterToken._type!=valueToken {
				return rv, customerr.Wrap(
					ExpectedValueErr,
					"Got: '%s' (%s)", iterToken.value, iterToken._type,
				), false
			}
			rv.B=iterToken.value
		case FlagArgType: break
		case MultiFlagArgType: break
		default:
			return rv, customerr.Wrap(
				InvalidArgType, "'%s' (%s)", rv.A.longFlag, rv.A.argType,
			), false
		}

		return rv, nil, true
	}
}
