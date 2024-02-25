// This package serves to implement the comparison and hash operations necessary
// to put built-in types into the containers defined in [containers]. All of the
// types in this package will implement the [containers.Widget] interface.
// Basically all of the code in this package is generated with go:generate comments.
package widgets

import "github.com/barbell-math/util/algo/hash"


type (
    // The interface that defines what it means to be a widget. This interface
    // is used by the containers in the [containers] package when performing
    // certian operations. Implementations of this interface are expected to
    // hold no state as the methods shown below will not be called in any
    // predetermined order.
    WidgetInterface[T any] interface {
        // A function that should return true if the current value equals other.
        Eq(l *T, r *T) bool
        // A function that should return true if the current value is less than other.
        Lt(l *T, r *T) bool
        // Returns a hash value that represent the currently wrapped value.
        Hash(v *T) hash.Hash
        // Zero's the supplied value. Equivalent to a destructor except Go does
        // not have manual memory management so this is mainly just to prenent
        // things like dangling pointers.
        Zero(v *T)
    }

    // The interface that defines what it means to be a widget that can do
    // baic arithmitic. This interface is used by the containers in the 
    // [container] package when performing certian operations. Implementations
    // of this interface are expected to hold no state as the methods will not
    // be called in any predetermined order.
    ArithWidgetInterface[T any] interface {
        // An arith widget is a superset of the [WidgetInterface].
        WidgetInterface[T]
        // Returns the value that represents "zero" for the underlying type.
        ZeroVal() T
        // Returns the value that represent "1" for the underlying type.
        UnitVal() T
        // Negates the value that is supplied to it.
        Neg(v *T)
        // Adds l and r and places the results in res. No uniqueness guarintees
        // are placed on res, l, and r. They may all be the same value. The
        // implementation of this interface needs to recognize this.
        Add(res *T, l *T, r *T)
        // Subtracts l and r and places the results in res. No uniqueness 
        // guarintees are placed on res, l, and r. They may all be the same 
        // value. The implementation of this interface needs to recognize this.
        Sub(res *T, l *T, r *T)
        // Multiplies l and r and places the results in res. No uniqueness 
        // guarintees are placed on res, l, and r. They may all be the same 
        // value. The implementation of this interface needs to recognize this.
        Mul(res *T, l *T, r *T)
        // Divides l and r and places the results in res. No uniqueness 
        // guarintees are placed on res, l, and r. They may all be the same 
        // value. The implementation of this interface needs to recognize this.
        Div(res *T, l *T, r *T)
    }

    // The base widget implementation that all the containers in the [containers]
    // package use as a type restriction. This type must be instantiaed with the
    // [NewWidget] function; zero valued Widget's are not valid and will result
    // in nil pointer errors. Internally, this struct will create an interface
    // value of type [WidgetInterface] that points to nil data. All methods on
    // widget are then very thin pass through functions that call the needed
    // methods on the interface value with the supplied values.
    Widget[T any, I WidgetInterface[T]] struct {
        iFace I
    }

    // An arithmitic aware version of the base [Widget] type that some of the
    // containers in the [container] package use as a type restriction. This
    // type must be instantiated with the [NewArithWidget] function; zero valued
    // ArithWidget's are not value and will result in nil pointer errors.
    // Internally, this struct will create an interface value of type 
    // [ArithWidgetInterface] that points to nil data. All methods on
    // widget are then very thin pass through functions that call the needed
    // methods on the interface value with the supplied values.
    ArithWidget[T any, I ArithWidgetInterface[T]] struct {
        iFace I
    }
)

// Creates a new widget and sets its internal state so that it is valid and can
// be used without error.
func NewWidget[T any, I WidgetInterface[T]]() Widget[T,I] {
    var iFaceImpl I
    return Widget[T, I]{iFace: iFaceImpl}
}

// Creates a new arithmitic aware widget and sets its internal state so that it
// is valid and can be used without error.
func NewArithWidget[T any, I ArithWidgetInterface[T]]() ArithWidget[T,I] {
    var iFaceImpl I
    return ArithWidget[T,I]{iFace: iFaceImpl}
}

// Zeros the supplied value using the logic defined by the interface that was
// supplied as a generic type.
func (w *Widget[T, I])Zero(v *T) {
    w.iFace.Zero(v)
}
// Zeros the supplied value using the logic defined by the interface that was
// supplied as a generic type.
func (w *ArithWidget[T, I])Zero(v *T) {
    w.iFace.Zero(v)
}

// Compares the left (l) and right (r) values and returns true is l==r using the
// Eq function from the interface that was supplied as a generic type.
func (w *Widget[T, I])Eq(l *T, r *T) bool {
    return w.iFace.Eq(l,r)
}
// Compares the left (l) and right (r) values and returns true is l==r using the
// Eq function from the interface that was supplied as a generic type.
func (w *ArithWidget[T, I])Eq(l *T, r *T) bool {
    return w.iFace.Eq(l,r)
}

// Compares the left (l) and right (r) values and returns true is l!=r using the
// Eq function from the interface that was supplied as a generic type.
func (w *Widget[T, I])Neq(l *T, r *T) bool {
    return !w.iFace.Eq(l,r)
}
// Compares the left (l) and right (r) values and returns true is l!=r using the
// Eq function from the interface that was supplied as a generic type.
func (w *ArithWidget[T, I])Neq(l *T, r *T) bool {
    return !w.iFace.Eq(l,r)
}

