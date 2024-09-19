package lexer

import (
	"fmt"
	"testing"

	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containers"
	"github.com/barbell-math/util/container/dynamicContainers"
	"github.com/barbell-math/util/test"
	"github.com/barbell-math/util/widgets"
)

func verifyGraph[A any, AI widgets.WidgetInterface[A]](
	nfaGraph dynamicContainers.ReadDirectedGraph[nfaNode, containers.HashSet[A, AI]],
	expectedLinks []basic.Triple[nfaNode, containers.HashSet[A, AI], nfaNode],
	t *testing.T,
) {
	test.Eq(len(expectedLinks), nfaGraph.NumLinks(), t)
	for _, v := range expectedLinks {
		test.True(nfaGraph.ContainsLinkPntr(&v.A, &v.C, &v.B), t)

		fullV := nfaNode{id: v.A.id}
		test.Nil(nfaGraph.GetVertex(&fullV), t)
		test.Eq(v.A.id, fullV.id, t)
		test.Eq(v.A.flags, fullV.flags, t)

		fullV = nfaNode{id: v.C.id}
		test.Nil(nfaGraph.GetVertex(&fullV), t)
		test.Eq(v.C.id, fullV.id, t)
		test.Eq(v.C.flags, fullV.flags, t)
	}
}

func TestNFAAppendTransition(t *testing.T) {
	n := NewNFA[byte, widgets.BuiltinByte]()
	n.AppendTransition('a')
	verifyGraph[byte, widgets.BuiltinByte](n,
		[]basic.Triple[nfaNode, containers.HashSet[byte, widgets.BuiltinByte], nfaNode]{
			{
				A: nfaNode{id: 0, flags: nfaSource},
				B: containers.HashSetValInit[byte, widgets.BuiltinByte]('a'),
				C: nfaNode{id: 1, flags: nfaSink},
			},
		},
		t,
	)

	n.AppendTransition('b')
	verifyGraph[byte, widgets.BuiltinByte](n,
		[]basic.Triple[nfaNode, containers.HashSet[byte, widgets.BuiltinByte], nfaNode]{
			{
				A: nfaNode{id: 0, flags: nfaSource},
				B: containers.HashSetValInit[byte, widgets.BuiltinByte]('a'),
				C: nfaNode{id: 1, flags: 0},
			},
			{
				A: nfaNode{id: 1, flags: 0},
				B: containers.HashSetValInit[byte, widgets.BuiltinByte]('b'),
				C: nfaNode{id: 2, flags: nfaSink},
			},
		},
		t,
	)

	n.AppendTransition('c')
	verifyGraph[byte, widgets.BuiltinByte](n,
		[]basic.Triple[nfaNode, containers.HashSet[byte, widgets.BuiltinByte], nfaNode]{
			{
				A: nfaNode{id: 0, flags: nfaSource},
				B: containers.HashSetValInit[byte, widgets.BuiltinByte]('a'),
				C: nfaNode{id: 1, flags: 0},
			},
			{
				A: nfaNode{id: 1, flags: 0},
				B: containers.HashSetValInit[byte, widgets.BuiltinByte]('b'),
				C: nfaNode{id: 2, flags: 0},
			},
			{
				A: nfaNode{id: 2, flags: 0},
				B: containers.HashSetValInit[byte, widgets.BuiltinByte]('c'),
				C: nfaNode{id: 3, flags: nfaSink},
			},
		},
		t,
	)
}

func TestNFAAppendNFAEmpty(t *testing.T) {
	op := func(
		n1 NFA[byte, widgets.BuiltinByte],
		n2 NFA[byte, widgets.BuiltinByte],
	) {
		n1.AppendNFA(n2)
		test.Eq(
			"map[0:{2 [{99 1}]} 1:{0 [{100 2}]} 2:{4 []}]",
			fmt.Sprint(n1),
			t,
		)
	}

	n1 := NewNFA[byte, widgets.BuiltinByte]()
	n2 := NewNFA[byte, widgets.BuiltinByte]()
	n2.AppendTransition('c')
	n2.AppendTransition('d')
	op(n1, n2)

	n1 = NewNFA[byte, widgets.BuiltinByte]()
	n1.AppendTransition('c')
	n1.AppendTransition('d')
	n2 = NewNFA[byte, widgets.BuiltinByte]()
	op(n1, n2)

	n1 = NewNFA[byte, widgets.BuiltinByte]()
	n1.AppendTransition('c')
	n1.AppendTransition('d')
	n2 = NewNFA[byte, widgets.BuiltinByte]()
	op(n1, n2)
}

