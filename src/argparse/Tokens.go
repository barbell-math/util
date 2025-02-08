package argparse

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/barbell-math/util/src/container/basic"
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/iter"
)

//go:generate ../../bin/structBaseWidget -type=token

type (
	// Represents an sequence of strings that can be translated into a sequence
	// of tokens.
	ArgvIter    iter.Iter[string]
	tokenIter   iter.Iter[token]
	argValPairs iter.Iter[basic.Pair[*arg, string]]

	token struct {
		//gen:structBaseWidget identity
		//gen:structBaseWidget baseTypeWidget widgets.BuiltinString
		//gen:structBaseWidget widgetPackage github.com/barbell-math/util/src/widgets
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

func (a argValPairs) ToIter() iter.Iter[basic.Pair[*arg, string]] {
	return iter.Iter[basic.Pair[*arg, string]](a)
}

// Translates the sequence of strings into tokens. No validation is done to
// check that the stream of tokens is valid.
func (a ArgvIter) ToTokens() tokenIter {
	expectingConfig := false
	tokens := []token{}

	return func(f iter.IteratorFeedback) (token, error, bool) {
		if f == iter.Break {
			tokens = []token{}
			return token{}, nil, false
		}

	iterStart:
		if len(tokens) > 0 {
			rv := tokens[len(tokens)-1]
			tokens = tokens[:len(tokens)-1]
			return rv, nil, true
		}

		s, err, cont := a(f)
		if err != nil || !cont {
			return token{}, err, cont
		}

		if expectingConfig {
			configTokens, err := generateConfigFileTokens(s)
			if err != nil {
				return token{}, err, false
			}
			for i := len(configTokens) - 1; i >= 0; i-- {
				tokens = append(tokens, configTokens[i])
			}
			expectingConfig = false
			goto iterStart
		} else if regexes[configEqualsFileFlag].MatchString(s) {
			parts := strings.Split(s, "=")
			configTokens, err := generateConfigFileTokens(parts[1])
			if err != nil {
				return token{}, err, false
			}
			for i := len(configTokens) - 1; i >= 0; i-- {
				tokens = append(tokens, configTokens[i])
			}
			goto iterStart
		} else if regexes[configSpaceFileFlag].MatchString(s) {
			expectingConfig = true
			goto iterStart
		} else if regexes[longEqualsFlag].MatchString(s) {
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
			for i := len(s) - 1; i >= 2; i-- {
				iterChar := string(s[i])
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

// Responsible for parsing a parser config file and returning the set of tokens
// that are represented in the file.
func generateConfigFileTokens(file string) ([]token, error) {
	wrapErr := func(err error) error {
		return customerr.AppendError(
			customerr.Wrap(ParserConfigFileErr, "File: %s", file),
			err,
		)
	}

	rv := []token{}

	f, err := os.Open(file)
	if err != nil {
		return rv, wrapErr(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	cntr := 0
	names := []string{}
	curName := ""
	for scanner.Scan() {
		trimmedLine := strings.TrimSpace(scanner.Text())

		if (len(trimmedLine) >= 2 && trimmedLine[0:2] == "//") || len(trimmedLine) == 0 {
			continue
		}

		splitLine := strings.SplitN(trimmedLine, " ", 2)
		for i, l := range splitLine {
			splitLine[i] = strings.TrimSpace(l)
		}

		if len(splitLine) != 2 && (len(splitLine) == 1 && splitLine[0] != "}") {
			return rv, wrapErr(customerr.WrapValueList(
				ParserConfigFileSyntaxErr,
				fmt.Sprintf("Syntax error on line %d. Expected one of the following formats", cntr+1),
				[]customerr.WrapListVal{
					customerr.WrapListVal{
						ItemName: "Argument definition",
						Item:     "'<name> <value>'",
					},
					customerr.WrapListVal{
						ItemName: "Argument group definition open",
						Item:     "'<name> {'",
					},
					customerr.WrapListVal{
						ItemName: "Argument group definition close",
						Item:     "'}'",
					},
					customerr.WrapListVal{
						ItemName: "Got",
						Item:     splitLine,
					},
					customerr.WrapListVal{
						ItemName: "Note",
						Item:     "leading white space on a line and white space inside a value are ignored but otherwise matters",
					},
				},
			))
		}

		if splitLine[0] == "}" {
			if len(names) == 0 {
				return rv, wrapErr(customerr.Wrap(
					ParserConfigFileSyntaxErr,
					"To many closing brackets. Noticed on line: %d",
					cntr+1,
				))
			}
			names = names[0 : len(names)-1]
			curName = strings.Join(names, "")
		} else if splitLine[1] == "{" {
			names = append(names, splitLine[0])
			curName = strings.Join(names, "")
		} else {
			rv = append(
				rv,
				token{value: curName + splitLine[0], _type: longFlagToken},
				token{value: splitLine[1], _type: valueToken},
			)
		}

		cntr++
	}
	if len(names) > 0 {
		return rv, wrapErr(customerr.Wrap(
			ParserConfigFileSyntaxErr,
			"Not enough closing brackets. Noticed at EOF.",
		))
	}

	return rv, nil
}

// Takes a sequence of tokens and turns it into a sequence of argument -> value
// pairs. This validates that the sequence of tokens is a valid sequence given
// the type of each token and the placement of each token.
func (t tokenIter) toArgValPairs(p *Parser) argValPairs {
	multiValue := false
	var multiValueToken *arg = nil

	getExpectedValue := func(f iter.IteratorFeedback) (string, error) {
		iterToken, err, cont := t(f)
		if err != nil {
			return "", err
		}
		if !cont {
			return "", customerr.AppendError(
				EndOfTokenStreamErr,
				ExpectedValueErr,
			)
		}
		if iterToken._type != valueToken {
			return "", customerr.Wrap(
				ExpectedValueErr,
				"Got: '%s' (%s)", iterToken.value, iterToken._type,
			)
		}
		return iterToken.value, nil
	}

	return func(f iter.IteratorFeedback) (basic.Pair[*arg, string], error, bool) {
		if f == iter.Break {
			return basic.Pair[*arg, string]{}, nil, false
		}

		rv := basic.Pair[*arg, string]{}
		iterToken, err, cont := token{}, error(nil), true

		iterToken, err, cont = t(f)
		if err != nil || !cont {
			return basic.Pair[*arg, string]{}, err, cont
		}

		switch iterToken._type {
		case shortFlagToken:
			if rv.A, err = p.getShortArg(iterToken.value[0]); err != nil {
				return rv, err, false
			}
			multiValue = false
		case longFlagToken:
			if rv.A, err = p.getLongArg(iterToken.value); err != nil {
				return rv, err, false
			}
			multiValue = false
		case valueToken:
			if !multiValue {
				return rv, customerr.Wrap(
					ExpectedArgumentErr,
					"Got: '%s' (%s)", iterToken.value, iterToken._type,
				), false
			} else {
				rv.A = multiValueToken
				rv.B = iterToken.value
			}
		default:
			return rv, customerr.Wrap(
				InvalidTokenType, "'%s' (%s)", iterToken.value, iterToken._type,
			), false
		}

		switch rv.A.argType {
		case ValueArgType:
			if rv.B, err = getExpectedValue(f); err != nil {
				return rv, err, false
			}
		case MultiValueArgType:
			if !multiValue {
				multiValueToken = rv.A
				multiValue = true
				if rv.B, err = getExpectedValue(f); err != nil {
					return rv, err, false
				}
			}
		case FlagArgType:
			break
		case MultiFlagArgType:
			break
		default:
			return rv, customerr.Wrap(
				InvalidArgType, "'%s' (%s)", rv.A.longFlag, rv.A.argType,
			), false
		}

		return rv, nil, true
	}
}
