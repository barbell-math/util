// This is a generator program that is used to create pass through widgets. A pass
// through widget is a widget that simply calls the methods of another widget
// without adding any additional logic. It is assumed that the two widget types can
// be cast to and from each other.
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
		Type     string `required:"t" help:"The type to make the alias for."`
		ShowInfo bool   `required:"f" default:"t" help:"Show debug info."`
	}
	CommentArgs struct {
		WidgetType     string `required:"t" help:"The type of widget to make. (Base, PartialOrder, Arith)."`
		BaseTypeWidget string `required:"t" help:"The base type widget to use when generating the new widget."`
		WidgetPackage  string `required:"t" help:"The package the base type widget resides in. If it is this package, put '.'"`
	}

	widgetType string
	ProgState  struct {
		widgetType widgetType
		baseType   string
		_package string
	}
	TemplateVals struct {
		GeneratorName  string
		Package        string
		Imports        []string
		AliasType      string
		BaseType       string
		BaseTypeWidget string
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
	INLINE_ARGS  InlineArgs
	COMMENT_ARGS CommentArgs
	PROG_STATE   ProgState
	TEMPLATES    common.GeneratedFilesRegistry = common.NewGeneratedFilesRegistryFromMap(
		map[string]string{
			"imports": `
import (
	"github.com/barbell-math/util/hash"
	{{range .Imports}} "{{ . }}" {{end}}
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
			"partialOrderArithFile": `
{{template "baseFile" .}}
{{template "ltFunc" .}}
{{template "zeroValFunc" .}}
{{template "unitValFunc" .}}
{{template "negFunc" .}}
{{template "addFunc" .}}
{{template "subFunc" .}}
{{template "mulFunc" .}}
{{template "divFunc" .}}
`,
			"arithFile": `
{{template "baseFile" .}}
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

	commentArgsFilter, foundArgs := common.DocArgsAstFilter(
		INLINE_ARGS.Type, &COMMENT_ARGS,
	)
	common.IterateOverAST(
		".",
		common.GenFileExclusionFilter,
		func(fSet *token.FileSet, file *ast.File, srcFile *os.File, node ast.Node) bool {
			commentArgsFilter(fSet, srcFile, node)
			switch node.(type) {
			case *ast.GenDecl:
				if PROG_STATE.baseType != "" {
					return false
				}
				gdNode := node.(*ast.GenDecl)
				if gdNode.Tok == token.TYPE {
					for _, spec := range gdNode.Specs {
						parseTypeSpec(
							fSet,
							srcFile,
							spec.(*ast.TypeSpec),
						)
					}
				}
				if PROG_STATE.baseType!="" {
					PROG_STATE._package=file.Name.Name
				}
				return false
			case *ast.FuncDecl:
				return false
			default:
				return true
			}
		},
	)

	if !*foundArgs {
		common.PrintRunningError(
			"The supplied type did not have comment args but they are required.",
		)
		os.Exit(1)
	}
	if PROG_STATE.baseType == "" {
		common.PrintRunningError("The base type was not found!")
		os.Exit(1)
	}

	if t := Unknown.FromString(COMMENT_ARGS.WidgetType); t == Unknown {
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
		PROG_STATE.widgetType = t
	}

	templateData := TemplateVals{
		GeneratorName:  os.Args[0],
		Package:        PROG_STATE._package,
		Imports:        []string{},
		AliasType:      INLINE_ARGS.Type,
		BaseType:       PROG_STATE.baseType,
		BaseTypeWidget: COMMENT_ARGS.BaseTypeWidget,
	}
	if COMMENT_ARGS.WidgetPackage != "." {
		templateData.Imports = append(templateData.Imports, COMMENT_ARGS.WidgetPackage)
	}

	if err := TEMPLATES.WriteToFile(
		fileName(),
		common.GeneratedSrcFileExt,
		map[widgetType]string{
			BaseWidget:              "baseFile",
			PartialOrderWidget:      "partialOrderFile",
			ArithWidget:             "arithFile",
			PartialOrderArithWidget: "partialOrderArithFile",
		}[PROG_STATE.widgetType],
		templateData,
	); err != nil {
		common.PrintRunningError("%s", err)
		os.Exit(1)
	}
}

func fileName() string {
	fileNameBaseType := []byte(PROG_STATE.baseType)
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

func parseTypeSpec(
	fSet *token.FileSet,
	srcFile *os.File,
	ts *ast.TypeSpec,
) {
	name, err := common.GetSourceTextFromExpr(
		fSet, srcFile, ts.Name,
	)
	fmt.Println(name)
	if err != nil {
		common.PrintRunningError("%s", err)
		os.Exit(1)
	}

	if name == INLINE_ARGS.Type {
		baseType, err := common.GetSourceTextFromExpr(
			fSet, srcFile, ts.Type,
		)
		if err != nil {
			common.PrintRunningError("%s", err)
			os.Exit(1)
		}

		PROG_STATE.baseType = baseType
	}
}
