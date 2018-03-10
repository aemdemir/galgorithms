package tree

import "fmt"

// Node is a block of data. Every node holds some data,
// and has a left and right child.
type Node struct {
	Value      int
	LeftChild  *Node
	RightChild *Node
}

// NewNode creates and returns a new node without children.
func NewNode(value int) *Node {
	return &Node{Value: value}
}

func (n *Node) count() int {
	if n == nil {
		return 0
	}
	return n.LeftChild.count() + 1 + n.RightChild.count()
}

func (n *Node) insert(value int) {
	if n == nil {
		return
	}

	switch {
	case n.Value > value:
		if n.LeftChild == nil {
			n.LeftChild = NewNode(value)
			return
		}
		n.LeftChild.insert(value)
	default:
		if n.RightChild == nil {
			n.RightChild = NewNode(value)
			return
		}
		n.RightChild.insert(value)
	}
}

func (n *Node) search(value int) *Node {
	if n == nil {
		return nil
	}

	switch {
	case n.Value > value:
		return n.LeftChild.search(value)
	case n.Value < value:
		return n.RightChild.search(value)
	default:
		return n
	}
}

func (n *Node) findDelete(value int) (node, parent *Node) {
	if n == nil {
		return nil, nil
	}

	node = n
	parent = n

	for node.Value != value {
		parent = node

		if node.Value > value {
			node = node.LeftChild
		} else {
			node = node.RightChild
		}

		if node == nil {
			return nil, nil
		}
	}

	return node, parent
}

func (n *Node) findSuccessor() *Node {
	if n == nil {
		return nil
	}

	sucParent := n
	suc := n.RightChild

	for ; suc.LeftChild != nil; suc = suc.LeftChild {
		sucParent = suc
	}

	if suc != n.RightChild {
		sucParent.LeftChild = suc.RightChild
		suc.RightChild = n.RightChild
	}

	return suc
}

// traverseInOrder first processes the left node, then the current node itself, and finally right node.
// It goes through the nodes in ascending order.
func (n *Node) traverseInOrder(process func(int)) {
	if n == nil {
		return
	}

	n.LeftChild.traverseInOrder(process)
	process(n.Value)
	n.RightChild.traverseInOrder(process)
}

// traversePreOrder first processes the current node, then visits its left and right node.
func (n *Node) traversePreOrder(process func(int)) {
	if n == nil {
		return
	}

	process(n.Value)
	n.LeftChild.traversePreOrder(process)
	n.RightChild.traversePreOrder(process)
}

// traversePostOrder first visits the left and right node, and processes the current node itself.
func (n *Node) traversePostOrder(process func(int)) {
	if n == nil {
		return
	}

	n.LeftChild.traversePostOrder(process)
	n.RightChild.traversePostOrder(process)
	process(n.Value)
}

// String prints Node in a nice format.
func (n *Node) String() string {
	if n == nil {
		return "nil node"
	}
	str := fmt.Sprintf("(%d)", n.Value)
	if n.LeftChild != nil {
		str += fmt.Sprintf("\n (%d)left  ->  %s", n.Value, n.LeftChild)
	}
	if n.RightChild != nil {
		str += fmt.Sprintf("\n (%d)right ->  %s", n.Value, n.RightChild)
	}
	return str
}

/*

	Binary Search Tree

*/

// BinarySearchTree is a special binary tree that performs insertions and deletions
// such that the tree is always sorted.
type BinarySearchTree struct {
	Root *Node
}

// NewBinarySearchTree creates and returns a new BinarySearchTree without root.
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

// NewBinarySearchTreeWith creates and returns a new BinarySearchTree.
func NewBinarySearchTreeWith(value int) *BinarySearchTree {
	return &BinarySearchTree{
		Root: &Node{Value: value},
	}
}

// NewBinarySearchTreeWithRoot creates and returns a new BinarySearchTree.
func NewBinarySearchTreeWithRoot(root *Node) *BinarySearchTree {
	return &BinarySearchTree{
		Root: root,
	}
}

// Count returns the number of nodes in the tree.
func (bst *BinarySearchTree) Count() int {
	if bst.Root == nil {
		return 0
	}
	return bst.Root.count()
}

// Insert inserts a new node into binary search tree.
// It keeps "always sorted" property.
func (bst *BinarySearchTree) Insert(value int) {
	if bst.Root == nil {
		bst.Root = NewNode(value)
		return
	}
	bst.Root.insert(value)
}

// Search searches for node that holds given value.
func (bst *BinarySearchTree) Search(value int) *Node {
	if bst.Root == nil {
		return nil
	}
	return bst.Root.search(value)
}

// Delete deletes node with given value.
func (bst *BinarySearchTree) Delete(value int) *Node {
	if bst.Root == nil {
		return nil
	}

	node, parent := bst.Root.findDelete(value)
	if node == nil || parent == nil {
		return nil
	}

	// Here we have node to be deleted and its parent.
	switch {
	// No children
	case node.LeftChild == nil && node.RightChild == nil:
		if node == bst.Root {
			bst.Root = nil
		} else if node == parent.LeftChild {
			parent.LeftChild = nil
		} else {
			parent.RightChild = nil
		}
	// One child: right
	case node.LeftChild == nil:
		if node == bst.Root {
			bst.Root = bst.Root.RightChild
		} else if node == parent.LeftChild {
			parent.LeftChild = node.RightChild
		} else {
			parent.RightChild = node.RightChild
		}
	// One child: left
	case node.RightChild == nil:
		if node == bst.Root {
			bst.Root = bst.Root.LeftChild
		} else if node == parent.LeftChild {
			parent.LeftChild = node.LeftChild
		} else {
			parent.RightChild = node.LeftChild
		}
	// Two children
	default:
		sucessor := node.findSuccessor()

		if sucessor == nil {
			return nil
		}

		if node == bst.Root {
			bst.Root = sucessor
		} else if node == parent.LeftChild {
			parent.LeftChild = sucessor
		} else {
			parent.RightChild = sucessor
		}

		sucessor.LeftChild = node.LeftChild
	}

	return node
}

// TraverseInOrder first processes the left node, then the current node itself, and finally right node.
// It goes through the nodes in ascending order.
func (bst *BinarySearchTree) TraverseInOrder(process func(int)) {
	if bst.Root == nil {
		return
	}

	bst.Root.traverseInOrder(process)
}

// TraversePreOrder first processes the current node, then visits its left and right node.
func (bst *BinarySearchTree) TraversePreOrder(process func(int)) {
	if bst.Root == nil {
		return
	}

	bst.Root.traversePreOrder(process)
}

// TraversePostOrder first visits the left and right node, and processes the current node itself.
func (bst *BinarySearchTree) TraversePostOrder(process func(int)) {
	if bst.Root == nil {
		return
	}

	bst.Root.traversePostOrder(process)
}

// String prints BinarySearchTree in a nice format.
func (bst *BinarySearchTree) String() string {
	return bst.Root.String()
}
