package strops

import (
	"fmt"
	"strings"

	"github.com/barbell-math/util/src/customerr"
)

type (
	// Options that define how to print a table in the [WriteTable] func.
	WriteTableOpts struct {
		ColWidths     []int
		ColSeparators []bool
		RowSeparators bool
	}
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
func WriteTable(
	sb *strings.Builder,
	table [][]string,
	opts WriteTableOpts,
) error {
	if len(table) == 0 {
		return nil
	}
	if len(table[0]) != len(opts.ColWidths) {
		return customerr.Wrap(
			customerr.DimensionsDoNotAgree,
			"The number of columns in the table (%d) does not equal the number of elems in the ColWidths slice (%d)",
			len(table[0]), len(opts.ColWidths),
		)
	}
	if len(table[0])+1 != len(opts.ColSeparators) {
		return customerr.Wrap(
			customerr.DimensionsDoNotAgree,
			"The number of columns in the table plus one (%d) does not equal the number of elems in the ColSeparators slice (%d)",
			len(table[0])+1, len(opts.ColSeparators),
		)
	}

	fmtDirectives := make([]string, len(opts.ColWidths))
	rowSeparators := make([]string, len(opts.ColWidths))
	for i := 0; i < len(opts.ColWidths); i++ {
		leadingColSeparator := opts.ColSeparators[i]
		trailingColSeparator := false
		if i+1 == len(opts.ColWidths) {
			trailingColSeparator = opts.ColSeparators[i+1]
		}

		if leadingColSeparator && trailingColSeparator {
			fmtDirectives[i] = fmt.Sprintf("| %%-%ds |", opts.ColWidths[i])
			rowSeparators[i] = fmt.Sprintf(
				"|-%s-|", strings.Repeat("-", opts.ColWidths[i]),
			)
		} else if leadingColSeparator {
			fmtDirectives[i] = fmt.Sprintf("| %%-%ds ", opts.ColWidths[i])
			rowSeparators[i] = fmt.Sprintf(
				"|-%s-", strings.Repeat("-", opts.ColWidths[i]),
			)
		} else {
			fmtDirectives[i] = fmt.Sprintf("%%-%ds ", opts.ColWidths[i])
			rowSeparators[i] = fmt.Sprintf(
				"-%s", strings.Repeat("-", opts.ColWidths[i]),
			)
		}
	}

	splitRow := make([][]string, len(opts.ColWidths))
	for _, row := range table {
		for j, _ := range splitRow {
			splitRow[j] = WordWrapLines(row[j], opts.ColWidths[j])
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

		for j := 0; j < len(splitRow) && opts.RowSeparators; j++ {
			sb.WriteString(rowSeparators[j])
		}
		sb.WriteByte('\n')
	}

	return nil
}
