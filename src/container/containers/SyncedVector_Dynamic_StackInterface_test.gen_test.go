package containers

// Code generated by ../../../bin/containerInterfaceTests - DO NOT EDIT.
import (
	"github.com/barbell-math/util/src/container/dynamicContainers"
	"github.com/barbell-math/util/src/container/tests"
	"testing"
)

func SyncedVectorToStackInterfaceFactory(capacity int) dynamicContainers.Stack[int] {
	v := generateSyncedVector(capacity)
	var rv dynamicContainers.Stack[int] = &v
	return rv
}

func TestSyncedVector_DynStackInterfaceSyncableInterface(t *testing.T) {
	tests.DynStackInterfaceSyncableInterface(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_DynStackInterfaceAddressableInterface(t *testing.T) {
	tests.DynStackInterfaceAddressableInterface(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_DynStackInterfaceLengthInterface(t *testing.T) {
	tests.DynStackInterfaceLengthInterface(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_DynStackInterfaceCapacityInterface(t *testing.T) {
	tests.DynStackInterfaceCapacityInterface(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_DynStackInterfaceClearInterface(t *testing.T) {
	tests.DynStackInterfaceClearInterface(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_DynStackInterfaceLastElemReadInterface(t *testing.T) {
	tests.DynStackInterfaceLastElemReadInterface(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_DynStackInterfaceLastElemWriteInterface(t *testing.T) {
	tests.DynStackInterfaceLastElemWriteInterface(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_DynStackInterfaceLastElemDeleteInterface(t *testing.T) {
	tests.DynStackInterfaceLastElemDeleteInterface(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_ReadDynStackInterface(t *testing.T) {
	tests.ReadDynStackInterface(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_WriteDynStackInterface(t *testing.T) {
	tests.WriteDynStackInterface(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_DynStackInterfaceInterface(t *testing.T) {
	tests.DynStackInterfaceInterface(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_DynStackInterfaceStaticCapacityInterface(t *testing.T) {
	tests.DynStackInterfaceStaticCapacityInterface(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_DynStackInterfaceClear(t *testing.T) {
	tests.DynStackInterfaceClear(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_DynStackInterfacePeekPntrBack(t *testing.T) {
	tests.DynStackInterfacePeekPntrBack(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_DynStackInterfacePeekBack(t *testing.T) {
	tests.DynStackInterfacePeekBack(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_DynStackInterfacePopBack(t *testing.T) {
	tests.DynStackInterfacePopBack(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_DynStackInterfacePushBack(t *testing.T) {
	tests.DynStackInterfacePushBack(SyncedVectorToStackInterfaceFactory, t)
}

func TestSyncedVector_DynStackInterfaceForcePushBack(t *testing.T) {
	tests.DynStackInterfaceForcePushBack(SyncedVectorToStackInterfaceFactory, t)
}
