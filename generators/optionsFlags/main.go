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
	Values struct {
		OptionsStruct string `required:"t" help:"The struct type that holds the options."`
		OptionsEnum   string `required:"t" help:"The type that holds the flags."`
		Package       string `required:"t" help:"The package the options enum type is in."`
		Debug         bool   `required:"f" default:"f" help:"Print diagnostic information to the console."`
		ShowInfo      bool   `required:"f" default:"t" help:"Show debug info."`
	}

	ProgState struct {
		// Enum type values
		prevType string
		enumVars map[string]string

		// Struct field values
		autoFields    []string
		defaults      map[string]string
		fieldTypes    map[string]string
		fieldComments map[string]string
		imports []string
	}

	TemplateVals struct {
		OptionsStruct string
		CapOptionsStruct string
		OptionsEnum   string
		Package       string
		EnumFlags     []EnumFlagTemplateVals
		StructFields  []StructFieldTemplateVals
		StructFieldDefaults []StructDefaultTemplateVals
		Imports []string
		GeneratorName string
	}
	EnumFlagTemplateVals struct {
		OptionsStruct string
		CapEnumFlag   string
		EnumFlag      string
		Comment       string
	}
	StructDefaultTemplateVals struct {
		StructField string
		StructFieldDefault string
	}
	StructFieldTemplateVals struct {
		OptionsStruct  string
		CapStructField string
		StructField    string
		FieldType      string
		Comment        string
	}
)

const (
	IgnoreEnumValue string = "//optionsFlags ignore"
)

var (
	VALS      Values
	TEMPLATES common.GeneratedFilesRegistry = common.NewGeneratedFilesRegistryFromMap(
		map[string]string{
			"flagOptionFunc": `
{{ .Comment }}
func (o *{{ .OptionsStruct }}) {{ .CapEnumFlag }}(b bool) *{{ .OptionsStruct }} {
	if b {
		o.flags |= {{ .EnumFlag }}
	} else {
		o.flags &= ^{{ .EnumFlag }}
	}
	return o
}
`,
			"getFlagFunc": `
func (o *{{ .OptionsStruct }}) getFlag(flag {{ .OptionsEnum }}) bool {
	return o.flags & flag>0
}
`,
			"structFieldFunc": `
{{ .Comment }}
func (o *{{ .OptionsStruct }}) {{ .CapStructField }}(v {{ .FieldType }}) *{{ .OptionsStruct }} {
	o.{{ .StructField }} = v
	return o
}
`,
			"defaultValueInit": `{{ .StructField }}: {{ .StructFieldDefault }},
`,
			"newFunc": `
// Returns a new {{ .OptionsStruct }} struct initialized with the default values.
func New{{ .CapOptionsStruct }}() *{{ .OptionsStruct }} {
	return &{{ .OptionsStruct }} {
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
{{template "getFlagFunc" .}}
{{range .EnumFlags}}
	{{template "flagOptionFunc" .}}
{{end}}
{{range .StructFields}}
	{{template "structFieldFunc" .}}
{{end}}
`,
		},
	)
)

