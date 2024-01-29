package tests

import (
	"testing"

	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func dequeReadInterface[T ~int, U any](c dynamicContainers.ReadDeque[T, U])   {}
func dequeWriteInterface[T ~int, U any](c dynamicContainers.WriteDeque[T, U]) {}
func dequeInterface[T ~int, U any](c dynamicContainers.Deque[T, U])           {}

// Tests that the value supplied by the factory implements the 
// [containerTypes.RWSyncable] interface.
func DequeInterfaceSyncableInterface[K ~int, V any](
	factory func() dynamicContainers.Deque[K, V],
	t *testing.T,
) {
	var container containerTypes.RWSyncable = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Length] interface.
func DequeInterfaceLengthInterface[K ~int, V any](
	factory func() dynamicContainers.Deque[K, V],
	t *testing.T,
) {
	var container containerTypes.Length = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Capacity] interface.
func DequeInterfaceCapacityInterface[K ~int, V any](
	factory func() dynamicContainers.Deque[K, V],
	t *testing.T,
) {
	var container containerTypes.Capacity = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Clear] interface.
func DequeInterfaceClearInterface[K ~int, V any](
	factory func() dynamicContainers.Deque[K, V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.FirstElemRead] interface.
func DequeInterfaceFirstElemReadInterface[K ~int, V any](
	factory func() dynamicContainers.Deque[K, V],
	t *testing.T,
) {
	var container containerTypes.FirstElemRead[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.FirstElemWrite] interface.
func DequeInterfaceFirstElemWriteInterface[K ~int, V any](
	factory func() dynamicContainers.Deque[K, V],
	t *testing.T,
) {
	var container containerTypes.FirstElemWrite[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.FirstElemDelete] interface.
func DequeInterfaceFirstElemDeleteInterface[K ~int, V any](
	factory func() dynamicContainers.Deque[K, V],
	t *testing.T,
) {
	var container containerTypes.FirstElemDelete[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.LastElemRead] interface.
func DequeInterfaceLastElemReadInterface[K ~int, V any](
	factory func() dynamicContainers.Deque[K, V],
	t *testing.T,
) {
	var container containerTypes.LastElemRead[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.LastElemWrite] interface.
func DequeInterfaceLastElemWriteInterface[K ~int, V any](
	factory func() dynamicContainers.Deque[K, V],
	t *testing.T,
) {
	var container containerTypes.LastElemWrite[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.LastElemDelete] interface.
func DequeInterfaceLastElemDeleteInterface[K ~int, V any](
	factory func() dynamicContainers.Deque[K, V],
	t *testing.T,
) {
	var container containerTypes.LastElemDelete[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.DequeRead] interface.
func ReadDequeInterface[K ~int, V any](
	factory func() dynamicContainers.Deque[K, V],
	t *testing.T,
) {
	dequeReadInterface[K, V](factory())
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.WriteDeque] interface.
func WriteDequeInterface[K ~int, V any](
	factory func() dynamicContainers.Deque[K, V],
	t *testing.T,
) {
	dequeWriteInterface[K, V](factory())
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.Deque] interface.
func DequeInterfaceInterface[K ~int, V any](
	factory func() dynamicContainers.Deque[K, V],
	t *testing.T,
) {
	dequeInterface[K, V](factory())
}

// Tests that the value supplied by the factory does not implement the 
// [staticContainers.Deque] interface.
func DequeInterfaceStaticCapacityInterface[K ~int, V any](
	factory func() dynamicContainers.Deque[K, V],
	t *testing.T,
) {
	test.Panics(
		func() {
			var c any
			c = factory()
			c2 := c.(containerTypes.StaticCapacity)
			_ = c2
		},
		"Code did not panic when casting a dynamic Deque to a static vector.", t,
	)
}

// Tests the Clear method functionality of a dynamic Deque.
func DequeInterfaceClear(
	factory func() dynamicContainers.Deque[int, int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 6; i++ {
		container.PushBack(i)
	}
	container.Clear()
	test.BasicTest(0, container.Length(), "Clear did not reset the underlying Deque.", t)
	test.BasicTest(0, container.Capacity(), "Clear did not reset the underlying Deque.", t)
}

// Tests the PeekPntrFront method functionality of a dynamic Deque.
func DequeInterfacePeekPntrFront(
	factory func() dynamicContainers.Deque[int, int],
	t *testing.T,
) {
	container := factory()
	_v, err := container.PeekPntrFront()
	test.BasicTest((*int)(nil), _v,
		"Peek pntr front did not return the correct value.", t,
	)
	test.ContainsError(customerr.ValOutsideRange, err,
		"Peek pntr front returned an incorrect error.", t,
	)
	container.PushBack(1)
	_v, err = container.PeekPntrFront()
	test.BasicTest(1, *_v,
		"Peek pntr front did not return the correct value.", t,
	)
	test.BasicTest(nil, err,
		"Peek pntr front returned an error when it shouldn't have.", t,
	)
	container.PushBack(2)
	_v, err = container.PeekPntrFront()
	test.BasicTest(1, *_v,
		"Peek pntr front did not return the correct value.", t,
	)
	test.BasicTest(nil, err,
		"Peek pntr front returned an error when it shouldn't have.", t,
	)
}

// Tests the PeekFront method functionality of a dynamic Deque.
func DequeInterfacePeekFront(
	factory func() dynamicContainers.Deque[int, int],
	t *testing.T,
) {
	container := factory()
	_, err := container.PeekFront()
	test.ContainsError(customerr.ValOutsideRange, err,
		"Peek front returned an incorrect error.", t,
	)
	container.PushBack(1)
	_v, err := container.PeekFront()
	test.BasicTest(1, _v,
		"Peek front did not return the correct value.", t,
	)
	test.BasicTest(nil, err,
		"Peek front returned an error when it shouldn't have.", t,
	)
	container.PushBack(2)
	_v, err = container.PeekFront()
	test.BasicTest(1, _v,
		"Peek front did not return the correct value.", t,
	)
	test.BasicTest(nil, err,
		"Peek front returned an error when it shouldn't have.", t,
	)
}

// Tests the PeekPntrBack method functionality of a dynamic Deque.
func DequeInterfacePeekPntrBack(
	factory func() dynamicContainers.Deque[int, int],
	t *testing.T,
) {
	container := factory()
	_v, err := container.PeekPntrBack()
	test.BasicTest((*int)(nil), _v,
		"Peek pntr back did not return the correct value.", t,
	)
	test.ContainsError(customerr.ValOutsideRange, err,
		"Peek pntr back returned an incorrect error.", t,
	)
	container.PushBack(1)
	_v, err = container.PeekPntrBack()
	test.BasicTest(1, *_v,
		"Peek pntr back did not return the correct value.", t,
	)
	test.BasicTest(nil, err,
		"Peek pntr back returned an error when it shouldn't have.", t,
	)
	container.PushBack(2)
	_v, err = container.PeekPntrBack()
	test.BasicTest(2, *_v,
		"Peek pntr back did not return the correct value.", t,
	)
	test.BasicTest(nil, err,
		"Peek pntr back returned an error when it shouldn't have.", t,
	)
}

// Tests the PeekBack method functionality of a dynamic Deque.
func DequeInterfacePeekBack(
	factory func() dynamicContainers.Deque[int, int],
	t *testing.T,
) {
	container := factory()
	_, err := container.PeekBack()
	test.ContainsError(customerr.ValOutsideRange, err,
		"Peek back returned an incorrect error.", t,
	)
	container.PushBack(1)
	_v, err := container.PeekBack()
	test.BasicTest(1, _v,
		"Peek back did not return the correct value.", t,
	)
	test.BasicTest(nil, err,
		"Peek back returned an error when it shouldn't have.", t,
	)
	container.PushBack(2)
	_v, err = container.PeekBack()
	test.BasicTest(2, _v,
		"Peek back did not return the correct value.", t,
	)
	test.BasicTest(nil, err,
		"Peek back returned an error when it shouldn't have.", t,
	)
}

// Tests the PopFront method functionality of a dynamic Deque.
func DequeInterfacePopFront(
	factory func() dynamicContainers.Deque[int, int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 4; i++ {
		container.PushBack(i)
	}
	for i := 0; i < 4; i++ {
		f, err := container.PopFront()
		test.BasicTest(i, f,
			"Pop front returned the incorrect value.", t,
		)
		test.BasicTest(nil, err,
			"Pop front returned an error when it shoudn't have.", t,
		)
	}
	_, err := container.PopFront()
	test.ContainsError(containerTypes.Empty, err,
		"Pop front returned an incorrect error.", t,
	)
}

// Tests the PopBack method functionality of a dynamic Deque.
func DequeInterfacePopBack(
	factory func() dynamicContainers.Deque[int, int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 4; i++ {
		container.PushBack(i)
	}
	for i := 3; i >= 0; i-- {
		f, err := container.PopBack()
		test.BasicTest(i, f,
			"Pop front returned the incorrect value.", t,
		)
		test.BasicTest(nil, err,
			"Pop front returned an error when it shoudn't have.", t,
		)
	}
	_, err := container.PopBack()
	test.ContainsError(containerTypes.Empty, err,
		"Pop front returned an incorrect error.", t,
	)
}

// Tests the PopFront method functionality of a dynamic Deque.
func DequeInterfacePushFront(
	factory func() dynamicContainers.Deque[int, int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 4; i++ {
		container.PushFront(i)
		test.BasicTest(i+1, container.Length(),
			"Push front did not add the value correctly.", t,
		)
		iterV, _ := container.PeekFront()
		test.BasicTest(i, iterV,
			"Push front did not put the value in the correct place.", t,
		)
		iterV, _ = container.PeekBack()
		test.BasicTest(0, iterV,
			"Push front did not put the value in the correct place.", t,
		)
	}
}

// Tests the ForcePopFront method functionality of a dynamic Deque.
func DequeInterfaceForcePushFront(
	factory func() dynamicContainers.Deque[int, int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 4; i++ {
		container.ForcePushFront(i)
		test.BasicTest(i+1, container.Length(),
			"Push front did not add the value correctly.", t,
		)
		iterV, _ := container.PeekFront()
		test.BasicTest(i, iterV,
			"Push front did not put the value in the correct place.", t,
		)
		iterV, _ = container.PeekBack()
		test.BasicTest(0, iterV,
			"Push front did not put the value in the correct place.", t,
		)
	}
}

// Tests the PushBack method functionality of a dynamic Deque.
func DequeInterfacePushBack(
	factory func() dynamicContainers.Deque[int, int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 4; i++ {
		container.PushBack(i)
		test.BasicTest(i+1, container.Length(),
			"Push back did not add the value correctly.", t,
		)
		iterV, _ := container.PeekBack()
		test.BasicTest(i, iterV,
			"Push front did not put the value in the correct place.", t,
		)
		iterV, _ = container.PeekFront()
		test.BasicTest(0, iterV,
			"Push front did not put the value in the correct place.", t,
		)
	}
}

// Tests the ForcePushBack method functionality of a dynamic Deque.
func DequeInterfaceForcePushBack(
	factory func() dynamicContainers.Deque[int, int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 4; i++ {
		container.ForcePushBack(i)
		test.BasicTest(i+1, container.Length(),
			"Push back did not add the value correctly.", t,
		)
		iterV, _ := container.PeekBack()
		test.BasicTest(i, iterV,
			"Push front did not put the value in the correct place.", t,
		)
		iterV, _ = container.PeekFront()
		test.BasicTest(0, iterV,
			"Push front did not put the value in the correct place.", t,
		)
	}
}

// func DequeInterfaceEq(t *testing.T){
//     v:=Deque[int,widgets.BuiltinInt]([]int{0,1,2,3})
//     v2:=Deque[int,widgets.BuiltinInt]([]int{0,1,2,3})
//     comp:=func(l *int, r *int) bool { return *l==*r }
//     test.BasicTest(true,v.Eq(&v2,comp),
// 	"Eq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Eq(&v,comp),
// 	"Eq returned a false negative.",t,
//     )
//     v.Delete(3)
//     test.BasicTest(false,v.Eq(&v2,comp),
// 	"Eq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Eq(&v,comp),
// 	"Eq returned a false positive.",t,
//     )
//     v=Deque[int,widgets.BuiltinInt]([]int{0})
//     v2=Deque[int,widgets.BuiltinInt]([]int{0})
//     test.BasicTest(true,v.Eq(&v2,comp),
// 	"Eq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Eq(&v,comp),
// 	"Eq returned a false negative.",t,
//     )
//     v.Delete(0)
//     test.BasicTest(false,v.Eq(&v2,comp),
// 	"Eq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Eq(&v,comp),
// 	"Eq returned a false positive.",t,
//     )
//     v=Deque[int,widgets.BuiltinInt]([]int{})
//     v2=Deque[int,widgets.BuiltinInt]([]int{})
//     test.BasicTest(true,v.Eq(&v2,comp),
// 	"Eq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Eq(&v,comp),
// 	"Eq returned a false negative.",t,
//     )
// }
//
// func DequeInterfaceNeq(t *testing.T){
//     v:=Deque[int,widgets.BuiltinInt]([]int{0,1,2,3})
//     v2:=Deque[int,widgets.BuiltinInt]([]int{0,1,2,3})
//     comp:=func(l *int, r *int) bool { return *l==*r }
//     test.BasicTest(false,v.Neq(&v2,comp),
// 	"Neq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Neq(&v,comp),
// 	"Neq returned a false positive.",t,
//     )
//     v.Delete(3)
//     test.BasicTest(true,v.Neq(&v2,comp),
// 	"Neq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Neq(&v,comp),
// 	"Neq returned a false negative.",t,
//     )
//     v=Deque[int,widgets.BuiltinInt]([]int{0})
//     v2=Deque[int,widgets.BuiltinInt]([]int{0})
//     test.BasicTest(false,v.Neq(&v2,comp),
// 	"Neq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Neq(&v,comp),
// 	"Neq returned a false positive.",t,
//     )
//     v.Delete(0)
//     test.BasicTest(true,v.Neq(&v2,comp),
// 	"Neq returned a false negative.",t,
//     )
//     test.BasicTest(true,v2.Neq(&v,comp),
// 	"Neq returned a false negative.",t,
//     )
//     v=Deque[int,widgets.BuiltinInt]([]int{})
//     v2=Deque[int,widgets.BuiltinInt]([]int{})
//     test.BasicTest(false,v.Neq(&v2,comp),
// 	"Neq returned a false positive.",t,
//     )
//     test.BasicTest(false,v2.Neq(&v,comp),
// 	"Neq returned a false positive.",t,
//     )
// }
