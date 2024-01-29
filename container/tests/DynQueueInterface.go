package tests

import (
	"testing"

	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func queueReadInterface[T ~int, U any](c dynamicContainers.ReadQueue[T, U])   {}
func queueWriteInterface[T ~int, U any](c dynamicContainers.WriteQueue[T, U]) {}
func queueInterface[T ~int, U any](c dynamicContainers.Queue[T, U])           {}

// Tests that the value supplied by the factory implements the 
// [containerTypes.RWSyncable] interface.
func QueueInterfaceSyncableInterface[K ~int, V any](
	factory func() dynamicContainers.Queue[K, V],
	t *testing.T,
) {
	var container containerTypes.RWSyncable = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Length] interface.
func QueueInterfaceLengthInterface[K ~int, V any](
	factory func() dynamicContainers.Queue[K, V],
	t *testing.T,
) {
	var container containerTypes.Length = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Capacity] interface.
func QueueInterfaceCapacityInterface[K ~int, V any](
	factory func() dynamicContainers.Queue[K, V],
	t *testing.T,
) {
	var container containerTypes.Capacity = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Clear] interface.
func QueueInterfaceClearInterface[K ~int, V any](
	factory func() dynamicContainers.Queue[K, V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.FirstElemRead] interface.
func QueueInterfaceFirstElemReadInterface[K ~int, V any](
	factory func() dynamicContainers.Queue[K, V],
	t *testing.T,
) {
	var container containerTypes.FirstElemRead[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.FirstElemDelete] interface.
func QueueInterfaceFirstElemDeleteInterface[K ~int, V any](
	factory func() dynamicContainers.Queue[K, V],
	t *testing.T,
) {
	var container containerTypes.FirstElemDelete[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.LastElemWrite] interface.
func QueueInterfaceLastElemWriteInterface[K ~int, V any](
	factory func() dynamicContainers.Queue[K, V],
	t *testing.T,
) {
	var container containerTypes.LastElemWrite[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.QueueRead] interface.
func ReadQueueInterface[K ~int, V any](
	factory func() dynamicContainers.Queue[K, V],
	t *testing.T,
) {
	queueReadInterface[K, V](factory())
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.WriteQueue] interface.
func WriteQueueInterface[K ~int, V any](
	factory func() dynamicContainers.Queue[K, V],
	t *testing.T,
) {
	queueWriteInterface[K, V](factory())
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.Queue] interface.
func QueueInterfaceInterface[K ~int, V any](
	factory func() dynamicContainers.Queue[K, V],
	t *testing.T,
) {
	queueInterface[K, V](factory())
}

// Tests that the value supplied by the factory does not implement the 
// [staticContainers.Queue] interface.
func QueueInterfaceStaticCapacityInterface[K ~int, V any](
	factory func() dynamicContainers.Queue[K, V],
	t *testing.T,
) {
	test.Panics(
		func() {
			var c any
			c = factory()
			c2 := c.(containerTypes.StaticCapacity)
			_ = c2
		},
		"Code did not panic when casting a dynamic Queue to a static vector.", t,
	)
}

// Tests the Clear method functionality of a dynamic Queue.
func QueueInterfaceClear(
	factory func() dynamicContainers.Queue[int, int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 6; i++ {
		container.PushBack(i)
	}
	container.Clear()
	test.BasicTest(0, container.Length(), "Clear did not reset the underlying Queue.", t)
	test.BasicTest(0, container.Capacity(), "Clear did not reset the underlying Queue.", t)
}

// Tests the PeekPntrFront method functionality of a dynamic Queue.
func QueueInterfacePeekPntrFront(
	factory func() dynamicContainers.Queue[int, int],
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

// Tests the PeekFront method functionality of a dynamic Queue.
func QueueInterfacePeekFront(
	factory func() dynamicContainers.Queue[int, int],
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

// Tests the PopFront method functionality of a dynamic Queue.
func QueueInterfacePopFront(
	factory func() dynamicContainers.Queue[int, int],
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

// Tests the PushBack method functionality of a dynamic Deque.
func QueueInterfacePushBack(
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
func QueueInterfaceForcePushBack(
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
