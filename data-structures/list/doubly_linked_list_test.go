package list

import (
	"testing"
)

func TestDoublyListPrepend(t *testing.T) {
	l := &DoublyLinkedList{}
	l.Prepend(16)
	l.Prepend(8)
	l.Prepend(4)

	newNode := NewDNode(2)
	l.PrependNode(newNode)

	if l.IsEmpty() {
		t.Errorf("List expected to be not empty, instead got IsEmpty: %v", l.IsEmpty())
	}
	if actual := l.Length(); actual != 4 {
		t.Errorf("Expected %v, instead got %v", 4, actual)
	}
	if actual := l.Head.Value; actual != 2 {
		t.Errorf("Expected %v, instead got %v", 2, actual)
	}
	if actual := l.NodeAt(1).Value; actual != 4 {
		t.Errorf("Expected %v, instead got %v", 4, actual)
	}
	if actual := l.Tail.Value; actual != 16 {
		t.Errorf("Expected %v, instead got %v", 16, actual)
	}
}

func TestDoublyListAppend(t *testing.T) {
	l := &DoublyLinkedList{}
	l.Append(16)
	l.Append(8)
	l.Append(4)

	newNode := NewDNode(2)
	l.AppendNode(newNode)

	if l.IsEmpty() {
		t.Errorf("List expected to be not empty, instead got IsEmpty: %v", l.IsEmpty())
	}
	if actual := l.Length(); actual != 4 {
		t.Errorf("Expected %v, instead got %v", 4, actual)
	}
	if actual := l.Head.Value; actual != 16 {
		t.Errorf("Expected %v, instead got %v", 16, actual)
	}
	if actual := l.NodeAt(1).Value; actual != 8 {
		t.Errorf("Expected %v, instead got %v", 8, actual)
	}
	if actual := l.Tail.Value; actual != 2 {
		t.Errorf("Expected %v, instead got %v", 2, actual)
	}
}

func TestDoublyListInsert(t *testing.T) {
	l := &DoublyLinkedList{}
	l.Insert(24, -100) // Does nothing.
	l.Insert(244, 100) // Does nothing.

	if actual := l.Length(); actual != 0 {
		t.Errorf("Expected %v, instead got %v", 0, actual)
	}

	l.Insert(2, 0)
	l.Insert(4, 1)
	l.Insert(8, 2)

	if actual := l.Length(); actual != 3 {
		t.Errorf("Expected %v, instead got %v", 3, actual)
	}
	if actual := l.Head.Value; actual != 2 {
		t.Errorf("Expected %v, instead got %v", 2, actual)
	}
	if actual := l.Tail.Value; actual != 8 {
		t.Errorf("Expected %v, instead got %v", 8, actual)
	}

	l.Insert(0, 0)

	if actual := l.Head.Value; actual != 0 {
		t.Errorf("Expected %v, instead got %v", 0, actual)
	}

	newNode := NewDNode(16)
	l.InsertNode(newNode, 4)

	if actual := l.Tail.Value; actual != 16 {
		t.Errorf("Expected %v, instead got %v", 16, actual)
	}
}

func TestDoublyListRemove(t *testing.T) {
	l := &DoublyLinkedList{}
	node2, node4, node8, node16 := NewDNode(2), NewDNode(4), NewDNode(8), NewDNode(16)
	l.AppendNode(node2)
	l.AppendNode(node4)
	l.AppendNode(node8)
	l.AppendNode(node16)

	var actualNode *DNode
	var actualValue int

	actualNode = l.Remove(4)
	if actualNode != node4 {
		t.Errorf("Expected %v, instead got %v", node4, actualNode)
	}

	actualValue = l.Length()
	if actualValue != 3 {
		t.Errorf("Expected %v, instead got %v", 3, actualValue)
	}

	actualNode = l.Find(4)
	if actualNode != nil {
		t.Errorf("Expected %v, instead got %v", nil, actualNode)
	}

	actualNode = l.RemoveAt(0)
	if actualNode.Value != 2 {
		t.Errorf("Expected %v, instead got %v", 2, actualNode)
	}

	actualValue = l.Head.Value
	if actualValue != 8 {
		t.Errorf("Expected %v, instead got %v", 8, actualValue)
	}

	actualValue = l.Tail.Value
	if actualValue != 16 {
		t.Errorf("Expected %v, instead got %v", 16, actualValue)
	}

	l.RemoveNode(node16)

	actualValue = l.Tail.Value
	if actualValue != 8 {
		t.Errorf("Expected %v, instead got %v", 8, actualValue)
	}

	l.RemoveNode(node8)

	actualNode = l.Head
	if actualNode != nil {
		t.Errorf("Expected %v, instead got %v", nil, actualNode)
	}

	actualNode = l.Tail
	if actualNode != nil {
		t.Errorf("Expected %v, instead got %v", nil, actualNode)
	}

	l.Remove(5)   // Does nothing.
	l.RemoveAt(0) // Does nothing.

	actualValue = l.Length()
	if actualValue != 0 {
		t.Errorf("Expected %v, instead got %v", 0, actualNode)
	}
}
