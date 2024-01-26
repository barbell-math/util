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
		"}\n\n",
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
