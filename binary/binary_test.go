package binary

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		haystack      []int
		needle        int
		expectedIndex int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 12, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 11, 10},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 6, 5},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 3, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 9, 8},
	}
	for i, test := range tests {
		actualIndex, _ := Search(test.needle, test.haystack)
		if actualIndex != test.expectedIndex {
			t.Errorf("%d, Error, expected %d, found %d", i, test.expectedIndex, actualIndex)
		}
	}
}
