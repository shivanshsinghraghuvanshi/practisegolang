package treesAndgraphs

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type BST struct {
	Root *TreeNode
}

func (b *BST) NewBST() *BST {
	return &BST{Root: nil}
}

func NewBST() *BST {
	return new(BST)
}

func (b *BST) Insert(value int) {
	if b.Root == nil {
		b.Root = new(TreeNode)
		b.Root.Value = value
		return
	}
	b.Root = insertNode(b.Root, value)
}

func (b *BST) InOrderTraversal() {
	if b.Root == nil {
		return
	}
	inOrderTraversal(b.Root)
	fmt.Printf("\n")
}

func inOrderTraversal(n *TreeNode) {
	// Put fmt here for pre Order
	// fmt.Printf("%d \t", n.Value)
	if n.Left != nil {
		inOrderTraversal(n.Left)
	}
	// fmt here for in order
	fmt.Printf("%d \t", n.Value)
	if n.Right != nil {
		inOrderTraversal(n.Right)
	}
	// fmt here for post order
	//	fmt.Printf("%d \t", n.Value)
}
func insertNode(n *TreeNode, val int) *TreeNode {
	if n == nil {
		n = new(TreeNode)
		n.Value = val
	}
	if val < n.Value {
		n.Left = insertNode(n.Left, val)
	}
	if val > n.Value {
		n.Right = insertNode(n.Right, val)
	}
	return n
}

func (b *BST) Get(value int) *TreeNode {
	if b.Root.Value == value {
		return b.Root
	}
	return find(b.Root, value)
}

func find(n *TreeNode, val int) *TreeNode {
	if n.Value == val {
		return n
	}
	if val < n.Value {
		find(n.Left, val)
	}
	if val > n.Value {
		find(n.Right, val)
	}
	return n
}

func (b *BST) GetMin() int {
	return getMin(b.Root)
}

func getMin(n *TreeNode) int {
	if n.Left == nil {
		return n.Value
	} else {
		return getMin(n.Left)
	}

}

func (b *BST) GetMax() int {
	return getMax(b.Root)
}

func getMax(n *TreeNode) int {
	if n.Right == nil {
		return n.Value
	} else {
		return getMax(n.Right)
	}
}

func (b *BST) Delete(value int) {
	n := b.Get(value)
	deleteNode(n)
}

func deleteNode(n *TreeNode) {

	// Case 1 : if the n is leaf node i.e. no left right parent
	// Case 2: Have only one child i.e. either left or right
	// Case 3: Have both child
	if n.Left == nil && n.Right == nil {
		n = nil
	} else if n.Left != nil && n.Right != nil {

	} else {

	}
}
