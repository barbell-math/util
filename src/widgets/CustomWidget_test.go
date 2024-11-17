package widgets

import (
	"testing"

	"github.com/barbell-math/util/src/hash"
	"github.com/barbell-math/util/src/test"
)

type customWidget struct {
	a int
	b float32
}

func (_ *customWidget) Eq(l *customWidget, r *customWidget) bool {
	return l.a == r.a
}
func (_ *customWidget) Lt(l *customWidget, r *customWidget) bool {
	return l.a < r.a
}
func (_ *customWidget) Hash(v *customWidget) hash.Hash {
	return hash.Hash(v.a)
}
func (_ *customWidget) Zero(v *customWidget) {
	*v = customWidget{}
}

func TestCustomWidget(t *testing.T) {
	v := customWidget{a: 10, b: 20}
	v2 := customWidget{a: 9, b: 20}
	w := PartialOrder[customWidget, *customWidget]{}
	test.True(w.Eq(&v, &v), t)
	test.False(w.Eq(&v, &v2), t)
	test.False(w.Lt(&v, &v), t)
	test.False(w.Lt(&v2, &v2), t)
	test.False(w.Lt(&v, &v2), t)
	test.True(w.Lt(&v2, &v), t)
	test.Eq(hash.Hash(10), w.Hash(&v), t)
	test.Eq(hash.Hash(9), w.Hash(&v2), t)

	v2 = customWidget{a: 9, b: 20}
	w2 := Base[customWidget, *customWidget]{}
	test.True(w2.Eq(&v, &v), t)
	test.False(w2.Eq(&v, &v2), t)
	test.Eq(hash.Hash(10), w2.Hash(&v), t)
	test.Eq(hash.Hash(9), w2.Hash(&v2), t)
}

func TestCustomWidgetPntr(t *testing.T) {
	vImpl := &customWidget{a: 10, b: 20}
	v2Impl := &customWidget{a: 9, b: 20}
	v := PartialOrder[
		*customWidget,
		PartialOrderPntr[customWidget, *customWidget],
	]{}
	test.True(v.Eq(&vImpl, &vImpl), t)
	test.False(v.Eq(&vImpl, &v2Impl), t)
	test.False(v.Lt(&vImpl, &vImpl), t)
	test.False(v.Lt(&v2Impl, &v2Impl), t)
	test.False(v.Lt(&vImpl, &v2Impl), t)
	test.True(v.Lt(&v2Impl, &vImpl), t)
	test.Eq(hash.Hash(10), v.Hash(&vImpl), t)
	test.Eq(hash.Hash(9), v.Hash(&v2Impl), t)
}
