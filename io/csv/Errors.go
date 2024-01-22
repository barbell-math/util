package csv;

import (
    customerr "github.com/barbell-math/util/err"
)

var MalformedCSVFile,IsMalformedCSVFile=customerr.ErrorFactory(
    "The CSV file cannot be converted to the requested struct.",
);
