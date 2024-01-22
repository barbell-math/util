package csv;

import (
	"fmt"
	stdReflect "reflect"
	"time"
	"unicode"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/reflect"
	customerr "github.com/barbell-math/util/err"
)

func StructToCSV[R any](elems iter.Iter[R],
        addHeaders bool,
        timeDateFormat string) iter.Iter[[]string] {
    var tmp R;
    if reflect.IsStructVal[R](&tmp) {
        return iter.ValElem(
            []string{},
            customerr.IncorrectType(
                fmt.Sprintf("Expected: Struct Got: %s",stdReflect.TypeOf(tmp)),
            ),
            1,
        );
    }
    capFilter:=func(idx int, thing string) bool {
        return len(thing)>0 && unicode.IsUpper(rune(thing[0]));
    }
    headers,_:=reflect.StructFieldNames[R](&tmp).Filter(capFilter).Collect()
    return iter.Next(elems,
    func(index int,
        val R,
        status iter.IteratorFeedback,
    ) (iter.IteratorFeedback, []string, error) {
        if status==iter.Break {
            return iter.Break,[]string{},nil;
        }
        sVals,_:=reflect.GetStructVals(&val,capFilter);
        valsAsStr,err:=getValsAsString(sVals,timeDateFormat);
        return iter.Continue,valsAsStr,err;
    }).Inject(func(idx int, val []string) ([]string,bool) {
        return headers,idx==0 && addHeaders;
    });
}

func getValsAsString(reflectVals []stdReflect.Value, timeDateFormat string) ([]string,error){
    valsAsStr:=make([]string,len(reflectVals));
    for i,v:=range(reflectVals) {
        if iterS,err:=getStringFromStructVal(v.Interface(),timeDateFormat); err==nil {
            valsAsStr[i]=iterS;
        } else {
            return []string{},err;
        }
    }
    return valsAsStr,nil;
}

//Only basic types are supported
func getStringFromStructVal[T any](val T, timeDateFormat string) (string,error) {
    switch any(val).(type) {
        case time.Time: return any(val).(time.Time).Format(timeDateFormat),nil;
        case bool: return fmt.Sprintf("%v",val),nil;
        case uint: return fmt.Sprintf("%v",val),nil;
        case uint8: return fmt.Sprintf("%v",val),nil;
        case uint16: return fmt.Sprintf("%v",val),nil;
        case uint32: return fmt.Sprintf("%v",val),nil;
        case uint64: return fmt.Sprintf("%v",val),nil;
        case int: return fmt.Sprintf("%v",val),nil;
        case int8: return fmt.Sprintf("%v",val),nil;
        case int16: return fmt.Sprintf("%v",val),nil;
        case int32: return fmt.Sprintf("%v",val),nil;
        case int64: return fmt.Sprintf("%v",val),nil;
        case float32: return fmt.Sprintf("%v",val),nil;
        case float64: return fmt.Sprintf("%v",val),nil;
        case string: return fmt.Sprintf("%v",val),nil;
        default: return "",UnsupportedType(fmt.Sprintf(
            "'%s'",stdReflect.TypeOf(val),
        ));
    }
}

