package lexer

import (
	"fmt"
	"regexp"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/dataStruct"
	"github.com/barbell-math/util/dataStruct/types"
)

type MatchStatus byte
const (
    NoMatch MatchStatus=0
    Match MatchStatus=1<<iota
    PossibleMatch
)

type TokenMatcher struct {
    re *regexp.Regexp
    endLine int
    endChar int
    totalChars int
    match MatchStatus
}

func NewTokenMatcher(re string) (TokenMatcher,error) {
    var err error
    rv:=TokenMatcher{};
    rv.re,err=regexp.Compile(re)
    return rv,err
}

func (t *TokenMatcher)incrementCounts(c byte){
    if c=='\n' {
        t.endLine+=1;
    } else {
        t.endChar+=1;
    }
    t.totalChars++
}


type TokenType int
const (
    UnknownToken TokenType=-1
)

type Token struct {
    Data string
    Type TokenType
    Line,Char int
}

type Parser struct {
    tokens map[TokenType]TokenMatcher
    line, char int
    leftoverChars dataStruct.Deque[byte]
    curToken []byte
    matchingToken TokenType
}

func NewParser() Parser {
    return Parser{
        tokens: make(map[TokenType]TokenMatcher,0),
    }
}

func (p Parser)TokenStream(charStream iter.Iter[byte]) iter.Iter[Token] {
    return iter.Map[byte,Token](
        charStream.Inject(func(idx int, val byte) (byte, bool) {
            if v,err:=p.leftoverChars.PopFront(); err==nil {
                return v,true
            }
            return byte(0),false
        }).Next(func(
            index int, iterChar byte, status iter.IteratorFeedback,
        ) (iter.IteratorFeedback, byte, error) {
            if status==iter.Break {
                var err error
                if p.leftoverChars.Length()>0 || len(p.curToken)>0 {
                    err=SyntaxError(fmt.Sprintf(
                        "Line: %d Char: %d | Enf of stream reached without completing last token.",
                        p.line,p.char,
                    ))
                }
                return iter.Break,byte(0),err
            }
            p.curToken=append(p.curToken, iterChar)
            num,err:=p.runMatchers(iterChar)
            return p.interpretResult(iterChar,num,err)
        }),
        func(index int, val byte) (Token, error) {
            return Token{
                Data: string(p.curToken),
                Line: p.line,
                Char: p.char,
                Type: p.matchingToken,
            }, nil
        },
    )
}

func (p Parser)runMatchers(iterChar byte) (int,error) {
    return iter.MapVals[TokenType,TokenMatcher](
        p.tokens,
    ).Filter(func(index int, val TokenMatcher) bool {
        if val.match==NoMatch {
            return false
        }
        // get match status: Match, Possible, NoMatch
        val.match=true  //HOW TF DO I DO THIS??? That aint it
        // return val.re.MatchString(p.curToken.String())
        return val.match==PossibleMatch || val.match==Match
    }).Next(func(
        index int, val TokenMatcher, status iter.IteratorFeedback,
    ) (iter.IteratorFeedback, TokenMatcher, error) {
        if status==iter.Break {
            return iter.Break, TokenMatcher{}, nil;
        }
        val.incrementCounts(iterChar)
        return iter.Continue, val, nil;
    }).Count();
}

func (p Parser)interpretResult(
    iterChar byte, 
    curMatches int, 
    err error,
) (iter.IteratorFeedback, byte, error) {
    if err!=nil {
        return iter.Break, byte(0), err
    }
    if curMatches>0 {  // Multiple regexp matches, continue with next char
        return iter.Iterate, byte(0), nil
    }
    _,maxMatchCount:=p.getMaxMatch()
    if maxMatchCount==1 {   // Single regexp matched, return token
        p.setTokenAndLeftovers()
        p.resetTokenState()
        return iter.Continue, iterChar, nil
    } else {    // Multiple or no regexp matched, return error
        return iter.Break, iterChar, SyntaxError(fmt.Sprintf(
            "Line: %d Char %d | No regular expression matched the supplied sequence.",
            p.line, p.char,
        ))
    }
}

func (p Parser)getMaxMatch() (int,int) {
    maxMatch:=0
    maxMatchCount:=0
    for k,v:=range(p.tokens) {
        if v.totalChars>maxMatch {
            maxMatch=v.totalChars
            maxMatchCount=1
            p.matchingToken=k
        } else if v.totalChars==maxMatch {
            maxMatchCount+=1
        }
    }
    return maxMatch,maxMatchCount
}

func (p Parser)setTokenAndLeftovers() {
    curTok:=[]byte(p.tokens[p.matchingToken].re.FindString(string(p.curToken)))
    p.leftoverChars.SetCapacity(
        p.leftoverChars.Length()+len(p.curToken)-len(curTok),
    )
    for i:=len(p.curToken)-1; i>=len(curTok); i-- {
    // for i:=len(curTok); i<len(p.curToken); i++ {
        p.leftoverChars.PushFront(p.curToken[i])
    }
    p.curToken=curTok
    p.line=p.tokens[p.matchingToken].endLine
    p.char=p.tokens[p.matchingToken].endChar
}

func (p Parser)resetTokenState() {
    for _,t:=range(p.tokens) {
        t.endLine=p.line
        t.endChar=p.char
        t.totalChars=0
        t.match=Match
    }
}
