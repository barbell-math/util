package dataStruct

import (
	"testing"
)

func TestVectorDynVectorTypeInterface(t *testing.T) {
    v:=Vector[int](make(Vector[int], 0))
    dynVectorInterfaceTypeCheck[int](&v);
}
