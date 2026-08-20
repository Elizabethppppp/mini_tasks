package functions

import "testing"

func TestFib(t *testing.T) {
	fib := Fib()

	t.Run("базовые случаи", func(t *testing.T) {
		tests := []struct {
			n    int
			want int
		}{
			{0, 0},
			{1, 1},
		}

		for _, tt := range tests {
			got := fib(tt.n)
			if got != tt.want {
				t.Errorf("fib(%d) = %d, want %d", tt.n, got, tt.want)
			}
		}
	})

	t.Run("первые числа Фибоначчи", func(t *testing.T) {
		expected := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

		for i, want := range expected {
			got := fib(i)
			if got != want {
				t.Errorf("fib(%d) = %d, want %d", i, got, want)
			}
		}
	})

	t.Run("Fib(40)", func(t *testing.T) {
		want := 102334155
		got := fib(40)
		if got != want {
			t.Errorf("fib(40) = %d, want %d", got, want)
		}
	})
}

func BenchmarkFib40(b *testing.B) {
	fib := Fib()

	fib(40)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib(40)
	}
}
