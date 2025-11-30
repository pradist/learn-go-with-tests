package iteration

import "testing"

func TestRepeat(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"a", "aaaaa"},
		{"b", "bbbbb"},
		{"c", "ccccc"},
	}

	for _, testCase := range testCases {
		t.Run("repeating "+testCase.input, func(t *testing.T) {
			repeated := Repeat(testCase.input)
			assertCorrectMessage(t, repeated, testCase.expected)
		})
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
