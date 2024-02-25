package containers

import (
	"sync"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containerTypes"
)

type (
    internalMapImpl[K any, V any] map[uint64]basic.Pair[K,V]

    // A type to represent a map that dynamically grows as key value pairs are 
    // added. The set will maintain uniqueness and is internally implemented 
    // with a hashing method. The type constraints on the generics define the 
    // logic for how value specific operations, such as equality comparisons, 
    // will be handled.
    Map[
        K any,
        V any,
        KI widgets.WidgetInterface[K],
        VI widgets.WidgetInterface[V],
    ] struct {
        internalMapImpl[K,V]
    }

    // A synchronized version of Map. All operations will be wrapped in the
    // appropriate calls the embedded RWMutex. A pointer to a RWMutex is 
    // embedded rather than a value to avoid copying the lock value.
    SyncedMap[
        K any,
        V any,
        KI widgets.WidgetInterface[K],
        VI widgets.WidgetInterface[V],
    ] struct {
        *sync.RWMutex
        Map[K,V,KI,VI]
    }
)

// Creates a new map initialized with enough memory to hold size elements. 
// Size must be >= 0, an error will be returned if it is not. If size is 0 the 
// map will be initialized with 0 elements.
func NewMap[
    K any, 
    V any, 
    KI widgets.WidgetInterface[K], 
    VI widgets.WidgetInterface[V],
](size int) (Map[K,V,KI,VI],error) {
    if size<0 {
        return Map[K, V, KI, VI]{},getSizeError(size)
    }
    return Map[K, V, KI, VI]{
        internalMapImpl: make(internalMapImpl[K, V],size),
    },nil
}

// Creates a new synced map initialized with enough memory to hold size 
// elements. Size must be >= 0, an error will be returned if it is not. If size 
// is 0 the map will be initialized with 0 elements. The underlying RWMutex 
// value will be fully unlocked upon initialization.
func NewSyncedMap[
    K any, 
    V any, 
    KI widgets.WidgetInterface[K], 
    VI widgets.WidgetInterface[V],
](size int) (SyncedMap[K,V,KI,VI],error) {
    rv,err:=NewMap[K,V,KI,VI](size)
    return SyncedMap[K, V, KI, VI]{
        RWMutex: &sync.RWMutex{},
        Map: rv,
    },err
}

// Converts the supplied map to a syncronized map. Beware: The original 
// non-synced map will remain useable.
func (m *Map[K, V, KI, VI])ToSynced() SyncedMap[K,V,KI,VI] {
    return SyncedMap[K, V, KI, VI]{
        RWMutex: &sync.RWMutex{},
        Map: *m,
    }
}

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (m *Map[K, V, KI, VI])Lock() { }

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (m *Map[K, V, KI, VI])Unlock() { }

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (m *Map[K, V, KI, VI])RLock() { }

// A empty pass through function that performs no action. Needed for the
// [containerTypes.Comparisons] interface.
func (m *Map[K, V, KI, VI])RUnlock() { }

// The SyncedMap method to override the Map pass through function and 
// actually apply the mutex operation.
func (m *SyncedMap[K, V, KI, VI])Lock() { m.RWMutex.Lock() }

// The SyncedMap method to override the Map pass through function and 
// actually apply the mutex operation.
func (m *SyncedMap[K, V, KI, VI])Unlock() { m.RWMutex.Unlock() }

// The SyncedMap method to override the Map pass through function and 
// actually apply the mutex operation.
func (m *SyncedMap[K, V, KI, VI])RLock() { m.RWMutex.RLock() }

// The SyncedMap method to override the Map pass through function and 
// actually apply the mutex operation.
func (m *SyncedMap[K, V, KI, VI])RUnlock() { m.RWMutex.RUnlock() }

// Returns false, maps are addressable.
func (m *Map[K, V, KI, VI])IsAddressable() bool { return false }

// Returns false, a map is not synced.
func (m *Map[K, V, KI, VI])IsSynced() bool { return false }

// Returns true, a synced map is synced.
func (m *SyncedMap[K, V, KI, VI])IsSynced() bool { return true }

func (m *Map[K, V, KI, VI])Length() int {
    return 0
}
func (m *SyncedMap[K, V, KI, VI])Length() int {
    return 0
}

func (m *Map[K, V, KI, VI])Contains(v V) bool {
    return false
}
func (m *SyncedMap[K, V, KI, VI])Contains(v V) bool {
    return false
}

func (m *Map[K, V, KI, VI])ContainsPntr(v *V) bool {
    return false
}
func (m *SyncedMap[K, V, KI, VI])ContainsPntr(v *V) bool {
    return false
}

func (m *Map[K, V, KI, VI])Get(k K) (V,error) {
    var tmp V
    return tmp,nil
}
func (m *SyncedMap[K, V, KI, VI])Get(k K) (V,error) {
    var tmp V
    return tmp,nil
}

func (m *Map[K, V, KI, VI])GetPntr(k K) (*V,error) {
    return nil,nil
}
func (m *SyncedMap[K, V, KI, VI])GetPntr(k K) (*V,error) {
    return nil,nil
}

func (m *Map[K, V, KI, VI])KeyOf(v V) (K,bool) {
    var tmp K
    return tmp,false
}
func (m *SyncedMap[K, V, KI, VI])KeyOf(v V) (K,bool) {
    var tmp K
    return tmp,false
}