func main() {
	common.Args(&VALS, os.Args)

	optionsStructFound := false
	progState := ProgState{
		enumVars:      map[string]string{},
		autoFields:    []string{},
		defaults:      map[string]string{},
		fieldTypes:    map[string]string{},
		fieldComments: map[string]string{},
		imports: []string{},
	}

	common.IterateOverAST(
		".",
		common.GenFileExclusionFilter,
		func(fSet *token.FileSet, file *ast.File, srcFile *os.File, node ast.Node) bool {
			switch node.(type) {
			case *ast.GenDecl:
				gdNode := node.(*ast.GenDecl)
				if gdNode.Tok == token.CONST {
					for _, spec := range gdNode.Specs {
						progState.prevType = parseValueSpec(
							fSet,
							srcFile,
							spec.(*ast.ValueSpec),
							&progState,
						)
					}
				} else if gdNode.Tok == token.TYPE && !optionsStructFound {
					gdNode := node.(*ast.GenDecl)
					for _, spec := range gdNode.Specs {
						optionsStructFound = parseTypeSpec(
							fSet,
							srcFile,
							spec.(*ast.TypeSpec),
							&progState,
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
		OptionsStruct: VALS.OptionsStruct,
		CapOptionsStruct: 
			strings.ToUpper(VALS.OptionsStruct[0:1])+VALS.OptionsStruct[1:],
		OptionsEnum:   VALS.OptionsEnum,
		Package:       VALS.Package,
		EnumFlags:     make([]EnumFlagTemplateVals, len(progState.enumVars)),
		StructFields:  make([]StructFieldTemplateVals, len(progState.autoFields)),
		StructFieldDefaults: make([]StructDefaultTemplateVals, len(progState.defaults)),
		Imports: progState.imports,
		GeneratorName: os.Args[0],
	}
	cntr := 0
	for e, c := range progState.enumVars {
		capEnumFlag := strings.ToUpper(e[0:1]) + e[1:]
		templateData.EnumFlags[cntr] = EnumFlagTemplateVals{
			OptionsStruct: VALS.OptionsStruct,
			CapEnumFlag:   capEnumFlag,
			EnumFlag:      e,
			Comment:       c,
		}
		cntr++
	}
	for i, v := range progState.autoFields {
		capStructField := strings.ToUpper(v[0:1]) + v[1:]
		templateData.StructFields[i] = StructFieldTemplateVals{
			OptionsStruct:  VALS.OptionsStruct,
			CapStructField: capStructField,
			StructField:    v,
			FieldType:      progState.fieldTypes[v],
			Comment:        progState.fieldComments[v],
		}
	}
	cntr=0
	for field, _default:=range progState.defaults {
		templateData.StructFieldDefaults[cntr]=StructDefaultTemplateVals{
			StructField: field,
			StructFieldDefault: _default,
		}
		cntr++
	}

	if err := TEMPLATES.WriteToFile(
		VALS.OptionsStruct,
		"file",
		templateData,
	); err != nil {
		common.PrintRunningError("%s", err)
	}
}

func parseValueSpec(
	fSet *token.FileSet,
	srcFile *os.File,
	vs *ast.ValueSpec,
	progState *ProgState,
) string {
	getType := func() (string, error) {
		var err error
		rv := progState.prevType
		if vs.Type != nil {
			rv, err = common.GetSourceTextFromExpr(
				fSet, srcFile, vs.Type,
			)
		} else if vs.Type == nil && progState.prevType != "" {
			rv = progState.prevType
		}
		return rv, err
	}

	if len(vs.Names) > 0 && vs.Names[0].Name == "_" {
		return progState.prevType
	}

	iterType, err := getType()
	if err != nil {
		common.PrintRunningError("%s", err)
		os.Exit(1)
	}
	comment, err := common.GetComment(fSet, srcFile, vs.Doc, vs.Comment)
	if err != nil {
		common.PrintRunningError("%s", err)
		os.Exit(1)
	}

	if len(vs.Names) > 0 &&
		iterType == VALS.OptionsEnum &&
		comment != IgnoreEnumValue {
		progState.enumVars[vs.Names[0].Name] = comment
	}

	return iterType
}

func parseTypeSpec(
	fSet *token.FileSet,
	srcFile *os.File,
	ts *ast.TypeSpec,
	progState *ProgState,
) bool {
	if ts.Name.Name == "_" {
		return false
	}
	if st, ok := ts.Type.(*ast.StructType); ok && ts.Name.Name == VALS.OptionsStruct {
		if st.Fields.List == nil {
			common.PrintRunningError("The supplied options struct has no fields.")
			os.Exit(1)
		}

		for _, field := range st.Fields.List {
			fieldName, err := common.GetSourceTextFromExpr(fSet, srcFile, field.Names[0])
			if err != nil {
				common.PrintRunningError(
					"could not get struct field name: %w",
					err,
				)
				os.Exit(1)
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

			progState.fieldTypes[fieldName] = fieldType
			progState.fieldComments[fieldName] = fieldComment

			// Remove the ticks
			tags := reflect.StructTag(rawFieldTag[1 : len(rawFieldTag)-1])

			if _default, ok := tags.Lookup("default"); !ok {
				common.PrintRunningError("all struct field tags must have a default value")
				os.Exit(1)
			} else if _default == "" {
				common.PrintRunningError("the default tag must not be an empty value")
				os.Exit(1)
			} else {
				progState.defaults[fieldName] = _default
			}

			auto, ok := tags.Lookup("auto")
			if fieldName == "flags" && ok {
				common.PrintRunningError("the flags field tag should not have an auto entry")
				os.Exit(1)
			} else if fieldName != "flags" && !ok {
				common.PrintRunningError("all field tags other than the flags field must have an auto entry")
				os.Exit(1)
			} else if fieldName != "flags" && ok {
				if b, err := strconv.ParseBool(auto); err != nil {
					common.PrintRunningError("the auto tag must be a bool type: %w", err)
					os.Exit(1)
				} else if b {
					progState.autoFields = append(progState.autoFields, fieldName)
				}
			}

			if _import,ok:=tags.Lookup("import"); ok {
				progState.imports = append(progState.imports, _import)
			}
		}

		return true
	}
	return false
}
