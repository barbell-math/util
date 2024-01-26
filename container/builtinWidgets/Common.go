// This package serves to implement the comparison and hash operations necessary
// to put built-in types into the containers defined in [containers]. All of the
// types in this package will implement the [containers.Widget] interface.
// Basically all of the code in this package is generated with go:generate comments.
package builtinWidgets

//go:generate go run widgetTemplate.go -package=builtinWidgets -type=int
//go:generate go run widgetTemplate.go -package=builtinWidgets -type=int8
//go:generate go run widgetTemplate.go -package=builtinWidgets -type=int16
//go:generate go run widgetTemplate.go -package=builtinWidgets -type=int32
//go:generate go run widgetTemplate.go -package=builtinWidgets -type=int64

//go:generate go run widgetTemplate.go -package=builtinWidgets -type=uint
//go:generate go run widgetTemplate.go -package=builtinWidgets -type=uint8
//go:generate go run widgetTemplate.go -package=builtinWidgets -type=uint16
//go:generate go run widgetTemplate.go -package=builtinWidgets -type=uint32
//go:generate go run widgetTemplate.go -package=builtinWidgets -type=uint64

//go:generate go run widgetTemplate.go -package=builtinWidgets -type=float32
//go:generate go run widgetTemplate.go -package=builtinWidgets -type=float64

//go:generate go run widgetTemplate.go -package=builtinWidgets -type=string
