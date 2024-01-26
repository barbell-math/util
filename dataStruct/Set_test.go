package dataStruct

import (
	// "testing"

	// "github.com/barbell-math/util/dataStruct/types/static"
	// // customerr "github.com/barbell-math/util/err"
	// "github.com/barbell-math/util/test"
)

type hashableStr[H ~uint32 | ~uint64] string
func (h *hashableStr[H])Hash() H {
    return H(1)
}
func (h *hashableStr[H])Eq(other *hashableStr[H]) bool {
    return *h==*other
}
func (h *hashableStr[H])Neq(other *hashableStr[H]) bool {
    return !h.Eq(other)
}

// func TestSetWriteInterface(t *testing.T){
//     s,_:=NewSet[uint32,hashableStr[uint32]](0)
//     s2,_:=NewSyncedSet[uint32,hashableStr[uint32]](5)
//     writeInterfaceTypeCeck[uint32,hashableStr[uint32]](&s);
//     writeInterfaceTypeCeck[uint32,hashableStr[uint32]](&s2);
// }
// 
// func TestSetReadInterface(t *testing.T){
//     s,_:=NewSet[uint32,string](0)
//     s2,_:=NewSyncedSet[uint32,string](5)
//     readInterfaceTypeCeck[int,string](&s);
//     readInterfaceTypeCeck[int,string](&s2);
// }
// 
// func TestSetDynSetTypeInterface(t *testing.T) {
//     s,_:=NewSet[uint32,string](0)
//     s2,_:=NewSyncedSet[uint32,string](5)
//     dynSetInterfaceTypeCheck[int](&s);
//     dynSetInterfaceTypeCheck[int](&s2);
// }
// 
// func TestSetDynStackTypeInterface(t *testing.T) {
//     s,_:=NewSet[uint32,string](0)
//     s2,_:=NewSyncedSet[uint32,string](5)
//     dynStackInterfaceTypeCheck[int](&s);
//     dynStackInterfaceTypeCheck[int](&s2);
// }
// 
// func TestSetDynQueueTypeInterface(t *testing.T) {
//     s,_:=NewSet[uint32,string](0)
//     s2,_:=NewSyncedSet[uint32,string](5)
//     dynQueueInterfaceTypeCheck[int](&s);
//     dynQueueInterfaceTypeCheck[int](&s2);
// }
// 
// func TestSetDynDequeTypeInterface(t *testing.T) {
//     s,_:=NewSet[uint32,string](0)
//     s2,_:=NewSyncedSet[uint32,string](5)
//     dynDequeInterfaceTypeCheck[int](&s);
//     dynDequeInterfaceTypeCheck[int](&s2);
// }
// 
// func TestSetEqualsTypeInterface(t *testing.T) {
//     s,_:=NewSet[uint32,string](0)
//     s2,_:=NewSyncedSet[uint32,string](5)
//     equalsInterfaceTypeCheck[Set[int],int](&s);
//     equalsInterfaceTypeCheck[Set[int],int](&s2);
// }
// 
// func TestSetStaticTypeInterface(t *testing.T){
//     test.Panics(
//         func () {
//             var c any
//             c,_=NewSet[int](5)
//             c2:=c.(static.Set[int])
//             _=c2
//         }, 
//         "Code did not panic when casting a dynamic vector to a static vector.",t,
//     )
//     test.Panics(
//         func () {
//             var c any
//             c,_=NewSet[int](5)
//             c2:=c.(static.Queue[int])
//             _=c2
//         }, 
//         "Code did not panic when casting a dynamic vector to a static queue.",t,
//     )
//     test.Panics(
//         func () {
//             var c any
//             c,_=NewSet[int](5)
//             c2:=c.(static.Stack[int])
//             _=c2
//         }, 
//         "Code did not panic when casting a dynamic vector to a static stack.",t,
//     )
//     test.Panics(
//         func () {
//             var c any
//             c,_=NewSet[int](5)
//             c2:=c.(static.Deque[int])
//             _=c2
//         }, 
//         "Code did not panic when casting a dynamic vector to a static deque.",t,
//     )
// }
