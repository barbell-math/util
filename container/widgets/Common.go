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
        Eq(other *T) bool
        // A function that should return true if the current value is less than other.
        Lt(other *T) bool
        // Unwraps the current value, returning a pointer to it. This is meant to 
        // be a very lightweight call so avoid copying values as much as possible. 
        // The [Widget] struct provides this method by default.
        Unwrap() *T
        // Wraps the value in the pointer. This is meant to be a very lightweight
        // call so avoid copying values as much as possible. The [Widget] struct
        // provides this method by default.
        Wrap(v *T)
        // Returns a hash value that represent the currently wrapped value.
        Hash() uint64
    }

    // The base widget implementation that all the containers in the [containers]
    // package use as a type restriction.
    Widget[T any] struct { v *T }
)

func (w Widget[T])Wrap(v *T) {
    w.v=v
}

func (w Widget[T])Unwrap() *T {
    return w.v
}

//go:generate go run widgetTemplate.go -package=widgets -type=int
//go:generate go run widgetTemplate.go -package=widgets -type=int8
//go:generate go run widgetTemplate.go -package=widgets -type=int16
//go:generate go run widgetTemplate.go -package=widgets -type=int32
//go:generate go run widgetTemplate.go -package=widgets -type=int64

//go:generate go run widgetTemplate.go -package=widgets -type=uint
//go:generate go run widgetTemplate.go -package=widgets -type=uint8
//go:generate go run widgetTemplate.go -package=widgets -type=uint16
//go:generate go run widgetTemplate.go -package=widgets -type=uint32
//go:generate go run widgetTemplate.go -package=widgets -type=uint64

//go:generate go run widgetTemplate.go -package=widgets -type=float32
//go:generate go run widgetTemplate.go -package=widgets -type=float64

//go:generate go run widgetTemplate.go -package=widgets -type=string
