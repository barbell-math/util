// This package defines several widget types that allow logic for certian
// common operations to be passed arround as state.
package widgets

import "github.com/barbell-math/util/hash"

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
		// A function that should return true if the current value is less than
		// other.
		Lt(l *T, r *T) bool
	}

	// The interface that defines what it means to be a widget that can do
	// basic arithmetic. Implementations of this interface are expected to hold 
	// no state that pertains to the widget functions as the methods will not
	// be called in any predetermined order.
	ArithInterface[T any] interface {
		// An arith widget is a superset of the [WidgetInterface].
		PartialOrderInterface[T]
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
		PartialOrder[T, I]
	}
)

// Zeros the supplied value using the logic defined by the interface that was
// supplied as a generic type.
func (w *Base[T, I]) Zero(v *T) {
	w.iFace.Zero(v)
}

// Zeros the supplied value using the logic defined by the interface that was
// supplied as a generic type.
func (w *Arith[T, I]) Zero(v *T) {
	w.iFace.Zero(v)
}

// Compares the left (l) and right (r) values and returns true is l==r using the
// Eq function from the interface that was supplied as a generic type.
func (w *Base[T, I]) Eq(l *T, r *T) bool {
	return w.iFace.Eq(l, r)
}

// Compares the left (l) and right (r) values and returns true is l==r using the
// Eq function from the interface that was supplied as a generic type.
func (w *Arith[T, I]) Eq(l *T, r *T) bool {
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

// Compares the left (l) and right (r) values and returns true is l<r using the
// Lt function from the interface that was supplied as a generic type.
func (w *PartialOrder[T, I]) Lt(l *T, r *T) bool {
	return w.iFace.Lt(l, r)
}

// Compares the left (l) and right (r) values and returns true is l<r using the
// Lt function from the interface that was supplied as a generic type.
func (w *Arith[T, I]) Lt(l *T, r *T) bool {
	return w.iFace.Lt(l, r)
}

// Compares the left (l) and right (r) values and returns true is l<=r using the
// Lt and Eq functions from the interface that was supplied as a generic type.
func (w *PartialOrder[T, I]) Lte(l *T, r *T) bool {
	return w.iFace.Lt(l, r) || w.iFace.Eq(l, r)
}

// Compares the left (l) and right (r) values and returns true is l<=r using the
// Lt and Eq functions from the interface that was supplied as a generic type.
func (w *Arith[T, I]) Lte(l *T, r *T) bool {
	return w.iFace.Lt(l, r) || w.iFace.Eq(l, r)
}

// Compares the left (l) and right (r) values and returns true is l>r using the
// Lt and Eq functions from the interface that was supplied as a generic type.
func (w *PartialOrder[T, I]) Gt(l *T, r *T) bool {
	return !w.iFace.Lt(l, r) && !w.iFace.Eq(l, r)
}

// Compares the left (l) and right (r) values and returns true is l>r using the
// Lt and Eq functions from the interface that was supplied as a generic type.
func (w *Arith[T, I]) Gt(l *T, r *T) bool {
	return !w.iFace.Lt(l, r) && !w.iFace.Eq(l, r)
}

// Compares the left (l) and right (r) values and returns true is l>=r using the
// Lt function from the interface that was supplied as a generic type.
func (w *PartialOrder[T, I]) Gte(l *T, r *T) bool {
	return !w.iFace.Lt(l, r)
}

// Compares the left (l) and right (r) values and returns true is l>=r using the
// Lt function from the interface that was supplied as a generic type.
func (w *Arith[T, I]) Gte(l *T, r *T) bool {
	return !w.iFace.Lt(l, r)
}

// Generates a hash for the given value using the hash function from the interface
// that was supplied as a generic type.
func (w *Base[T, I]) Hash(v *T) hash.Hash {
	return w.iFace.Hash(v)
}

// Generates a hash for the given value using the hash function from the interface
// that was supplied as a generic type.
func (w *Arith[T, I]) Hash(v *T) hash.Hash {
	return w.iFace.Hash(v)
}

// Returns the "zero" value for the given widget type using the function from
// the interface that was supplied as a generic type.
func (w *Arith[T, I]) ZeroVal() T {
	return w.iFace.ZeroVal()
}

// Returns the "1" value for the given widget type using the function from the
// interface that was supplied as a generic type.
func (w *Arith[T, I]) UnitVal() T {
	return w.iFace.UnitVal()
}

// Negates the supplied value using the function from the interface that was
// supplied as a generic type.
func (w *Arith[T, I]) Neg(v *T) {
	w.iFace.Neg(v)
}

// Adds the supplied values using the function from the interface that was
// supplied as a generic type.
func (w *Arith[T, I]) Add(res *T, l *T, r *T) {
	w.iFace.Add(res, l, r)
}

// Subtracts the supplied values using the function from the interface that was
// supplied as a generic type.
func (w *Arith[T, I]) Sub(res *T, l *T, r *T) {
	w.iFace.Sub(res, l, r)
}

// Multiplies the supplied values using the function from the interface that was
// supplied as a generic type.
func (w *Arith[T, I]) Mul(res *T, l *T, r *T) {
	w.iFace.Mul(res, l, r)
}

// Divides the supplied values using the function from the interface that was
// supplied as a generic type.
func (w *Arith[T, I]) Div(res *T, l *T, r *T) {
	w.iFace.Div(res, l, r)
}

//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=bool
//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=byte
//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=uintptr
//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=string

//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=int
//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=int8
//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=int16
//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=int32
//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=int64

//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=uint
//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=uint8
//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=uint16
//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=uint32
//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=uint64

//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=float32
//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=float64

//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=complex64
//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=complex128

// This is a special case that is only allowed because the widget package itself
// relies on hash.Hash, making it so the hash.Hash package cannot implement the
// widget interface on itself, would create circular imports.
//go:generate ../bin/widgetInterfaceImpl -package=widgets -type=hash.Hash
