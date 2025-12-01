package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	testCases := []struct {
		numbers  []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{5, 5, 5, 5, 5}, 25},
		{[]int{1, 2, 3}, 6},
		{[]int{}, 0},
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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
