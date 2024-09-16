package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"

	"github.com/barbell-math/util/generators/common"
)

type (
	InlineArgs struct {
		Type string `required:"t" help:"The type to make the alias for."`
		ShowInfo       bool   `required:"f" default:"t" help:"Show debug info."`
	}
	CommentArgs struct {
		WidgetType string `required:"t" help:"The type of widget to make. (Base, PartialOrder, Arith)."`
		Package        string `required:"t" help:"The packge to put the files in."`
		BaseType       string `required:"t" help:"The base type to generate the widget for."`
		BaseTypeWidget string `required:"t" help:"The base type widget to use when generating the new widget."`
		WidgetPackage  string `required:"t" help:"The package the base type widget resides in. If it is this package, put '.'"`
	}

	widgetType string
	ProgState struct {
		widgetType widgetType
	}
	TemplateVals struct {
		GeneratorName string
		Package string
		Imports []string
		AliasType string
		BaseType string
		BaseTypeWidget string
	}
)

const (
	BaseWidget widgetType="Base"
	PartialOrderWidget widgetType="PartialOrder"
	ArithWidget widgetType="Arith"
	Unknown widgetType="UNKNOWN"
)

var (
	INLINE_ARGS InlineArgs
	COMMENT_ARGS CommentArgs
	PROG_STATE ProgState
	TEMPLATES common.GeneratedFilesRegistry=common.NewGeneratedFilesRegistryFromMap(
		map[string]string{
			"imports": `
import (
	"github.com/barbell-math/util/hash"
	{{range .Imports}}
		"{{ . }}"
	{{end}}
)
`,
			"eqFunc": `
// Returns true if l equals r. Uses the Eq operator provided by the 
// {{ .BaseTypeWidget }} widget internally.
func (_ *{{ .AliasType }}) Eq(l *{{ .AliasType }}, r *{{ .AliasType }}) bool {
	var tmp {{ .BaseTypeWidget }}
	return tmp.Eq((*{{ .BaseType }})(l), (*{{ .BaseType }})(r))
}
`,
			"ltFunc": `
// Returns true if l is less than r. Uses the Lt operator provided by the 
// {{ .BaseTypeWidget }} widget internally.
func (_ *{{ .AliasType }}) Lt(l *{{ .AliasType }}, r *{{ .AliasType }}) bool {
	var tmp {{ .BaseTypeWidget }}
	return tmp.Lt((*{{ .BaseType }})(l), (*{{ .BaseType }})(r))
}
`,
			"hashFunc": `
// Returns a hash to represent other. The hash that is returned will be supplied
// by the {{ .BaseTypeWidget }} widget internally.
func (_ *{{ .AliasType }}) Hash(other *{{ .AliasType }}) hash.Hash {
	var tmp {{ .BaseTypeWidget }}
	return tmp.Hash((*{{ .BaseType }})(other))
}
`,
			"zeroFunc": `
// Zeros the supplied value. The operation that is performed will be determined
// by the {{ .BaseTypeWidget }} widget internally.
func (_ *{{ .AliasType }}) Zero(other *{{ .AliasType }}) {
	var tmp {{ .BaseTypeWidget }}
	tmp.Zero((*{{ .BaseType }})(other))
}
`,
			"zeroValFunc": `
// Returns the zero value for the underlying type. The value that is performed
// will be determined by the {{ .BaseTypeWidget }} widget internally.
func (_ *{{ .AliasType }}) ZeroVal() {{ .AliasType }} {
	var tmp {{ .BaseTypeWidget }}
	return ({{ .AliasType }})(tmp.ZeroVal())
}
`,
			"unitValFunc": `
// Returns the value that represent "1" for the underlying type. The value that
// is performed will be determined by the {{ .BaseTypeWidget }} widget internally.
func (_ *{{ .AliasType }}) UnitVal() {{ .AliasType }} {
	var tmp {{ .BaseTypeWidget }}
	return ({{ .AliasType }})(tmp.ZeroVal())
}
`,
			"negFunc": `
// Negates the value that is supplied to it. The value that is returned will be
// determined by the {{ .BaseTypeWidget }} widget internally.
func (_ *{{ .AliasType }}) Neg(v *{{ .AliasType }}) {{ .AliasType }} {
	var tmp {{ .BaseTypeWidget }}
	return ({{ .AliasType }})(tmp.Neg((*{{ .BaseTypeWidget }})(v)))
}
`,
			"addFunc": `
// Adds l and r and places the results in res. The value that is returned will
// be determined by the {{ .BaseTypeWidget }} widget internally.
func (_ *{{ .AliasType}}) Add(res *{{ .AliasType }}, l *{{ .AliasType }}, r *{{ .AliasType }}) {
	var tmp {{ .BaseTypeWidget }}
	return ({{ .AliasType }})(tmp.Add(
		(*{{ .BaseTypeWidget }})(res),
		(*{{ .BaseTypeWidget }})(l),
		(*{{ .BaseTypeWidget }})(r),
	))
}
`,
			"subFunc": `
// Subtracts l and r and places the results in res. The value that is returned
// will be determined by the {{ .BaseTypeWidget }} widget internally.
func (_ *{{ .AliasType}}) Sub(res *{{ .AliasType }}, l *{{ .AliasType }}, r *{{ .AliasType }}) {
	var tmp {{ .BaseTypeWidget }}
	return ({{ .AliasType }})(tmp.Sub(
		(*{{ .BaseTypeWidget }})(res),
		(*{{ .BaseTypeWidget }})(l),
		(*{{ .BaseTypeWidget }})(r),
	))
}
`,
			"mulFunc": `
// Multiplys l and r and places the results in res. The value that is returned
// will be determined by the {{ .BaseTypeWidget }} widget internally.
func (_ *{{ .AliasType}}) Mul(res *{{ .AliasType }}, l *{{ .AliasType }}, r *{{ .AliasType }}) {
	var tmp {{ .BaseTypeWidget }}
	return ({{ .AliasType }})(tmp.Mul(
		(*{{ .BaseTypeWidget }})(res),
		(*{{ .BaseTypeWidget }})(l),
		(*{{ .BaseTypeWidget }})(r),
	))
}
`,
			"divFunc": `
// Divides l and r and places the results in res. The value that is returned
// will be determined by the {{ .BaseTypeWidget }} widget internally.
func (_ *{{ .AliasType}}) Div(res *{{ .AliasType }}, l *{{ .AliasType }}, r *{{ .AliasType }}) {
	var tmp {{ .BaseTypeWidget }}
	return ({{ .AliasType }})(tmp.Div(
		(*{{ .BaseTypeWidget }})(res),
		(*{{ .BaseTypeWidget }})(l),
		(*{{ .BaseTypeWidget }})(r),
	))
}
`,
			"arithFile": `
{{template "partialOrderFile" .}}
{{template "zeroValFunc" .}}
{{template "unitValFunc" .}}
{{template "negFunc" .}}
{{template "addFunc" .}}
{{template "subFunc" .}}
{{template "mulFunc" .}}
{{template "divFunc" .}}
`,
			"partialOrderFile": `
{{template "baseFile" .}}
{{template "ltFunc" .}}
`,
			"baseFile": `
package {{ .Package }}

{{template "autoGenComment" .}}
{{template "imports" .}}

{{template "eqFunc" .}}
{{template "hashFunc" .}}
{{template "zeroFunc" .}}
`,
		},
	)
)

