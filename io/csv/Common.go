package csv

import (
	"encoding/csv"
	"io"
	"os"
	"reflect"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/customerr"
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
    rowNum:=idx+1
    if opts.hasHeaders {
        rowNum++
    }
    return customerr.Wrap(
        MalformedCSVFile,
        "Not all CSV rows had the same number of columns. Prev rows had %d elems and this row (row %d) had %d elems.",
        expLen,rowNum,gotLen,
    )
}


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
// 
// func Flatten(elems iter.Iter[[]string], sep string) iter.Iter[string] {
//     return iter.Next(elems,
//     func(index int,
//         val []string,
//         status iter.IteratorFeedback,
//     ) (iter.IteratorFeedback, string, error) {
//         if status==iter.Break {
//             return iter.Break,"",nil;
//         }
//         var sb strings.Builder;
//         for i,v:=range(val) {
//             sb.WriteString(v);
//             if i!=len(val)-1 {
//                 sb.WriteString(sep);
//             }
//         }
//         return iter.Continue,sb.String(),nil;
//     });
// }

func Parse(src string, opts *options) iter.Iter[[]string] {
    var reader *csv.Reader=nil;
    file,err:=os.Open(src);
    if err==nil {
        reader=csv.NewReader(file);
        reader.Comma=opts.delimiter;
        reader.Comment=opts.comment;
    } else {
        return iter.ValElem([]string{},err,1);
    };
    return func(f iter.IteratorFeedback) ([]string, error, bool) {
        if f==iter.Break || err!=nil {
            file.Close();
            return []string{},err,false;
        }
        if cols,readerErr:=reader.Read(); readerErr==nil {
            return cols,readerErr,true;
        } else {
            if readerErr==io.EOF {
                return cols,nil,false;
            }
            return []string{},readerErr,false;
        }
    }
}
