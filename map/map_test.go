package _map

import (
	"reflect"
	"testing"
)

func TestTopFreq(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected []string
	}{
		{
			name:     "обычный текст с разной частотой",
			text:     "apple banana apple cherry banana apple date",
			expected: []string{"apple", "banana", "cherry", "date"},
		},
		{
			name:     "равные частоты - сортировка по алфавиту",
			text:     "cat dog bird cat bird fish",
			expected: []string{"bird", "cat", "dog", "fish"},
		},
		{
			name:     "все слова с одинаковой частотой",
			text:     "one two three four",
			expected: []string{"four", "one", "three", "two"},
		},
		{
			name:     "пустой текст",
			text:     "",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			freq := CountWords(tt.text)
			got := TopFreq(freq)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("TopFreq() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGroupByLen(t *testing.T) {
	tests := []struct {
		name     string
		text     []string
		expected map[int][]string
	}{
		{
			name:     "обычный текст",
			text:     []string{"го", "юг", "тест"},
			expected: map[int][]string{2: {"го", "юг"}, 4: {"тест"}},
		},
		{
			name:     "пустой текст",
			text:     []string{},
			expected: map[int][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupByLen(tt.text)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("GroupByLen() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkContainsSlice(b *testing.B) {
	s := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		s[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsSlice(s, 50000)
	}
}

func BenchmarkContainsMap(b *testing.B) {
	s := make(map[int]struct{}, 100000)
	for i := 0; i < 100000; i++ {
		s[i] = struct{}{}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsMap(s, 50000)
	}
}

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		text     []string
		expected [][]string
	}{
		{
			name:     "обычный текст с разной частотой",
			text:     []string{"кот", "ток", "дом", "мод", "окт"},
			expected: [][]string{{"кот", "ток", "окт"}, {"дом", "мод"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupAnagrams(tt.text)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("TopFreq() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected map[string][]int
	}{
		{
			name:     "обычный текст с разной частотой",
			text:     "го учит го",
			expected: map[string][]int{"го": {0, 2}, "учит": {1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Index(tt.text)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("TopFreq() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMergeSum(t *testing.T) {
	tests := []struct {
		name string
		a    map[string]int
		b    map[string]int
		want map[string]int
	}{
		{
			name: "одинаковые символы",
			a:    map[string]int{"a": 1, "b": 2, "c": 3},
			b:    map[string]int{"a": 1, "b": 2, "c": 3},
			want: map[string]int{"a": 2, "b": 4, "c": 6},
		},
		{
			name: "пустой b",
			a:    map[string]int{"a": 1, "b": 2, "c": 3},
			b:    map[string]int{},
			want: map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name: "пустой a",
			a:    map[string]int{},
			b:    map[string]int{"a": 1, "b": 2, "c": 3},
			want: map[string]int{"a": 1, "b": 2, "c": 3},
		},
		{
			name: "разные ключи",
			a:    map[string]int{"a": 1, "x": 2, "y": 3},
			b:    map[string]int{"a": 1, "b": 2, "c": 3},
			want: map[string]int{"a": 2, "b": 2, "c": 3, "x": 2, "y": 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSum(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopFreq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapsEqual(t *testing.T) {
	tests := []struct {
		name     string
		a        map[string]int
		b        map[string]int
		expected bool
	}{
		{
			name:     "равные",
			a:        map[string]int{"a": 1, "b": 2, "c": 3},
			b:        map[string]int{"a": 1, "b": 2, "c": 3},
			expected: true,
		},
		{
			name:     "разной длины",
			a:        map[string]int{"a": 1, "b": 2, "c": 3},
			b:        map[string]int{"a": 1, "b": 2},
			expected: false,
		},
		{
			name:     "одинаковые ключи разные значения",
			a:        map[string]int{"a": 1, "b": 2, "c": 3},
			b:        map[string]int{"a": 4, "b": 5, "c": 3},
			expected: false,
		},
		{
			name:     "nil and nil",
			a:        nil,
			b:        nil,
			expected: true,
		},
		{
			name:     "nil and empty",
			a:        nil,
			b:        map[string]int{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapsEqual(tt.a, tt.b)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("MapsEqual() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSteps(t *testing.T) {
	tests := []struct {
		name     string
		runes    []rune
		expected map[[2]int]int
	}{
		{
			name:  "пустой путь",
			runes: []rune(""),
			expected: map[[2]int]int{
				{0, 0}: 1,
			},
		},
		{
			name:  "сложный маршрут с множеством возвратов",
			runes: []rune("→←→←↑↓↑↓"),
			expected: map[[2]int]int{
				{0, 0}: 5,
				{1, 0}: 2,
				{0, 1}: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Steps(tt.runes)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("TopFreq() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCircleIterationOrderInOneRun(t *testing.T) {
	m := Circle()

	var firstOrder []int
	for key := range m {
		firstOrder = append(firstOrder, key)
	}

	var secondOrder []int
	for key := range m {
		secondOrder = append(secondOrder, key)
	}

	var thirdOrder []int
	for key := range m {
		thirdOrder = append(thirdOrder, key)
	}

	t.Logf("Первый обход:  %v", firstOrder)
	t.Logf("Второй обход: %v", secondOrder)
	t.Logf("Третий обход: %v", thirdOrder)
}

func TestCircleIterationOrderPrint(t *testing.T) {
	m := Circle()

	var keys []int
	for key := range m {
		keys = append(keys, key)
	}

	t.Logf("Порядок ключей в этом запуске: %v", keys)
}

func TestSetGrade(t *testing.T) {
	tests := []struct {
		name    string
		student string
		subject string
		grade   int
		setup   func(map[string]map[string]int)
		want    map[string]map[string]int
	}{
		{
			name:    "новый студент, новый предмет",
			student: "Анна",
			subject: "Математика",
			grade:   5,
			setup:   func(m map[string]map[string]int) {},
			want: map[string]map[string]int{
				"Анна": {"Математика": 5},
			},
		},
		{
			name:    "существующий студент, новый предмет",
			student: "Анна",
			subject: "Физика",
			grade:   4,
			setup: func(m map[string]map[string]int) {
				// ВАЖНО: создаем внутреннюю карту правильно
				m["Анна"] = map[string]int{
					"Математика": 5,
				}
			},
			want: map[string]map[string]int{
				"Анна": {
					"Математика": 5,
					"Физика":     4,
				},
			},
		},
		{
			name:    "существующий студент, существующий предмет (перезапись)",
			student: "Анна",
			subject: "Математика",
			grade:   3,
			setup: func(m map[string]map[string]int) {
				m["Анна"] = map[string]int{
					"Математика": 5,
				}
			},
			want: map[string]map[string]int{
				"Анна": {"Математика": 3},
			},
		},
		{
			name:    "несколько студентов",
			student: "Петр",
			subject: "Химия",
			grade:   5,
			setup: func(m map[string]map[string]int) {
				m["Анна"] = map[string]int{
					"Математика": 5,
				}
			},
			want: map[string]map[string]int{
				"Анна": {"Математика": 5},
				"Петр": {"Химия": 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grades := make(map[string]map[string]int)

			if tt.setup != nil {
				tt.setup(grades)
			}

			SetGrade(grades, tt.student, tt.subject, tt.grade)
			if !reflect.DeepEqual(grades, tt.want) {
				t.Errorf("SetGrade() = %v, want %v", grades, tt.want)
			}
		})
	}
}

func TestGetGrade(t *testing.T) {
	tests := []struct {
		name    string
		student string
		subject string
		setup   func(map[string]map[string]int)
		want    int
		wantOk  bool
	}{
		{
			name:    "существующий студент, существующий предмет",
			student: "Анна",
			subject: "Математика",
			setup: func(m map[string]map[string]int) {
				m["Анна"] = map[string]int{
					"Математика": 5,
				}
			},
			want:   5,
			wantOk: true,
		},
		{
			name:    "существующий студент, несуществующий предмет",
			student: "Анна",
			subject: "Физика",
			setup: func(m map[string]map[string]int) {
				m["Анна"] = map[string]int{
					"Математика": 5,
				}
			},
			want:   0,
			wantOk: false,
		},
		{
			name:    "несуществующий студент",
			student: "Петр",
			subject: "Математика",
			setup: func(m map[string]map[string]int) {
				m["Анна"] = map[string]int{
					"Математика": 5,
				}
			},
			want:   0,
			wantOk: false,
		},
		{
			name:    "пустая карта",
			student: "Анна",
			subject: "Математика",
			setup:   func(m map[string]map[string]int) {},
			want:    0,
			wantOk:  false,
		},
		{
			name:    "существующий студент с пустой внутренней картой",
			student: "Анна",
			subject: "Математика",
			setup: func(m map[string]map[string]int) {
				m["Анна"] = make(map[string]int)
			},
			want:   0,
			wantOk: false,
		},
		{
			name:    "несколько студентов и предметов",
			student: "Петр",
			subject: "Химия",
			setup: func(m map[string]map[string]int) {
				m["Анна"] = map[string]int{
					"Математика": 5,
					"Физика":     4,
				}
				m["Петр"] = map[string]int{
					"Химия":    5,
					"Биология": 3,
				}
			},
			want:   5,
			wantOk: true,
		},
		{
			name:    "существующий студент с nil внутренней картой",
			student: "Анна",
			subject: "Математика",
			setup: func(m map[string]map[string]int) {
				m["Анна"] = nil
			},
			want:   0,
			wantOk: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grades := make(map[string]map[string]int)

			if tt.setup != nil {
				tt.setup(grades)
			}

			got, ok := GetGrade(grades, tt.student, tt.subject)
			if got != tt.want {
				t.Errorf("GetGrade() grade = %v, want %v", got, tt.want)
			}
			if ok != tt.wantOk {
				t.Errorf("GetGrade() ok = %v, want %v", ok, tt.wantOk)
			}
		})
	}
}
