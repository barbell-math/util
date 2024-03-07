package tests

import (
	"testing"

	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func queueReadInterface[U any](c dynamicContainers.ReadQueue[U])   {}
func queueWriteInterface[U any](c dynamicContainers.WriteQueue[U]) {}
func queueInterface[U any](c dynamicContainers.Queue[U])           {}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Length] interface.
func DynQueueInterfaceLengthInterface[V any](
	factory func(capacity int) dynamicContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.Length = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Capacity] interface.
func DynQueueInterfaceCapacityInterface[V any](
	factory func(capacity int) dynamicContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.Capacity = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Clear] interface.
func DynQueueInterfaceClearInterface[V any](
	factory func(capacity int) dynamicContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.FirstElemRead] interface.
func DynQueueInterfaceFirstElemReadInterface[V any](
	factory func(capacity int) dynamicContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.FirstElemRead[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.FirstElemDelete] interface.
func DynQueueInterfaceFirstElemDeleteInterface[V any](
	factory func(capacity int) dynamicContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.FirstElemDelete[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.LastElemWrite] interface.
func DynQueueInterfaceLastElemWriteInterface[V any](
	factory func(capacity int) dynamicContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.LastElemWrite[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.QueueRead] interface.
func ReadDynQueueInterface[V any](
	factory func(capacity int) dynamicContainers.Queue[V],
	t *testing.T,
) {
	queueReadInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.WriteQueue] interface.
func WriteDynQueueInterface[V any](
	factory func(capacity int) dynamicContainers.Queue[V],
	t *testing.T,
) {
	queueWriteInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.Queue] interface.
func DynQueueInterfaceInterface[V any](
	factory func(capacity int) dynamicContainers.Queue[V],
	t *testing.T,
) {
	queueInterface[V](factory(0))
}

// Tests that the value supplied by the factory does not implement the 
// [staticContainers.Queue] interface.
func DynQueueInterfaceStaticCapacityInterface[V any](
	factory func(capacity int) dynamicContainers.Queue[V],
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

// Tests the Clear method functionality of a dynamic Queue.
func DynQueueInterfaceClear(
	factory func(capacity int) dynamicContainers.Queue[int],
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

// Tests the PeekPntrFront method functionality of a dynamic Queue.
func DynQueueInterfacePeekPntrFront(
	factory func(capacity int) dynamicContainers.Queue[int],
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

// Tests the PeekFront method functionality of a dynamic Queue.
func DynQueueInterfacePeekFront(
	factory func(capacity int) dynamicContainers.Queue[int],
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

// Tests the PopFront method functionality of a dynamic Queue.
func DynQueueInterfacePopFront(
	factory func(capacity int) dynamicContainers.Queue[int],
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

// Tests the PushBack method functionality of a dynamic Queue.
func DynQueueInterfacePushBack(
	factory func(capacity int) dynamicContainers.Queue[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 4; i++ {
		container.PushBack(i)
		test.Eq(i+1, container.Length(),t)
		iterV, _ := container.PeekFront()
		test.Eq(0, iterV,t)
	}
	container=factory(0)
	for i := 0; i < 6; i+=2 {
		container.PushBack(i,i+1)
		test.Eq(i+2, container.Length(),t)
		iterV, _ := container.PeekFront()
		test.Eq(0, iterV,t)
	}
}

// Tests the ForcePushBack method functionality of a dynamic Queue.
func DynQueueInterfaceForcePushBack(
	factory func(capacity int) dynamicContainers.Queue[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 4; i++ {
		container.ForcePushBack(i)
		test.Eq(i+1, container.Length(),t)
		iterV, _ := container.PeekFront()
		test.Eq(0, iterV,t)
	}
	container=factory(0)
	for i := 0; i < 6; i+=2 {
		container.ForcePushBack(i,i+1)
		test.Eq(i+2, container.Length(),t)
		iterV, _ := container.PeekFront()
		test.Eq(0, iterV,t)
	}
}
