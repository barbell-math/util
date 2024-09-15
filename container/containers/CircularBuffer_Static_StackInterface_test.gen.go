package containers

// Code generated by ../../bin/containerInterfaceTests - DO NOT EDIT.
import (
	"testing"
	"github.com/barbell-math/util/container/tests"
	"github.com/barbell-math/util/container/staticContainers"
)

func CircularBufferToStackInterfaceFactory(capacity int) staticContainers.Stack[int] {
	v := generateCircularBuffer(capacity)
	var rv staticContainers.Stack[int] = &v
	return rv
}

func TestCircularBuffer_StaticStackInterfaceStaticCapacity(t *testing.T) {
	tests.StaticStackInterfaceStaticCapacity(CircularBufferToStackInterfaceFactory, t)
}

func TestCircularBuffer_StaticStackInterfaceLengthInterface(t *testing.T) {
	tests.StaticStackInterfaceLengthInterface(CircularBufferToStackInterfaceFactory, t)
}

func TestCircularBuffer_StaticStackInterfaceCapacityInterface(t *testing.T) {
	tests.StaticStackInterfaceCapacityInterface(CircularBufferToStackInterfaceFactory, t)
}

func TestCircularBuffer_StaticStackInterfaceClearInterface(t *testing.T) {
	tests.StaticStackInterfaceClearInterface(CircularBufferToStackInterfaceFactory, t)
}

func TestCircularBuffer_StaticStackInterfaceLastElemReadInterface(t *testing.T) {
	tests.StaticStackInterfaceLastElemReadInterface(CircularBufferToStackInterfaceFactory, t)
}

func TestCircularBuffer_StaticStackInterfaceLastElemWriteInterface(t *testing.T) {
	tests.StaticStackInterfaceLastElemWriteInterface(CircularBufferToStackInterfaceFactory, t)
}

func TestCircularBuffer_StaticStackInterfaceLastElemDeleteInterface(t *testing.T) {
	tests.StaticStackInterfaceLastElemDeleteInterface(CircularBufferToStackInterfaceFactory, t)
}

func TestCircularBuffer_ReadStaticStackInterface(t *testing.T) {
	tests.ReadStaticStackInterface(CircularBufferToStackInterfaceFactory, t)
}

func TestCircularBuffer_WriteStaticStackInterface(t *testing.T) {
	tests.WriteStaticStackInterface(CircularBufferToStackInterfaceFactory, t)
}

func TestCircularBuffer_StaticStackInterfaceInterface(t *testing.T) {
	tests.StaticStackInterfaceInterface(CircularBufferToStackInterfaceFactory, t)
}

func TestCircularBuffer_StaticStackInterfaceClear(t *testing.T) {
	tests.StaticStackInterfaceClear(CircularBufferToStackInterfaceFactory, t)
}

func TestCircularBuffer_StaticStackInterfacePeekPntrBack(t *testing.T) {
	tests.StaticStackInterfacePeekPntrBack(CircularBufferToStackInterfaceFactory, t)
}

func TestCircularBuffer_StaticStackInterfacePeekBack(t *testing.T) {
	tests.StaticStackInterfacePeekBack(CircularBufferToStackInterfaceFactory, t)
}

func TestCircularBuffer_StaticStackInterfacePopBack(t *testing.T) {
	tests.StaticStackInterfacePopBack(CircularBufferToStackInterfaceFactory, t)
}

func TestCircularBuffer_StaticStackInterfacePushBack(t *testing.T) {
	tests.StaticStackInterfacePushBack(CircularBufferToStackInterfaceFactory, t)
}

func TestCircularBuffer_StaticStackInterfaceForcePushBack(t *testing.T) {
	tests.StaticStackInterfaceForcePushBack(CircularBufferToStackInterfaceFactory, t)
}
