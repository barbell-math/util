package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/barbell-math/util/generators/common"
)

type (
	Values struct {
		Package  string `required:"t" help:"The package to put the files in."`
		Type     string `required:"t" help:"The underlying type to generate the widget for."`
		CapType  string `required:"f" default:"" help:"The type with the first letter capitilized. This will be generated automatically if not supplied."`
		ShowInfo bool   `required:"f" default:"t" help:"Show debug info."`
	}
)

var (
	VALS        Values
	VALID_TYPES []string = []string{
		"bool",
		"byte",
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"float32",
		"float64",
		"complex64",
		"complex128",
		"string",
		"uintptr",
		// This is a special case that is only allowed because the widget package itself
		// relies on hash.Hash, making it so the hash.Hash package cannot implement the
		// widget interface on itself, would create circular imports.
		"hash.Hash",
	}
)

func main() {
	common.Args(&VALS, os.Args)
	checkSuppliedType()

	VALS.CapType = VALS.Type
	dotSplit := strings.SplitN(VALS.CapType, ".", 2)
	if len(dotSplit) > 1 {
		VALS.CapType = dotSplit[len(dotSplit)-1]
	}
	VALS.CapType = fmt.Sprintf("%s%s", strings.ToUpper(VALS.CapType)[:1], VALS.CapType[1:])

	fName := fmt.Sprintf("Builtin%s.go", VALS.CapType)
	f, err := os.Create(fName)
	if err != nil {
		common.PrintRunningError("Could not open to write to it.", fName)
		os.Exit(1)
	}

	t, err := template.New("builtin").Parse(
		"package {{ .Package }}\n\n" +
			"// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.\n\n" +
			generateImports() +
			generateGlobals() +
			"// A widget to represent the built-in type {{ .Type }}\n" +
			"// This is meant to be used with the containers from the [containers] package.\n" +
			"type Builtin{{ .CapType }} struct{}\n\n" +
			"// Returns true if both {{ .Type }}'s are equal. Uses the standard == operator internally.\n" +
			"func (_ Builtin{{ .CapType }}) Eq(l *{{ .Type }}, r *{{ .Type }}) bool {\n" +
			"\treturn *l == *r\n" +
			"}\n\n" +
			"// Returns true if a is less than r. Uses the standard < operator internally.\n" +
			generateLtFunction() +
			"// Provides a hash function for the value that it is wrapping.\n" +
			generateHashFunction() +
			"\n" +
			"// Zeros the supplied value.\n" +
			generateZeroFunction() +
			"\n" +
			generateArithFuncs(),
	)
	if err != nil {
		common.PrintRunningError(
			"An error occurred parsing the template: %s",
			err,
		)
		f.Close()
		os.Exit(1)
	}

	err = t.Execute(f, VALS)
	if err != nil {
		common.PrintRunningError(
			"An error occurred when populating the template: %s",
			err,
		)
		f.Close()
		os.Exit(1)
	}
	f.Close()
}

func checkSuppliedType() {
	foundType := false
	for _, v := range VALID_TYPES {
		if foundType = (v == VALS.Type); foundType {
			break
		}
	}
	if !foundType {
		fmt.Println("ERROR | The supplied type was not one of the types recognized by this tool.")
		fmt.Println("The following types are recognized: ", VALID_TYPES)
		fmt.Println("The following type was received: ", VALS.Type)
		os.Exit(1)
	}
}

func generateImports() string {
	commonImport := "import \"github.com/barbell-math/util/hash\"\n\n"
	if VALS.Type == "string" {
		return "import \"hash/maphash\"\n" + commonImport
	}
	if VALS.Type == "complex64" || VALS.Type == "complex128" {
		return "import \"github.com/barbell-math/util/math/basic\"\n" + commonImport
	}
	return commonImport
}

func generateGlobals() string {
	if VALS.Type == "string" {
		return "// The random seed will be different every time the program runs" +
			"// meaning that between runs the hash values will not be consistent.\n" +
			"// This was done for security purposes.\n" +
			"var RANDOM_SEED_{{ .Type }} maphash.Seed = maphash.MakeSeed()\n\n"
	}
	return ""
}

