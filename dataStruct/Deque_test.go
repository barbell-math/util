package dataStruct

import "testing"


func TestDequeDynDequeTypeInterface(t *testing.T) {
    v,_:=NewDeque[int](5)
    v2,_:=NewSyncedDeque[int](5)
    dynDequeInterfaceTypeCheck[int](&v);
    dynDequeInterfaceTypeCheck[int](&v2);
}

func TestDequeDynVectorTypeInterface(t *testing.T) {
    v,_:=NewDeque[int](5)
    v2,_:=NewSyncedDeque[int](5)
    dynVectorInterfaceTypeCheck[int](&v);
    dynVectorInterfaceTypeCheck[int](&v2);
}
