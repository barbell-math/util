//go:build ignore

package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Values struct {
	Package string
	Type string
	CapType string
}

func main() {
	if len(os.Args)!=3 {
		fmt.Println(
			"ERROR | "+
			"widgetTemplate expects two arguments: the name of the package and "+
			"the name of the type to generate the widget for.",
		)
		fmt.Println("Recieved: ",os.Args[1:])
		fmt.Println("Re-run go generate after fixing the problem.")
		os.Exit(1)
	}

	vals:=Values{
		Package: os.Args[1],
		Type: os.Args[2],
		CapType: fmt.Sprintf("%s%s",strings.ToUpper(os.Args[2])[:1],os.Args[2][1:]),
	}
	fmt.Println("Recieved the following values:")
	fmt.Println("  Package: ",vals.Package)
	fmt.Println("  Type: ",vals.Type)
	fmt.Println("  CapType: ",vals.CapType)

	fName:=fmt.Sprintf("Builtin%s.go",vals.CapType)
	f,err:=os.Create(fName)
	if err!=nil {
		fmt.Println("ERROR | Could not open ",fName," to write to it.")
		os.Exit(1)
	}

	t,err:=template.New("builtin").Parse(
		"package {{ .Package }}\n\n"+
		generateImports(vals.Type)+
		generateGlobals(vals.Type)+
		"type Builtin{{ .CapType }} {{.Type}}\n\n"+
		"func (a *Builtin{{ .CapType }})Eq(r *{{ .Type }}) bool {\n"+
		"    return ({{ .Type }}(*a))==*r\n"+
		"}\n\n"+
		"func (a *Builtin{{ .CapType }})Lt(r *{{ .Type }}) bool {\n"+
		"    return ({{ .Type }}(*a))<*r\n"+
		"}\n\n"+
		"func (a *Builtin{{ .CapType }})Unwrap() {{ .Type }} {\n"+
		"    return {{ .Type }}(*a)\n"+
		"}\n\n"+
		"func (a *Builtin{{ .CapType }})Wrap(v *{{ .Type }}) {\n"+
		"    *a=Builtin{{ .CapType }}(*v)\n"+
		"}\n\n"+
		generateHashFunction(vals.Type),
	)
	if err!=nil {
		fmt.Println("ERROR | An error occurred parsing the template.")
		f.Close()
		os.Exit(1)
	}

	if err:=t.Execute(f,vals); err!=nil {
		fmt.Println("ERROR | An error occurred when populating the template.")
	}
	f.Close()
}

func generateImports(typeName string) string {
	if typeName=="string" {
		return "import \"hash/maphash\"\n\n"
	}
	return ""
}

func generateGlobals(typeName string) string {
	if typeName=="string" {
		return "var RANDOM_SEED_{{ .Type }} maphash.Seed=maphash.MakeSeed()\n\n"
	}
	return ""
}

func generateHashFunction(typeName string) string {
	switch typeName {
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
			return "func (a *Builtin{{ .CapType }})Hash() uint64 {\n"+
				"    return uint64(a.Unwrap())\n"+
			    "}\n\n"
		case "float32": fallthrough
		case "float64":
			return "func (a *Builtin{{ .CapType }})Hash() uint64 {\n"+
				"    panic(\"Floats are not hashable!\")\n"+
			    "}\n\n"
		case "string":
			return "func (a *Builtin{{ .CapType }})Hash() uint64 {\n"+
				"    return maphash.String(RANDOM_SEED_{{ .Type }},string(*a))\n"+
			    "}\n\n"
		default:
			return "func (a *Builtin{{ .CapType }})Hash() uint64 {\n"+
				"    // this will fail compilation (on purpose!)\n"+
				"    // the supplied type was not hashable!\n"+
				"    return int(-1)\n"+
			    "}\n\n"
	}
}
