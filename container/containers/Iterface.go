package containers

// import (
// 	"github.com/barbell-math/util/algo/iter"
// 	"github.com/barbell-math/util/container/staticContainers"
// )

// TODO - impl and test
// // Window will take the parent iterator and return a window of it's cached
// // values of length equal to the allowed capacity of the supplied queue (q).
// // Note that a static queue is expected to be passed instead of a dynamic one.
// // If allowPartials is true then windows that are not full will be returned.
// // Setting allowPartials to false will enforce all returned windows to have
// // length equal to the allowed capacity of the supplied queue. An error will
// // stop iteration.
// func Window[T any](
//     i iter.Iter[T],
//     q interface{ staticContainers.Queue[T]; staticContainers.Vector[T] },
//     allowPartials bool,
// ) iter.Iter[staticContainers.Vector[T]] {
//     return iter.Next(
//         i,
//         func(
//             index int, val T, status iter.IteratorFeedback,
//         ) (iter.IteratorFeedback, staticContainers.Vector[T], error) {
//             if status==Break {
//                 return Break,q,nil;
//             }
//             q.ForcePushBack(val);
//             if !allowPartials && q.Length()!=q.Capacity() {
//                 return iter.Iterate,q,nil;
//             }
//             return iter.Continue,q,nil;
//         });
// }
