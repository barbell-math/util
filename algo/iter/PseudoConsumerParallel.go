package iter;

// This funciton is a consumer.
//
// This function will filter each value from it's parent iterator in parallel.
// The provided operation function will be run in parallel using numThreads go
// routines. The operation function will return true should return true if
// the value should be kept and false if it should not be kept. Use FilterParallel 
// when the operation that determines which values to filter takes a non-trivial
// amount of time.
func (i Iter[T])FilterParallel(op func(val T) bool, numThreads int) ([]T,error) {
    rv:=make([]T,0);
    err:=Parallel(i,func(val T) (bool,error) {
        return op(val),nil;
    },func(val T, res bool, err error){
        if res {
            rv=append(rv,val);
        }
    },numThreads);
    return rv,err;
}
