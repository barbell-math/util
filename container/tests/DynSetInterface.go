package tests

import (
	"testing"

	"github.com/barbell-math/util/iter"
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func dynamicSetReadInterface[U any](c dynamicContainers.ReadSet[U])   {}
func dynamicSetWriteInterface[U any](c dynamicContainers.WriteSet[U]) {}
func dynamicSetInterface[U any](c dynamicContainers.Set[U])           {}

// Tests that the value supplied by the factory implements the
// [containerTypes.RWSyncable] interface.
func DynSetInterfaceSyncableInterface[V any](
	factory func(capacity int) dynamicContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.RWSyncable = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.RWSyncable] interface.
func DynSetInterfaceAddressableInterface[V any](
	factory func(capacity int) dynamicContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.Addressable = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Length] interface.
func DynSetInterfaceLengthInterface[V any](
	factory func(capacity int) dynamicContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.Length = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Clear] interface.
func DynSetInterfaceClearInterface[V any](
	factory func(capacity int) dynamicContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.WriteUniqueOps] interface.
func DynSetInterfaceWriteUniqueOpsInterface[V any](
	factory func(capacity int) dynamicContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.WriteUniqueOps[uint64, V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.ReadOps] interface.
func DynSetInterfaceReadOpsInterface[V any](
	factory func(capacity int) dynamicContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.ReadOps[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.DeleteOps] interface.
func DynSetInterfaceDeleteOpsInterface[V any](
	factory func(capacity int) dynamicContainers.Set[V],
	t *testing.T,
) {
	var container containerTypes.DeleteOps[uint64, V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.VectorRead] interface.
func ReadDynSetInterface[V any](
	factory func(capacity int) dynamicContainers.Set[V],
	t *testing.T,
) {
	dynamicSetReadInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.WriteVector] interface.
func WriteDynSetInterface[V any](
	factory func(capacity int) dynamicContainers.Set[V],
	t *testing.T,
) {
	dynamicSetWriteInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.Vector] interface.
func DynSetInterfaceInterface[V any](
	factory func(capacity int) dynamicContainers.Set[V],
	t *testing.T,
) {
	dynamicSetInterface[V](factory(0))
}

// Tests that the value supplied by the factory does not implement the
// [containerTypes.StaticCapacity] interface.
func DynSetInterfaceStaticCapacityInterface[V any](
	factory func(capacity int) dynamicContainers.Set[V],
	t *testing.T,
) {
	test.Panics(
		func() {
			var c any
			c = factory(0)
			c2 := c.(containerTypes.StaticCapacity)
			_ = c2
		},
		t,
	)
}

func dynSetValsHelper(
	factory func(capacity int) dynamicContainers.Set[int],
	l int,
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < l; i++ {
		container.AppendUnique(i)
	}
	cnt := 0
	container.Vals().ForEach(func(index, val int) (iter.IteratorFeedback, error) {
		cnt++
		test.True(container.Contains(val), t)
		return iter.Continue, nil
	})
	test.Eq(l, cnt, t)
}

// Tests the Vals method functionality of a dynamic set.
func DynSetInterfaceVals(
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	dynSetValsHelper(factory, 0, t)
	dynSetValsHelper(factory, 1, t)
	dynSetValsHelper(factory, 2, t)
	dynSetValsHelper(factory, 5, t)
}

func dynSetPntrValsHelper(
	factory func(capacity int) dynamicContainers.Set[int],
	l int,
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < l; i++ {
		container.AppendUnique(i)
	}
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
	test.Eq(l, cnt, t)
}

// Tests the ValPntrs method functionality of a dynamic set.
func DynSetInterfaceValPntrs(
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	container := factory(0)
	if container.IsAddressable() {
		dynSetPntrValsHelper(factory, 0, t)
		dynSetPntrValsHelper(factory, 1, t)
		dynSetPntrValsHelper(factory, 2, t)
	} else {
		test.Panics(func() { container.ValPntrs() }, t)
	}
}

// Tests the ContainsPntr method functionality of a dynamic set.
func DynSetInterfaceContainsPntr(
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 5; i++ {
		container.AppendUnique(i)
	}
	for i := 0; i < 5; i++ {
		test.True(container.ContainsPntr(&i), t)
	}
	v := 5
	test.False(container.ContainsPntr(&v), t)
	v = -1
	test.False(container.ContainsPntr(&v), t)
	container.Pop(0)
	v = 0
	test.False(container.ContainsPntr(&v), t)
	for i := 1; i < 5; i++ {
		test.True(container.ContainsPntr(&i), t)
	}
}

// Tests the GetUnique method functionality of a dynamic set.
func DynSetInterfaceGetUnique(
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 5; i++ {
		container.AppendUnique(i)
	}
	tmp := 0
	for i := 0; i < 5; i++ {
		tmp = i
		test.True(container.Contains(i), t)
		test.Nil(container.GetUnique(&tmp), t)
		test.Eq(i, tmp, t)
	}
	tmp = -1
	test.ContainsError(containerTypes.ValueError, container.GetUnique(&tmp), t)
}

// Tests the Contains method functionality of a dynamic set.
func DynSetInterfaceContains(
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 5; i++ {
		container.AppendUnique(i)
	}
	for i := 0; i < 5; i++ {
		test.True(container.Contains(i), t)
	}
	test.False(container.Contains(5), t)
	test.False(container.Contains(-1), t)
	container.Pop(0)
	test.False(container.Contains(0), t)
	for i := 1; i < 5; i++ {
		test.True(container.Contains(i), t)
	}
}

// Tests the Clear method functionality of a dynamic set.
func DynSetInterfaceClear(
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	container := factory(0)
	test.Eq(0, container.Length(), t)
	for i := 0; i < 5; i++ {
		container.AppendUnique(i)
	}
	test.Eq(5, container.Length(), t)
	container.Clear()
	test.Eq(0, container.Length(), t)
}

// Tests the AppendUnique method functionality of a dynamic set.
func DynSetInterfaceAppendUnique(
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	container := factory(0)
	test.Eq(0, container.Length(), t)
	for i := 0; i < 5; i++ {
		err := container.AppendUnique(i)
		test.Nil(err, t)
	}
	for i := 0; i < 5; i++ {
		test.True(container.Contains(i), t)
	}
	for i := 0; i < 5; i++ {
		err := container.AppendUnique(i)
		test.Nil(err, t)
		test.Eq(5, container.Length(), t)
	}
	container = factory(0)
	test.Eq(0, container.Length(), t)
	for i := 0; i < 6; i += 2 {
		container.AppendUnique(i, i+1)
	}
	for i := 0; i < 6; i++ {
		test.True(container.Contains(i), t)
	}
}

// Tests the UpdateUnique method functionality of a dynamic set.
func DynSetInterfaceUpdateUnique(
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	type testStruct struct {
		id    int
		other int
	}
	container := factory(0)
	test.Eq(0, container.Length(), t)
	for i := 0; i < 5; i++ {
		err := container.AppendUnique(i)
		test.Nil(err, t)
	}
	for i := 0; i < 5; i++ {
		test.True(container.Contains(i), t)
	}
	for i := 0; i < 5; i++ {
		err := container.UpdateUnique(i, func(orig *int) {})
		test.Nil(err, t)
		test.True(container.Contains(i), t)
		test.Eq(5, container.Length(), t)
	}
	for i := 0; i < 5; i++ {
		err := container.UpdateUnique(i, func(orig *int) { *orig = *orig + 1 })
		test.ContainsError(containerTypes.UpdateViolation, err, t)
		test.ContainsError(customerr.InvalidValue, err, t)
	}
	err := container.UpdateUnique(6, func(orig *int) {})
	test.ContainsError(containerTypes.ValueError, err, t)
}

// Tests the Pop method functionality of a dynamic set.
func DynSetInterfacePop(
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	container := factory(0)
	test.Eq(0, container.Length(), t)
	for i := 0; i < 5; i++ {
		container.AppendUnique(i)
	}
	for i := 0; i < 5; i++ {
		test.True(container.Contains(i), t)
		container.Pop(i)
		test.False(container.Contains(i), t)
		test.Eq(4-i, container.Length(), t)
	}
}

// Tests the PopPntr method functionality of a dynamic set.
func DynSetInterfacePopPntr(
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	container := factory(0)
	test.Eq(0, container.Length(), t)
	for i := 0; i < 5; i++ {
		container.AppendUnique(i)
	}
	for i := 0; i < 5; i++ {
		test.True(container.Contains(i), t)
		container.PopPntr(&i)
		test.False(container.Contains(i), t)
		test.Eq(4-i, container.Length(), t)
	}
}

// Tests the UnorderedEq method functionality of a dynamic set.
func DynSetInterfaceUnorderedEq(
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	v := factory(0)
	v.AppendUnique(1, 2, 3)
	v2 := factory(0)
	v2.AppendUnique(1, 2, 3)
	test.True(v.UnorderedEq(v2), t)
	test.True(v2.UnorderedEq(v), t)
	v.Pop(3)
	test.False(v.UnorderedEq(v2), t)
	test.False(v2.UnorderedEq(v), t)
	v.AppendUnique(3)
	v2 = factory(0)
	v2.AppendUnique(3, 1, 2)
	test.True(v.UnorderedEq(v2), t)
	test.True(v2.UnorderedEq(v), t)
	v.Pop(3)
	test.False(v.UnorderedEq(v2), t)
	test.False(v2.UnorderedEq(v), t)
	v.AppendUnique(3)
	v2 = factory(0)
	v2.AppendUnique(2, 3, 1)
	test.True(v.UnorderedEq(v2), t)
	test.True(v2.UnorderedEq(v), t)
	v.Pop(3)
	test.False(v.UnorderedEq(v2), t)
	test.False(v2.UnorderedEq(v), t)
	v = factory(0)
	v.AppendUnique(0)
	v2 = factory(0)
	v2.AppendUnique(0)
	test.True(v.UnorderedEq(v2), t)
	test.True(v2.UnorderedEq(v), t)
	v.Pop(0)
	test.False(v.UnorderedEq(v2), t)
	test.False(v2.UnorderedEq(v), t)
	v = factory(0)
	v2 = factory(0)
	test.True(v.UnorderedEq(v2), t)
	test.True(v2.UnorderedEq(v), t)
}

func dynSetIntersectionHelper(
	res dynamicContainers.Set[int],
	l dynamicContainers.Set[int],
	r dynamicContainers.Set[int],
	exp []int,
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	tester := func(c dynamicContainers.Set[int]) {
		test.Eq(len(exp), c.Length(), t)
		for _, v := range exp {
			test.True(c.Contains(v), t)
		}
	}
	res.Intersection(l, r)
	tester(res)
	res.Intersection(r, l)
	tester(res)
}

// Tests the Intersection method functionality of a dynamic set.
func DynSetInterfaceIntersection(
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	v := factory(0)
	v2 := factory(0)
	dynSetIntersectionHelper(factory(0), v, v2, []int{}, factory, t)
	v.AppendUnique(1)
	dynSetIntersectionHelper(factory(0), v, v2, []int{}, factory, t)
	v2.AppendUnique(1)
	dynSetIntersectionHelper(factory(0), v, v2, []int{1}, factory, t)
	v2.AppendUnique(2)
	dynSetIntersectionHelper(factory(0), v, v2, []int{1}, factory, t)
	v.AppendUnique(2)
	dynSetIntersectionHelper(factory(0), v, v2, []int{1, 2}, factory, t)
	v.AppendUnique(3)
	dynSetIntersectionHelper(factory(0), v, v2, []int{1, 2}, factory, t)
	v2.AppendUnique(3)
	dynSetIntersectionHelper(factory(0), v, v2, []int{1, 2, 3}, factory, t)

	if !v.IsSynced() {
		v = factory(0)
		v2 = factory(0)
		v.AppendUnique(1, 2, 3, 4)
		v2.AppendUnique(2, 4)
		dynSetIntersectionHelper(v, v, v2, []int{2, 4}, factory, t)
	}
}

func dynSetUnionHelper(
	res dynamicContainers.Set[int],
	l dynamicContainers.Set[int],
	r dynamicContainers.Set[int],
	exp []int,
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	tester := func(c dynamicContainers.Set[int]) {
		test.Eq(len(exp), c.Length(), t)
		for _, v := range exp {
			test.True(c.Contains(v), t)
		}
	}
	res.Union(l, r)
	tester(res)
	res.Union(r, l)
	tester(res)
}

// Tests the Union method functionality of a dynamic set.
func DynSetInterfaceUnion(
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	v := factory(0)
	v2 := factory(0)
	dynSetUnionHelper(factory(0), v, v2, []int{}, factory, t)
	v.AppendUnique(1)
	dynSetUnionHelper(factory(0), v, v2, []int{1}, factory, t)
	v2.AppendUnique(1)
	dynSetUnionHelper(factory(0), v, v2, []int{1}, factory, t)
	v2.AppendUnique(2)
	dynSetUnionHelper(factory(0), v, v2, []int{1, 2}, factory, t)
	v.AppendUnique(2)
	dynSetUnionHelper(factory(0), v, v2, []int{1, 2}, factory, t)
	v.AppendUnique(3)
	dynSetUnionHelper(factory(0), v, v2, []int{1, 2, 3}, factory, t)
	v2.AppendUnique(3)
	dynSetUnionHelper(factory(0), v, v2, []int{1, 2, 3}, factory, t)

	if !v.IsSynced() {
		v = factory(0)
		v2 = factory(0)
		v.AppendUnique(1, 2, 3, 4)
		v2.AppendUnique(2, 4, 5, 6)
		dynSetUnionHelper(v, v, v2, []int{1, 2, 3, 4, 5, 6}, factory, t)
	}
}

func dynSetDifferenceHelper(
	res dynamicContainers.Set[int],
	l dynamicContainers.Set[int],
	r dynamicContainers.Set[int],
	exp []int,
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	res.Difference(l, r)
	test.Eq(len(exp), res.Length(), t)
	for _, v := range exp {
		test.True(res.Contains(v), t)
	}
}

// Tests the Difference method functionality of a dynamic set.
func DynSetInterfaceDifference(
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	innerFactory := func(vals ...int) dynamicContainers.Set[int] {
		rv := factory(0)
		rv.AppendUnique(vals...)
		return rv
	}
	v := innerFactory()
	v2 := innerFactory()
	dynSetDifferenceHelper(factory(0), v, v2, []int{}, factory, t)
	dynSetDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(1)
	v2 = innerFactory()
	dynSetDifferenceHelper(factory(0), v, v2, []int{1}, factory, t)
	dynSetDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(1)
	v2 = innerFactory(1)
	dynSetDifferenceHelper(factory(0), v, v2, []int{}, factory, t)
	dynSetDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(1)
	v2 = innerFactory(1, 2)
	dynSetDifferenceHelper(factory(0), v, v2, []int{}, factory, t)
	dynSetDifferenceHelper(factory(0), v2, v, []int{2}, factory, t)

	v = innerFactory(1, 2)
	v2 = innerFactory(1, 2)
	dynSetDifferenceHelper(factory(0), v, v2, []int{}, factory, t)
	dynSetDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(1, 2, 3)
	v2 = innerFactory(1, 2)
	dynSetDifferenceHelper(factory(0), v, v2, []int{3}, factory, t)
	dynSetDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(1, 2, 3)
	v2 = innerFactory(1, 2, 3)
	dynSetDifferenceHelper(factory(0), v, v2, []int{}, factory, t)
	dynSetDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(1, 2, 3, 4, 5, 6)
	v2 = innerFactory(1, 2, 3)
	dynSetDifferenceHelper(factory(0), v, v2, []int{4, 5, 6}, factory, t)
	dynSetDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	if !v.IsSynced() {
		v = innerFactory(1, 2, 3, 4)
		v2 = innerFactory(2, 4)
		dynSetDifferenceHelper(v, v, v2, []int{1, 3}, factory, t)
	}
}

func DynSetInterfaceIsSuperset(
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	innerFactory := func(vals ...int) dynamicContainers.Set[int] {
		rv := factory(0)
		rv.AppendUnique(vals...)
		return rv
	}

	v := innerFactory()
	v2 := innerFactory()
	test.True(v.IsSuperset(v2), t)

	v = innerFactory(1)
	v2 = innerFactory()
	test.True(v.IsSuperset(v2), t)
	test.False(v2.IsSuperset(v), t)

	v = innerFactory(1)
	v2 = innerFactory(1)
	test.True(v.IsSuperset(v2), t)
	test.True(v2.IsSuperset(v), t)

	v = innerFactory(1)
	v2 = innerFactory(1, 2)
	test.False(v.IsSuperset(v2), t)
	test.True(v2.IsSuperset(v), t)

	v = innerFactory(1)
	v2 = innerFactory(1, 2, 3)
	test.False(v.IsSuperset(v2), t)
	test.True(v2.IsSuperset(v), t)

	v = innerFactory(1, 2, 3)
	v2 = innerFactory(1, 2, 3)
	test.True(v.IsSuperset(v2), t)
	test.True(v2.IsSuperset(v), t)

	v = innerFactory(1, 2, 3, 4, 5)
	v2 = innerFactory(1, 2, 3)
	test.True(v.IsSuperset(v2), t)
	test.False(v2.IsSuperset(v), t)
}

func DynSetInterfaceIsSubset(
	factory func(capacity int) dynamicContainers.Set[int],
	t *testing.T,
) {
	innerFactory := func(vals ...int) dynamicContainers.Set[int] {
		rv := factory(0)
		rv.AppendUnique(vals...)
		return rv
	}

	v := innerFactory()
	v2 := innerFactory()
	test.True(v.IsSubset(v2), t)

	v = innerFactory(1)
	v2 = innerFactory()
	test.False(v.IsSubset(v2), t)
	test.True(v2.IsSubset(v), t)

	v = innerFactory(1)
	v2 = innerFactory(1)
	test.True(v.IsSubset(v2), t)
	test.True(v2.IsSubset(v), t)

	v = innerFactory(1)
	v2 = innerFactory(1, 2)
	test.True(v.IsSubset(v2), t)
	test.False(v2.IsSubset(v), t)

	v = innerFactory(1)
	v2 = innerFactory(1, 2, 3)
	test.True(v.IsSubset(v2), t)
	test.False(v2.IsSubset(v), t)

	v = innerFactory(1, 2, 3)
	v2 = innerFactory(1, 2, 3)
	test.True(v.IsSubset(v2), t)
	test.True(v2.IsSubset(v), t)

	v = innerFactory(1, 2, 3, 4, 5)
	v2 = innerFactory(1, 2, 3)
	test.False(v.IsSubset(v2), t)
	test.True(v2.IsSubset(v), t)
}
