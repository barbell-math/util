package tests

import (
	"testing"

	"github.com/barbell-math/util/src/container/basic"
	"github.com/barbell-math/util/src/container/containerTypes"
	"github.com/barbell-math/util/src/container/dynamicContainers"
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/iter"
	"github.com/barbell-math/util/src/test"
)

func dynamicVectorReadInterface[U any](c dynamicContainers.ReadVector[U])   {}
func dynamicVectorWriteInterface[U any](c dynamicContainers.WriteVector[U]) {}
func dynamicVectorInterface[U any](c dynamicContainers.Vector[U])           {}

// Tests that the value supplied by the factory implements the
// [containerTypes.RWSyncable] interface.
func DynVectorInterfaceSyncableInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.RWSyncable = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Addressable] interface.
func DynVectorInterfaceAddressableInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.Addressable = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Length] interface.
func DynVectorInterfaceLengthInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.Length = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Capacity] interface.
func DynVectorInterfaceCapacityInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.Capacity = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Clear] interface.
func DynVectorInterfaceClearInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.WriteOps] interface.
func DynVectorInterfaceWriteOpsInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.WriteOps[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.WriteKeyedOps] interface.
func DynVectorInterfaceWriteKeyedOpsInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.WriteKeyedOps[int, V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.WriteKeyedSequentialOps] interface.
func DynVectorInterfaceWriteKeyedSequentialOpsInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.WriteKeyedSequentialOps[int, V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedWriteOps] interface.
func DynVectorInterfaceWriteDynKeyedOpsInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.WriteDynKeyedOps[int, V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.ReadOps] interface.
func DynVectorInterfaceReadOpsInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.ReadOps[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedReadOps] interface.
func DynVectorInterfaceReadKeyedOpsInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.ReadKeyedOps[int, V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.DeleteOps] interface.
func DynVectorInterfaceDeleteOpsInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.DeleteOps[int, V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedDeleteOps] interface.
func DynVectorInterfaceDeleteKeyedOpsInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.DeleteKeyedOps[int, V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedDeleteOps] interface.
func DynVectorInterfaceDeleteSequentialOpsInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.DeleteSequentialOps[int, V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedDeleteOps] interface.
func DynVectorInterfaceDeleteKeyedSequentialOpsInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.DeleteKeyedSequentialOps[int, V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.VectorRead] interface.
func ReadDynVectorInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	dynamicVectorReadInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.WriteVector] interface.
func WriteDynVectorInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	dynamicVectorWriteInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.Vector] interface.
func DynVectorInterfaceInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
	t *testing.T,
) {
	dynamicVectorInterface[V](factory(0))
}

