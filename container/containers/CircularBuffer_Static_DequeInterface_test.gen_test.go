package containers

// Code generated by ../../bin/containerInterfaceTests - DO NOT EDIT.
import (
	"github.com/barbell-math/util/container/staticContainers"
	"github.com/barbell-math/util/container/tests"
	"testing"
)

func CircularBufferToDequeInterfaceFactory(capacity int) staticContainers.Deque[int] {
	v := generateCircularBuffer(capacity)
	var rv staticContainers.Deque[int] = &v
	return rv
}

func TestCircularBuffer_StaticDequeInterfaceStaticCapacity(t *testing.T) {
	tests.StaticDequeInterfaceStaticCapacity(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfaceLengthInterface(t *testing.T) {
	tests.StaticDequeInterfaceLengthInterface(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfaceCapacityInterface(t *testing.T) {
	tests.StaticDequeInterfaceCapacityInterface(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfaceClearInterface(t *testing.T) {
	tests.StaticDequeInterfaceClearInterface(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfaceFirstElemReadInterface(t *testing.T) {
	tests.StaticDequeInterfaceFirstElemReadInterface(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfaceFirstElemWriteInterface(t *testing.T) {
	tests.StaticDequeInterfaceFirstElemWriteInterface(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfaceFirstElemDeleteInterface(t *testing.T) {
	tests.StaticDequeInterfaceFirstElemDeleteInterface(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfaceLastElemReadInterface(t *testing.T) {
	tests.StaticDequeInterfaceLastElemReadInterface(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfaceLastElemWriteInterface(t *testing.T) {
	tests.StaticDequeInterfaceLastElemWriteInterface(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfaceLastElemDeleteInterface(t *testing.T) {
	tests.StaticDequeInterfaceLastElemDeleteInterface(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_ReadStaticDequeInterface(t *testing.T) {
	tests.ReadStaticDequeInterface(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_WriteStaticDequeInterface(t *testing.T) {
	tests.WriteStaticDequeInterface(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfaceInterface(t *testing.T) {
	tests.StaticDequeInterfaceInterface(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfaceClear(t *testing.T) {
	tests.StaticDequeInterfaceClear(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfacePeekPntrFront(t *testing.T) {
	tests.StaticDequeInterfacePeekPntrFront(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfacePeekFront(t *testing.T) {
	tests.StaticDequeInterfacePeekFront(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfacePeekPntrBack(t *testing.T) {
	tests.StaticDequeInterfacePeekPntrBack(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfacePeekBack(t *testing.T) {
	tests.StaticDequeInterfacePeekBack(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfacePopFront(t *testing.T) {
	tests.StaticDequeInterfacePopFront(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfacePopBack(t *testing.T) {
	tests.StaticDequeInterfacePopBack(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfacePushFront(t *testing.T) {
	tests.StaticDequeInterfacePushFront(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfaceForcePushFront(t *testing.T) {
	tests.StaticDequeInterfaceForcePushFront(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfacePushBack(t *testing.T) {
	tests.StaticDequeInterfacePushBack(CircularBufferToDequeInterfaceFactory, t)
}

func TestCircularBuffer_StaticDequeInterfaceForcePushBack(t *testing.T) {
	tests.StaticDequeInterfaceForcePushBack(CircularBufferToDequeInterfaceFactory, t)
}