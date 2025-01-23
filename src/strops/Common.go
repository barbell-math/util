package strops

import (
	"fmt"
	"strings"

	"github.com/barbell-math/util/src/customerr"
)

// Breaks a body of text into separate lines that are no longer than maxLen.
// Line splits are made on word (a.k.a. white space delimited) boundaries.
func WordWrapLines(text string, maxLen int) []string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return []string{text}
	}

	line := words[0]
	wrapped := []string{}
	spaceLeft := maxLen - len(line)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped = append(wrapped, line)
			line = word
			spaceLeft = maxLen - len(word)
		} else {
			line += " " + word
			spaceLeft -= len(word) + 1
		}
	}
	if len(line) > 0 {
		wrapped = append(wrapped, line)
	}
	return wrapped
}

// Writes a table to the supplied string builder given the supplied data and
// column widths. All data within the rows will be split using [WordWrapLines].
func WriteTable(sb *strings.Builder, table [][]string, colWidths []int) error {
	if len(table) == 0 {
		return nil
	}
	if len(table[0]) != len(colWidths) {
		return customerr.Wrap(
			customerr.DimensionsDoNotAgree,
			"The number of columns in the table (%d) does not equal the number of elems in the colWidths slice (%d)",
			len(table[0]), len(colWidths),
		)
	}

	fmtDirectives := make([]string, len(colWidths))
	for i := 0; i < len(colWidths); i++ {
		fmtDirectives[i] = fmt.Sprintf("%%-%ds ", colWidths[i])
	}

	splitRow := make([][]string, len(colWidths))
	for _, row := range table {
		for j, _ := range splitRow {
			splitRow[j] = WordWrapLines(row[j], colWidths[j])
		}

		for lineCntr := 0; ; lineCntr++ {
			somethingToPrint := false
			for j := 0; j < len(splitRow) && !somethingToPrint; j++ {
				if len(splitRow[j]) > lineCntr {
					somethingToPrint = true
				}
			}
			if !somethingToPrint {
				break
			}
			for j := 0; j < len(splitRow); j++ {
				if len(splitRow[j]) > lineCntr {
					sb.WriteString(fmt.Sprintf(
						fmtDirectives[j], splitRow[j][lineCntr],
					))
				} else {
					sb.WriteString(fmt.Sprintf(fmtDirectives[j], ""))
				}
			}
			sb.WriteByte('\n')
		}
	}

	return nil
}
