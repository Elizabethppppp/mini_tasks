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
