package translators

import (
	"github.com/barbell-math/util/math/basic"
)

type (
	FlagCntr[T basic.Int | basic.Uint] struct {
		val T
	}
)

func (f *FlagCntr[T]) Translate(arg string) (T, error) {
	f.val++
	return f.val, nil
}
