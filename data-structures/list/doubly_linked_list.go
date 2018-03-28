package list

import "fmt"

// DNode defines doubly linked list node.
// It holds value, and points to next and previous nodes.
type DNode struct {
	Value    interface{}
	Next     *DNode
	Previous *DNode
}

// NewDNode returns a new doubly linked list node.
func NewDNode(value interface{}) *DNode {
	return &DNode{Value: value}
}

// String prints DNode in a nice format.
func (n *DNode) String() string {
	if n == nil {
		return "Nil Node"
	}
	return fmt.Sprintf("%v", n.Value)
}

// DoublyLinkedList is a sequence of connected nodes.
type DoublyLinkedList struct {
	Head *DNode
	Tail *DNode
}

// NewDoublyLinkedList returns a new doubly linked list.
func NewDoublyLinkedList(value interface{}) *DoublyLinkedList {
	node := NewDNode(value)
	return &DoublyLinkedList{Head: node, Tail: node}
}

// IsEmpty tells whether the list is empty or not.
func (l *DoublyLinkedList) IsEmpty() bool {
	return l.Head == nil
}

// Length returns the number of nodes in the list.
func (l *DoublyLinkedList) Length() int {
	length := 0
	for n := l.Head; n != nil; n = n.Next {
		length++
	}
	return length
}

// Prepend prepends a new node with given value.
func (l *DoublyLinkedList) Prepend(value interface{}) {
	newNode := NewDNode(value)
	l.PrependNode(newNode)
}

// PrependNode prepends a new node into the list.
func (l *DoublyLinkedList) PrependNode(newNode *DNode) {
	if newNode == nil { return }

	if l.Head != nil {
		newNode.Next = l.Head
		l.Head.Previous = newNode
	} else {
		l.Tail = newNode
	}

	l.Head = newNode
}

// Append appends a new node with given value.
func (l *DoublyLinkedList) Append(value interface{}) {
	newNode := NewDNode(value)
	l.AppendNode(newNode)
}

// AppendNode appends a new node into the list.
func (l *DoublyLinkedList) AppendNode(newNode *DNode) {
	if newNode == nil { return }

	if l.Tail != nil {
		newNode.Previous = l.Tail
		l.Tail.Next = newNode
	} else {
		l.Head = newNode
	}

	l.Tail = newNode
}

// Insert inserts a new node with given value at given index.
func (l *DoublyLinkedList) Insert(value interface{}, index int) {
	newNode := NewDNode(value)
	l.InsertNode(newNode, index)
}

// InsertNode inserts a new node at given index in the list.
func (l *DoublyLinkedList) InsertNode(newNode *DNode, index int) {
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

		newNode.Previous = prevNode
		newNode.Next = nextNode

		prevNode.Next = newNode
		nextNode.Previous = newNode
	}
}

// NodeAt returns the node at given index.
func (l *DoublyLinkedList) NodeAt(index int) *DNode {
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
func (l *DoublyLinkedList) Find(value interface{}) *DNode {
	for n := l.Head; n != nil; n = n.Next {
		if n.Value == value {
			return n
		}
	}
	return nil
}

// Remove removes node with given value.
func (l *DoublyLinkedList) Remove(value interface{}) *DNode {
	node := l.Find(value)
	return l.RemoveNode(node)
}

// RemoveAt removes node at given index.
func (l *DoublyLinkedList) RemoveAt(index int) *DNode {
	node := l.NodeAt(index)
	return l.RemoveNode(node)
}

// RemoveNode removes the given node.
func (l *DoublyLinkedList) RemoveNode(node *DNode) *DNode {
	if l.Head == nil { return nil }
	if node == nil { return nil }

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
func (l *DoublyLinkedList) RemoveAll() {
	l.Head = nil
	l.Tail = nil
}

// String prints doubly linked list in a nice format.
func (l *DoublyLinkedList) String() string {
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
