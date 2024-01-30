package tests

import (
	"testing"

	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func vectorReadInterface[U any](c dynamicContainers.ReadVector[U])   {}
func vectorWriteInterface[U any](c dynamicContainers.WriteVector[U]) {}
func vectorInterface[U any](c dynamicContainers.Vector[U])           {}

// Tests that the value supplied by the factory implements the
// [containerTypes.RWSyncable] interface.
func VectorInterfaceSyncableInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.RWSyncable = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Length] interface.
func VectorInterfaceLengthInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.Length = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Capacity] interface.
func VectorInterfaceCapacityInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.Capacity = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Clear] interface.
func VectorInterfaceClearInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.WriteOps] interface.
func VectorInterfaceWriteOpsInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.WriteOps[int, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedWriteOps] interface.
func VectorInterfaceWriteKeyedOpsInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.WriteKeyedOps[int, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.ReadOps] interface.
func VectorInterfaceReadOpsInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.ReadOps[int, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedReadOps] interface.
func VectorInterfaceReadKeyedOpsInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.ReadKeyedOps[int, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.DeleteOps] interface.
func VectorInterfaceDeleteOpsInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.DeleteOps[int, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedDeleteOps] interface.
func VectorInterfaceDeleteKeyedOpsInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	var container containerTypes.DeleteKeyedOps[int, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.VectorRead] interface.
func ReadVectorInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	vectorReadInterface[V](factory())
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.WriteVector] interface.
func WriteVectorInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	vectorWriteInterface[V](factory())
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.Vector] interface.
func VectorInterfaceInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	vectorInterface[V](factory())
}

// Tests that the value supplied by the factory does not implement the
// [staticContainers.Vector] interface.
func VectorInterfaceStaticCapacityInterface[V any](
	factory func() dynamicContainers.Vector[V],
	t *testing.T,
) {
	test.Panics(
		func() {
			var c any
			c = factory()
			c2 := c.(containerTypes.StaticCapacity)
			_ = c2
		},
		"Code did not panic when casting a dynamic vector to a static vector.", t,
	)
}

// Tests the Get method functionality of a dynamic vector.
func VectorInterfaceGet(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory()
	_, err := container.Get(0)
	test.ContainsError(customerr.ValOutsideRange, err,
		"Get did not return the correct error with invalid index.", t,
	)
	for i := 0; i < 5; i++ {
		container.Append(i)
	}
	for i := 0; i < 5; i++ {
		_v, err := container.Get(i)
		test.BasicTest(i, _v,
			"Get did not return the correct value.", t,
		)
		test.BasicTest(nil, err,
			"Get returned an error when it shouldn't have.", t,
		)
	}
	_, err = container.Get(-1)
	test.ContainsError(customerr.ValOutsideRange, err,
		"Get did not return the correct error with invalid index.", t,
	)
	_, err = container.Get(6)
	test.ContainsError(customerr.ValOutsideRange, err,
		"Get did not return the correct error with invalid index.", t,
	)
}

// Tests the GetPntr method functionality of a dynamic vector.
func VectorInterfaceGetPntr(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory()
	_, err := container.GetPntr(0)
	test.ContainsError(customerr.ValOutsideRange, err,
		"Get did not return the correct error with invalid index.", t,
	)
	for i := 0; i < 5; i++ {
		container.Append(i)
	}
	for i := 0; i < 5; i++ {
		_v, err := container.GetPntr(i)
		test.BasicTest(i, *_v,
			"Get did not return the correct value.", t,
		)
		test.BasicTest(nil, err,
			"Get returned an error when it shouldn't have.", t,
		)
	}
	_, err = container.GetPntr(-1)
	test.ContainsError(customerr.ValOutsideRange, err,
		"Get pntr did not return the correct error with invalid index.", t,
	)
	_, err = container.GetPntr(6)
	test.ContainsError(customerr.ValOutsideRange, err,
		"Get pntr did not return the correct error with invalid index.", t,
	)
}

