package containers

import "github.com/barbell-math/util/algo/widgets"

//go:generate go run interfaceTest.go -type=HashGraph -category=dynamic -interface=DirectedGraph -genericDecl=[int,int] -factory=generateHashGraph
//go:generate go run interfaceTest.go -type=SyncedHashGraph -category=dynamic -interface=DirectedGraph -genericDecl=[int,int] -factory=generateSyncedHashGraph
// //go:generate go run interfaceTest.go -type=HashGraph -category=dynamic -interface=UndirectedGraph -genericDecl=[int,int] -factory=generateHashGraph
// //go:generate go run interfaceTest.go -type=SyncedHashGraph -category=dynamic -interface=UndirectedGraph -genericDecl=[int,int] -factory=generateSyncedHashGraph

func generateHashGraph(
	capacity int,
) HashGraph[int, int, widgets.BuiltinInt, widgets.BuiltinInt] {
	v, _ := NewHashGraph[int, int, widgets.BuiltinInt, widgets.BuiltinInt](
		capacity,
		capacity,
	)
	return v
}

func generateSyncedHashGraph(
	capacity int,
) SyncedHashGraph[int, int, widgets.BuiltinInt, widgets.BuiltinInt] {
	v, _ := NewSyncedHashGraph[int, int, widgets.BuiltinInt, widgets.BuiltinInt](
		capacity,
		capacity,
	)
	return v
}
