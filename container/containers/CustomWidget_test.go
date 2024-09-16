package containers

import (
	"testing"

	"github.com/barbell-math/util/hash"
	"github.com/barbell-math/util/test"
	"github.com/barbell-math/util/widgets"
)

type customWidget struct {
	a int
	b float32
}

func (c *customWidget) Eq(l *customWidget, r *customWidget) bool {
	return l.a == r.a
}
func (c *customWidget) Lt(l *customWidget, r *customWidget) bool {
	return l.a < r.a
}
func (c *customWidget) Hash(v *customWidget) hash.Hash {
	return hash.Hash(v.a)
}
func (c *customWidget) Zero(v *customWidget) {
	*v = customWidget{}
}

func TestCustomWidgetInVector(t *testing.T) {
	v := make(Vector[customWidget, *customWidget], 0)
	v.Append(customWidget{a: 10, b: 20}, customWidget{a: 9, b: 20})
	test.True(v[0].Eq(&v[0], &v[0]), t)
	test.False(v[0].Eq(&v[0], &v[1]), t)
	test.False(v[0].Lt(&v[0], &v[0]), t)
	test.False(v[0].Lt(&v[1], &v[1]), t)
	test.False(v[0].Lt(&v[0], &v[1]), t)
	test.True(v[0].Lt(&v[1], &v[0]), t)
	test.Eq(hash.Hash(10), v[0].Hash(&v[0]), t)
	test.Eq(hash.Hash(9), v[0].Hash(&v[1]), t)
}

func TestCustomWidgetPntr(t *testing.T) {
	v := make(Vector[
		*customWidget,
		widgets.BasePntr[customWidget, *customWidget],
	], 0)
	v.Append(&customWidget{a: 10, b: 20}, &customWidget{a: 9, b: 20})
	test.True(v[0].Eq(v[0], v[0]), t)
	test.False(v[0].Eq(v[0], v[1]), t)
	test.False(v[0].Lt(v[0], v[0]), t)
	test.False(v[0].Lt(v[1], v[1]), t)
	test.False(v[0].Lt(v[0], v[1]), t)
	test.True(v[0].Lt(v[1], v[0]), t)
	test.Eq(hash.Hash(10), v[0].Hash(v[0]), t)
	test.Eq(hash.Hash(9), v[0].Hash(v[1]), t)

	v2 := make(Vector[
		*customWidget,
		widgets.BasePntr[customWidget, *customWidget],
	], 0)
	v2.Append(&customWidget{a: 10, b: 20}, &customWidget{a: 9, b: 20})
	test.True(v.Eq(&v, &v2), t)
	v2[0] = &customWidget{a: 100, b: 200}
	test.False(v.Eq(&v, &v2), t)
}