func (m *Map[K, V, KI, VI])Set(kvPairs ...basic.Pair[K,V]) error {
    return nil
}
func (m *SyncedMap[K, V, KI, VI])Set(kvPairs ...basic.Pair[K,V]) error {
    return nil
}

func (m *Map[K, V, KI, VI])SetSequential(k K, v ...V) error {
    return nil
}
func (m *SyncedMap[K, V, KI, VI])SetSequential(k K, v ...V) error {
    return nil
}

// Standard push func you would think of
func (m *Map[K, V, KI, VI])Emplace(vals ...basic.Pair[K,V]) error {
    return nil
}
func (m *SyncedMap[K, V, KI, VI])Emplace(vals ...basic.Pair[K,V]) error {
    return nil
}

func (m *Map[K, V, KI, VI])EmplaceSequential(k K, vals ...V) error {
    return nil
}
func (m *SyncedMap[K, V, KI, VI])EmplaceSequential(k K, vals ...V) error {
    return nil
}

// Pushes all keys to the right by "1"
func (m *Map[K, V, KI, VI])Insert(vals ...basic.Pair[K,V]) error {
    return nil
}
func (m *SyncedMap[K, V, KI, VI])Insert(vals ...basic.Pair[K,V]) error {
    return nil
}

func (m *Map[K, V, KI, VI])InsertSequential(k K, vals ...V) error {
    return nil
}
func (m *SyncedMap[K, V, KI, VI])InsertSequential(k K, vals ...V) error {
    return nil
}

func (m *Map[K, V, KI, VI])Pop(v V, num int) int {
    return 0
}
func (m *SyncedMap[K, V, KI, VI])Pop(v V, num int) int {
    return 0
}

func (m *Map[K, V, KI, VI])Delete(k K) error {
    return nil
}
func (m *SyncedMap[K, V, KI, VI])Delete(k K) error {
    return nil
}

func (m *Map[K, V, KI, VI])Clear() {

}
func (m *SyncedMap[K, V, KI, VI])Clear() {

}

func (m *Map[K, V, KI, VI])Keys() iter.Iter[K] {
    return iter.NoElem[K]()
}
func (m *SyncedMap[K, V, KI, VI])Keys() iter.Iter[K] {
    return iter.NoElem[K]()
}

func (m *Map[K, V, KI, VI])Vals() iter.Iter[V] {
    return iter.NoElem[V]()
}
func (m *SyncedMap[K, V, KI, VI])Vals() iter.Iter[V] {
    return iter.NoElem[V]()
}

func (m *Map[K, V, KI, VI])ValPntrs() iter.Iter[*V] {
    return iter.NoElem[*V]()
}
func (m *SyncedMap[K, V, KI, VI])ValPntrs() iter.Iter[*V] {
    return iter.NoElem[*V]()
}

func (m *Map[K, V, KI, VI])KeyedEq(
    other containerTypes.KeyedComparisonsOtherConstraint[K,V],
) bool {
    return false
}
func (m *SyncedMap[K, V, KI, VI])KeyedEq(
    other containerTypes.KeyedComparisonsOtherConstraint[K,V],
) bool {
    return false
}

func (m *Map[K, V, KI, VI])UnorderedEq(
    other containerTypes.ComparisonsOtherConstraint[V],
) bool {
    return false
}
func (m *SyncedMap[K, V, KI, VI])UnorderedEq(
    other containerTypes.ComparisonsOtherConstraint[V],
) bool {
    return false
}

func (m *Map[K, V, KI, VI])Union(
    l containerTypes.ComparisonsOtherConstraint[V], 
    r containerTypes.ComparisonsOtherConstraint[V],
) {

}
func (m *SyncedMap[K, V, KI, VI])Union(
    l containerTypes.ComparisonsOtherConstraint[V], 
    r containerTypes.ComparisonsOtherConstraint[V],
) {

}

func (m *Map[K, V, KI, VI])Intersection(
    l containerTypes.ComparisonsOtherConstraint[V], 
    r containerTypes.ComparisonsOtherConstraint[V],
) {

}
func (m *SyncedMap[K, V, KI, VI])Intersection(
    l containerTypes.ComparisonsOtherConstraint[V], 
    r containerTypes.ComparisonsOtherConstraint[V],
) {

}

func (m *Map[K, V, KI, VI])Difference(
    l containerTypes.ComparisonsOtherConstraint[V], 
    r containerTypes.ComparisonsOtherConstraint[V],
) {

}
func (m *SyncedMap[K, V, KI, VI])Differnce(
    l containerTypes.ComparisonsOtherConstraint[V], 
    r containerTypes.ComparisonsOtherConstraint[V],
) {

}

func (m *Map[K, V, KI, VI])IsSuperset(
    other containerTypes.ComparisonsOtherConstraint[V],
) bool {
    return false
}
func (m *SyncedMap[K, V, KI, VI])IsSuperset(
    other containerTypes.ComparisonsOtherConstraint[V],
) bool {
    return false
}

func (m *Map[K, V, KI, VI])IsSubset(
    other containerTypes.ComparisonsOtherConstraint[V],
) bool {
    return false
}
func (m *SyncedMap[K, V, KI, VI])IsSubset(
    other containerTypes.ComparisonsOtherConstraint[V],
) bool {
    return false
}
