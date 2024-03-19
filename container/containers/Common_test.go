package containers

import (
	"github.com/barbell-math/util/algo/hash"
	"github.com/barbell-math/util/algo/widgets"
)

type (
	// badBuiltinInt is a type that forces hash collisions to occur with integer
	// types. It does this by providing a *horrible* hash function.
	// !!! This type should be used for testing only. !!!
	badBuiltinInt struct {
		widgets.BuiltinInt
	}
)

func (b badBuiltinInt) Hash(v *int) hash.Hash {
	return hash.Hash(*v % 2)
}
