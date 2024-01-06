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

type Lexer struct {
    tokens map[TokenType]TokenMatcher
    line, char int
    leftoverChars dataStruct.Deque[byte]
    curToken []byte
    matchingToken TokenType
}

func NewParser() Lexer {
    return Lexer{
        tokens: make(map[TokenType]TokenMatcher,0),
    }
}

func (l *Lexer)TokenStream(charStream iter.Iter[byte]) iter.Iter[Token] {
    return iter.Map[byte,Token](
        charStream.Inject(func(idx int, val byte) (byte, bool) {
            if v,err:=l.leftoverChars.PopFront(); err==nil {
                return v,true
            }
            return byte(0),false
        }).Next(func(
            index int, iterChar byte, status iter.IteratorFeedback,
        ) (iter.IteratorFeedback, byte, error) {
            if status==iter.Break {
                var err error
                if l.leftoverChars.Length()>0 || len(l.curToken)>0 {
                    err=SyntaxError(fmt.Sprintf(
                        "Line: %d Char: %d | Enf of stream reached without completing last token.",
                        l.line,l.char,
                    ))
                }
                return iter.Break,byte(0),err
            }
            l.curToken=append(l.curToken, iterChar)
            num,err:=l.runMatchers(iterChar)
            return l.interpretResult(iterChar,num,err)
        }),
        func(index int, val byte) (Token, error) {
            return Token{
                Data: string(l.curToken),
                Line: l.line,
                Char: l.char,
                Type: l.matchingToken,
            }, nil
        },
    )
}

func (l *Lexer)runMatchers(iterChar byte) (int,error) {
    return iter.MapVals[TokenType,TokenMatcher](
        l.tokens,
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

func (l *Lexer)interpretResult(
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
    _,maxMatchCount:=l.getMaxMatch()
    if maxMatchCount==1 {   // Single regexp matched, return token
        l.setTokenAndLeftovers()
        l.resetTokenState()
        return iter.Continue, iterChar, nil
    } else {    // Multiple or no regexp matched, return error
        return iter.Break, iterChar, SyntaxError(fmt.Sprintf(
            "Line: %d Char %d | No regular expression matched the supplied sequence.",
            l.line, l.char,
        ))
    }
}

func (l *Lexer)getMaxMatch() (int,int) {
    maxMatch:=0
    maxMatchCount:=0
    for k,v:=range(l.tokens) {
        if v.totalChars>maxMatch {
            maxMatch=v.totalChars
            maxMatchCount=1
            l.matchingToken=k
        } else if v.totalChars==maxMatch {
            maxMatchCount+=1
        }
    }
    return maxMatch,maxMatchCount
}

func (l *Lexer)setTokenAndLeftovers() {
    curTok:=[]byte(l.tokens[l.matchingToken].re.FindString(string(l.curToken)))
    l.leftoverChars.SetCapacity(
        l.leftoverChars.Length()+len(l.curToken)-len(curTok),
    )
    for i:=len(l.curToken)-1; i>=len(curTok); i-- {
    // for i:=len(curTok); i<len(p.curToken); i++ {
        l.leftoverChars.PushFront(l.curToken[i])
    }
    l.curToken=curTok
    l.line=l.tokens[l.matchingToken].endLine
    l.char=l.tokens[l.matchingToken].endChar
}

func (l *Lexer)resetTokenState() {
    for _,t:=range(l.tokens) {
        t.endLine=l.line
        t.endChar=l.char
        t.totalChars=0
        t.match=Match
    }
}
