package common

import (
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
	"text/template"
)

type (
	GeneratedFilesRegistry struct {
		template.Template
	}
)

const (
	GeneratedFileExt string = ".gen.go"
)

func NewGeneratedFilesRegistryFromMap(
	vals map[string]string,
) GeneratedFilesRegistry {
	rv := GeneratedFilesRegistry{}
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
	templateName string,
	data any,
) error {
	if strings.Contains(outFile, ".") {
		return fmt.Errorf(
			"%w: generated files should not include an extension.",
			InvalidGeneratorFileName,
		)
	}

	fileName := fmt.Sprintf("%s%s", outFile, GeneratedFileExt)

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
	printer.Fprint(f, fSet, ast)
	return nil
}