func (w widgetType) FromString(s string) widgetType {
	switch s {
	case string(BaseWidget): return BaseWidget
	case string(PartialOrderWidget): return PartialOrderWidget
	case string(ArithWidget): return ArithWidget
	default: return Unknown
	}
}

func main() {
	common.InlineArgs(&INLINE_ARGS, os.Args)

	commentArgsFilter:=common.CommentArgsAstFilter(INLINE_ARGS.Type, &COMMENT_ARGS)
	common.IterateOverAST(
		".",
		common.GenFileExclusionFilter,
		func(fSet *token.FileSet, file *ast.File, srcFile *os.File, node ast.Node) bool {
			commentArgsFilter(fSet, srcFile, node)
			switch node.(type) {
			case *ast.GenDecl: return false
			case *ast.FuncDecl: return false
			default: return true
			}
		},
	)

	if t:=Unknown.FromString(COMMENT_ARGS.WidgetType); t==Unknown {
		common.PrintRunningError(
			"The supplied widget type was not recognized. Must be one of %v",
			[]widgetType{BaseWidget, PartialOrderWidget, ArithWidget},
		)
		os.Exit(1)
	} else {
		PROG_STATE.widgetType=t
	}

	templateData:=TemplateVals{
		GeneratorName: os.Args[0],
		Package: COMMENT_ARGS.Package,
		Imports: []string{},
		AliasType: INLINE_ARGS.Type,
		BaseType: COMMENT_ARGS.BaseType,
		BaseTypeWidget: COMMENT_ARGS.BaseTypeWidget,
	}
	if COMMENT_ARGS.WidgetPackage!="." {
		templateData.Imports = append(templateData.Imports, COMMENT_ARGS.WidgetPackage)
	}

	if err := TEMPLATES.WriteToFile(
		fileName(),
		common.GeneratedSrcFileExt,
		map[widgetType]string{
			BaseWidget: "baseFile",
			PartialOrderWidget: "partialOrderFile",
			ArithWidget: "arithFile",
		}[PROG_STATE.widgetType],
		templateData,
	); err != nil {
		common.PrintRunningError("%s", err)
		os.Exit(1)
	}
}

func fileName() string {
	fileNameBaseType := []byte(COMMENT_ARGS.BaseType)
	fileNameAliasType := []byte(INLINE_ARGS.Type)
	bannedChars := map[byte]struct{}{
		'[':  {},
		']':  {},
		'{':  {},
		'}':  {},
		':':  {},
		';':  {},
		'<':  {},
		'>':  {},
		',':  {},
		'.':  {},
		'/':  {},
		'\\': {},
		'|':  {},
		'*':  {},
		'?':  {},
		'%':  {},
		'"':  {},
		' ':  {},
	}
	for i, c := range fileNameBaseType {
		if _, ok := bannedChars[c]; ok {
			fileNameBaseType[i] = '_'
		}
	}
	for i, c := range fileNameAliasType {
		if _, ok := bannedChars[c]; ok {
			fileNameAliasType[i] = '_'
		}
	}

	return fmt.Sprintf(
		"TypeAliasPassThroughWidget_%s_to_%s",
		fileNameAliasType,
		fileNameBaseType,
	)
}
