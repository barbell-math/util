package reflect;

import (
    customerr "github.com/barbell-math/util/err"
)

var NonStructValue,IsNonStructValue=customerr.ErrorFactory(
    "A struct value was expected but was not recieved.",
);

var InAddressableField,IsInAddressableField=customerr.ErrorFactory(
    "The address could not be calculated of a field in the given struct.",
)
