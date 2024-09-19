// This is a generator program that is used internally by the widgets package to
// generate the widgets for the builtin types.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/barbell-math/util/generators/common"
)

type (
	InlineArgs struct {
		Type       string `required:"t" help:"The underlying type to generate the widget for."`
		WidgetType string `required:"t" help:"The type of widget to generate. (Base, PartialOrder, Arith)"`
		ShowInfo   bool   `required:"f" default:"t" help:"Show debug info."`
	}

	TypeInfo struct {
		ZeroValue  string
		HashTempl  string
		Imports    []string
		GlobalVars []string
	}

	widgetType string
	ProgState  struct {
		WidgetType widgetType
		CapType    string
	}

	TemplateVals struct {
		GeneratorName string
		Package       string
		Imports       []string
		GlobalVars    []string
		Type          string
		CapType       string
		ZeroValue     string
		HashTemplRes  string
	}
)

const (
	BaseWidget              widgetType = "Base"
	PartialOrderWidget      widgetType = "PartialOrder"
	ArithWidget             widgetType = "Arith"
	PartialOrderArithWidget widgetType = "PartialOrderArith"
	Unknown                 widgetType = "UNKNOWN"
)

var (
	INLINE_ARGS InlineArgs
	PROG_STATE  ProgState
	VALID_TYPES map[string]TypeInfo = map[string]TypeInfo{
		"bool": {
			ZeroValue: "false",
			Imports:   []string{},
			HashTempl: `if *v { return hash.Hash(1) } 
return hash.Hash(0) `,
		},
		"byte": {
			ZeroValue: "0",
			Imports:   []string{},
			HashTempl: `return hash.Hash(*v)`,
		},
		"int": {
			ZeroValue: "0",
			Imports:   []string{},
			HashTempl: `return hash.Hash(*v)`,
		},
		"int8": {
			ZeroValue: "0",
			Imports:   []string{},
			HashTempl: `return hash.Hash(*v)`,
		},
		"int16": {
			ZeroValue: "0",
			Imports:   []string{},
			HashTempl: `return hash.Hash(*v)`,
		},
		"int32": {
			ZeroValue: "0",
			Imports:   []string{},
			HashTempl: `return hash.Hash(*v)`,
		},
		"int64": {
			ZeroValue: "0",
			Imports:   []string{},
			HashTempl: `return hash.Hash(*v)`,
		},
		"uint": {
			ZeroValue: "0",
			Imports:   []string{},
			HashTempl: `return hash.Hash(*v)`,
		},
		"uint8": {
			ZeroValue: "0",
			Imports:   []string{},
			HashTempl: `return hash.Hash(*v)`,
		},
		"uint16": {
			ZeroValue: "0",
			Imports:   []string{},
			HashTempl: `return hash.Hash(*v)`,
		},
		"uint32": {
			ZeroValue: "0",
			Imports:   []string{},
			HashTempl: `return hash.Hash(*v)`,
		},
		"uint64": {
			ZeroValue: "0",
			Imports:   []string{},
			HashTempl: `return hash.Hash(*v)`,
		},
		"float32": {
			ZeroValue: "0",
			Imports:   []string{"unsafe"},
			HashTempl: `return hash.Hash(*(*uint32)(unsafe.Pointer(v)))`,
		},
		"float64": {
			ZeroValue: "0",
			Imports:   []string{"unsafe"},
			HashTempl: `return hash.Hash(*(*uint64)(unsafe.Pointer(v)))`,
		},
		"complex64": {
			ZeroValue: "0",
			Imports:   []string{"github.com/barbell-math/util/math/basic"},
			HashTempl: `return hash.Hash(basic.LossyConv[float32, int32](basic.RealPart[complex64, float32](*v))).
	Combine(hash.Hash(basic.LossyConv[float32, int32](basic.ImaginaryPart[complex64, float32](*v))))
`,
		},
		"complex128": {
			ZeroValue: "0",
			Imports:   []string{"github.com/barbell-math/util/math/basic"},
			HashTempl: `return hash.Hash(basic.LossyConv[float64, int64](basic.RealPart[complex128, float64](*v))).
	Combine(hash.Hash(basic.LossyConv[float64, int64](basic.ImaginaryPart[complex128, float64](*v))))
`,
		},
		"string": {
			ZeroValue: "\"\"",
			Imports:   []string{"hash/maphash"},
			HashTempl: `return hash.Hash(maphash.String(RANDOM_SEED_STRING, *(v)))`,
			GlobalVars: []string{
				`// The random seed will be different every time the program runs
// meaning that between runs the hash values will not be consistent.
// This was done for security purposes.
RANDOM_SEED_STRING maphash.Seed = maphash.MakeSeed()
`,
			},
		},
		"uintptr": {
			ZeroValue: "0",
			Imports:   []string{},
			HashTempl: `return hash.Hash(*v)`,
		},
		// This is a special case that is only allowed because the widget package itself
		// relies on hash.Hash, making it so the hash.Hash package cannot implement the
		// widget interface on itself (would create circular imports)
		"hash.Hash": {
			ZeroValue: "0",
			Imports:   []string{},
			HashTempl: `return *v`,
		},
	}
	TEMPLATES common.GeneratedFilesRegistry = common.NewGeneratedFilesRegistryFromMap(
		map[string]string{
			"imports": `
import (
	"github.com/barbell-math/util/hash"
	{{range .Imports}}
		"{{ . }}"
	{{end}}
)
`,
			"globalVars": `
var (
	{{range .GlobalVars}} {{ . }} {{end}}
)
`,
			"eqFunc": `
// Returns true if both {{ .Type }}'s are equal. Uses the standard == operator internally.
func (_ Builtin{{ .CapType }}) Eq(l *{{ .Type }}, r *{{ .Type }}) bool {
	return *l == *r
}
`,
			"ltFunc": `
// Returns true if l<r. Uses the standard < operator internally.
func (_ Builtin{{ .CapType }}) Lt(l *{{ .Type }}, r *{{ .Type }}) bool {
	return *l < *r
}
`,
			"hashFunc": `
// Provides a hash function for the value that it is wrapping.
func (_ Builtin{{ .CapType }}) Hash(v *{{ .Type }}) hash.Hash {
	{{ .HashTemplRes }}
}
`,
			"zeroFunc": `
// Zeros the supplied value.
func (_ Builtin{{ .CapType }}) Zero(other *{{ .Type }}) {
	*other = ({{ .Type }})({{ .ZeroValue }})
}
`,
			"arithFuncs": `
// Returns the zero value for the {{ .Type }} type.
func (_ Builtin{{ .CapType }}) ZeroVal() {{ .Type }} {
	return {{ .Type }}(0)
}

// Returns the unit value for the {{ .Type }} type.
func (_ Builtin{{ .CapType }}) UnitVal() {{ .Type }} {
	return {{ .Type }}(1)
}

// Negates v, updating the value that v points to.
func (_ Builtin{{ .CapType }}) Neg(v *{{ .Type }}) {
	*v = -(*v)
}

// Adds l to r, placing the result in the value that res points to.
func (_ Builtin{{ .CapType }}) Add(res *{{ .Type }}, l *{{ .Type }}, r *{{ .Type }}) {
	*res = *l + *r
}

// Subtracts l to r, placing the result in the value that res points to.
func (_ Builtin{{ .CapType }}) Sub(res *{{ .Type }}, l *{{ .Type }}, r *{{ .Type }}) {
	*res = *l - *r
}

// Multiplies l to r, placing the result in the value that res points to.
func (_ Builtin{{ .CapType }}) Mul(res *{{ .Type }}, l *{{ .Type }}, r *{{ .Type }}) {
	*res = *l * *r
}

// Divides l to r, placing the result in the value that res points to.
func (_ Builtin{{ .CapType }}) Div(res *{{ .Type }}, l *{{ .Type }}, r *{{ .Type }}) {
	*res = *l / *r
}
`,
			"partialOrderArithFile": `
{{template "baseFile" .}}
{{template "ltFunc" .}}
{{template "arithFuncs" .}}
`,
			"arithFile": `
{{template "baseFile" .}}
{{template "arithFuncs" .}}
`,
			"partialOrderFile": `
{{template "baseFile" .}}
{{template "ltFunc" .}}
`,
			"baseFile": `
package widgets

{{template "autoGenComment" .}}
{{template "imports" .}}
{{template "globalVars" .}}

// A widget to represent the built-in {{ .Type }} type
type (
	Builtin{{ .CapType }} struct{}
)

{{template "eqFunc" .}}
{{template "hashFunc" .}}
{{template "zeroFunc" .}}
`,
		},
	)
)