func generateLtFunction() string {
	switch VALS.Type {
	case "bool":
		return "func (_ Builtin{{ .CapType }}) Lt(l *{{ .Type }}, r *{{ .Type }}) bool {\n" +
			"\treturn (!*l && *r)\n" +
			"}\n\n"
	case "complex64":
		fallthrough
	case "complex128":
		return "func (_ Builtin{{ .CapType }}) Lt(l *{{ .Type }}, r *{{ .Type }}) bool {\n" +
			"\tpanic(\"Complex values cannot be compared relative to each other!\")\n" +
			"}\n\n"
	case "byte":
		fallthrough
	case "int":
		fallthrough
	case "int8":
		fallthrough
	case "int16":
		fallthrough
	case "int32":
		fallthrough
	case "int64":
		fallthrough
	case "uint":
		fallthrough
	case "uint8":
		fallthrough
	case "uint16":
		fallthrough
	case "uint32":
		fallthrough
	case "uint64":
		fallthrough
	case "hash.Hash":
		fallthrough
	case "float32":
		fallthrough
	case "float64":
		fallthrough
	case "string":
		fallthrough
	case "uintptr":
		fallthrough
	default:
		return "func (_ Builtin{{ .CapType }}) Lt(l *{{ .Type }}, r *{{ .Type }}) bool {\n" +
			"\treturn *l < *r\n" +
			"}\n\n"
	}
}

func generateHashFunction() string {
	switch VALS.Type {
	case "bool":
		return "func (_ Builtin{{ .CapType }}) Hash(v *{{ .Type }}) hash.Hash {\n" +
			"\tif *v {\n" +
			"\t\treturn hash.Hash(1)\n" +
			"\t}\n" +
			"\treturn hash.Hash(0)\n" +
			"}\n"
	case "complex64":
		return "func (_ Builtin{{ .CapType }}) Hash(v *{{ .Type }}) hash.Hash {\n" +
			"\treturn hash.Hash(basic.LossyConv[float32, int32](basic.RealPart[complex64, float32](*v))).\n" +
			"\t\tCombine(hash.Hash(basic.LossyConv[float32, int32](basic.ImaginaryPart[complex64, float32](*v))))\n" +
			"}\n"
	case "complex128":
		return "func (_ Builtin{{ .CapType }}) Hash(v *{{ .Type }}) hash.Hash {\n" +
			"\treturn hash.Hash(basic.LossyConv[float64, int64](basic.RealPart[complex128, float64](*v))).\n" +
			"\t\tCombine(hash.Hash(basic.LossyConv[float64, int64](basic.ImaginaryPart[complex128, float64](*v))))\n" +
			"}\n"
	case "byte":
		fallthrough
	case "int":
		fallthrough
	case "int8":
		fallthrough
	case "int16":
		fallthrough
	case "int32":
		fallthrough
	case "int64":
		fallthrough
	case "uint":
		fallthrough
	case "uint8":
		fallthrough
	case "uint16":
		fallthrough
	case "uint32":
		fallthrough
	case "uint64":
		fallthrough
	case "uintptr":
		fallthrough
	case "hash.Hash":
		return "func (_ Builtin{{ .CapType }}) Hash(v *{{ .Type }}) hash.Hash {\n" +
			"\treturn hash.Hash(*v)\n" +
			"}\n"
	case "float32":
		fallthrough
	case "float64":
		return "func (_ Builtin{{ .CapType }}) Hash(v *{{ .Type }}) hash.Hash {\n" +
			"\tpanic(\"Floats are not hashable!\")\n" +
			"}\n"
	case "string":
		return "func (_ Builtin{{ .CapType }}) Hash(v *{{ .Type }}) hash.Hash {\n" +
			"\treturn hash.Hash(maphash.String(RANDOM_SEED_{{ .Type }}, *(v)))\n" +
			"}\n"
	default:
		return "func (_ Builtin{{ .CapType }}) Hash(v *{{ .Type }}) hash.Hash {\n" +
			"\t// this will fail compilation (on purpose!)\n" +
			"\t// the supplied type was not hashable!\n" +
			"\tpanic(\"Supplied type was not hashable!\")\n" +
			"\treturn hash.Hash(-1)\n" +
			"}\n"
	}
}

