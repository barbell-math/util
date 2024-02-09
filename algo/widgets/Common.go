// This package serves to implement the comparison and hash operations necessary
// to put built-in types into the containers defined in [containers]. All of the
// types in this package will implement the [containers.Widget] interface.
// Basically all of the code in this package is generated with go:generate comments.
package widgets


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
        Hash(v *T) uint64
        // Zero's the supplied value. Equivalent to a destructor except Go does
        // not have manual memory management so this is mainly just to prenent
        // things like dangling pointers.
        Zero(v *T)
    }

    // The base widget implementation that all the containers in the [containers]
    // package use as a type restriction. This type must be instantiaed with the
    // [NewWidget] function; zero valued Widget's are not valid and will result
    // in nil pointer errors. The size of this struct will equal the size of the 
    // underlying WidgetInterface[T] value. This has the implication that if the
    // interface type is zero bytes the the widget will also be zero bytes. This
    // is important because containers will create widgets as needed, so minimizing
    // the size of the widget will reduce memory allocations. It is also for this
    // reason that any type that implements the widget interface 'for itself'
    // should have pointer recievers for the widget interface methods unless the 
    // underlying type is small enough to make this concern mute.
    Widget[T any, I WidgetInterface[T]] struct {
        iFace I
    }
)

// Creates a new widget and sets its internal state so that it is valid and can
// be used without error.
func NewWidget[T any, I WidgetInterface[T]]() Widget[T,I] {
    var iFaceImpl I
    return Widget[T, I]{iFace: iFaceImpl}
}

// Zeros the supplied value using the logic defined by the interface that was
// supplied as a generic type.
func (w *Widget[T, I])Zero(v *T) {
    w.iFace.Zero(v)
}

// Compares the left (l) and right (r) values and returns true is l==r using the
// Eq function from the interface that was supplied as a generic type.
func (w *Widget[T, I])Eq(l *T, r *T) bool {
    return w.iFace.Eq(l,r)
}

// Compares the left (l) and right (r) values and returns true is l!=r using the
// Eq function from the interface that was supplied as a generic type.
func (w *Widget[T, I])Neq(l *T, r *T) bool {
    return !w.iFace.Eq(l,r)
}

// Compares the left (l) and right (r) values and returns true is l<r using the
// Lt function from the interface that was supplied as a generic type.
func (w *Widget[T, I])Lt(l *T, r *T) bool {
    return w.iFace.Lt(l,r)
}

// Compares the left (l) and right (r) values and returns true is l<=r using the
// Lt and Eq functions from the interface that was supplied as a generic type.
func (w *Widget[T, I])Lte(l *T, r *T) bool {
    return w.iFace.Lt(l,r) || w.iFace.Eq(l,r)
}

// Compares the left (l) and right (r) values and returns true is l>r using the
// Lt and Eq functions from the interface that was supplied as a generic type.
func (w *Widget[T, I])Gt(l *T, r *T) bool {
    return !w.iFace.Lt(l,r) && !w.iFace.Eq(l,r)
}

// Compares the left (l) and right (r) values and returns true is l>=r using the
// Lt function from the interface that was supplied as a generic type.
func (w *Widget[T, I])Gte(l *T, r *T) bool {
    return !w.iFace.Lt(l,r)
}

// Generates a hash for the given value using the hash function from the interface
// that was supplied as a generic type.
func (w *Widget[T, I])Hash(v *T) uint64 {
    return w.iFace.Hash(v)
}

// TODO - add pntr types
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
