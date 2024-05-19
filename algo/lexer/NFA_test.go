package lexer

import (
	"fmt"
	"testing"

	"github.com/barbell-math/util/test"
)

func TestNFAAppendTransitionEmpty(t *testing.T) {
	n:=NFA{}
	n.AppendTransition('a')
	test.Eq("{map[0:{2 [{97 1}]} 1:{4 []}]}",fmt.Sprint(n),t)
}

func TestNFAAppendTransition(t *testing.T) {
	n:=NewNFA()
	n.AppendTransition('a')
	test.Eq("{map[0:{2 [{97 1}]} 1:{4 []}]}",fmt.Sprint(n),t)

	n.AppendTransition('b')
	test.Eq(
		"{map[0:{2 [{97 1}]} 1:{0 [{98 2}]} 2:{4 []}]}",
		fmt.Sprint(n),
		t,
	)

	n.AppendTransition('c')
	test.Eq(
		"{map["+
			"0:{2 [{97 1}]} "+
			"1:{0 [{98 2}]} "+
			"2:{0 [{99 3}]} "+
			"3:{4 []}"+
		"]}",
		fmt.Sprint(n),
		t,
	)
}

func TestNFAAppendNFAEmpty(t *testing.T) {
	op:=func(n1 *NFA, n2 *NFA) {
		n1.AppendNFA(n2)
		test.Eq(
			"&{map[0:{2 [{99 1}]} 1:{0 [{100 2}]} 2:{4 []}]}",
			fmt.Sprint(n1),
			t,
		)
	}

	n1:=NFA{}
	n2:=NewNFA()
	n2.AppendTransition('c')
	n2.AppendTransition('d')
	op(&n1,&n2)

	n1=NewNFA()
	n2=NewNFA()
	n2.AppendTransition('c')
	n2.AppendTransition('d')
	op(&n1,&n2)

	n1=NewNFA()
	n1.AppendTransition('c')
	n1.AppendTransition('d')
	n2=NFA{}
	op(&n1,&n2)

	n1=NewNFA()
	n1.AppendTransition('c')
	n1.AppendTransition('d')
	n2=NewNFA()
	op(&n1,&n2)
}

func TestNFAAppendNFA(t *testing.T){
	n1:=NewNFA()
	n1.AppendTransition('a')
	n1.AppendTransition('b')

	n2:=NewNFA()
	n2.AppendTransition('c')
	n2.AppendTransition('d')

	n1.AppendNFA(&n2)
	test.Eq(
		"{map["+
			"0:{2 [{97 1}]} "+
			"1:{0 [{98 2}]} "+
			"2:{0 [{99 3}]} "+
			"3:{0 [{100 4}]} "+
			"4:{4 []}"+
		"]}",
		fmt.Sprint(n1),
		t,
	)
}

func TestNFAApplyKleeneEmpty(t *testing.T) {
	n1:=NewNFA()
	n1.ApplyKleene()
	test.Eq(
		"{map[0:{2 [{1 1}]} 1:{0 [{1 2}]} 2:{4 [{1 1}]}]}",
		fmt.Sprint(n1),
		t,
	)
}

func TestNFAApplyKleene(t *testing.T) {
	n1:=NewNFA()
	n1.AppendTransition('a')
	n1.ApplyKleene()
	test.Eq(
		"{map[0:{2 [{1 2}]} 1:{0 [{1 3}]} 2:{0 [{97 1}]} 3:{4 [{1 2}]}]}",
		fmt.Sprint(n1),
		t,
	)

	n1=NewNFA()
	n1.AppendTransition('a')
	n1.AppendTransition('b')
	n1.ApplyKleene()
	test.Eq(
		"{map["+
			"0:{2 [{1 3}]} "+
			"1:{0 [{98 2}]} "+
			"2:{0 [{1 4}]} "+
			"3:{0 [{97 1}]} "+
			"4:{4 [{1 3}]}"+
		"]}",
		fmt.Sprint(n1),
		t,
	)
}

func TestNFAApplyKleeneMultipleTimes(t *testing.T) {
	n1:=NewNFA()
	n1.AppendTransition('a')
	n1.ApplyKleene()
	n1.ApplyKleene()
	test.Eq(
		"{map[0:{2 [{1 2}]} 1:{0 [{1 3}]} 2:{0 [{97 1}]} 3:{4 [{1 2}]}]}",
		fmt.Sprint(n1),
		t,
	)
}

func TestNFAAddBranchEmpty(t *testing.T) {
	op:=func(n1 *NFA, n2 *NFA) {
		n1.AddBranch(n2)
		test.Eq(
			"&{map[0:{2 [{99 1}]} 1:{0 [{100 2}]} 2:{4 []}]}", 
			fmt.Sprint(n1),
			t,
		)
	}

	n1:=NFA{}
	n2:=NewNFA()
	n2.AppendTransition('c')
	n2.AppendTransition('d')
	op(&n1,&n2)

	n1=NewNFA()
	n2=NewNFA()
	n2.AppendTransition('c')
	n2.AppendTransition('d')
	op(&n1,&n2)

	n1=NewNFA()
	n1.AppendTransition('c')
	n1.AppendTransition('d')
	n2=NFA{}
	op(&n1,&n2)

	n1=NewNFA()
	n1.AppendTransition('c')
	n1.AppendTransition('d')
	n2=NewNFA()
	op(&n1,&n2)
}

