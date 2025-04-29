# stacks [![GoDoc](https://pkg.go.dev/badge/github.com/byExist/stacks.svg)](https://pkg.go.dev/github.com/byExist/stacks) [![Go Report Card](https://goreportcard.com/badge/github.com/byExist/stacks)](https://goreportcard.com/report/github.com/byExist/stacks)


## What is "stacks"?

stacks is a lightweight generic stack package written in Go. It wraps a slice internally to provide efficient and simple stack operations like Push and Pop.

## Features

- Supports generic types
- Fast Push/Pop operations
- Automatically grows capacity when needed
- Can be cleared and reused
- Provides a stack clone function
- Supports iteration over values using iter.Seq

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

## Usage

The stacks package makes it easy to create and manage lightweight stacks. You can quickly perform operations like adding (Push), removing (Pop), peeking at the top element (Peek), and iterating over all elements (Values).

Internally, it wraps a slice to provide simple and efficient stack behavior.

## API Overview

### Constructors

- `New[T]() *Stack[T]`
- `Collect[T](seq iter.Seq[T]) *Stack[T]`

### Core Methods

- `Push(s *Stack[T], item T)`
- `Pop(s *Stack[T]) (T, bool)`
- `Peek(s *Stack[T]) (T, bool)`
- `Clear(s *Stack[T])`
- `Clone(s *Stack[T]) *Stack[T]`
- `Values(s *Stack[T]) iter.Seq[T]`
- `Len(s *Stack[T]) int`

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.