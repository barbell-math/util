//go:build ignore

package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

type Values struct {
	Package        string
	AliasType      string
	BaseType       string
	BaseTypeWidget string
	WidgetPackage  string
	ShowInfo       bool
}

var VALS Values
var REQUIRED_ARGS []string = []string{
	"package",
	"aliasType",
	"baseType",
	"baseTypeWidget",
	"widgetPackage",
}

func main() {
	setupFlags()
	parseArgs()
	checkRequiredArgs()

	fmt.Printf(
		"Making pass through type alias widget | Alias Type: %-30s | Base Type: %-40s | Base Type Widget: %-40s\n",
		VALS.AliasType,
		VALS.BaseType,
		VALS.BaseTypeWidget,
	)
	if VALS.ShowInfo {
		fmt.Println("Received the following values:")
		fmt.Println("\tPackage: ", VALS.Package)
		fmt.Println("\tAlias Type: ", VALS.AliasType)
		fmt.Println("\tBase Type: ", VALS.BaseType)
		fmt.Println("\tBase Type Widget: ", VALS.BaseTypeWidget)
	}

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
		fmt.Println("ERROR | Could not open ", fName, " to write to it.")
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

func setupFlags() {
	flag.StringVar(&VALS.Package, "package", "", "The packge to put the files in.")
	flag.StringVar(&VALS.AliasType, "aliasType", "", "The alias type to generate the widget for.")
	flag.StringVar(&VALS.BaseType, "baseType", "", "The base type to generate the widget for.")
	flag.StringVar(&VALS.BaseTypeWidget, "baseTypeWidget", "", "The base type widget to use when generating the new widget.")
	flag.StringVar(&VALS.WidgetPackage, "widgetPackage", "", "The package the base type widget resides in. If it is this package, put '.'")
	flag.BoolVar(&VALS.ShowInfo, "info", false, "Print debug information.")
}

func parseArgs() {
	if len(os.Args) < 6 {
		fmt.Println("ERROR | Not enough arguments.")
		fmt.Println("Received: ", os.Args[1:])
		flag.PrintDefaults()
		fmt.Println("Re-run go generate after fixing the problem.")
		os.Exit(1)
	}
	flag.Parse()
}

func checkRequiredArgs() {
	requiredCopy := append([]string{}, REQUIRED_ARGS...)
	flag.Visit(func(f *flag.Flag) {
		for i, v := range requiredCopy {
			if f.Name == v {
				requiredCopy = append(requiredCopy[:i], requiredCopy[i+1:]...)
			}
		}
	})
	if len(requiredCopy) > 0 {
		fmt.Println("ERROR | Not all required flags were passed.")
		fmt.Println("The following flags must be added: ", requiredCopy)

		cntr:=0
		fmt.Println("Received: ")
		flag.Visit(func(f *flag.Flag) {
			cntr++
			fmt.Printf(" %d. %s: %+v\n",cntr,f.Name,f.Value)
		})
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func generateImports() string {
	commonImport := "import \"github.com/barbell-math/util/algo/hash\"\n"
	if VALS.WidgetPackage != "." {

		commonImport = commonImport + "import \"{{ .WidgetPackage }}\"\n"
	}
	return commonImport + "\n"
}
