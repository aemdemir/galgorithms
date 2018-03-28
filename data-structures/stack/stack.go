package stack

import (
	"fmt"
)

// Stack is a data structure similar to arrays with limited functionalities.
// It makes a pLIFO(last in, first out) order.
type Stack struct {
	slice []interface{}
}

// New creates and returns a new Stack.
func New() *Stack {
	s := &Stack{}
	s.slice = make([]interface{}, 0)
	return s
}

// String prints stack in a nice format.
func (s *Stack) String() string {
	begin := fmt.Sprintf("%s\n", "----Stack----")
	elements := ""
	for _, element := range s.slice {
		elements += fmt.Sprintf("%v\n", element)
	}
	end := fmt.Sprintf("%s\n", "-------------")
	return begin + elements + end
}

// Push adds element onto the stack.
func (s *Stack) Push(element interface{}) {
	s.slice = append(s.slice, element)
}

// Pop removes the top element from the stack.
func (s *Stack) Pop() interface{} {
	if len(s.slice) == 0 {
		return nil
	}

	element := s.slice[len(s.slice)-1]

	s.slice[len(s.slice)-1] = nil
	s.slice = s.slice[:len(s.slice)-1]

	return element
}

// Peek returns the top element of the stack without removing.
func (s *Stack) Peek() interface{} {
	if len(s.slice) == 0 {
		return nil
	}
	return s.slice[len(s.slice)-1]
}

// Length returns how many elements are currently in the stack.
func (s *Stack) Length() int {
	return len(s.slice)
}

// IsEmpty tells wheter the stack is empty or not.
func (s *Stack) IsEmpty() bool {
	return len(s.slice) == 0
}

// ToSlice converts stack to a slice.
func (s *Stack) ToSlice() []interface{} {
	copySlice := make([]interface{}, len(s.slice))
	copy(copySlice, s.slice)
	return copySlice
}