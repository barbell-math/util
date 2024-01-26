package dataStruct

import (
	"fmt"
	"sync"

	"github.com/barbell-math/util/algo"
	"github.com/barbell-math/util/algo/hash"
	customerr "github.com/barbell-math/util/err"
)

type (
    Set[
        K ~uint32 | ~uint64,
        V any, 
        CONSTRAINT interface { *V; algo.Equals[V]; hash.Hashable[K] },
    ] map[K]V

    SyncedSet[
        K ~uint32 | ~uint64, 
        V any, 
        CONSTRAINT interface { *V; algo.Equals[V]; hash.Hashable[K] },
    ] struct {
        *sync.RWMutex
        Set[K,V,CONSTRAINT]
    }
)

func NewSet[
    K ~uint32 | ~uint64, 
    V any, 
    CONSTRAINT interface{ *V; algo.Equals[V]; hash.Hashable[K] },
](size int) (Set[K,V,CONSTRAINT],error) {
    if size<0 {
        return Set[K, V, CONSTRAINT]{},customerr.ValOutsideRange(
            fmt.Sprintf("Size of vector must be >=0 | Have: %d",size),
        );
    }
    return Set[K,V,CONSTRAINT](make(map[K]V,size)),nil
}

func NewSyncedSet[
    K ~uint32 | ~uint64, 
    V any, 
    CONSTRAINT interface{ *V; algo.Equals[V]; hash.Hashable[K] },
](size int) (SyncedSet[K,V,CONSTRAINT],error) {
    rv,err:=NewSet[K,V, CONSTRAINT](size)
    return SyncedSet[K,V, CONSTRAINT]{
        Set: rv,
        RWMutex: &sync.RWMutex{},
    },err
}

// A empty pass through function that performs no action. Set will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedSet can simply embed a Set and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from Set.
func (v *Set[K,V,CONSTRAINT])Lock() { }

// A empty pass through function that performs no action. Set will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedSet can simply embed a Set and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from Set.
func (v *Set[K,V,CONSTRAINT])Unlock() { }

// A empty pass through function that performs no action. Set will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedSet can simply embed a Set and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from Set.
func (v *Set[K,V,CONSTRAINT])RLock() { }

// A empty pass through function that performs no action. Set will call all 
// the appropriate locking methods despite not being synced, just nothing will
// happen. This is done so that SyncedSet can simply embed a Set and
// override the appropriate locking methods to implement the correct behavior
// without needing to make any additional changes such as wrapping every single
// method from Set.
func (v *Set[K,V,CONSTRAINT])RUnlock() { }

// Vhe SyncedSet method to override the Set pass through function and 
// actually apply the mutex operation.
func (v *SyncedSet[K,V,CONSTRAINT])Lock() { v.RWMutex.Lock() }

// Vhe SyncedSet method to override the Set pass through function and 
// actually apply the mutex operation.
func (v *SyncedSet[K,V,CONSTRAINT])Unlock() { v.RWMutex.Unlock() }

// Vhe SyncedSet method to override the Set pass through function and 
// actually apply the mutex operation.
func (v *SyncedSet[K,V,CONSTRAINT])RLock() { v.RWMutex.RLock() }

// Vhe SyncedSet method to override the Set pass through function and 
// actually apply the mutex operation.
func (v *SyncedSet[K,V,CONSTRAINT])RUnlock() { v.RWMutex.RUnlock() }

func (s *Set[K,V,CONSTRAINT])Length() int {
    s.RLock()
    defer s.RUnlock()
    return len(*s)
}

func (s *Set[K, V, CONSTRAINT])Append(vals ...V) error {
    return nil
}

func (s *Set[K, V, CONSTRAINT])Insert(idx K, v ...V) error {
    return nil
}

// func (s *Set[K, V])Set(idx K, v V) error {
//     return nil
// }