func vectorContainsHelper(
	v dynamicContainers.Vector[int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Append(i)
	}
	for i := 0; i < l; i++ {
		test.BasicTest(true, v.Contains(i),
			"Contains returned a false negative.", t,
		)
	}
	test.BasicTest(false, v.Contains(-1),
		"Contains returned a false positive.", t,
	)
	test.BasicTest(false, v.Contains(l),
		"Contains returned a false positive.", t,
	)
}

// Tests the Contains method functionality of a dynamic vector.
func VectorInterfaceContains(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	vectorContainsHelper(factory(), 0, t)
	vectorContainsHelper(factory(), 1, t)
	vectorContainsHelper(factory(), 2, t)
	vectorContainsHelper(factory(), 5, t)
}

func vectorContainsPntrHelper(
	v dynamicContainers.Vector[int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Append(i)
	}
	for i := 0; i < l; i++ {
		test.BasicTest(true, v.ContainsPntr(&i),
			"Contains returned a false negative.", t,
		)
	}
	tmp:=-1
	test.BasicTest(false, v.ContainsPntr(&tmp),
		"Contains returned a false positive.", t,
	)
	test.BasicTest(false, v.ContainsPntr(&l),
		"Contains returned a false positive.", t,
	)
}

// Tests the ContainsPntr method functionality of a dynamic vector.
func VectorInterfaceContainsPntr(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	vectorContainsHelper(factory(), 0, t)
	vectorContainsHelper(factory(), 1, t)
	vectorContainsHelper(factory(), 2, t)
	vectorContainsHelper(factory(), 5, t)
}

func vectorKeyOfHelper(
	v dynamicContainers.Vector[int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Append(i)
	}
	for i := 0; i < l; i++ {
		k, found := v.KeyOf(i)
		test.BasicTest(i, k,
			"KeyOf did not return the correct index.", t,
		)
		test.BasicTest(true, found,
			"KeyOf returned a false negative.", t,
		)
	}
	_, found := v.KeyOf(-1)
	test.BasicTest(false, found,
		"KeyOf returned a false positive.", t,
	)
	_, found = v.KeyOf(-1)
	test.BasicTest(false, v.Contains(l),
		"KeyOf returned a false positive.", t,
	)
}

// Tests the KeyOf method functionality of a dynamic vector.
func VectorInterfaceKeyOf(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	vectorKeyOfHelper(factory(), 0, t)
	vectorKeyOfHelper(factory(), 1, t)
	vectorKeyOfHelper(factory(), 2, t)
	vectorKeyOfHelper(factory(), 5, t)
}

// Tests the Emplace method functionality of a dynamic vector.
func VectorInterfaceEmplace(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory()
	err := container.Emplace(0, 6)
	test.ContainsError(customerr.ValOutsideRange, err,
		"Emplace did not return the correct error with invalid index.", t,
	)
	for i := 0; i < 5; i++ {
		container.Append(i)
	}
	for i := 0; i < 5; i++ {
		err := container.Emplace(i, i+1)
		test.BasicTest(nil, err,
			"Get returned an error when it shouldn't have.", t,
		)
	}
	for i := 0; i < 5; i++ {
		iterV, _ := container.Get(i)
		test.BasicTest(i+1, iterV,
			"Emplace did not set the value correctly.", t,
		)
	}
	err = container.Emplace(-1, 6)
	test.ContainsError(customerr.ValOutsideRange, err,
		"Emplace did not return the correct error with invalid index.", t,
	)
	err = container.Emplace(6, 6)
	test.ContainsError(customerr.ValOutsideRange, err,
		"Emplace did not return the correct error with invalid index.", t,
	)
}

// Tests the Append method functionality of a dynamic vector.
func VectorInterfaceAppend(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 5; i++ {
		container.Append(i)
	}
	for i := 0; i < 5; i++ {
		iterV, _ := container.Get(i)
		test.BasicTest(i, iterV,
			"Append did not add the value correctly.", t,
		)
	}
	container.Append(5, 6, 7)
	for i := 0; i < 8; i++ {
		iterV, _ := container.Get(i)
		test.BasicTest(i, iterV,
			"Append did not add the value correctly.", t,
		)
	}
}