// Tests that the value supplied by the factory does not implement the
// [staticContainers.Vector] interface.
func DynVectorInterfaceStaticCapacityInterface[V any](
	factory func(capacity int) dynamicContainers.Vector[V],
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

// Tests the Get method functionality of a dynamic vector.
func DynVectorInterfaceGet(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory(0)
	_, err := container.Get(0)
	test.ContainsError(customerr.ValOutsideRange, err, t)
	for i := 0; i < 5; i++ {
		container.Append(i)
	}
	for i := 0; i < 5; i++ {
		_v, err := container.Get(i)
		test.Nil(err, t)
		test.Eq(i, _v, t)
	}
	_, err = container.Get(-1)
	test.ContainsError(customerr.ValOutsideRange, err, t)
	_, err = container.Get(6)
	test.ContainsError(customerr.ValOutsideRange, err, t)
}

// Tests the GetPntr method functionality of a dynamic vector.
func DynVectorInterfaceGetPntr(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory(0)
	if container.IsAddressable() {
		v, err := container.GetPntr(0)
		test.ContainsError(customerr.ValOutsideRange, err, t)
		test.NilPntr[int](v, t)
		for i := 0; i < 5; i++ {
			container.Append(i)
		}
		for i := 0; i < 5; i++ {
			_v, err := container.GetPntr(i)
			test.Eq(i, *_v, t)
			test.Eq(nil, err, t)
		}
		v, err = container.GetPntr(-1)
		test.ContainsError(customerr.ValOutsideRange, err, t)
		test.NilPntr[int](v, t)
		v, err = container.GetPntr(6)
		test.ContainsError(customerr.ValOutsideRange, err, t)
		test.NilPntr[int](v, t)
	} else {
		test.Panics(
			func() {
				container := factory(0)
				container.GetPntr(1)
			},
			t,
		)
	}
}

func dynVectorContainsHelper(
	v dynamicContainers.Vector[int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Append(i)
	}
	for i := 0; i < l; i++ {
		test.True(v.Contains(i), t)
	}
	test.False(v.Contains(-1), t)
	test.False(v.Contains(l), t)
}

// Tests the Contains method functionality of a dynamic vector.
func DynVectorInterfaceContains(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	dynVectorContainsHelper(factory(0), 0, t)
	dynVectorContainsHelper(factory(0), 1, t)
	dynVectorContainsHelper(factory(0), 2, t)
	dynVectorContainsHelper(factory(0), 5, t)
}

func dynVectorContainsPntrHelper(
	v dynamicContainers.Vector[int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Append(i)
	}
	for i := 0; i < l; i++ {
		test.True(v.ContainsPntr(&i), t)
	}
	tmp := -1
	test.False(v.ContainsPntr(&tmp), t)
	test.False(v.ContainsPntr(&l), t)
}

// Tests the ContainsPntr method functionality of a dynamic vector.
func DynVectorInterfaceContainsPntr(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	dynVectorContainsPntrHelper(factory(0), 0, t)
	dynVectorContainsPntrHelper(factory(0), 1, t)
	dynVectorContainsPntrHelper(factory(0), 2, t)
	dynVectorContainsPntrHelper(factory(0), 5, t)
}

func dynVectorKeyOfHelper(
	v dynamicContainers.Vector[int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Append(i)
	}
	for i := 0; i < l; i++ {
		k, found := v.KeyOf(i)
		test.Eq(i, k, t)
		test.True(found, t)
	}
	_, found := v.KeyOf(-1)
	test.False(found, t)
	_, found = v.KeyOf(-1)
	test.False(v.Contains(l), t)
}

// Tests the KeyOf method functionality of a dynamic vector.
func DynVectorInterfaceKeyOf(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	dynVectorKeyOfHelper(factory(0), 0, t)
	dynVectorKeyOfHelper(factory(0), 1, t)
	dynVectorKeyOfHelper(factory(0), 2, t)
	dynVectorKeyOfHelper(factory(0), 5, t)
}

func dynVectorKeyOfPntrHelper(
	v dynamicContainers.Vector[int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Append(i)
	}
	for i := 0; i < l; i++ {
		k, found := v.KeyOfPntr(&i)
		test.Eq(i, k, t)
		test.True(found, t)
	}
	tmp := -1
	_, found := v.KeyOfPntr(&tmp)
	test.False(found, t)
	_, found = v.KeyOfPntr(&tmp)
	test.False(v.Contains(l), t)
}

// Tests the KeyOfPntr method functionality of a dynamic vector.
func DynVectorInterfaceKeyOfPntr(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	dynVectorKeyOfPntrHelper(factory(0), 0, t)
	dynVectorKeyOfPntrHelper(factory(0), 1, t)
	dynVectorKeyOfPntrHelper(factory(0), 2, t)
	dynVectorKeyOfPntrHelper(factory(0), 5, t)
}

// Tests the Set method functionality of a dynamic vector.
func DynVectorInterfaceSet(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory(0)
	err := container.Set(basic.Pair[int, int]{0, 6})
	test.ContainsError(customerr.ValOutsideRange, err, t)
	for i := 0; i < 5; i++ {
		container.Append(i)
	}
	for i := 0; i < 5; i++ {
		err := container.Set(basic.Pair[int, int]{i, i + 1})
		test.Nil(err, t)
	}
	for i := 0; i < 5; i++ {
		iterV, err := container.Get(i)
		test.Nil(err, t)
		test.Eq(i+1, iterV, t)
	}
	err = container.Set(basic.Pair[int, int]{-1, 6})
	test.ContainsError(customerr.ValOutsideRange, err, t)
	err = container.Set(basic.Pair[int, int]{6, 6})
	test.ContainsError(customerr.ValOutsideRange, err, t)
}

func dynVectorSetSequentialHelper(
	factory func(capacity int) dynamicContainers.Vector[int],
	idx int,
	vals []int,
	l int,
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < l; i++ {
		container.Append(i)
	}
	err := container.SetSequential(idx, vals...)
	test.Nil(err, t)
	for i := 0; i < container.Length(); i++ {
		iterV, _ := container.Get(i)
		if i >= idx && i < idx+len(vals) {
			test.Eq(vals[i-idx], iterV, t)
		} else {
			test.Eq(i, iterV, t)
		}
	}
}

// Tests the SetSequential method functionality of a dynamic vector.
func DynVectorInterfaceSetSequential(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	for i := 0; i < 20; i++ {
		for j := 0; j < i; j++ {
			vals := []int{}
			for k := 0; k < i-j; k++ {
				vals = append(vals, k)
				dynVectorSetSequentialHelper(factory, j, vals, i, t)
			}
		}
	}
}

// Tests the Append method functionality of a dynamic vector.
func DynVectorInterfaceAppend(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 5; i++ {
		container.Append(i)
	}
	for i := 0; i < 5; i++ {
		iterV, _ := container.Get(i)
		test.Eq(i, iterV, t)
	}
	container.Append(5, 6, 7)
	for i := 0; i < 8; i++ {
		iterV, _ := container.Get(i)
		test.Eq(i, iterV, t)
	}
}

func dynVectorInsertHelper(
	factory func(capacity int) dynamicContainers.Vector[int],
	vals []basic.Pair[int, int],
	l int,
	t *testing.T,
) {
	indexContained := func(idx int) (basic.Pair[int, int], bool) {
		for _, v := range vals {
			if idx == v.A {
				return v, true
			}
		}
		return basic.Pair[int, int]{}, false
	}
	container := factory(0)
	for i := 0; i < l; i++ {
		err := container.Append(i)
		test.Nil(err, t)
	}
	err := container.Insert(vals...)
	test.Nil(err, t)
	test.Eq(l+len(vals), container.Length(), t)
	offset := 0
	for i := 0; i < container.Length(); i++ {
		iterV, err := container.Get(i)
		test.Nil(err, t)
		if v, ok := indexContained(i); ok {
			test.Eq(v.B, iterV, t)
			offset++
		} else {
			test.Eq(i-offset, iterV, t)
		}
	}
}

// Tests the Insert method functionality of a dynamic vector.
func DynVectorInterfaceInsert(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory(0)
	err := container.Insert(basic.Pair[int, int]{1, 0})
	test.ContainsError(containerTypes.KeyError, err, t)
	for i := 0; i < 20; i++ {
		vals := []basic.Pair[int, int]{}
		dynVectorInsertHelper(factory, vals, i, t)
		for j := 0; j < 20; j++ {
			vals = append(vals, basic.Pair[int, int]{j, j})
			dynVectorInsertHelper(factory, vals, i, t)
		}
	}
	for i := 0; i < 20; i++ {
		vals := []basic.Pair[int, int]{}
		dynVectorInsertHelper(factory, vals, i, t)
		for j := 0; j < i; j += 2 {
			vals = append(vals, basic.Pair[int, int]{j, j})
			dynVectorInsertHelper(factory, vals, i, t)
		}
	}
}

func dynVectorInsertSequentialHelper(
	factory func(capacity int) dynamicContainers.Vector[int],
	idx int,
	l int,
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < l-1; i++ {
		container.Append(i)
	}
	err := container.InsertSequential(idx, l-1)
	test.Nil(err, t)
	test.Eq(l, container.Length(), t)
	for i := 0; i < container.Length(); i++ {
		var exp int
		v, _ := container.Get(i)
		if i < idx {
			exp = i
		} else if i == idx {
			exp = l - 1
		} else {
			exp = i - 1
		}
		test.Eq(exp, v, t)
	}
	err = container.InsertSequential(l+1, -1)
	test.ContainsError(customerr.ValOutsideRange, err, t)
	test.Eq(l, container.Length(), t)
	err = container.InsertSequential(-1, -1)
	test.ContainsError(customerr.ValOutsideRange, err, t)
	test.Eq(l, container.Length(), t)
}

// Tests the InsertSequential method functionality of a dynamic vector.
func DynVectorInterfaceInsertSequential(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 2; i >= 0; i-- {
		err := container.InsertSequential(0, i)
		test.Nil(err, t)
	}
	for i := 3; i < 5; i++ {
		err := container.InsertSequential(container.Length(), i)
		test.Nil(err, t)
	}
	for i := 0; i < 5; i++ {
		iterV, err := container.Get(i)
		test.Nil(err, t)
		test.Eq(i, iterV, t)
	}
	container = factory(0)
	container.InsertSequential(0, 0, 1, 2)
	container.InsertSequential(3, 4, 5)
	container.InsertSequential(3, 3)
	for i := 0; i < 6; i++ {
		iterV, err := container.Get(i)
		test.Nil(err, t)
		test.Eq(i, iterV, t)
	}
	for i := 0; i < 5; i++ {
		dynVectorInsertSequentialHelper(factory, i, 5, t)
	}
}

func dynVectorPopSequentialHelper(
	factory func(capacity int) dynamicContainers.Vector[int],
	l int,
	num int,
	t *testing.T,
) {
	// fmt.Println("Permutation: l: ",l," num: ",num)
	container := factory(0)
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			container.Append(-1)
		} else {
			container.Append(i)
		}
	}
	// fmt.Println("Init:   ",v)
	n := container.PopSequential(-1, num)
	exp := factory(0)
	cntr := 0
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			if cntr < num {
				cntr++
				continue
			} else {
				exp.Append(-1)
			}
		} else {
			exp.Append(i)
		}
	}
	test.Eq(exp.Length(), container.Length(), t)
	test.Eq(cntr, n, t)
	// fmt.Println("EXP:    ",exp)
	// fmt.Println("Final:  ",v)
	for i := 0; i < container.Length(); i++ {
		iterV, found := container.Get(i)
		expIterV, foundExp := exp.Get(i)
		test.Nil(found, t)
		test.Nil(foundExp, t)
		test.Eq(expIterV, iterV, t)
	}
}

