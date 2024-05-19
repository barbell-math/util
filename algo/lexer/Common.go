package lexer

import "github.com/barbell-math/util/container/basic"

const (
	LambdaChar byte=0
)

type dfaNode struct {
	flags int
	transitions []basic.Pair[byte,int]
}

type DFA map[int]dfaNode

type Token struct {
	Line int
	Char int
	Id int
}

type Lexer map[Regex]Token

