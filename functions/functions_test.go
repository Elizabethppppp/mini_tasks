package functions

import (
	"reflect"
	"testing"
)

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

func TestMapFilterReduce(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	isEven := func(x int) bool { return x%2 == 0 }
	square := func(x int) int { return x * x }
	sum := func(acc, x int) int { return acc + x }

	t.Run("по отдельности", func(t *testing.T) {
		filtered := Filter(numbers, isEven)
		expectedFiltered := []int{2, 4, 6, 8, 10}
		if !reflect.DeepEqual(filtered, expectedFiltered) {
			t.Errorf("Filter() = %v, want %v", filtered, expectedFiltered)
		}

		mapped := Map(numbers, square)
		expectedMapped := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
		if !reflect.DeepEqual(mapped, expectedMapped) {
			t.Errorf("Map() = %v, want %v", mapped, expectedMapped)
		}

		reduced := Reduce(numbers, 0, sum)
		expectedReduced := 55
		if reduced != expectedReduced {
			t.Errorf("Reduce() = %d, want %d", reduced, expectedReduced)
		}
	})

	t.Run("цепочка: сумма квадратов четных = 220", func(t *testing.T) {

		result := Reduce(
			Map(
				Filter(numbers, isEven),
				square,
			),
			0,
			sum,
		)

		expected := 220
		if result != expected {
			t.Errorf("Цепочка: сумма квадратов четных = %d, want %d", result, expected)
		}
	})

	t.Run("цепочка: сумма квадратов нечетных = 165", func(t *testing.T) {

		isOdd := func(x int) bool { return x%2 != 0 }

		result := Reduce(
			Map(
				Filter(numbers, isOdd),
				square,
			),
			0,
			sum,
		)

		expected := 165
		if result != expected {
			t.Errorf("Цепочка: сумма квадратов нечетных = %d, want %d", result, expected)
		}
	})

	t.Run("цепочка: произведение четных чисел", func(t *testing.T) {

		multiply := func(acc, x int) int { return acc * x }

		result := Reduce(
			Filter(numbers, isEven),
			1,
			multiply,
		)

		expected := 3840
		if result != expected {
			t.Errorf("Произведение четных чисел = %d, want %d", result, expected)
		}
	})

	t.Run("проверка что функции не изменяют исходный слайс", func(t *testing.T) {
		original := []int{1, 2, 3, 4, 5}
		originalCopy := make([]int, len(original))
		copy(originalCopy, original)

		_ = Filter(original, isEven)
		_ = Map(original, square)
		_ = Reduce(original, 0, sum)

		if !reflect.DeepEqual(original, originalCopy) {
			t.Errorf("Исходный слайс изменился! Было: %v, Стало: %v", originalCopy, original)
		}
	})
}

func TestMultiplier(t *testing.T) {
	t.Run("умножение на 2", func(t *testing.T) {
		double := Multiplier(2)

		tests := []struct {
			x    int
			want int
		}{
			{0, 0},
			{1, 2},
			{5, 10},
			{10, 20},
			{-3, -6},
		}

		for _, tt := range tests {
			got := double(tt.x)
			if got != tt.want {
				t.Errorf("double(%d) = %d, want %d", tt.x, got, tt.want)
			}
		}
	})

	t.Run("умножение на 3", func(t *testing.T) {
		triple := Multiplier(3)

		tests := []struct {
			x    int
			want int
		}{
			{0, 0},
			{1, 3},
			{5, 15},
			{10, 30},
		}

		for _, tt := range tests {
			got := triple(tt.x)
			if got != tt.want {
				t.Errorf("triple(%d) = %d, want %d", tt.x, got, tt.want)
			}
		}
	})

	t.Run("умножение на 0", func(t *testing.T) {
		zero := Multiplier(0)

		if got := zero(100); got != 0 {
			t.Errorf("zero(100) = %d, want 0", got)
		}
	})

	t.Run("умножение на отрицательное", func(t *testing.T) {
		neg := Multiplier(-2)

		if got := neg(5); got != -10 {
			t.Errorf("neg(5) = %d, want -10", got)
		}
	})
}

func TestCompose(t *testing.T) {
	double := Multiplier(2)
	triple := Multiplier(3)

	t.Run("Compose(double, triple)(5) = 30", func(t *testing.T) {
		composed := Compose(double, triple)
		got := composed(5)
		want := 30

		if got != want {
			t.Errorf("Compose(double, triple)(5) = %d, want %d", got, want)
		}
	})

	t.Run("Compose(triple, double)(5) = 30", func(t *testing.T) {
		composed := Compose(triple, double)
		got := composed(5)
		want := 30

		if got != want {
			t.Errorf("Compose(triple, double)(5) = %d, want %d", got, want)
		}
	})

	t.Run("Compose(double, triple)(10) = 60", func(t *testing.T) {
		composed := Compose(double, triple)
		got := composed(10)
		want := 60

		if got != want {
			t.Errorf("Compose(double, triple)(10) = %d, want %d", got, want)
		}
	})
}

func TestComposeOrderMatters(t *testing.T) {
	t.Run("порядок важен: f∘g ≠ g∘f", func(t *testing.T) {

		f := func(x int) int {
			return x + 10
		}

		g := func(x int) int {
			return x * 2
		}

		composedFG := Compose(f, g)
		resultFG := composedFG(5)
		composedGF := Compose(g, f)
		resultGF := composedGF(5)

		if resultFG == resultGF {
			t.Errorf("f∘g и g∘f дают одинаковый результат: %d. Это означает, что порядок не имеет значения для этих функций, но для задачи нужны функции, где порядок важен.", resultFG)
		}

		t.Logf("f∘g(5) = %d, g∘f(5) = %d", resultFG, resultGF)
		t.Logf("Порядок важен: f∘g ≠ g∘f ✓")

		expected := 20
		if resultFG != expected {
			t.Errorf("Compose(f, g)(5) = %d, want %d", resultFG, expected)
		}
	})
}

