package err

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
func ChainedErrorOps(ops ...func(results ...any) (any,error)) error {
    var err error=nil;
    res:=make([]any,len(ops));
    for i:=0; err==nil && i<len(ops); i++ {
        res[i],err=ops[i](res[:i]...);
    }
    return err;
}

// A run-time assert function that will panic if the given operation function
// returns an error. The panic will be given the error that the operation
// function returns. If the operation function returns nil then it will not panic.
func Assert(op func() error){
    err:=op();
    if err!=nil {
        panic(err);
    }
}

// Wraps an error with a predetermined format, as shown below.
//
//  <original error>
//    |- <wrapped information>
//
// This allows for consistent error formatting.
func Wrap(origErr error, fmtStr string, vals ...any) error {
    fmtStrWithErr:=fmt.Sprintf("%%w\n  |- %s\n",fmtStr)
    args:=[]interface{}{origErr}
    return fmt.Errorf(fmtStrWithErr,append(args,vals...)...)
}

// Unwraps an error. A simple helper function to provide a clean error interface
// in this module.
func Unwrap(err error) error {
    return errors.Unwrap(err)
}

func ArrayDimsArgree[N any, P any](l []N, r []P) error {
    if ll,lr:=len(l),len(r); ll!=lr {
        return Wrap(DimensionsDoNotAgree, "len(one)=%d len(two)=%d", ll,lr);
    }
    return nil;
}

// Given two errors it will append them with a predetermined format, as shown
// below.
//
//  Multiple errors have occured.
//  First error: <original first error>
//    |- <wrapped information>
//  Second error: <original second error>
//    |- <wrapped information>
//
// This allows for consistent error formatting. If either first or second is nil
// then second or first will be respectively returned without wrapping them with
// additional information. If both first and second are nil then nil will be
// returned.
func AppendError(first error, second error) error {
    if second!=nil && first==nil {
        first=second;
    } else if second!=nil && first!=nil {
        return fmt.Errorf(
            "Multiple errors have occurred.\nFirst error: %w\nSecond error: %w\n",
            first,second,
        )
    }
    return first;
}
