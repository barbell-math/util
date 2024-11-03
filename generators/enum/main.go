// A generator program that creates methods surrounding a enum type.
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
		Type     string `required:"t" help:"The type that is used to represent the enum values."`
		Package  string `required:"t" help:"The package the options enum type is in."`
		ShowInfo bool   `required:"f" default:"t" help:"Show debug info."`
	}

	CommentArgs struct {
		UnknownValue string `required:"t" help:"The value that will be returned should an invalid enum type be used."`
		Default      string `required:"t" help:"The default value the new function should return."`
	}

	ProgState struct {
		prevType string
		enumVars []EnumVar
	}
	EnumVar struct {
		name    string
		comment string
		String  string `required:"t" help:"The string representation of this flag. Used for json marshaling and unmarshaling."`
	}

	TemplateVals struct {
		GeneratorName      string
		Package            string
		EnumType           string
		CapEnumType        string
		AllCapsEnumType    string
		EnumFlags          []EnumFlagTemplateVals
		EnumFlagSetters    []EnumFlagTemplateVals
		UnknownValue       string
		UnknownValueString string
		Default            string
	}
	EnumFlagTemplateVals struct {
		EnumType       string
		CapEnumFlag    string
		EnumFlag       string
		EnumFlagString string
		Comment        string
	}
)

var (
	INLINE_ARGS  InlineArgs
	COMMENT_ARGS CommentArgs
	PROG_STATE   ProgState = ProgState{
		enumVars: []EnumVar{},
	}
	TEMPLATES common.GeneratedFilesRegistry = common.NewGeneratedFilesRegistryFromMap(
		map[string]string{
			"validFunc": `
func (o {{ .EnumType }}) Valid() error {
	switch o {
	{{range .EnumFlags }}
			case {{ .EnumFlag }}: return nil
	{{end}}
	default: return Invalid{{ .CapEnumType }}
	}
}
`,
			"marshalJSONFunc": `
func (o {{ .EnumType }}) MarshalJSON() ([]byte, error) {
	switch o {
	{{range .EnumFlags }}
			case {{ .EnumFlag }}: return []byte("{{ .EnumFlagString}}"), nil
	{{end}}
	default: return []byte("{{ .UnknownValueString }}"), Invalid{{ .CapEnumType }}
	}
}
`,
			"stringFunc": `
func (o {{ .EnumType }}) String() string {
	switch o {
	{{range .EnumFlags }} case {{ .EnumFlag }}: return "{{ .EnumFlagString }}"
	{{end}}
	default: return "{{ .UnknownValueString }}"
	}
}
`,
			"unmarshalJSONFunc": `
func (o *{{ .EnumType }}) UnmarshalJSON(b []byte) error {
	switch string(b) {
	{{range .EnumFlags}}
	case "{{ .EnumFlagString }}":
		*o={{ .EnumFlag }}
		return nil
	{{end}}
	default:
		*o={{ .UnknownValue }}
		return fmt.Errorf("%w: %s",Invalid{{ .CapEnumType }}, string(b))
	}
}
`,
			"fromStringFunc": `
func (o *{{ .EnumType }}) FromString(s string) error {
	switch s {
	{{range .EnumFlags}}
	case "{{ .EnumFlagString }}":
		*o={{ .EnumFlag }}
		return nil
	{{end}}
	default:
		*o={{ .UnknownValue }}
		return fmt.Errorf("%w: %s",Invalid{{ .CapEnumType }}, s)
	}
}
`,
			"defaultFunc": `
func New{{ .CapEnumType }}() {{ .EnumType }} {
	return {{ .Default }}
}
`,
			"file": `
package {{ .Package }}
{{template "autoGenComment" .}}
import (
	"fmt"
	"errors"
)

var (
	Invalid{{ .CapEnumType }}=errors.New("Invalid {{ .EnumType }}")
	{{ .AllCapsEnumType }} []{{ .EnumType }}=[]{{ .EnumType }}{
		{{range .EnumFlags}} {{.EnumFlag}},
		{{end}}
	}
)

{{template "defaultFunc" .}}
{{template "validFunc" .}}
{{template "stringFunc" .}}
{{template "marshalJSONFunc" .}}
{{template "fromStringFunc" .}}
{{template "unmarshalJSONFunc" .}}
`,
		},
	)
)

