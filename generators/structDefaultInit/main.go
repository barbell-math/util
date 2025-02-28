// A generator program that provides boiler plate code for struct initialization
// and basic methods.
package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"strings"

	"github.com/barbell-math/util/generators/common"
)

type (
	InlineArgs struct {
		Struct   string `required:"t" help:"The struct type to generate code for."`
		ShowInfo bool   `required:"f" default:"t" help:"Show debug info."`
	}
	StructArgs struct {
		NewReturns string `required:"t" help:"Controls if the new function for the struct returns a pointer or value."`
	}
	StructFieldArgs struct {
		Default       string `required:"t" help:"The default value for this field. Must be valid go syntax."`
		Setter        bool   `required:"f" default:"f" help:"Controls if a setter is added for this field."`
		Getter        bool   `required:"f" default:"f" help:"Controls if a getter is added for this field."`
		PointerSetter bool   `required:"f" default:"f" help:"Controls if a setter is added for the pointers underlying value for this field."`
		Imports       string `required:"f" default:"" help:"A space separated list of packages to import"`
	}

	ProgState struct {
		shortStructGenerics string
		longStructGenerics  string
		newReturns          string
		fieldSetters        []string
		fieldGetters        []string
		fieldPointerSetters []string
		fieldDefaults       []DefaultInfo
		fieldTypes          map[string]string
		fieldComments       map[string]string
		imports             map[string]struct{}
		_package            string
	}
	DefaultInfo struct {
		name  string
		value string
	}

	TemplateVals struct {
		StructName                string
		CapStructName             string
		ShortStructGenerics       string
		LongStructGenerics        string
		Package                   string
		NewRet                    string
		NewRetVal                 string
		StructFieldSetters        []StructFieldTemplateVals
		StructFieldGetters        []StructFieldTemplateVals
		StructFieldPointerSetters []StructFieldTemplateVals
		StructFieldDefaults       []StructDefaultTemplateVals
		Imports                   []string
		GeneratorName             string
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
		StructName          string
		CapStructField      string
		ShortStructGenerics string
		LongStructGenerics  string
		StructField         string
		FieldType           string
		Comment             string
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
		imports:       map[string]struct{}{},
	}
	TEMPLATES common.GeneratedFilesRegistry = common.NewGeneratedFilesRegistryFromMap(
		map[string]string{
			"structFieldSetter": `
{{ .Comment }}
func (o *{{ .StructName }}{{ .ShortStructGenerics }}) Set{{ .CapStructField }}(v {{ .FieldType }}) *{{ .StructName }}{{ .ShortStructGenerics }} {
	o.{{ .StructField }} = v
	return o
}
`,
			"structFieldPointerSetter": `
{{ .Comment }}
func (o *{{ .StructName }}{{ .ShortStructGenerics }}) Set{{ .CapStructField }}PntrVal(v {{ .FieldType }}) *{{ .StructName }}{{ .ShortStructGenerics }} {
	*o.{{ .StructField }} = *v
	return o
}
`,
			"structFieldGetter": `
{{ .Comment }}
func (o *{{ .StructName }}{{ .ShortStructGenerics }}) Get{{ .CapStructField }}() {{ .FieldType }} {
	return o.{{ .StructField }}
}
`,
			"defaultValueInit": `{{ .StructField }}: {{ .StructFieldDefault }},
`,
			"newFunc": `
// Returns a new {{ .StructName }} struct initialized with the default values.
func New{{ .CapStructName }}{{ .LongStructGenerics }}() {{ .NewRet }}{{ .ShortStructGenerics }} {
	return {{ .NewRetVal }}{{ .ShortStructGenerics }} {
		{{range .StructFieldDefaults}} {{template "defaultValueInit" .}} {{end}}
	}
}
`,
			"file": `
package {{ .Package }}
{{template "autoGenComment" .}}
import (
	{{range .Imports -}}
		"{{ . }}"
	{{end -}}
)

{{template "newFunc" .}}
{{range .StructFieldSetters}}
	{{template "structFieldSetter" .}}
{{end}}
{{range .StructFieldPointerSetters}}
	{{template "structFieldPointerSetter" .}}
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

	requestedStructFound := false

	common.IterateOverAST(
		".",
		common.GenFileExclusionFilter,
		func(fSet *token.FileSet, file *ast.File, srcFile *os.File, node ast.Node) bool {
			switch node.(type) {
			case *ast.GenDecl:
				gdNode := node.(*ast.GenDecl)
				if gdNode.Tok == token.TYPE && !requestedStructFound {
					for _, spec := range gdNode.Specs {
						requestedStructFound = parseTypeSpec(
							fSet,
							srcFile,
							spec.(*ast.TypeSpec),
						)
						if requestedStructFound {
							break
						}
					}
					if requestedStructFound {
						PROG_STATE._package = file.Name.Name
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

	if !requestedStructFound {
		common.PrintRunningError(
			"The supplied type was not found or was not a struct but is required to be.",
		)
		os.Exit(1)
	}

	templateData := TemplateVals{
		StructName:                INLINE_ARGS.Struct,
		CapStructName:             strings.ToUpper(INLINE_ARGS.Struct[0:1]) + INLINE_ARGS.Struct[1:],
		ShortStructGenerics:       PROG_STATE.shortStructGenerics,
		LongStructGenerics:        PROG_STATE.longStructGenerics,
		Package:                   PROG_STATE._package,
		StructFieldSetters:        make([]StructFieldTemplateVals, len(PROG_STATE.fieldSetters)),
		StructFieldPointerSetters: make([]StructFieldTemplateVals, len(PROG_STATE.fieldPointerSetters)),
		StructFieldGetters:        make([]StructFieldTemplateVals, len(PROG_STATE.fieldGetters)),
		StructFieldDefaults:       make([]StructDefaultTemplateVals, len(PROG_STATE.fieldDefaults)),
		GeneratorName:             os.Args[0],
	}
	if PROG_STATE.newReturns == "val" {
		templateData.NewRet = INLINE_ARGS.Struct
		templateData.NewRetVal = INLINE_ARGS.Struct
	} else if PROG_STATE.newReturns == "pntr" {
		templateData.NewRet = fmt.Sprintf("*%s", INLINE_ARGS.Struct)
		templateData.NewRetVal = fmt.Sprintf("&%s", INLINE_ARGS.Struct)
	}

	for k, _ := range PROG_STATE.imports {
		templateData.Imports = append(templateData.Imports, k)
	}

	for i, v := range PROG_STATE.fieldSetters {
		capStructField := strings.ToUpper(v[0:1]) + v[1:]
		templateData.StructFieldSetters[i] = StructFieldTemplateVals{
			StructName:          INLINE_ARGS.Struct,
			CapStructField:      capStructField,
			ShortStructGenerics: templateData.ShortStructGenerics,
			StructField:         v,
			FieldType:           PROG_STATE.fieldTypes[v],
			Comment:             PROG_STATE.fieldComments[v],
		}
	}
	for i, v := range PROG_STATE.fieldGetters {
		capStructField := strings.ToUpper(v[0:1]) + v[1:]
		templateData.StructFieldGetters[i] = StructFieldTemplateVals{
			StructName:          INLINE_ARGS.Struct,
			CapStructField:      capStructField,
			ShortStructGenerics: templateData.ShortStructGenerics,
			StructField:         v,
			FieldType:           PROG_STATE.fieldTypes[v],
			Comment:             PROG_STATE.fieldComments[v],
		}
	}
	for i, v := range PROG_STATE.fieldPointerSetters {
		capStructField := strings.ToUpper(v[0:1]) + v[1:]
		templateData.StructFieldPointerSetters[i] = StructFieldTemplateVals{
			StructName:          INLINE_ARGS.Struct,
			CapStructField:      capStructField,
			ShortStructGenerics: templateData.ShortStructGenerics,
			StructField:         v,
			FieldType:           PROG_STATE.fieldTypes[v],
			Comment:             PROG_STATE.fieldComments[v],
		}
	}

	cntr := 0
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
) bool {
	if ts.Name.Name == "_" {
		return false
	}

	if st, ok := ts.Type.(*ast.StructType); ok && ts.Name.Name == INLINE_ARGS.Struct {
		commentArgs, err := common.GetDocArgVals(fSet, srcFile, ts.Doc)
		if err != nil {
			common.PrintRunningError("%s", err)
			os.Exit(1)
		}
		structArgs := StructArgs{}
		if err := common.CommentArgs(&structArgs, commentArgs); err != nil {
			common.PrintRunningError("%s", err)
			os.Exit(1)
		}
		PROG_STATE.newReturns = structArgs.NewReturns
		if PROG_STATE.newReturns != "pntr" && PROG_STATE.newReturns != "val" {
			common.PrintRunningError("The newReturns argument must be either 'pntr' or 'val'")
			os.Exit(1)
		}

		if st.Fields.List == nil {
			common.PrintRunningError("The supplied options struct has no fields.")
			os.Exit(1)
		}

		if PROG_STATE.shortStructGenerics, err = common.GetShortGenericsString(
			fSet, srcFile, ts,
		); err != nil {
			common.PrintRunningError("could not get generic type: %w", err)
			os.Exit(1)
		}
		if PROG_STATE.longStructGenerics, err = common.GetLongGenericsString(
			fSet, srcFile, ts,
		); err != nil {
			common.PrintRunningError("could not get generic type: %w", err)
			os.Exit(1)
		}

		setFieldVals(fSet, srcFile, st)

		return true
	}
	return false
}

func setFieldVals(
	fSet *token.FileSet,
	srcFile *os.File,
	st *ast.StructType,
) {
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

		commentArgs, err := common.GetDocArgVals(fSet, srcFile, field.Doc)
		if err != nil {
			common.PrintRunningError("%s", err)
			os.Exit(1)
		}
		structArgs := StructFieldArgs{}
		if err := common.CommentArgs(&structArgs, commentArgs); err != nil {
			common.PrintRunningError("%s", err)
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

		if structArgs.Default == "" {
			common.PrintRunningError("the default value must not be an empty value")
			os.Exit(1)
		} else {
			PROG_STATE.fieldDefaults = append(
				PROG_STATE.fieldDefaults,
				DefaultInfo{
					name:  fieldName,
					value: structArgs.Default,
				},
			)
		}

		if structArgs.Setter {
			PROG_STATE.fieldSetters = append(PROG_STATE.fieldSetters, fieldName)
		}
		if structArgs.Getter {
			PROG_STATE.fieldGetters = append(PROG_STATE.fieldGetters, fieldName)
		}
		if structArgs.PointerSetter {
			PROG_STATE.fieldPointerSetters = append(PROG_STATE.fieldPointerSetters, fieldName)
		}

		if len(structArgs.Imports) > 0 {
			for _, i := range strings.Split(structArgs.Imports, " ") {
				PROG_STATE.imports[i] = struct{}{}
			}
		}
	}
}
