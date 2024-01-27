package containers

import (
	"unsafe"

	"github.com/barbell-math/util/container/staticContainers"
)

// The type is used to determine which value is in a [Variant].
type VariantFlag int

const (
	A VariantFlag = iota	// A value
	B						// B value
)

// A variant type that can hold one of two types at a single time. The underlying
// value can be accessed through methods of the variant. Trying to access the
// wrong value will result in a panic unless a default value is supplied. The
// size of the struct will be equal to the largest type plus the size of the
// [VariantFlag].
type Variant[T any, U any] struct {
	data []byte
	aOrB VariantFlag
}

// The return type for these two functions has to be types.Variant because that
// is what the interface expects. The interface cannot use a specific return
// value because that would require the interface to import this package, creating
// circular imports.

// Sets the variant to hold value type A, initilized with the value passed to 
// the function. After calling this method the variant will panic if value
// type B is attempted to be accessed.
func (v Variant[T, U]) SetValA(newVal T) staticContainers.Variant[T, U] {
	var tmpA T
	var tmpB U
	size:=unsafe.Sizeof(tmpA)
	if bSize:=unsafe.Sizeof(tmpB); bSize>size {
		size=bSize
	}
	v.data=make([]byte,size)
	*(*T)(unsafe.Pointer(&v.data))=newVal
	v.aOrB = A
	return v
}

// Sets the variant to hold value type B, initilized with the value passed to 
// the function. After calling this method the variant will panic if value
// type A is attempted to be accessed.
func (v Variant[T, U]) SetValB(newVal U) staticContainers.Variant[T, U] {
	var tmpA T
	var tmpB U
	size:=unsafe.Sizeof(tmpA)
	if bSize:=unsafe.Sizeof(tmpB); bSize>size {
		size=bSize
	}
	v.data=make([]byte,size)
	*(*U)(unsafe.Pointer(&v.data))=newVal
	v.aOrB = B
	return v
}

// Returns true if the variant holds value A.
func (v Variant[T, U]) HasA() bool { return v.aOrB == A }
// Returns true if the variant holds value B.
func (v Variant[T, U]) HasB() bool { return v.aOrB == B }

// Attempts to return value A from the variant. Panics if the variant does not
// hold type A.
func (v Variant[T, U]) ValA() T {
	if v.aOrB!=A{
		panic("Variant does not contain type A!")
	}
	return *(*T)(unsafe.Pointer(&v.data))
}

// Attempts to return value B from the variant. Panics if the variant does not
// hold type B.
func (v Variant[T, U]) ValB() U {
	if v.aOrB!=B{
		panic("Variant does not contain type B!")
	}
	return *(*U)(unsafe.Pointer(&v.data))
}

// Attempts to return value A from the variant. If the variant does not hold
// type A then it will return the default value.
func (v Variant[T, U]) ValAOr(_default T) T {
	if v.aOrB == A {
		return v.ValA()
	}
	return _default
}

// Attempts to return value B from the variant. If the variant does not hold
// type B then it will return the default value.
func (v Variant[T, U]) ValBOr(_default U) U {
	if v.aOrB == B {
		return v.ValB()
	}
	return _default
}
