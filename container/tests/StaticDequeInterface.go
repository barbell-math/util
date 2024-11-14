package tests

import (
	"testing"

	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/staticContainers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func staticDequeReadInterface[U any](c staticContainers.ReadDeque[U])   {}
func staticDequeWriteInterface[U any](c staticContainers.WriteDeque[U]) {}
func staticDequeInterface[U any](c staticContainers.Deque[U])           {}

// Tests that the value supplied by the factory implements the
// [containerTypes.StaticCapacity] interface.
func StaticDequeInterfaceStaticCapacity[V any](
	factory func(capacity int) staticContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.StaticCapacity = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Length] interface.
func StaticDequeInterfaceLengthInterface[V any](
	factory func(capacity int) staticContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.Length = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Capacity] interface.
func StaticDequeInterfaceCapacityInterface[V any](
	factory func(capacity int) staticContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.Capacity = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Clear] interface.
func StaticDequeInterfaceClearInterface[V any](
	factory func(capacity int) staticContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.FirstElemRead] interface.
func StaticDequeInterfaceFirstElemReadInterface[V any](
	factory func(capacity int) staticContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.FirstElemRead[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.FirstElemWrite] interface.
func StaticDequeInterfaceFirstElemWriteInterface[V any](
	factory func(capacity int) staticContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.FirstElemWrite[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.FirstElemDelete] interface.
func StaticDequeInterfaceFirstElemDeleteInterface[V any](
	factory func(capacity int) staticContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.FirstElemDelete[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.LastElemRead] interface.
func StaticDequeInterfaceLastElemReadInterface[V any](
	factory func(capacity int) staticContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.LastElemRead[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.LastElemWrite] interface.
func StaticDequeInterfaceLastElemWriteInterface[V any](
	factory func(capacity int) staticContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.LastElemWrite[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.LastElemDelete] interface.
func StaticDequeInterfaceLastElemDeleteInterface[V any](
	factory func(capacity int) staticContainers.Deque[V],
	t *testing.T,
) {
	var container containerTypes.LastElemDelete[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [staticContainers.DequeRead] interface.
func ReadStaticDequeInterface[V any](
	factory func(capacity int) staticContainers.Deque[V],
	t *testing.T,
) {
	staticDequeReadInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the
// [staticContainers.WriteDeque] interface.
func WriteStaticDequeInterface[V any](
	factory func(capacity int) staticContainers.Deque[V],
	t *testing.T,
) {
	staticDequeWriteInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the
// [staticContainers.Deque] interface.
func StaticDequeInterfaceInterface[V any](
	factory func(capacity int) staticContainers.Deque[V],
	t *testing.T,
) {
	staticDequeInterface[V](factory(0))
}

// Tests the Clear method functionality of a static Deque.
func StaticDequeInterfaceClear(
	factory func(capacity int) staticContainers.Deque[int],
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

// Tests the PeekPntrFront method functionality of a static Deque.
func StaticDequeInterfacePeekPntrFront(
	factory func(capacity int) staticContainers.Deque[int],
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

// Tests the PeekFront method functionality of a static Deque.
func StaticDequeInterfacePeekFront(
	factory func(capacity int) staticContainers.Deque[int],
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

// Tests the PeekPntrBack method functionality of a static Deque.
func StaticDequeInterfacePeekPntrBack(
	factory func(capacity int) staticContainers.Deque[int],
	t *testing.T,
) {
	container := factory(5)
	_v, err := container.PeekPntrBack()
	test.NilPntr[int](_v, t)
	test.ContainsError(customerr.ValOutsideRange, err, t)
	container.PushBack(1)
	_v, err = container.PeekPntrBack()
	test.Eq(1, *_v, t)
	test.Nil(err, t)
	container.PushBack(2)
	_v, err = container.PeekPntrBack()
	test.Eq(2, *_v, t)
	test.Nil(err, t)
	container.PushBack(3, 4, 5)
	_v, err = container.PeekPntrBack()
	test.Eq(5, *_v, t)
	test.Nil(err, t)
}

// Tests the PeekBack method functionality of a static Deque.
func StaticDequeInterfacePeekBack(
	factory func(capacity int) staticContainers.Deque[int],
	t *testing.T,
) {
	container := factory(5)
	_, err := container.PeekBack()
	test.ContainsError(customerr.ValOutsideRange, err, t)
	test.Nil(container.PushBack(1), t)
	_v, err := container.PeekBack()
	test.Eq(1, _v, t)
	test.Nil(err, t)
	test.Nil(container.PushBack(2), t)
	_v, err = container.PeekBack()
	test.Eq(2, _v, t)
	test.Nil(err, t)
	test.Nil(container.PushBack(3, 4, 5), t)
	_v, err = container.PeekBack()
	test.Eq(5, _v, t)
	test.Nil(err, t)
}

func staticDequePopFrontHelper(
	factory func(capacity int) staticContainers.Deque[int],
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

// Tests the PopFront method functionality of a static Deque.
func StaticDequeInterfacePopFront(
	factory func(capacity int) staticContainers.Deque[int],
	t *testing.T,
) {
	staticDequePopFrontHelper(factory, 5, 5, t)
	staticDequePopFrontHelper(factory, 5, 10, t)
}

func staticDequePopBackHelper(
	factory func(capacity int) staticContainers.Deque[int],
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
	for i := l - 1; i >= 0; i-- {
		f, err := container.PopBack()
		test.Eq(i, container.Length(), t)
		test.Eq(c, container.Capacity(), t)
		test.Eq(i, f, t)
		test.Nil(err, t)
	}
	test.Eq(0, container.Length(), t)
	test.Eq(c, container.Capacity(), t)
	_, err := container.PopBack()
	test.ContainsError(containerTypes.Empty, err, t)
}

// Tests the PopBack method functionality of a static Deque.
func StaticDequeInterfacePopBack(
	factory func(capacity int) staticContainers.Deque[int],
	t *testing.T,
) {
	staticDequePopBackHelper(factory, 5, 5, t)
	staticDequePopBackHelper(factory, 5, 10, t)
}

func dequePushFrontHelper(
	factory func(capacity int) staticContainers.Deque[int],
	l int,
	c int,
	t *testing.T,
) {
	container := factory(c)
	for i := 0; i < l; i++ {
		test.Nil(container.PushFront(i), t)
		test.Eq(i+1, container.Length(), t)
		test.Eq(c, container.Capacity(), t)
		iterV, err := container.PeekFront()
		test.Nil(err, t)
		test.Eq(i, iterV, t)
		iterV, err = container.PeekBack()
		test.Nil(err, t)
		test.Eq(0, iterV, t)
	}
	test.Eq(l, container.Length(), t)
	test.Eq(c, container.Capacity(), t)
}

// Tests the PopFront method functionality of a static Deque.
func StaticDequeInterfacePushFront(
	factory func(capacity int) staticContainers.Deque[int],
	t *testing.T,
) {
	dequePushFrontHelper(factory, 5, 5, t)
	dequePushFrontHelper(factory, 5, 10, t)
	container := factory(6)
	for i := 0; i < 6; i += 2 {
		container.PushFront(i, i+1)
		test.Eq(i+2, container.Length(), t)
		iterV, err := container.PeekFront()
		test.Nil(err, t)
		test.Eq(i, iterV, t)
		iterV, err = container.PeekBack()
		test.Nil(err, t)
		test.Eq(1, iterV, t)
	}
	test.Eq(6, container.Length(), t)
	test.Eq(6, container.Capacity(), t)
}

func staticDequeForcePushFrontHelper(
	factory func(capacity int) staticContainers.Deque[int],
	l int,
	c int,
	t *testing.T,
) {
	container := factory(c)
	for i := 0; i < l; i++ {
		container.ForcePushFront(i)
		test.Eq(i+1, container.Length(), t)
		test.Eq(c, container.Capacity(), t)
		iterV, err := container.PeekFront()
		test.Nil(err, t)
		test.Eq(i, iterV, t)
		iterV, err = container.PeekBack()
		test.Nil(err, t)
		test.Eq(0, iterV, t)
	}
	for i := l; i < c; i++ {
		container.ForcePushFront(i)
		if c <= l {
			test.Eq(l, container.Length(), t)
		} else {
			test.Eq(i+1, container.Length(), t)
		}
		test.Eq(c, container.Capacity(), t)
		iterV, err := container.PeekFront()
		test.Nil(err, t)
		test.Eq(i, iterV, t)
		iterV, err = container.PeekBack()
		test.Nil(err, t)
		if i >= c {
			test.Eq(i-l+1, iterV, t)
		} else {
			test.Eq(0, iterV, t)
		}
	}
}

// Tests the ForcePopFront method functionality of a static Deque.
func StaticDequeInterfaceForcePushFront(
	factory func(capacity int) staticContainers.Deque[int],
	t *testing.T,
) {
	staticDequeForcePushFrontHelper(factory, 5, 5, t)
	staticDequeForcePushFrontHelper(factory, 5, 10, t)

	container := factory(5)
	container.ForcePushFront(1)
	v, err := container.PeekFront()
	test.Nil(err, t)
	test.Eq(1, v, t)
	v, err = container.PeekBack()
	test.Nil(err, t)
	test.Eq(1, v, t)

	container = factory(5)
	container.ForcePushFront(1, 2, 3)
	v, err = container.PeekFront()
	test.Nil(err, t)
	test.Eq(1, v, t)
	v, err = container.PeekBack()
	test.Nil(err, t)
	test.Eq(3, v, t)

	container = factory(5)
	container.ForcePushFront(1, 2, 3, 4, 5, 6, 7)
	v, err = container.PeekFront()
	test.Nil(err, t)
	test.Eq(1, v, t)
	v, err = container.PeekBack()
	test.Nil(err, t)
	test.Eq(5, v, t)
}

func staticDequePushBackHelper(
	factory func(capacity int) staticContainers.Deque[int],
	l int,
	c int,
	t *testing.T,
) {
	container := factory(c)
	for i := 0; i < l; i++ {
		test.Nil(container.PushBack(i), t)
		test.Eq(i+1, container.Length(), t)
		test.Eq(c, container.Capacity(), t)
		iterV, err := container.PeekBack()
		test.Nil(err, t)
		test.Eq(i, iterV, t)
		iterV, err = container.PeekFront()
		test.Nil(err, t)
		test.Eq(0, iterV, t)
	}
	test.Eq(l, container.Length(), t)
	test.Eq(c, container.Capacity(), t)
}

// Tests the PushBack method functionality of a static Deque.
func StaticDequeInterfacePushBack(
	factory func(capacity int) staticContainers.Deque[int],
	t *testing.T,
) {
	staticDequePushBackHelper(factory, 5, 5, t)
	staticDequePushBackHelper(factory, 5, 10, t)
	container := factory(6)
	for i := 0; i < 6; i += 2 {
		test.Nil(container.PushBack(i, i+1), t)
		test.Eq(i+2, container.Length(), t)
		iterV, err := container.PeekBack()
		test.Nil(err, t)
		test.Eq(i+1, iterV, t)
		iterV, err = container.PeekFront()
		test.Nil(err, t)
		test.Eq(0, iterV, t)
	}
	test.Eq(6, container.Length(), t)
	test.Eq(6, container.Capacity(), t)
}

func staticDequeForcePushBackHelper(
	factory func(capacity int) staticContainers.Deque[int],
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
		iterV, err = container.PeekBack()
		test.Nil(err, t)
		test.Eq(i, iterV, t)
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
		iterV, err = container.PeekBack()
		test.Nil(err, t)
		test.Eq(i, iterV, t)
	}
}

// Tests the ForcePopFront method functionality of a static Deque.
func StaticDequeInterfaceForcePushBack(
	factory func(capacity int) staticContainers.Deque[int],
	t *testing.T,
) {
	staticDequeForcePushBackHelper(factory, 5, 5, t)
	staticDequeForcePushBackHelper(factory, 5, 10, t)

	container := factory(5)
	container.ForcePushBack(1)
	v, err := container.PeekFront()
	test.Nil(err, t)
	test.Eq(1, v, t)
	v, err = container.PeekBack()
	test.Nil(err, t)
	test.Eq(1, v, t)

	container = factory(5)
	container.ForcePushBack(1, 2, 3)
	v, err = container.PeekFront()
	test.Nil(err, t)
	test.Eq(1, v, t)
	v, err = container.PeekBack()
	test.Nil(err, t)
	test.Eq(3, v, t)

	container = factory(5)
	container.ForcePushBack(1, 2, 3, 4, 5, 6, 7)
	v, err = container.PeekFront()
	test.Nil(err, t)
	test.Eq(3, v, t)
	v, err = container.PeekBack()
	test.Nil(err, t)
	test.Eq(7, v, t)
}
