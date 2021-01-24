package main

import (
	"fmt"
	. "github.com/shivanshsinghraghuvanshi/practisegolang/amazon/linkedlist"
	"sort"
)

func middleNode(head *ListNode) *ListNode {

	// fid the length of the array
	length := 0
	p := head
	for p != nil {
		length += 1
		p = p.Next
	}

	m := length % 2
	pos := 0
	if m == 0 {
		pos = (length / 2) + 1
	} else {
		pos = length / 2
	}
	var f *ListNode
	for i := 0; i < length; i++ {
		if i == pos {
			f = head
		}
		head = head.Next
	}
	return f
}

func pairwiseSwap(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	res := &ListNode{
		Val:  0,
		Next: head,
	}
	p := res
	for p.Next != nil && p.Next.Next != nil {
		f := p.Next
		s := f.Next

		t := s.Next

		s.Next = f
		f.Next = t
		p.Next = s
		p.Next = s
		p = p.Next.Next
		//prev := p.Next
		//p.Next, prev, p.Next.Next = p.Next.Next, p.Next, prev
	}
	return res.Next
}

//func main() {
//	//output := MinPartitions("32")
//	//fmt.Println(output)
//	//points := [][]int{
//	//	{1, 3}, {-2, -2},
//	//}
//	//fmt.Println(KClosest(points, 1))
//
//	//	fmt.Println(MyAtoi("-21474836460"))
//	//fmt.Println(MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
//	//IntToRoman(3)
//	//RomanToInt("III")
//	// ThreeSum([]int{-1, 0, 1, 2, -1, -4})
//	l1 := new(ListNode)
//	l1.Val = 1
//	l1.Next = &ListNode{
//		Val: 2,
//		Next: &ListNode{
//			Val: 3,
//			Next: &ListNode{
//				Val:  4,
//				Next: nil,
//			},
//		},
//	}
//	l2 := new(ListNode)
//	l2.Val = 1
//	l2.Next = &ListNode{
//		Val: 3,
//		Next: &ListNode{
//			Val:  4,
//			Next: nil,
//		},
//	}
//
//	l3 := new(ListNode)
//	l3.Val = 2
//	l3.Next = &ListNode{
//		Val:  6,
//		Next: nil,
//	}
//	pairwiseSwap(l1)
//	//KmergeKLists([]*ListNode{l1, l2, l3})
//	//MergeTwoLists(l1, l2)
//
//	//l := NewMyLinkedList()
//	//l.Add(10)
//	//l.Add(20)
//	//l.Add(9)
//	//l.Add(11)
//	// l.Print()
//	//l.PrintOnlyValues()
//	//fmt.Printf("\n %d", l.Length())
//	//
//	//fmt.Printf("\n %t", l.Search(10))
//	//
//	//fmt.Printf("\n %t \n", l.Search(40))
//	//
//	//found,node := l.FindNode(3)
//	//fmt.Printf("%t \t %v",found,node)
//
//	//i, node := l.Middle()
//	//fmt.Printf("%d \t %v \n", i, node)
//	//l.PrintOnlyValues()
//	//fmt.Printf("\n")
//	//_, _ = l.InsertAt(1, 18)
//	//l.PrintOnlyValues()
//	//fmt.Printf("\n")
//	//_, _ = l.DeleteAt(1)
//	//l.PrintOnlyValues()
//	//fmt.Printf("\n")
//	//l.Reverse()
//	//l.PrintOnlyValues()
//}

// reverse a string

//func main() {
//	// BST Tester
//
//	b := NewBST()
//	b.Insert(25)
//	b.Insert(20)
//	b.Insert(15)
//	b.Insert(27)
//	b.Insert(30)
//	b.Insert(29)
//	b.Insert(26)
//	b.Insert(22)
//	b.Insert(32)
//	fmt.Println("end of programme")
//
//	b.InOrderTraversal()
//
//	//x :=b.Get(20)
//
//	fmt.Println(IsValidBST(b.Root))
//
//	//fmt.Printf(" Minimum value of this BST is %d \n",b.GetMin())
//	//fmt.Printf(" Maximum value of this BST is %d \n",b.GetMax())
//
//
//}

func main() {

	h:= newHash(5)
	h.Add("Shirt")
	h.Add("Trouser")
	fmt.Println(h.GetAll())
	h.Delete("Shirt")
	fmt.Println(h.GetAll())
}

// arr[3,7,5,6,2]
func subsetA(arr []int) []int {

	sort.Ints(arr)
	var sumL int
	sumL = 0
	//i := len(arr) - 1
	for i := len(arr) - 1; i >= 0; i-- {
		sumL = sumL + arr[i]
		if sumL > sumRemaining(arr[:i]) {
			return arr[i+1:]
		}
	}

	return []int{}
}

func sumRemaining(rem []int) int {
	var sum int
	sum = 0
	for _, n := range rem {
		sum += n
	}
	return sum
}

type hash struct {
	d []string
	l int
	c int
}

func newHash(c int) *hash {
	return &hash{d: make([]string,c), l: 0, c: c}
}



func (h *hash) Add(name string) {
	f, _ := found(h.d, name)
	if h.l < h.c && !f {
		h.d[h.l] = name
		h.l += 1
	}
}

func (h *hash) GetAll() []string {
	return h.d
}

func (h *hash) Delete(s string) {

	f, i := found(h.d, s)
	if h.l < h.c && f {
		remove(h.d, i)
		h.l -= 1
	}
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func found(a []string, s string) (bool, int) {
	for i, x := range a {
		if x == s {
			return true, i
		}
	}
	return false, -1
}
