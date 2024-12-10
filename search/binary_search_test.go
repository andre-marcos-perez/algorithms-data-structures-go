package search

import "testing"

func TestBinarySearch(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 8, 9, 10}

	t.Run("found", func(t *testing.T) {
		for _, v := range []int{3, 1, 10} {
			found := BinarySearch(v, &arr)
			if found == false {
				t.Fail()
			}
		}
	})

	t.Run("not found", func(t *testing.T) {
		for _, v := range []int{7, -1, 11} {
			found := BinarySearch(v, &arr)
			if found == true {
				t.Fail()
			}
		}
	})
}
