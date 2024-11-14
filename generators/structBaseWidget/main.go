// A generator program that creates methods for a given struct that implement the
// [widgets.BaseInterface].
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
		Type     string `required:"t" help:"The type to make the alias for."`
		ShowInfo bool   `required:"f" default:"t" help:"Show debug info."`
	}
	CommentArgs struct {
		name           string
		Identity       bool   `required:"t" help:"Set flag to true if the value should be part of the new widgets identity."`
		BaseTypeWidget string `required:"t" help:"The underlying widget to use when making comparisons."`
		WidgetPackage  string `required:"t" help:"The package the base type widget is in. Use '.' for the current package."`
	}

	ProgState struct {
		fieldArgs           []CommentArgs
		shortStructGenerics string
		longStructGenerics  string
		_package            string
	}

	TemplateVals struct {
		GeneratorName       string
		Package             string
		Type                string
		ShortStructGenerics string
		LongStructGenerics  string
		Imports             []string
		IdentityFields      []IdentityFields
	}
	IdentityFields struct {
		Name   string
		Widget string
	}
)

var (
	INLINE_ARGS InlineArgs
	PROG_STATE  ProgState
	TEMPLATES   common.GeneratedFilesRegistry = common.NewGeneratedFilesRegistryFromMap(
		map[string]string{
			"imports": `
import (
	"github.com/barbell-math/util/hash"
	{{range .Imports}} "{{ . }}" {{end}}
)
`,
			"eqFunc": `
// Returns true if l equals r using the widgets associated with the following
// fields:
{{range .IdentityFields -}}
//  - {{ .Name }}
{{end -}}
func (_ *{{ .Type }}{{ .ShortStructGenerics }}) Eq(l *{{ .Type }}{{ .ShortStructGenerics }}, r *{{ .Type }}{{ .ShortStructGenerics }}) bool {
	var (
		{{range .IdentityFields -}}
			{{ .Name }}Widget {{ .Widget }}
		{{end -}}
		rv bool=true
	)

	{{range .IdentityFields -}}
		rv=rv && {{ .Name }}Widget.Eq(&l.{{ .Name }}, &r.{{ .Name }})
	{{end -}}
	return rv
}
`,
			"hashFunc": `
// Returns a hash that represents other by calling [hash.CombineIgnoreZero] with
// the values generated from using the widgets associated with the following
// fields:
{{range .IdentityFields -}}
//  - {{ .Name }}
{{end -}}
func (_ *{{ .Type }}{{ .ShortStructGenerics }}) Hash(other *{{ .Type }}{{ .ShortStructGenerics }}) hash.Hash {
	var (
		{{range .IdentityFields -}}
			{{ .Name }}Widget {{ .Widget }}
		{{end -}}
		rv hash.Hash
	)

	{{range .IdentityFields -}}
		rv=rv.CombineIgnoreZero({{ .Name }}Widget.Hash(&other.{{ .Name }}))
	{{end -}}
	return rv
}
`,
			"zeroFunc": `
// Zeros the following fields in the struct using the widgets associated with
// the following fields:
{{range .IdentityFields -}}
//  - {{ .Name }}
{{end -}}
func (_ *{{ .Type }}{{ .ShortStructGenerics }}) Zero (other *{{ .Type }}{{ .ShortStructGenerics }}) {
	var (
		{{range .IdentityFields -}}
			{{ .Name }}Widget {{ .Widget }}
		{{end -}}
	)

	{{range .IdentityFields -}}
		{{ .Name }}Widget.Zero(&other.{{ .Name }})
	{{end -}}
}
`,
			"file": `
package {{ .Package }}

{{template "autoGenComment" .}}
{{template "imports" .}}

{{template "eqFunc" .}}
{{template "hashFunc" .}}
{{template "zeroFunc" .}}
`,
		},
	)
)

func main() {
	common.InlineArgs(&INLINE_ARGS, os.Args)

	found := false
	common.IterateOverAST(
		".",
		common.GenFileExclusionFilter,
		func(fSet *token.FileSet, file *ast.File, srcFile *os.File, node ast.Node) bool {
			switch node.(type) {
			case *ast.GenDecl:
				gdNode := node.(*ast.GenDecl)
				if gdNode.Tok == token.TYPE {
					for _, spec := range gdNode.Specs {
						found = (found || parseTypeSpec(fSet, srcFile, spec.(*ast.TypeSpec)))
					}
				}
				if found {
					PROG_STATE._package = file.Name.Name
				}
				return found
			case *ast.FuncDecl:
				return false
			default:
				return true
			}
		},
	)

	if !found {
		common.PrintRunningError(
			"The supplied type was not found or was not a struct but is required to be.",
		)
		os.Exit(1)
	}
	if len(PROG_STATE.fieldArgs) == 0 {
		common.PrintRunningError(
			"The found struct definition had no fields with an identity flag. At least one is required.",
		)
		os.Exit(1)
	}

	templateData := TemplateVals{
		GeneratorName:       os.Args[0],
		Package:             PROG_STATE._package,
		Type:                INLINE_ARGS.Type,
		ShortStructGenerics: PROG_STATE.shortStructGenerics,
		LongStructGenerics:  PROG_STATE.longStructGenerics,
	}
	for _, fArgs := range PROG_STATE.fieldArgs {
		if fArgs.WidgetPackage != "." {
			templateData.Imports = append(
				templateData.Imports, fArgs.WidgetPackage,
			)
		}
		if fArgs.Identity {
			templateData.IdentityFields = append(
				templateData.IdentityFields,
				IdentityFields{
					Name:   fArgs.name,
					Widget: fArgs.BaseTypeWidget,
				},
			)
		}
	}

	if err := TEMPLATES.WriteToFile(
		common.CleanFileName(fmt.Sprintf("%sWidget", INLINE_ARGS.Type)),
		common.GeneratedSrcFileExt,
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
	structType, ok := ts.Type.(*ast.StructType)
	if !ok {
		return false
	}

	name, err := common.GetSourceTextFromExpr(fSet, srcFile, ts.Name)
	if err != nil {
		common.PrintRunningError("%s", err)
		os.Exit(1)
	}

	if name != INLINE_ARGS.Type {
		return false
	}

	if PROG_STATE.shortStructGenerics, err = common.GetShortGenericsString(
		fSet, srcFile, ts,
	); err != nil {
		common.PrintRunningError("could not get generic type: %w", err)
		os.Exit(1)
	}
	if PROG_STATE.longStructGenerics, err = common.GetLongGenericsString(
		fSet, srcFile, ts,
	); err != nil {
		common.PrintRunningError("could not get generic type: %w", err)
		os.Exit(1)
	}

	for i := 0; i < structType.Fields.NumFields(); i++ {
		field := structType.Fields.List[i]
		rawDocArgs, err := common.GetDocArgVals(fSet, srcFile, field.Doc)
		if err != nil {
			common.PrintRunningError("%s", err)
			os.Exit(1)
		}

		if _, ok := rawDocArgs["identity"]; !ok {
			continue
		}

		docArgs := CommentArgs{}
		err = common.CommentArgs(&docArgs, rawDocArgs)
		if err != nil && docArgs.Identity {
			common.PrintRunningError("%s", err)
			os.Exit(1)
		}

		docArgs.name = field.Names[0].Name
		PROG_STATE.fieldArgs = append(PROG_STATE.fieldArgs, docArgs)
	}

	return true
}
