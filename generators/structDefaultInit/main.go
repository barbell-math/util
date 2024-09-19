package main

import (
	"go/ast"
	"go/token"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/barbell-math/util/generators/common"
)

type (
	InlineArgs struct {
		Struct   string `required:"t" help:"The struct type to generate code for."`
		Package  string `required:"t" help:"The package the options enum type is in."`
		ShowInfo bool   `required:"f" default:"t" help:"Show debug info."`
	}

	ProgState struct {
		fieldSetters  []string
		fieldGetters  []string
		fieldDefaults []DefaultInfo
		fieldTypes    map[string]string
		fieldComments map[string]string
		imports       []string
	}
	DefaultInfo struct {
		name  string
		value string
	}

	TemplateVals struct {
		StructName          string
		CapStructName       string
		Package             string
		StructFieldSetters  []StructFieldTemplateVals
		StructFieldGetters  []StructFieldTemplateVals
		StructFieldDefaults []StructDefaultTemplateVals
		Imports             []string
		GeneratorName       string
	}
	EnumFlagTemplateVals struct {
		StructName  string
		CapEnumFlag string
		EnumFlag    string
		Comment     string
	}
	StructDefaultTemplateVals struct {
		StructField        string
		StructFieldDefault string
	}
	StructFieldTemplateVals struct {
		StructName     string
		CapStructField string
		StructField    string
		FieldType      string
		Comment        string
	}
)

var (
	INLINE_ARGS InlineArgs
	PROG_STATE  ProgState = ProgState{
		fieldSetters:  []string{},
		fieldGetters:  []string{},
		fieldDefaults: []DefaultInfo{},
		fieldTypes:    map[string]string{},
		fieldComments: map[string]string{},
		imports:       []string{},
	}
	TEMPLATES common.GeneratedFilesRegistry = common.NewGeneratedFilesRegistryFromMap(
		map[string]string{
			"structFieldSetter": `
{{ .Comment }}
func (o *{{ .StructName }}) {{ .CapStructField }}(v {{ .FieldType }}) *{{ .StructName }} {
	o.{{ .StructField }} = v
	return o
}
`,
			"structFieldGetter": `
{{ .Comment }}
func (o *{{ .StructName }}) Get{{ .CapStructField }}() {{ .FieldType }} {
	return o.{{ .StructField }}
}
`,
			"defaultValueInit": `{{ .StructField }}: {{ .StructFieldDefault }},
`,
			"newFunc": `
// Returns a new {{ .StructName }} struct initialized with the default values.
func New{{ .CapStructName }}() *{{ .StructName }} {
	return &{{ .StructName }} {
		{{range .StructFieldDefaults}} {{template "defaultValueInit" .}} {{end}}
	}
}
`,
			"file": `
package {{ .Package }}
{{template "autoGenComment" .}}
import (
	{{range .Imports}}{{.}}{{end}}
)

{{template "newFunc" .}}
{{range .StructFieldSetters}}
	{{template "structFieldSetter" .}}
{{end}}
{{range .StructFieldGetters}}
	{{template "structFieldGetter" .}}
{{end}}
`,
		},
	)
)

func main() {
	common.InlineArgs(&INLINE_ARGS, os.Args)

	optionsStructFound := false

	common.IterateOverAST(
		".",
		common.GenFileExclusionFilter,
		func(fSet *token.FileSet, file *ast.File, srcFile *os.File, node ast.Node) bool {
			switch node.(type) {
			case *ast.GenDecl:
				gdNode := node.(*ast.GenDecl)
				if gdNode.Tok == token.TYPE && !optionsStructFound {
					for _, spec := range gdNode.Specs {
						optionsStructFound = parseTypeSpec(
							fSet,
							srcFile,
							spec.(*ast.TypeSpec),
							&PROG_STATE,
						)
						if optionsStructFound {
							break
						}
					}
					return false
				}
				return false
			case *ast.FuncDecl:
				return false
			default:
				return true
			}
		},
	)

	templateData := TemplateVals{
		StructName:          INLINE_ARGS.Struct,
		CapStructName:       strings.ToUpper(INLINE_ARGS.Struct[0:1]) + INLINE_ARGS.Struct[1:],
		Package:             INLINE_ARGS.Package,
		StructFieldSetters:  make([]StructFieldTemplateVals, len(PROG_STATE.fieldSetters)),
		StructFieldGetters:  make([]StructFieldTemplateVals, len(PROG_STATE.fieldGetters)),
		StructFieldDefaults: make([]StructDefaultTemplateVals, len(PROG_STATE.fieldDefaults)),
		Imports:             PROG_STATE.imports,
		GeneratorName:       os.Args[0],
	}
	cntr := 0
	for i, v := range PROG_STATE.fieldSetters {
		capStructField := strings.ToUpper(v[0:1]) + v[1:]
		templateData.StructFieldSetters[i] = StructFieldTemplateVals{
			StructName:     INLINE_ARGS.Struct,
			CapStructField: capStructField,
			StructField:    v,
			FieldType:      PROG_STATE.fieldTypes[v],
			Comment:        PROG_STATE.fieldComments[v],
		}
	}
	for i, v := range PROG_STATE.fieldGetters {
		capStructField := strings.ToUpper(v[0:1]) + v[1:]
		templateData.StructFieldGetters[i] = StructFieldTemplateVals{
			StructName:     INLINE_ARGS.Struct,
			CapStructField: capStructField,
			StructField:    v,
			FieldType:      PROG_STATE.fieldTypes[v],
			Comment:        PROG_STATE.fieldComments[v],
		}
	}
	cntr = 0
	for _, defaultInfo := range PROG_STATE.fieldDefaults {
		templateData.StructFieldDefaults[cntr] = StructDefaultTemplateVals{
			StructField:        defaultInfo.name,
			StructFieldDefault: defaultInfo.value,
		}
		cntr++
	}

	if err := TEMPLATES.WriteToFile(
		INLINE_ARGS.Struct,
		common.GeneratedSrcFileExt,
		"file",
		templateData,
	); err != nil {
		common.PrintRunningError("%s", err)
		os.Exit(1)
	}
}

