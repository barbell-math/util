package containers

// Code generated by ../../../bin/containerInterfaceTests - DO NOT EDIT.
import (
	"github.com/barbell-math/util/src/container/dynamicContainers"
	"github.com/barbell-math/util/src/container/tests"
	"testing"
)

func VectorToSetInterfaceFactory(capacity int) dynamicContainers.Set[int] {
	v := generateVector(capacity)
	var rv dynamicContainers.Set[int] = &v
	return rv
}

func TestVector_DynSetInterfaceSyncableInterface(t *testing.T) {
	tests.DynSetInterfaceSyncableInterface(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceAddressableInterface(t *testing.T) {
	tests.DynSetInterfaceAddressableInterface(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceLengthInterface(t *testing.T) {
	tests.DynSetInterfaceLengthInterface(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceClearInterface(t *testing.T) {
	tests.DynSetInterfaceClearInterface(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceWriteUniqueOpsInterface(t *testing.T) {
	tests.DynSetInterfaceWriteUniqueOpsInterface(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceReadOpsInterface(t *testing.T) {
	tests.DynSetInterfaceReadOpsInterface(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceDeleteOpsInterface(t *testing.T) {
	tests.DynSetInterfaceDeleteOpsInterface(VectorToSetInterfaceFactory, t)
}

func TestVector_ReadDynSetInterface(t *testing.T) {
	tests.ReadDynSetInterface(VectorToSetInterfaceFactory, t)
}

func TestVector_WriteDynSetInterface(t *testing.T) {
	tests.WriteDynSetInterface(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceInterface(t *testing.T) {
	tests.DynSetInterfaceInterface(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceStaticCapacityInterface(t *testing.T) {
	tests.DynSetInterfaceStaticCapacityInterface(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceVals(t *testing.T) {
	tests.DynSetInterfaceVals(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceValPntrs(t *testing.T) {
	tests.DynSetInterfaceValPntrs(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceContainsPntr(t *testing.T) {
	tests.DynSetInterfaceContainsPntr(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceGetUnique(t *testing.T) {
	tests.DynSetInterfaceGetUnique(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceContains(t *testing.T) {
	tests.DynSetInterfaceContains(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceClear(t *testing.T) {
	tests.DynSetInterfaceClear(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceAppendUnique(t *testing.T) {
	tests.DynSetInterfaceAppendUnique(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceUpdateUnique(t *testing.T) {
	tests.DynSetInterfaceUpdateUnique(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfacePop(t *testing.T) {
	tests.DynSetInterfacePop(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfacePopPntr(t *testing.T) {
	tests.DynSetInterfacePopPntr(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceUnorderedEq(t *testing.T) {
	tests.DynSetInterfaceUnorderedEq(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceIntersection(t *testing.T) {
	tests.DynSetInterfaceIntersection(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceUnion(t *testing.T) {
	tests.DynSetInterfaceUnion(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceDifference(t *testing.T) {
	tests.DynSetInterfaceDifference(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceIsSuperset(t *testing.T) {
	tests.DynSetInterfaceIsSuperset(VectorToSetInterfaceFactory, t)
}

func TestVector_DynSetInterfaceIsSubset(t *testing.T) {
	tests.DynSetInterfaceIsSubset(VectorToSetInterfaceFactory, t)
}
