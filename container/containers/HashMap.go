package containers

import (
	"fmt"
	"sync"

	"github.com/barbell-math/util/algo/hash"
	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containerTypes"
)

type (
    internalHashMapImpl[K any, V any] map[hash.Hash]basic.Pair[K,V]

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

// Description: Contains will return true if the supplied value is in the 
// map, false otherwise. All equality comparisons are performed by the 
// generic VI widget type that the vector was initialized with. 
//
// Time Complexity: O(n) (linear search)
func (m *HashMap[K, V, KI, VI])Contains(v V) bool {
    return m.ContainsPntr(&v)
}
// Description: Places a read lock on the underlying map and then calls the 
// underlying map [Map.ContainsPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (m *SyncedHashMap[K, V, KI, VI])Contains(v V) bool {
    m.RLock()
    defer m.RUnlock()
    return m.ContainsPntr(&v)
}

// Description: ContainsPntr will return true if the supplied value is in the 
// map, false otherwise. All equality comparisons are performed by the 
// generic VI widget type that the map was initialized with.
//
// Time Complexity: O(n) (linear search)
func (m *HashMap[K, V, KI, VI])ContainsPntr(v *V) bool {
    w:=widgets.NewWidget[V,VI]()
    for _,iterV:=range(m.internalHashMapImpl) {
        if w.Eq(v,&iterV.B) {
            return true
        }
    }
    return false
}
// Description: Places a read lock on the underlying map and then calls the 
// underlying maps [Map.ContainsPntr] method.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (m *SyncedHashMap[K, V, KI, VI])ContainsPntr(v *V) bool {
    m.RLock()
    defer m.RUnlock()
    return m.HashMap.ContainsPntr(v)
}

// Description: Gets the value at the specified key. Returns a 
// [containerTypes.KeyError] if the key is not found in the hash map.
//
// Time Complexity: O(1)
func (m *HashMap[K, V, KI, VI])Get(k K) (V,error) {
    if h,ok:=m.getHashPosition(&k); ok {
        return m.internalHashMapImpl[h].B,nil
    }
    var tmp V
    return tmp,getKeyError[K](&k)
}
// Description: Places a read lock on the underlying hash map and then gets the 
// value at the specified key. Exhibits the same behavior as the [HashMap.Get]
// method. The underlying [HashMap.Get] method is not called to avoid copying 
// the return value twice, which could be inefficient with a large value for the 
// K generic.
//
// Lock Type: Read
//
// Time Complexity: O(1)
func (m *SyncedHashMap[K, V, KI, VI])Get(k K) (V,error) {
    m.RLock()
    defer m.RUnlock()
    if h,ok:=m.getHashPosition(&k); ok {
        return m.internalHashMapImpl[h].B,nil
    }
    var tmp V
    return tmp,getKeyError[K](&k)
}

func (h *HashMap[K, V, KI, VI])getHashPosition(k *K) (hash.Hash,bool) {
    w:=widgets.NewWidget[K,KI]()
    for i:=w.Hash(k); ; i++ {
        fmt.Println("Considering hash: ",i)
        if iterV,found:=h.internalHashMapImpl[i]; found && w.Eq(k,&iterV.A) {
            fmt.Println("Found acceptable value.")
            return i,true 
        } else if !found {
            fmt.Println("No value found.")
            return hash.Hash(0),false
        }
    }
}

// Panics, hash maps are not addressable.
func (m *HashMap[K, V, KI, VI])GetPntr(k K) (*V,error) {
    panic(getNonAddressablePanicText("hash map"))
}

// Description: KeyOf will return the key of the first occurrence of the 
// supplied value in the map. If the value is not found then the returned 
// key will be a zero initialized key value and the boolean flag will be set to 
// false. If the value is found then the boolean flag will be set to true. All 
// equality comparisons are performed by the generic VI widget type that the map 
// was initialized with.
//
// Time Complexity: O(n) (linear search)
func (m *HashMap[K, V, KI, VI])KeyOf(v V) (K,bool) {
    var tmp K
    return tmp,m.keyOfImpl(&tmp,&v)
}
// Description: Places a read lock on the underlying vector and then calls the 
// underlying vectors [Map.KeyOf] implemenation method. The [Map.KeyOf] method 
// is not called directly to avoid copying the val variable twice, which could 
// be expensive with a large type for the V generic.
//
// Lock Type: Read
//
// Time Complexity: O(n) (linear search)
func (m *SyncedHashMap[K, V, KI, VI])KeyOf(v V) (K,bool) {
    m.RLock()
    defer m.RUnlock()
    var tmp K
    return tmp,m.keyOfImpl(&tmp,&v)
}

func (m *HashMap[K, V, KI, VI])keyOfImpl(k *K, v *V) bool {
    w:=widgets.NewWidget[V,VI]()
    for _,iterV:=range(m.internalHashMapImpl) {
        if w.Eq(v,&iterV.B) {
            *k=iterV.A
            return true
        }
    }
    return false
}