// Tests the PopSequential method functionality of a dynamic vector.
func DynVectorInterfacePopSequential(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	for i := 0; i < 13; i++ {
		for j := 0; j < 13; j++ {
			dynVectorPopSequentialHelper(factory, i, j, t)
		}
	}
}

func dynVectorPopHelper(
	factory func(capacity int) dynamicContainers.Vector[int],
	l int,
	t *testing.T,
) {
	// fmt.Println("Permutation: l: ",l," num: ",num)
	container := factory(0)
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			container.Append(-1)
		} else {
			container.Append(i)
		}
	}
	// fmt.Println("Init:   ",v)
	n := container.Pop(-1)
	exp := factory(0)
	cntr := 0
	for i := 0; i < l; i++ {
		if i%4 != 0 {
			exp.Append(i)
		} else {
			cntr++
		}
	}
	test.Eq(exp.Length(), container.Length(), t)
	test.Eq(cntr, n, t)
	// fmt.Println("EXP:    ",exp)
	// fmt.Println("Final:  ",v)
	for i := 0; i < container.Length(); i++ {
		iterV, found := container.Get(i)
		expIterV, foundExp := exp.Get(i)
		test.Nil(found, t)
		test.Nil(foundExp, t)
		test.Eq(expIterV, iterV, t)
	}
}

