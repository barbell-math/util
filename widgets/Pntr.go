package widgets

import "github.com/barbell-math/util/hash"

type (
	BasePntr[T any, I BaseInterface[T]] struct {
		iFace I
	}
	PartialOrderPntr[T any, I PartialOrderInterface[T]] struct {
		iFace I
	}
)

func (p BasePntr[T, I]) Eq(l **T, r **T) bool {
	return p.iFace.Eq(*l, *r)
}
func (p PartialOrderPntr[T, I]) Eq(l **T, r **T) bool {
	return p.iFace.Eq(*l, *r)
}

func (p PartialOrderPntr[T, I]) Lt(l **T, r **T) bool {
	return p.iFace.Lt(*l, *r)
}

func (p BasePntr[T, I]) Hash(other **T) hash.Hash {
	return p.iFace.Hash(*other)
}
func (p PartialOrderPntr[T, I]) Hash(other **T) hash.Hash {
	return p.iFace.Hash(*other)
}

func (p BasePntr[T, I]) Zero(other **T) {
	p.iFace.Zero(*other)
}
func (p PartialOrderPntr[T, I]) Zero(other **T) {
	p.iFace.Zero(*other)
}
