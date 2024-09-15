package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"regexp"
	"strings"

	"github.com/barbell-math/util/generators/common"
)

type (
	category byte

	Values struct {
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
		Cat category
		ViableFuncs []*ast.FuncDecl
	}
	TemplateVals struct {
		Cat string
		Type string
		Interface string
		GenericDecl string
		FuncNames []FuncTemplateVals
		GeneratorName string
		Factory string
	}
	FuncTemplateVals struct {
		Name string
		Type string
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
	VALS          Values
	REQUIRED_ARGS []string = []string{
		"type",
		"category",
		"interface",
		"factory",
		"genericDecl",
	}
	PROG_STATE ProgState=ProgState{
		ViableFuncs: []*ast.FuncDecl{},
	}
	TEMPLATES common.GeneratedFilesRegistry=common.NewGeneratedFilesRegistryFromMap(
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
	"github.com/barbell-math/util/container/tests"
	"github.com/barbell-math/util/container/{{ .Cat }}Containers"
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
	common.Args(&VALS, os.Args)
	checkGenericDecl()
	VALS.CapType = strings.ToUpper(VALS.Type)[:1] + VALS.Type[1:]
	PROG_STATE.Cat=Dynamic.FromString(VALS.Category)

	common.IterateOverAST(
		"../tests/",
		common.GenFileExclusionFilter,
		func(fSet *token.FileSet, file *ast.File, srcFile *os.File, node ast.Node) bool {
			switch node.(type) {
				case *ast.FuncDecl:
					fd:=node.(*ast.FuncDecl)
					if VALS.Debug {
						fmt.Print("Found func: ", fd.Name, "| Status: ")
					}
					if ok, info := hasViableSignature(fSet, srcFile, fd); ok {
						if VALS.Debug {
							fmt.Println("Accepted")
						}
						PROG_STATE.ViableFuncs=append(PROG_STATE.ViableFuncs, fd)
					} else if VALS.Debug {
						fmt.Println("Rejected | Reason:", info)
					}
					return false
				case *ast.GenDecl: return false
				default: return true
			}
		},
	)
	common.PrintRunningInfo("Num Funcs: %3d", len(PROG_STATE.ViableFuncs))

	templateData:=TemplateVals{
		Cat: PROG_STATE.Cat.String(),
		Type: VALS.Type,
		Interface: VALS.Interface,
		GenericDecl: VALS.GenericDecl,
		FuncNames: make([]FuncTemplateVals, len(PROG_STATE.ViableFuncs)),
		GeneratorName: os.Args[0],
		Factory: VALS.Factory,
	}
	for i,f:=range(PROG_STATE.ViableFuncs) {
		templateData.FuncNames[i]=FuncTemplateVals{
			Name: f.Name.Name,
			Type: VALS.Type,
			Interface: VALS.Interface,
		}
	}

	if err := TEMPLATES.WriteToFile(
		fileName(),
		"file",
		templateData,
	); err != nil {
		common.PrintRunningError("%s", err)
		os.Exit(1)
	}
}

func checkGenericDecl() {
	if len(VALS.GenericDecl) < 2 ||
		VALS.GenericDecl[0] != '[' ||
		VALS.GenericDecl[len(VALS.GenericDecl)-1] != ']' {
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
		VALS.CapType,
		category,
		VALS.Interface,
	)
}

func hasViableSignature(fSet *token.FileSet, srcFile *os.File, fn *ast.FuncDecl) (bool, string) {
	if len(fn.Type.Params.List) != 2 {
		return false, "Did not have two parameters."
	}
	if fn.Type.Results != nil && len(fn.Type.Results.List) > 0 {
		return false, "Specified a return value when none was expected."
	}
	if ok, info := viableFirstParam(fn, srcFile, fSet); !ok {
		return false, fmt.Sprintf("First Param error: %s", info)
	}
	if ok, info := viableSecondParam(fn, srcFile, fSet); !ok {
		return false, fmt.Sprintf("Second Param error: %s", info)
	}
	return true, ""
}

// todo reorder func params
func viableFirstParam(fn *ast.FuncDecl, srcFile *os.File, fSet *token.FileSet) (bool, string) {
	if fn.Type.Params.List[0].Names[0].Name != FirstParamName {
		return false, fmt.Sprintf("Was not named %s", FirstParamName)
	}
	if f, ok := fn.Type.Params.List[0].Type.(*ast.FuncType); ok {
		if res, info := isViableFactory(f, srcFile, fSet); !res {
			return false, info
		}
	} else {
		return false, "Expected a function"
	}
	return true, ""
}

func isViableFactory(f *ast.FuncType, srcFile *os.File, fSet *token.FileSet) (bool, string) {
	if f.Params == nil || f.Params.NumFields() != 1 {
		return false, "Expected a function that accepted one argument."
	}
	if _, err := srcFile.Seek(int64(fSet.Position(f.Params.List[0].Type.Pos()).Offset), 0); err != nil {
		return false, fmt.Sprintf("An error occurred seeking to the required location in the src file.\n%s", err.Error())
	}
	src := make([]byte, f.Params.List[0].Type.End()-f.Params.List[0].Type.Pos())
	if _, err := srcFile.Read(src); err == nil {
		if string(src) != "int" {
			return false, fmt.Sprintf("Factory argument was not correct .\nExpected: 'int'\nGot: '%s'\n", string(src))
		}
	} else {
		return false, fmt.Sprintf("An error ocurred reading it's arguments from the src file.\n%s", err.Error())
	}
	if f.Results.NumFields() != 1 {
		return false, "Expected a function that returned a single value."
	}
	if _, err := srcFile.Seek(int64(fSet.Position(f.Results.Pos()).Offset), 0); err != nil {
		return false, fmt.Sprintf("An error occurred seeking to the required location in the src file.\n%s", err.Error())
	}
	src = make([]byte, f.Results.End()-f.Results.Pos())
	if _, err := srcFile.Read(src); err == nil {
		expSrcType := fmt.Sprintf("%sContainers.%s*", PROG_STATE.Cat.String(), VALS.Interface)
		if match, _ := regexp.Match(expSrcType, src); !match {
			return false, fmt.Sprintf("Src type was not correct.\nExpected: '%s'\nGot: '%s'\n", expSrcType, string(src))
		}
	} else {
		return false, fmt.Sprintf("An error ocurred reading it's return type from the src file.\n%s", err.Error())
	}
	return true, ""
}

func viableSecondParam(fn *ast.FuncDecl, srcFile *os.File, fSet *token.FileSet) (bool, string) {
	if fn.Type.Params.List[1].Names[0].Name != SecondParamName {
		return false, fmt.Sprintf("Was not named: %s", SecondParamName)
	}
	if _, err := srcFile.Seek(int64(fSet.Position(fn.Type.Params.List[1].Pos()).Offset), 0); err != nil {
		return false, fmt.Sprintf("An error occurred seeking to the required location in the src file.\n%s", err.Error())
	}
	src := make([]byte, fn.Type.Params.List[1].End()-fn.Type.Params.List[1].Pos())
	if _, err := srcFile.Read(src); err == nil {
		expSrcType := "\\*testing.T"
		if match, _ := regexp.Match(expSrcType, src); !match {
			return false, fmt.Sprintf("Src type was not correct.\nExpected: %s\nGot: %s\n", expSrcType, string(src))
		}
	} else {
		return false, fmt.Sprintf("An error ocurred reading it's return type from the src file.\n%s", err.Error())
	}
	return true, ""
}
