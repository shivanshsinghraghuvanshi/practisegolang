package main

import (
	"fmt"
	. "github.com/shivanshsinghraghuvanshi/practisegolang/amazon/linkedlist"
)

func main() {
	//output := MinPartitions("32")
	//fmt.Println(output)
	//points := [][]int{
	//	{1, 3}, {-2, -2},
	//}
	//fmt.Println(KClosest(points, 1))

	//	fmt.Println(MyAtoi("-21474836460"))
	//fmt.Println(MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	//IntToRoman(3)
	//RomanToInt("III")
	// ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	//l1 := new(ListNode)
	//l1.Val = 1
	//l1.Next = &ListNode{
	//	Val: 2,
	//	Next: &ListNode{
	//		Val:  4,
	//		Next: nil,
	//	},
	//}
	//l2 := new(ListNode)
	//l2.Val = 1
	//l2.Next = &ListNode{
	//	Val: 3,
	//	Next: &ListNode{
	//		Val:  4,
	//		Next: nil,
	//	},
	//}
	//MergeTwoLists(l1, l2)

	l := NewMyLinkedList()
	l.Add(10)
	l.Add(20)
	l.Add(9)
	l.Add(11)
	// l.Print()
	//l.PrintOnlyValues()
	//fmt.Printf("\n %d", l.Length())
	//
	//fmt.Printf("\n %t", l.Search(10))
	//
	//fmt.Printf("\n %t \n", l.Search(40))
	//
	//found,node := l.FindNode(3)
	//fmt.Printf("%t \t %v",found,node)

	i, node := l.Middle()
	fmt.Printf("%d \t %v \n", i, node)
	l.PrintOnlyValues()
	fmt.Printf("\n")
	_, _ = l.InsertAt(1, 18)
	l.PrintOnlyValues()
	fmt.Printf("\n")
	_, _ = l.DeleteAt(1)
	l.PrintOnlyValues()
	fmt.Printf("\n")
}

// reverse a string
