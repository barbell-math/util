package symbolic;

import (
    customerr "github.com/barbell-math/util/err"
)

var InvalidOperation,IsInvalidOperation=customerr.ErrorFactory(
    "The supplied types cannot perform the requested operation.",
);
