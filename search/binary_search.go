package search

func BinarySearch(k int, sequence *[]int) bool {

	var (
		head   int
		middle int
		tail   int
	)

	head = 0
	tail = len(*sequence) - 1

	for {

		if head > tail {
			break
		}

		if (head+tail)%2 == 0 {
			middle = (head + tail) / 2
		} else {
			middle = (((head + tail) - 1) / 2) + 1
		}

		if k > (*sequence)[middle] {
			head = middle + 1
		} else if k < (*sequence)[middle] {
			tail = middle - 1
		} else {
			return true
		}

	}

	return false
}
