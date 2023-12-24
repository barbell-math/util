package argTypes

import "github.com/barbell-math/util/err"

var FlagDoesNotTakeArgs,IsFlagDoesNotTakeArgs=err.ErrorFactory(
    "A flag was given arguments. Flag does not accept arguments.",
)
