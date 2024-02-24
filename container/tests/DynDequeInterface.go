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
func DequeInterfaceLengthInterface[V any](
	factory func() dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.Length = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Capacity] interface.
func DequeInterfaceCapacityInterface[V any](
	factory func() dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.Capacity = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Clear] interface.
func DequeInterfaceClearInterface[V any](
	factory func() dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.FirstElemRead] interface.
func DequeInterfaceFirstElemReadInterface[V any](
	factory func() dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.FirstElemRead[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.FirstElemWrite] interface.
func DequeInterfaceFirstElemWriteInterface[V any](
	factory func() dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.FirstElemWrite[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.FirstElemDelete] interface.
func DequeInterfaceFirstElemDeleteInterface[V any](
	factory func() dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.FirstElemDelete[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.LastElemRead] interface.
func DequeInterfaceLastElemReadInterface[V any](
	factory func() dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.LastElemRead[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.LastElemWrite] interface.
func DequeInterfaceLastElemWriteInterface[V any](
	factory func() dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.LastElemWrite[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.LastElemDelete] interface.
func DequeInterfaceLastElemDeleteInterface[V any](
	factory func() dynamicContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.LastElemDelete[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.DequeRead] interface.
func ReadDequeInterface[V any](
	factory func() dynamicContainers.Deque[V],
	t *testing.T,
) {
	dequeReadInterface[V](factory())
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.WriteDeque] interface.
func WriteDequeInterface[V any](
	factory func() dynamicContainers.Deque[V],
	t *testing.T,
) {
	dequeWriteInterface[V](factory())
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.Deque] interface.
func DequeInterfaceInterface[V any](
	factory func() dynamicContainers.Deque[V],
	t *testing.T,
) {
	dequeInterface[V](factory())
}

// Tests that the value supplied by the factory does not implement the 
// [staticContainers.Deque] interface.
func DequeInterfaceStaticCapacityInterface[V any](
	factory func() dynamicContainers.Deque[V],
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

// Tests the Clear method functionality of a dynamic Deque.
func DequeInterfaceClear(
	factory func() dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 6; i++ {
		container.PushBack(i)
	}
	container.Clear()
	test.Eq(0, container.Length(), t)
	test.Eq(0, container.Capacity(), t)
}

// Tests the PeekPntrFront method functionality of a dynamic Deque.
func DequeInterfacePeekPntrFront(
	factory func() dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory()
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
func DequeInterfacePeekFront(
	factory func() dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory()
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
func DequeInterfacePeekPntrBack(
	factory func() dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory()
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
func DequeInterfacePeekBack(
	factory func() dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory()
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
func DequeInterfacePopFront(
	factory func() dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory()
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
func DequeInterfacePopBack(
	factory func() dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory()
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
func DequeInterfacePushFront(
	factory func() dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 4; i++ {
		container.PushFront(i)
		test.Eq(i+1, container.Length(),t)
		iterV, _ := container.PeekFront()
		test.Eq(i, iterV,t)
		iterV, _ = container.PeekBack()
		test.Eq(0, iterV,t)
	}
	container=factory()
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
func DequeInterfaceForcePushFront(
	factory func() dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 4; i++ {
		container.ForcePushFront(i)
		test.Eq(i+1, container.Length(),t)
		iterV, _ := container.PeekFront()
		test.Eq(i, iterV,t)
		iterV, _ = container.PeekBack()
		test.Eq(0, iterV,t)
	}
	container=factory()
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
func DequeInterfacePushBack(
	factory func() dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 4; i++ {
		container.PushBack(i)
		test.Eq(i+1, container.Length(),t)
		iterV, _ := container.PeekBack()
		test.Eq(i, iterV,t)
		iterV, _ = container.PeekFront()
		test.Eq(0, iterV,t)
	}
	container=factory()
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
func DequeInterfaceForcePushBack(
	factory func() dynamicContainers.Deque[int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 4; i++ {
		container.ForcePushBack(i)
		test.Eq(i+1, container.Length(),t)
		iterV, _ := container.PeekBack()
		test.Eq(i, iterV,t)
		iterV, _ = container.PeekFront()
		test.Eq(0, iterV,t)
	}
	container=factory()
	for i := 0; i < 6; i+=2 {
		container.ForcePushBack(i,i+1)
		test.Eq(i+2, container.Length(),t)
		iterV, _ := container.PeekBack()
		test.Eq(i+1, iterV,t)
		iterV, _ = container.PeekFront()
		test.Eq(0, iterV,t)
	}
}