func vectorPushHelper(
	v func() dynamicContainers.Vector[int],
	idx int,
	l int,
	t *testing.T,
) {
	container := v()
	for i := 0; i < l-1; i++ {
		container.Append(i)
	}
	err := container.Push(idx, l-1)
	test.BasicTest(nil, err,
		"Push returned an error when it shouldn't have.", t,
	)
	test.BasicTest(l, container.Length(),
		"Push did not increment the number of elements.", t,
	)
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
		test.BasicTest(exp, v,
			"Push did not put the value in the correct place.", t,
		)
	}
}

// Tests the Push method functionality of a dynamic vector.
func VectorInterfacePush(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory()
	for i := 2; i >= 0; i-- {
		container.Push(0, i)
	}
	for i := 3; i < 5; i++ {
		container.Push(container.Length(), i)
	}
	for i := 0; i < 5; i++ {
		iterV, _ := container.Get(i)
		test.BasicTest(i, iterV,
			"Push did not put the values in the correct place.", t,
		)
	}
	for i := 0; i < 5; i++ {
		vectorPushHelper(factory, i, 5, t)
	}
}

func vectorPopHelper(
	factory func() dynamicContainers.Vector[int],
	l int,
	num int,
	t *testing.T,
) {
	// fmt.Println("Permutation: l: ",l," num: ",num)
	container := factory()
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			container.Append(-1)
		} else {
			container.Append(i)
		}
	}
	// fmt.Println("Init:   ",v)
	n := container.Pop(-1, num)
	exp := factory()
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
	test.BasicTest(exp.Length(), container.Length(),
		"Pop did not remove value from the list correctly.", t,
	)
	test.BasicTest(cntr, n,
		"Pop did not pop the correct number of values.", t,
	)
	// fmt.Println("EXP:    ",exp)
	// fmt.Println("Final:  ",v)
	for i := 0; i < container.Length(); i++ {
		iterV, _ := container.Get(i)
		expIterV, _ := exp.Get(i)
		test.BasicTest(expIterV, iterV,
			"Pop did not shift the values correctly.", t,
		)
	}
}

// Tests the Pop method functionality of a dynamic vector.
func VectorInterfacePop(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	for i := 0; i < 13; i++ {
		for j := 0; j < 13; j++ {
			vectorPopHelper(factory, i, j, t)
		}
	}
}

// Tests the Delete method functionality of a dynamic vector.
func VectorInterfaceDelete(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 6; i++ {
		container.Append(i)
	}
	for i := container.Length() - 1; i >= 0; i-- {
		container.Delete(i)
		test.BasicTest(i, container.Length(), "Delete removed to many values.", t)
		for j := 0; j < i; j++ {
			iterV, _ := container.Get(j)
			test.BasicTest(j, iterV, "Delete changed the wrong value.", t)
		}
	}
	err := container.Delete(0)
	test.ContainsError(customerr.ValOutsideRange, err,
		"Delete returned an incorrect error.", t,
	)
}

// Tests the Clear method functionality of a dynamic vector.
func VectorInterfaceClear(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 6; i++ {
		container.Append(i)
	}
	container.Clear()
	test.BasicTest(0, container.Length(), "Clear did not reset the underlying vector.", t)
	test.BasicTest(0, container.Capacity(), "Clear did not reset the underlying vector.", t)
}