// Tests the Pop method functionality of a dynamic vector.
func DynVectorInterfacePop(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	for i := 0; i < 13; i++ {
		dynVectorPopHelper(factory, i, t)
	}
}

func dynVectorPopPntrHelper(
	factory func(capacity int) dynamicContainers.Vector[int],
	l int,
	t *testing.T,
) {
	// fmt.Println("Permutation: l: ",l," num: ",num)
	container := factory(0)
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			container.Append(-1)
		} else {
			container.Append(i)
		}
	}
	// fmt.Println("Init:   ",v)
	tmp := -1
	n := container.PopPntr(&tmp)
	exp := factory(0)
	cntr := 0
	for i := 0; i < l; i++ {
		if i%4 != 0 {
			exp.Append(i)
		} else {
			cntr++
		}
	}
	test.Eq(exp.Length(), container.Length(), t)
	test.Eq(cntr, n, t)
	// fmt.Println("EXP:    ",exp)
	// fmt.Println("Final:  ",v)
	for i := 0; i < container.Length(); i++ {
		iterV, found := container.Get(i)
		expIterV, foundExp := exp.Get(i)
		test.Nil(found, t)
		test.Nil(foundExp, t)
		test.Eq(expIterV, iterV, t)
	}
}

// Tests the PopPntr method functionality of a dynamic vector.
func DynVectorInterfacePopPntr(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	for i := 0; i < 13; i++ {
		dynVectorPopPntrHelper(factory, i, t)
	}
}

