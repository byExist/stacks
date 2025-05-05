package stacks_test

import (
	"encoding/json"
	"fmt"
	"slices"
	"testing"

	"github.com/byExist/stacks"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	s := stacks.New[int]()
	require.NotNil(t, s)
	require.Equal(t, 0, stacks.Len(s))
}

func TestCollect(t *testing.T) {
	seq := slices.Values([]int{1, 2, 3, 4, 5})
	s := stacks.Collect(seq)
	require.Equal(t, 5, stacks.Len(s))

	vals := slices.Collect(stacks.Values(s))
	require.Equal(t, []int{1, 2, 3, 4, 5}, vals)
}

func TestPush(t *testing.T) {
	s := stacks.New[string]()
	stacks.Push(s, "hello")
	require.Equal(t, 1, stacks.Len(s))

	top, ok := stacks.Peek(s)
	require.True(t, ok)
	require.Equal(t, "hello", top)
}

func TestPop(t *testing.T) {
	s := stacks.New[int]()
	stacks.Push(s, 42)
	stacks.Push(s, 100)

	val, ok := stacks.Pop(s)
	require.True(t, ok)
	require.Equal(t, 100, val)
	require.Equal(t, 1, stacks.Len(s))

	val, ok = stacks.Pop(s)
	require.True(t, ok)
	require.Equal(t, 42, val)
	require.Equal(t, 0, stacks.Len(s))

	_, ok = stacks.Pop(s)
	require.False(t, ok)
}

func TestPeek(t *testing.T) {
	s := stacks.New[int]()
	_, ok := stacks.Peek(s)
	require.False(t, ok)

	stacks.Push(s, 10)
	stacks.Push(s, 20)

	val, ok := stacks.Peek(s)
	require.True(t, ok)
	require.Equal(t, 20, val)
	require.Equal(t, 2, stacks.Len(s)) // Peek does not change length
}

func TestClone(t *testing.T) {
	s := stacks.New[string]()
	stacks.Push(s, "a")
	stacks.Push(s, "b")

	clone := stacks.Clone(s)
	require.Equal(t, stacks.Len(s), stacks.Len(clone))

	originalVals := []string{}
	for v := range stacks.Values(s) {
		originalVals = append(originalVals, v)
	}

	cloneVals := []string{}
	for v := range stacks.Values(clone) {
		cloneVals = append(cloneVals, v)
	}

	require.Equal(t, originalVals, cloneVals)

	// Modify original and ensure clone is unaffected
	stacks.Push(s, "c")
	require.NotEqual(t, stacks.Len(s), stacks.Len(clone))
}

func TestLen(t *testing.T) {
	s := stacks.New[int]()
	require.Equal(t, 0, stacks.Len(s))

	stacks.Push(s, 1)
	require.Equal(t, 1, stacks.Len(s))

	stacks.Push(s, 2)
	require.Equal(t, 2, stacks.Len(s))

	_, _ = stacks.Pop(s)
	require.Equal(t, 1, stacks.Len(s))
}

func TestValues(t *testing.T) {
	s := stacks.New[int]()
	for i := 1; i <= 3; i++ {
		stacks.Push(s, i)
	}

	vals := []int{}
	for v := range stacks.Values(s) {
		vals = append(vals, v)
	}
	require.Equal(t, []int{1, 2, 3}, vals)
}

func TestClear(t *testing.T) {
	s := stacks.New[int]()
	for i := 0; i < 5; i++ {
		stacks.Push(s, i)
	}
	require.Equal(t, 5, stacks.Len(s))

	stacks.Clear(s)
	require.Equal(t, 0, stacks.Len(s))

	// After clear, stack should be reusable
	stacks.Push(s, 100)
	require.Equal(t, 1, stacks.Len(s))

	val, ok := stacks.Peek(s)
	require.True(t, ok)
	require.Equal(t, 100, val)
}

func TestStackWithDifferentTypes(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		s := stacks.New[int]()
		stacks.Push(s, 10)
		stacks.Push(s, 20)

		val, ok := stacks.Pop(s)
		require.True(t, ok)
		require.Equal(t, 20, val)

		val, ok = stacks.Pop(s)
		require.True(t, ok)
		require.Equal(t, 10, val)
	})

	t.Run("string", func(t *testing.T) {
		s := stacks.New[string]()
		stacks.Push(s, "foo")
		stacks.Push(s, "bar")

		val, ok := stacks.Pop(s)
		require.True(t, ok)
		require.Equal(t, "bar", val)

		val, ok = stacks.Pop(s)
		require.True(t, ok)
		require.Equal(t, "foo", val)
	})

	t.Run("struct", func(t *testing.T) {
		type Point struct {
			X, Y int
		}

		s := stacks.New[Point]()
		stacks.Push(s, Point{X: 1, Y: 2})
		stacks.Push(s, Point{X: 3, Y: 4})

		val, ok := stacks.Pop(s)
		require.True(t, ok)
		require.Equal(t, Point{X: 3, Y: 4}, val)

		val, ok = stacks.Pop(s)
		require.True(t, ok)
		require.Equal(t, Point{X: 1, Y: 2}, val)
	})
}

