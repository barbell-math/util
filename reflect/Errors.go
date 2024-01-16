package reflect;

import (
    customerr "github.com/barbell-math/util/err"
)

var IncorrectType,IsIncorrectType=customerr.ErrorFactory(
    "An incorrect type was recieved.",
);

var InAddressableField,IsInAddressableField=customerr.ErrorFactory(
    "The address could not be calculated of a field in the given value.",
)
