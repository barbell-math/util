package csv

import (
	"fmt"
	stdReflect "reflect"
	"strings"
	"time"

	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/iter"
	"github.com/barbell-math/util/src/reflect"
)

// Takes an iterator stream of structs with type specified by the generic R
// type, treats each struct as a row of a row of a csv file, and maps that
// stream of struts to a stream of string slices. The options argument controls
// the behavior of the mapping process, see [NewOptions] for more information.
// If an error is encountered while processing the stream of structs then
// iteration will stop and that error will be returned.
//
// The data that is placed in the csv file based either on the struct field
// names, the struct field tags, or a combination of both depending on the
// options passed as well as what is available from the struct definition. The
// ordering of the columns is either determined by the ordering of the struct
// fields or by the options passed.
//
// When writing the data to the file, the following data types are supported:
//
//   - All integer types (int,int8,int16,int32,int64)
//   - All unsigned integer types (uint,uint8,uint16,uint32,uint64)
//   - All float types (float32,float64)
//   - Strings
//   - Booleans
//   - Date time formats (format specified by options)
//
// The above data types represent the basic types present within the language,
// and this mirrors the limitations of a CSV file. For more expressive
// representations of arrays and objects consider using JSON instead.
func FromStructs[R any](src iter.Iter[R], opts *options) iter.Iter[[]string] {
	var tmp R
	if !reflect.IsStructVal[R](&tmp) {
		return iter.ValElem[[]string](
			[]string{},
			customerr.AppendError(getBadRowTypeError(&tmp), src.Stop()),
			1,
		)
	}

	structHeaderMapping, err := newStructHeaderMapping[R](&tmp, opts)
	if err != nil {
		return iter.ValElem[[]string](
			[]string{},
			customerr.AppendError(MalformedCSVStruct, err, src.Stop()),
			1,
		)
	}

	idxMapping, err := newStructToHeaderIndexMapping[R](structHeaderMapping, &tmp, opts)
	if err != nil {
		return iter.ValElem[[]string](
			[]string{},
			customerr.AppendError(MalformedCSVStruct, err, src.Stop()),
			1,
		)
	}

	return iter.Next[R, []string](
		src,
		func(
			index int,
			val R,
			status iter.IteratorFeedback,
		) (iter.IteratorFeedback, []string, error) {
			if status == iter.Break {
				return iter.Break, []string{}, nil
			}
			rv := make([]string, len(idxMapping))
			for sIdx, hIdx := range idxMapping {
				if v, err := getValAsString[R](&val, sIdx, opts); err == nil {
					rv[hIdx] = v
				} else {
					return iter.Break, rv, err
				}
			}
			return iter.Continue, rv, nil
		},
	).Inject(func(idx int, val []string, injectedPrev bool) ([]string, error, bool) {
		if idx > 0 || !opts.GetFlag(writeHeaders) {
			return []string{}, nil, false
		}
		if opts.GetFlag(headersSupplied) {
			return opts.headers, nil, true
		}
		rv := make([]string, len(idxMapping))
		for name, sIdx := range structHeaderMapping {
			// Not all struct fields will be mapped to a header, so no error here
			if hIdx, ok := idxMapping[sIdx]; ok {
				rv[hIdx] = name
			}
		}
		return rv, nil, true
	})
}

func getValAsString[R any](r *R, sIdx structIndex, opts *options) (string, error) {
	fmtStr := "%v"
	v := stdReflect.ValueOf(r).Elem().Field(int(sIdx))
	actualVal := v.Interface()

	if v.Type() == stdReflect.TypeOf((*time.Time)(nil)).Elem() {
		if actualVal.(time.Time).Equal(time.Time{}) && !opts.GetFlag(writeZeroValues) {
			return "", nil
		}
		return v.Interface().(time.Time).Format(opts.dateTimeFormat), nil
	}

	zeroVal := stdReflect.Zero(v.Type())
	switch v.Kind() {
	case stdReflect.Bool:
		fmtStr = "%t"
	case stdReflect.Uint:
		fallthrough
	case stdReflect.Uint8:
		fallthrough
	case stdReflect.Uint16:
		fallthrough
	case stdReflect.Uint32:
		fallthrough
	case stdReflect.Uint64:
		fallthrough
	case stdReflect.Int:
		fallthrough
	case stdReflect.Int8:
		fallthrough
	case stdReflect.Int16:
		fallthrough
	case stdReflect.Int32:
		fallthrough
	case stdReflect.Int64:
		fmtStr = "%d"
	case stdReflect.Float32:
		fallthrough
	case stdReflect.Float64:
		fmtStr = "%f"
	case stdReflect.String:
		fmtStr = "%s"
		if len(actualVal.(string)) == 0 {
			break
		}
		if actualVal.(string)[0] == '"' && actualVal.(string)[len(actualVal.(string))-1] == '"' {
			actualVal = fmt.Sprintf(
				"\"\"\"%s\"\"\"",
				strings.ReplaceAll(
					actualVal.(string)[1:len(actualVal.(string))-1],
					"\"",
					"\"\"",
				),
			)
		} else if strings.ContainsAny(actualVal.(string), ",\n") {
			actualVal = fmt.Sprintf(
				"\"%s\"",
				strings.ReplaceAll(
					actualVal.(string),
					"\"",
					"\"\"",
				),
			)
		}
	default:
		return "", customerr.Wrap(
			customerr.UnsupportedType,
			"Struct field: %s Type: %s",
			stdReflect.TypeOf(r).Elem().Field(int(sIdx)).Name, v.Kind().String(),
		)
	}
	if actualVal == zeroVal.Interface() && !opts.GetFlag(writeZeroValues) {
		return "", nil
	}
	return fmt.Sprintf(fmtStr, actualVal), nil
}
