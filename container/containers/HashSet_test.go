package containers

import "github.com/barbell-math/util/algo/widgets"

//go:generate go run interfaceTest.go -type=HashSet -category=dynamic -interface=Set -genericDecl=[int] -factory=generateHashSet
//go:generate go run interfaceTest.go -type=SyncedHashSet -category=dynamic -interface=Set -genericDecl=[int] -factory=generateSyncedHashSet

func generateHashSet() HashSet[int,widgets.BuiltinInt] {
    v,_:=NewHashSet[int,widgets.BuiltinInt](0)
    return v
}

func generateSyncedHashSet() SyncedHashSet[int,widgets.BuiltinInt] {
    v,_:=NewSyncedHashSet[int,widgets.BuiltinInt](0)
    return v
}
