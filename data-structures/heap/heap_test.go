package heap

import (
	"testing"
)

func isEqual(sliceA, sliceB []interface{}) bool {
	if sliceA == nil && sliceB == nil {
		return true
	}

	if sliceA == nil || sliceB == nil {
		return false
	}

	if len(sliceA) != len(sliceB) {
		return false
	}

	for i := range sliceA {
		if sliceA[i] != sliceB[i] {
			return false
		}
	}

	return true
}

func TestNewHeap(t *testing.T) {
	maxPriorityFunc := func(lhs interface{}, rhs interface{}) bool {
		return lhs.(int) > rhs.(int)
	}
	minPriorityFunc := func(lhs interface{}, rhs interface{}) bool {
		return lhs.(int) < rhs.(int)
	}

	emptyElements := []interface{}{}
	maxEmptyHeap := NewHeap(emptyElements, maxPriorityFunc)
	minEmptyHeap := NewHeap(emptyElements, minPriorityFunc)

	if actual := maxEmptyHeap.IsEmpty(); actual != true {
		t.Errorf("Expected %v, instead got %v", true, actual)
	}
	if actual := minEmptyHeap.Count(); actual != 0 {
		t.Errorf("Expected %v, instead got %v", 0, actual)
	}

	elements1 := []interface{}{1}
	expected := []interface{}{1}
	h1 := NewHeap(elements1, maxPriorityFunc)

	if actual := isEqual(h1.elements, expected); actual != true {
		t.Errorf("Expected %v, instead got %v", expected, actual)
	}

	elements2 := []interface{}{2, 1}
	expected = []interface{}{1, 2}
	h2 := NewHeap(elements2, minPriorityFunc)

	if actual := isEqual(h2.elements, expected); actual != true {
		t.Errorf("Expected %v, instead got %v", expected, actual)
	}

	elements3 := []interface{}{3, 2, 8, 5, 0}
	expected = []interface{}{8, 5, 3, 2, 0}
	h3 := NewHeap(elements3, maxPriorityFunc)

	if actual := isEqual(h3.elements, expected); actual != true {
		t.Errorf("Expected %v, instead got %v", expected, actual)
	}

	elements4 := []interface{}{3, 2, 8, 5, 0}
	expected = []interface{}{0, 2, 8, 5, 3}
	h4 := NewHeap(elements4, minPriorityFunc)

	if actual := isEqual(h4.elements, expected); actual != true {
		t.Errorf("Expected %v, instead got %v", expected, actual)
	}
}

func TestHeapInsert(t *testing.T) {
	maxPriorityFunc := func(lhs interface{}, rhs interface{}) bool {
		return lhs.(float64) > rhs.(float64)
	}

	h1 := NewHeap([]interface{}{}, maxPriorityFunc)
	h1.Insert(7.0)
	expected := []interface{}{7.0}

	if actual := isEqual(h1.elements, expected); actual != true {
		t.Errorf("Expected %v, instead got %v", expected, h1.elements)
	}

	h1.Insert(2.0)
	expected = []interface{}{7.0, 2.0}

	if actual := isEqual(h1.elements, expected); actual != true {
		t.Errorf("Expected %v, instead got %v", expected, h1.elements)
	}

	h1.Insert(9.0)
	expected = []interface{}{9.0, 2.0, 7.0}

	if actual := isEqual(h1.elements, expected); actual != true {
		t.Errorf("Expected %v, instead got %v", expected, h1.elements)
	}

	h1.Insert(1.0)
	expected = []interface{}{9.0, 2.0, 7.0, 1.0}

	if actual := isEqual(h1.elements, expected); actual != true {
		t.Errorf("Expected %v, instead got %v", expected, h1.elements)
	}

	h2 := NewHeap([]interface{}{9.0, 2.0, 7.0, 1.0}, maxPriorityFunc)
	h2.Insert(20.0)

	if actual := h2.Highest(); actual != 20.0 {
		t.Errorf("Expected %v, instead got %v", 20.0, actual)
	}
}

func TestHeapExtract(t *testing.T) {
	maxPriorityFunc := func(lhs interface{}, rhs interface{}) bool {
		return lhs.(string) > rhs.(string)
	}

	h := NewHeap([]interface{}{"p", "c", "k", "a"}, maxPriorityFunc)

	if actual := h.Extract(); actual != "p" {
		t.Errorf("Expected %v, instead got %v", "p", actual)
	}

	expected := []interface{}{"k", "c", "a"}
	if actual := isEqual(h.elements, expected); actual != true {
		t.Errorf("Expected %v, instead got %v", expected, actual)
	}

	h.Extract()
	expected = []interface{}{"c", "a"}
	if actual := isEqual(h.elements, expected); actual != true {
		t.Errorf("Expected %v, instead got %v", expected, actual)
	}

	h.Extract()
	expected = []interface{}{"a"}
	if actual := isEqual(h.elements, expected); actual != true {
		t.Errorf("Expected %v, instead got %v", expected, actual)
	}

	h.Extract()
	expected = []interface{}{}
	if actual := isEqual(h.elements, expected); actual != true {
		t.Errorf("Expected %v, instead got %v", expected, actual)
	}

	h.Extract()
	expected = []interface{}{}
	if actual := isEqual(h.elements, expected); actual != true {
		t.Errorf("Expected %v, instead got %v", expected, actual)
	}

	if actual := h.Extract(); actual != nil {
		t.Errorf("Expected %v, instead got %v", nil, actual)
	}
}
