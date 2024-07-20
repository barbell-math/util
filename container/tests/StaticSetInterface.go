package tests

import (
	"testing"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/staticContainers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func staticSetReadInterface[U any](c staticContainers.ReadSet[U])   {}
func staticSetWriteInterface[U any](c staticContainers.WriteSet[U]) {}
func staticSetInterface[U any](c staticContainers.Set[U])           {}

// Tests that the value supplied by the factory implements the
// [containerTypes.StaticCapacity] interface.
func StaticSetInterfaceStaticCapacity[V any](
	factory func(capacity int) staticContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.StaticCapacity = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Length] interface.
func StaticSetInterfaceLengthInterface[V any](
	factory func(capacity int) staticContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.Length = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Clear] interface.
func StaticSetInterfaceClearInterface[V any](
	factory func(capacity int) staticContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.WriteUniqueOps] interface.
func StaticSetInterfaceWriteUniqueOpsInterface[V any](
	factory func(capacity int) staticContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.WriteUniqueOps[uint64, V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.ReadOps] interface.
func StaticSetInterfaceReadOpsInterface[V any](
	factory func(capacity int) staticContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.ReadOps[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.DeleteOps] interface.
func StaticSetInterfaceDeleteOpsInterface[V any](
	factory func(capacity int) staticContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.DeleteOps[uint64, V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [staticContainers.VectorRead] interface.
func ReadStaticSetInterface[V any](
	factory func(capacity int) staticContainers.Set[V],
	t *testing.T,
) {
	staticSetReadInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the
// [staticContainers.WriteVector] interface.
func WriteStaticSetInterface[V any](
	factory func(capacity int) staticContainers.Set[V],
	t *testing.T,
) {
	staticSetWriteInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the
// [staticContainers.Vector] interface.
func StaticSetInterfaceInterface[V any](
	factory func(capacity int) staticContainers.Set[V],
	t *testing.T,
) {
	staticSetInterface[V](factory(0))
}

func staticSetValsHelper(
	factory func(capacity int) staticContainers.Set[int],
	l int,
	c int,
	t *testing.T,
) {
	container := factory(c)
	for i := 0; i < l; i++ {
		test.Nil(container.AppendUnique(i), t)
	}
	test.Eq(c, container.Capacity(), t)
	test.Eq(max(0, l), container.Length(), t)
	cnt := 0
	container.Vals().ForEach(func(index, val int) (iter.IteratorFeedback, error) {
		cnt++
		test.True(container.Contains(val), t)
		return iter.Continue, nil
	})
	test.Eq(max(0, l), cnt, t)
}

// Tests the Vals method functionality of a static set.
func StaticSetInterfaceVals(
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	staticSetValsHelper(factory, -1, 0, t)
	staticSetValsHelper(factory, 0, 0, t)
	staticSetValsHelper(factory, 0, 10, t)
	staticSetValsHelper(factory, 1, 1, t)
	staticSetValsHelper(factory, 1, 10, t)
	staticSetValsHelper(factory, 2, 2, t)
	staticSetValsHelper(factory, 2, 10, t)
	staticSetValsHelper(factory, 5, 5, t)
	staticSetValsHelper(factory, 5, 10, t)
}

func staticSetPntrValsHelper(
	factory func(capacity int) staticContainers.Set[int],
	l int,
	c int,
	t *testing.T,
) {
	container := factory(c)
	for i := 0; i < l; i++ {
		test.Nil(container.AppendUnique(i), t)
	}
	test.Eq(max(0, l), container.Length(), t)
	test.Eq(c, container.Capacity(), t)
	cnt := 0
	container.ValPntrs().ForEach(func(index int, val *int) (iter.IteratorFeedback, error) {
		cnt++
		test.Eq(index, *val, t)
		*val = 100
		return iter.Continue, nil
	})
	container.Vals().ForEach(func(index int, val int) (iter.IteratorFeedback, error) {
		test.Eq(100, val, t)
		return iter.Continue, nil
	})
	test.Eq(max(l, 0), cnt, t)
}

// Tests the ValPntrs method functionality of a static set.
func StaticSetInterfaceValPntrs(
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	container := factory(0)
	if container.IsAddressable() {
		staticSetPntrValsHelper(factory, -1, 0, t)
		staticSetPntrValsHelper(factory, 0, 0, t)
		staticSetPntrValsHelper(factory, 0, 10, t)
		staticSetPntrValsHelper(factory, 1, 1, t)
		staticSetPntrValsHelper(factory, 1, 10, t)
		staticSetPntrValsHelper(factory, 2, 2, t)
		staticSetPntrValsHelper(factory, 2, 10, t)
		staticSetPntrValsHelper(factory, 5, 5, t)
		staticSetPntrValsHelper(factory, 5, 10, t)
	} else {
		test.Panics(func() { container.ValPntrs() }, t)
	}
}

func staticSetContainsPntrHelper(
	factory func(capacity int) staticContainers.Set[int],
	l int,
	c int,
	t *testing.T,
) {
	container := factory(c)
	for i := 0; i < l; i++ {
		test.Nil(container.AppendUnique(i), t)
	}
	for i := 0; i < l; i++ {
		test.True(container.ContainsPntr(&i), t)
	}
	v := l
	test.False(container.ContainsPntr(&v), t)
	v = -1
	test.False(container.ContainsPntr(&v), t)
	container.Pop(0)
	v = 0
	test.False(container.ContainsPntr(&v), t)
	for i := 1; i < l; i++ {
		test.True(container.ContainsPntr(&i), t)
	}
}

// Tests the ContainsPntr method functionality of a static set.
func StaticSetInterfaceContainsPntr(
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	staticSetContainsPntrHelper(factory, 5, 5, t)
	staticSetContainsPntrHelper(factory, 5, 10, t)
}

func staticSetContainsHelper(
	factory func(capacity int) staticContainers.Set[int],
	l int,
	c int,
	t *testing.T,
) {
	container := factory(c)
	for i := 0; i < l; i++ {
		test.Nil(container.AppendUnique(i), t)
	}
	for i := 0; i < l; i++ {
		test.True(container.Contains(i), t)
	}
	test.False(container.Contains(l), t)
	test.False(container.Contains(-1), t)
	container.Pop(0)
	test.False(container.Contains(0), t)
	for i := 1; i < l; i++ {
		test.True(container.Contains(i), t)
	}
}

// Tests the Contains method functionality of a static set.
func StaticSetInterfaceContains(
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	staticSetContainsHelper(factory, 5, 5, t)
	staticSetContainsHelper(factory, 5, 10, t)
}

// Tests the Clear method functionality of a static set.
func StaticSetInterfaceClear(
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	container := factory(5)
	test.Eq(0, container.Length(), t)
	for i := 0; i < 5; i++ {
		test.Nil(container.AppendUnique(i), t)
	}
	test.Eq(5, container.Length(), t)
	container.Clear()
	test.Eq(0, container.Length(), t)
	test.Eq(5, container.Capacity(), t)
}

func staticSetAppendUniqueHelper(
	factory func(capacity int) staticContainers.Set[int],
	l int,
	c int,
	t *testing.T,
) {
	container := factory(c)
	test.Eq(0, container.Length(), t)
	test.Eq(c, container.Capacity(), t)
	for i := 0; i < l; i++ {
		err := container.AppendUnique(i)
		test.Nil(err, t)
	}
	for i := 0; i < l; i++ {
		test.True(container.Contains(i), t)
	}
	for i := 0; i < l; i++ {
		err := container.AppendUnique(i)
		test.Nil(err, t)
		test.Eq(5, container.Length(), t)
	}
	test.Eq(c, container.Capacity(), t)
	if l == c {
		test.ContainsError(containerTypes.Full, container.AppendUnique(-1), t)
		test.ContainsError(containerTypes.Full, container.AppendUnique(5), t)
	} else {
		test.Nil(container.AppendUnique(5), t)
	}
}

// Tests the AppendUnique method functionality of a static set.
func StaticSetInterfaceAppendUnique(
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	staticSetAppendUniqueHelper(factory, 5, 5, t)
	staticSetAppendUniqueHelper(factory, 5, 10, t)

	container := factory(6)
	test.Eq(0, container.Length(), t)
	for i := 0; i < 6; i += 2 {
		container.AppendUnique(i, i+1)
	}
	for i := 0; i < 6; i++ {
		test.True(container.Contains(i), t)
	}
	for i := 0; i < 6; i++ {
		err := container.AppendUnique(i)
		test.Nil(err, t)
		test.Eq(6, container.Length(), t)
	}
	test.ContainsError(containerTypes.Full, container.AppendUnique(-1), t)
	test.ContainsError(containerTypes.Full, container.AppendUnique(6), t)
}

// Tests the UpdateUnique method functionality of a static set.
func StaticSetInterfaceUpdateUnique(
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	container := factory(5)
	test.Eq(0, container.Length(), t)
	for i := 0; i < 5; i++ {
		err := container.AppendUnique(i)
		test.Nil(err, t)
	}
	for i := 0; i < 5; i++ {
		test.True(container.Contains(i), t)
	}
	for i:=0; i<5; i++ {
		err:=container.UpdateUnique(i, func(orig *int) {})
		test.Nil(err,t)
		test.True(container.Contains(i),t)
		test.Eq(5,container.Length(),t)
	}
	for i:=0; i<5; i++ {
		err:=container.UpdateUnique(i, func(orig *int) {*orig=*orig+1})
		test.ContainsError(containerTypes.UpdateViolation, err, t)
		test.ContainsError(customerr.InvalidValue, err, t)
	}
	err:=container.UpdateUnique(6, func(orig *int) {})
	test.ContainsError(containerTypes.ValueError, err, t)
}

// Tests the Pop method functionality of a static set.
func StaticSetInterfacePop(
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	container := factory(5)
	test.Eq(0, container.Length(), t)
	for i := 0; i < 5; i++ {
		test.Nil(container.AppendUnique(i), t)
	}
	for i := 0; i < 5; i++ {
		test.True(container.Contains(i), t)
		container.Pop(i)
		test.False(container.Contains(i), t)
		test.Eq(4-i, container.Length(), t)
		test.Eq(5, container.Capacity(), t)
	}
}

// Tests the UnorderedEq method functionality of a static set.
func StaticSetInterfaceUnorderedEq(
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	innerFactory := func(_cap int, vals ...int) staticContainers.Set[int] {
		rv := factory(_cap)
		rv.AppendUnique(vals...)
		return rv
	}

	v := innerFactory(3, 1, 2, 3)
	v2 := innerFactory(3, 1, 2, 3)
	test.True(v.UnorderedEq(v2), t)
	test.True(v2.UnorderedEq(v), t)

	v = innerFactory(6, 1, 2, 3)
	v2 = innerFactory(3, 1, 2, 3)
	test.True(v.UnorderedEq(v2), t)
	test.True(v2.UnorderedEq(v), t)

	v = innerFactory(3, 1, 2, 3)
	v2 = innerFactory(3, 1, 2)
	test.False(v.UnorderedEq(v2), t)
	test.False(v2.UnorderedEq(v), t)

	v = innerFactory(3, 1, 2, 3)
	v2 = innerFactory(6, 1, 2)
	test.False(v.UnorderedEq(v2), t)
	test.False(v2.UnorderedEq(v), t)

	v = innerFactory(3, 1, 2, 3)
	v2 = innerFactory(3, 3, 1, 2)
	test.True(v.UnorderedEq(v2), t)
	test.True(v2.UnorderedEq(v), t)

	v = innerFactory(3, 1, 2)
	v2 = innerFactory(3, 3, 1, 2)
	test.False(v.UnorderedEq(v2), t)
	test.False(v2.UnorderedEq(v), t)

	v = innerFactory(3, 1, 2, 3)
	v2 = innerFactory(3, 2, 3, 1)
	test.True(v.UnorderedEq(v2), t)
	test.True(v2.UnorderedEq(v), t)

	v = innerFactory(3, 1, 2)
	v2 = innerFactory(3, 2, 3, 1)
	test.False(v.UnorderedEq(v2), t)
	test.False(v2.UnorderedEq(v), t)

	v = innerFactory(3, 1)
	v2 = innerFactory(3, 1)
	test.True(v.UnorderedEq(v2), t)
	test.True(v2.UnorderedEq(v), t)

	v = innerFactory(1, 1)
	v2 = innerFactory(1, 1)
	test.True(v.UnorderedEq(v2), t)
	test.True(v2.UnorderedEq(v), t)

	v = innerFactory(1)
	v2 = innerFactory(1, 1)
	test.False(v.UnorderedEq(v2), t)
	test.False(v2.UnorderedEq(v), t)

	v = innerFactory(1)
	v2 = innerFactory(1)
	test.True(v.UnorderedEq(v2), t)
	test.True(v2.UnorderedEq(v), t)
}

func staticSetIntersectionHelper(
	res staticContainers.Set[int],
	l staticContainers.Set[int],
	r staticContainers.Set[int],
	exp []int,
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	tester := func(expCap int, c staticContainers.Set[int]) {
		test.Eq(len(exp), c.Length(), t)
		test.Eq(expCap, c.Capacity(), t)
		for _, v := range exp {
			test.True(c.Contains(v), t)
		}
	}
	expCap := l.Length() + r.Length()
	res.Intersection(l, r)
	tester(expCap, res)
	expCap = l.Length() + r.Length()
	res.Intersection(r, l)
	tester(expCap, res)
}

// Tests the Intersection method functionality of a static set.
func StaticSetInterfaceIntersection(
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	v := factory(3)
	v2 := factory(3)
	staticSetIntersectionHelper(factory(0), v, v2, []int{}, factory, t)
	v.AppendUnique(1)
	staticSetIntersectionHelper(factory(0), v, v2, []int{}, factory, t)
	v2.AppendUnique(1)
	staticSetIntersectionHelper(factory(1), v, v2, []int{1}, factory, t)
	v2.AppendUnique(2)
	staticSetIntersectionHelper(factory(1), v, v2, []int{1}, factory, t)
	v.AppendUnique(2)
	staticSetIntersectionHelper(factory(2), v, v2, []int{1, 2}, factory, t)
	v.AppendUnique(3)
	staticSetIntersectionHelper(factory(2), v, v2, []int{1, 2}, factory, t)
	v2.AppendUnique(3)
	staticSetIntersectionHelper(factory(3), v, v2, []int{1, 2, 3}, factory, t)

	if !v.IsSynced() {
		v = factory(4)
		v2 = factory(2)
		v.AppendUnique(1, 2, 3, 4)
		v2.AppendUnique(2, 4)
		staticSetIntersectionHelper(v, v, v2, []int{2, 4}, factory, t)
	}
}

func staticSetUnionHelper(
	res staticContainers.Set[int],
	l staticContainers.Set[int],
	r staticContainers.Set[int],
	exp []int,
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	tester := func(expCap int, c staticContainers.Set[int]) {
		test.Eq(len(exp), c.Length(), t)
		test.Eq(expCap, c.Capacity(), t)
		for _, v := range exp {
			test.True(c.Contains(v), t)
		}
	}
	expCap := l.Length() + r.Length()
	res.Union(l, r)
	tester(expCap, res)
	expCap = l.Length() + r.Length()
	res.Union(r, l)
	tester(expCap, res)
}

// Tests the Union method functionality of a static set.
func StaticSetInterfaceUnion(
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	v := factory(3)
	v2 := factory(3)
	staticSetUnionHelper(factory(0), v, v2, []int{}, factory, t)
	v.AppendUnique(1)
	staticSetUnionHelper(factory(0), v, v2, []int{1}, factory, t)
	v2.AppendUnique(1)
	staticSetUnionHelper(factory(0), v, v2, []int{1}, factory, t)
	v2.AppendUnique(2)
	staticSetUnionHelper(factory(0), v, v2, []int{1, 2}, factory, t)
	v.AppendUnique(2)
	staticSetUnionHelper(factory(0), v, v2, []int{1, 2}, factory, t)
	v.AppendUnique(3)
	staticSetUnionHelper(factory(0), v, v2, []int{1, 2, 3}, factory, t)
	v2.AppendUnique(3)
	staticSetUnionHelper(factory(0), v, v2, []int{1, 2, 3}, factory, t)

	if !v.IsSynced() {
		v = factory(4)
		v2 = factory(4)
		v.AppendUnique(1, 2, 3, 4)
		v2.AppendUnique(2, 4, 5, 6)
		staticSetUnionHelper(v, v, v2, []int{1, 2, 3, 4, 5, 6}, factory, t)
	}
}

func staticSetDifferenceHelper(
	res staticContainers.Set[int],
	l staticContainers.Set[int],
	r staticContainers.Set[int],
	exp []int,
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	expCap := l.Length()
	res.Difference(l, r)
	test.Eq(expCap, res.Capacity(), t)
	test.Eq(len(exp), res.Length(), t)
	for _, v := range exp {
		test.True(res.Contains(v), t)
	}
}

// Tests the Difference method functionality of a static set.
func StaticSetInterfaceDifference(
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	innerFactory := func(_cap int, vals ...int) staticContainers.Set[int] {
		rv := factory(_cap)
		rv.AppendUnique(vals...)
		return rv
	}
	v := innerFactory(6)
	v2 := innerFactory(3)
	staticSetDifferenceHelper(factory(0), v, v2, []int{}, factory, t)
	staticSetDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(3)
	v2 = innerFactory(3)
	staticSetDifferenceHelper(factory(0), v, v2, []int{}, factory, t)
	staticSetDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(6, 1)
	v2 = innerFactory(3)
	staticSetDifferenceHelper(factory(0), v, v2, []int{1}, factory, t)
	staticSetDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(6, 1)
	v2 = innerFactory(3, 1)
	staticSetDifferenceHelper(factory(0), v, v2, []int{}, factory, t)
	staticSetDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(6, 1)
	v2 = innerFactory(3, 1, 2)
	staticSetDifferenceHelper(factory(0), v, v2, []int{}, factory, t)
	staticSetDifferenceHelper(factory(0), v2, v, []int{2}, factory, t)

	v = innerFactory(6, 1, 2)
	v2 = innerFactory(3, 1, 2)
	staticSetDifferenceHelper(factory(0), v, v2, []int{}, factory, t)
	staticSetDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(6, 1, 2, 3)
	v2 = innerFactory(3, 1, 2)
	staticSetDifferenceHelper(factory(0), v, v2, []int{3}, factory, t)
	staticSetDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(6, 1, 2, 3)
	v2 = innerFactory(3, 1, 2, 3)
	staticSetDifferenceHelper(factory(0), v, v2, []int{}, factory, t)
	staticSetDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(6, 1, 2, 3, 4, 5, 6)
	v2 = innerFactory(3, 1, 2, 3)
	staticSetDifferenceHelper(factory(0), v, v2, []int{4, 5, 6}, factory, t)
	staticSetDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	if !v.IsSynced() {
		v = innerFactory(4, 1, 2, 3, 4)
		v2 = innerFactory(4, 2, 4)
		staticSetDifferenceHelper(v, v, v2, []int{1, 3}, factory, t)
	}
}

func StaticSetInterfaceIsSuperset(
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	innerFactory := func(_cap int, vals ...int) staticContainers.Set[int] {
		rv := factory(_cap)
		rv.AppendUnique(vals...)
		return rv
	}
	v := innerFactory(5)
	v2 := innerFactory(3)
	test.True(v.IsSuperset(v2), t)

	v = innerFactory(3)
	v2 = innerFactory(3)
	test.True(v.IsSuperset(v2), t)

	v = innerFactory(5, 1)
	v2 = innerFactory(3)
	test.True(v.IsSuperset(v2), t)
	test.False(v2.IsSuperset(v), t)

	v = innerFactory(5, 1)
	v2 = innerFactory(3, 1)
	test.True(v.IsSuperset(v2), t)
	test.True(v2.IsSuperset(v), t)

	v = innerFactory(5, 1)
	v2 = innerFactory(3, 1, 2)
	test.False(v.IsSuperset(v2), t)
	test.True(v2.IsSuperset(v), t)

	v = innerFactory(5, 1)
	v2 = innerFactory(3, 1, 2, 3)
	test.False(v.IsSuperset(v2), t)
	test.True(v2.IsSuperset(v), t)

	v = innerFactory(5, 1, 2, 3)
	v2 = innerFactory(3, 1, 2, 3)
	test.True(v.IsSuperset(v2), t)
	test.True(v2.IsSuperset(v), t)

	v = innerFactory(5, 1, 2, 3, 4, 5)
	v2 = innerFactory(3, 1, 2, 3)
	test.True(v.IsSuperset(v2), t)
	test.False(v2.IsSuperset(v), t)
}

func StaticSetInterfaceIsSubset(
	factory func(capacity int) staticContainers.Set[int],
	t *testing.T,
) {
	innerFactory := func(_cap int, vals ...int) staticContainers.Set[int] {
		rv := factory(_cap)
		rv.AppendUnique(vals...)
		return rv
	}
	v := innerFactory(5)
	v2 := innerFactory(3)
	test.True(v.IsSubset(v2), t)

	v = innerFactory(3)
	v2 = innerFactory(3)
	test.True(v.IsSubset(v2), t)

	v = innerFactory(5, 1)
	v2 = innerFactory(3)
	test.False(v.IsSubset(v2), t)
	test.True(v2.IsSubset(v), t)

	v = innerFactory(5, 1)
	v2 = innerFactory(3, 1)
	test.True(v.IsSubset(v2), t)
	test.True(v2.IsSubset(v), t)

	v = innerFactory(5, 1)
	v2 = innerFactory(3, 1, 2)
	test.True(v.IsSubset(v2), t)
	test.False(v2.IsSubset(v), t)

	v = innerFactory(5, 1)
	v2 = innerFactory(3, 1, 2, 3)
	test.True(v.IsSubset(v2), t)
	test.False(v2.IsSubset(v), t)

	v = innerFactory(5, 1, 2, 3)
	v2 = innerFactory(3, 1, 2, 3)
	test.True(v.IsSubset(v2), t)
	test.True(v2.IsSubset(v), t)

	v = innerFactory(5, 1, 2, 3, 4, 5)
	v2 = innerFactory(3, 1, 2, 3)
	test.False(v.IsSubset(v2), t)
	test.True(v2.IsSubset(v), t)
}
