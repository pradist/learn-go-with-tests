package arrays

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	testCases := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{"given array of five numbers when sum then returns correct total", []int{1, 2, 3, 4, 5}, 15},
		{"given array of same numbers when sum then returns correct total", []int{5, 5, 5, 5, 5}, 25},
		{"given array of three numbers when sum then returns correct total", []int{1, 2, 3}, 6},
		{"given empty array when sum then returns zero", []int{}, 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := Sum(testCase.numbers)
			if got != testCase.expected {
				t.Errorf("got %d want %d given, %v", got, testCase.expected, testCase.numbers)
			}
		})
	}
}

func TestSumAll(t *testing.T) {
	testCases := []struct {
		name         string
		numbersToSum [][]int
		expected     []int
	}{
		{"given multiple slices when sum all then returns correct totals", [][]int{{1, 2}, {3, 4, 5}}, []int{3, 12}},
		{"given empty slice when sum all then returns zero", [][]int{{}, {0, 9}}, []int{0, 9}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := SumAll(testCase.numbersToSum...)
			if !slices.Equal(got, testCase.expected) {
				t.Errorf("got %v want %v given, %v", got, testCase.expected, testCase.numbersToSum)
			}
		})
	}
}

func TestSumAllTails(t *testing.T) {
	testCases := []struct {
		name         string
		numbersToSum [][]int
		expected     []int
	}{
		{"make the sums of tails of", [][]int{{1, 2}, {0, 9}}, []int{2, 9}},
		{"safely sum empty slices", [][]int{{}, {3, 4, 5}}, []int{0, 9}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := SumAllTails(testCase.numbersToSum...)
			if !reflect.DeepEqual(got, testCase.expected) {
				t.Errorf("got %v want %v given, %v", got, testCase.expected, testCase.numbersToSum)
			}
		})
	}
}
