package view

import (
	"fmt"
	"testing"
	"unsafe"
)

func argPassing(v Vector[int]) {}

func TestView(t *testing.T){
	v:=View[int,int,ReadKeyedAddressable[int,int],NoCreate,NoUpdate,NoDelete]{}
	fmt.Println(unsafe.Sizeof(v))
	var b bool=v.Read.Contains(0)
	fmt.Println(b)

	vec:=Vector[int]{}
	vec.View=v
}
