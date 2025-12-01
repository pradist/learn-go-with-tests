package iteration

import "testing"

func TestRepeat(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"repeat a five times", "a", "aaaaa"},
		{"repeat b five times", "b", "bbbbb"},
		{"repeat c five times", "c", "ccccc"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
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

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
