package widgets

import "github.com/barbell-math/util/hash"

type (
	BasePntr[T any, I BaseInterface[T]] struct {
		iFace I
	}
	PartialOrderPntr[T any, I PartialOrderInterface[T]] struct {
		iFace I
	}
	ArithPntr[T any, I ArithInterface[T]] struct {
		iFace I
	}
	PartialOrderArithPntr[T any, I PartialOrderArithInterface[T]] struct {
		iFace I
	}
)

func (p BasePntr[T, I]) Eq(l **T, r **T) bool {
	return p.iFace.Eq(*l, *r)
}
func (p PartialOrderPntr[T, I]) Eq(l **T, r **T) bool {
	return p.iFace.Eq(*l, *r)
}
func (p ArithPntr[T, I]) Eq(l **T, r **T) bool {
	return p.iFace.Eq(*l, *r)
}
func (p PartialOrderArithPntr[T, I]) Eq(l **T, r **T) bool {
	return p.iFace.Eq(*l, *r)
}

func (p PartialOrderPntr[T, I]) Lt(l **T, r **T) bool {
	return p.iFace.Lt(*l, *r)
}
func (p PartialOrderArithPntr[T, I]) Lt(l **T, r **T) bool {
	return p.iFace.Lt(*l, *r)
}

func (p BasePntr[T, I]) Hash(other **T) hash.Hash {
	return p.iFace.Hash(*other)
}
func (p PartialOrderPntr[T, I]) Hash(other **T) hash.Hash {
	return p.iFace.Hash(*other)
}
func (p ArithPntr[T, I]) Hash(other **T) hash.Hash {
	return p.iFace.Hash(*other)
}
func (p PartialOrderArithPntr[T, I]) Hash(other **T) hash.Hash {
	return p.iFace.Hash(*other)
}

func (p BasePntr[T, I]) Zero(other **T) {
	p.iFace.Zero(*other)
}
func (p PartialOrderPntr[T, I]) Zero(other **T) {
	p.iFace.Zero(*other)
}
func (p ArithPntr[T, I]) Zero(other **T) {
	p.iFace.Zero(*other)
}
func (p PartialOrderArithPntr[T, I]) Zero(other **T) {
	p.iFace.Zero(*other)
}

func (p ArithPntr[T, I]) ZeroVal() *T {
	v := p.iFace.ZeroVal()
	return &v
}
func (p PartialOrderArithPntr[T, I]) ZeroVal() *T {
	v := p.iFace.ZeroVal()
	return &v
}

func (p ArithPntr[T, I]) UnitVal() *T {
	v := p.iFace.UnitVal()
	return &v
}
func (p PartialOrderArithPntr[T, I]) UnitVal() *T {
	v := p.iFace.UnitVal()
	return &v
}

func (p ArithPntr[T, I]) Neg(v **T) {
	p.iFace.Neg(*v)
}
func (p PartialOrderArithPntr[T, I]) Neg(v **T) {
	p.iFace.Neg(*v)
}

func (p ArithPntr[T, I]) Add(res **T, l **T, r **T) {
	p.iFace.Add(*res, *l, *r)
}
func (p PartialOrderArithPntr[T, I]) Add(res **T, l **T, r **T) {
	p.iFace.Add(*res, *l, *r)
}

func (p ArithPntr[T, I]) Sub(res **T, l **T, r **T) {
	p.iFace.Sub(*res, *l, *r)
}
func (p PartialOrderArithPntr[T, I]) Sub(res **T, l **T, r **T) {
	p.iFace.Sub(*res, *l, *r)
}

func (p ArithPntr[T, I]) Mul(res **T, l **T, r **T) {
	p.iFace.Mul(*res, *l, *r)
}
func (p PartialOrderArithPntr[T, I]) Mul(res **T, l **T, r **T) {
	p.iFace.Mul(*res, *l, *r)
}

func (p ArithPntr[T, I]) Div(res **T, l **T, r **T) {
	p.iFace.Div(*res, *l, *r)
}
func (p PartialOrderArithPntr[T, I]) Div(res **T, l **T, r **T) {
	p.iFace.Div(*res, *l, *r)
}
