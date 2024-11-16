package tests

import (
	"testing"

	"github.com/barbell-math/util/src/container/containerTypes"
	"github.com/barbell-math/util/src/container/staticContainers"
	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/test"
)

func staticStackReadInterface[U any](c staticContainers.ReadStack[U])   {}
func staticStackWriteInterface[U any](c staticContainers.WriteStack[U]) {}
func staticStackInterface[U any](c staticContainers.Stack[U])           {}

// Tests that the value supplied by the factory implements the
// [containerTypes.StaticCapacity] interface.
func StaticStackInterfaceStaticCapacity[V any](
	factory func(capacity int) staticContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.StaticCapacity = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Length] interface.
func StaticStackInterfaceLengthInterface[V any](
	factory func(capacity int) staticContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.Length = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Capacity] interface.
func StaticStackInterfaceCapacityInterface[V any](
	factory func(capacity int) staticContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.Capacity = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Clear] interface.
func StaticStackInterfaceClearInterface[V any](
	factory func(capacity int) staticContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.LastElemRead] interface.
func StaticStackInterfaceLastElemReadInterface[V any](
	factory func(capacity int) staticContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.LastElemRead[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.LastElemWrite] interface.
func StaticStackInterfaceLastElemWriteInterface[V any](
	factory func(capacity int) staticContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.LastElemWrite[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.LastElemDelete] interface.
func StaticStackInterfaceLastElemDeleteInterface[V any](
	factory func(capacity int) staticContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.LastElemDelete[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [staticContainers.StackRead] interface.
func ReadStaticStackInterface[V any](
	factory func(capacity int) staticContainers.Stack[V],
	t *testing.T,
) {
	staticStackReadInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the
// [staticContainers.WriteStack] interface.
func WriteStaticStackInterface[V any](
	factory func(capacity int) staticContainers.Stack[V],
	t *testing.T,
) {
	staticStackWriteInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the
// [staticContainers.Stack] interface.
func StaticStackInterfaceInterface[V any](
	factory func(capacity int) staticContainers.Stack[V],
	t *testing.T,
) {
	staticStackInterface[V](factory(0))
}

// Tests the Clear method functionality of a static Stack.
func StaticStackInterfaceClear(
	factory func(capacity int) staticContainers.Stack[int],
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

// Tests the PeekPntrBack method functionality of a static Stack.
func StaticStackInterfacePeekPntrBack(
	factory func(capacity int) staticContainers.Stack[int],
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

// Tests the PeekBack method functionality of a static Stack.
func StaticStackInterfacePeekBack(
	factory func(capacity int) staticContainers.Stack[int],
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

func staticStackPopBackHelper(
	factory func(capacity int) staticContainers.Stack[int],
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

// Tests the PopBack method functionality of a static Stack.
func StaticStackInterfacePopBack(
	factory func(capacity int) staticContainers.Stack[int],
	t *testing.T,
) {
	staticStackPopBackHelper(factory, 5, 5, t)
	staticStackPopBackHelper(factory, 5, 10, t)
}

func staticStackPushBackHelper(
	factory func(capacity int) staticContainers.Stack[int],
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
	}
	test.Eq(l, container.Length(), t)
	test.Eq(c, container.Capacity(), t)
}

// Tests the PushBack method functionality of a static Stack.
func StaticStackInterfacePushBack(
	factory func(capacity int) staticContainers.Stack[int],
	t *testing.T,
) {
	staticStackPushBackHelper(factory, 5, 5, t)
	staticStackPushBackHelper(factory, 5, 10, t)
	container := factory(6)
	for i := 0; i < 6; i += 2 {
		test.Nil(container.PushBack(i, i+1), t)
		test.Eq(i+2, container.Length(), t)
		iterV, _ := container.PeekBack()
		test.Eq(i+1, iterV, t)
	}
	test.Eq(6, container.Length(), t)
	test.Eq(6, container.Capacity(), t)
}

func staticStackForcePushBackHelper(
	factory func(capacity int) staticContainers.Stack[int],
	l int,
	c int,
	t *testing.T,
) {
	container := factory(c)
	for i := 0; i < l; i++ {
		container.ForcePushBack(i)
		test.Eq(i+1, container.Length(), t)
		test.Eq(c, container.Capacity(), t)
		iterV, err := container.PeekBack()
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
		iterV, err := container.PeekBack()
		test.Nil(err, t)
		test.Eq(i, iterV, t)
	}
}

// Tests the ForcePopFront method functionality of a static Stack.
func StaticStackInterfaceForcePushBack(
	factory func(capacity int) staticContainers.Stack[int],
	t *testing.T,
) {
	staticStackForcePushBackHelper(factory, 5, 5, t)
	staticStackForcePushBackHelper(factory, 5, 10, t)

	container := factory(5)
	container.ForcePushBack(1)
	v, err := container.PeekBack()
	test.Nil(err, t)
	test.Eq(1, v, t)

	container = factory(5)
	container.ForcePushBack(1, 2, 3)
	v, err = container.PeekBack()
	test.Nil(err, t)
	test.Eq(3, v, t)

	container = factory(5)
	container.ForcePushBack(1, 2, 3, 4, 5, 6, 7)
	v, err = container.PeekBack()
	test.Nil(err, t)
	test.Eq(7, v, t)
}
