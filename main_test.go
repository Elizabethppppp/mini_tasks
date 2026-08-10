package main

import (
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