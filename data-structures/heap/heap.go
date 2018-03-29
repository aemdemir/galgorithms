package heap

import "fmt"

// Heap is a tree, and all of the nodes in the tree have 0, 1 or 2 children.
// Elements in a heap are partially sorted by their priority.
// Every node in the tree has a higher priority than its children.
// Heap's priority can be either maxheap or minheap.
type Heap struct {
	elements     []interface{}
	priorityFunc func(interface{}, interface{}) bool
}

// NewHeap returns a new heap prioritized with given function.
func NewHeap(elements []interface{}, priorityFunc func(interface{}, interface{}) bool) *Heap {
	h := &Heap{elements, priorityFunc}
	h.BuildHeap()
	return h
}

// BuildHeap converts an array into heap.
func (h *Heap) BuildHeap() {
	for i := (h.Count() / 2) - 1; i >= 0; i-- {
		h.HeapifyDown(i)
	}
}

// Highest returns the first element in the heap(element with the highest priority).
func (h *Heap) Highest() interface{} {
	if h.IsEmpty() {
		return nil
	}
	return h.elements[0]
}

// Insert inserts a new element into heap.
func (h *Heap) Insert(element interface{}) {
	h.elements = append(h.elements, element)
	h.HeapifyUp(h.Count() - 1)
}

// HeapifyUp sifts element up if it's priority higher than it's parent's.
func (h *Heap) HeapifyUp(index int) {
	if h.IsRoot(index) {
		return
	}

	parent := h.Parent(index)

	if h.IsHigherPriority(index, parent) {
		h.Swap(index, parent)
		h.HeapifyUp(parent)
	}
}

// Extract extracts the highest priority element.
func (h *Heap) Extract() interface{} {
	if h.IsEmpty() {
		return nil
	}

	element := h.elements[0]
	lastIndex := h.Count() - 1

	h.Swap(0, lastIndex)
	h.elements[lastIndex] = nil
	h.elements = h.elements[:lastIndex]

	if !h.IsEmpty() {
		h.HeapifyDown(0)
	}

	return element
}

// HeapifyDown sifts element down if it's priority lower than it's children's.
func (h *Heap) HeapifyDown(index int) {
	highest := h.HighestPriorityIndexFor(index)

	if index == highest {
		return
	}

	h.Swap(index, highest)
	h.HeapifyDown(highest)
}

// String prints heap in a nice format.
func (h *Heap) String() string {
	if h.Count() == 0 {
		return "[]"
	}

	return fmt.Sprint(h.elements)
}

/*
	Helper Functions
*/

// Count returns how many elements are currently in the heap.
func (h *Heap) Count() int {
	return len(h.elements)
}

// IsEmpty tells whether the heap is empty or not.
func (h *Heap) IsEmpty() bool {
	return h.Count() == 0
}

// IsRoot tells whether the given index is root index or not.
func (h *Heap) IsRoot(index int) bool {
	return index == 0
}

// Left returns left child index of i.
func (h *Heap) Left(i int) int {
	return (2 * i) + 1
}

// Right returns right child index of i.
func (h *Heap) Right(i int) int {
	return (2 * i) + 2
}

// Parent returns parent index of i.
func (h *Heap) Parent(i int) int {
	return (i - 1) / 2
}

// Swap swaps element at firstIndex and secondIndex.
func (h *Heap) Swap(firstIndex, secondIndex int) {
	if firstIndex == secondIndex {
		return
	}
	h.elements[firstIndex], h.elements[secondIndex] = h.elements[secondIndex], h.elements[firstIndex]
}

// IsHigherPriority returns true if the element at firstIndex has
// higher priority than the element at second index.
func (h *Heap) IsHigherPriority(firstIndex, secondIndex int) bool {
	return h.priorityFunc(h.elements[firstIndex], h.elements[secondIndex])
}

// HighestPriorityIndex returns the index which has highest priority
// among given parent and child indices.
func (h *Heap) HighestPriorityIndex(parentIndex, childIndex int) int {
	if childIndex < h.Count() && h.IsHigherPriority(childIndex, parentIndex) {
		return childIndex
	}
	return parentIndex
}

// HighestPriorityIndexFor returns the index with highest priority
// among parent, left and right child indices.
func (h *Heap) HighestPriorityIndexFor(parentIndex int) int {
	return h.HighestPriorityIndex(h.HighestPriorityIndex(parentIndex, h.Left(parentIndex)), h.Right(parentIndex))
}
