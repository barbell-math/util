package containers

// Code generated by ../../../bin/containerInterfaceTests - DO NOT EDIT.
import (
	"github.com/barbell-math/util/src/container/dynamicContainers"
	"github.com/barbell-math/util/src/container/tests"
	"testing"
)

func SyncedVectorToSetInterfaceFactory(capacity int) dynamicContainers.Set[int] {
	v := generateSyncedVector(capacity)
	var rv dynamicContainers.Set[int] = &v
	return rv
}

func TestSyncedVector_DynSetInterfaceSyncableInterface(t *testing.T) {
	tests.DynSetInterfaceSyncableInterface(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceAddressableInterface(t *testing.T) {
	tests.DynSetInterfaceAddressableInterface(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceLengthInterface(t *testing.T) {
	tests.DynSetInterfaceLengthInterface(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceClearInterface(t *testing.T) {
	tests.DynSetInterfaceClearInterface(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceWriteUniqueOpsInterface(t *testing.T) {
	tests.DynSetInterfaceWriteUniqueOpsInterface(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceReadOpsInterface(t *testing.T) {
	tests.DynSetInterfaceReadOpsInterface(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceDeleteOpsInterface(t *testing.T) {
	tests.DynSetInterfaceDeleteOpsInterface(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_ReadDynSetInterface(t *testing.T) {
	tests.ReadDynSetInterface(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_WriteDynSetInterface(t *testing.T) {
	tests.WriteDynSetInterface(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceInterface(t *testing.T) {
	tests.DynSetInterfaceInterface(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceStaticCapacityInterface(t *testing.T) {
	tests.DynSetInterfaceStaticCapacityInterface(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceVals(t *testing.T) {
	tests.DynSetInterfaceVals(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceValPntrs(t *testing.T) {
	tests.DynSetInterfaceValPntrs(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceContainsPntr(t *testing.T) {
	tests.DynSetInterfaceContainsPntr(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceGetUnique(t *testing.T) {
	tests.DynSetInterfaceGetUnique(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceContains(t *testing.T) {
	tests.DynSetInterfaceContains(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceClear(t *testing.T) {
	tests.DynSetInterfaceClear(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceAppendUnique(t *testing.T) {
	tests.DynSetInterfaceAppendUnique(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceUpdateUnique(t *testing.T) {
	tests.DynSetInterfaceUpdateUnique(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfacePop(t *testing.T) {
	tests.DynSetInterfacePop(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfacePopPntr(t *testing.T) {
	tests.DynSetInterfacePopPntr(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceUnorderedEq(t *testing.T) {
	tests.DynSetInterfaceUnorderedEq(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceIntersection(t *testing.T) {
	tests.DynSetInterfaceIntersection(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceUnion(t *testing.T) {
	tests.DynSetInterfaceUnion(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceDifference(t *testing.T) {
	tests.DynSetInterfaceDifference(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceIsSuperset(t *testing.T) {
	tests.DynSetInterfaceIsSuperset(SyncedVectorToSetInterfaceFactory, t)
}

func TestSyncedVector_DynSetInterfaceIsSubset(t *testing.T) {
	tests.DynSetInterfaceIsSubset(SyncedVectorToSetInterfaceFactory, t)
}