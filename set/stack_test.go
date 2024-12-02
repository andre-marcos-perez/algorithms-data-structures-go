package set

import (
	"errors"
	"testing"
)

func TestStack(t *testing.T) {

	t.Run("create stack", func(t *testing.T) {
		stack := NewStack()
		if stack == nil {
			t.Fail()
		}
	})

	t.Run("push stack", func(t *testing.T) {
		expected := 10
		stack := NewStack()
		err := stack.Push(expected)
		if err != nil {
			t.Fail()
		}
		got, err := stack.Peek()
		if err != nil {
			t.Fail()
		}
		if got != expected {
			t.Fail()
		}
	})

	t.Run("pop stack", func(t *testing.T) {
		stack := NewStack()
		expectedOnStack := 10
		err := stack.Push(expectedOnStack)
		if err != nil {
			t.Fail()
		}
		expectedPopped := 20
		err = stack.Push(expectedPopped)
		if err != nil {
			t.Fail()
		}
		got, err := stack.Pop()
		if err != nil {
			t.Fail()
		}
		if got != expectedPopped {
			t.Fail()
		}
		got, err = stack.Peek()
		if err != nil {
			t.Fail()
		}
		if got != expectedOnStack {
			t.Fail()
		}
	})

	t.Run("stack underflow", func(t *testing.T) {
		stack := NewStack()
		_, err := stack.Pop()
		if err == nil {
			t.Fail()
		} else {
			if !errors.Is(err, ErrStackUnderflow) {
				t.Fail()
			}
		}
	})

	t.Run("stack overflow", func(t *testing.T) {
		stack := NewStack()
		var err error
		for i := range len(stack.arr) + 1 {
			err = stack.Push(i)
			if err != nil {
				break
			}
		}
		if err == nil {
			t.Fail()
		} else {
			if !errors.Is(err, ErrStackOverflow) {
				t.Fail()
			}
		}
	})

	t.Run("invert array", func(t *testing.T) {
		var v int
		var err error
		stack := NewStack()
		arrA := []int{1, 2, 3, 4, 5}
		arrB := []int{0, 0, 0, 0, 0}
		for _, v = range arrA {
			err = stack.Push(v)
			if err != nil {
				t.Fail()
			}
		}
		for i := range arrA {
			v, err = stack.Pop()
			if err != nil {
				t.Fail()
			}
			arrB[i] = v
		}
		for i := range arrA {
			if arrA[i] != arrB[len(arrB)-1-i] {
				t.Fail()
			}
		}
	})
}
