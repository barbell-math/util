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
		TypeToCheck string `required:"t" help:"The type that should implement the requested interface."`
		ShowInfo    bool   `required:"f" default:"t" help:"Show debug info."`
	}
	CommentArgs struct {
		IfaceName string `required:"t" help:"The interface the type should implement."`
		Imports   string `required:"f" default:"" help:"The package the required interface is defined in."`
		ValOrPntr string `required:"f" default:"val" help:"Whether the type, a pointer to the type, or both implement the requested interface."`
		Generics  string `required:"f" default:"" help:"The generic parameter values that should be used for the type in the unit tests."`
	}
	ProgState struct {
		Package    string
		Imports    map[string]struct{}
		ValOfType  bool
		PntrToType bool
	}
	TemplateVals struct {
		Package           string
		Imports           []string
		IfaceName         string
		FuncNameIfaceName string
		TypeName          string
		FuncNameTypeName  string
		ValOfType         bool
		PntrToType        bool
		GeneratorName     string
	}
)

var (
	INLINE_ARGS  InlineArgs
	COMMENT_ARGS CommentArgs
	PROG_STATE   ProgState = ProgState{
		Imports: map[string]struct{}{},
	}
	TEMPLATES common.GeneratedFilesRegistry = common.NewGeneratedFilesRegistryFromMap(
		map[string]string{
			"valImpl": `
func Test{{ .FuncNameTypeName }}ValueImplements{{ .FuncNameIfaceName }}(t *testing.T) {
	var typeThing {{ .TypeName }}
	var iFaceThing {{ .IfaceName }} = typeThing
	_ = iFaceThing
}
`,
			"pntrImpl": `
func Test{{ .FuncNameTypeName }}PntrImplements{{ .FuncNameIfaceName }}(t *testing.T) {
	var typeThing {{ .TypeName }}
	var iFaceThing {{ .IfaceName }} = &typeThing
	_ = iFaceThing
}
`,
			"file": `
package {{ .Package }}
{{template "autoGenComment" .}}
import (
	"testing"
	{{range .Imports -}}
		{{ . }}
	{{end -}}
)

{{ if .ValOfType }}
	{{template "valImpl" .}}
{{ end }}

{{ if .PntrToType }}
	{{template "pntrImpl" .}}
{{ end }}
`,
		},
	)
)

func main() {
	common.InlineArgs(&INLINE_ARGS, os.Args)

	requestedTypeFound := false

	common.IterateOverAST(
		".",
		common.GenFileExclusionFilter,
		func(fSet *token.FileSet, file *ast.File, srcFile *os.File, node ast.Node) bool {
			switch node.(type) {
			case *ast.GenDecl:
				gdNode := node.(*ast.GenDecl)
				if gdNode.Tok == token.TYPE && !requestedTypeFound {
					for _, spec := range gdNode.Specs {
						requestedTypeFound = parseTypeSpec(
							fSet,
							srcFile,
							spec.(*ast.TypeSpec),
						)
						if requestedTypeFound {
							break
						}
					}
					if requestedTypeFound {
						PROG_STATE.Package = file.Name.Name
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

	if !requestedTypeFound {
		common.PrintRunningError(
			"The supplied type was not found or was not a struct but is required to be.",
		)
		os.Exit(1)
	}

	templateData := TemplateVals{
		Package:   PROG_STATE.Package,
		Imports:   make([]string, len(PROG_STATE.Imports)),
		IfaceName: COMMENT_ARGS.IfaceName,
		FuncNameIfaceName: common.CapitilizeName(common.CleanName(
			COMMENT_ARGS.IfaceName,
		)),
		TypeName: INLINE_ARGS.TypeToCheck + COMMENT_ARGS.Generics,
		FuncNameTypeName: common.CapitilizeName(common.CleanName(
			INLINE_ARGS.TypeToCheck,
		)),
		ValOfType:     PROG_STATE.ValOfType,
		PntrToType:    PROG_STATE.PntrToType,
		GeneratorName: os.Args[0],
	}

	cntr := 0
	for _import, _ := range PROG_STATE.Imports {
		templateData.Imports[cntr] = _import
		cntr++
	}

	if err := TEMPLATES.WriteToFile(
		fmt.Sprintf(
			"%sImpls%sTest",
			templateData.FuncNameTypeName,
			templateData.FuncNameIfaceName,
		),
		common.GeneratedTestFileExt,
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
	if ts.Name.Name != INLINE_ARGS.TypeToCheck {
		return false
	}

	rawCommentArgs, err := common.GetDocArgVals(fSet, srcFile, ts.Doc)
	if err != nil {
		common.PrintRunningError("%s", err)
		os.Exit(1)
	}
	if err := common.CommentArgs(&COMMENT_ARGS, rawCommentArgs); err != nil {
		common.PrintRunningError("%s", err)
		os.Exit(1)
	}

	if COMMENT_ARGS.ValOrPntr == "val" {
		PROG_STATE.ValOfType = true
		PROG_STATE.PntrToType = false
	} else if COMMENT_ARGS.ValOrPntr == "pntr" {
		PROG_STATE.ValOfType = false
		PROG_STATE.PntrToType = true
	} else if COMMENT_ARGS.ValOrPntr == "both" {
		PROG_STATE.ValOfType = true
		PROG_STATE.PntrToType = true
	} else {
		common.PrintRunningError(fmt.Sprintf(
			"The value supplied to valOrPntr must be one of: [val, pntr, both] Got: %s",
			COMMENT_ARGS.ValOrPntr,
		))
		os.Exit(1)
	}

	common.ParseImports(PROG_STATE.Imports, COMMENT_ARGS.Imports)
	return true
}
