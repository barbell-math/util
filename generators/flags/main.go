// A generator program that creates methods surrounding a bit-flag enum type.
package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"strings"

	"github.com/barbell-math/util/src/generators/common"
)

type (
	InlineArgs struct {
		Type     string `required:"t" help:"The type that is used to represent the flag values."`
		Package  string `required:"t" help:"The package the options enum type is in."`
		ShowInfo bool   `required:"f" default:"t" help:"Show debug info."`
	}

	ProgState struct {
		prevType string
		enumVars []EnumVar
	}
	EnumVar struct {
		name     string
		comment  string
		NoSetter bool `required:"f" default:"f" help:"Set to true to not generate the setter code for this flag."`
	}

	TemplateVals struct {
		GeneratorName   string
		Package         string
		EnumType        string
		CapEnumType     string
		AllCapsEnumType string
		EnumFlags       []EnumFlagTemplateVals
		EnumFlagSetters []EnumFlagTemplateVals
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
	INLINE_ARGS InlineArgs
	PROG_STATE  ProgState = ProgState{
		enumVars: []EnumVar{},
	}
	TEMPLATES common.GeneratedFilesRegistry = common.NewGeneratedFilesRegistryFromMap(
		map[string]string{
			"flagOptionFunc": `
{{ .Comment }}
func (o {{ .EnumType }}) {{ .CapEnumFlag }}(b bool) {{ .EnumType }} {
	if b {
		o |= {{ .EnumFlag }}
	} else {
		o &= ^{{ .EnumFlag }}
	}
	return o
}
`,
			"getFlagFunc": `
// Returns the supplied flags status
func (o {{ .EnumType }}) GetFlag(flag {{ .EnumType }}) bool {
	return o & flag>0
}
`,
			"file": `
package {{ .Package }}
{{template "autoGenComment" .}}

{{template "getFlagFunc" .}}
{{range .EnumFlagSetters}}
	{{template "flagOptionFunc" .}}
{{end}}
`,
		},
	)
)

func main() {
	common.InlineArgs(&INLINE_ARGS, os.Args)

	common.IterateOverAST(
		".",
		common.GenFileExclusionFilter,
		func(fSet *token.FileSet, file *ast.File, srcFile *os.File, node ast.Node) bool {
			if file.Name.Name != INLINE_ARGS.Package {
				return false
			}
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

	allCapsEnumType := getAllCapsType()
	templateData := TemplateVals{
		GeneratorName:   os.Args[0],
		Package:         INLINE_ARGS.Package,
		EnumType:        INLINE_ARGS.Type,
		CapEnumType:     strings.ToUpper(INLINE_ARGS.Type[0:1]) + INLINE_ARGS.Type[1:],
		AllCapsEnumType: allCapsEnumType,
		EnumFlags:       make([]EnumFlagTemplateVals, len(PROG_STATE.enumVars)),
		EnumFlagSetters: make([]EnumFlagTemplateVals, 0),
	}
	cntr := 0
	for _, v := range PROG_STATE.enumVars {
		capEnumFlag := strings.ToUpper(v.name[0:1]) + v.name[1:]
		iterV := EnumFlagTemplateVals{
			EnumType:    INLINE_ARGS.Type,
			CapEnumFlag: capEnumFlag,
			EnumFlag:    v.name,
			Comment:     v.comment,
		}
		templateData.EnumFlags[cntr] = iterV
		if !v.NoSetter {
			templateData.EnumFlagSetters = append(
				templateData.EnumFlagSetters, iterV,
			)
		}
		cntr++
	}

	if err := TEMPLATES.WriteToFile(
		fmt.Sprintf("%sFlags", INLINE_ARGS.Type),
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
