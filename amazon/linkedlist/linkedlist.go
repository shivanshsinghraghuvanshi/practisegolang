package linkedlist

import "fmt"

type myNode struct {
	// val value that is hold int the node
	val int
	// the next pointer
	next *myNode
	// list to which this node belongs to
	// list *myLinkedList
}

type myLinkedList struct {
	head *myNode
	// tail *myNode
	// length int
}

func NewMyLinkedList() *myLinkedList {
	return new(myLinkedList)
}

func (l *myLinkedList) Search(val int) bool {
	if l.Length() == 0 {
		return false
	}
	p := l.head
	for l.head != nil && p != nil {
		if p.val == val {
			return true
		}
		p = p.next
	}
	return false
}
func (l *myLinkedList) Length() int {
	if l.head == nil {
		return 0
	}
	p := l.head
	length := 0
	for l.head != nil && p != nil {
		length += 1
		p = p.next
	}
	return length
}
func (l *myLinkedList) Print() {
	if l.head == nil {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for l.head != nil && ptr != nil {
		fmt.Println("Node: ", ptr)
		ptr = ptr.next
	}
}

func (l *myLinkedList) PrintOnlyValues() {
	if l.head == nil {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for l.head != nil && ptr != nil {
		fmt.Printf("%d \t", ptr.val)
		ptr = ptr.next
	}
}
func (l *myLinkedList) Add(val int) {
	// Prepare Insert Object
	n := myNode{
		val:  val,
		next: nil,
	}
	if l.head == nil {
		l.head = &n
		return
	}
	p := l.head
	for p != nil {
		if p.next == nil {
			p.next = &n
			return
		}
		p = p.next
	}

}

func (l *myLinkedList) FindNode(pos int) (bool, *myNode) {
	length := l.Length()
	if pos < 0 || pos > length {
		return false, new(myNode)
	}

	p := l.head
	for i := 0; i < length; i++ {
		if i == pos-1 {
			return true, p
		}
		p = p.next
	}
	return false, nil
}

func (l *myLinkedList) Middle() (int, *myNode) {
	m := l.Length() % 2
	if m == 0 {
		return -1, nil
	}
	if ok, node := l.FindNode(m + 1); ok {
		return m + 1, node
	}
	return -1, nil
}

func (l *myLinkedList) InsertAt(pos, val int) (bool, int) {
	if pos < 0 || pos > l.Length() {
		return false, l.Length()
	}
	p := l.head
	for i := 0; i < l.Length(); i++ {
		if i == pos {
			next := p.next
			n := myNode{
				val:  val,
				next: next,
			}
			p.next = &n
			p = p.next
			return true, l.Length()
		}
	}
	return false, l.Length()
}

func (l *myLinkedList) DeleteAt(pos int) (bool, int) {
	if pos < 0 || pos > l.Length() {
		return false, l.Length()
	}
	p := l.head
	for i := 0; i < l.Length(); i++ {
		if i == pos {
			n := p.next.next
			p.next = n
			p = p.next
			return true, l.Length()
		}
	}
	return false, l.Length()
}

func (l *myLinkedList) Reverse() {
	node := l.head
	var prev *myNode

	for node != nil {
		node, prev, node.next = node.next, node, prev
	}
	// if you have to return a node return prev pointer
	//return prev
	l.head = prev
}

func Reverse(head *myNode) *myNode {
	node := head
	var prev *myNode
	for node != nil {
		node, prev, node.next = node.next, node, prev
	}
	return prev
}

func (l *myLinkedList) Sort() {
	node := l.head
	l.head = mergeSort(node, l.Length())
}

func mergeSort(l *myNode, length int) *myNode {
	return split(l, 0, length)
}

func split(l *myNode, start, end int) *myNode {
	if end-start < 2 {
		return l
	}
	// call the left recursion
	// call the right recursion
	var l1, l2 *myNode
	// call the left recursion for l1
	// TODO
	// call the right recursion l2
	// TODO
	return merge(l1, l2)
}

func merge(l1, l2 *myNode) *myNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.val < l2.val {
		l1.next = merge(l1.next, l2)
		return l1
	} else {
		l2.next = merge(l1, l2.next)
		return l2
	}

}