func parseTypeSpec(
	fSet *token.FileSet,
	srcFile *os.File,
	ts *ast.TypeSpec,
	PROG_STATE *ProgState,
) bool {
	if ts.Name.Name == "_" {
		return false
	}
	if st, ok := ts.Type.(*ast.StructType); ok && ts.Name.Name == INLINE_ARGS.Struct {
		if st.Fields.List == nil {
			common.PrintRunningError("The supplied options struct has no fields.")
			os.Exit(1)
		}

		for _, field := range st.Fields.List {
			var err error
			var fieldName string
			if len(field.Names) > 0 {
				fieldName, err = common.GetSourceTextFromExpr(
					fSet, srcFile, field.Names[0],
				)
				if err != nil {
					common.PrintRunningError(
						"could not get struct field name: %w",
						err,
					)
					os.Exit(1)
				}
			} else { // Embed type
				fieldName, err = common.GetSourceTextFromExpr(
					fSet, srcFile, field.Type,
				)
				if err != nil {
					common.PrintRunningError(
						"could not get embeded struct field name: %w",
						err,
					)
					os.Exit(1)
				}
				if idx := strings.LastIndex(fieldName, "."); idx != -1 {
					fieldName = fieldName[idx+1:]
				}
			}

			rawFieldTag, err := common.GetSourceTextFromExpr(fSet, srcFile, field.Tag)
			if err != nil {
				common.PrintRunningError(
					"could not get struct field tag: %w",
					err,
				)
				os.Exit(1)
			}
			fieldType, err := common.GetSourceTextFromExpr(fSet, srcFile, field.Type)
			if err != nil {
				common.PrintRunningError(
					"could not get struct field type: %w",
					err,
				)
				os.Exit(1)
			}
			fieldComment, err := common.GetComment(fSet, srcFile, field.Doc, field.Comment)
			if err != nil {
				common.PrintRunningError(
					"could not get struct field type: %w",
					err,
				)
				os.Exit(1)
			}

			PROG_STATE.fieldTypes[fieldName] = fieldType
			PROG_STATE.fieldComments[fieldName] = fieldComment

			// Remove the ticks
			tags := reflect.StructTag(rawFieldTag[1 : len(rawFieldTag)-1])

			if _default, ok := tags.Lookup("default"); !ok {
				common.PrintRunningError("all struct field tags must have a default value")
				os.Exit(1)
			} else if _default == "" {
				common.PrintRunningError("the default tag must not be an empty value")
				os.Exit(1)
			} else {
				PROG_STATE.fieldDefaults = append(
					PROG_STATE.fieldDefaults,
					DefaultInfo{
						name:  fieldName,
						value: _default,
					},
				)
			}

			if auto, ok := tags.Lookup("setter"); !ok {
				common.PrintRunningError("all field tags the must have a setter entry")
				os.Exit(1)
			} else {
				if b, err := strconv.ParseBool(auto); err != nil {
					common.PrintRunningError("the setter tag must be a bool type: %w", err)
					os.Exit(1)
				} else if b {
					PROG_STATE.fieldSetters = append(PROG_STATE.fieldSetters, fieldName)
				}
			}

			if auto, ok := tags.Lookup("getter"); !ok {
				common.PrintRunningError("all field tags the must have a getter entry")
				os.Exit(1)
			} else {
				if b, err := strconv.ParseBool(auto); err != nil {
					common.PrintRunningError("the getter tag must be a bool type: %w", err)
					os.Exit(1)
				} else if b {
					PROG_STATE.fieldGetters = append(PROG_STATE.fieldGetters, fieldName)
				}
			}

			if _import, ok := tags.Lookup("import"); ok {
				PROG_STATE.imports = append(PROG_STATE.imports, _import)
			}
		}

		return true
	}
	return false
}
