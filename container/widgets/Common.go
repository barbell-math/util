// This package serves to implement the comparison and hash operations necessary
// to put built-in types into the containers defined in [containers]. All of the
// types in this package will implement the [containers.Widget] interface.
// Basically all of the code in this package is generated with go:generate comments.
package widgets


type (
    // The type of values that the containers will act upon. This interface 
    // enforces all the required information is exposed by the underlying types 
    // held in the container.
    WidgetInterface[T any] interface {
        // A function that should return true if the current value equals other.
        Eq(l *T, r *T) bool
        // A function that should return true if the current value is less than other.
        Lt(l *T, r *T) bool
        // Returns a hash value that represent the currently wrapped value.
        Hash(v *T) uint64
    }

    // The base widget implementation that all the containers in the [containers]
    // package use as a type restriction.
    Widget[T any, I WidgetInterface[T]] struct { 
        v *T
        iFace I
    }
)

func NewWidget[T any, I WidgetInterface[T]](i I) Widget[T,I] {
    return Widget[T, I]{iFace: i}
}

func (w *Widget[T, I])Eq(l *T, r *T) bool {
    return w.iFace.Eq(l,r)
}

func (w *Widget[T, I])Lt(l *T, r *T) bool {
    return w.iFace.Lt(l,r)
}

func (w *Widget[T, I])Hash() uint64 {
    return w.iFace.Hash(w.v)
}

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