func TestForEach(t *testing.T) {
	t.Run("остановка на первом кратном 4", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		var visited []int

		ForEach(s, func(x int) bool {
			visited = append(visited, x)
			return x%4 != 0
		})

		expected := []int{1, 2, 3, 4}
		if !reflect.DeepEqual(visited, expected) {
			t.Errorf("ForEach посетил %v, ожидалось %v", visited, expected)
		}
	})
}

func TestMySortFunc(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "разные длины",
			input: []string{"яблоко", "киви", "ананас", "манго"},
			want:  []string{"киви", "манго", "ананас", "яблоко"},
		},
		{
			name:  "разные языки",
			input: []string{"яблоко", "apple", "banana", "банан"},
			want:  []string{"apple", "банан", "banana", "яблоко"},
		},
		{
			name:  "слова с одинаковой длиной",
			input: []string{"дом", "сон", "лес", "кот"},
			want:  []string{"дом", "кот", "лес", "сон"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]string, len(tt.input))
			copy(input, tt.input)

			MySortFunc(input)
			if !reflect.DeepEqual(input, tt.want) {
				t.Errorf("SortByLengthThenAlphabet() = %v, want %v", input, tt.want)
			}
		})
	}
}

func TestSelfCheck(t *testing.T) {
	even := func(x int) bool { return x%2 == 0 }
	positive := func(x int) bool { return x > 0 }

	t.Run("And(even, Not(positive))", func(t *testing.T) {
		predicate := And(
			And(even, Not(positive)),
			func(x int) bool { return x != 0 },
		)

		numbers := []int{-4, -3, -2, -1, 0, 1, 2, 3, 4}

		result := Filter2(numbers, predicate)

		expected := []int{-4, -2}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Filter(numbers, And(even, Not(positive))) = %v, want %v", result, expected)
		}
	})
}

func Filter2(s []int, keep Pred) []int {
	result := make([]int, 0, len(s))
	for _, v := range s {
		if keep(v) {
			result = append(result, v)
		}
	}
	return result
}

func TestBSearchRec(t *testing.T) {
	tests := []struct {
		name   string
		s      []int
		search int
		want   int
	}{
		{
			name:   "3 ",
			s:      []int{1, 2, 3, 4, 5, 6},
			search: 3,
			want:   2,
		},
		{
			name:   "не найдено",
			s:      []int{1, 2, 3, 4, 5, 6},
			search: 7,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BSearchRec(tt.s, tt.search)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BSearchRec(%d) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestBSearchIter(t *testing.T) {
	tests := []struct {
		name   string
		s      []int
		search int
		want   int
	}{
		{
			name:   "3 ",
			s:      []int{1, 2, 3, 4, 5, 6},
			search: 3,
			want:   2,
		},
		{
			name:   "не найдено",
			s:      []int{1, 2, 3, 4, 5, 6},
			search: 7,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BSearchIter(tt.s, tt.search)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BSearchIter(%d) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestHanoi(t *testing.T) {
	t.Run("0 дисков", func(t *testing.T) {
		Hanoi(0, "A", "C", "B")
	})

	t.Run("1 диск", func(t *testing.T) {
		Hanoi(1, "A", "C", "B")
	})

	t.Run("2 диска", func(t *testing.T) {
		Hanoi(2, "A", "C", "B")
	})

	t.Run("3 диска", func(t *testing.T) {
		Hanoi(3, "A", "C", "B")
	})
}

func TestSafeDiv(t *testing.T) {
	tests := []struct {
		name  string
		a     int
		b     int
		want  int
		wantB bool
	}{
		{
			name:  "всё хорошо",
			a:     9,
			b:     3,
			want:  3,
			wantB: true,
		},
		{
			name:  "деление на ноль",
			a:     9,
			b:     0,
			want:  0,
			wantB: false,
		},
		{
			name:  "ноль делится на число",
			a:     0,
			b:     9,
			want:  0,
			wantB: true,
		},
		{
			name:  "деление на отрицательное",
			a:     9,
			b:     -1,
			want:  -9,
			wantB: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := SafeDiv(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("SafeDiv(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
			if ok != tt.wantB {
				t.Errorf("SafeDiv(%d, %d) ok = %v, want %v", tt.a, tt.b, ok, tt.wantB)
			}
		})
	}
}

func TestAt(t *testing.T) {
	tests := []struct {
		name  string
		s     []int
		i     int
		want  int
		wantB bool
	}{
		{
			name:  "всё хорошо",
			s:     []int{9, 5, 4, 3, 2, 1},
			i:     3,
			want:  3,
			wantB: true,
		},
		{
			name:  "такого индекса нет",
			s:     []int{9, 5, 4, 3, 2, 1},
			i:     6,
			want:  0,
			wantB: false,
		},
		{
			name:  "отрицательный индекс",
			s:     []int{9, 5, 4, 3, 2, 1},
			i:     -2,
			want:  0,
			wantB: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := At(tt.s, tt.i)

			if got != tt.want {
				t.Errorf("SafeDiv(%d, %d) = %d, want %d", tt.s, tt.i, got, tt.want)
			}
			if ok != tt.wantB {
				t.Errorf("SafeDiv(%d, %d) ok = %v, want %v", tt.s, tt.i, ok, tt.wantB)
			}
		})
	}
}
