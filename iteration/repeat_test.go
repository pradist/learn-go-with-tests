package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("print a five times", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("print b five times", func(t *testing.T) {
		repeated := Repeat("b")
		expected := "bbbbb"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("print c five times", func(t *testing.T) {
		repeated := Repeat("c")
		expected := "ccccc"

		assertCorrectMessage(t, repeated, expected)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
