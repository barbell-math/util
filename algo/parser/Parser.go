package parser

import (
	"regexp"
	"strings"

	"github.com/barbell-math/util/algo/iter"
)

type TokenMatcher string

func (t TokenMatcher)matches(s string) bool {
    match, _:=regexp.MatchString(string(t),s)
    return match
}

type TokenType int
const (
    UnknownToken TokenType=-1
)

type Token struct {
    data string
    _type TokenType
}

type Parser map[TokenType]TokenMatcher

func (t *Parser)TokenStream(charStream iter.Iter[byte]) iter.Iter[Token] {
    var curToken strings.Builder
    var updatedToken strings.Builder
    index:=0
    for t:=range(charStream) {
        updatedToken.WriteByte(t)
        // noMatch,err:=iter.
        for _,m:=range(t) {
            
        }
        tType:=UnknownToken

        s <- token {
            data: t,
            index: i,
            _type: tType,
        }
    }
    return nil
}
