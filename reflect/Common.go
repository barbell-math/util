package reflect

import (
	"fmt"
	"reflect"
)

func isKindOrReflectValKind[T any, U reflect.Value | *T](
    t U, 
    expType reflect.Kind,
) bool {
    switch reflect.TypeOf(t) {
        case reflect.TypeOf(reflect.Value{}):
            if refVal:=any(t).(reflect.Value); refVal.Kind()==reflect.Ptr {
                return refVal.Elem().Kind()==expType
            } else {
                return refVal.Kind()==expType
            }
        default: return reflect.ValueOf(t).Elem().Kind()==expType
    }
}

func valError[T any, U reflect.Value | *T](
    t U, 
    expKind reflect.Kind, 
    kindChecker func(t U) bool,
) error {
    if kindChecker(t) {
        return nil;
    }
    var fString string
    switch reflect.TypeOf(t) {
        case reflect.TypeOf(reflect.Value{}):
            if refVal:=any(t).(reflect.Value); refVal.Kind()==reflect.Ptr {
                fString=fmt.Sprintf(
                    "Got a reflect.Value pointer to: %s",
                    refVal.Elem().Kind().String(),
                )
            } else {
                fString=fmt.Sprintf(
                    "Got a reflect.Value containing: %s",
                    refVal.Kind().String(),
                )
            }
        default:
            fString=fmt.Sprintf(
                "Got a pointer to: %s",
                reflect.ValueOf(t).Elem().Kind().String(),
            )
    }
    return IncorrectType(fmt.Sprintf(
        "Function requires a %s as target. | %s",
        expKind.String(),fString,
    ));
}

func homogonizeValue[T any, U reflect.Value | *T](
    t U,
    valError func(t U) error,
) (reflect.Value,error) {
    switch reflect.TypeOf(t) {
        case reflect.TypeOf(reflect.Value{}):
            if err:=valError(t); err!=nil {
                return reflect.Value{},err
            }
            if refVal:=any(t).(reflect.Value); refVal.Kind()==reflect.Ptr {
                return refVal.Elem(),nil
            } else {
                return refVal,nil
            }
        default:
            if err:=valError(t); err!=nil {
                return reflect.Value{},err
            }
            return reflect.ValueOf(any(t).(*T)).Elem(),nil
    }
}