// func TestNFAAppendNFA(t *testing.T) {
// 	n1 := NewNFA()
// 	n1.AppendTransition('a')
// 	n1.AppendTransition('b')
//
// 	n2 := NewNFA()
// 	n2.AppendTransition('c')
// 	n2.AppendTransition('d')
//
// 	n1.AppendNFA(n2)
// 	test.Eq(
// 		"map["+
// 			"0:{2 [{97 1}]} "+
// 			"1:{0 [{98 2}]} "+
// 			"2:{0 [{99 3}]} "+
// 			"3:{0 [{100 4}]} "+
// 			"4:{4 []}"+
// 			"]",
// 		fmt.Sprint(n1),
// 		t,
// 	)
// }
//
// func TestNFAApplyKleeneEmpty(t *testing.T) {
// 	n1 := NewNFA()
// 	n1.ApplyKleene()
// 	test.Eq(
// 		"map[0:{2 [{128 1}]} 1:{0 [{128 2}]} 2:{4 [{128 1}]}]",
// 		fmt.Sprint(n1),
// 		t,
// 	)
// }
//
// func TestNFAApplyKleene(t *testing.T) {
// 	n1 := NewNFA()
// 	n1.AppendTransition('a')
// 	n1.ApplyKleene()
// 	test.Eq(
// 		"map[0:{2 [{128 2}]} 1:{0 [{128 3}]} 2:{0 [{97 1}]} 3:{4 [{128 2}]}]",
// 		fmt.Sprint(n1),
// 		t,
// 	)
//
// 	n1 = NewNFA()
// 	n1.AppendTransition('a')
// 	n1.AppendTransition('b')
// 	n1.ApplyKleene()
// 	test.Eq(
// 		"map["+
// 			"0:{2 [{128 3}]} "+
// 			"1:{0 [{98 2}]} "+
// 			"2:{0 [{128 4}]} "+
// 			"3:{0 [{97 1}]} "+
// 			"4:{4 [{128 3}]}"+
// 			"]",
// 		fmt.Sprint(n1),
// 		t,
// 	)
// }
//
// func TestNFAApplyKleeneMultipleTimes(t *testing.T) {
// 	n1 := NewNFA()
// 	n1.AppendTransition('a')
// 	n1.ApplyKleene()
// 	n1.ApplyKleene()
// 	test.Eq(
// 		"map[0:{2 [{128 2}]} 1:{0 [{128 3}]} 2:{0 [{97 1}]} 3:{4 [{128 2}]}]",
// 		fmt.Sprint(n1),
// 		t,
// 	)
// }
//
// func TestNFAAddBranchEmpty(t *testing.T) {
// 	op := func(n1 NFA, n2 NFA) {
// 		n1.AddBranch(n2)
// 		test.Eq(
// 			"map[0:{2 [{99 1}]} 1:{0 [{100 2}]} 2:{4 []}]",
// 			fmt.Sprint(n1),
// 			t,
// 		)
// 	}
//
// 	n1 := NFA{}
// 	n2 := NewNFA()
// 	n2.AppendTransition('c')
// 	n2.AppendTransition('d')
// 	op(n1, n2)
//
// 	n1 = NewNFA()
// 	n2 = NewNFA()
// 	n2.AppendTransition('c')
// 	n2.AppendTransition('d')
// 	op(n1, n2)
//
// 	n1 = NewNFA()
// 	n1.AppendTransition('c')
// 	n1.AppendTransition('d')
// 	n2 = NFA{}
// 	op(n1, n2)
//
// 	n1 = NewNFA()
// 	n1.AppendTransition('c')
// 	n1.AppendTransition('d')
// 	n2 = NewNFA()
// 	op(n1, n2)
// }
//
// func TestNFAAddBranch(t *testing.T) {
// 	n1 := NewNFA()
// 	n1.AppendTransition('a')
// 	n1.AppendTransition('b')
//
// 	n2 := NewNFA()
// 	n2.AppendTransition('c')
// 	n2.AppendTransition('d')
//
// 	n1.AddBranch(n2)
// 	test.Eq(
// 		"map["+
// 			"0:{2 [{97 1} {99 3}]} "+
// 			"1:{0 [{98 2}]} "+
// 			"2:{0 [{128 5}]} "+
// 			"3:{0 [{100 4}]} "+
// 			"4:{0 [{128 5}]} "+
// 			"5:{4 []}"+
// 			"]",
// 		fmt.Sprint(n1),
// 		t,
// 	)
// }
//
// func TestNFAKleeneAndAppendNFA(t *testing.T) {
// 	n1 := NewNFA()
// 	n1.AppendTransition('a')
// 	n1.ApplyKleene()
//
// 	n2 := NewNFA()
// 	n2.AppendTransition('b')
// 	n2.AppendTransition('c')
// 	n2.ApplyKleene()
//
// 	n1.AppendNFA(n2)
// 	test.Eq(
// 		"map["+
// 			"0:{2 [{128 2}]} "+
// 			"1:{0 [{128 3}]} "+
// 			"2:{0 [{97 1}]} "+
// 			"3:{0 [{128 2} {128 6}]} "+
// 			"4:{0 [{99 5}]} "+
// 			"5:{0 [{128 7}]} "+
// 			"6:{0 [{98 4}]} "+
// 			"7:{4 [{128 6}]}"+
// 			"]",
// 		fmt.Sprint(n1),
// 		t,
// 	)
// }
//
// func TestNFABranchWithKleenedNFA(t *testing.T) {
// 	n1 := NewNFA()
// 	n1.AppendTransition('a')
// 	n1.ApplyKleene()
//
// 	n2 := NewNFA()
// 	n2.AppendTransition('b')
//
// 	n1.AddBranch(n2)
// 	test.Eq(
// 		"map["+
// 			"0:{2 [{128 2} {98 4}]} "+
// 			"1:{0 [{128 3}]} "+
// 			"2:{0 [{97 1}]} "+
// 			"3:{0 [{128 2} {128 5}]} "+
// 			"4:{0 [{128 5}]} "+
// 			"5:{4 []}"+
// 			"]",
// 		fmt.Sprint(n1),
// 		t,
// 	)
//
// 	n1 = NewNFA()
// 	n1.AppendTransition('a')
// 	n1.ApplyKleene()
//
// 	n2 = NewNFA()
// 	n2.AppendTransition('b')
// 	n2.ApplyKleene()
//
// 	n1.AddBranch(n2)
// 	test.Eq(
// 		"map["+
// 			"0:{2 [{128 2} {128 5}]} "+
// 			"1:{0 [{128 3}]} "+
// 			"2:{0 [{97 1}]} "+
// 			"3:{0 [{128 2} {128 7}]} "+
// 			"4:{0 [{128 6}]} "+
// 			"5:{0 [{98 4}]} "+
// 			"6:{0 [{128 5} {128 7}]} "+
// 			"7:{4 []}"+
// 			"]",
// 		fmt.Sprint(n1),
// 		t,
// 	)
// }
//
// func TestNFAAppendNFAWithBranches(t *testing.T) {
// 	n1 := NewNFA()
// 	n1.AppendTransition('a')
//
// 	n2 := NewNFA()
// 	n2.AppendTransition('b')
//
// 	n3 := NewNFA()
// 	n3.AppendTransition('c')
//
// 	n1.AddBranch(n2)
// 	n1.AppendNFA(n3)
//
// 	test.Eq(
// 		"map["+
// 			"0:{2 [{97 1} {98 2}]} "+
// 			"1:{0 [{128 3}]} "+
// 			"2:{0 [{128 3}]} "+
// 			"3:{0 [{99 4}]} "+
// 			"4:{4 []}"+
// 			"]",
// 		fmt.Sprint(n1),
// 		t,
// 	)
// }
//
// func TestNFANestedBranches(t *testing.T) {
// 	n1 := NewNFA()
// 	n1.AppendTransition('a')
//
// 	n2 := NewNFA()
// 	n2.AppendTransition('b')
//
// 	n3 := NewNFA()
// 	n3.AppendTransition('c')
//
// 	n4 := NewNFA()
// 	n4.AppendTransition('d')
//
// 	n1.AddBranch(n2)
// 	n3.AddBranch(n4)
// 	n1.AddBranch(n3)
//
// 	test.Eq(
// 		"map["+
// 			"0:{2 [{97 1} {98 2} {99 4} {100 5}]} "+
// 			"1:{0 [{128 3}]} "+
// 			"2:{0 [{128 3}]} "+
// 			"3:{0 [{128 7}]} "+
// 			"4:{0 [{128 6}]} "+
// 			"5:{0 [{128 6}]} "+
// 			"6:{0 [{128 7}]} "+
// 			"7:{4 []}"+
// 			"]",
// 		fmt.Sprint(n1),
// 		t,
// 	)
// }
//
// func TestNFANestedBranchesWithKleene(t *testing.T) {
// 	n1 := NewNFA()
// 	n1.AppendTransition('a')
// 	n1.ApplyKleene()
//
// 	n2 := NewNFA()
// 	n2.AppendTransition('b')
// 	n2.ApplyKleene()
//
// 	n3 := NewNFA()
// 	n3.AppendTransition('c')
// 	n3.ApplyKleene()
//
// 	n4 := NewNFA()
// 	n4.AppendTransition('d')
// 	n4.ApplyKleene()
//
// 	n1.AddBranch(n2)
// 	test.Eq(
// 		"map["+
// 			"0:{2 [{128 2} {128 5}]} "+
// 			"1:{0 [{128 3}]} "+
// 			"2:{0 [{97 1}]} "+
// 			"3:{0 [{128 2} {128 7}]} "+
// 			"4:{0 [{128 6}]} "+
// 			"5:{0 [{98 4}]} "+
// 			"6:{0 [{128 5} {128 7}]} "+
// 			"7:{4 []}"+
// 			"]",
// 		fmt.Sprint(n1),
// 		t,
// 	)
//
// 	n3.AddBranch(n4)
// 	test.Eq(
// 		"map["+
// 			"0:{2 [{128 2} {128 5}]} "+
// 			"1:{0 [{128 3}]} "+
// 			"2:{0 [{99 1}]} "+
// 			"3:{0 [{128 2} {128 7}]} "+
// 			"4:{0 [{128 6}]} "+
// 			"5:{0 [{100 4}]} "+
// 			"6:{0 [{128 5} {128 7}]} "+
// 			"7:{4 []}"+
// 			"]",
// 		fmt.Sprint(n3),
// 		t,
// 	)
//
// 	n1.AddBranch(n3)
// 	test.Eq(
// 		"map["+
// 			"0:{2 [{128 2} {128 5} {128 9} {128 12}]} "+
// 			"1:{0 [{128 3}]} "+
// 			"2:{0 [{97 1}]} "+
// 			"3:{0 [{128 2} {128 7}]} "+
// 			"4:{0 [{128 6}]} "+
// 			"5:{0 [{98 4}]} "+
// 			"6:{0 [{128 5} {128 7}]} "+
// 			"7:{0 [{128 15}]} "+
// 			"8:{0 [{128 10}]} "+
// 			"9:{0 [{99 8}]} "+
// 			"10:{0 [{128 9} {128 14}]} "+
// 			"11:{0 [{128 13}]} "+
// 			"12:{0 [{100 11}]} "+
// 			"13:{0 [{128 12} {128 14}]} "+
// 			"14:{0 [{128 15}]} "+
// 			"15:{4 []"+
// 			"}]",
// 		fmt.Sprint(n1),
// 		t,
// 	)
// }
//
// func TestNFAWrapBranchesInKleene(t *testing.T) {
// 	n1 := NewNFA()
// 	n1.AppendTransition('a')
// 	n1.ApplyKleene()
//
// 	n2 := NewNFA()
// 	n2.AppendTransition('b')
// 	n2.ApplyKleene()
//
// 	n1.AddBranch(n2)
// 	test.Eq(
// 		"map["+
// 			"0:{2 [{128 2} {128 5}]} "+
// 			"1:{0 [{128 3}]} "+
// 			"2:{0 [{97 1}]} "+
// 			"3:{0 [{128 2} {128 7}]} "+
// 			"4:{0 [{128 6}]} "+
// 			"5:{0 [{98 4}]} "+
// 			"6:{0 [{128 5} {128 7}]} "+
// 			"7:{4 []}"+
// 			"]",
// 		fmt.Sprint(n1),
// 		t,
// 	)
//
// 	n1.ApplyKleene()
// 	test.Eq(
// 		"map["+
// 			"0:{2 [{128 8}]} "+
// 			"1:{0 [{128 3}]} "+
// 			"2:{0 [{97 1}]} "+
// 			"3:{0 [{128 2} {128 7}]} "+
// 			"4:{0 [{128 6}]} "+
// 			"5:{0 [{98 4}]} "+
// 			"6:{0 [{128 5} {128 7}]} "+
// 			"7:{0 [{128 9}]} "+
// 			"8:{0 [{128 2} {128 5}]} "+
// 			"9:{4 [{128 8}]}"+
// 			"]",
// 		fmt.Sprint(n1),
// 		t,
// 	)
// }
//
// func TestNFAAllParts(t *testing.T) {
// 	n1 := NewNFA()
// 	n1.AppendTransition('a')
// 	n1.ApplyKleene()
// 	n1.AppendTransition('b')
//
// 	n2 := NewNFA()
// 	n2.AppendTransition('c')
// 	n2.ApplyKleene()
//
// 	n1.AppendNFA(n2)
//
// 	n3 := NewNFA()
// 	n3.AppendTransition('d')
//
// 	n1.AddBranch(n3)
// 	n1.ApplyKleene()
// 	n1.AppendTransition('e')
//
// 	test.Eq(
// 		"map["+
// 			"0:{2 [{128 10}]} "+
// 			"1:{0 [{128 3}]} "+
// 			"2:{0 [{97 1}]} "+
// 			"3:{0 [{128 2} {98 4}]} "+
// 			"4:{0 [{128 6}]} "+
// 			"5:{0 [{128 7}]} "+
// 			"6:{0 [{99 5}]} "+
// 			"7:{0 [{128 6} {128 9}]} "+
// 			"8:{0 [{128 9}]} "+
// 			"9:{0 [{128 11}]} "+
// 			"10:{0 [{128 2} {100 8}]} "+
// 			"11:{0 [{128 10} {101 12}]} "+
// 			"12:{4 []"+
// 			"}]",
// 		fmt.Sprint(n1),
// 		t,
// 	)
// }