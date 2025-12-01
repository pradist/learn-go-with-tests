package arrays

import (
	"reflect"
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
	testcases := []struct {
		numbersToSum [][]int
		expected     []int
	}{
		{[][]int{{1, 2}, {3, 4, 5}}, []int{3, 12}},
		{[][]int{{}, {0, 9}}, []int{0, 9}},
	}

	for _, testtestcases := range testcases {
		t.Run("given multiple slices when sum all then returns correct totals", func(t *testing.T) {
			got := SumAll(testtestcases.numbersToSum...)
			if !slices.Equal(got, testtestcases.expected) {
				t.Errorf("got %v want %v given, %v", got, testtestcases.expected, testtestcases.numbersToSum)
			}
		})
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
