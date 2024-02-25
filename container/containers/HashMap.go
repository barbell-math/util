package containers

import (
	"sync"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containerTypes"
)

type (
    internalHashMapImpl[K any, V any] map[uint64]basic.Pair[K,V]

    // A type to represent a map that dynamically grows as key value pairs are 
    // added. The set will maintain uniqueness and is internally implemented 
    // with a hashing method. The type constraints on the generics define the 
    // logic for how value specific operations, such as equality comparisons, 
    // will be handled.
    HashMap[
        K any,
        V any,
        KI widgets.WidgetInterface[K],
        VI widgets.WidgetInterface[V],
    ] struct {
        internalHashMapImpl[K,V]
    }

    // A synchronized version of HashMap. All operations will be wrapped in the
    // appropriate calls the embedded RWMutex. A pointer to a RWMutex is 
    // embedded rather than a value to avoid copying the lock value.
    SyncedHashMap[
        K any,
        V any,
        KI widgets.WidgetInterface[K],
        VI widgets.WidgetInterface[V],
    ] struct {
        *sync.RWMutex
        HashMap[K,V,KI,VI]
    }
)

// Creates a new map initialized with enough memory to hold size elements. 
// Size must be >= 0, an error will be returned if it is not. If size is 0 the 
// map will be initialized with 0 elements.
func NewHashMap[
    K any, 
    V any, 
    KI widgets.WidgetInterface[K], 
    VI widgets.WidgetInterface[V],
](size int) (HashMap[K,V,KI,VI],error) {
    if size<0 {
        return HashMap[K, V, KI, VI]{},getSizeError(size)
    }
    return HashMap[K, V, KI, VI]{
        internalHashMapImpl: make(internalHashMapImpl[K, V],size),
    },nil
}

// Creates a new synced map initialized with enough memory to hold size 
// elements. Size must be >= 0, an error will be returned if it is not. If size 
// is 0 the map will be initialized with 0 elements. The underlying RWMutex 
// value will be fully unlocked upon initialization.
func NewSyncedHashMap[
    K any, 
    V any, 
    KI widgets.WidgetInterface[K],
    VI widgets.WidgetInterface[V],
](size int) (SyncedHashMap[K,V,KI,VI],error) {
    rv,err:=NewHashMap[K,V,KI,VI](size)
    return SyncedHashMap[K, V, KI, VI]{
        RWMutex: &sync.RWMutex{},
        HashMap: rv,
    },err
}

// Converts the supplied map to a syncronized map. Beware: The original 
// non-synced map will remain useable.
func (m *HashMap[K, V, KI, VI])ToSynced() SyncedHashMap[K,V,KI,VI] {
    return SyncedHashMap[K, V, KI, VI]{
        RWMutex: &sync.RWMutex{},
        HashMap: *m,
    }
}

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (m *HashMap[K, V, KI, VI])Lock() { }

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (m *HashMap[K, V, KI, VI])Unlock() { }

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (m *HashMap[K, V, KI, VI])RLock() { }

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (m *HashMap[K, V, KI, VI])RUnlock() { }

// The SyncedHashMap method to override the HashMap pass through function and 
// actually apply the mutex operation.
func (m *SyncedHashMap[K, V, KI, VI])Lock() { m.RWMutex.Lock() }

// The SyncedHashMap method to override the HashMap pass through function and 
// actually apply the mutex operation.
func (m *SyncedHashMap[K, V, KI, VI])Unlock() { m.RWMutex.Unlock() }

// The SyncedHashMap method to override the HashMap pass through function and 
// actually apply the mutex operation.
func (m *SyncedHashMap[K, V, KI, VI])RLock() { m.RWMutex.RLock() }

// The SyncedHashMap method to override the HashMap pass through function and 
// actually apply the mutex operation.
func (m *SyncedHashMap[K, V, KI, VI])RUnlock() { m.RWMutex.RUnlock() }

// Returns false, maps are addressable.
func (m *HashMap[K, V, KI, VI])IsAddressable() bool { return false }

// Returns false, a map is not synced.
func (m *HashMap[K, V, KI, VI])IsSynced() bool { return false }

// Returns true, a synced map is synced.
func (m *SyncedHashMap[K, V, KI, VI])IsSynced() bool { return true }

// Description: Returns the number of elements in the hash map.
//
// Time Complexity: O(1)
func (m *HashMap[K, V, KI, VI])Length() int {
    return len(m.internalHashMapImpl)
}
// Description: Places a read lock on the underlying hash map and then calls the 
// underlying hash map [HashMap.Length] method.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (m *SyncedHashMap[K, V, KI, VI])Length() int {
    m.RLock()
    defer m.RUnlock()
    return m.HashMap.Length()
}

func (m *HashMap[K, V, KI, VI])Contains(v V) bool {
    return false
}
func (m *SyncedHashMap[K, V, KI, VI])Contains(v V) bool {
    return false
}

func (m *HashMap[K, V, KI, VI])ContainsPntr(v *V) bool {
    return false
}
func (m *SyncedHashMap[K, V, KI, VI])ContainsPntr(v *V) bool {
    return false
}

