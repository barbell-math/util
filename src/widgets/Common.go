// This package defines several widget types that allow logic for certian
// common operations to be passed arround as state.
package widgets

import "github.com/barbell-math/util/src/hash"

type (
	// The interface that defines what it means to be a widget. Implementations
	// of this interface are expected to hold no state that pertains to the
	// widget functions as the methods shown below will not be called in any
	// predetermined order.
	BaseInterface[T any] interface {
		// A function that should return true if the current value equals other.
		Eq(l *T, r *T) bool
		// Returns a hash value that represent the currently wrapped value.
		Hash(v *T) hash.Hash
		// Zero's the supplied value. Equivalent to a destructor except Go does
		// not have manual memory management so this is mainly just to prevent
		// things like dangling pointers.
		Zero(v *T)
	}

	// The interface that defines what it means to be a widget that can be
	// compared to other widgets. Implementations of this interface are expected
	// to hold no state that pertains to the widget functions as the methods
	// shown below will not be called in any predetermined order.
	PartialOrderInterface[T any] interface {
		BaseInterface[T]
		// A function that should return true if l<r.
		Lt(l *T, r *T) bool
	}

	// The interface that defines what it means to be a widget that can do
	// basic arithmetic. Implementations of this interface are expected to hold
	// no state that pertains to the widget functions as the methods will not
	// be called in any predetermined order.
	ArithInterface[T any] interface {
		BaseInterface[T]
		// Returns the value that represents "zero" for the underlying type.
		ZeroVal() T
		// Returns the value that represent "1" for the underlying type.
		UnitVal() T
		// Negates the value that is supplied to it.
		Neg(v *T)
		// Adds l and r and places the results in res. No uniqueness guarantees
		// are placed on res, l, and r. They may all be the same value. The
		// implementation of this interface needs to recognize this.
		Add(res *T, l *T, r *T)
		// Subtracts l and r and places the results in res. No uniqueness
		// guarantees are placed on res, l, and r. They may all be the same
		// value. The implementation of this interface needs to recognize this.
		Sub(res *T, l *T, r *T)
		// Multiplies l and r and places the results in res. No uniqueness
		// guarantees are placed on res, l, and r. They may all be the same
		// value. The implementation of this interface needs to recognize this.
		Mul(res *T, l *T, r *T)
		// Divides l and r and places the results in res. No uniqueness
		// guarantees are placed on res, l, and r. They may all be the same
		// value. The implementation of this interface needs to recognize this.
		Div(res *T, l *T, r *T)
	}

	// The interface that defines what it means to be a widget that supports
	// both the arithmetic and partial order interface. Implementations of this
	// interface are expected to hold no state that pertains to the widget
	// functions as the methods will not be called in any predetermined order.
	PartialOrderArithInterface[T any] interface {
		ArithInterface[T]
		PartialOrderInterface[T]
	}

	// A concrete base widget. Internally, this struct will create an interface
	// value of type [BaseInterface] that points to nil data. All methods on
	// widget are then very thin pass through functions that call the needed
	// methods on the interface value with the supplied values.
	Base[T any, I BaseInterface[T]] struct {
		iFace I
	}

	// A concrete partial order widget. Internally, this struct will create an
	// interface value of type [PartialOrderInterface] that points to nil data.
	// All methods on widget are then very thin pass through functions that call
	// the needed methods on the interface value with the supplied values.
	PartialOrder[T any, I PartialOrderInterface[T]] struct {
		Base[T, I]
	}

	// A concrete arithmitic widget. Internally, this struct will create an
	// interface value of type [ArithInterface] that points to nil data. All
	// methods on widget are then very thin pass through functions that call the
	// needed methods on the interface value with the supplied values.
	Arith[T any, I ArithInterface[T]] struct {
		Base[T, I]
	}

	// A concrete arithmitic widget. Internally, this struct will create an
	// interface value of type [PartialOrderArith] that points to nil data. All
	// methods on widget are then very thin pass through functions that call the
	// needed methods on the interface value with the supplied values.
	PartialOrderArith[T any, I PartialOrderArithInterface[T]] struct {
		Base[T, I]
	}
)

// Zeros the supplied value using the logic defined by the interface that was
// supplied as a generic type.
func (w *Base[T, I]) Zero(v *T) {
	w.iFace.Zero(v)
}

