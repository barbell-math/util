package containers

// Code generated by ../../../bin/containerInterfaceTests - DO NOT EDIT.
import (
	"github.com/barbell-math/util/src/container/dynamicContainers"
	"github.com/barbell-math/util/src/container/tests"
	"testing"
)

func VectorToVectorInterfaceFactory(capacity int) dynamicContainers.Vector[int] {
	v := generateVector(capacity)
	var rv dynamicContainers.Vector[int] = &v
	return rv
}

func TestVector_DynVectorInterfaceSyncableInterface(t *testing.T) {
	tests.DynVectorInterfaceSyncableInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceAddressableInterface(t *testing.T) {
	tests.DynVectorInterfaceAddressableInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceLengthInterface(t *testing.T) {
	tests.DynVectorInterfaceLengthInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceCapacityInterface(t *testing.T) {
	tests.DynVectorInterfaceCapacityInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceClearInterface(t *testing.T) {
	tests.DynVectorInterfaceClearInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceWriteOpsInterface(t *testing.T) {
	tests.DynVectorInterfaceWriteOpsInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceWriteKeyedOpsInterface(t *testing.T) {
	tests.DynVectorInterfaceWriteKeyedOpsInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceWriteKeyedSequentialOpsInterface(t *testing.T) {
	tests.DynVectorInterfaceWriteKeyedSequentialOpsInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceWriteDynKeyedOpsInterface(t *testing.T) {
	tests.DynVectorInterfaceWriteDynKeyedOpsInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceReadOpsInterface(t *testing.T) {
	tests.DynVectorInterfaceReadOpsInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceReadKeyedOpsInterface(t *testing.T) {
	tests.DynVectorInterfaceReadKeyedOpsInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceDeleteOpsInterface(t *testing.T) {
	tests.DynVectorInterfaceDeleteOpsInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceDeleteKeyedOpsInterface(t *testing.T) {
	tests.DynVectorInterfaceDeleteKeyedOpsInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceDeleteSequentialOpsInterface(t *testing.T) {
	tests.DynVectorInterfaceDeleteSequentialOpsInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceDeleteKeyedSequentialOpsInterface(t *testing.T) {
	tests.DynVectorInterfaceDeleteKeyedSequentialOpsInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_ReadDynVectorInterface(t *testing.T) {
	tests.ReadDynVectorInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_WriteDynVectorInterface(t *testing.T) {
	tests.WriteDynVectorInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceInterface(t *testing.T) {
	tests.DynVectorInterfaceInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceStaticCapacityInterface(t *testing.T) {
	tests.DynVectorInterfaceStaticCapacityInterface(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceGet(t *testing.T) {
	tests.DynVectorInterfaceGet(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceGetPntr(t *testing.T) {
	tests.DynVectorInterfaceGetPntr(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceContains(t *testing.T) {
	tests.DynVectorInterfaceContains(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceContainsPntr(t *testing.T) {
	tests.DynVectorInterfaceContainsPntr(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceKeyOf(t *testing.T) {
	tests.DynVectorInterfaceKeyOf(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceKeyOfPntr(t *testing.T) {
	tests.DynVectorInterfaceKeyOfPntr(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceSet(t *testing.T) {
	tests.DynVectorInterfaceSet(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceSetSequential(t *testing.T) {
	tests.DynVectorInterfaceSetSequential(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceAppend(t *testing.T) {
	tests.DynVectorInterfaceAppend(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceInsert(t *testing.T) {
	tests.DynVectorInterfaceInsert(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceInsertSequential(t *testing.T) {
	tests.DynVectorInterfaceInsertSequential(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfacePopSequential(t *testing.T) {
	tests.DynVectorInterfacePopSequential(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfacePop(t *testing.T) {
	tests.DynVectorInterfacePop(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfacePopPntr(t *testing.T) {
	tests.DynVectorInterfacePopPntr(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceDelete(t *testing.T) {
	tests.DynVectorInterfaceDelete(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceDeleteSequential(t *testing.T) {
	tests.DynVectorInterfaceDeleteSequential(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceClear(t *testing.T) {
	tests.DynVectorInterfaceClear(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceVals(t *testing.T) {
	tests.DynVectorInterfaceVals(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceValPntrs(t *testing.T) {
	tests.DynVectorInterfaceValPntrs(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceKeys(t *testing.T) {
	tests.DynVectorInterfaceKeys(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceKeyedEq(t *testing.T) {
	tests.DynVectorInterfaceKeyedEq(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceUnorderedEq(t *testing.T) {
	tests.DynVectorInterfaceUnorderedEq(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceIntersection(t *testing.T) {
	tests.DynVectorInterfaceIntersection(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceUnion(t *testing.T) {
	tests.DynVectorInterfaceUnion(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceDifference(t *testing.T) {
	tests.DynVectorInterfaceDifference(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceIsSuperset(t *testing.T) {
	tests.DynVectorInterfaceIsSuperset(VectorToVectorInterfaceFactory, t)
}

func TestVector_DynVectorInterfaceIsSubset(t *testing.T) {
	tests.DynVectorInterfaceIsSubset(VectorToVectorInterfaceFactory, t)
}