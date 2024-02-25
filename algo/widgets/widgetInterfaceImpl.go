//go:build ignore

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Values struct {
	Package string
	Type string
	CapType string
	ShowInfo bool
}

var VALS Values
var REQUIRED_ARGS []string=[]string{"package","type"}

func main() {
	setupFlags()
	parseArgs()
	checkRequiredArgs()
	VALS.CapType=fmt.Sprintf("%s%s",strings.ToUpper(VALS.Type)[:1],VALS.Type[1:])

	if VALS.ShowInfo {
		fmt.Println("Making widget for type ",VALS.Type, " using the below options.")
		fmt.Println("Recieved the following values:")
		fmt.Println("  Package: ",VALS.Package)
		fmt.Println("  Type: ",VALS.Type)
		fmt.Println("  CapType: ",VALS.CapType)
	}

	fName:=fmt.Sprintf("Builtin%s.go",VALS.CapType)
	f,err:=os.Create(fName)
	if err!=nil {
		fmt.Println("ERROR | Could not open ",fName," to write to it.")
		os.Exit(1)
	}

	t,err:=template.New("builtin").Parse(
		"package {{ .Package }}\n\n"+
		"// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.\n\n"+
		generateImports()+
		generateGlobals()+
		"// A widget to represent the built-in type {{ .Type }}\n"+
		"// This is meant to be used with the containers from the [containers] package.\n"+
		"type Builtin{{ .CapType }} struct{}\n\n"+
		"// Returns true if both {{ .Type }}'s are equal. Uses the standard == operator internally.\n"+
		"func (a Builtin{{ .CapType }})Eq(l *{{ .Type }}, r *{{ .Type }}) bool {\n"+
		"    return *l==*r\n"+
		"}\n\n"+
		"// Returns true if a is less than r. Uses the standard < operator internally.\n"+
		"func (a Builtin{{ .CapType }})Lt(l *{{ .Type }}, r *{{ .Type }}) bool {\n"+
		"    return *l<*r\n"+
		"}\n\n"+
		"// Provides a hash function for the value that it is wrapping.\n"+
		generateHashFunction()+
		"// Zeros the supplied value.\n"+
		generateZeroFunction(),
		// generateArithFuncs(),
	)
	if err!=nil {
		fmt.Println("ERROR | An error occurred parsing the template.")
		f.Close()
		os.Exit(1)
	}

	err=t.Execute(f,VALS)
    if err!=nil {
		fmt.Println("ERROR | An error occurred when populating the template.")
	    f.Close()
        os.Exit(1)
	}
    f.Close()
}

func setupFlags() {
	flag.StringVar(&VALS.Package,"package","","The packge to put the files in.")
	flag.StringVar(&VALS.Type,"type","","The underlying type to generate the widget for.")
	flag.BoolVar(&VALS.ShowInfo,"info",false,"Print debug information.")
}

func parseArgs() {
	if len(os.Args)<3 {
		fmt.Println("ERROR | Not enough arguments.")
		fmt.Println("Recieved: ",os.Args[1:])
		flag.PrintDefaults()
		fmt.Println("Re-run go generate after fixing the problem.")
		os.Exit(1)
	}
	flag.Parse()
}

func checkRequiredArgs() {
	requiredCopy:=append([]string{},REQUIRED_ARGS...)
	flag.Visit(func(f *flag.Flag) {
		for i,v:=range(requiredCopy) {
			if f.Name==v {
				requiredCopy=append(requiredCopy[:i],requiredCopy[i+1:]...)
			}
		}
	})
	if len(requiredCopy)>0 {
		fmt.Println("ERROR | Not all required flags were passed.")
		fmt.Println("The following flags must be added: ",requiredCopy)
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func generateImports() string {
	commonImport:="import \"github.com/barbell-math/util/algo/hash\"\n\n"
	if VALS.Type=="string" {
		return "import \"hash/maphash\"\n"+commonImport
	}
	return commonImport
}

func generateGlobals() string {
	if VALS.Type=="string" {
		return "// The random seed will be differrent every time the program runs"+
			"// meaning that between runs the hash values will not be consistent.\n"+
			"// This was done for security purposes.\n"+
			"var RANDOM_SEED_{{ .Type }} maphash.Seed=maphash.MakeSeed()\n\n"
	}
	return ""
}

func generateHashFunction() string {
	switch VALS.Type {
		case "byte": fallthrough
		case "int": fallthrough
		case "int8": fallthrough
		case "int16": fallthrough
		case "int32": fallthrough
		case "int64": fallthrough
		case "uint": fallthrough
		case "uint8": fallthrough
		case "uint16": fallthrough
		case "uint32": fallthrough
		case "uint64":
			return "func (a Builtin{{ .CapType }})Hash(v *{{ .Type }}) hash.Hash {\n"+
				"    return hash.Hash(*v)\n"+
			    "}\n\n"
		case "float32": fallthrough
		case "float64":
			return "func (a Builtin{{ .CapType }})Hash(v *{{ .Type }}) hash.Hash {\n"+
				"    panic(\"Floats are not hashable!\")\n"+
			    "}\n\n"
		case "string":
			return "func (a Builtin{{ .CapType }})Hash(v *{{ .Type }}) hash.Hash {\n"+
				"    return hash.Hash(maphash.String(RANDOM_SEED_{{ .Type }},*(v)))\n"+
			    "}\n\n"
		default:
			return "func (a Builtin{{ .CapType }})Hash(v *{{ .Type }}) hash.Hash {\n"+
				"    // this will fail compilation (on purpose!)\n"+
				"    // the supplied type was not hashable!\n"+
				"    panic(\"Supplied type was not hashable!\")\n"+
				"    return hash.Hash(-1)\n"+
			    "}\n\n"
	}
}

func generateZeroFunction() string {
	switch VALS.Type {
		case "byte": fallthrough
		case "int": fallthrough
		case "int8": fallthrough
		case "int16": fallthrough
		case "int32": fallthrough
		case "int64": fallthrough
		case "uint": fallthrough
		case "uint8": fallthrough
		case "uint16": fallthrough
		case "uint32": fallthrough
		case "uint64": fallthrough
		case "float32": fallthrough
		case "float64":
			return "func (a Builtin{{ .CapType }})Zero(v *{{ .Type }}) {\n"+
				"    *v={{ .Type }}(0)\n"+
				"}\n\n"
		case "string":
			return "func (a Builtin{{ .CapType }})Zero(v *{{ .Type }}) {\n"+
				"    *v=\"\"\n"+
				"}\n\n"
		default:
			return "func (a Builtin{{ .CapType }})Zero(v *{{ .Type }}) {\n"+
				"    // this will fail compilation (on purpose!)\n"+
				"    // the supplied type was not found in the zero table!\n"+
				"	panic(\"The supplied type does not have a zeor value.\")\n"+
				"    return int(-1)\n"+
			    "}\n\n"
	}
}
