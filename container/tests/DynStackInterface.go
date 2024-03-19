package tests

import (
	"testing"

	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func dynStackReadInterface[U any](c dynamicContainers.ReadStack[U])   {}
func dynStackWriteInterface[U any](c dynamicContainers.WriteStack[U]) {}
func dynStackInterface[U any](c dynamicContainers.Stack[U])           {}

// Tests that the value supplied by the factory implements the
// [containerTypes.Length] interface.
func DynStackInterfaceLengthInterface[V any](
	factory func(capacity int) dynamicContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.Length = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Capacity] interface.
func DynStackInterfaceCapacityInterface[V any](
	factory func(capacity int) dynamicContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.Capacity = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.Clear] interface.
func DynStackInterfaceClearInterface[V any](
	factory func(capacity int) dynamicContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.LastElemRead] interface.
func DynStackInterfaceLastElemReadInterface[V any](
	factory func(capacity int) dynamicContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.LastElemRead[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.LastElemWrite] interface.
func DynStackInterfaceLastElemWriteInterface[V any](
	factory func(capacity int) dynamicContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.LastElemWrite[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [containerTypes.LastElemDelete] interface.
func DynStackInterfaceLastElemDeleteInterface[V any](
	factory func(capacity int) dynamicContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.LastElemDelete[V] = factory(0)
	_ = container
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.StackRead] interface.
func ReadDynStackInterface[V any](
	factory func(capacity int) dynamicContainers.Stack[V],
	t *testing.T,
) {
	dynStackReadInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.WriteStack] interface.
func WriteDynStackInterface[V any](
	factory func(capacity int) dynamicContainers.Stack[V],
	t *testing.T,
) {
	dynStackWriteInterface[V](factory(0))
}

// Tests that the value supplied by the factory implements the
// [dynamicContainers.Stack] interface.
func DynStackInterfaceInterface[V any](
	factory func(capacity int) dynamicContainers.Stack[V],
	t *testing.T,
) {
	dynStackInterface[V](factory(0))
}

// Tests that the value supplied by the factory does not implement the
// [staticContainers.Stack] interface.
func DynStackInterfaceStaticCapacityInterface[V any](
	factory func(capacity int) dynamicContainers.Stack[V],
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

// Tests the Clear method functionality of a dynamic Stack.
func DynStackInterfaceClear(
	factory func(capacity int) dynamicContainers.Stack[int],
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

// Tests the PeekPntrBack method functionality of a dynamic Stack.
func DynStackInterfacePeekPntrBack(
	factory func(capacity int) dynamicContainers.Stack[int],
	t *testing.T,
) {
	container := factory(0)
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
}

// Tests the PeekBack method functionality of a dynamic Stack.
func DynStackInterfacePeekBack(
	factory func(capacity int) dynamicContainers.Stack[int],
	t *testing.T,
) {
	container := factory(0)
	_, err := container.PeekBack()
	test.ContainsError(customerr.ValOutsideRange, err, t)
	container.PushBack(1)
	_v, err := container.PeekBack()
	test.Eq(1, _v, t)
	test.Nil(err, t)
	container.PushBack(2)
	_v, err = container.PeekBack()
	test.Eq(2, _v, t)
	test.Nil(err, t)
}

// Tests the PopBack method functionality of a dynamic Stack.
func DynStackInterfacePopBack(
	factory func(capacity int) dynamicContainers.Stack[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 4; i++ {
		container.PushBack(i)
	}
	for i := 3; i >= 0; i-- {
		f, err := container.PopBack()
		test.Eq(i, f, t)
		test.Nil(err, t)
	}
	_, err := container.PopBack()
	test.ContainsError(containerTypes.Empty, err, t)
}

// Tests the PushBack method functionality of a dynamic Stack.
func DynStackInterfacePushBack(
	factory func(capacity int) dynamicContainers.Stack[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 4; i++ {
		container.PushBack(i)
		test.Eq(i+1, container.Length(), t)
		iterV, _ := container.PeekBack()
		test.Eq(i, iterV, t)
	}
	container = factory(0)
	for i := 0; i < 6; i += 2 {
		container.PushBack(i, i+1)
		test.Eq(i+2, container.Length(), t)
		iterV, _ := container.PeekBack()
		test.Eq(i+1, iterV, t)
	}
}

// Tests the ForcePushBack method functionality of a dynamic Stack.
func DynStackInterfaceForcePushBack(
	factory func(capacity int) dynamicContainers.Stack[int],
	t *testing.T,
) {
	container := factory(0)
	for i := 0; i < 4; i++ {
		container.ForcePushBack(i)
		test.Eq(i+1, container.Length(), t)
		iterV, _ := container.PeekBack()
		test.Eq(i, iterV, t)
	}
	container = factory(0)
	for i := 0; i < 6; i += 2 {
		container.ForcePushBack(i, i+1)
		test.Eq(i+2, container.Length(), t)
		iterV, _ := container.PeekBack()
		test.Eq(i+1, iterV, t)
	}
}
