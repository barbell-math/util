package containers

import (
	"github.com/barbell-math/util/hash"
	"github.com/barbell-math/util/widgets"
)

type (
	// badBuiltinInt is a type that forces hash collisions to occur with integer
	// types. It does this by providing a *horrible* hash function.
	// !!! This type should be used for testing only. !!!
	badBuiltinInt struct {
		widgets.BuiltinInt
	}

	// badBuiltinInt2 is a type that forces hash collisions to occur with
	// integer types. It does this by providing a *horrible* hash function.
	// !!! This type should be used for testing only. !!!
	badBuiltinInt2 struct {
		widgets.BuiltinInt
	}
)

func (b badBuiltinInt) Hash(v *int) hash.Hash {
	return hash.Hash(*v % 2)
}
func (b badBuiltinInt2) Hash(v *int) hash.Hash {
	return hash.Hash(*v % 4)
}
