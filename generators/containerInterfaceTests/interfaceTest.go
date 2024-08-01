package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"regexp"
	"strings"

	"github.com/barbell-math/util/generators/common"
)

type (
	category byte

	Values struct {
		Type        string `required:"t" default:"" help:"The underlying type to generate the widget for."`
		Interface   string `required:"t" default:"" help:"The packge to put the files in."`
		GenericDecl string `required:"t" default:"" help:"The generic type signature to use."`
		Factory     string `required:"t" default:"" help:"The factory that will produce containers to test."`
		Category    string `required:"t" default:"" help:"Either static or dynamic."`
		CapType     string `required:"f" default:"" help:"The type but the first letter is capitilized. This will be calculated if left blank."`
		Debug       bool   `required:"f" default:"false" help:"Print diagonistic information to the console."`

		cat category
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

	VALS.cat = Dynamic.FromString(VALS.Category)
	VALS.CapType = strings.ToUpper(VALS.Type)[:1] + VALS.Type[1:]

	testFuncs := viableFuncs()
	common.PrintRunningInfo("Num Funcs: %3d", len(testFuncs))

	fName := fileName()
	f, err := os.Create(fName)
	if err != nil {
		common.PrintRunningError("Could not open %s to write to it.", fName)
		os.Exit(1)
	}

	f.WriteString("package containers\n\n")
	f.WriteString("// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.\n\n")
	f.WriteString("import (\n")
	f.WriteString("    \"testing\"\n")
	f.WriteString("    \"github.com/barbell-math/util/container/tests\"\n")
	f.WriteString(fmt.Sprintf(
		"    \"github.com/barbell-math/util/container/%sContainers\"\n",
		VALS.cat.String(),
	))
	f.WriteString(")\n\n")
	f.WriteString(fmt.Sprintf(
		"func %sTo%sInterfaceFactory(capacity int) %sContainers.%s%s {\n",
		VALS.Type, VALS.Interface, VALS.cat.String(), VALS.Interface, VALS.GenericDecl,
	))
	f.WriteString(fmt.Sprintf("    v:= %s(capacity)\n", VALS.Factory))
	f.WriteString(fmt.Sprintf(
		"    var rv %sContainers.%s%s=&v\n",
		VALS.cat.String(), VALS.Interface, VALS.GenericDecl,
	))
	f.WriteString("    return rv\n")
	f.WriteString("}\n\n")

	for _, iterFunc := range testFuncs {
		f.WriteString(fmt.Sprintf(
			"func Test%s_%s(t *testing.T){\n"+
				"    tests.%s(%sTo%sInterfaceFactory,t)\n"+
				"}\n\n",
			VALS.Type, iterFunc.Name, iterFunc.Name, VALS.Type, VALS.Interface,
		))
	}

	f.Close()
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
	switch VALS.cat {
	case Dynamic:
		category = "Dynamic"
	case Static:
		category = "Static"
	}
	return fmt.Sprintf(
		"./%s_%s_%sInterface_test.go",
		VALS.CapType,
		category,
		VALS.Interface,
	)
}

func viableFuncs() []*ast.FuncDecl {
	fSet := token.NewFileSet()
	if VALS.Debug {
		fmt.Println("Locating appropriate funcs from tests dir.")
	}
	packs, err := parser.ParseDir(fSet, "../tests/", nil, 0)
	if err != nil {
		common.PrintRunningError("Failed to parse package:", err)
		os.Exit(1)
	}
	rv := []*ast.FuncDecl{}
	for _, pack := range packs {
		for fileName, f := range pack.Files {
			srcFile, err := os.OpenFile(fileName, os.O_RDONLY, 666)
			if err != nil {
				common.PrintRunningError(
					"Could not open file %s to parse source.",
					fileName,
				)
			}
			ast.Inspect(f, func(n ast.Node) bool {
				if fd, ok := n.(*ast.FuncDecl); ok {
					if VALS.Debug {
						fmt.Print("Found func: ", fd.Name, "| Status: ")
					}
					if ok, info := hasViableSignature(fd, srcFile, fSet); ok {
						if VALS.Debug {
							fmt.Println("Accepted")
						}
						rv = append(rv, fd)
					} else if VALS.Debug {
						fmt.Println("Rejected | Reason:", info)
					}
				}
				return true
			})
			srcFile.Close()
		}
	}
	return rv
}

func hasViableSignature(fn *ast.FuncDecl, srcFile *os.File, fSet *token.FileSet) (bool, string) {
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
		expSrcType := fmt.Sprintf("%sContainers.%s*", VALS.cat.String(), VALS.Interface)
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
