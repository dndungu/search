package binary

// Search takes a needle and haystack and returns the index of the needle in the haystack
// and true or -1 and false when the needle does not exist in the haystack
func Search(needle int, haystack []int) (int, bool) {
	low := 0
	high := len(haystack) - 1
	for {
		if low > high {
			return -1, false
		}
		middle := int((low + high) / 2)
		if needle == haystack[middle] {
			return middle, true
		}
		if needle < haystack[middle] {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
}
