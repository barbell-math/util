package csv;

import (
    customerr "github.com/barbell-math/util/err"
)

var MalformedCSVFile,IsMalformedCSVFile=customerr.ErrorFactory(
    "The CSV file cannot be converted to the requested struct.",
);

var NonStructValue,IsNonStructValue=customerr.ErrorFactory(
    "A struct value was expected but was not recieved.",
);

var UnsupportedType,IsUnsupportedType=customerr.ErrorFactory(
    "The supplied type is not supported.",
);
