package lexer

import (
	"fmt"
	"testing"

	"github.com/barbell-math/util/test"
)

func getFullRangeOfChars() string {
	var offset byte=0
	b:=make([]byte,127-32-len(specialCharEncoding)-1)
	var i byte=32
	for ; i<127; i++ {
		if _,ok:=specialCharEncoding[i]; !ok && i!='\\' {
			b[i-32-offset]=i
		} else {
			offset++
		}
	}
	return string(b)
}

func TestRegexTokenStreamNominalCase(t *testing.T) {
	s:=getFullRangeOfChars()
	r:=Regex(s)
	tokens,err:=r.toTokenStream().Collect()
	test.Nil(err,t)
	test.Eq(len(s),len(tokens),t)
	for i:=0; i<len(tokens); i++ {
		test.Eq(s[i],tokens[i],t)
	}
}

func TestEscapeCharacters(t *testing.T) {
	for c,_:=range(escapeChars) {
		r:=Regex(fmt.Sprintf("\\%c",c))
		tokens,err:=r.toTokenStream().Collect()
		test.Nil(err,t)
		test.Eq(fmt.Sprintf("%c",c),string(tokens),t)
	}
}

func TestMultipleEscapeChars(t *testing.T){
	r:=Regex("abc\\|abc\\*")
	tokens,err:=r.toTokenStream().Collect()
	test.Nil(err,t)
	test.Eq("abc|abc*",string(tokens),t)

	r=Regex("\\|abc\\*abc")
	tokens,err=r.toTokenStream().Collect()
	test.Nil(err,t)
	test.Eq("|abc*abc",string(tokens),t)

	r=Regex("\\|\\*abcabc")
	tokens,err=r.toTokenStream().Collect()
	test.Nil(err,t)
	test.Eq("|*abcabc",string(tokens),t)

	r=Regex("abcabc\\|\\*")
	tokens,err=r.toTokenStream().Collect()
	test.Nil(err,t)
	test.Eq("abcabc|*",string(tokens),t)
}

func TestInvalidEscapeChar(t *testing.T) {
	r:=Regex("\\a")
	tokens,err:=r.toTokenStream().Collect()
	test.ContainsError(RegexSyntaxError,err,t)
	test.ContainsError(RegexInvalidEscapeSequence,err,t)
	test.Eq("",string(tokens),t)

	r=Regex("abc\\a")
	tokens,err=r.toTokenStream().Collect()
	test.ContainsError(RegexSyntaxError,err,t)
	test.ContainsError(RegexInvalidEscapeSequence,err,t)
	test.Eq("abc",string(tokens),t)

	r=Regex("\\aabc")
	tokens,err=r.toTokenStream().Collect()
	test.ContainsError(RegexSyntaxError,err,t)
	test.ContainsError(RegexInvalidEscapeSequence,err,t)
	test.Eq("",string(tokens),t)
}

func TestSpecialCharEncoding(t *testing.T){
	r:=Regex("_|(a)|b|c*")
	tokens,err:=r.toTokenStream().Collect()
	test.Nil(err,t)
	test.SlicesMatch[byte](
		[]byte{
			lambdaChar,barChar,lParenChar,
			'a',rParenChar,barChar,
			'b',barChar,'c',starChar,
		},
		tokens,
		t,
	)
}
