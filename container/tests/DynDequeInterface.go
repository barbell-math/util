package tests

import (
	"testing"

	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func dequeReadInterface[U any](c dynamicContainers.ReadDeque[U])   {}
func dequeWriteInterface[U any](c dynamicContainers.WriteDeque[U]) {}
func dequeInterface[U any](c dynamicContainers.Deque[U])           {}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Length] interface.
func DynDequeInterfaceLengthInterface[V any](
	factory func(capacity int) dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.Length = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Capacity] interface.
func DynDequeInterfaceCapacityInterface[V any](
	factory func(capacity int) dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.Capacity = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Clear] interface.
func DynDequeInterfaceClearInterface[V any](
	factory func(capacity int) dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.FirstElemRead] interface.
func DynDequeInterfaceFirstElemReadInterface[V any](
	factory func(capacity int) dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.FirstElemRead[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.FirstElemWrite] interface.
func DynDequeInterfaceFirstElemWriteInterface[V any](
	factory func(capacity int) dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.FirstElemWrite[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.FirstElemDelete] interface.
func DynDequeInterfaceFirstElemDeleteInterface[V any](
	factory func(capacity int) dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.FirstElemDelete[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.LastElemRead] interface.
func DynDequeInterfaceLastElemReadInterface[V any](
	factory func(capacity int) dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.LastElemRead[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.LastElemWrite] interface.
func DynDequeInterfaceLastElemWriteInterface[V any](
	factory func(capacity int) dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.LastElemWrite[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.LastElemDelete] interface.
func DynDequeInterfaceLastElemDeleteInterface[V any](
	factory func(capacity int) dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.LastElemDelete[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.DequeRead] interface.
func ReadDynDequeInterface[V any](
	factory func(capacity int) dynamicContainers.Deque[V],
	t *testing.T,
) {
	dequeReadInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.WriteDeque] interface.
func WriteDynDequeInterface[V any](
	factory func(capacity int) dynamicContainers.Deque[V],
	t *testing.T,
) {
	dequeWriteInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.Deque] interface.
func DynDequeInterfaceInterface[V any](
	factory func(capacity int) dynamicContainers.Deque[V],
	t *testing.T,
) {
	dequeInterface[V](factory(0))
}

// Tests that the value supplied by the factory does not implement the 
// [staticContainers.Deque] interface.
func DynDequeInterfaceStaticCapacityInterface[V any](
	factory func(capacity int) dynamicContainers.Deque[V],
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

// Tests the Clear method functionality of a dynamic Deque.
func DynDequeInterfaceClear(
	factory func(capacity int) dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 6; i++ {
		container.PushBack(i)
	}
	container.Clear()
	test.Eq(0, container.Length(), t)
	test.Eq(0, container.Capacity(), t)
}

// Tests the PeekPntrFront method functionality of a dynamic Deque.
func DynDequeInterfacePeekPntrFront(
	factory func(capacity int) dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory(0)
	_v, err := container.PeekPntrFront()
	test.NilPntr[int](_v,t)
	test.ContainsError(customerr.ValOutsideRange, err,t)
	container.PushBack(1)
	_v, err = container.PeekPntrFront()
	test.Eq(1, *_v,t)
	test.Nil(err,t)
	container.PushBack(2)
	_v, err = container.PeekPntrFront()
	test.Eq(1, *_v,t)
	test.Nil(err,t)
}

// Tests the PeekFront method functionality of a dynamic Deque.
func DynDequeInterfacePeekFront(
	factory func(capacity int) dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory(0)
	_, err := container.PeekFront()
	test.ContainsError(customerr.ValOutsideRange, err,t)
	container.PushBack(1)
	_v, err := container.PeekFront()
	test.Eq(1, _v,t)
	test.Nil(err,t)
	container.PushBack(2)
	_v, err = container.PeekFront()
	test.Eq(1, _v,t)
	test.Nil(err,t)
}

// Tests the PeekPntrBack method functionality of a dynamic Deque.
func DynDequeInterfacePeekPntrBack(
	factory func(capacity int) dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory(0)
	_v, err := container.PeekPntrBack()
	test.NilPntr[int](_v,t)
	test.ContainsError(customerr.ValOutsideRange, err,t)
	container.PushBack(1)
	_v, err = container.PeekPntrBack()
	test.Eq(1, *_v,t)
	test.Nil(err,t)
	container.PushBack(2)
	_v, err = container.PeekPntrBack()
	test.Eq(2, *_v,t)
	test.Nil(err,t)
}

// Tests the PeekBack method functionality of a dynamic Deque.
func DynDequeInterfacePeekBack(
	factory func(capacity int) dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory(0)
	_, err := container.PeekBack()
	test.ContainsError(customerr.ValOutsideRange, err,t)
	container.PushBack(1)
	_v, err := container.PeekBack()
	test.Eq(1, _v,t)
	test.Nil(err,t)
	container.PushBack(2)
	_v, err = container.PeekBack()
	test.Eq(2, _v,t)
	test.Nil(err,t)
}

// Tests the PopFront method functionality of a dynamic Deque.
func DynDequeInterfacePopFront(
	factory func(capacity int) dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 4; i++ {
		container.PushBack(i)
	}
	for i := 0; i < 4; i++ {
		f, err := container.PopFront()
		test.Eq(i, f,t)
		test.Nil(err,t)
	}
	_, err := container.PopFront()
	test.ContainsError(containerTypes.Empty, err,t)
}

// Tests the PopBack method functionality of a dynamic Deque.
func DynDequeInterfacePopBack(
	factory func(capacity int) dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 4; i++ {
		container.PushBack(i)
	}
	for i := 3; i >= 0; i-- {
		f, err := container.PopBack()
		test.Eq(i, f,t)
		test.Nil(err,t)
	}
	_, err := container.PopBack()
	test.ContainsError(containerTypes.Empty, err,t)
}

// Tests the PopFront method functionality of a dynamic Deque.
func DynDequeInterfacePushFront(
	factory func(capacity int) dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 4; i++ {
		container.PushFront(i)
		test.Eq(i+1, container.Length(),t)
		iterV, _ := container.PeekFront()
		test.Eq(i, iterV,t)
		iterV, _ = container.PeekBack()
		test.Eq(0, iterV,t)
	}
	container=factory(0)
	for i := 0; i < 6; i+=2 {
		container.PushFront(i,i+1)
		test.Eq(i+2, container.Length(),t)
		iterV, _ := container.PeekFront()
		test.Eq(i, iterV,t)
		iterV, _ = container.PeekBack()
		test.Eq(1, iterV,t)
	}
}

// Tests the ForcePopFront method functionality of a dynamic Deque.
func DynDequeInterfaceForcePushFront(
	factory func(capacity int) dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 4; i++ {
		container.ForcePushFront(i)
		test.Eq(i+1, container.Length(),t)
		iterV, _ := container.PeekFront()
		test.Eq(i, iterV,t)
		iterV, _ = container.PeekBack()
		test.Eq(0, iterV,t)
	}
	container=factory(0)
	for i := 0; i < 6; i+=2 {
		container.PushFront(i,i+1)
		test.Eq(i+2, container.Length(),t)
		iterV, _ := container.PeekFront()
		test.Eq(i, iterV,t)
		iterV, _ = container.PeekBack()
		test.Eq(1, iterV,t)
	}
}

// Tests the PushBack method functionality of a dynamic Deque.
func DynDequeInterfacePushBack(
	factory func(capacity int) dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 4; i++ {
		container.PushBack(i)
		test.Eq(i+1, container.Length(),t)
		iterV, _ := container.PeekBack()
		test.Eq(i, iterV,t)
		iterV, _ = container.PeekFront()
		test.Eq(0, iterV,t)
	}
	container=factory(0)
	for i := 0; i < 6; i+=2 {
		container.PushBack(i,i+1)
		test.Eq(i+2, container.Length(),t)
		iterV, _ := container.PeekBack()
		test.Eq(i+1, iterV,t)
		iterV, _ = container.PeekFront()
		test.Eq(0, iterV,t)
	}
}

// Tests the ForcePushBack method functionality of a dynamic Deque.
func DynDequeInterfaceForcePushBack(
	factory func(capacity int) dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 4; i++ {
		container.ForcePushBack(i)
		test.Eq(i+1, container.Length(),t)
		iterV, _ := container.PeekBack()
		test.Eq(i, iterV,t)
		iterV, _ = container.PeekFront()
		test.Eq(0, iterV,t)
	}
	container=factory(0)
	for i := 0; i < 6; i+=2 {
		container.ForcePushBack(i,i+1)
		test.Eq(i+2, container.Length(),t)
		iterV, _ := container.PeekBack()
		test.Eq(i+1, iterV,t)
		iterV, _ = container.PeekFront()
		test.Eq(0, iterV,t)
	}
}
