# stacks [![GoDoc](https://pkg.go.dev/badge/github.com/byExist/stacks.svg)](https://pkg.go.dev/github.com/byExist/stacks) [![Go Report Card](https://goreportcard.com/badge/github.com/byExist/stacks)](https://goreportcard.com/report/github.com/byExist/stacks)


## What is "stacks"?

stacks is a lightweight generic stack package written in Go. It wraps a slice internally to provide efficient and simple stack operations like Push and Pop. It supports generic types, fast Push/Pop operations, automatic growth when needed, clear and reuse behavior, cloning, and iteration over values using iter.Seq. The stack also supports string and JSON representations.

## Installation

To install, use the following command:

```bash
go get github.com/byExist/stacks
```

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/byExist/stacks"
)

func main() {
	s := stacks.New[int]()

	// Push elements
	stacks.Push(s, 1)
	stacks.Push(s, 2)
	stacks.Push(s, 3)

	// Peek at the top element
	v, ok := stacks.Peek(s)
	if ok {
		fmt.Println("Peek:", v)
	}

	// Pop elements
	for {
		v, ok := stacks.Pop(s)
		if !ok {
			break
		}
		fmt.Println("Pop:", v)
	}

	// Check length
	fmt.Println("Length:", stacks.Len(s))
}
```

```output
Peek: 3
Pop: 3
Pop: 2
Pop: 1
Length: 0
```

## API Overview

### Constructors

| Function                                | Description                          | Time Complexity |
|-----------------------------------------|--------------------------------------|----------------|
| `New[T]()`                              | Create a new empty stack             | O(1)           |
| `Collect[T](seq iter.Seq[T])`           | Build a stack from an iterator       | O(n)           |

### Operations

| Function                                | Description                          | Time Complexity |
|-----------------------------------------|--------------------------------------|----------------|
| `Push(s *Stack[T], item T)`             | Push an item onto the stack          | O(1)           |
| `Pop(s *Stack[T]) (T, bool)`            | Pop the top item from the stack      | O(1)           |
| `Peek(s *Stack[T]) (T, bool)`           | Peek at the top item without removing it | O(1)       |
| `Clear(s *Stack[T])`                    | Remove all items                     | O(1)           |
| `Clone(s *Stack[T]) *Stack[T]`          | Create a shallow copy                | O(n)           |

### Introspection

| Function                                | Description                          | Time Complexity |
|-----------------------------------------|--------------------------------------|----------------|
| `Len(s *Stack[T]) int`                  | Return number of items in the stack  | O(1)           |

### Iteration

| Function                                | Description                          | Time Complexity |
|-----------------------------------------|--------------------------------------|----------------|
| `Values(s *Stack[T]) iter.Seq[T]`       | Return an iterator for the stack     | O(1)           |

### Methods

| Method                                  | Description                          | Time Complexity |
|-----------------------------------------|--------------------------------------|----------------|
| `(*Stack[T]) String() string`           | Return a string representation       | O(n)           |
| `(*Stack[T]) MarshalJSON() ([]byte, error)` | Serialize the stack to JSON      | O(n)           |
| `(*Stack[T]) UnmarshalJSON([]byte) error`   | Deserialize JSON into the stack   | O(n)           |

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.