// Compares the left (l) and right (r) values and returns true is l==r using the
// Eq function from the interface that was supplied as a generic type.
func (w *Base[T, I]) Eq(l *T, r *T) bool {
	return w.iFace.Eq(l, r)
}

// Compares the left (l) and right (r) values and returns true is l!=r using the
// Eq function from the interface that was supplied as a generic type.
func (w *Base[T, I]) Neq(l *T, r *T) bool {
	return !w.iFace.Eq(l, r)
}

// Compares the left (l) and right (r) values and returns true is l!=r using the
// Eq function from the interface that was supplied as a generic type.
func (w *Arith[T, I]) Neq(l *T, r *T) bool {
	return !w.iFace.Eq(l, r)
}

// Compares the left (l) and right (r) values and returns true is l!=r using the
// Eq function from the interface that was supplied as a generic type.
func (w *PartialOrderArith[T, I]) Neq(l *T, r *T) bool {
	return !w.iFace.Eq(l, r)
}

// Compares the left (l) and right (r) values and returns true is l<r using the
// Lt function from the interface that was supplied as a generic type.
func (w *PartialOrder[T, I]) Lt(l *T, r *T) bool {
	return w.iFace.Lt(l, r)
}

// Compares the left (l) and right (r) values and returns true is l<r using the
// Lt function from the interface that was supplied as a generic type.
func (w *PartialOrderArith[T, I]) Lt(l *T, r *T) bool {
	return w.iFace.Lt(l, r)
}

// Compares the left (l) and right (r) values and returns true is l<=r using the
// Lt and Eq functions from the interface that was supplied as a generic type.
func (w *PartialOrder[T, I]) Lte(l *T, r *T) bool {
	return w.iFace.Lt(l, r) || w.iFace.Eq(l, r)
}

// Compares the left (l) and right (r) values and returns true is l<=r using the
// Lt and Eq functions from the interface that was supplied as a generic type.
func (w *PartialOrderArith[T, I]) Lte(l *T, r *T) bool {
	return w.iFace.Lt(l, r) || w.iFace.Eq(l, r)
}

// Compares the left (l) and right (r) values and returns true is l>r using the
// Lt and Eq functions from the interface that was supplied as a generic type.
func (w *PartialOrder[T, I]) Gt(l *T, r *T) bool {
	return !w.iFace.Lt(l, r) && !w.iFace.Eq(l, r)
}

// Compares the left (l) and right (r) values and returns true is l>r using the
// Lt and Eq functions from the interface that was supplied as a generic type.
func (w *PartialOrderArith[T, I]) Gt(l *T, r *T) bool {
	return !w.iFace.Lt(l, r) && !w.iFace.Eq(l, r)
}

// Compares the left (l) and right (r) values and returns true is l>=r using the
// Lt function from the interface that was supplied as a generic type.
func (w *PartialOrder[T, I]) Gte(l *T, r *T) bool {
	return !w.iFace.Lt(l, r)
}

// Compares the left (l) and right (r) values and returns true is l>=r using the
// Lt function from the interface that was supplied as a generic type.
func (w *PartialOrderArith[T, I]) Gte(l *T, r *T) bool {
	return !w.iFace.Lt(l, r)
}

// Generates a hash for the given value using the hash function from the interface
// that was supplied as a generic type.
func (w *Base[T, I]) Hash(v *T) hash.Hash {
	return w.iFace.Hash(v)
}

// Returns the "zero" value for the given widget type using the function from
// the interface that was supplied as a generic type.
func (w *Arith[T, I]) ZeroVal() T {
	return w.iFace.ZeroVal()
}

// Returns the "zero" value for the given widget type using the function from
// the interface that was supplied as a generic type.
func (w *PartialOrderArith[T, I]) ZeroVal() T {
	return w.iFace.ZeroVal()
}

// Returns the "1" value for the given widget type using the function from the
// interface that was supplied as a generic type.
func (w *Arith[T, I]) UnitVal() T {
	return w.iFace.UnitVal()
}

// Returns the "1" value for the given widget type using the function from the
// interface that was supplied as a generic type.
func (w *PartialOrderArith[T, I]) UnitVal() T {
	return w.iFace.UnitVal()
}

