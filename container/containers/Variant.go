package containers

import (
	"reflect"
	"unsafe"

	"github.com/barbell-math/util/container/staticContainers"
)


type (
	// A variant type that can hold one of two types at a single time. The 
	// underlying value can be accessed through methods of the variant. Trying 
	// to access the wrong value will result in a panic unless a default value 
	// is supplied. The size of the struct will be equal to the largest type 
	// plus one byte which is used to specify which value is in the variant.
	Variant[A any, B any] struct {
		data []byte
	}

	variantFlag byte
)

const (
	aVal variantFlag = iota
	bVal
)

func (v *Variant[A,B])initData() {
	if len(v.data)==0 {
		v.data=make([]byte,v.dataSize()+1)
	}
}

func (v Variant[A,B])dataSize() uintptr {
	size:=reflect.TypeOf((*A)(nil)).Elem().Size()
	if bSize:=reflect.TypeOf((*B)(nil)).Elem().Size(); bSize>size {
		size=bSize
	}
	return size
}

func (v Variant[A,B])dataStart() unsafe.Pointer {
	return unsafe.Pointer(&v.data[1])
}

// The return type for these two functions has to be types.Variant because that
// is what the interface expects. The interface cannot use a specific return
// value because that would require the interface to import this package, 
// creating circular imports.

// Sets the variant to hold value type A, initilized with the value passed to 
// the function. After calling this method the variant will panic if value
// type B is attempted to be accessed.
func (v Variant[A,B]) SetValA(newVal A) staticContainers.Variant[A,B] {
	v.initData()
	v.data[0]=byte(aVal)
	*(*A)(v.dataStart())=newVal
	return v
}

// Sets the variant to hold value type B, initilized with the value passed to 
// the function. After calling this method the variant will panic if value
// type A is attempted to be accessed.
func (v Variant[A,B]) SetValB(newVal B) staticContainers.Variant[A,B] {
	v.initData()
	v.data[0]=byte(bVal)
	*(*B)(v.dataStart())=newVal
	return v
}

// Returns true if the variant holds value A.
func (v Variant[A,B]) HasA() bool { return v.data[0]==byte(aVal) }
// Returns true if the variant holds value B.
func (v Variant[A,B]) HasB() bool { return v.data[0]==byte(bVal) }

// Attempts to return value A from the variant. Panics if the variant does not
// hold type A.
func (v Variant[A,B]) ValA() A {
	if !v.HasA() {
		panic("Variant does not contain type A!")
	}
	return *(*A)(v.dataStart())
}

// Attempts to return value B from the variant. Panics if the variant does not
// hold type B.
func (v Variant[A,B]) ValB() B {
	if !v.HasB() {
		panic("Variant does not contain type B!")
	}
	return *(*B)(v.dataStart())
}

// Attempts to return value A from the variant. If the variant does not hold
// type A then it will return the default value.
func (v Variant[A,B]) ValAOr(_default A) A {
	if v.HasA() {
		return v.ValA()
	}
	return _default
}

// Attempts to return value B from the variant. If the variant does not hold
// type B then it will return the default value.
func (v Variant[A,B]) ValBOr(_default B) B {
	if v.HasB() {
		return v.ValB()
	}
	return _default
}