// Tests the Delete method functionality of a dynamic vector.
func DynVectorInterfaceDelete(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory(0)
	err := container.Delete(0)
	test.ContainsError(customerr.ValOutsideRange, err, t)
	for i := 0; i < 6; i++ {
		container.Append(i)
	}
	err = container.Delete(-1)
	test.ContainsError(customerr.ValOutsideRange, err, t)
	err = container.Delete(7)
	test.ContainsError(customerr.ValOutsideRange, err, t)
	for i := container.Length() - 1; i >= 0; i-- {
		container.Delete(i)
		test.Eq(i, container.Length(), t)
		for j := 0; j < i; j++ {
			iterV, err := container.Get(j)
			test.Nil(err, t)
			test.Eq(j, iterV, t)
		}
	}
	err = container.Delete(0)
	test.ContainsError(customerr.ValOutsideRange, err, t)
}

func dynVectorDeleteSequentialHelper(
	factory func(capacity int) dynamicContainers.Vector[int],
	start int,
	end int,
	l int,
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < l; i++ {
		container.Append(i)
	}
	container.DeleteSequential(start, end)
	test.Eq(l-(end-start), container.Length(), t)
	for i := 0; i < l; i++ {
		if i < start {
			v, err := container.Get(i)
			test.Nil(err, nil)
			test.Eq(i, v, t)
		} else if i >= end {
			v, err := container.Get(i - (end - start))
			test.Nil(err, nil)
			test.Eq(i, v, t)
		} else {
			test.False(container.Contains(i), t)
		}
	}
}

// Tests the DeleteSequential method functionality of a dynamic vector.
func DynVectorInterfaceDeleteSequential(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory(0)
	container.Append(0, 1, 2, 3)
	err := container.DeleteSequential(-1, 3)
	test.ContainsError(customerr.ValOutsideRange, err, t)
	test.Eq(4, container.Length(), t)
	err = container.DeleteSequential(0, 5)
	test.ContainsError(customerr.ValOutsideRange, err, t)
	test.Eq(4, container.Length(), t)
	err = container.DeleteSequential(2, 1)
	test.ContainsError(customerr.InvalidValue, err, t)
	test.Eq(4, container.Length(), t)
	for i := 0; i < 20; i++ {
		for j := 0; j < i; j++ {
			for k := j; k < i; k++ {
				dynVectorDeleteSequentialHelper(factory, j, k, i, t)
			}
		}
	}
}