// Tests the UnorderedEq method functionality of a dynamic vector.
func VectorInterfaceUnorderedEq(
	factory func() dynamicContainers.Vector[int],
	t *testing.T,
) {
	v := factory()
	v.Append(1, 2, 3)
	v2 := factory()
	v2.Append(1, 2, 3)
	test.BasicTest(true, v.UnorderedEq(v2), 
		"UnorderedEq returned a false negative.", t,
	)
	test.BasicTest(true, v2.UnorderedEq(v), 
		"UnorderedEq returned a false negative.", t,
	)
	v.Pop(3,1)
	test.BasicTest(false, v.UnorderedEq(v2), 
		"UnorderedEq returned a false positive.", t,
	)
	test.BasicTest(false, v2.UnorderedEq(v), 
		"UnorderedEq returned a false positive.", t,
	)

	// v.Append(3)
	// v2 = factory()
	// v.Append(3, 1, 2)
	// test.BasicTest(true, v.UnorderedEq(v2), 
	// 	"UnorderedEq returned a false negative.", t,
	// )
	// test.BasicTest(true, v2.UnorderedEq(v), 
	// 	"UnorderedEq returned a false negative.", t,
	// )
	// v.Pop(3,1)
	// test.BasicTest(false, v.UnorderedEq(v2), 
	// 	"UnorderedEq returned a false positive.", t,
	// )
	// test.BasicTest(false, v2.UnorderedEq(v), 
	// 	"UnorderedEq returned a false positive.", t,
	// )
	// v.Append(3)
	// v2 = factory()
	// v.Append(2, 3, 1)
	// test.BasicTest(true, v.UnorderedEq(v2), 
	// 	"UnorderedEq returned a false negative.", t,
	// )
	// test.BasicTest(true, v2.UnorderedEq(v), 
	// 	"UnorderedEq returned a false negative.", t,
	// )
	// v.Pop(3,1)
	// test.BasicTest(false, v.UnorderedEq(v2), 
	// 	"UnorderedEq returned a false positive.", t,
	// )
	// test.BasicTest(false, v2.UnorderedEq(v), 
	// 	"UnorderedEq returned a false positive.", t,
	// )
	// v = factory()
	// v.Append(0)
	// v2 = factory()
	// v2.Append(0)
	// test.BasicTest(true, v.UnorderedEq(v2), 
	// 	"UnorderedEq returned a false negative.", t,
	// )
	// test.BasicTest(true, v2.UnorderedEq(v), 
	// 	"UnorderedEq returned a false negative.", t,
	// )
	// v.Delete(0)
	// test.BasicTest(false, v.UnorderedEq(v2), 
	// 	"UnorderedEq returned a false positive.", t,
	// )
	// test.BasicTest(false, v2.UnorderedEq(v), 
	// 	"UnorderedEq returned a false positive.", t,
	// )
	// v = factory()
	// v2 = factory()
	// test.BasicTest(true, v.UnorderedEq(v2), 
	// 	"UnorderedEq returned a false negative.", t,
	// )
	// test.BasicTest(true, v2.UnorderedEq(v), 
	// 	"UnorderedEq returned a false negative.", t,
	// )
}

// func VectorInterfaceNeq(t *testing.T){
//     v:=Vector[int,widgets.BuiltinInt]([]int{0,1,2,3})
//     v2:=Vector[int,widgets.BuiltinInt]([]int{0,1,2,3})
//     comp:=func(l *int, r *int) bool { return *l==*r }
//     test.BasicTest(false,v.Neq(&v2),
// 	"Neq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Neq(&v),
// 	"Neq returned a false positive.",t,
//     )
//     v.Delete(3)
//     test.BasicTest(true,v.Neq(&v2),
// 	"Neq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Neq(&v),
// 	"Neq returned a false negative.",t,
//     )
//     v=Vector[int,widgets.BuiltinInt]([]int{0})
//     v2=Vector[int,widgets.BuiltinInt]([]int{0})
//     test.BasicTest(false,v.Neq(&v2),
// 	"Neq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Neq(&v),
// 	"Neq returned a false positive.",t,
//     )
//     v.Delete(0)
//     test.BasicTest(true,v.Neq(&v2),
// 	"Neq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Neq(&v),
// 	"Neq returned a false negative.",t,
//     )
//     v=Vector[int,widgets.BuiltinInt]([]int{})
//     v2=Vector[int,widgets.BuiltinInt]([]int{})
//     test.BasicTest(false,v.Neq(&v2),
// 	"Neq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Neq(&v),
// 	"Neq returned a false positive.",t,
//     )
// }
