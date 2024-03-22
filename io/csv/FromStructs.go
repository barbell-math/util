package csv

import (
	"fmt"
	stdReflect "reflect"
	"time"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/reflect"
)

func FromStructs[R any](src iter.Iter[R], opts *options) iter.Iter[[]string] {
    var tmp R;
    if !reflect.IsStructVal[R](&tmp) {
        return iter.ValElem[[]string](
            []string{},
            customerr.AppendError(getBadRowTypeError(&tmp),src.Stop()),
            1,
        )
    }

    structHeaderMapping,err:=newStructHeaderMapping[R](&tmp,opts)
    if err!=nil {
        return iter.ValElem[[]string](
            []string{},
            customerr.AppendError(MalformedCSVStruct,err,src.Stop()),
            1,
        )
    }

    idxMapping,err:=newStructToHeaderIndexMapping[R](structHeaderMapping,&tmp,opts)
    if err!=nil {
        return iter.ValElem[[]string](
            []string{},
            customerr.AppendError(MalformedCSVStruct,err,src.Stop()),
            1,
        )
    }

    return iter.Next[R,[]string](
        src,
        func(
            index int,
            val R,
            status iter.IteratorFeedback,
        ) (iter.IteratorFeedback, []string, error) {
            if status==iter.Break {
                return iter.Break,[]string{},nil
            }
            rv:=make([]string,len(idxMapping))
            for sIdx,hIdx:=range(idxMapping) {
                if v,err:=getValAsString[R](&val,sIdx,opts); err==nil {
                    rv[hIdx]=v
                } else {
                    return iter.Break,rv,err
                }
            }
            return iter.Continue,rv,nil
        },
    ).Inject(func(idx int, val []string, injectedPrev bool) ([]string, error, bool) {
        if idx>0 || !opts.writeHeaders {
            return []string{},nil,false
        }
        if opts.headersSupplied {
            return opts.headers,nil,true
        }
        v:=stdReflect.TypeOf(&tmp).Elem()
        rv:=make([]string,len(idxMapping))
        for sIdx,hIdx:=range(idxMapping) {
            rv[hIdx]=v.Field(int(sIdx)).Name
        }
        return rv,nil,true
    })
}

func getValAsString[R any](r *R, sIdx structIndex, opts *options) (string,error) {
    v:=stdReflect.ValueOf(r).Elem().Field(int(sIdx))
    if v.Type()==stdReflect.TypeOf((*time.Time)(nil)).Elem() {
        return v.Interface().(time.Time).Format(opts.dateTimeFormat),nil
    }
    switch v.Kind() {
        case stdReflect.Bool: fallthrough
        case stdReflect.Uint: fallthrough
        case stdReflect.Uint8: fallthrough
        case stdReflect.Uint16: fallthrough
        case stdReflect.Uint32: fallthrough
        case stdReflect.Uint64: fallthrough
        case stdReflect.Int: fallthrough
        case stdReflect.Int8: fallthrough
        case stdReflect.Int16: fallthrough
        case stdReflect.Int32: fallthrough
        case stdReflect.Int64: fallthrough
        case stdReflect.Float32: fallthrough
        case stdReflect.Float64: fallthrough
        case stdReflect.String: return fmt.Sprintf("%v",v.Interface()),nil
        default: return "",customerr.Wrap(
            customerr.UnsupportedType,
            "Struct field: %s Type: %s",
            stdReflect.TypeOf(r).Elem().Field(int(sIdx)).Name,v.Kind().String(),
        )
    }
}
