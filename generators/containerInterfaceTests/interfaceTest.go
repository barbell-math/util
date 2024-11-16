// This is a generator program that is used internally by the container/containers
// package to generate unit tests for containers.
package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"regexp"
	"strings"

	"github.com/barbell-math/util/src/generators/common"
)

type (
	category byte

	InlineArgs struct {
		Type        string `required:"t" help:"The underlying type to generate the widget for."`
		Interface   string `required:"t" help:"The packge to put the files in."`
		GenericDecl string `required:"t" help:"The generic type signature to use."`
		Factory     string `required:"t" help:"The factory that will produce containers to test."`
		Category    string `required:"t" help:"Either static or dynamic."`
		CapType     string `required:"f" default:"" help:"The type but the first letter is capitilized. This will be calculated if left blank."`
		Debug       bool   `required:"f" default:"false" help:"Print diagonistic information to the console."`
		ShowInfo    bool   `required:"f" default:"t" help:"Show debug info."`
	}

	ProgState struct {
		Cat         category
		ViableFuncs []*ast.FuncDecl
	}
	TemplateVals struct {
		Cat           string
		Type          string
		Interface     string
		GenericDecl   string
		FuncNames     []FuncTemplateVals
		GeneratorName string
		Factory       string
	}
	FuncTemplateVals struct {
		Name      string
		Type      string
		Interface string
	}
)

const (
	Static category = iota
	Dynamic
	Unknown

	FirstParamName  string = "factory"
	SecondParamName string = "t"
)

var (
	INLINE_ARGS   InlineArgs
	REQUIRED_ARGS []string = []string{
		"type",
		"category",
		"interface",
		"factory",
		"genericDecl",
	}
	PROG_STATE ProgState = ProgState{
		ViableFuncs: []*ast.FuncDecl{},
	}
	TEMPLATES common.GeneratedFilesRegistry = common.NewGeneratedFilesRegistryFromMap(
		map[string]string{
			"funcTemplate": `
func Test{{ .Type }}_{{ .Name }}(t *testing.T){
	tests.{{ .Name }}({{ .Type }}To{{ .Interface }}InterfaceFactory,t)
}`,
			"file": `
package containers

{{template "autoGenComment" .}}
import (
	"testing"
	"github.com/barbell-math/util/src/container/tests"
	"github.com/barbell-math/util/src/container/{{ .Cat }}Containers"
)

func {{ .Type }}To{{ .Interface }}InterfaceFactory(capacity int) {{ .Cat }}Containers.{{ .Interface }}{{ .GenericDecl }} {
	v:= {{ .Factory }}(capacity)
	var rv {{ .Cat }}Containers.{{ .Interface }}{{ .GenericDecl }}=&v
	return rv
}
{{range .FuncNames}}
	{{template "funcTemplate" .}}
{{end}}
`,
		},
	)
)

func (c category) String() string {
	switch c {
	case Static:
		return "static"
	case Dynamic:
		return "dynamic"
	default:
		return ""
	}
}

func (c category) FromString(s string) category {
	switch strings.ToLower(s) {
	case "dynamic":
		return Dynamic
	case "static":
		return Static
	default:
		return Unknown
	}
}

func main() {
	common.InlineArgs(&INLINE_ARGS, os.Args)
	checkGenericDecl()
	INLINE_ARGS.CapType = strings.ToUpper(INLINE_ARGS.Type)[:1] + INLINE_ARGS.Type[1:]
	PROG_STATE.Cat = Dynamic.FromString(INLINE_ARGS.Category)

	common.IterateOverAST(
		"../tests/",
		common.GenFileExclusionFilter,
		func(fSet *token.FileSet, file *ast.File, srcFile *os.File, node ast.Node) bool {
			switch node.(type) {
			case *ast.FuncDecl:
				fd := node.(*ast.FuncDecl)
				if INLINE_ARGS.Debug {
					fmt.Print("Found func: ", fd.Name, "| Status: ")
				}
				if ok, info := hasViableSignature(fSet, srcFile, fd); ok {
					if INLINE_ARGS.Debug {
						fmt.Println("Accepted")
					}
					PROG_STATE.ViableFuncs = append(PROG_STATE.ViableFuncs, fd)
				} else if INLINE_ARGS.Debug {
					fmt.Println("Rejected | Reason:", info)
				}
				return false
			case *ast.GenDecl:
				return false
			default:
				return true
			}
		},
	)
	common.PrintRunningInfo("Num Funcs: %3d", len(PROG_STATE.ViableFuncs))

	templateData := TemplateVals{
		Cat:           PROG_STATE.Cat.String(),
		Type:          INLINE_ARGS.Type,
		Interface:     INLINE_ARGS.Interface,
		GenericDecl:   INLINE_ARGS.GenericDecl,
		FuncNames:     make([]FuncTemplateVals, len(PROG_STATE.ViableFuncs)),
		GeneratorName: os.Args[0],
		Factory:       INLINE_ARGS.Factory,
	}
	for i, f := range PROG_STATE.ViableFuncs {
		templateData.FuncNames[i] = FuncTemplateVals{
			Name:      f.Name.Name,
			Type:      INLINE_ARGS.Type,
			Interface: INLINE_ARGS.Interface,
		}
	}

	if err := TEMPLATES.WriteToFile(
		fileName(),
		common.GeneratedTestFileExt,
		"file",
		templateData,
	); err != nil {
		common.PrintRunningError("%s", err)
		os.Exit(1)
	}
}

