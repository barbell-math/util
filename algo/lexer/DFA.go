package lexer

import "github.com/barbell-math/util/container/basic"

type (
	dfaID   int
	dfaFlag int
	dfaNode struct {
		flags       int
		transitions []basic.Pair[byte, int]
	}

	DFA map[int]dfaNode
)

const (
	dfaStart  dfaID   = 0
	dfaSource dfaFlag = 1 << iota
	dfaSink
)

// func (d DFA) FromNFA(n NFA) error {
// 	return nil
// }
