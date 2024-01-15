package reflect

import (
	"fmt"
	"reflect"
)

func isKindOrReflectValKind[T any](t *T, expType reflect.Kind) bool {
    switch reflect.TypeOf(t) {
        case reflect.TypeOf(&reflect.Value{}):
            if refVal:=any(t).(*reflect.Value); refVal.Kind()==reflect.Ptr {
                return refVal.Elem().Kind()==expType
            } else {
                return refVal.Kind()==expType
            }
        default: return reflect.ValueOf(t).Elem().Kind()==expType
    }
}

func valError[T any](
    t *T, 
    expKind reflect.Kind, 
    kindChecker func(t *T) bool,
) error {
    if kindChecker(t) {
        return nil;
    }
    var fString string
    switch reflect.TypeOf(t) {
        case reflect.TypeOf(&reflect.Value{}):
            if refVal:=any(t).(*reflect.Value); refVal.Kind()==reflect.Ptr {
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
