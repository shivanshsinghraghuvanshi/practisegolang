package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{
		Val:  0,
		Next: nil,
	}
	z := r

	v := 0
	for l1 != nil || l2 != nil || v != 0 {

		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}

		curVal := (v1 + v2 + v) % 10
		v = (v1 + v2 + v) / 10

		z.Next = &ListNode{
			Val:  curVal,
			Next: nil,
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		z = z.Next
	}
	return r
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