func TestNFAAddBranch(t *testing.T) {
	n1:=NewNFA()
	n1.AppendTransition('a')
	n1.AppendTransition('b')

	n2:=NewNFA()
	n2.AppendTransition('c')
	n2.AppendTransition('d')

	n1.AddBranch(&n2)
	test.Eq(
		"{map["+
			"0:{2 [{97 1} {99 3}]} "+
			"1:{0 [{98 2}]} "+
			"2:{0 [{1 5}]} "+
			"3:{0 [{100 4}]} "+
			"4:{0 [{1 5}]} "+
			"5:{4 []}"+
		"]}",
		fmt.Sprint(n1),
		t,
	)
}

func TestNFAKleeneAndAppendNFA(t *testing.T) {
	n1:=NewNFA()
	n1.AppendTransition('a')
	n1.ApplyKleene()

	n2:=NewNFA()
	n2.AppendTransition('b')
	n2.AppendTransition('c')
	n2.ApplyKleene()

	n1.AppendNFA(&n2)
	test.Eq(
		"{map["+
			"0:{2 [{1 2}]} "+
			"1:{0 [{1 3}]} "+
			"2:{0 [{97 1}]} "+
			"3:{0 [{1 2} {1 6}]} "+
			"4:{0 [{99 5}]} "+
			"5:{0 [{1 7}]} "+
			"6:{0 [{98 4}]} "+
			"7:{4 [{1 6}]}"+
		"]}",
		fmt.Sprint(n1),
		t,
	)
}

func TestNFABranchWithKleenedNFA(t *testing.T) {
	n1:=NewNFA()
	n1.AppendTransition('a')
	n1.ApplyKleene()

	n2:=NewNFA()
	n2.AppendTransition('b')

	n1.AddBranch(&n2)
	test.Eq(
		"{map["+
			"0:{2 [{1 2} {98 4}]} "+
			"1:{0 [{1 3}]} "+
			"2:{0 [{97 1}]} "+
			"3:{0 [{1 2} {1 5}]} "+
			"4:{0 [{1 5}]} "+
			"5:{4 []}"+
		"]}",
		fmt.Sprint(n1),
		t,
	)

	n1=NewNFA()
	n1.AppendTransition('a')
	n1.ApplyKleene()

	n2=NewNFA()
	n2.AppendTransition('b')
	n2.ApplyKleene()

	n1.AddBranch(&n2)
	test.Eq(
		"{map["+
			"0:{2 [{1 2} {1 5}]} "+
			"1:{0 [{1 3}]} "+
			"2:{0 [{97 1}]} "+
			"3:{0 [{1 2} {1 7}]} "+
			"4:{0 [{1 6}]} "+
			"5:{0 [{98 4}]} "+
			"6:{0 [{1 5} {1 7}]} "+
			"7:{4 []}"+
		"]}",
		fmt.Sprint(n1),
		t,
	)
}

func TestNFAAppendNFAWithBranches(t *testing.T){
	n1:=NewNFA()
	n1.AppendTransition('a')

	n2:=NewNFA()
	n2.AppendTransition('b')

	n3:=NewNFA()
	n3.AppendTransition('c')

	n1.AddBranch(&n2)
	n1.AppendNFA(&n3)

	test.Eq(
		"{map["+
			"0:{2 [{97 1} {98 2}]} "+
			"1:{0 [{1 3}]} "+
			"2:{0 [{1 3}]} "+
			"3:{0 [{99 4}]} "+
			"4:{4 []}"+
		"]}",
		fmt.Sprint(n1),
		t,
	)
}

func TestNFANestedBranches(t *testing.T) {
	n1:=NewNFA()
	n1.AppendTransition('a')

	n2:=NewNFA()
	n2.AppendTransition('b')

	n3:=NewNFA()
	n3.AppendTransition('c')

	n4:=NewNFA()
	n4.AppendTransition('d')

	n1.AddBranch(&n2)
	n3.AddBranch(&n4)
	n1.AddBranch(&n3)

	test.Eq(
		"{map["+
			"0:{2 [{97 1} {98 2} {99 4} {100 5}]} "+
			"1:{0 [{1 3}]} "+
			"2:{0 [{1 3}]} "+
			"3:{0 [{1 7}]} "+
			"4:{0 [{1 6}]} "+
			"5:{0 [{1 6}]} "+
			"6:{0 [{1 7}]} "+
			"7:{4 []}"+
		"]}",
		fmt.Sprint(n1),
		t,
	)
}