func generateZeroFunction() string {
	switch VALS.Type {
	case "bool":
		return "func (_ Builtin{{ .CapType }}) Zero(v *{{ .Type }}) {\n" +
			"\t*v = false\n" +
			"}\n"
	case "byte":
		fallthrough
	case "int":
		fallthrough
	case "int8":
		fallthrough
	case "int16":
		fallthrough
	case "int32":
		fallthrough
	case "int64":
		fallthrough
	case "uint":
		fallthrough
	case "uint8":
		fallthrough
	case "uint16":
		fallthrough
	case "uint32":
		fallthrough
	case "uint64":
		fallthrough
	case "float32":
		fallthrough
	case "float64":
		fallthrough
	case "complex64":
		fallthrough
	case "complex128":
		fallthrough
	case "uintptr":
		fallthrough
	case "hash.Hash":
		return "func (_ Builtin{{ .CapType }}) Zero(v *{{ .Type }}) {\n" +
			"\t*v = {{ .Type }}(0)\n" +
			"}\n"
	case "string":
		return "func (_ Builtin{{ .CapType }}) Zero(v *{{ .Type }}) {\n" +
			"\t*v = \"\"\n" +
			"}\n"
	default:
		return "func (_ Builtin{{ .CapType }}) Zero(v *{{ .Type }}) {\n" +
			"\t// this will fail compilation (on purpose!)\n" +
			"\t// the supplied type was not found in the zero table!\n" +
			"\tpanic(\"The supplied type does not have a zeor value.\")\n" +
			"\treturn int(-1)\n" +
			"}\n"
	}
}

func generateArithFuncs() string {
	switch VALS.Type {
	case "bool":
		return "// A bool is not an arithmetic aware widget. Bools are only base widgets.\n"
	case "byte":
		fallthrough
	case "int":
		fallthrough
	case "int8":
		fallthrough
	case "int16":
		fallthrough
	case "int32":
		fallthrough
	case "int64":
		fallthrough
	case "uint":
		fallthrough
	case "uint8":
		fallthrough
	case "uint16":
		fallthrough
	case "uint32":
		fallthrough
	case "uint64":
		fallthrough
	case "float32":
		fallthrough
	case "float64":
		fallthrough
	case "complex64":
		fallthrough
	case "complex128":
		fallthrough
	case "uintptr":
		fallthrough
	case "hash.Hash":
		return "func (_ Builtin{{ .CapType }}) ZeroVal() {{ .Type }} {\n" +
			"\treturn {{ .Type }}(0)\n" +
			"}\n\n" +
			"func (_ Builtin{{ .CapType }}) UnitVal() {{ .Type }} {\n" +
			"\treturn {{ .Type }}(1)\n" +
			"}\n\n" +
			"func (_ Builtin{{ .CapType }}) Neg(v *{{ .Type }}) {\n" +
			"\t*v = -(*v)\n" +
			"}\n\n" +
			"func (_ Builtin{{ .CapType }}) Add(res *{{ .Type }}, l *{{ .Type }}, r *{{ .Type }}) {\n" +
			"\t*res = *l + *r\n" +
			"}\n\n" +
			"func (_ Builtin{{ .CapType }}) Sub(res *{{ .Type }}, l *{{ .Type }}, r *{{ .Type }}) {\n" +
			"\t*res = *l - *r\n" +
			"}\n\n" +
			"func (_ Builtin{{ .CapType }}) Mul(res *{{ .Type }}, l *{{ .Type }}, r *{{ .Type }}) {\n" +
			"\t*res = *l * *r\n" +
			"}\n\n" +
			"func (_ Builtin{{ .CapType }}) Div(res *{{ .Type }}, l *{{ .Type }}, r *{{ .Type }}) {\n" +
			"\t*res = *l / *r\n" +
			"}\n"
	case "string":
		return "// A string is not an arithmetic aware widget. Strings are only base widgets.\n"
	default:
		return ""
	}
}