// Negates the supplied value using the function from the interface that was
// supplied as a generic type.
func (w *Arith[T, I]) Neg(v *T) {
	w.iFace.Neg(v)
}

// Negates the supplied value using the function from the interface that was
// supplied as a generic type.
func (w *PartialOrderArith[T, I]) Neg(v *T) {
	w.iFace.Neg(v)
}

// Adds the supplied values using the function from the interface that was
// supplied as a generic type.
func (w *Arith[T, I]) Add(res *T, l *T, r *T) {
	w.iFace.Add(res, l, r)
}

// Adds the supplied values using the function from the interface that was
// supplied as a generic type.
func (w *PartialOrderArith[T, I]) Add(res *T, l *T, r *T) {
	w.iFace.Add(res, l, r)
}

// Subtracts the supplied values using the function from the interface that was
// supplied as a generic type.
func (w *Arith[T, I]) Sub(res *T, l *T, r *T) {
	w.iFace.Sub(res, l, r)
}

// Subtracts the supplied values using the function from the interface that was
// supplied as a generic type.
func (w *PartialOrderArith[T, I]) Sub(res *T, l *T, r *T) {
	w.iFace.Sub(res, l, r)
}

// Multiplies the supplied values using the function from the interface that was
// supplied as a generic type.
func (w *Arith[T, I]) Mul(res *T, l *T, r *T) {
	w.iFace.Mul(res, l, r)
}

// Multiplies the supplied values using the function from the interface that was
// supplied as a generic type.
func (w *PartialOrderArith[T, I]) Mul(res *T, l *T, r *T) {
	w.iFace.Mul(res, l, r)
}

// Divides the supplied values using the function from the interface that was
// supplied as a generic type.
func (w *Arith[T, I]) Div(res *T, l *T, r *T) {
	w.iFace.Div(res, l, r)
}

// Divides the supplied values using the function from the interface that was
// supplied as a generic type.
func (w *PartialOrderArith[T, I]) Div(res *T, l *T, r *T) {
	w.iFace.Div(res, l, r)
}

//go:generate ../../bin/widgetInterfaceImpl -widgetType=Base -type=bool
//go:generate ../../bin/widgetInterfaceImpl -widgetType=PartialOrderArith -type=byte
//go:generate ../../bin/widgetInterfaceImpl -widgetType=PartialOrderArith -type=uintptr
//go:generate ../../bin/widgetInterfaceImpl -widgetType=PartialOrder -type=string

//go:generate ../../bin/widgetInterfaceImpl -widgetType=PartialOrderArith -type=int
//go:generate ../../bin/widgetInterfaceImpl -widgetType=PartialOrderArith -type=int8
//go:generate ../../bin/widgetInterfaceImpl -widgetType=PartialOrderArith -type=int16
//go:generate ../../bin/widgetInterfaceImpl -widgetType=PartialOrderArith -type=int32
//go:generate ../../bin/widgetInterfaceImpl -widgetType=PartialOrderArith -type=int64

//go:generate ../../bin/widgetInterfaceImpl -widgetType=PartialOrderArith -type=uint
//go:generate ../../bin/widgetInterfaceImpl -widgetType=PartialOrderArith -type=uint8
//go:generate ../../bin/widgetInterfaceImpl -widgetType=PartialOrderArith -type=uint16
//go:generate ../../bin/widgetInterfaceImpl -widgetType=PartialOrderArith -type=uint32
//go:generate ../../bin/widgetInterfaceImpl -widgetType=PartialOrderArith -type=uint64

//go:generate ../../bin/widgetInterfaceImpl -widgetType=PartialOrderArith -type=float32
//go:generate ../../bin/widgetInterfaceImpl -widgetType=PartialOrderArith -type=float64

//go:generate ../../bin/widgetInterfaceImpl -widgetType=Arith -type=complex64
//go:generate ../../bin/widgetInterfaceImpl -widgetType=Arith -type=complex128

// This is a special case that is only allowed because the widget package itself
// relies on hash.Hash, making it so the hash.Hash package cannot implement the
// widget interface on itself (would create circular imports).
//go:generate ../../bin/widgetInterfaceImpl -widgetType=PartialOrderArith -type=hash.Hash
