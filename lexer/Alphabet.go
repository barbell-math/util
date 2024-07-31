package lexer

import (
	"github.com/barbell-math/util/container/containers"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/widgets"
)

type (
	// A type that defines what alphabet the lexer will be working with
	Alphabet[A any] dynamicContainers.ReadSet[A]
)

var (
	ASCII_ALPHABET dynamicContainers.ReadSet[byte] = makeAsciiAlphabet()
)

func makeAsciiAlphabet() dynamicContainers.ReadSet[byte] {
	rv, _ := containers.NewHashSet[byte, widgets.BuiltinByte](0)
	var i byte = 20
	for ; i < 127; i++ {
		rv.AppendUnique(i)
	}
	rv.AppendUnique('\n', '\r')
	return &rv
}
