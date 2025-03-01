package common

import "strings"

func ParseImports(existingImports map[string]struct{}, rawImports string) {
	for _, _import := range strings.Split(rawImports, " ") {
		// The default value provided by the arg parser is a empty string, skip
		if _import == "" {
			continue
		}

		_import = strings.Replace(_import, "->", " ", 1)
		splitImport := strings.SplitN(_import, " ", 2)
		if len(splitImport) == 2 {
			splitImport[1] = "\"" + splitImport[1] + "\""
		} else {
			splitImport[0] = "\"" + splitImport[0] + "\""
		}
		existingImports[strings.Join(splitImport, " ")] = struct{}{}
	}
}