// Compares the left (l) and right (r) values and returns true is l<r using the
// Lt function from the interface that was supplied as a generic type.
func (w *Widget[T, I])Lt(l *T, r *T) bool {
    return w.iFace.Lt(l,r)
}
// Compares the left (l) and right (r) values and returns true is l<r using the
// Lt function from the interface that was supplied as a generic type.
func (w *ArithWidget[T, I])Lt(l *T, r *T) bool {
    return w.iFace.Lt(l,r)
}

// Compares the left (l) and right (r) values and returns true is l<=r using the
// Lt and Eq functions from the interface that was supplied as a generic type.
func (w *Widget[T, I])Lte(l *T, r *T) bool {
    return w.iFace.Lt(l,r) || w.iFace.Eq(l,r)
}
// Compares the left (l) and right (r) values and returns true is l<=r using the
// Lt and Eq functions from the interface that was supplied as a generic type.
func (w *ArithWidget[T, I])Lte(l *T, r *T) bool {
    return w.iFace.Lt(l,r) || w.iFace.Eq(l,r)
}

// Compares the left (l) and right (r) values and returns true is l>r using the
// Lt and Eq functions from the interface that was supplied as a generic type.
func (w *Widget[T, I])Gt(l *T, r *T) bool {
    return !w.iFace.Lt(l,r) && !w.iFace.Eq(l,r)
}
// Compares the left (l) and right (r) values and returns true is l>r using the
// Lt and Eq functions from the interface that was supplied as a generic type.
func (w *ArithWidget[T, I])Gt(l *T, r *T) bool {
    return !w.iFace.Lt(l,r) && !w.iFace.Eq(l,r)
}

// Compares the left (l) and right (r) values and returns true is l>=r using the
// Lt function from the interface that was supplied as a generic type.
func (w *Widget[T, I])Gte(l *T, r *T) bool {
    return !w.iFace.Lt(l,r)
}
// Compares the left (l) and right (r) values and returns true is l>=r using the
// Lt function from the interface that was supplied as a generic type.
func (w *ArithWidget[T, I])Gte(l *T, r *T) bool {
    return !w.iFace.Lt(l,r)
}

// Generates a hash for the given value using the hash function from the interface
// that was supplied as a generic type.
func (w *Widget[T, I])Hash(v *T) hash.Hash {
    return w.iFace.Hash(v)
}
// Generates a hash for the given value using the hash function from the interface
// that was supplied as a generic type.
func (w *ArithWidget[T, I])Hash(v *T) hash.Hash {
    return w.iFace.Hash(v)
}

// Returns the "zero" value for the given widget type using the function from
// the interface that was supplied as a generic type.
func (w *ArithWidget[T, I])ZeroVal() T {
    return w.iFace.ZeroVal()
}

// Returns the "1" value for the given widget type using the function from the 
// interface that was supplied as a generic type.
func (w *ArithWidget[T, I])UnitVal() T {
    return w.iFace.UnitVal()
}

// Negates the supplied value using the function from the interface that was 
// supplied as a generic type.
func (w *ArithWidget[T, I])Neg(v *T) {
    w.iFace.Neg(v)
}

// Adds the supplied values using the function from the interface that was 
// supplied as a generic type.
func (w *ArithWidget[T, I])Add(res *T, l *T, r *T) {
    w.iFace.Add(res,l,r)
}

// Subtracts the supplied values using the function from the interface that was 
// supplied as a generic type.
func (w *ArithWidget[T, I])Sub(res *T, l *T, r *T) {
    w.iFace.Sub(res,l,r)
}

// Multiplies the supplied values using the function from the interface that was 
// supplied as a generic type.
func (w *ArithWidget[T, I])Mul(res *T, l *T, r *T) {
    w.iFace.Mul(res,l,r)
}

// Divides the supplied values using the function from the interface that was 
// supplied as a generic type.
func (w *ArithWidget[T, I])Div(res *T, l *T, r *T) {
    w.iFace.Div(res,l,r)
}

//go:generate go run widgetInterfaceImpl.go -package=widgets -type=byte

//go:generate go run widgetInterfaceImpl.go -package=widgets -type=int
//go:generate go run widgetInterfaceImpl.go -package=widgets -type=int8
//go:generate go run widgetInterfaceImpl.go -package=widgets -type=int16
//go:generate go run widgetInterfaceImpl.go -package=widgets -type=int32
//go:generate go run widgetInterfaceImpl.go -package=widgets -type=int64

//go:generate go run widgetInterfaceImpl.go -package=widgets -type=uint
//go:generate go run widgetInterfaceImpl.go -package=widgets -type=uint8
//go:generate go run widgetInterfaceImpl.go -package=widgets -type=uint16
//go:generate go run widgetInterfaceImpl.go -package=widgets -type=uint32
//go:generate go run widgetInterfaceImpl.go -package=widgets -type=uint64

//go:generate go run widgetInterfaceImpl.go -package=widgets -type=float32
//go:generate go run widgetInterfaceImpl.go -package=widgets -type=float64

//go:generate go run widgetInterfaceImpl.go -package=widgets -type=string
