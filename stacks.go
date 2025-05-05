package stacks

import (
	"encoding/json"
	"fmt"
	"iter"
	"slices"
	"strings"
)

// Stack is a generic, dynamically resizing stack.
type Stack[T any] struct {
	items []T
}

// String returns a string representation of the stack.
func (s *Stack[T]) String() string {
	var b strings.Builder
	b.WriteString("Stack{")
	for i, v := range s.items {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprint(v))
	}
	b.WriteString("}")
	return b.String()
}

// MarshalJSON implements json.Marshaler for Stack.
func (s *Stack[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.items)
}

// UnmarshalJSON implements json.Unmarshaler for Stack.
func (s *Stack[T]) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &s.items); err != nil {
		return err
	}
	return nil
}

// New creates a new empty stack.
func New[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

// Collect builds a stack from a given sequence of elements.
func Collect[T any](i iter.Seq[T]) *Stack[T] {
	s := New[T]()
	for e := range i {
		Push(s, e)
	}
	return s
}

// Clone creates a shallow copy of the stack.
func Clone[T any](s *Stack[T]) *Stack[T] {
	return &Stack[T]{
		items: slices.Clone(s.items),
	}
}

// Peek returns the top element without removing it. Returns false if the stack is empty.
func Peek[T any](s *Stack[T]) (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// Push adds an element to the top of the stack.
func Push[T any](s *Stack[T], item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element of the stack. Returns false if the stack is empty.
func Pop[T any](s *Stack[T]) (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	idx := len(s.items) - 1
	item := s.items[idx]
	s.items = s.items[:idx]
	return item, true
}

// Len returns the number of elements in the stack.
func Len[T any](s *Stack[T]) int {
	return len(s.items)
}

// Values returns a sequence that yields all elements in the stack from bottom to top.
func Values[T any](s *Stack[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range s.items {
			if !yield(v) {
				break
			}
		}
	}
}

// Clear removes all elements but keeps the allocated capacity.
func Clear[T any](s *Stack[T]) {
	s.items = make([]T, 0, cap(s.items))
}
