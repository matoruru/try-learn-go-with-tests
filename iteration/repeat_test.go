package iteration

import "fmt"
import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat charactet specified times(5)", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("repeat charactet specified times(8)", func(t *testing.T) {
		repeated := Repeat("a", 8)
		expected := "aaaaaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
}

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	// Output: aaaaa
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
