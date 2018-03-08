package queue

import "fmt"

// Queue is a data structure similar to arrays.
// It can only insert new items at back and can only remove items from the front.
type Queue struct {
	slice []interface{}
}

// New creates and returns a new Queue.
func New() *Queue {
	q := &Queue{}
	q.slice = make([]interface{}, 0)
	return q
}

// String prints queue in a nice format.
func (q *Queue) String() string {
	begin := fmt.Sprintf("%s\n", "----Queue----")
	elements := ""
	for _, element := range q.slice {
		elements += fmt.Sprintf("%v\n", element)
	}
	end := fmt.Sprintf("%s\n", "-------------")
	return begin + elements + end
}

// Enqueue appends element.
func (q *Queue) Enqueue(element interface{}) {
	q.slice = append(q.slice, element)
}

// Dequeue removes the first element from the queue.
func (q *Queue) Dequeue() interface{} {
	if len(q.slice) == 0 {
		return nil
	}

	element := q.slice[0]

	q.slice[0] = nil
	q.slice = q.slice[1:]

	return element
}

// Peek returns the element at the beginning of the queue without removing it.
func (q *Queue) Peek() interface{} {
	if len(q.slice) == 0 {
		return nil
	}
	return q.slice[0]
}

// Length returns how many elements are currently in the queue.
func (q *Queue) Length() int {
	return len(q.slice)
}

// IsEmpty tells wheter the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(q.slice) == 0
}