func main() {
	common.InlineArgs(&INLINE_ARGS, os.Args)

	op, found := common.DocArgsAstFilter(INLINE_ARGS.Type, &COMMENT_ARGS)
	common.IterateOverAST(
		".",
		common.GenFileExclusionFilter,
		func(fSet *token.FileSet, file *ast.File, srcFile *os.File, node ast.Node) bool {
			if file.Name.Name != INLINE_ARGS.Package {
				return false
			}
			op(fSet, srcFile, node)
			switch node.(type) {
			case *ast.GenDecl:
				gdNode := node.(*ast.GenDecl)
				if gdNode.Tok == token.CONST {
					PROG_STATE.prevType = ""
					for _, spec := range gdNode.Specs {
						PROG_STATE.prevType = parseValueSpec(
							fSet,
							srcFile,
							spec.(*ast.ValueSpec),
						)
					}
				}
				return false
			case *ast.FuncDecl:
				return false
			default:
				return true
			}
		},
	)

	if !*found {
		common.PrintRunningError(
			"The supplied type did not have comment args but they are required.",
		)
		os.Exit(1)
	}

	foundUnknownValue := false
	for _, v := range PROG_STATE.enumVars {
		foundUnknownValue = (v.name == COMMENT_ARGS.UnknownValue)
		if foundUnknownValue {
			break
		}
	}
	if !foundUnknownValue {
		common.PrintRunningError(
			"The supplied unknown value was not found in the list of enum flag values.",
		)
		os.Exit(1)
	}

	allCapsEnumType := getAllCapsType()
	templateData := TemplateVals{
		GeneratorName:   os.Args[0],
		Package:         INLINE_ARGS.Package,
		EnumType:        INLINE_ARGS.Type,
		CapEnumType:     strings.ToUpper(INLINE_ARGS.Type[0:1]) + INLINE_ARGS.Type[1:],
		AllCapsEnumType: allCapsEnumType,
		EnumFlags:       make([]EnumFlagTemplateVals, len(PROG_STATE.enumVars)),
		EnumFlagSetters: make([]EnumFlagTemplateVals, 0),
		Default:         COMMENT_ARGS.Default,
	}
	cntr := 0
	for _, v := range PROG_STATE.enumVars {
		capEnumFlag := strings.ToUpper(v.name[0:1]) + v.name[1:]
		iterV := EnumFlagTemplateVals{
			EnumType:       INLINE_ARGS.Type,
			CapEnumFlag:    capEnumFlag,
			EnumFlag:       v.name,
			Comment:        v.comment,
			EnumFlagString: v.String,
		}
		templateData.EnumFlags[cntr] = iterV
		if v.name == COMMENT_ARGS.UnknownValue {
			templateData.UnknownValue = v.name
			templateData.UnknownValueString = v.String
		}
		cntr++
	}

	if err := TEMPLATES.WriteToFile(
		fmt.Sprintf("%sEnum", INLINE_ARGS.Type),
		common.GeneratedSrcFileExt,
		"file",
		templateData,
	); err != nil {
		common.PrintRunningError("%s", err)
		os.Exit(1)
	}
}

func parseValueSpec(
	fSet *token.FileSet,
	srcFile *os.File,
	vs *ast.ValueSpec,
) string {
	getType := func() (string, error) {
		var err error
		rv := PROG_STATE.prevType
		if vs.Type != nil {
			rv, err = common.GetSourceTextFromExpr(
				fSet, srcFile, vs.Type,
			)
		} else if vs.Type == nil && PROG_STATE.prevType != "" {
			rv = PROG_STATE.prevType
		}
		return rv, err
	}

	if len(vs.Names) > 0 && vs.Names[0].Name == "_" {
		return PROG_STATE.prevType
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

	if len(vs.Names) <= 0 || iterType != INLINE_ARGS.Type {
		return iterType
	}

	ev := EnumVar{
		name:    vs.Names[0].Name,
		comment: comment,
	}
	commentArgs, err := common.GetDocArgVals(fSet, srcFile, vs.Doc)
	if err != nil {
		common.PrintRunningError("%s", err)
		os.Exit(1)
	}
	if err := common.CommentArgs(&ev, commentArgs); err != nil {
		common.PrintRunningError("%s", err)
		os.Exit(1)
	}

	PROG_STATE.enumVars = append(PROG_STATE.enumVars, ev)

	return iterType
}

func getAllCapsType() string {
	prevIndex := 0
	allCapsEnumType := ""
	for i, v := range INLINE_ARGS.Type {
		if strings.ToUpper(string(v)) == string(v) {
			allCapsEnumType += "_"
			allCapsEnumType += strings.ToUpper(INLINE_ARGS.Type[prevIndex:i])
			prevIndex = i
		}
	}
	allCapsEnumType += "_"
	allCapsEnumType += strings.ToUpper(INLINE_ARGS.Type[prevIndex:])
	if allCapsEnumType[0] == '_' {
		allCapsEnumType = allCapsEnumType[1:]
	}
	return allCapsEnumType
}