func (w widgetType) FromString(s string) widgetType {
	switch s {
	case string(BaseWidget):
		return BaseWidget
	case string(PartialOrderWidget):
		return PartialOrderWidget
	case string(ArithWidget):
		return ArithWidget
	case string(PartialOrderArithWidget):
		return PartialOrderArithWidget
	default:
		return Unknown
	}
}

func main() {
	common.InlineArgs(&INLINE_ARGS, os.Args)

	if _, ok := VALID_TYPES[INLINE_ARGS.Type]; !ok {
		fmt.Println("ERROR | The supplied type was not one of the types recognized by this tool.")
		fmt.Println("The following types are recognized: ", VALID_TYPES)
		fmt.Println("The following type was received: ", INLINE_ARGS.Type)
		os.Exit(1)
	}

	if t := Unknown.FromString(INLINE_ARGS.WidgetType); t == Unknown {
		common.PrintRunningError(
			"The supplied widget type was not recognized. Must be one of %v",
			[]widgetType{
				BaseWidget,
				PartialOrderWidget,
				ArithWidget,
				PartialOrderArithWidget,
			},
		)
		os.Exit(1)
	} else {
		PROG_STATE.WidgetType = t
	}

	PROG_STATE.CapType = INLINE_ARGS.Type
	dotSplit := strings.SplitN(PROG_STATE.CapType, ".", 2)
	if len(dotSplit) > 1 {
		PROG_STATE.CapType = dotSplit[len(dotSplit)-1]
	}
	PROG_STATE.CapType = fmt.Sprintf("%s%s", strings.ToUpper(PROG_STATE.CapType)[:1], PROG_STATE.CapType[1:])

	templateData := TemplateVals{
		GeneratorName: os.Args[0],
		Type:          INLINE_ARGS.Type,
		CapType:       PROG_STATE.CapType,
		ZeroValue:     VALID_TYPES[INLINE_ARGS.Type].ZeroValue,
		HashTemplRes:  VALID_TYPES[INLINE_ARGS.Type].HashTempl,
		Imports:       VALID_TYPES[INLINE_ARGS.Type].Imports,
		GlobalVars:    VALID_TYPES[INLINE_ARGS.Type].GlobalVars,
	}

	if err := TEMPLATES.WriteToFile(
		fmt.Sprintf("Builtin%s", PROG_STATE.CapType),
		common.GeneratedSrcFileExt,
		map[widgetType]string{
			BaseWidget:              "baseFile",
			PartialOrderWidget:      "partialOrderFile",
			ArithWidget:             "arithFile",
			PartialOrderArithWidget: "partialOrderArithFile",
		}[PROG_STATE.WidgetType],
		templateData,
	); err != nil {
		common.PrintRunningError("%s", err)
		os.Exit(1)
	}
}
