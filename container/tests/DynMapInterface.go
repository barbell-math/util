package tests

import (
	// "fmt"
	"fmt"
	"testing"
	//
	// 	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/dynamicContainers"

	// "github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func mapReadInterface[T any, U any](c dynamicContainers.ReadMap[T,U])   {}
func mapWriteInterface[T any, U any](c dynamicContainers.WriteMap[T,U]) {}
func mapInterface[T any, U any](c dynamicContainers.Map[T,U])           {}

// Tests that the value supplied by the factory implements the
// [containerTypes.RWSyncable] interface.
func MapInterfaceSyncableInterface[K any, V any](
	factory func() dynamicContainers.Map[K,V],
	t *testing.T,
) {
	var container containerTypes.RWSyncable = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Length] interface.
func MapInterfaceLengthInterface[K any, V any](
	factory func() dynamicContainers.Map[K,V],
	t *testing.T,
) {
	var container containerTypes.Length = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Clear] interface.
func MapInterfaceClearInterface[K any, V any](
	factory func() dynamicContainers.Map[K,V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedWriteOps] interface.
func MapInterfaceWriteKeyedOpsInterface[K any, V any](
	factory func() dynamicContainers.Map[K,V],
	t *testing.T,
) {
	var container containerTypes.WriteKeyedOps[K, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.ReadOps] interface.
func MapInterfaceReadOpsInterface[K any, V any](
	factory func() dynamicContainers.Map[K,V],
	t *testing.T,
) {
	var container containerTypes.ReadOps[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedReadOps] interface.
func MapInterfaceReadKeyedOpsInterface[K any, V any](
	factory func() dynamicContainers.Map[K,V],
	t *testing.T,
) {
	var container containerTypes.ReadKeyedOps[K, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.DeleteOps] interface.
func MapInterfaceDeleteOpsInterface[K any, V any](
	factory func() dynamicContainers.Map[K,V],
	t *testing.T,
) {
	var container containerTypes.DeleteOps[K, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.KeyedDeleteOps] interface.
func MapInterfaceDeleteKeyedOpsInterface[K any, V any](
	factory func() dynamicContainers.Map[K,V],
	t *testing.T,
) {
	var container containerTypes.DeleteKeyedOps[K, V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.MapRead] interface.
func ReadMapInterface[K any, V any](
	factory func() dynamicContainers.Map[K,V],
	t *testing.T,
) {
	mapReadInterface[K,V](factory())
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.WriteMap] interface.
func WriteMapInterface[K any, V any](
	factory func() dynamicContainers.Map[K,V],
	t *testing.T,
) {
	mapWriteInterface[K,V](factory())
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.Map] interface.
func MapInterfaceInterface[K any, V any](
	factory func() dynamicContainers.Map[K,V],
	t *testing.T,
) {
	mapInterface[K,V](factory())
}

// Tests that the value supplied by the factory does not implement the
// [staticContainers.Map] interface.
func MapInterfaceStaticCapacityInterface[K any, V any](
	factory func() dynamicContainers.Map[K,V],
	t *testing.T,
) {
	test.Panics(
		func() {
			var c any
			c = factory()
			c2 := c.(containerTypes.StaticCapacity)
			_ = c2
		},
		t,
	)
}

// Tests the Get method functionality of a dynamic map.
func MapInterfaceGet(
	factory func() dynamicContainers.Map[int,int],
	t *testing.T,
) {
	container := factory()
	_, err := container.Get(0)
	test.ContainsError(containerTypes.KeyError, err,t)
	for i := 0; i < 5; i++ {
		container.Emplace(basic.Pair[int, int]{i,i})
	}
	for i := 0; i < 5; i++ {
		_v, err := container.Get(i)
		test.Eq(i, _v,t)
		test.Nil(err,t)
	}
	_, err = container.Get(-1)
	test.ContainsError(containerTypes.KeyError, err,t)
	_, err = container.Get(6)
	test.ContainsError(containerTypes.KeyError, err,t)
}

// Tests the GetPntr method functionality of a dynamic map.
func MapInterfaceGetPntr(
	factory func() dynamicContainers.Map[int,int],
	t *testing.T,
) {
	container := factory()
	if container.IsAddressable() {
		_, err := container.GetPntr(0)
		test.ContainsError(containerTypes.KeyError, err,t)
		for i := 0; i < 5; i++ {
			container.Emplace(basic.Pair[int, int]{i,i})
		}
		for i := 0; i < 5; i++ {
			_v, err := container.GetPntr(i)
			test.Eq(i, *_v,t)
			test.Eq(nil, err,t)
		}
		_, err = container.GetPntr(-1)
		test.ContainsError(containerTypes.KeyError, err,t)
		_, err = container.GetPntr(6)
		test.ContainsError(containerTypes.KeyError, err,t)
	} else {
		test.Panics(
			func() {
				container:=factory()
				container.GetPntr(1)
			},
			t,
		)
	}
}

// Tests the Set method functionality of a dynamic map.
func MapInterfaceSet(
	factory func() dynamicContainers.Map[int,int],
	t *testing.T,
) {
	container := factory()
	err := container.Set(basic.Pair[int, int]{0,6})
	test.ContainsError(containerTypes.KeyError, err,t)
	for i := 0; i < 5; i++ {
		container.Emplace(basic.Pair[int, int]{i,i})
	}
	for i := 0; i < 5; i++ {
		err := container.Set(basic.Pair[int,int]{i, i+1})
		test.Nil(err,t)
	}
	for i := 0; i < 5; i++ {
		iterV, _ := container.Get(i)
		test.Eq(i+1, iterV,t)
	}
	err = container.Set(basic.Pair[int,int]{-1, 6})
	test.ContainsError(containerTypes.KeyError, err,t)
	err = container.Set(basic.Pair[int,int]{6, 6})
	test.ContainsError(containerTypes.KeyError, err,t)
}

func mapContainsHelper(
	v dynamicContainers.Map[int,int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Emplace(basic.Pair[int, int]{i,i})
	}
	for i := 0; i < l; i++ {
		test.True(v.Contains(i),t)
	}
	test.False(v.Contains(-1),t)
	test.False(v.Contains(l),t)
}
// Tests the Contains method functionality of a dynamic map.
func MapInterfaceContains(
	factory func() dynamicContainers.Map[int,int],
	t *testing.T,
) {
	mapContainsHelper(factory(), 0, t)
	mapContainsHelper(factory(), 1, t)
	mapContainsHelper(factory(), 2, t)
	mapContainsHelper(factory(), 5, t)
}

func mapContainsPntrHelper(
	v dynamicContainers.Map[int,int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Emplace(basic.Pair[int, int]{i,i})
	}
	for i := 0; i < l; i++ {
		test.True(v.ContainsPntr(&i),t)
	}
	tmp:=-1
	test.False(v.ContainsPntr(&tmp),t)
	test.False(v.ContainsPntr(&l),t)
}
// Tests the ContainsPntr method functionality of a dynamic map.
func MapInterfaceContainsPntr(
	factory func() dynamicContainers.Map[int,int],
	t *testing.T,
) {
	mapContainsHelper(factory(), 0, t)
	mapContainsHelper(factory(), 1, t)
	mapContainsHelper(factory(), 2, t)
	mapContainsHelper(factory(), 5, t)
}

func mapKeyOfHelper(
	v dynamicContainers.Map[int,int],
	l int,
	t *testing.T,
) {
	for i := 0; i < l; i++ {
		v.Emplace(basic.Pair[int, int]{i,i})
	}
	for i := 0; i < l; i++ {
		k, found := v.KeyOf(i)
		test.Eq(i, k,t)
		test.True(found,t)
	}
	_, found := v.KeyOf(-1)
	test.False(found,t)
	_, found = v.KeyOf(-1)
	test.False(v.Contains(l),t)
}
// Tests the KeyOf method functionality of a dynamic map.
func MapInterfaceKeyOf(
	factory func() dynamicContainers.Map[int,int],
	t *testing.T,
) {
	mapKeyOfHelper(factory(), 0, t)
	mapKeyOfHelper(factory(), 1, t)
	mapKeyOfHelper(factory(), 2, t)
	mapKeyOfHelper(factory(), 5, t)
}

func mapPopHelper(
	factory func() dynamicContainers.Map[int,int],
	l int,
	t *testing.T,
) {
	// fmt.Println("Permutation: l: ",l," num: ",num)
	container := factory()
	for i := 0; i < l; i++ {
		if i%4 == 0 {
			container.Emplace(basic.Pair[int, int]{i,-1})
		} else {
			container.Emplace(basic.Pair[int, int]{i,i})
		}
	}
	fmt.Println("Init:   ",container)
	n := container.Pop(-1)
	fmt.Println("After pop: ",container)
	// exp := factory()
	cntr := 0
	expLength:=0
	for i := 0; i < l; i++ {
		if i%4 != 0 {
			// exp.Emplace(basic.Pair[int, int]{i,i})
			expLength++
			fmt.Println("Getting ",i)
			iterV, found:=container.Get(i)
			test.Nil(found,t)
			test.Eq(i,iterV,t)
		} else {
			cntr++
		}
	}
	test.Eq(cntr,n,t)
	test.Eq(expLength, container.Length(),t)
}
// Tests the Pop method functionality of a dynamic map.
func MapInterfacePop(
	factory func() dynamicContainers.Map[int,int],
	t *testing.T,
) {
	for i := 0; i < 13; i++ {
		mapPopHelper(factory, i, t)
	}
}

// Tests the Delete method functionality of a dynamic map.
func MapInterfaceDelete(
	factory func() dynamicContainers.Map[int,int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 6; i++ {
		container.Emplace(basic.Pair[int, int]{i,i})
	}
	for i := container.Length() - 1; i >= 0; i-- {
		container.Delete(i)
		test.Eq(i, container.Length(), t)
		for j := 0; j < i; j++ {
			iterV, err := container.Get(j)
			test.Eq(j, iterV, t)
			test.Nil(err,t)
		}
	}
	err := container.Delete(0)
	test.ContainsError(containerTypes.KeyError, err,t)
}

// Tests the Clear method functionality of a dynamic map.
func MapInterfaceClear(
	factory func() dynamicContainers.Map[int,int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 6; i++ {
		container.Emplace(basic.Pair[int,int]{i,i})
	}
	container.Clear()
	test.Eq(0, container.Length(), t)
}

// func mapContainsHelper(
// 	v dynamicContainers.Map[int,int],
// 	l int,
// 	t *testing.T,
// ) {
// 	for i := 0; i < l; i++ {
// 		v.Emplace(i,i)
// 	}
// 	for i := 0; i < l; i++ {
// 		test.BasicTest(true, v.Contains(i),
// 			"Contains returned a false negative.", t,
// 		)
// 	}
// 	test.BasicTest(false, v.Contains(-1),
// 		"Contains returned a false positive.", t,
// 	)
// 	test.BasicTest(false, v.Contains(l),
// 		"Contains returned a false positive.", t,
// 	)
// }
// 
// // Tests the Contains method functionality of a dynamic map.
// func MapInterfaceContains(
// 	factory func() dynamicContainers.Map[int,int],
// 	t *testing.T,
// ) {
// 	mapContainsHelper(factory(), 0, t)
// 	mapContainsHelper(factory(), 1, t)
// 	mapContainsHelper(factory(), 2, t)
// 	mapContainsHelper(factory(), 5, t)
// }
// 
// func mapContainsPntrHelper(
// 	v dynamicContainers.Map[int,int],
// 	l int,
// 	t *testing.T,
// ) {
// 	for i := 0; i < l; i++ {
// 		v.Emplace(i,i)
// 	}
// 	for i := 0; i < l; i++ {
// 		test.BasicTest(true, v.ContainsPntr(&i),
// 			"Contains returned a false negative.", t,
// 		)
// 	}
// 	tmp:=-1
// 	test.BasicTest(false, v.ContainsPntr(&tmp),
// 		"Contains returned a false positive.", t,
// 	)
// 	test.BasicTest(false, v.ContainsPntr(&l),
// 		"Contains returned a false positive.", t,
// 	)
// }
// 
// // Tests the ContainsPntr method functionality of a dynamic map.
// func MapInterfaceContainsPntr(
// 	factory func() dynamicContainers.Map[int,int],
// 	t *testing.T,
// ) {
// 	mapContainsHelper(factory(), 0, t)
// 	mapContainsHelper(factory(), 1, t)
// 	mapContainsHelper(factory(), 2, t)
// 	mapContainsHelper(factory(), 5, t)
// }
// 
// func mapKeyOfHelper(
// 	v dynamicContainers.Map[int,int],
// 	l int,
// 	t *testing.T,
// ) {
// 	for i := 0; i < l; i++ {
// 		v.Emplace(i,i)
// 	}
// 	for i := 0; i < l; i++ {
// 		k, found := v.KeyOf(i)
// 		test.BasicTest(i, k,
// 			"KeyOf did not return the correct index.", t,
// 		)
// 		test.BasicTest(true, found,
// 			"KeyOf returned a false negative.", t,
// 		)
// 	}
// 	_, found := v.KeyOf(-1)
// 	test.BasicTest(false, found,
// 		"KeyOf returned a false positive.", t,
// 	)
// 	_, found = v.KeyOf(-1)
// 	test.BasicTest(false, v.Contains(l),
// 		"KeyOf returned a false positive.", t,
// 	)
// }
// 
// // Tests the KeyOf method functionality of a dynamic map.
// func MapInterfaceKeyOf(
// 	factory func() dynamicContainers.Map[int,int],
// 	t *testing.T,
// ) {
// 	mapKeyOfHelper(factory(), 0, t)
// 	mapKeyOfHelper(factory(), 1, t)
// 	mapKeyOfHelper(factory(), 2, t)
// 	mapKeyOfHelper(factory(), 5, t)
// }
// 
// // Tests the Emplace method functionality of a dynamic map.
// func MapInterfaceEmplace(
// 	factory func() dynamicContainers.Map[int,int],
// 	t *testing.T,
// ) {
// 	container := factory()
// 	err := container.Emplace(0, 6)
// 	test.ContainsError(customerr.ValOutsideRange, err,
// 		"Emplace did not return the correct error with invalid index.", t,
// 	)
// 	for i := 0; i < 5; i++ {
// 		container.Emplace(i,i)
// 	}
// 	for i := 0; i < 5; i++ {
// 		err := container.Emplace(i, i+1)
// 		test.BasicTest(nil, err,
// 			"Get returned an error when it shouldn't have.", t,
// 		)
// 	}
// 	for i := 0; i < 5; i++ {
// 		iterV, _ := container.Get(i)
// 		test.BasicTest(i+1, iterV,
// 			"Emplace did not set the value correctly.", t,
// 		)
// 	}
// 	err = container.Emplace(-1, 6)
// 	test.ContainsError(customerr.ValOutsideRange, err,
// 		"Emplace did not return the correct error with invalid index.", t,
// 	)
// 	err = container.Emplace(6, 6)
// 	test.ContainsError(customerr.ValOutsideRange, err,
// 		"Emplace did not return the correct error with invalid index.", t,
// 	)
// }
// 
// // TODO - REMOVE, DOES NOT MAKE SENSE
// func mapPushHelper(
// 	v func() dynamicContainers.Map[int,int],
// 	idx int,
// 	l int,
// 	t *testing.T,
// ) {
// 	container := v()
// 	for i := 0; i < l-1; i++ {
// 		container.Emplace(i,i)
// 	}
// 	err := container.Push(idx, l-1)
// 	test.BasicTest(nil, err,
// 		"Push returned an error when it shouldn't have.", t,
// 	)
// 	test.BasicTest(l, container.Length(),
// 		"Push did not increment the number of elements.", t,
// 	)
// 	for i := 0; i < container.Length(); i++ {
// 		var exp int
// 		v, _ := container.Get(i)
// 		if i < idx {
// 			exp = i
// 		} else if i == idx {
// 			exp = l - 1
// 		} else {
// 			exp = i - 1
// 		}
// 		test.BasicTest(exp, v,
// 			"Push did not put the value in the correct place.", t,
// 		)
// 	}
// }
// 
// // Tests the Push method functionality of a dynamic map.
// func MapInterfacePush(
// 	factory func() dynamicContainers.Map[int,int],
// 	t *testing.T,
// ) {
// 	container := factory()
// 	for i := 2; i >= 0; i-- {
// 		container.Push(0, i)
// 	}
// 	for i := 3; i < 5; i++ {
// 		container.Push(container.Length(), i)
// 	}
// 	for i := 0; i < 5; i++ {
// 		iterV, _ := container.Get(i)
// 		test.BasicTest(i, iterV,
// 			"Push did not put the values in the correct place.", t,
// 		)
// 	}
// 	for i := 0; i < 5; i++ {
// 		mapPushHelper(factory, i, 5, t)
// 	}
// }
// 
// func mapPopHelper(
// 	factory func() dynamicContainers.Map[int,int],
// 	l int,
// 	num int,
// 	t *testing.T,
// ) {
// 	// fmt.Println("Permutation: l: ",l," num: ",num)
// 	container := factory()
// 	for i := 0; i < l; i++ {
// 		if i%4 == 0 {
// 			container.Emplace(i,-1)
// 		} else {
// 			container.Emplace(i,i)
// 		}
// 	}
// 	// fmt.Println("Init:   ",v)
// 	n := container.Pop(-1, num)
// 	exp := factory()
// 	cntr := 0
// 	for i := 0; i < l; i++ {
// 		if i%4 == 0 {
// 			if cntr < num {
// 				cntr++
// 				continue
// 			} else {
// 				exp.Emplace(i,-1)
// 			}
// 		} else {
// 			exp.Emplace(i,i)
// 		}
// 	}
// 	test.BasicTest(exp.Length(), container.Length(),
// 		"Pop did not remove value from the list correctly.", t,
// 	)
// 	test.BasicTest(cntr, n,
// 		"Pop did not pop the correct number of values.", t,
// 	)
// 	// fmt.Println("EXP:    ",exp)
// 	// fmt.Println("Final:  ",v)
// 	for i := 0; i < container.Length(); i++ {
// 		iterV, _ := container.Get(i)
// 		expIterV, _ := exp.Get(i)
// 		test.BasicTest(expIterV, iterV,
// 			"Pop did not shift the values correctly.", t,
// 		)
// 	}
// }
// 
// // Tests the Pop method functionality of a dynamic map.
// func MapInterfacePop(
// 	factory func() dynamicContainers.Map[int,int],
// 	t *testing.T,
// ) {
// 	for i := 0; i < 13; i++ {
// 		for j := 0; j < 13; j++ {
// 			mapPopHelper(factory, i, j, t)
// 		}
// 	}
// }
// 
// // Tests the Delete method functionality of a dynamic map.
// func MapInterfaceDelete(
// 	factory func() dynamicContainers.Map[int,int],
// 	t *testing.T,
// ) {
// 	container := factory()
// 	for i := 0; i < 6; i++ {
// 		container.Emplace(i,i)
// 	}
// 	for i := container.Length() - 1; i >= 0; i-- {
// 		container.Delete(i)
// 		test.BasicTest(i, container.Length(), "Delete removed to many values.", t)
// 		for j := 0; j < i; j++ {
// 			iterV, _ := container.Get(j)
// 			test.BasicTest(j, iterV, "Delete changed the wrong value.", t)
// 		}
// 	}
// 	err := container.Delete(0)
// 	test.ContainsError(customerr.ValOutsideRange, err,
// 		"Delete returned an incorrect error.", t,
// 	)
// }
// 
// // Tests the Clear method functionality of a dynamic map.
// func MapInterfaceClear(
// 	factory func() dynamicContainers.Map[int,int],
// 	t *testing.T,
// ) {
// 	container := factory()
// 	for i := 0; i < 6; i++ {
// 		container.Emplace(i,i)
// 	}
// 	container.Clear()
// 	test.BasicTest(0, container.Length(), "Clear did not reset the underlying map.", t)
// }
// 
// func testMapValsHelper(
//     factory func() dynamicContainers.Map[int,int],
//     l int, 
//     t *testing.T,
// ){
// 	container:=factory()
//     for i:=0; i<l; i++ {
//         container.Emplace(i,i);
//     }
//     cnt:=0
//     container.Vals().ForEach(func(index, val int) (iter.IteratorFeedback, error) {
//         cnt++
//         test.BasicTest(index,val,"Element was skipped while iterating.",t);
//         return iter.Continue,nil;
//     });
//     test.BasicTest(l,cnt,
//         "All the elements were not iterated over.",t,
//     )
// }
// func TestMapVals(
// 	factory func() dynamicContainers.Map[int,int],
// 	t *testing.T,
// ){
//     testMapValsHelper(factory,0,t);
//     testMapValsHelper(factory,1,t);
//     testMapValsHelper(factory,2,t);
//     testMapValsHelper(factory,5,t);
// }
// 
// func testMapPntrValsHelper(
// 	factory func() dynamicContainers.Map[int,int],
//     l int, 
//     t *testing.T,
// ){
// 	container:=factory()
//     for i:=0; i<l; i++ {
//         container.Emplace(i,i);
//     }
//     cnt:=0
//     container.ValPntrs().ForEach(func(index int, val *int) (iter.IteratorFeedback, error) {
//         cnt++
//         test.BasicTest(index,*val,"Element was skipped while iterating.",t);
//         *val=100;
//         return iter.Continue,nil;
//     });
//     container.Vals().ForEach(func(index int, val int) (iter.IteratorFeedback, error) {
//         test.BasicTest(100,val,"Element was not updated while iterating.",t);
//         return iter.Continue,nil;
//     });
//     test.BasicTest(l,cnt,
//         "All the elements were not iterated over.",t,
//     )
// }
// func TestMapValPntrs(
// 	factory func() dynamicContainers.Map[int,int],
// 	t *testing.T,
// ){
//     testMapPntrValsHelper(factory,0,t);
//     testMapPntrValsHelper(factory,1,t);
//     testMapPntrValsHelper(factory,2,t);
// }
// 
// func testMapKeysHelper(
//     factory func() dynamicContainers.Map[int,int],
//     l int, 
//     t *testing.T,
// ){
// 	container:=factory()
//     for i:=0; i<l; i++ {
//         container.Emplace(i,i);
//     }
//     cnt:=0
//     container.Keys().ForEach(func(index, val int) (iter.IteratorFeedback, error) {
//         cnt++
//         test.BasicTest(index,val,"Keys were skipped while iterating.",t);
//         return iter.Continue,nil;
//     });
//     test.BasicTest(l,cnt,
//         "All the keys were not iterated over.",t,
//     )
// }
// func TestMapKeys(
// 	factory func() dynamicContainers.Map[int,int],
// 	t *testing.T,
// ){
//     testMapKeysHelper(factory,0,t);
//     testMapKeysHelper(factory,1,t);
//     testMapKeysHelper(factory,2,t);
//     testMapKeysHelper(factory,5,t);
// }
// 
// // // Tests the KeyedEq method functionality of a dynamic map.
// // func MapInterfaceKeyedEq(
// // 	factory func() dynamicContainers.Map[int,int],
// // 	t *testing.T,
// // ) {
// // 	v := factory()
// // 	v.Emplace(0,1)	// TODO - replace with variadic emplace
// // 	v.Emplace(1,2)
// // 	v.Emplace(2,3)
// // 	v2 := factory()
// // 	v2.Append(1, 2, 3)
// // 	test.BasicTest(true, v.KeyedEq(v2), 
// // 		"KeyedEq returned a false negative.", t,
// // 	)
// // 	test.BasicTest(true, v2.KeyedEq(v), 
// // 		"KeyedEq returned a false negative.", t,
// // 	)
// // 	v.Pop(3,1)
// // 	test.BasicTest(false, v.KeyedEq(v2), 
// // 		"KeyedEq returned a false positive.", t,
// // 	)
// // 	test.BasicTest(false, v2.KeyedEq(v), 
// // 		"KeyedEq returned a false positive.", t,
// // 	)
// // 	v.Append(3)
// // 	v2 = factory()
// // 	v2.Append(3, 1, 2)
// // 	test.BasicTest(false, v.KeyedEq(v2), 
// // 		"KeyedEq returned a false positive.", t,
// // 	)
// // 	test.BasicTest(false, v2.KeyedEq(v), 
// // 		"KeyedEq returned a false positive.", t,
// // 	)
// // 	v.Pop(3,1)
// // 	test.BasicTest(false, v.KeyedEq(v2), 
// // 		"KeyedEq returned a false positive.", t,
// // 	)
// // 	test.BasicTest(false, v2.KeyedEq(v), 
// // 		"KeyedEq returned a false positive.", t,
// // 	)
// // 	v = factory()
// // 	v.Append(0)
// // 	v2 = factory()
// // 	v2.Append(0)
// // 	test.BasicTest(true, v.KeyedEq(v2), 
// // 		"KeyedEq returned a false negative.", t,
// // 	)
// // 	test.BasicTest(true, v2.KeyedEq(v), 
// // 		"KeyedEq returned a false negative.", t,
// // 	)
// // 	v.Pop(0,1)
// // 	test.BasicTest(false, v.KeyedEq(v2), 
// // 		"KeyedEq returned a false positive.", t,
// // 	)
// // 	test.BasicTest(false, v2.KeyedEq(v), 
// // 		"KeyedEq returned a false positive.", t,
// // 	)
// // 	v = factory()
// // 	v2 = factory()
// // 	test.BasicTest(true, v.KeyedEq(v2), 
// // 		"KeyedEq returned a false negative.", t,
// // 	)
// // 	test.BasicTest(true, v2.KeyedEq(v), 
// // 		"KeyedEq returned a false negative.", t,
// // 	)
// // }
