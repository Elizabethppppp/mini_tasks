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
