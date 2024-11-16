package tests

import (
	"testing"

	"github.com/barbell-math/util/src/container/containerTypes"
	"github.com/barbell-math/util/src/container/staticContainers"
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/test"
)

func staticQueueReadInterface[U any](c staticContainers.ReadQueue[U])   {}
func staticQueueWriteInterface[U any](c staticContainers.WriteQueue[U]) {}
func staticQueueInterface[U any](c staticContainers.Queue[U])           {}

// Tests that the value supplied by the factory implements the
// [containerTypes.StaticCapacity] interface.
func StaticQueueInterfaceStaticCapacity[V any](
	factory func(capacity int) staticContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.StaticCapacity = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Length] interface.
func StaticQueueInterfaceLengthInterface[V any](
	factory func(capacity int) staticContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.Length = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Capacity] interface.
func StaticQueueInterfaceCapacityInterface[V any](
	factory func(capacity int) staticContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.Capacity = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Clear] interface.
func StaticQueueInterfaceClearInterface[V any](
	factory func(capacity int) staticContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.FirstElemRead] interface.
func StaticQueueInterfaceFirstElemReadInterface[V any](
	factory func(capacity int) staticContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.FirstElemRead[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.FirstElemDelete] interface.
func StaticQueueInterfaceFirstElemDeleteInterface[V any](
	factory func(capacity int) staticContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.FirstElemDelete[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.LastElemWrite] interface.
func StaticQueueInterfaceLastElemWriteInterface[V any](
	factory func(capacity int) staticContainers.Queue[V],
	t *testing.T,
) {
	var container containerTypes.LastElemWrite[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [staticContainers.QueueRead] interface.
func ReadStaticQueueInterface[V any](
	factory func(capacity int) staticContainers.Queue[V],
	t *testing.T,
) {
	staticQueueReadInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the
// [staticContainers.WriteQueue] interface.
func WriteStaticQueueInterface[V any](
	factory func(capacity int) staticContainers.Queue[V],
	t *testing.T,
) {
	staticQueueWriteInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the
// [staticContainers.Queue] interface.
func StaticQueueInterfaceInterface[V any](
	factory func(capacity int) staticContainers.Queue[V],
	t *testing.T,
) {
	staticQueueInterface[V](factory(0))
}

// Tests the Clear method functionality of a static Queue.
func StaticQueueInterfaceClear(
	factory func(capacity int) staticContainers.Queue[int],
	t *testing.T,
) {
	container := factory(6)
	for i := 0; i < 6; i++ {
		container.PushBack(i)
	}
	container.Clear()
	test.Eq(0, container.Length(), t)
	test.Eq(6, container.Capacity(), t)
}

// Tests the PeekPntrFront method functionality of a static Queue.
func StaticQueueInterfacePeekPntrFront(
	factory func(capacity int) staticContainers.Queue[int],
	t *testing.T,
) {
	container := factory(5)
	_v, err := container.PeekPntrFront()
	test.NilPntr[int](_v, t)
	test.ContainsError(customerr.ValOutsideRange, err, t)
	test.Nil(container.PushBack(1), t)
	_v, err = container.PeekPntrFront()
	test.Eq(1, *_v, t)
	test.Nil(err, t)
	test.Nil(container.PushBack(2), t)
	_v, err = container.PeekPntrFront()
	test.Eq(1, *_v, t)
	test.Nil(err, t)
	test.Nil(container.PushBack(3, 4, 5), t)
	_v, err = container.PeekPntrFront()
	test.Eq(1, *_v, t)
	test.Nil(err, t)
}

// Tests the PeekFront method functionality of a static Queue.
func StaticQueueInterfacePeekFront(
	factory func(capacity int) staticContainers.Queue[int],
	t *testing.T,
) {
	container := factory(5)
	_, err := container.PeekFront()
	test.ContainsError(customerr.ValOutsideRange, err, t)
	test.Nil(container.PushBack(1), t)
	_v, err := container.PeekFront()
	test.Eq(1, _v, t)
	test.Nil(err, t)
	test.Nil(container.PushBack(2), t)
	_v, err = container.PeekFront()
	test.Eq(1, _v, t)
	test.Nil(err, t)
	test.Nil(container.PushBack(3, 4, 5), t)
	_v, err = container.PeekFront()
	test.Eq(1, _v, t)
	test.Nil(err, t)
}

func staticQueuePopFrontHelper(
	factory func(capacity int) staticContainers.Queue[int],
	l int,
	c int,
	t *testing.T,
) {
	container := factory(c)
	for i := 0; i < l; i++ {
		test.Nil(container.PushBack(i), t)
	}
	test.Eq(l, container.Length(), t)
	test.Eq(c, container.Capacity(), t)
	for i := 0; i < l; i++ {
		f, err := container.PopFront()
		test.Eq(c, container.Capacity(), t)
		test.Eq(l-i-1, container.Length(), t)
		test.Eq(i, f, t)
		test.Nil(err, t)
	}
	test.Eq(0, container.Length(), t)
	test.Eq(c, container.Capacity(), t)
	_, err := container.PopFront()
	test.ContainsError(containerTypes.Empty, err, t)
}

// Tests the PopFront method functionality of a static Queue.
func StaticQueueInterfacePopFront(
	factory func(capacity int) staticContainers.Queue[int],
	t *testing.T,
) {
	staticQueuePopFrontHelper(factory, 5, 5, t)
	staticQueuePopFrontHelper(factory, 5, 10, t)
}

func staticQueuePushBackHelper(
	factory func(capacity int) staticContainers.Queue[int],
	l int,
	c int,
	t *testing.T,
) {
	container := factory(c)
	for i := 0; i < l; i++ {
		test.Nil(container.PushBack(i), t)
		test.Eq(i+1, container.Length(), t)
		test.Eq(c, container.Capacity(), t)
		iterV, err := container.PeekFront()
		test.Nil(err, t)
		test.Eq(0, iterV, t)
	}
	test.Eq(l, container.Length(), t)
	test.Eq(c, container.Capacity(), t)
}

// Tests the PushBack method functionality of a static Queue.
func StaticQueueInterfacePushBack(
	factory func(capacity int) staticContainers.Queue[int],
	t *testing.T,
) {
	staticQueuePushBackHelper(factory, 5, 5, t)
	staticQueuePushBackHelper(factory, 5, 10, t)
	container := factory(6)
	for i := 0; i < 6; i += 2 {
		test.Nil(container.PushBack(i, i+1), t)
		test.Eq(i+2, container.Length(), t)
		iterV, err := container.PeekFront()
		test.Nil(err, t)
		test.Eq(0, iterV, t)
	}
	test.Eq(6, container.Length(), t)
	test.Eq(6, container.Capacity(), t)
}

func staticQueueForcePushBackHelper(
	factory func(capacity int) staticContainers.Queue[int],
	l int,
	c int,
	t *testing.T,
) {
	container := factory(c)
	for i := 0; i < l; i++ {
		container.ForcePushBack(i)
		test.Eq(i+1, container.Length(), t)
		test.Eq(c, container.Capacity(), t)
		iterV, err := container.PeekFront()
		test.Nil(err, t)
		test.Eq(0, iterV, t)
	}
	for i := l; i < c; i++ {
		container.ForcePushBack(i)
		if c <= l {
			test.Eq(l, container.Length(), t)
		} else {
			test.Eq(i+1, container.Length(), t)
		}
		test.Eq(c, container.Capacity(), t)
		iterV, err := container.PeekFront()
		test.Nil(err, t)
		if i >= c {
			test.Eq(i-l+1, iterV, t)
		} else {
			test.Eq(0, iterV, t)
		}
	}
}

// Tests the ForcePopFront method functionality of a static Queue.
func StaticQueueInterfaceForcePushBack(
	factory func(capacity int) staticContainers.Queue[int],
	t *testing.T,
) {
	staticQueueForcePushBackHelper(factory, 5, 5, t)
	staticQueueForcePushBackHelper(factory, 5, 10, t)

	container := factory(5)
	container.ForcePushBack(1)
	v, err := container.PeekFront()
	test.Nil(err, t)
	test.Eq(1, v, t)

	container = factory(5)
	container.ForcePushBack(1, 2, 3)
	v, err = container.PeekFront()
	test.Nil(err, t)
	test.Eq(1, v, t)

	container = factory(5)
	container.ForcePushBack(1, 2, 3, 4, 5, 6, 7)
	v, err = container.PeekFront()
	test.Nil(err, t)
	test.Eq(3, v, t)
}
