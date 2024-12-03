package set

import (
	"errors"
	"testing"
)

func TestList(t *testing.T) {

	t.Run("create list", func(t *testing.T) {
		list := NewList()
		if list == nil {
			t.Fail()
		}
	})

	t.Run("insert list", func(t *testing.T) {
		list := NewList()
		arr := []int{9, 16, 4, 1}
		for _, v := range arr {
			expected := v
			list.Insert(expected)
			got, err := list.Peek()
			if err != nil {
				t.Fail()
			}
			if got.k != expected {
				t.Fail()
			}
		}
		i := 0
		for e, _ := list.Peek(); e != nil; e = e.next {
			if e.k != arr[len(arr)-1-i] {
				t.Fail()
			}
			i += 1
		}
	})

	t.Run("find list", func(t *testing.T) {
		list := NewList()
		arr := []int{9, 16, 4, 1}
		for _, v := range arr {
			expected := v
			list.Insert(expected)
			got, err := list.Peek()
			if err != nil {
				t.Fail()
			}
			if got.k != expected {
				t.Fail()
			}
		}
		expected := 4
		got, err := list.Find(expected)
		if err != nil {
			t.Fail()
		}
		if got.k != expected {
			t.Fail()
		}
		expected = 10
		_, err = list.Find(expected)
		if !errors.Is(err, ErrListNotFound) {
			t.Fail()
		}
	})

	t.Run("delete list", func(t *testing.T) {
		list := NewList()
		arr := []int{9, 16, 4, 1}
		for _, v := range arr {
			expected := v
			list.Insert(expected)
			got, err := list.Peek()
			if err != nil {
				t.Fail()
			}
			if got.k != expected {
				t.Fail()
			}
		}
		expected := 4
		got, err := list.Delete(expected)
		if err != nil {
			t.Fail()
		}
		if got.k != expected {
			t.Fail()
		}
		_, err = list.Find(expected)
		if !errors.Is(err, ErrListNotFound) {
			t.Fail()
		}
	})
}
