package tests

import (
	"testing"

	"github.com/barbell-math/util/container/containerTypes"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func stackReadInterface[U any](c dynamicContainers.ReadStack[U])   {}
func stackWriteInterface[U any](c dynamicContainers.WriteStack[U]) {}
func stackInterface[U any](c dynamicContainers.Stack[U])           {}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Length] interface.
func StackInterfaceLengthInterface[V any](
	factory func() dynamicContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.Length = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Capacity] interface.
func StackInterfaceCapacityInterface[V any](
	factory func() dynamicContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.Capacity = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.Clear] interface.
func StackInterfaceClearInterface[V any](
	factory func() dynamicContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.Clear = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.LastElemRead] interface.
func StackInterfaceLastElemReadInterface[V any](
	factory func() dynamicContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.LastElemRead[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.LastElemWrite] interface.
func StackInterfaceLastElemWriteInterface[V any](
	factory func() dynamicContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.LastElemWrite[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [containerTypes.LastElemDelete] interface.
func StackInterfaceLastElemDeleteInterface[V any](
	factory func() dynamicContainers.Stack[V],
	t *testing.T,
) {
	var container containerTypes.LastElemDelete[V] = factory()
	_ = container
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.StackRead] interface.
func ReadStackInterface[V any](
	factory func() dynamicContainers.Stack[V],
	t *testing.T,
) {
	stackReadInterface[V](factory())
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.WriteStack] interface.
func WriteStackInterface[V any](
	factory func() dynamicContainers.Stack[V],
	t *testing.T,
) {
	stackWriteInterface[V](factory())
}

// Tests that the value supplied by the factory implements the 
// [dynamicContainers.Stack] interface.
func StackInterfaceInterface[V any](
	factory func() dynamicContainers.Stack[V],
	t *testing.T,
) {
	stackInterface[V](factory())
}

// Tests that the value supplied by the factory does not implement the 
// [staticContainers.Stack] interface.
func StackInterfaceStaticCapacityInterface[V any](
	factory func() dynamicContainers.Stack[V],
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

// Tests the Clear method functionality of a dynamic Stack.
func StackInterfaceClear(
	factory func() dynamicContainers.Stack[int],
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

// Tests the PeekPntrBack method functionality of a dynamic Stack.
func StackInterfacePeekPntrBack(
	factory func() dynamicContainers.Stack[int],
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

// Tests the PeekBack method functionality of a dynamic Stack.
func StackInterfacePeekBack(
	factory func() dynamicContainers.Stack[int],
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

// Tests the PopBack method functionality of a dynamic Stack.
func StackInterfacePopBack(
	factory func() dynamicContainers.Stack[int],
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

// Tests the PushBack method functionality of a dynamic Stack.
func StackInterfacePushBack(
	factory func() dynamicContainers.Stack[int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 4; i++ {
		container.PushBack(i)
		test.Eq(i+1, container.Length(),t)
		iterV, _ := container.PeekBack()
		test.Eq(i, iterV,t)
	}
	container=factory()
	for i := 0; i < 6; i+=2 {
		container.PushBack(i,i+1)
		test.Eq(i+2, container.Length(),t)
		iterV, _ := container.PeekBack()
		test.Eq(i+1, iterV,t)
	}
}

// Tests the ForcePushBack method functionality of a dynamic Stack.
func StackInterfaceForcePushBack(
	factory func() dynamicContainers.Stack[int],
	t *testing.T,
) {
	container := factory()
	for i := 0; i < 4; i++ {
		container.ForcePushBack(i)
		test.Eq(i+1, container.Length(),t)
		iterV, _ := container.PeekBack()
		test.Eq(i, iterV,t)
	}
	container=factory()
	for i := 0; i < 6; i+=2 {
		container.ForcePushBack(i,i+1)
		test.Eq(i+2, container.Length(),t)
		iterV, _ := container.PeekBack()
		test.Eq(i+1, iterV,t)
	}
}
