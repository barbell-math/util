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
func QueueInterfaceLengthInterface[V any](
	factory func() dynamicContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.Length = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Capacity] interface.
func QueueInterfaceCapacityInterface[V any](
	factory func() dynamicContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.Capacity = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Clear] interface.
func QueueInterfaceClearInterface[V any](
	factory func() dynamicContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.FirstElemRead] interface.
func QueueInterfaceFirstElemReadInterface[V any](
	factory func() dynamicContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.FirstElemRead[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.FirstElemDelete] interface.
func QueueInterfaceFirstElemDeleteInterface[V any](
	factory func() dynamicContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.FirstElemDelete[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.LastElemWrite] interface.
func QueueInterfaceLastElemWriteInterface[V any](
	factory func() dynamicContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.LastElemWrite[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.QueueRead] interface.
func ReadQueueInterface[V any](
	factory func() dynamicContainers.Queue[V],
	t *testing.T,
) {
	queueReadInterface[V](factory())
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.WriteQueue] interface.
func WriteQueueInterface[V any](
	factory func() dynamicContainers.Queue[V],
	t *testing.T,
) {
	queueWriteInterface[V](factory())
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.Queue] interface.
func QueueInterfaceInterface[V any](
	factory func() dynamicContainers.Queue[V],
	t *testing.T,
) {
	queueInterface[V](factory())
}

// Tests that the value supplied by the factory does not implement the 
// [staticContainers.Queue] interface.
func QueueInterfaceStaticCapacityInterface[V any](
	factory func() dynamicContainers.Queue[V],
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
	factory func() dynamicContainers.Queue[int],
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
	factory func() dynamicContainers.Queue[int],
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
	factory func() dynamicContainers.Queue[int],
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
	factory func() dynamicContainers.Queue[int],
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

// Tests the PushBack method functionality of a dynamic Queue.
func QueueInterfacePushBack(
	factory func() dynamicContainers.Queue[int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 4; i++ {
		container.PushBack(i)
		test.BasicTest(i+1, container.Length(),
			"Push back did not add the value correctly.", t,
		)
		iterV, _ := container.PeekFront()
		test.BasicTest(0, iterV,
			"Push front did not put the value in the correct place.", t,
		)
	}
	container=factory()
	for i := 0; i < 6; i+=2 {
		container.PushBack(i,i+1)
		test.BasicTest(i+2, container.Length(),
			"Push front did not add the value correctly.", t,
		)
		iterV, _ := container.PeekFront()
		test.BasicTest(0, iterV,
			"Push front did not put the value in the correct place.", t,
		)
	}
}

// Tests the ForcePushBack method functionality of a dynamic Queue.
func QueueInterfaceForcePushBack(
	factory func() dynamicContainers.Queue[int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 4; i++ {
		container.ForcePushBack(i)
		test.BasicTest(i+1, container.Length(),
			"Push back did not add the value correctly.", t,
		)
		iterV, _ := container.PeekFront()
		test.BasicTest(0, iterV,
			"Push front did not put the value in the correct place.", t,
		)
	}
	container=factory()
	for i := 0; i < 6; i+=2 {
		container.ForcePushBack(i,i+1)
		test.BasicTest(i+2, container.Length(),
			"Push front did not add the value correctly.", t,
		)
		iterV, _ := container.PeekFront()
		test.BasicTest(0, iterV,
			"Push front did not put the value in the correct place.", t,
		)
	}
}
