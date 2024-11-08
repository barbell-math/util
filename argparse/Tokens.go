package argparse

import (
	"strings"

	"github.com/barbell-math/util/iter"
)

//go:generate ../bin/structBaseWidget -type=token

type (
	ArgvIter iter.Iter[string]
	tokenIter iter.Iter[token]

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

func (a ArgvIter) ToTokens() tokenIter {
	tokens:=[]token{}

	return func(f iter.IteratorFeedback) (token, error, bool) {
		if f == iter.Break {
			tokens=[]token{}
			return token{}, nil, false
		}

		if len(tokens)>0 {
			rv:=tokens[len(tokens)-1]
			tokens=tokens[:len(tokens)-1]
			return rv, nil, true
		}

		s, err, cont:=a(f)
		if err!=nil || !cont {
			return token{}, err, cont
		}
		
		if regexes[longEqualsFlag].MatchString(s) {
			parts:=strings.Split(s, "=")
			tokens = append(tokens, token{
				value: strings.TrimSpace(parts[1]),
				_type: argumentToken,
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
		} else if regexes[shortFlag].MatchString(s) {
			return token{
				value: s[1:],
				_type: shortFlagToken,
			}, nil, true
		} else {
			return token{
				value: s,
				_type: argumentToken,
			}, nil, true
		}
	}
}

// func (t tokenIter) ToPairs() iter.Iter[staticContainers.Vector[token]] {
// 	q, _ := containers.NewCircularBuffer[token, *token](2)
// 	tokenPairs:=containers.SteppingWindow[token](t.ToIter(), &q)
// 	return func(f iter.IteratorFeedback) (staticContainers.Vector[token], error, bool) {
// 		if f == iter.Break {
// 			q.Clear()
// 			return &q, nil, false
// 		}
// 
// 		vec, err, cont:=tokenPairs(f)
// 		if err!=nil || !cont {
// 			return &q, nil, false
// 		}
// 
// 		val, err:=vec.GetPntr(0)
// 		if err!=nil {
// 			return &q, err, false
// 		}
// 		if _, ok:=flagTokens[val._type]; !ok {
// 			return &q, fmt.Errorf(
// 				"%w: Got: '%s'(%s)",
// 				ExpectedFlag, val.value, val._type,
// 			), false
// 		}
// 
// 		val, err=vec.GetPntr(1)
// 		if err!=nil {
// 			return &q, err, false
// 		}
// 		if _, ok:=argumentTokens[val._type]; !ok {
// 			return &q, fmt.Errorf(
// 				"%w: Got: '%s'(%s)",
// 				ExpectedFlag, val.value, val._type,
// 			), false
// 		}
// 
// 		return vec, nil, true
// 	}
// }