// Tests the Clear method functionality of a dynamic vector.
func DynVectorInterfaceClear(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 6; i++ {
		container.Append(i)
	}
	container.Clear()
	test.Eq(0, container.Length(), t)
	test.Eq(0, container.Capacity(), t)
}

func dynVectorValsHelper(
	factory func(capacity int) dynamicContainers.Vector[int],
	l int,
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < l; i++ {
		container.Append(i)
	}
	cnt := 0
	container.Vals().ForEach(func(index, val int) (iter.IteratorFeedback, error) {
		cnt++
		test.Eq(index, val, t)
		return iter.Continue, nil
	})
	test.Eq(max(0, l), cnt, t)
}

// Tests the Vals method functionality of a dynamic vector.
func DynVectorInterfaceVals(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	dynVectorValsHelper(factory, -1, t)
	dynVectorValsHelper(factory, 0, t)
	dynVectorValsHelper(factory, 1, t)
	dynVectorValsHelper(factory, 2, t)
	dynVectorValsHelper(factory, 5, t)
}

func dynVectorPntrValsHelper(
	factory func(capacity int) dynamicContainers.Vector[int],
	l int,
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < l; i++ {
		container.Append(i)
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
	test.Eq(max(0, l), cnt, t)
}

// Tests the ValPntrs method functionality of a dynamic vector.
func DynVectorInterfaceValPntrs(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory(0)
	if container.IsAddressable() {
		dynVectorPntrValsHelper(factory, -1, t)
		dynVectorPntrValsHelper(factory, 0, t)
		dynVectorPntrValsHelper(factory, 1, t)
		dynVectorPntrValsHelper(factory, 2, t)
		dynVectorPntrValsHelper(factory, 5, t)
	} else {
		test.Panics(func() { container.ValPntrs() }, t)
	}
}

func dynVectorKeysHelper(
	factory func(capacity int) dynamicContainers.Vector[int],
	l int,
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < l; i++ {
		container.Append(i)
	}
	cnt := 0
	container.Keys().ForEach(func(index, val int) (iter.IteratorFeedback, error) {
		cnt++
		test.Eq(index, val, t)
		return iter.Continue, nil
	})
	test.Eq(max(l, 0), cnt, t)
}

// Tests the Keys method functionality of a dynamic vector.
func DynVectorInterfaceKeys(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	dynVectorKeysHelper(factory, -1, t)
	dynVectorKeysHelper(factory, 0, t)
	dynVectorKeysHelper(factory, 1, t)
	dynVectorKeysHelper(factory, 2, t)
	dynVectorKeysHelper(factory, 5, t)
}

// Tests the KeyedEq method functionality of a dynamic vector.
func DynVectorInterfaceKeyedEq(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	innerFactory := func(vals ...int) dynamicContainers.Vector[int] {
		rv := factory(0)
		rv.Append(vals...)
		return rv
	}
	v := innerFactory(1, 2, 3)
	v2 := innerFactory(1, 2, 3)
	test.True(v.KeyedEq(v2), t)
	test.True(v2.KeyedEq(v), t)

	v = innerFactory(1, 2)
	v2 = innerFactory(1, 2, 3)
	test.False(v.KeyedEq(v2), t)
	test.False(v2.KeyedEq(v), t)

	v = innerFactory(1, 2, 3)
	v2 = innerFactory(3, 1, 2)
	test.False(v.KeyedEq(v2), t)
	test.False(v2.KeyedEq(v), t)

	v = innerFactory(1, 2)
	v2 = innerFactory(3, 1, 2)
	test.False(v.KeyedEq(v2), t)
	test.False(v2.KeyedEq(v), t)

	v = innerFactory(0)
	v2 = innerFactory(0)
	test.True(v.KeyedEq(v2), t)
	test.True(v2.KeyedEq(v), t)

	v = innerFactory()
	v2 = innerFactory(0)
	test.False(v.KeyedEq(v2), t)
	test.False(v2.KeyedEq(v), t)

	v = innerFactory()
	v2 = innerFactory()
	test.True(v.KeyedEq(v2), t)
	test.True(v2.KeyedEq(v), t)
}

// Tests the UnorderedEq method functionality of a dynamic vector.
func DynVectorInterfaceUnorderedEq(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	v := factory(0)
	v.Append(1, 2, 3)
	v2 := factory(0)
	v2.Append(1, 2, 3)
	test.True(v.UnorderedEq(v2), t)
	test.True(v2.UnorderedEq(v), t)
	v.Pop(3)
	test.False(v.UnorderedEq(v2), t)
	test.False(v2.UnorderedEq(v), t)
	v.Append(3)
	v2 = factory(0)
	v2.Append(3, 1, 2)
	test.True(v.UnorderedEq(v2), t)
	test.True(v2.UnorderedEq(v), t)
	v.Pop(3)
	test.False(v.UnorderedEq(v2), t)
	test.False(v2.UnorderedEq(v), t)
	v.Append(3)
	v2 = factory(0)
	v2.Append(2, 3, 1)
	test.True(v.UnorderedEq(v2), t)
	test.True(v2.UnorderedEq(v), t)
	v.Pop(3)
	test.False(v.UnorderedEq(v2), t)
	test.False(v2.UnorderedEq(v), t)
	v = factory(0)
	v.Append(0)
	v2 = factory(0)
	v2.Append(0)
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

func dynVectorIntersectionHelper(
	res dynamicContainers.Vector[int],
	l dynamicContainers.Vector[int],
	r dynamicContainers.Vector[int],
	exp []int,
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	tester := func(c dynamicContainers.Vector[int]) {
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

// Tests the Intersection method functionality of a dynamic vector.
func DynVectorInterfaceIntersection(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	v := factory(0)
	v2 := factory(0)
	dynVectorIntersectionHelper(factory(0), v, v2, []int{}, factory, t)
	v.Append(1)
	dynVectorIntersectionHelper(factory(0), v, v2, []int{}, factory, t)
	v2.Append(1)
	dynVectorIntersectionHelper(factory(0), v, v2, []int{1}, factory, t)
	v2.Append(2)
	dynVectorIntersectionHelper(factory(0), v, v2, []int{1}, factory, t)
	v.Append(2)
	dynVectorIntersectionHelper(factory(0), v, v2, []int{1, 2}, factory, t)
	v.Append(3)
	dynVectorIntersectionHelper(factory(0), v, v2, []int{1, 2}, factory, t)
	v2.Append(3)
	dynVectorIntersectionHelper(factory(0), v, v2, []int{1, 2, 3}, factory, t)

	if !v.IsSynced() {
		v = factory(0)
		v2 = factory(0)
		v.Append(1, 2, 3, 4)
		v2.Append(2, 4)
		dynVectorIntersectionHelper(v, v, v2, []int{2, 4}, factory, t)
	}
}

func dynVectorUnionHelper(
	res dynamicContainers.Vector[int],
	l dynamicContainers.Vector[int],
	r dynamicContainers.Vector[int],
	exp []int,
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	tester := func(c dynamicContainers.Vector[int]) {
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

// Tests the Union method functionality of a dynamic vector.
func DynVectorInterfaceUnion(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	v := factory(0)
	v2 := factory(0)
	dynVectorUnionHelper(factory(0), v, v2, []int{}, factory, t)
	v.Append(1)
	dynVectorUnionHelper(factory(0), v, v2, []int{1}, factory, t)
	v2.Append(1)
	dynVectorUnionHelper(factory(0), v, v2, []int{1}, factory, t)
	v2.Append(2)
	dynVectorUnionHelper(factory(0), v, v2, []int{1, 2}, factory, t)
	v.Append(2)
	dynVectorUnionHelper(factory(0), v, v2, []int{1, 2}, factory, t)
	v.Append(3)
	dynVectorUnionHelper(factory(0), v, v2, []int{1, 2, 3}, factory, t)
	v2.Append(3)
	dynVectorUnionHelper(factory(0), v, v2, []int{1, 2, 3}, factory, t)

	if !v.IsSynced() {
		v = factory(0)
		v2 = factory(0)
		v.Append(1, 2, 3, 4)
		v2.Append(2, 4, 5, 6)
		dynVectorUnionHelper(v, v, v2, []int{1, 2, 3, 4, 5, 6}, factory, t)
	}
}

func dynVectorDifferenceHelper(
	res dynamicContainers.Vector[int],
	l dynamicContainers.Vector[int],
	r dynamicContainers.Vector[int],
	exp []int,
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	res.Difference(l, r)
	test.Eq(len(exp), res.Length(), t)
	for _, v := range exp {
		test.True(res.Contains(v), t)
	}
}

// Tests the Difference method functionality of a dynamic vector.
func DynVectorInterfaceDifference(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	innerFactory := func(vals ...int) dynamicContainers.Vector[int] {
		rv := factory(0)
		rv.Append(vals...)
		return rv
	}
	v := innerFactory()
	v2 := innerFactory()
	dynVectorDifferenceHelper(factory(0), v, v2, []int{}, factory, t)
	dynVectorDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(1)
	v2 = innerFactory()
	dynVectorDifferenceHelper(factory(0), v, v2, []int{1}, factory, t)
	dynVectorDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(1)
	v2 = innerFactory(1)
	dynVectorDifferenceHelper(factory(0), v, v2, []int{}, factory, t)
	dynVectorDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(1)
	v2 = innerFactory(1, 2)
	dynVectorDifferenceHelper(factory(0), v, v2, []int{}, factory, t)
	dynVectorDifferenceHelper(factory(0), v2, v, []int{2}, factory, t)

	v = innerFactory(1, 2)
	v2 = innerFactory(1, 2)
	dynVectorDifferenceHelper(factory(0), v, v2, []int{}, factory, t)
	dynVectorDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(1, 2, 3)
	v2 = innerFactory(1, 2)
	dynVectorDifferenceHelper(factory(0), v, v2, []int{3}, factory, t)
	dynVectorDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(1, 2, 3)
	v2 = innerFactory(1, 2, 3)
	dynVectorDifferenceHelper(factory(0), v, v2, []int{}, factory, t)
	dynVectorDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	v = innerFactory(1, 2, 3, 4, 5, 6)
	v2 = innerFactory(1, 2, 3)
	dynVectorDifferenceHelper(factory(0), v, v2, []int{4, 5, 6}, factory, t)
	dynVectorDifferenceHelper(factory(0), v2, v, []int{}, factory, t)

	if !v.IsSynced() {
		v = innerFactory(1, 2, 3, 4)
		v2 = innerFactory(2, 4)
		dynVectorDifferenceHelper(v, v, v2, []int{1, 3}, factory, t)
	}
}

func DynVectorInterfaceIsSuperset(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	innerFactory := func(vals ...int) dynamicContainers.Vector[int] {
		rv := factory(0)
		rv.Append(vals...)
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

func DynVectorInterfaceIsSubset(
	factory func(capacity int) dynamicContainers.Vector[int],
	t *testing.T,
) {
	innerFactory := func(vals ...int) dynamicContainers.Vector[int] {
		rv := factory(0)
		rv.Append(vals...)
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