func TestNFANestedBranchesWithKleene(t *testing.T) {
	n1:=NewNFA()
	n1.AppendTransition('a')
	n1.ApplyKleene()

	n2:=NewNFA()
	n2.AppendTransition('b')
	n2.ApplyKleene()

	n3:=NewNFA()
	n3.AppendTransition('c')
	n3.ApplyKleene()

	n4:=NewNFA()
	n4.AppendTransition('d')
	n4.ApplyKleene()

	n1.AddBranch(&n2)
	test.Eq(
		"{map["+
			"0:{2 [{1 2} {1 5}]} "+
			"1:{0 [{1 3}]} "+
			"2:{0 [{97 1}]} "+
			"3:{0 [{1 2} {1 7}]} "+
			"4:{0 [{1 6}]} "+
			"5:{0 [{98 4}]} "+
			"6:{0 [{1 5} {1 7}]} "+
			"7:{4 []}"+
		"]}",
		fmt.Sprint(n1),
		t,
	)

	n3.AddBranch(&n4)
	test.Eq(
		"{map["+
			"0:{2 [{1 2} {1 5}]} "+
			"1:{0 [{1 3}]} "+
			"2:{0 [{99 1}]} "+
			"3:{0 [{1 2} {1 7}]} "+
			"4:{0 [{1 6}]} "+
			"5:{0 [{100 4}]} "+
			"6:{0 [{1 5} {1 7}]} "+
			"7:{4 []}"+
		"]}",
		fmt.Sprint(n3),
		t,
	)

	n1.AddBranch(&n3)
	test.Eq(
		"{map["+
			"0:{2 [{1 2} {1 5} {1 9} {1 12}]} "+
			"1:{0 [{1 3}]} "+
			"2:{0 [{97 1}]} "+
			"3:{0 [{1 2} {1 7}]} "+
			"4:{0 [{1 6}]} "+
			"5:{0 [{98 4}]} "+
			"6:{0 [{1 5} {1 7}]} "+
			"7:{0 [{1 15}]} "+
			"8:{0 [{1 10}]} "+
			"9:{0 [{99 8}]} "+
			"10:{0 [{1 9} {1 14}]} "+
			"11:{0 [{1 13}]} "+
			"12:{0 [{100 11}]} "+
			"13:{0 [{1 12} {1 14}]} "+
			"14:{0 [{1 15}]} "+
			"15:{4 []"+
		"}]}",
		fmt.Sprint(n1),
		t,
	)
}

func TestNFAWrapBranchesInKleene(t *testing.T) {
	n1:=NewNFA()
	n1.AppendTransition('a')
	n1.ApplyKleene()

	n2:=NewNFA()
	n2.AppendTransition('b')
	n2.ApplyKleene()

	n1.AddBranch(&n2)
	test.Eq(
		"{map["+
			"0:{2 [{1 2} {1 5}]} "+
			"1:{0 [{1 3}]} "+
			"2:{0 [{97 1}]} "+
			"3:{0 [{1 2} {1 7}]} "+
			"4:{0 [{1 6}]} "+
			"5:{0 [{98 4}]} "+
			"6:{0 [{1 5} {1 7}]} "+
			"7:{4 []}"+
		"]}",
		fmt.Sprint(n1),
		t,
	)

	n1.ApplyKleene()
	test.Eq(
		"{map["+
			"0:{2 [{1 8}]} "+
			"1:{0 [{1 3}]} "+
			"2:{0 [{97 1}]} "+
			"3:{0 [{1 2} {1 7}]} "+
			"4:{0 [{1 6}]} "+
			"5:{0 [{98 4}]} "+
			"6:{0 [{1 5} {1 7}]} "+
			"7:{0 [{1 9}]} "+
			"8:{0 [{1 2} {1 5}]} "+
			"9:{4 [{1 8}]}"+
		"]}",
		fmt.Sprint(n1),
		t,
	)
}

func TestNFAAllParts(t *testing.T) {
	n1:=NewNFA()
	n1.AppendTransition('a')
	n1.ApplyKleene()
	n1.AppendTransition('b')

	n2:=NewNFA()
	n2.AppendTransition('c')
	n2.ApplyKleene()

	n1.AppendNFA(&n2)

	n3:=NewNFA()
	n3.AppendTransition('d')

	n1.AddBranch(&n3)
	n1.ApplyKleene()
	n1.AppendTransition('e')

	test.Eq(
		"{map["+
			"0:{2 [{1 10}]} "+
			"1:{0 [{1 3}]} "+
			"2:{0 [{97 1}]} "+
			"3:{0 [{1 2} {98 4}]} "+
			"4:{0 [{1 6}]} "+
			"5:{0 [{1 7}]} "+
			"6:{0 [{99 5}]} "+
			"7:{0 [{1 6} {1 9}]} "+
			"8:{0 [{1 9}]} "+
			"9:{0 [{1 11}]} "+
			"10:{0 [{1 2} {100 8}]} "+
			"11:{0 [{1 10} {101 12}]} "+
			"12:{4 []"+
		"}]}",
		fmt.Sprint(n1),
		t,
	)
}
