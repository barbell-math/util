package containers

// Code generated by ../../../bin/containerInterfaceTests - DO NOT EDIT.
import (
	"github.com/barbell-math/util/src/container/staticContainers"
	"github.com/barbell-math/util/src/container/tests"
	"testing"
)

func SyncedCircularBufferToVectorInterfaceFactory(capacity int) staticContainers.Vector[int] {
	v := generateSyncedCircularBuffer(capacity)
	var rv staticContainers.Vector[int] = &v
	return rv
}

func TestSyncedCircularBuffer_StaticVectorInterfaceStaticCapacity(t *testing.T) {
	tests.StaticVectorInterfaceStaticCapacity(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceAddressableInterface(t *testing.T) {
	tests.StaticVectorInterfaceAddressableInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceLengthInterface(t *testing.T) {
	tests.StaticVectorInterfaceLengthInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceCapacityInterface(t *testing.T) {
	tests.StaticVectorInterfaceCapacityInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceClearInterface(t *testing.T) {
	tests.StaticVectorInterfaceClearInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceWriteOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceWriteOpsInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceWriteKeyedOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceWriteKeyedOpsInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceWriteKeyedSequentialOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceWriteKeyedSequentialOpsInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceWriteDynKeyedOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceWriteDynKeyedOpsInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceReadOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceReadOpsInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceReadKeyedOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceReadKeyedOpsInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceDeleteOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceDeleteOpsInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceDeleteKeyedOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceDeleteKeyedOpsInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceDeleteSequentialOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceDeleteSequentialOpsInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceDeleteKeyedSequentialOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceDeleteKeyedSequentialOpsInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_ReadStaticVectorInterface(t *testing.T) {
	tests.ReadStaticVectorInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_WriteStaticVectorInterface(t *testing.T) {
	tests.WriteStaticVectorInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceInterface(t *testing.T) {
	tests.StaticVectorInterfaceInterface(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceGet(t *testing.T) {
	tests.StaticVectorInterfaceGet(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceGetPntr(t *testing.T) {
	tests.StaticVectorInterfaceGetPntr(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceContains(t *testing.T) {
	tests.StaticVectorInterfaceContains(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceContainsPntr(t *testing.T) {
	tests.StaticVectorInterfaceContainsPntr(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceKeyOf(t *testing.T) {
	tests.StaticVectorInterfaceKeyOf(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceKeyOfPntr(t *testing.T) {
	tests.StaticVectorInterfaceKeyOfPntr(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceSet(t *testing.T) {
	tests.StaticVectorInterfaceSet(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceSetSequential(t *testing.T) {
	tests.StaticVectorInterfaceSetSequential(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceAppend(t *testing.T) {
	tests.StaticVectorInterfaceAppend(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceInsert(t *testing.T) {
	tests.StaticVectorInterfaceInsert(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceInsertSequential(t *testing.T) {
	tests.StaticVectorInterfaceInsertSequential(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfacePopSequential(t *testing.T) {
	tests.StaticVectorInterfacePopSequential(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfacePop(t *testing.T) {
	tests.StaticVectorInterfacePop(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfacePopPntr(t *testing.T) {
	tests.StaticVectorInterfacePopPntr(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceDelete(t *testing.T) {
	tests.StaticVectorInterfaceDelete(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceDeleteSequential(t *testing.T) {
	tests.StaticVectorInterfaceDeleteSequential(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceClear(t *testing.T) {
	tests.StaticVectorInterfaceClear(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceVals(t *testing.T) {
	tests.StaticVectorInterfaceVals(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceValPntrs(t *testing.T) {
	tests.StaticVectorInterfaceValPntrs(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceKeys(t *testing.T) {
	tests.StaticVectorInterfaceKeys(SyncedCircularBufferToVectorInterfaceFactory, t)
}

func TestSyncedCircularBuffer_StaticVectorInterfaceKeyedEq(t *testing.T) {
	tests.StaticVectorInterfaceKeyedEq(SyncedCircularBufferToVectorInterfaceFactory, t)
}
