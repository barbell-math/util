// This package implements a csv parser that extends the one supplied by the std
// lib to include marshalling and un-marshaling to struct values. Many options
// are provided beyond what the std lib csv package presents. This package
// relies on the iterator package to create streams of structs and string slices
// which are used to represent un-marshalling and marshalling respectively.
//
// Within a csv file, string literals can be specified that follow the same
// rules as the std lib csv parser. (A quote literal can escape newlines and
// commas, a double double quote "" escapes a quote within a quote literal,
// etc...)
package csv

import (
	"encoding/csv"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/iter"
)

type (
	structIndex int
	headerIndex int
)

func getBadRowTypeError[R any](r R) error {
	return customerr.Wrap(
		customerr.IncorrectType,
		"Expected: %s Got: %s",
		reflect.Struct, reflect.TypeOf(r).Elem().Kind(),
	)
}

func getColCountError(idx int, expLen int, gotLen int, opts *options) error {
	rowNum := idx + 1
	if opts.getFlag(hasHeaders) {
		rowNum++
	}
	return customerr.Wrap(
		MalformedCSVFile,
		"Not all CSV rows had the same number of columns. Prev rows had %d elems and this row (row %d) had %d elems.",
		expLen, rowNum, gotLen,
	)
}

// TODO - still needed??
// func CSVGenerator(sep string, callback func(iter int) (string,bool)) string {
//     var sb strings.Builder;
//     var temp string;
//     cont:=true;
//     for i:=0; cont; i++ {
//         temp,cont=callback(i);
//         sb.WriteString(temp);
//         if cont {
//             sb.WriteString(sep);
//         }
//     }
//     return sb.String();
// }

// Maps a stream of string slices to a stream of strings using a csv format,
// which can then be written directly to a file. The options argument controls
// the behavior of the mapping process, see [NewOptions] for more information.
// If an error is encountered while mapping the stream of strings slices then
// iteration will stop and that error will be returned.
func Flatten(elems iter.Iter[[]string], opts *options) iter.Iter[string] {
	return iter.Next(elems,
		func(index int,
			val []string,
			status iter.IteratorFeedback,
		) (iter.IteratorFeedback, string, error) {
			if status == iter.Break {
				return iter.Break, "", nil
			}
			var sb strings.Builder
			for i, v := range val {
				sb.WriteString(v)
				if i != len(val)-1 {
					sb.WriteString(string(opts.delimiter))
				}
			}
			return iter.Continue, sb.String(), nil
		})
}

// Parses the supplied csv file to produce a stream of string slices, which can
// then be passed to the [ToStructs] function to unmarshall the data into
// structs. The options argument controls the behavior of the parsing process,
// see [NewOptions] for more information. If an error is encountered while
// processing the stream of structs then iteration will stop and that error will
// be returned.
func Parse(src string, opts *options) iter.Iter[[]string] {
	var reader *csv.Reader = nil
	file, err := os.Open(src)
	if err == nil {
		reader = csv.NewReader(file)
		reader.Comma = opts.delimiter
		reader.Comment = opts.comment
	} else {
		return iter.ValElem([]string{}, err, 1)
	}
	return func(f iter.IteratorFeedback) ([]string, error, bool) {
		if f == iter.Break || err != nil {
			file.Close()
			return []string{}, err, false
		}
		if cols, readerErr := reader.Read(); readerErr == nil {
			return cols, readerErr, true
		} else {
			if readerErr == io.EOF {
				return cols, nil, false
			}
			return []string{}, readerErr, false
		}
	}
}
