// This package serves to implement the comparison and hash operations necessary
// to put built-in types into the containers defined in [containers]. All of the
// types in this package will implement the [containers.Widget] interface.
// Basically all of the code in this package is generated with go:generate comments.
package builtinWidgets

//go:generate go run widgetTemplate.go builtinWidgets int
//go:generate go run widgetTemplate.go builtinWidgets int8
//go:generate go run widgetTemplate.go builtinWidgets int16
//go:generate go run widgetTemplate.go builtinWidgets int32
//go:generate go run widgetTemplate.go builtinWidgets int64

//go:generate go run widgetTemplate.go builtinWidgets uint
//go:generate go run widgetTemplate.go builtinWidgets uint8
//go:generate go run widgetTemplate.go builtinWidgets uint16
//go:generate go run widgetTemplate.go builtinWidgets uint32
//go:generate go run widgetTemplate.go builtinWidgets uint64

//go:generate go run widgetTemplate.go builtinWidgets float32
//go:generate go run widgetTemplate.go builtinWidgets float64

//go:generate go run widgetTemplate.go builtinWidgets string
