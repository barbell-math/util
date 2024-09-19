package common

import (
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"text/template"
)

type (
	GeneratedExtension     string
	GeneratedFilesRegistry struct {
		template.Template
	}
)

const (
	GeneratedSrcFileExt      GeneratedExtension = ".gen.go"
	GeneratedTestFileExt     GeneratedExtension = ".gen_test.go"
	GeneratedCommentTemplate string             = "autoGenComment"
	AutoGeneratedTemplate    string             = "// Code generated by {{ .GeneratorName }} - DO NOT EDIT."
)

func NewGeneratedFilesRegistryFromMap(
	vals map[string]string,
) GeneratedFilesRegistry {
	rv := GeneratedFilesRegistry{}
	rv.New(GeneratedCommentTemplate).Parse(AutoGeneratedTemplate)
	for n, t := range vals {
		if _, err := rv.New(n).Parse(t); err != nil {
			PrintRunningError("%s", err)
			os.Exit(1)
		}
	}
	return rv
}

func (g *GeneratedFilesRegistry) WriteToFile(
	outFile string,
	extension GeneratedExtension,
	templateName string,
	data any,
) error {
	if strings.Contains(outFile, ".") {
		return fmt.Errorf(
			"%w: generated files should not include an extension.",
			InvalidGeneratorFileName,
		)
	}

	fileName := fmt.Sprintf("%s%s", outFile, extension)

	sb := strings.Builder{}
	if err := g.Template.ExecuteTemplate(&sb, templateName, data); err != nil {
		return fmt.Errorf("Template error: %w", err)
	}
	fSet := token.NewFileSet()
	ast, err := parser.ParseFile(fSet, fileName, sb.String(), parser.ParseComments)
	if err != nil {
		return fmt.Errorf("Parser error: %w", err)
	}

	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("File error: %w", err)
	}
	format.Node(f, fSet, ast)
	return nil
}