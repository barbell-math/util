package csv

import (
	stdReflect "reflect"
	"strconv"
	"time"

	"github.com/barbell-math/util/src/customerr"
	"github.com/barbell-math/util/src/iter"
	"github.com/barbell-math/util/src/reflect"
)

// Takes an iterator stream of string slices, treats each slice as a row of a
// csv file, and maps that stream of slices to a stream of structs with the
// type specified by the generic R type. The options argument controls the
// behavior of the mapping process, see [NewOptions] for more information. If an
// error is encountered while processing the stream of slices then iteration
// will stop and that error will be returned.
//
// Data is placed in struct fields based on either the CSV headers or the
// ordering of the fields in the struct, depending on the options passed to the
// function. Any fields of the struct that are not set given the header mapping
// will be left zero-initilized. Any blank values in the CSV file will also be
// zero initilized.
//
// When parsing data the following data types are supported:
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
func ToStructs[R any](src iter.Iter[[]string], opts *options) iter.Iter[R] {
	var tmp R
	if !reflect.IsStructVal[R](&tmp) {
		return iter.ValElem[R](
			tmp,
			customerr.AppendError(getBadRowTypeError(&tmp), src.Stop()),
			1,
		)
	}

	structHeaderMapping, err := newStructHeaderMapping[R](&tmp, opts)
	if err != nil {
		return iter.ValElem[R](
			tmp,
			customerr.AppendError(MalformedCSVStruct, err, src.Stop()),
			1,
		)
	}

	idxMapping, err := newHeaderToStructIndexMapping[R](src, structHeaderMapping, &tmp, opts)
	if err != nil {
		return iter.ValElem[R](
			tmp,
			customerr.AppendError(MalformedCSVFile, err, src.Stop()),
			1,
		)
	}

	prevLen := len(idxMapping)
	return iter.Next[[]string, R](
		src,
		func(
			index int,
			val []string,
			status iter.IteratorFeedback,
		) (iter.IteratorFeedback, R, error) {
			if status == iter.Break {
				return iter.Break, tmp, nil
			}
			// If there were no headers then there is an incorrect header cnt,
			// so set it to the correct value here.
			if index == 0 && !opts.GetFlag(hasHeaders) {
				prevLen = len(val)
			}
			if len(val) != prevLen {
				return iter.Break, tmp, getColCountError(
					index, prevLen, len(val), opts,
				)
			}
			if rv, err := rowToStruct[R](idxMapping, val, opts); err == nil {
				return iter.Continue, rv, nil
			} else {
				return iter.Break, rv, customerr.AppendError(
					customerr.Wrap(MalformedCSVFile, "Line: %d", index+1),
					err,
				)
			}
		},
	)
}

func rowToStruct[R any](
	idxMapping headerToStructIndexMapping[R],
	data []string,
	opts *options,
) (R, error) {
	var rv R
	var err error = nil
	for i := 0; err == nil && i < len(data); i++ {
		if len(data[i]) > 0 {
			err = setStructValue(&rv, int(idxMapping[headerIndex(i)]), data[i], opts)
		}
	}
	return rv, err
}

func setStructValue[R any](
	row *R,
	fieldIdx int,
	val string,
	opts *options,
) error {
	var err error = nil
	f := stdReflect.ValueOf(row).Elem().Field(fieldIdx)
	if f.IsValid() && f.CanSet() {
		if tmp, ok := f.Interface().(time.Time); ok {
			tmp, err = time.Parse(opts.dateTimeFormat, val)
			f.Set(stdReflect.ValueOf(tmp))
			return err
		}

		switch f.Kind() {
		case stdReflect.Bool:
			var tmp bool
			tmp, err = strconv.ParseBool(val)
			f.SetBool(tmp)
		case stdReflect.Uint:
			err = setUint[uint](f, val)
		case stdReflect.Uint8:
			err = setUint[uint8](f, val)
		case stdReflect.Uint16:
			err = setUint[uint16](f, val)
		case stdReflect.Uint32:
			err = setUint[uint32](f, val)
		case stdReflect.Uint64:
			err = setUint[uint64](f, val)
		case stdReflect.Int:
			err = setInt[int](f, val)
		case stdReflect.Int8:
			err = setInt[int8](f, val)
		case stdReflect.Int16:
			err = setInt[int16](f, val)
		case stdReflect.Int32:
			err = setInt[int32](f, val)
		case stdReflect.Int64:
			err = setInt[int64](f, val)
		case stdReflect.Float32:
			err = setFloat[float32](f, val)
		case stdReflect.Float64:
			err = setFloat[float32](f, val)
		case stdReflect.String:
			f.SetString(val)
		default:
			err = customerr.Wrap(customerr.UnsupportedType, "'%s'", f.Kind())
		}
	} else {
		err = customerr.Wrap(
			InvalidStructField,
			"Field: %s",
			stdReflect.TypeOf(row).Elem().Field(fieldIdx).Name,
		)
	}
	return err
}

func setUint[N ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](
	f stdReflect.Value,
	v string,
) error {
	tmp, err := strconv.ParseUint(v, 10, 64)
	f.SetUint(tmp)
	return err
}
func setInt[N ~int | ~int8 | ~int16 | ~int32 | ~int64](
	f stdReflect.Value,
	v string,
) error {
	tmp, err := strconv.ParseInt(v, 10, 64)
	f.SetInt(tmp)
	return err
}
func setFloat[N ~float32 | ~float64](f stdReflect.Value, v string) error {
	tmp, err := strconv.ParseFloat(v, 64)
	f.SetFloat(tmp)
	return err
}
