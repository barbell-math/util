package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/barbell-math/util/generators/common"
)

type (
	Values struct {
		Package        string `required:"t" help:"The packge to put the files in."`
		AliasType      string `required:"t" help:"The alias type to generate the widget for."`
		BaseType       string `required:"t" help:"The base type to generate the widget for."`
		BaseTypeWidget string `required:"t" help:"The base type widget to use when generating the new widget."`
		WidgetPackage  string `required:"t" help:"The package the base type widget resides in. If it is this package, put '.'"`
	}
)

var (
	VALS Values
)

func main() {
	common.Args(&VALS, os.Args)

	fileNameBaseType := []byte(VALS.BaseType)
	fileNameAliasType := []byte(VALS.AliasType)
	bannedChars := map[byte]struct{}{
		'[':  struct{}{},
		']':  struct{}{},
		'{':  struct{}{},
		'}':  struct{}{},
		':':  struct{}{},
		';':  struct{}{},
		'<':  struct{}{},
		'>':  struct{}{},
		',':  struct{}{},
		'.':  struct{}{},
		'/':  struct{}{},
		'\\': struct{}{},
		'|':  struct{}{},
		'*':  struct{}{},
		'?':  struct{}{},
		'%':  struct{}{},
		'"':  struct{}{},
		' ':  struct{}{},
	}
	for i, c := range fileNameBaseType {
		if _, ok := bannedChars[c]; ok {
			fileNameBaseType[i] = '_'
		}
	}
	for i, c := range fileNameAliasType {
		if _, ok := bannedChars[c]; ok {
			fileNameAliasType[i] = '_'
		}
	}

	fName := fmt.Sprintf(
		"TypeAliasPassThroughWidget_%s_to_%s.go",
		fileNameAliasType,
		fileNameBaseType,
	)
	f, err := os.Create(fName)
	if err != nil {
		common.PrintRunningError("Could not open to write to it.", fName)
		os.Exit(1)
	}

	t, err := template.New("builtin").Parse(
		"package {{ .Package }}\n\n" +
			"// THIS FILE IS AUTO-GENERATED. DO NOT EDIT AND EXPECT CHANGES TO PERSIST.\n\n" +
			generateImports() +
			"// A pass through widget to represent the aliased type {{ .AliasType }}\n" +
			"// This is meant to be used with the containers from the [containers] package.\n" +
			"// Returns true if both {{ .AliasType }}'s are equal. Uses the Eq operator provided by the {{ .BaseTypeWidget }} widget internally.\n" +
			"func (_ *{{ .AliasType }}) Eq(l *{{ .AliasType }}, r *{{ .AliasType }}) bool {\n" +
			"\tvar tmp {{ .BaseTypeWidget }}\n" +
			"\treturn tmp.Eq((*{{ .BaseType }})(l), (*{{ .BaseType }})(r))\n" +
			"}\n\n" +
			"// Returns true if a is less than r. Uses the Lt operator provided by the {{ .BaseTypeWidget }} widget internally.\n" +
			"func (_ *{{ .AliasType }}) Lt(l *{{ .AliasType }}, r *{{ .AliasType }}) bool {\n" +
			"\tvar tmp {{ .BaseTypeWidget }}\n" +
			"\treturn tmp.Lt((*{{ .BaseType }})(l), (*{{ .BaseType }})(r))\n" +
			"}\n\n" +
			"// Provides a hash function for the value that it is wrapping. The value that is returned will be supplied by the {{ .BaseTypeWidget }} widget internally.\n" +
			"func (_ *{{ .AliasType }}) Hash(other *{{ .AliasType }}) hash.Hash {\n" +
			"\tvar tmp {{ .BaseTypeWidget }}\n" +
			"\treturn tmp.Hash((*{{ .BaseType }})(other))\n" +
			"}\n\n" +
			"// Zeros the supplied value. The operation that is performed will be determined by the {{ .BaseTypeWidget }} widget internally.\n" +
			"func (_ *{{ .AliasType }}) Zero(other *{{ .AliasType }}) {\n" +
			"\tvar tmp {{ .BaseTypeWidget }}\n" +
			"\ttmp.Zero((*{{ .BaseType }})(other))\n" +
			"}\n",
	)
	if err != nil {
		fmt.Println("ERROR | An error occurred parsing the template.")
		fmt.Println(err)
		f.Close()
		os.Exit(1)
	}

	err = t.Execute(f, VALS)
	if err != nil {
		fmt.Println("ERROR | An error occurred when populating the template.")
		fmt.Println(err)
		f.Close()
		os.Exit(1)
	}
	f.Close()
}

func generateImports() string {
	commonImport := "import \"github.com/barbell-math/util/hash\"\n"
	if VALS.WidgetPackage != "." {

		commonImport = commonImport + "import \"{{ .WidgetPackage }}\"\n"
	}
	return commonImport + "\n"
}
