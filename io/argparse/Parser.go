package argparse

import "regexp"

type tokenMatch string

func (t tokenMatch)matches(s string) bool {
    match, _:=regexp.MatchString(string(t),s)
    return match
}

type tokenType int
const (
    ShortNameToken tokenType=iota
    LongNameToken tokenType=iota
    ArgumentToken tokenType=iota
)

type token struct {
    data string
    _type tokenType
}

type tokenStream map[tokenType]tokenMatch

var parser tokenStream=map[tokenType]tokenMatch{
    ShortNameToken: tokenMatch("-?"),
    LongNameToken: tokenMatch("--*"),
    ArgumentToken: tokenMatch("*"),
}

func (t *tokenStream)tokenStream(tokenSequence []string, s chan<- token) error {
    for i,t:=range(tokenSequence) {
        tType:=ArgumentToken
        if p.isArg(t) {
            tType=CommandToken
        }
        s <- token {
            data: t,
            index: i,
            _type: tType,
        }
    }
    return nil
}

func (p *Parser)isShortNameToken(token string) bool {
    return (
        (len(token)==2 && token[0]==TOKEN_LEADER && token[1]!=TOKEN_LEADER) ||
        (len(token)>2 && token[0]==TOKEN_LEADER && token[1]==TOKEN_LEADER))
}
