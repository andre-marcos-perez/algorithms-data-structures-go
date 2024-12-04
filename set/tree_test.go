package set

import (
	"errors"
	"testing"
)

func TestTree(t *testing.T) {
	t.Parallel()

	t.Run("create", func(t *testing.T) {
		tree := NewTree()
		if tree == nil {
			t.Fail()
		}
	})

	t.Run("insert", func(t *testing.T) {
		tree := NewTree()
		arr := []int{3, 2, 1, 7, 6, 5, 4, 8}
		for _, v := range arr {
			tree.Insert(v)
		}
	})

	t.Run("find", func(t *testing.T) {
		tree := NewTree()
		arr := []int{3, 2, 1, 7, 6, 5, 4, 8}
		for _, v := range arr {
			tree.Insert(v)
		}
		for _, v := range arr {
			expected := v
			got, err := tree.Find(tree.root, expected)
			if err != nil {
				t.Fail()
			}
			if got.k != expected {
				t.Fail()
			}
		}
		expected := 10
		_, err := tree.Find(tree.root, expected)
		if !errors.Is(err, ErrTreeNotFound) {
			t.Fail()
		}
	})

	t.Run("min", func(t *testing.T) {
		tree := NewTree()
		_, err := tree.Min(tree.root)
		if !errors.Is(err, ErrTreeEmpty) {
			t.Fail()
		}
		arr := []int{3, 2, 1, 7, 6, 5, 4, 8}
		for _, v := range arr {
			tree.Insert(v)
		}
		expected := 1
		got, err := tree.Min(tree.root)
		if err != nil {
			t.Fail()
		}
		if got.k != expected {
			t.Fail()
		}
	})

	t.Run("max", func(t *testing.T) {
		tree := NewTree()
		_, err := tree.Min(tree.root)
		if !errors.Is(err, ErrTreeEmpty) {
			t.Fail()
		}
		arr := []int{3, 2, 1, 7, 6, 5, 4, 8}
		for _, v := range arr {
			tree.Insert(v)
		}
		expected := 8
		got, err := tree.Max(tree.root)
		if err != nil {
			t.Fail()
		}
		if got.k != expected {
			t.Fail()
		}
	})

}