func TestStackString(t *testing.T) {
	s := stacks.New[int]()
	stacks.Push(s, 1)
	stacks.Push(s, 2)
	require.Contains(t, s.String(), "Stack{1, 2}")
}

func TestStackMarshalUnmarshalJSON(t *testing.T) {
	s := stacks.New[int]()
	stacks.Push(s, 10)
	stacks.Push(s, 20)

	data, err := json.Marshal(s)
	require.NoError(t, err)

	var s2 stacks.Stack[int]
	err = json.Unmarshal(data, &s2)
	require.NoError(t, err)

	vals1 := slices.Collect(stacks.Values(s))
	vals2 := slices.Collect(stacks.Values(&s2))
	require.Equal(t, vals1, vals2)
}

func ExampleNew() {
	s := stacks.New[int]()
	stacks.Push(s, 1)
	stacks.Push(s, 2)

	for v := range stacks.Values(s) {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
}

func ExampleCollect() {
	seq := slices.Values([]int{10, 20, 30})
	s := stacks.Collect(seq)

	for v := range stacks.Values(s) {
		fmt.Println(v)
	}
	// Output:
	// 10
	// 20
	// 30
}

func ExampleClone() {
	s := stacks.New[string]()
	stacks.Push(s, "x")
	stacks.Push(s, "y")

	clone := stacks.Clone(s)
	stacks.Push(s, "z") // Modify original

	for v := range stacks.Values(clone) {
		fmt.Println(v)
	}
	// Output:
	// x
	// y
}

func ExamplePush() {
	s := stacks.New[string]()
	stacks.Push(s, "alpha")
	stacks.Push(s, "beta")

	for v := range stacks.Values(s) {
		fmt.Println(v)
	}
	// Output:
	// alpha
	// beta
}

func ExamplePop() {
	s := stacks.New[int]()
	stacks.Push(s, 1)
	stacks.Push(s, 2)

	val, ok := stacks.Pop(s)
	if ok {
		fmt.Println(val)
	}
	// Output:
	// 2
}

func ExamplePeek() {
	s := stacks.New[int]()
	stacks.Push(s, 42)

	val, ok := stacks.Peek(s)
	if ok {
		fmt.Println(val)
	}
	// Output:
	// 42
}

func ExampleLen() {
	s := stacks.New[int]()
	fmt.Println(stacks.Len(s))

	stacks.Push(s, 1)
	stacks.Push(s, 2)
	fmt.Println(stacks.Len(s))

	_, _ = stacks.Pop(s)
	fmt.Println(stacks.Len(s))
	// Output:
	// 0
	// 2
	// 1
}

func ExampleValues() {
	s := stacks.New[int]()
	stacks.Push(s, 5)
	stacks.Push(s, 10)

	for v := range stacks.Values(s) {
		fmt.Println(v)
	}
	// Output:
	// 5
	// 10
}

func ExampleClear() {
	s := stacks.New[int]()
	stacks.Push(s, 123)
	stacks.Push(s, 456)
	fmt.Println("Before clear:", stacks.Len(s))

	stacks.Clear(s)
	fmt.Println("After clear:", stacks.Len(s))
	// Output:
	// Before clear: 2
	// After clear: 0
}

func ExampleStack_String() {
	s := stacks.New[int]()
	stacks.Push(s, 1)
	stacks.Push(s, 2)
	fmt.Println(s)
	// Output:
	// Stack{1, 2}
}

func ExampleStack_MarshalJSON() {
	s := stacks.New[int]()
	stacks.Push(s, 1)
	stacks.Push(s, 2)
	b, _ := json.Marshal(s)
	fmt.Println(string(b))
	// Output:
	// [1,2]
}

func ExampleStack_UnmarshalJSON() {
	var s stacks.Stack[int]
	_ = json.Unmarshal([]byte(`[5,10]`), &s)
	for v := range stacks.Values(&s) {
		fmt.Println(v)
	}
	// Output:
	// 5
	// 10
}