func checkGenericDecl() {
	if len(INLINE_ARGS.GenericDecl) < 2 ||
		INLINE_ARGS.GenericDecl[0] != '[' ||
		INLINE_ARGS.GenericDecl[len(INLINE_ARGS.GenericDecl)-1] != ']' {
		common.PrintRunningError("The supplied generic declaration was not valid.")
		common.PrintRunningError("Expected a value of the form '[*]' where * represent the generic types.")
		os.Exit(1)
	}
}

func fileName() string {
	category := ""
	switch PROG_STATE.Cat {
	case Dynamic:
		category = "Dynamic"
	case Static:
		category = "Static"
	}
	return fmt.Sprintf(
		"%s_%s_%sInterface_test",
		INLINE_ARGS.CapType,
		category,
		INLINE_ARGS.Interface,
	)
}

func hasViableSignature(fSet *token.FileSet, srcFile *os.File, fd *ast.FuncDecl) (bool, string) {
	if len(fd.Type.Params.List) != 2 {
		return false, "Did not have two parameters."
	}
	if fd.Type.Results != nil && len(fd.Type.Results.List) > 0 {
		return false, "Specified a return value when none was expected."
	}
	if ok, info := viableFirstParam(fSet, srcFile, fd); !ok {
		return false, fmt.Sprintf("First Param error: %s", info)
	}
	if ok, info := viableSecondParam(fSet, srcFile, fd); !ok {
		return false, fmt.Sprintf("Second Param error: %s", info)
	}
	return true, ""
}

func viableFirstParam(fSet *token.FileSet, srcFile *os.File, fd *ast.FuncDecl) (bool, string) {
	if fd.Type.Params.List[0].Names[0].Name != FirstParamName {
		return false, fmt.Sprintf("Was not named %s", FirstParamName)
	}
	if ft, ok := fd.Type.Params.List[0].Type.(*ast.FuncType); ok {
		if res, info := isViableFactory(fSet, srcFile, ft); !res {
			return false, info
		}
	} else {
		return false, "Expected a function"
	}
	return true, ""
}

func isViableFactory(fSet *token.FileSet, srcFile *os.File, ft *ast.FuncType) (bool, string) {
	if ft.Params == nil || ft.Params.NumFields() != 1 {
		return false, "Expected a function that accepted one argument."
	}

	if src, err := common.GetSourceTextFromExpr(fSet, srcFile, ft.Params.List[0].Type); err != nil {
		return false, fmt.Sprintf("An error ocurred reading it's arguments from the src file.\n%s", err.Error())
	} else if src != "int" {
		return false, fmt.Sprintf("Factory argument was not correct .\nExpected: 'int'\nGot: '%s'\n", string(src))
	}

	if ft.Results.NumFields() != 1 {
		return false, "Expected a function that returned a single value."
	}

	expSrcType := fmt.Sprintf("%sContainers.%s*", PROG_STATE.Cat.String(), INLINE_ARGS.Interface)
	if src, err := common.GetSourceTextFromExpr(fSet, srcFile, ft.Results); err != nil {
		return false, fmt.Sprintf("An error ocurred reading it's return type from the src file.\n%s", err.Error())
	} else if match, _ := regexp.Match(expSrcType, []byte(src)); !match {
		return false, fmt.Sprintf("Src type was not correct.\nExpected: '%s'\nGot: '%s'\n", expSrcType, string(src))
	}

	return true, ""
}

func viableSecondParam(fSet *token.FileSet, srcFile *os.File, fd *ast.FuncDecl) (bool, string) {
	if fd.Type.Params.List[1].Names[0].Name != SecondParamName {
		return false, fmt.Sprintf("Was not named: %s", SecondParamName)
	}

	expSrcType := "\\*testing.T"
	if src, err := common.GetSourceTextFromExpr(fSet, srcFile, fd.Type.Params.List[1].Type); err != nil {
		return false, fmt.Sprintf("An error ocurred reading it's arguments from the src file.\n%s", err.Error())
	} else if match, _ := regexp.Match(expSrcType, []byte(src)); !match {
		return false, fmt.Sprintf("Src type was not correct.\nExpected: %s\nGot: %s\n", expSrcType, string(src))
	}

	return true, ""
}
