package csv

import (
	"fmt"
	stdReflect "reflect"
	"strconv"
	"time"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/dataStruct"
	"github.com/barbell-math/util/dataStruct/types/static"
	customerr "github.com/barbell-math/util/err"
	"github.com/barbell-math/util/reflect"
)

type Options struct {
    dtFormat string
    useJsonTag bool
}

func CSVToStruct[R any](src iter.Iter[[]string], opts Options) iter.Iter[R] {
    var tmp R;
    if reflect.IsStructVal[R](&tmp) {
        return iter.ValElem[R](tmp,customerr.IncorrectType(
            fmt.Sprintf("Expected: Struct Got: %s",stdReflect.TypeOf(tmp)),
        ),1)
    }
    colNames,err:=validColNames[R](&tmp,&opts)
    if err!=nil {
        return iter.ValElem[R](tmp,err,1)
    }
    return iter.Next[[]string,R](src,func(
        index int, 
        val []string, 
        status iter.IteratorFeedback,
    ) (iter.IteratorFeedback, R, error) {
        if status==iter.Break {
            return iter.Break,tmp,nil;
        }
        if index==0 {
            if err:=validateHeaders(val,colNames); err!=nil {
                return iter.Break,tmp,err
            }
            return iter.Iterate,tmp,nil;
        }
    })
}

func validColNames[R any](s *R, opts *Options) ([]string,error) {
    return iter.Map[static.Pair[string,stdReflect.StructTag],string](
        iter.Zip[string,stdReflect.StructTag](
            reflect.StructFieldNames[R](s),
            reflect.StructFieldTags[R](s),
            func() static.Pair[string, stdReflect.StructTag] {
                return &dataStruct.Pair[string,stdReflect.StructTag]{}
            },
        ),
        func(index int, val static.Pair[string, stdReflect.StructTag]) (string, error) {
            if opts.useJsonTag {
                if v,ok:=val.GetB().Lookup("json"); ok {
                    return v,nil
                }
            }
            if v,ok:=val.GetB().Lookup("csv"); ok {
                return v,nil
            }
            return val.GetA(),nil
        },
    ).Collect()
}

// Same length and contents - compare sets
func validateHeaders(headers []string, colNames []string) error {
    
}
