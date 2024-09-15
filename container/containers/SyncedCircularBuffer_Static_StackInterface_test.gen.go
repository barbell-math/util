package containers

// Code generated by ../../bin/containerInterfaceTests - DO NOT EDIT.
import (
	"testing"
	"github.com/barbell-math/util/container/tests"
	"github.com/barbell-math/util/container/staticContainers"
)

func SyncedCircularBufferToStackInterfaceFactory(capacity int) staticContainers.Stack[int] {
	v := generateSyncedCircularBuffer(capacity)
	var rv staticContainers.Stack[int] = &v
	return rv
}

func TestSyncedCircularBuffer_StaticStackInterfaceStaticCapacity(t *testing.T) {
	tests.StaticStackInterfaceStaticCapacity(SyncedCircularBufferToStackInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticStackInterfaceLengthInterface(t *testing.T) {
	tests.StaticStackInterfaceLengthInterface(SyncedCircularBufferToStackInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticStackInterfaceCapacityInterface(t *testing.T) {
	tests.StaticStackInterfaceCapacityInterface(SyncedCircularBufferToStackInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticStackInterfaceClearInterface(t *testing.T) {
	tests.StaticStackInterfaceClearInterface(SyncedCircularBufferToStackInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticStackInterfaceLastElemReadInterface(t *testing.T) {
	tests.StaticStackInterfaceLastElemReadInterface(SyncedCircularBufferToStackInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticStackInterfaceLastElemWriteInterface(t *testing.T) {
	tests.StaticStackInterfaceLastElemWriteInterface(SyncedCircularBufferToStackInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticStackInterfaceLastElemDeleteInterface(t *testing.T) {
	tests.StaticStackInterfaceLastElemDeleteInterface(SyncedCircularBufferToStackInterfaceFactory, t)
}

func TestSyncedCircularBuffer_ReadStaticStackInterface(t *testing.T) {
	tests.ReadStaticStackInterface(SyncedCircularBufferToStackInterfaceFactory, t)
}

func TestSyncedCircularBuffer_WriteStaticStackInterface(t *testing.T) {
	tests.WriteStaticStackInterface(SyncedCircularBufferToStackInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticStackInterfaceInterface(t *testing.T) {
	tests.StaticStackInterfaceInterface(SyncedCircularBufferToStackInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticStackInterfaceClear(t *testing.T) {
	tests.StaticStackInterfaceClear(SyncedCircularBufferToStackInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticStackInterfacePeekPntrBack(t *testing.T) {
	tests.StaticStackInterfacePeekPntrBack(SyncedCircularBufferToStackInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticStackInterfacePeekBack(t *testing.T) {
	tests.StaticStackInterfacePeekBack(SyncedCircularBufferToStackInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticStackInterfacePopBack(t *testing.T) {
	tests.StaticStackInterfacePopBack(SyncedCircularBufferToStackInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticStackInterfacePushBack(t *testing.T) {
	tests.StaticStackInterfacePushBack(SyncedCircularBufferToStackInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticStackInterfaceForcePushBack(t *testing.T) {
	tests.StaticStackInterfaceForcePushBack(SyncedCircularBufferToStackInterfaceFactory, t)
}
