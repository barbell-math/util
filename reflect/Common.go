// This package defines many helper functions that make it easier to get values
// from the standard libraries reflect package. To do this the [iter] package is
// used frequently.
package reflect

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/barbell-math/util/customerr"
)

type ValInfo struct {
	v    reflect.Value
	Type reflect.Type
	Kind reflect.Kind
	// The underlying value. This is a copy of the value, not the
	// original value. To access the original value use the Pntr field.
	Val func() (any, bool)
	// Returns a pointer to the value, if possible. An error will be returned if
	// the value is not addressable.
	Pntr func() (any, error)
}

var (
	InaddressableValueErr = errors.New("Inaddressable value")
)

// Makes a new val info struct with the supplied reflect value. Set keepVal to
// true to make a copy of the underlying reflect value, false to not make a
// copy. The addressableErr will be the error that is returned when trying to
// access get the address of a non-addressable reflect value.
func NewValInfo(v reflect.Value, keepVal bool, addressableErr error) ValInfo {
	return ValInfo{
		v:    v,
		Type: v.Type(),
		Kind: v.Kind(),
		Val: func() (any, bool) {
			if keepVal {
				return v.Interface(), true
			}
			return nil, false
		},
		Pntr: func() (any, error) {
			if v.CanAddr() {
				return v.Addr().Interface(), nil
			}
			return nil, addressableErr
		},
	}
}

// Checks if the supplied value can be nil. Can be used as a safe guard to check
// if the stdlib's [reflect.Value.IsNil] function will panic or not.
func CanBeNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan:
		fallthrough
	case reflect.Func:
		fallthrough
	case reflect.Interface:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Pointer:
		fallthrough
	case reflect.Slice:
		return true
	default:
		return false
	}
}

func isKindOrReflectValKind[T any, U reflect.Value | *T](
	t U,
	expType reflect.Kind,
) bool {
	switch reflect.TypeOf(t) {
	case reflect.TypeOf((*reflect.Value)(nil)).Elem():
		if refVal := any(t).(reflect.Value); refVal.Kind() == reflect.Ptr ||
			refVal.Kind() == reflect.Interface {
			return refVal.Elem().Kind() == expType
		} else {
			return refVal.Kind() == expType
		}
	default:
		return reflect.ValueOf(t).Elem().Kind() == expType
	}
}

func valError[T any, U reflect.Value | *T](
	t U,
	expKind reflect.Kind,
	kindChecker func(t U) bool,
) error {
	if kindChecker(t) {
		return nil
	}
	var fString string
	switch reflect.TypeOf(t) {
	case reflect.TypeOf((*reflect.Value)(nil)).Elem():
		if refVal := any(t).(reflect.Value); refVal.Kind() == reflect.Ptr ||
			refVal.Kind() == reflect.Interface {
			fString = fmt.Sprintf(
				"Got a reflect.Value pointer to: %s",
				refVal.Elem().Kind().String(),
			)
		} else {
			fString = fmt.Sprintf(
				"Got a reflect.Value containing: %s",
				refVal.Kind().String(),
			)
		}
	default:
		fString = fmt.Sprintf(
			"Got a pointer to: %s",
			reflect.ValueOf(t).Elem().Kind().String(),
		)
	}
	return customerr.Wrap(
		customerr.IncorrectType,
		"Function requires a %s as target. | %s",
		expKind.String(), fString,
	)
}

func homogonizeValue[T any, U reflect.Value | *T](
	t U,
	valError func(t U) error,
) (reflect.Value, error) {
	switch reflect.TypeOf(t) {
	case reflect.TypeOf((*reflect.Value)(nil)).Elem():
		if err := valError(t); err != nil {
			return reflect.Value{}, err
		}
		if refVal := any(t).(reflect.Value); refVal.Kind() == reflect.Ptr ||
			refVal.Kind() == reflect.Interface {
			return refVal.Elem(), nil
		} else {
			return refVal, nil
		}
	default:
		if err := valError(t); err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(any(t).(*T)).Elem(), nil
	}
}
