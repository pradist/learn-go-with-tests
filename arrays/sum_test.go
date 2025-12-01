package arrays

import "testing"

func TestSum(t *testing.T) {

	testCases := []struct {
		numbers  [5]int
		expected int
	}{
		{[5]int{1, 2, 3, 4, 5}, 15},
		{[5]int{5, 5, 5, 5, 5}, 25},
	}

	for _, testCase := range testCases {
		t.Run("given array of five numbers when sum then returns correct total", func(t *testing.T) {
			got := Sum(testCase.numbers)
			if got != testCase.expected {
				t.Errorf("got %d want %d given, %v", got, testCase.expected, testCase.numbers)
			}
		})
	}

}
