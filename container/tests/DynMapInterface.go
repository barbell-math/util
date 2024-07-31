package tests

import (
	"testing"

	"github.com/barbell-math/util/iter"
	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/test"
)

func mapReadInterface[T any, U any](c dynamicContainers.ReadMap[T, U])   {}
func mapWriteInterface[T any, U any](c dynamicContainers.WriteMap[T, U]) {}
func mapInterface[T any, U any](c dynamicContainers.Map[T, U])           {}

// Tests that the value supplied by the factory implements the
// [containerTypes.RWSyncable] interface.
func DynMapInterfaceSyncableInterface[K any, V any](
	factory func(capacity int) dynamicContainers.Map[K, V],
	t *testing.T,
) {
	var container containerTypes.RWSyncable = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.RWSyncable] interface.
func DynMapInterfaceAddressableInterface[K any, V any](
	factory func(capacity int) dynamicContainers.Map[K, V],
	t *testing.T,
) {
	var container containerTypes.Addressable = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Length] interface.
func DynMapInterfaceLengthInterface[K any, V any](
	factory func(capacity int) dynamicContainers.Map[K, V],
	t *testing.T,
) {
	var container containerTypes.Length = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Clear] interface.
func DynMapInterfaceClearInterface[K any, V any](
	factory func(capacity int) dynamicContainers.Map[K, V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedWriteOps] interface.
func DynMapInterfaceWriteKeyedOpsInterface[K any, V any](
	factory func(capacity int) dynamicContainers.Map[K, V],
	t *testing.T,
) {
	var container containerTypes.WriteKeyedOps[K, V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.ReadOps] interface.
func DynMapInterfaceReadOpsInterface[K any, V any](
	factory func(capacity int) dynamicContainers.Map[K, V],
	t *testing.T,
) {
	var container containerTypes.ReadOps[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedReadOps] interface.
func DynMapInterfaceReadKeyedOpsInterface[K any, V any](
	factory func(capacity int) dynamicContainers.Map[K, V],
	t *testing.T,
) {
	var container containerTypes.ReadKeyedOps[K, V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.DeleteOps] interface.
func DynMapInterfaceDeleteOpsInterface[K any, V any](
	factory func(capacity int) dynamicContainers.Map[K, V],
	t *testing.T,
) {
	var container containerTypes.DeleteOps[K, V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedDeleteOps] interface.
func DynMapInterfaceDeleteKeyedOpsInterface[K any, V any](
	factory func(capacity int) dynamicContainers.Map[K, V],
	t *testing.T,
) {
	var container containerTypes.DeleteKeyedOps[K, V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.MapRead] interface.
func ReadDynMapInterface[K any, V any](
	factory func(capacity int) dynamicContainers.Map[K, V],
	t *testing.T,
) {
	mapReadInterface[K, V](factory(0))
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.WriteMap] interface.
func WriteDynMapInterface[K any, V any](
	factory func(capacity int) dynamicContainers.Map[K, V],
	t *testing.T,
) {
	mapWriteInterface[K, V](factory(0))
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.Map] interface.
func DynMapInterfaceInterface[K any, V any](
	factory func(capacity int) dynamicContainers.Map[K, V],
	t *testing.T,
) {
	mapInterface[K, V](factory(0))
}

// Tests that the value supplied by the factory does not implement the
// [staticContainers.Map] interface.
func DynMapInterfaceStaticCapacityInterface[K any, V any](
	factory func(capacity int) dynamicContainers.Map[K, V],
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

// Tests the Get method functionality of a dynamic map.
func DynMapInterfaceGet(
	factory func(capacity int) dynamicContainers.Map[int, int],
	t *testing.T,
) {
	container := factory(0)
	_, err := container.Get(0)
	test.ContainsError(containerTypes.KeyError, err, t)
	for i := 0; i < 5; i++ {
		container.Emplace(basic.Pair[int, int]{i, i})
	}
	for i := 0; i < 5; i++ {
		_v, err := container.Get(i)
		test.Eq(i, _v, t)
		test.Nil(err, t)
	}
	_, err = container.Get(-1)
	test.ContainsError(containerTypes.KeyError, err, t)
	_, err = container.Get(6)
	test.ContainsError(containerTypes.KeyError, err, t)
}

// Tests the GetPntr method functionality of a dynamic map.
func DynMapInterfaceGetPntr(
	factory func(capacity int) dynamicContainers.Map[int, int],
	t *testing.T,
) {
	container := factory(0)
	if container.IsAddressable() {
		_, err := container.GetPntr(0)
		test.ContainsError(containerTypes.KeyError, err, t)
		for i := 0; i < 5; i++ {
			container.Emplace(basic.Pair[int, int]{i, i})
		}
		for i := 0; i < 5; i++ {
			_v, err := container.GetPntr(i)
			test.Eq(i, *_v, t)
			test.Eq(nil, err, t)
		}
		_, err = container.GetPntr(-1)
		test.ContainsError(containerTypes.KeyError, err, t)
		_, err = container.GetPntr(6)
		test.ContainsError(containerTypes.KeyError, err, t)
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

// Tests the Set method functionality of a dynamic map.
func DynMapInterfaceSet(
	factory func(capacity int) dynamicContainers.Map[int, int],
	t *testing.T,
) {
	container := factory(0)
	err := container.Set(basic.Pair[int, int]{0, 6})
	test.ContainsError(containerTypes.KeyError, err, t)
	for i := 0; i < 5; i++ {
		container.Emplace(basic.Pair[int, int]{i, i})
	}
	for i := 0; i < 5; i++ {
		err := container.Set(basic.Pair[int, int]{i, i + 1})
		test.Nil(err, t)
	}
	for i := 0; i < 5; i++ {
		iterV, _ := container.Get(i)
		test.Eq(i+1, iterV, t)
	}
	err = container.Set(basic.Pair[int, int]{-1, 6})
	test.ContainsError(containerTypes.KeyError, err, t)
	err = container.Set(basic.Pair[int, int]{6, 6})
	test.ContainsError(containerTypes.KeyError, err, t)
}

func mapContainsHelper(
	v dynamicContainers.Map[int, int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Emplace(basic.Pair[int, int]{i, i})
	}
	for i := 0; i < l; i++ {
		test.True(v.Contains(i), t)
	}
	test.False(v.Contains(-1), t)
	test.False(v.Contains(l), t)
}

// Tests the Contains method functionality of a dynamic map.
func DynMapInterfaceContains(
	factory func(capacity int) dynamicContainers.Map[int, int],
	t *testing.T,
) {
	mapContainsHelper(factory(0), 0, t)
	mapContainsHelper(factory(0), 1, t)
	mapContainsHelper(factory(0), 2, t)
	mapContainsHelper(factory(0), 5, t)
}

func mapContainsPntrHelper(
	v dynamicContainers.Map[int, int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Emplace(basic.Pair[int, int]{i, i})
	}
	for i := 0; i < l; i++ {
		test.True(v.ContainsPntr(&i), t)
	}
	tmp := -1
	test.False(v.ContainsPntr(&tmp), t)
	test.False(v.ContainsPntr(&l), t)
}

// Tests the ContainsPntr method functionality of a dynamic map.
func DynMapInterfaceContainsPntr(
	factory func(capacity int) dynamicContainers.Map[int, int],
	t *testing.T,
) {
	mapContainsPntrHelper(factory(0), 0, t)
	mapContainsPntrHelper(factory(0), 1, t)
	mapContainsPntrHelper(factory(0), 2, t)
	mapContainsPntrHelper(factory(0), 5, t)
}

func mapKeyOfHelper(
	v dynamicContainers.Map[int, int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Emplace(basic.Pair[int, int]{i, i})
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

// Tests the KeyOf method functionality of a dynamic map.
func DynMapInterfaceKeyOf(
	factory func(capacity int) dynamicContainers.Map[int, int],
	t *testing.T,
) {
	mapKeyOfHelper(factory(0), 0, t)
	mapKeyOfHelper(factory(0), 1, t)
	mapKeyOfHelper(factory(0), 2, t)
	mapKeyOfHelper(factory(0), 5, t)
}

func mapKeyOfPntrHelper(
	v dynamicContainers.Map[int, int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Emplace(basic.Pair[int, int]{i, i})
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

// Tests the KeyOfPntr method functionality of a dynamic map.
func DynMapInterfaceKeyOfPntr(
	factory func(capacity int) dynamicContainers.Map[int, int],
	t *testing.T,
) {
	mapKeyOfPntrHelper(factory(0), 0, t)
	mapKeyOfPntrHelper(factory(0), 1, t)
	mapKeyOfPntrHelper(factory(0), 2, t)
	mapKeyOfPntrHelper(factory(0), 5, t)
}

func mapPopHelper(
	factory func(capacity int) dynamicContainers.Map[int, int],
	l int,
	t *testing.T,
) {
	// fmt.Println("Permutation: l: ",l," num: ",num)
	container := factory(0)
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			container.Emplace(basic.Pair[int, int]{i, -1})
		} else {
			container.Emplace(basic.Pair[int, int]{i, i})
		}
	}
	// fmt.Println("Init:   ",container)
	n := container.Pop(-1)
	// fmt.Println("After pop: ",container)
	cntr := 0
	expLength := 0
	for i := 0; i < l; i++ {
		if i%4 != 0 {
			expLength++
			// fmt.Println("Getting ",i)
			iterV, found := container.Get(i)
			test.Nil(found, t)
			test.Eq(i, iterV, t)
		} else {
			cntr++
		}
	}
	test.Eq(cntr, n, t)
	test.Eq(expLength, container.Length(), t)
}

// Tests the Pop method functionality of a dynamic map.
func DynMapInterfacePop(
	factory func(capacity int) dynamicContainers.Map[int, int],
	t *testing.T,
) {
	for i := 0; i < 13; i++ {
		mapPopHelper(factory, i, t)
	}
}

func mapPopPntrHelper(
	factory func(capacity int) dynamicContainers.Map[int, int],
	l int,
	t *testing.T,
) {
	// fmt.Println("Permutation: l: ",l," num: ",num)
	container := factory(0)
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			container.Emplace(basic.Pair[int, int]{i, -1})
		} else {
			container.Emplace(basic.Pair[int, int]{i, i})
		}
	}
	// fmt.Println("Init:   ",container)
	tmp := -1
	n := container.PopPntr(&tmp)
	// fmt.Println("After pop: ",container)
	cntr := 0
	expLength := 0
	for i := 0; i < l; i++ {
		if i%4 != 0 {
			expLength++
			// fmt.Println("Getting ",i)
			iterV, found := container.Get(i)
			test.Nil(found, t)
			test.Eq(i, iterV, t)
		} else {
			cntr++
		}
	}
	test.Eq(cntr, n, t)
	test.Eq(expLength, container.Length(), t)
}

// Tests the PopPntr method functionality of a dynamic map.
func DynMapInterfacePopPntr(
	factory func(capacity int) dynamicContainers.Map[int, int],
	t *testing.T,
) {
	for i := 0; i < 13; i++ {
		mapPopPntrHelper(factory, i, t)
	}
}

// Tests the Delete method functionality of a dynamic map.
func DynMapInterfaceDelete(
	factory func(capacity int) dynamicContainers.Map[int, int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 6; i++ {
		container.Emplace(basic.Pair[int, int]{i, i})
	}
	for i := container.Length() - 1; i >= 0; i-- {
		container.Delete(i)
		test.Eq(i, container.Length(), t)
		for j := 0; j < i; j++ {
			iterV, err := container.Get(j)
			test.Eq(j, iterV, t)
			test.Nil(err, t)
		}
	}
	err := container.Delete(0)
	test.ContainsError(containerTypes.KeyError, err, t)
}

// Tests the Clear method functionality of a dynamic map.
func DynMapInterfaceClear(
	factory func(capacity int) dynamicContainers.Map[int, int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 6; i++ {
		container.Emplace(basic.Pair[int, int]{i, i})
	}
	container.Clear()
	test.Eq(0, container.Length(), t)
}

func testHashMapKeysHelper(
	factory func(capacity int) dynamicContainers.Map[int, int],
	l int,
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < l; i++ {
		container.Emplace(basic.Pair[int, int]{i, i})
	}
	cnt := 0
	container.Keys().ForEach(func(index, val int) (iter.IteratorFeedback, error) {
		cnt++
		v, err := container.Get(val)
		test.Nil(err, t)
		test.Eq(val, v, t)
		return iter.Continue, nil
	})
	test.Eq(l, cnt, t)
}

// Tests the Keys method functionality of a dynamic map.
func DynMapInterfaceKeys(
	factory func(capacity int) dynamicContainers.Map[int, int],
	t *testing.T,
) {
	testHashMapKeysHelper(factory, 0, t)
	testHashMapKeysHelper(factory, 1, t)
	testHashMapKeysHelper(factory, 2, t)
	testHashMapKeysHelper(factory, 5, t)
}

func testHashMapValsHelper(
	factory func(capacity int) dynamicContainers.Map[int, int],
	l int,
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < l; i++ {
		container.Emplace(basic.Pair[int, int]{i, i})
	}
	cnt := 0
	container.Keys().ForEach(func(index, val int) (iter.IteratorFeedback, error) {
		cnt++
		test.True(container.Contains(val), t)
		return iter.Continue, nil
	})
	test.Eq(l, cnt, t)
}

// Tests the Vals method functionality of a dynamic map.
func DynMapInterfaceVals(
	factory func(capacity int) dynamicContainers.Map[int, int],
	t *testing.T,
) {
	testHashMapValsHelper(factory, 0, t)
	testHashMapValsHelper(factory, 1, t)
	testHashMapValsHelper(factory, 2, t)
	testHashMapValsHelper(factory, 5, t)
}

func testHashMapValPntrsHelper(
	factory func(capacity int) dynamicContainers.Map[int, int],
	l int,
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < l; i++ {
		container.Emplace(basic.Pair[int, int]{i, i})
	}
	cnt := 0
	container.ValPntrs().ForEach(func(index int, val *int) (iter.IteratorFeedback, error) {
		cnt++
		test.True(container.Contains(*val), t)
		return iter.Continue, nil
	})
	test.Eq(l, cnt, t)
}

// Tests the Vals method functionality of a dynamic map.
func DynMapInterfaceValPntrs(
	factory func(capacity int) dynamicContainers.Map[int, int],
	t *testing.T,
) {
	container := factory(0)
	if container.IsAddressable() {
		testHashMapValPntrsHelper(factory, 0, t)
		testHashMapValPntrsHelper(factory, 1, t)
		testHashMapValPntrsHelper(factory, 2, t)
		testHashMapValPntrsHelper(factory, 5, t)
	} else {
		test.Panics(func() { container.ValPntrs() }, t)
	}
}

// Tests the KeyedEq method functionality of a dynamic hash map.
func HashDynMapInterfaceKeyedEq(
	factory func(capacity int) dynamicContainers.Map[int, int],
	t *testing.T,
) {
	v := factory(0)
	v.Emplace(
		basic.Pair[int, int]{1, 1},
		basic.Pair[int, int]{2, 2},
		basic.Pair[int, int]{3, 3},
	)
	v2 := factory(0)
	v2.Emplace(
		basic.Pair[int, int]{1, 1},
		basic.Pair[int, int]{2, 2},
		basic.Pair[int, int]{3, 3},
	)
	test.True(v.KeyedEq(v2), t)
	test.True(v2.KeyedEq(v), t)
	v.Pop(3)
	test.False(v.KeyedEq(v2), t)
	test.False(v2.KeyedEq(v), t)
	v.Emplace(basic.Pair[int, int]{3, 3})
	v2 = factory(0)
	v2.Emplace(
		basic.Pair[int, int]{3, 3},
		basic.Pair[int, int]{1, 1},
		basic.Pair[int, int]{2, 2},
	)
	test.True(v.KeyedEq(v2), t)
	test.True(v2.KeyedEq(v), t)
	v.Pop(3)
	test.False(v.KeyedEq(v2), t)
	test.False(v2.KeyedEq(v), t)
	v = factory(0)
	v.Emplace(
		basic.Pair[int, int]{0, 0},
	)
	v2 = factory(0)
	v2.Emplace(
		basic.Pair[int, int]{0, 0},
	)
	test.True(v.KeyedEq(v2), t)
	test.True(v2.KeyedEq(v), t)
	v.Pop(0)
	test.False(v.KeyedEq(v2), t)
	test.False(v2.KeyedEq(v), t)
	v = factory(0)
	v2 = factory(0)
	test.True(v.KeyedEq(v2), t)
	test.True(v2.KeyedEq(v), t)
}
