package linkedlist

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	dummy := &Node{}
	c1 := head
	c2 := dummy
	for c1 != nil {
		newNode := &Node{Val: c1.Val}
		m[c1] = newNode
		c2.Next = newNode
		c1 = c1.Next
		c2 = c2.Next
	}
	c1 = head
	c2 = dummy.Next
	for c1 != nil {
		c2.Random = m[c1.Random]
		c1 = c1.Next
		c2 = c2.Next
	}
	return dummy.Next
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
