package containers

// Code generated by ../../../bin/containerInterfaceTests - DO NOT EDIT.
import (
	"github.com/barbell-math/util/src/container/dynamicContainers"
	"github.com/barbell-math/util/src/container/tests"
	"testing"
)

func VectorToQueueInterfaceFactory(capacity int) dynamicContainers.Queue[int] {
	v := generateVector(capacity)
	var rv dynamicContainers.Queue[int] = &v
	return rv
}

func TestVector_DynQueueInterfaceSyncableInterface(t *testing.T) {
	tests.DynQueueInterfaceSyncableInterface(VectorToQueueInterfaceFactory, t)
}

func TestVector_DynQueueInterfaceAddressableInterface(t *testing.T) {
	tests.DynQueueInterfaceAddressableInterface(VectorToQueueInterfaceFactory, t)
}

func TestVector_DynQueueInterfaceLengthInterface(t *testing.T) {
	tests.DynQueueInterfaceLengthInterface(VectorToQueueInterfaceFactory, t)
}

func TestVector_DynQueueInterfaceCapacityInterface(t *testing.T) {
	tests.DynQueueInterfaceCapacityInterface(VectorToQueueInterfaceFactory, t)
}

func TestVector_DynQueueInterfaceClearInterface(t *testing.T) {
	tests.DynQueueInterfaceClearInterface(VectorToQueueInterfaceFactory, t)
}

func TestVector_DynQueueInterfaceFirstElemReadInterface(t *testing.T) {
	tests.DynQueueInterfaceFirstElemReadInterface(VectorToQueueInterfaceFactory, t)
}

func TestVector_DynQueueInterfaceFirstElemDeleteInterface(t *testing.T) {
	tests.DynQueueInterfaceFirstElemDeleteInterface(VectorToQueueInterfaceFactory, t)
}

func TestVector_DynQueueInterfaceLastElemWriteInterface(t *testing.T) {
	tests.DynQueueInterfaceLastElemWriteInterface(VectorToQueueInterfaceFactory, t)
}

func TestVector_ReadDynQueueInterface(t *testing.T) {
	tests.ReadDynQueueInterface(VectorToQueueInterfaceFactory, t)
}

func TestVector_WriteDynQueueInterface(t *testing.T) {
	tests.WriteDynQueueInterface(VectorToQueueInterfaceFactory, t)
}

func TestVector_DynQueueInterfaceInterface(t *testing.T) {
	tests.DynQueueInterfaceInterface(VectorToQueueInterfaceFactory, t)
}

func TestVector_DynQueueInterfaceStaticCapacityInterface(t *testing.T) {
	tests.DynQueueInterfaceStaticCapacityInterface(VectorToQueueInterfaceFactory, t)
}

func TestVector_DynQueueInterfaceClear(t *testing.T) {
	tests.DynQueueInterfaceClear(VectorToQueueInterfaceFactory, t)
}

func TestVector_DynQueueInterfacePeekPntrFront(t *testing.T) {
	tests.DynQueueInterfacePeekPntrFront(VectorToQueueInterfaceFactory, t)
}

func TestVector_DynQueueInterfacePeekFront(t *testing.T) {
	tests.DynQueueInterfacePeekFront(VectorToQueueInterfaceFactory, t)
}

func TestVector_DynQueueInterfacePopFront(t *testing.T) {
	tests.DynQueueInterfacePopFront(VectorToQueueInterfaceFactory, t)
}

func TestVector_DynQueueInterfacePushBack(t *testing.T) {
	tests.DynQueueInterfacePushBack(VectorToQueueInterfaceFactory, t)
}

func TestVector_DynQueueInterfaceForcePushBack(t *testing.T) {
	tests.DynQueueInterfaceForcePushBack(VectorToQueueInterfaceFactory, t)
}
