package list

import "fmt"

// Node is a block of data. Every node holds some data,
// and points to next and previous node.
type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

// NewNode returns a new node.
func NewNode(value int) *Node {
	return &Node{Value: value}
}

// String prints Node in a nice format.
func (n *Node) String() string {
	if n == nil {
		return "Nil Node"
	}
	return fmt.Sprintf("(%d)", n.Value)
}

// LinkedList is a sequence of connected nodes.
type LinkedList struct {
	Head *Node
	Tail *Node
}

// NewLinkedList returns a new linked list.
func NewLinkedList(value int) *LinkedList {
	node := NewNode(value)
	return &LinkedList{Head: node, Tail: node}
}

// IsEmpty tells whether the list is empty or not.
func (l *LinkedList) IsEmpty() bool {
	return l.Head == nil
}

// Append appends a new node into the list.
func (l *LinkedList) Append(value int) {
	newNode := NewNode(value)

	if l.Tail != nil {
		newNode.Previous = l.Tail
		l.Tail.Next = newNode
	} else {
		l.Head = newNode
	}

	l.Tail = newNode
}

// NodeAt returns the node at given index.
func (l *LinkedList) NodeAt(index int) *Node {
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
func (l *LinkedList) Find(value int) *Node {
	for n := l.Head; n != nil; n = n.Next {
		if n.Value == value {
			return n
		}
	}
	return nil
}

// Remove removes the node with given value.
func (l *LinkedList) Remove(value int) *Node {
	node := l.Find(value)
	if node == nil {
		return nil
	}

	prevNode := node.Previous
	nextNode := node.Next

	if prevNode != nil {
		prevNode.Next = nextNode
	} else {
		l.Head = nextNode
	}

	if nextNode != nil {
		nextNode.Previous = prevNode
	} else {
		l.Tail = prevNode
	}

	node.Previous = nil
	node.Next = nil

	return node
}

// RemoveAll removes all nodes.
func (l *LinkedList) RemoveAll() {
	l.Head = nil
	l.Tail = nil
}

// String prints stack in a nice format.
func (l *LinkedList) String() string {
	str := ""
	for n := l.Head; n != nil; n = n.Next {
		str += n.String() + " "
	}
	return str
}
