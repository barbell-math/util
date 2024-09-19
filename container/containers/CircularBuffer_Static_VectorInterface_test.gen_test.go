package containers

// Code generated by ../../bin/containerInterfaceTests - DO NOT EDIT.
import (
	"github.com/barbell-math/util/container/staticContainers"
	"github.com/barbell-math/util/container/tests"
	"testing"
)

func CircularBufferToVectorInterfaceFactory(capacity int) staticContainers.Vector[int] {
	v := generateCircularBuffer(capacity)
	var rv staticContainers.Vector[int] = &v
	return rv
}

func TestCircularBuffer_StaticVectorInterfaceStaticCapacity(t *testing.T) {
	tests.StaticVectorInterfaceStaticCapacity(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceAddressableInterface(t *testing.T) {
	tests.StaticVectorInterfaceAddressableInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceLengthInterface(t *testing.T) {
	tests.StaticVectorInterfaceLengthInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceCapacityInterface(t *testing.T) {
	tests.StaticVectorInterfaceCapacityInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceClearInterface(t *testing.T) {
	tests.StaticVectorInterfaceClearInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceWriteOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceWriteOpsInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceWriteKeyedOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceWriteKeyedOpsInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceWriteKeyedSequentialOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceWriteKeyedSequentialOpsInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceWriteDynKeyedOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceWriteDynKeyedOpsInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceReadOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceReadOpsInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceReadKeyedOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceReadKeyedOpsInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceDeleteOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceDeleteOpsInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceDeleteKeyedOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceDeleteKeyedOpsInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceDeleteSequentialOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceDeleteSequentialOpsInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceDeleteKeyedSequentialOpsInterface(t *testing.T) {
	tests.StaticVectorInterfaceDeleteKeyedSequentialOpsInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_ReadStaticVectorInterface(t *testing.T) {
	tests.ReadStaticVectorInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_WriteStaticVectorInterface(t *testing.T) {
	tests.WriteStaticVectorInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceInterface(t *testing.T) {
	tests.StaticVectorInterfaceInterface(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceGet(t *testing.T) {
	tests.StaticVectorInterfaceGet(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceGetPntr(t *testing.T) {
	tests.StaticVectorInterfaceGetPntr(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceContains(t *testing.T) {
	tests.StaticVectorInterfaceContains(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceContainsPntr(t *testing.T) {
	tests.StaticVectorInterfaceContainsPntr(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceKeyOf(t *testing.T) {
	tests.StaticVectorInterfaceKeyOf(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceKeyOfPntr(t *testing.T) {
	tests.StaticVectorInterfaceKeyOfPntr(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceSet(t *testing.T) {
	tests.StaticVectorInterfaceSet(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceSetSequential(t *testing.T) {
	tests.StaticVectorInterfaceSetSequential(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceAppend(t *testing.T) {
	tests.StaticVectorInterfaceAppend(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceInsert(t *testing.T) {
	tests.StaticVectorInterfaceInsert(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceInsertSequential(t *testing.T) {
	tests.StaticVectorInterfaceInsertSequential(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfacePopSequential(t *testing.T) {
	tests.StaticVectorInterfacePopSequential(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfacePop(t *testing.T) {
	tests.StaticVectorInterfacePop(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfacePopPntr(t *testing.T) {
	tests.StaticVectorInterfacePopPntr(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceDelete(t *testing.T) {
	tests.StaticVectorInterfaceDelete(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceDeleteSequential(t *testing.T) {
	tests.StaticVectorInterfaceDeleteSequential(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceClear(t *testing.T) {
	tests.StaticVectorInterfaceClear(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceVals(t *testing.T) {
	tests.StaticVectorInterfaceVals(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceValPntrs(t *testing.T) {
	tests.StaticVectorInterfaceValPntrs(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceKeys(t *testing.T) {
	tests.StaticVectorInterfaceKeys(CircularBufferToVectorInterfaceFactory, t)
}

func TestCircularBuffer_StaticVectorInterfaceKeyedEq(t *testing.T) {
	tests.StaticVectorInterfaceKeyedEq(CircularBufferToVectorInterfaceFactory, t)
}