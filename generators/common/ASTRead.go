package common

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"strings"
)

func GenFileExclusionFilter(f fs.FileInfo) bool {
	return (!strings.HasSuffix(f.Name(), string(GeneratedSrcFileExt)) &&
		!strings.HasSuffix(f.Name(), string(GeneratedTestFileExt)))
}

func IterateOverAST(
	path string,
	filter func(f fs.FileInfo) bool,
	op func(
		fSet *token.FileSet,
		file *ast.File,
		srcFile *os.File,
		node ast.Node,
	) bool,
) {
	fSet := token.NewFileSet()
	packs, err := parser.ParseDir(fSet, path, filter, parser.ParseComments)
	if err != nil {
		PrintRunningError("Failed to parse package:", err)
		os.Exit(1)
	}

	for _, pack := range packs {
		for fileName, f := range pack.Files {
			srcFile, err := os.OpenFile(fileName, os.O_RDONLY, 666)
			if err != nil {
				PrintRunningError(
					"Could not open file %s to parse source.",
					fileName,
				)
			}
			ast.Inspect(f, func(n ast.Node) bool {
				if n != nil {
					return op(fSet, f, srcFile, n)
				}
				return false
			})
			srcFile.Close()
		}
	}
}

func GetSourceTextFromExpr(
	fSet *token.FileSet,
	srcFile *os.File,
	field ast.Node,
) (string, error) {
	if field == nil {
		return "", nil
	}

	if _, err := srcFile.Seek(
		int64(fSet.Position(field.Pos()).Offset),
		0,
	); err != nil {
		return "", err
	}
	src := make([]byte, field.End()-field.Pos())
	if _, err := srcFile.Read(src); err != nil {
		return "", err
	}
	return string(src), nil
}

func GetComment(
	fSet *token.FileSet,
	srcFile *os.File,
	doc *ast.CommentGroup,
	comment *ast.CommentGroup,
) (string, error) {
	var err error
	origComments := ""
	if doc != nil {
		origComments, err = GetSourceTextFromExpr(
			fSet, srcFile, doc,
		)
		if err != nil {
			return "", err
		}
	}

	if comment != nil {
		temp, err := GetSourceTextFromExpr(
			fSet, srcFile, comment,
		)
		origComments += temp
		if err != nil {
			return "", err
		}
	}

	rvLines := strings.Split(origComments, "\n")
	for i, _ := range rvLines {
		rvLines[i] = strings.TrimSpace(rvLines[i])
	}
	return strings.Join(rvLines, "\n"), nil
}
