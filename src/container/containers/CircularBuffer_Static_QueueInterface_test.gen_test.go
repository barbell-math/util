package containers

// Code generated by ../../../bin/containerInterfaceTests - DO NOT EDIT.
import (
	"github.com/barbell-math/util/src/container/staticContainers"
	"github.com/barbell-math/util/src/container/tests"
	"testing"
)

func CircularBufferToQueueInterfaceFactory(capacity int) staticContainers.Queue[int] {
	v := generateCircularBuffer(capacity)
	var rv staticContainers.Queue[int] = &v
	return rv
}

func TestCircularBuffer_StaticQueueInterfaceStaticCapacity(t *testing.T) {
	tests.StaticQueueInterfaceStaticCapacity(CircularBufferToQueueInterfaceFactory, t)
}

func TestCircularBuffer_StaticQueueInterfaceLengthInterface(t *testing.T) {
	tests.StaticQueueInterfaceLengthInterface(CircularBufferToQueueInterfaceFactory, t)
}

func TestCircularBuffer_StaticQueueInterfaceCapacityInterface(t *testing.T) {
	tests.StaticQueueInterfaceCapacityInterface(CircularBufferToQueueInterfaceFactory, t)
}

func TestCircularBuffer_StaticQueueInterfaceClearInterface(t *testing.T) {
	tests.StaticQueueInterfaceClearInterface(CircularBufferToQueueInterfaceFactory, t)
}

func TestCircularBuffer_StaticQueueInterfaceFirstElemReadInterface(t *testing.T) {
	tests.StaticQueueInterfaceFirstElemReadInterface(CircularBufferToQueueInterfaceFactory, t)
}

func TestCircularBuffer_StaticQueueInterfaceFirstElemDeleteInterface(t *testing.T) {
	tests.StaticQueueInterfaceFirstElemDeleteInterface(CircularBufferToQueueInterfaceFactory, t)
}

func TestCircularBuffer_StaticQueueInterfaceLastElemWriteInterface(t *testing.T) {
	tests.StaticQueueInterfaceLastElemWriteInterface(CircularBufferToQueueInterfaceFactory, t)
}

func TestCircularBuffer_ReadStaticQueueInterface(t *testing.T) {
	tests.ReadStaticQueueInterface(CircularBufferToQueueInterfaceFactory, t)
}

func TestCircularBuffer_WriteStaticQueueInterface(t *testing.T) {
	tests.WriteStaticQueueInterface(CircularBufferToQueueInterfaceFactory, t)
}

func TestCircularBuffer_StaticQueueInterfaceInterface(t *testing.T) {
	tests.StaticQueueInterfaceInterface(CircularBufferToQueueInterfaceFactory, t)
}

func TestCircularBuffer_StaticQueueInterfaceClear(t *testing.T) {
	tests.StaticQueueInterfaceClear(CircularBufferToQueueInterfaceFactory, t)
}

func TestCircularBuffer_StaticQueueInterfacePeekPntrFront(t *testing.T) {
	tests.StaticQueueInterfacePeekPntrFront(CircularBufferToQueueInterfaceFactory, t)
}

func TestCircularBuffer_StaticQueueInterfacePeekFront(t *testing.T) {
	tests.StaticQueueInterfacePeekFront(CircularBufferToQueueInterfaceFactory, t)
}

func TestCircularBuffer_StaticQueueInterfacePopFront(t *testing.T) {
	tests.StaticQueueInterfacePopFront(CircularBufferToQueueInterfaceFactory, t)
}

func TestCircularBuffer_StaticQueueInterfacePushBack(t *testing.T) {
	tests.StaticQueueInterfacePushBack(CircularBufferToQueueInterfaceFactory, t)
}

func TestCircularBuffer_StaticQueueInterfaceForcePushBack(t *testing.T) {
	tests.StaticQueueInterfaceForcePushBack(CircularBufferToQueueInterfaceFactory, t)
}
