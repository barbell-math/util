package argparse

import "github.com/barbell-math/util/err"

var InvalidLongName,IsInvalidLongName=err.ErrorFactory(
    "The supplied long name is invalid.",
);

var InvalidShortName,IsInvalidShortName=err.ErrorFactory(
    "The supplied short name is invalid.",
);

var InvalidToken,IsInvalidToken=err.ErrorFactory(
    "A token did not parse as an argument or flag.",
)
