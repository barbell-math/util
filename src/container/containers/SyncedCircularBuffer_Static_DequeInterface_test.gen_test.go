package containers

// Code generated by ../../../bin/containerInterfaceTests - DO NOT EDIT.
import (
	"github.com/barbell-math/util/src/container/staticContainers"
	"github.com/barbell-math/util/src/container/tests"
	"testing"
)

func SyncedCircularBufferToDequeInterfaceFactory(capacity int) staticContainers.Deque[int] {
	v := generateSyncedCircularBuffer(capacity)
	var rv staticContainers.Deque[int] = &v
	return rv
}

func TestSyncedCircularBuffer_StaticDequeInterfaceStaticCapacity(t *testing.T) {
	tests.StaticDequeInterfaceStaticCapacity(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfaceLengthInterface(t *testing.T) {
	tests.StaticDequeInterfaceLengthInterface(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfaceCapacityInterface(t *testing.T) {
	tests.StaticDequeInterfaceCapacityInterface(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfaceClearInterface(t *testing.T) {
	tests.StaticDequeInterfaceClearInterface(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfaceFirstElemReadInterface(t *testing.T) {
	tests.StaticDequeInterfaceFirstElemReadInterface(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfaceFirstElemWriteInterface(t *testing.T) {
	tests.StaticDequeInterfaceFirstElemWriteInterface(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfaceFirstElemDeleteInterface(t *testing.T) {
	tests.StaticDequeInterfaceFirstElemDeleteInterface(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfaceLastElemReadInterface(t *testing.T) {
	tests.StaticDequeInterfaceLastElemReadInterface(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfaceLastElemWriteInterface(t *testing.T) {
	tests.StaticDequeInterfaceLastElemWriteInterface(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfaceLastElemDeleteInterface(t *testing.T) {
	tests.StaticDequeInterfaceLastElemDeleteInterface(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_ReadStaticDequeInterface(t *testing.T) {
	tests.ReadStaticDequeInterface(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_WriteStaticDequeInterface(t *testing.T) {
	tests.WriteStaticDequeInterface(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfaceInterface(t *testing.T) {
	tests.StaticDequeInterfaceInterface(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfaceClear(t *testing.T) {
	tests.StaticDequeInterfaceClear(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfacePeekPntrFront(t *testing.T) {
	tests.StaticDequeInterfacePeekPntrFront(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfacePeekFront(t *testing.T) {
	tests.StaticDequeInterfacePeekFront(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfacePeekPntrBack(t *testing.T) {
	tests.StaticDequeInterfacePeekPntrBack(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfacePeekBack(t *testing.T) {
	tests.StaticDequeInterfacePeekBack(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfacePopFront(t *testing.T) {
	tests.StaticDequeInterfacePopFront(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfacePopBack(t *testing.T) {
	tests.StaticDequeInterfacePopBack(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfacePushFront(t *testing.T) {
	tests.StaticDequeInterfacePushFront(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfaceForcePushFront(t *testing.T) {
	tests.StaticDequeInterfaceForcePushFront(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfacePushBack(t *testing.T) {
	tests.StaticDequeInterfacePushBack(SyncedCircularBufferToDequeInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticDequeInterfaceForcePushBack(t *testing.T) {
	tests.StaticDequeInterfaceForcePushBack(SyncedCircularBufferToDequeInterfaceFactory, t)
}