func (m *HashMap[K, V, KI, VI])Get(k K) (V,error) {
    var tmp V
    return tmp,nil
}
func (m *SyncedHashMap[K, V, KI, VI])Get(k K) (V,error) {
    var tmp V
    return tmp,nil
}

func (m *HashMap[K, V, KI, VI])GetPntr(k K) (*V,error) {
    return nil,nil
}
func (m *SyncedHashMap[K, V, KI, VI])GetPntr(k K) (*V,error) {
    return nil,nil
}

func (m *HashMap[K, V, KI, VI])KeyOf(v V) (K,bool) {
    var tmp K
    return tmp,false
}
func (m *SyncedHashMap[K, V, KI, VI])KeyOf(v V) (K,bool) {
    var tmp K
    return tmp,false
}

func (m *HashMap[K, V, KI, VI])Set(kvPairs ...basic.Pair[K,V]) error {
    return nil
}
func (m *SyncedHashMap[K, V, KI, VI])Set(kvPairs ...basic.Pair[K,V]) error {
    return nil
}

func (m *HashMap[K, V, KI, VI])Emplace(vals ...basic.Pair[K,V]) error {
    return nil
}
func (m *SyncedHashMap[K, V, KI, VI])Emplace(vals ...basic.Pair[K,V]) error {
    return nil
}

func (m *HashMap[K, V, KI, VI])Pop(v V, num int) int {
    return 0
}
func (m *SyncedHashMap[K, V, KI, VI])Pop(v V, num int) int {
    return 0
}

func (m *HashMap[K, V, KI, VI])Delete(k K) error {
    return nil
}
func (m *SyncedHashMap[K, V, KI, VI])Delete(k K) error {
    return nil
}

func (m *HashMap[K, V, KI, VI])Clear() {

}
func (m *SyncedHashMap[K, V, KI, VI])Clear() {

}

func (m *HashMap[K, V, KI, VI])Keys() iter.Iter[K] {
    return iter.NoElem[K]()
}
func (m *SyncedHashMap[K, V, KI, VI])Keys() iter.Iter[K] {
    return iter.NoElem[K]()
}

func (m *HashMap[K, V, KI, VI])Vals() iter.Iter[V] {
    return iter.NoElem[V]()
}
func (m *SyncedHashMap[K, V, KI, VI])Vals() iter.Iter[V] {
    return iter.NoElem[V]()
}

func (m *HashMap[K, V, KI, VI])ValPntrs() iter.Iter[*V] {
    return iter.NoElem[*V]()
}
func (m *SyncedHashMap[K, V, KI, VI])ValPntrs() iter.Iter[*V] {
    return iter.NoElem[*V]()
}

func (m *HashMap[K, V, KI, VI])KeyedEq(
    other containerTypes.KeyedComparisonsOtherConstraint[K,V],
) bool {
    return false
}
func (m *SyncedHashMap[K, V, KI, VI])KeyedEq(
    other containerTypes.KeyedComparisonsOtherConstraint[K,V],
) bool {
    return false
}

func (m *HashMap[K, V, KI, VI])UnorderedEq(
    other containerTypes.ComparisonsOtherConstraint[V],
) bool {
    return false
}
func (m *SyncedHashMap[K, V, KI, VI])UnorderedEq(
    other containerTypes.ComparisonsOtherConstraint[V],
) bool {
    return false
}

func (m *HashMap[K, V, KI, VI])Union(
    l containerTypes.ComparisonsOtherConstraint[V], 
    r containerTypes.ComparisonsOtherConstraint[V],
) {

}
func (m *SyncedHashMap[K, V, KI, VI])Union(
    l containerTypes.ComparisonsOtherConstraint[V], 
    r containerTypes.ComparisonsOtherConstraint[V],
) {

}

func (m *HashMap[K, V, KI, VI])Intersection(
    l containerTypes.ComparisonsOtherConstraint[V], 
    r containerTypes.ComparisonsOtherConstraint[V],
) {

}
func (m *SyncedHashMap[K, V, KI, VI])Intersection(
    l containerTypes.ComparisonsOtherConstraint[V], 
    r containerTypes.ComparisonsOtherConstraint[V],
) {

}

func (m *HashMap[K, V, KI, VI])Difference(
    l containerTypes.ComparisonsOtherConstraint[V], 
    r containerTypes.ComparisonsOtherConstraint[V],
) {

}
func (m *SyncedHashMap[K, V, KI, VI])Differnce(
    l containerTypes.ComparisonsOtherConstraint[V], 
    r containerTypes.ComparisonsOtherConstraint[V],
) {

}

func (m *HashMap[K, V, KI, VI])IsSuperset(
    other containerTypes.ComparisonsOtherConstraint[V],
) bool {
    return false
}
func (m *SyncedHashMap[K, V, KI, VI])IsSuperset(
    other containerTypes.ComparisonsOtherConstraint[V],
) bool {
    return false
}

func (m *HashMap[K, V, KI, VI])IsSubset(
    other containerTypes.ComparisonsOtherConstraint[V],
) bool {
    return false
}
func (m *SyncedHashMap[K, V, KI, VI])IsSubset(
    other containerTypes.ComparisonsOtherConstraint[V],
) bool {
    return false
}
