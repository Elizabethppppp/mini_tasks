package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want bool
	}{
		{
			name: "равные срезы",
			a:    []int{1, 2, 3},
			b:    []int{1, 2, 3},
			want: true,
		},
		{
			name: "разные срезы",
			a:    []int{1, 2, 3},
			b:    []int{1, 2, 4},
			want: false,
		},
		{
			name: "разная длина",
			a:    []int{1, 2, 3},
			b:    []int{1, 2},
			want: false,
		},
		{
			name: "оба пустых",
			a:    []int{},
			b:    []int{},
			want: true,
		},
		{
			name: "nil и пустой",
			a:    nil,
			b:    []int{},
			want: true,
		},
		{
			name: "пустой и nil",
			a:    []int{},
			b:    nil,
			want: true,
		},
		{
			name: "nil и nil",
			a:    nil,
			b:    nil,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Equal(tt.a, tt.b)
			wantSlices := slices.Equal(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("Equal(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
			if got != wantSlices {
				t.Errorf("Ваша функция вернула %v, а slices.Equal вернул %v", got, wantSlices)
			}
		})
	}
}

func TestProcess(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "положительное число меньше 100",
			s:    "50",
			want: "ok",
		},
		{
			name: "положительное число больше 99",
			s:    "100",
			want: "big",
		},
		{
			name: "положительное число больше 100",
			s:    "150",
			want: "big",
		},
		{
			name: "отрицательное число",
			s:    "-5",
			want: "neg",
		},
		{
			name: "ноль",
			s:    "0",
			want: "neg",
		},
		{
			name: "не число",
			s:    "abc",
			want: "err",
		},
		{
			name: "пустая строка",
			s:    "",
			want: "err",
		},
		{
			name: "число с минусом",
			s:    "-10",
			want: "neg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := process(tt.s)
			if got != tt.want {
				t.Errorf("process(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestReadAge(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    int
		wantErr bool
		wantOut bool
	}{
		{
			name:    "положительное число меньше 150",
			s:       "50",
			want:    50,
			wantErr: false,
		},
		{
			name:    "положительное число больше 150",
			s:       "151",
			want:    151,
			wantErr: true,
			wantOut: true,
		},
		{
			name:    "отрицательное число",
			s:       "-5",
			want:    0,
			wantErr: true,
			wantOut: true,
		},
		{
			name:    "ноль",
			s:       "0",
			want:    0,
			wantErr: false,
		},
		{
			name:    "не число",
			s:       "abc",
			want:    0,
			wantErr: true,
			wantOut: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadAge(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAge(%q) error = %v, wantErr %v", tt.s, err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.want {
				t.Errorf("ReadAge(%q) = %d, want %d", tt.s, got, tt.want)
			}

			if tt.wantErr && err == nil {
				t.Errorf("ReadAge(%q) ожидалась ошибка, но вернулась nil", tt.s)
			}

		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want map[string]int
	}{
		{
			name: "русские слова",
			s:    "го го го учит го",
			want: map[string]int{"го": 4, "учит": 1},
		},
		{
			name: "",
			s:    "",
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountWords(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Freq(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