// Description: Sets the values at the specified keys. Returns an error if the 
// key is not in the map. Stops setting values as soon as an error is 
// encountered.
//
// Time Complexity: O(m), where m=len(vals)
func (m *HashMap[K, V, KI, VI])Set(kvPairs ...basic.Pair[K,V]) error {
    return m.setImpl(kvPairs)
}
// Description: Places a write lock on the underlying map and then calls the 
// underlying vectors [Map.Set] implementaiton method. The [Map.Set] method is 
// not called directly to avoid copying the vals varargs twice, which could be 
// expensive with a large types for the K or V generics or a large number of 
// values.
//
// Lock Type: Write
//
// Time Complexity: O(m), where m=len(vals)
func (m *SyncedHashMap[K, V, KI, VI])Set(kvPairs ...basic.Pair[K,V]) error {
    m.Lock()
    defer m.Unlock()
    return m.setImpl(kvPairs)
}

func (m *HashMap[K, V, KI, VI])setImpl(kvPairs []basic.Pair[K,V]) error {
    for i:=0; i<len(kvPairs); i++ {
        if h,found:=m.getHashPosition(&kvPairs[i].A); found {
            m.internalHashMapImpl[h]=kvPairs[i]
        } else {
            return getKeyError[K](&kvPairs[i].A)
        }
    }
    return nil
}

// Description: Emplace will insert the supplied values into the hash map if 
// they do not exist and will set they keys value if it already exists in the
// hash map. The values will be inserted in the order that they are given. 
//
// Time Complexity: O(m), where m=len(vals)
func (m *HashMap[K, V, KI, VI])Emplace(vals ...basic.Pair[K,V]) error {
    for i:=0; i<len(vals); i++ {
        m.emplaceImpl(&vals[i])
    }
    return nil
}
// Description: Places a write lock on the underlying hash map and then calls 
// the underlying hash maps [HashMap.Insert] implementation method. The 
// [HashMap.Insert] method is not called directly to avoid copying the vals 
// varargs twice, which could be expensive with a large types for the K or V 
// generics or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: O(n*m), where m=len(vals)
func (m *SyncedHashMap[K, V, KI, VI])Emplace(vals ...basic.Pair[K,V]) error {
    m.Lock()
    defer m.Unlock()
    for i:=0; i<len(vals); i++ {
        m.emplaceImpl(&vals[i])
    }
    return nil
}

func (m *HashMap[K, V, KI, VI])emplaceImpl(v *basic.Pair[K,V]) {
    w:=widgets.NewWidget[K,KI]()
    for i:=w.Hash(&v.A); ; i++ {
        if iterV,found:=m.internalHashMapImpl[i]; !found {
            m.internalHashMapImpl[i]=*v
            break
        } else if w.Eq(&v.A,&iterV.A) {
            m.internalHashMapImpl[i]=*v
            break
        }
    }
}

// Description: Pop will remove the first num occurrences of val in the map. 
// All equality comparisons are performed by the generic VI widget type that the 
// map was initialized with. If num is <=0 then no values will be poped and 
// the map will not change.
//
// Time Complexity: O(n)
func (m *HashMap[K, V, KI, VI])Pop(v V) int {
    return m.popImpl(&v)
}
// Description: Places a write lock on the underlying map and then calls the 
// underlying map [Map.Pop] implementation method. The [Map.Pop] method is not 
// called directly to avoid copying the value twice, which could be expensive 
// with a large type for the V generic or a large number of values.
//
// Lock Type: Write
//
// Time Complexity: O(n)
func (m *SyncedHashMap[K, V, KI, VI])Pop(v V) int {
    m.Lock()
    defer m.Unlock()
    return m.popImpl(&v)
}

func (m *HashMap[K, V, KI, VI])popImpl(v *V) int {
    rv:=0
    vw:=widgets.NewWidget[V,VI]()
    for iterH,iterV:=range(m.internalHashMapImpl) {
        if vw.Eq(&iterV.B,v) {
            m.removeValue(iterH,v)
            rv++
        }
    }
    return rv
}

func (m *HashMap[K, V, KI, VI])removeValue(h hash.Hash, v *V) {
    kw:=widgets.NewWidget[K,KI]()
    vw:=widgets.NewWidget[V,VI]()
    delete(m.internalHashMapImpl,h)
    curPos:=h
    for j:=h+1; ; j++ {
        fmt.Println("Looking at hash: ",j)
        if iterV,found:=m.internalHashMapImpl[j]; found {
            // If the hash of iterV is > the curPos then do not move it, it 
            // is already in the correct position, but continue on this 
            // collision chain to check for other values that could move.
            if kw.Hash(&iterV.A)>curPos {
                continue
            }
            // If the value is eq to the value we are deleting then also skip it
            // because deleting it and then re-creating it does not guarintee
            // it will be iterated on later.
            // Scenario 1: The value is the last value of the collision chain
            //   The value will be orphaned an unaccessable using the standard
            //   Get method but will be deleted by a future iteration of the
            //   popImpl function. If it were moved this 're-teration' would not
            //   be guarinteed.
            // Scenario 2: The value is in the middle of a collision chain. in
            //   this case the hole that is created will be filled by an 
            //   appripriate value later in the collision chain.
            if vw.Eq(v,&iterV.B) {
                continue
            }
            m.internalHashMapImpl[curPos]=m.internalHashMapImpl[j]
            delete(m.internalHashMapImpl,j)
            curPos=j 
        } else {
            break
        }
    }
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
