package list

import "fmt"

// SNode defines singly linked list node.
// It holds value, and only has a reference to the next node.
type SNode struct {
	Value int
	Next  *SNode
}

// NewSNode returns a new singly linked list node.
func NewSNode(value int) *SNode {
	return &SNode{Value: value}
}

// String prints SNode in a nice format.
func (n *SNode) String() string {
	if n == nil {
		return "Nil Node"
	}
	return fmt.Sprintf("%d", n.Value)
}

// SinglyLinkedList is a sequence of connected nodes.
type SinglyLinkedList struct {
	Head *SNode
	Tail *SNode
}

// NewSinglyLinkedList returns a new singly linked list.
func NewSinglyLinkedList(value int) *SinglyLinkedList {
	node := NewSNode(value)
	return &SinglyLinkedList{Head: node, Tail: node}
}

// IsEmpty tells whether the list is empty or not.
func (l *SinglyLinkedList) IsEmpty() bool {
	return l.Head == nil
}

// Length returns the number of nodes in the list.
func (l *SinglyLinkedList) Length() int {
	length := 0
	for n := l.Head; n != nil; n = n.Next {
		length++
	}
	return length
}

// Prepend prepends a new node with given value.
func (l *SinglyLinkedList) Prepend(value int) {
	newNode := NewSNode(value)
	l.PrependNode(newNode)
}

// PrependNode prepends a new node into the list.
func (l *SinglyLinkedList) PrependNode(newNode *SNode) {
	if newNode == nil { return }

	if l.Head != nil {
		newNode.Next = l.Head
	} else {
		l.Tail = newNode
	}

	l.Head = newNode
}

// Append appends a new node with given value.
func (l *SinglyLinkedList) Append(value int) {
	newNode := NewSNode(value)
	l.AppendNode(newNode)
}

// AppendNode appends a new node into the list.
func (l *SinglyLinkedList) AppendNode(newNode *SNode) {
	if newNode == nil { return }

	if l.Tail != nil {
		l.Tail.Next = newNode
	} else {
		l.Head = newNode
	}

	l.Tail = newNode
}

// Insert inserts a new node with given value at given index.
func (l *SinglyLinkedList) Insert(value, index int) {
	newNode := NewSNode(value)
	l.InsertNode(newNode, index)
}

// InsertNode inserts a new node at given index in the list.
func (l *SinglyLinkedList) InsertNode(newNode *SNode, index int) {
	if newNode == nil { return }

	if index < 0 || index > l.Length() {
		return
	}

	switch {
	case index == 0:
		l.PrependNode(newNode)
	case index == l.Length():
		l.AppendNode(newNode)
	default:
		prevNode := l.NodeAt(index - 1)
		nextNode := prevNode.Next

		prevNode.Next = newNode
		newNode.Next = nextNode
	}
}

// NodeAt returns the node at given index.
func (l *SinglyLinkedList) NodeAt(index int) *SNode {
	if index >= 0 {
		for n := l.Head; n != nil; n = n.Next {
			if index == 0 {
				return n
			}
			index--
		}
	}
	return nil
}

// Find searches for node with given value.
func (l *SinglyLinkedList) Find(value int) *SNode {
	for n := l.Head; n != nil; n = n.Next {
		if n.Value == value {
			return n
		}
	}
	return nil
}

// Remove removes node with given value.
func (l *SinglyLinkedList) Remove(value int) *SNode {
	node := l.Find(value)
	return l.RemoveNode(node)
}

// RemoveAt removes node at given index.
func (l *SinglyLinkedList) RemoveAt(index int) *SNode {
	node := l.NodeAt(index)
	return l.RemoveNode(node)
}

// RemoveNode removes the given node.
func (l *SinglyLinkedList) RemoveNode(node *SNode) *SNode {
	if l.Head == nil { return nil }
	if node == nil { return nil }

	var prevNode *SNode

	for n := l.Head; n != node; n = n.Next {
		if n.Next == nil {
			return nil
		}
		prevNode = n
	}

	if prevNode != nil {
		prevNode.Next = node.Next
	} else {
		l.Head = node.Next
	}

	if node.Next == nil {
		l.Tail = prevNode
	}

	node.Next = nil

	return node
}

// RemoveAll removes all nodes.
func (l *SinglyLinkedList) RemoveAll() {
	l.Head = nil
	l.Tail = nil
}

// String prints singly linked list in a nice format.
func (l *SinglyLinkedList) String() string {
	if l.Head == nil {
		return "[]"
	}

	str := fmt.Sprintf("[%s", l.Head.String())
	for n := l.Head.Next; n != nil; n = n.Next {
		str += ", " + n.String()
	}
	str += "]"
	return str
}
