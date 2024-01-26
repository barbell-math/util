package containerTypes


// The type of values that the containers will act upon. This interface enforces
// all the required information is exposed by the underlying types held in the 
// container.
type Widget[T any] interface {
    Eq(l *T, r *T) bool
    Lt(l *T, r *T) bool
    Unwrap() T
}

type RWSyncable interface {
    Lock()
    Unlock()
    RLock()
    RUnlock()
}
