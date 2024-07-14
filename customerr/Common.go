package customerr

import (
	"errors"
	"fmt"
)

// Will continue to call the provided operation functions (ops) in the order
// that they were given until one of them returns an error. The return values of
// all previous operations will be passed as argument to the current operation
// in the order that the values were generated. Index 0 represents the return
// value from the first operation function, index 1 represent the return value
// from the second operation, etc.
func ChainedErrorOps(ops ...func(results ...any) (any, error)) error {
	var err error = nil
	res := make([]any, len(ops))
	for i := 0; err == nil && i < len(ops); i++ {
		res[i], err = ops[i](res[:i]...)
	}
	return err
}

// A run-time assert function that will panic if the given operation function
// returns an error. The panic will be given the error that the operation
// function returns. If the operation function returns nil then it will not panic.
func Assert(op func() error) {
	err := op()
	if err != nil {
		panic(err)
	}
}

// Wraps an error with a predetermined format, as shown below.
//
//	<original error>
//	  |- <wrapped information>
//
// This allows for consistent error formatting.
func Wrap(origErr error, fmtStr string, vals ...any) error {
	fmtStrWithErr := fmt.Sprintf("%%w\n  |- %s", fmtStr)
	args := []interface{}{origErr}
	return fmt.Errorf(fmtStrWithErr, append(args, vals...)...)
}

// Unwraps an error. A simple helper function to provide a clean error interface
// in this module.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func ArrayDimsArgree[N any, P any](l []N, r []P) error {
	if ll, lr := len(l), len(r); ll != lr {
		return Wrap(DimensionsDoNotAgree, "len(one)=%d len(two)=%d", ll, lr)
	}
	return nil
}

// Given a list of errors it will append them with a predetermined format, as
// shown below.
//
//	<original first error>
//	  |- <wrapped information>
//	...
//	<original nth error>
//	  |- <wrapped information>
//
// This allows for consistent error formatting. Special cases are as follows:
//   - All supplied errors are nil: The returned value will be nil.
//   - Only one of the supplied errors is nil: The returned value will be the error that is not nil.
//   - Multiple errors are not nil: The returned error will be a [MultipleErrorsOccurred] error with all of the sub-errors wrapped in it following the above format.
func AppendError(errs ...error) error {
	var rv error
	cntr := 0
	for _, e := range errs {
		if e != nil {
			if cntr == 0 {
				rv = e
			} else {
				rv = fmt.Errorf("%w\n%w", rv, e)
			}
			cntr++
		}
	}
	return rv
}
