package _interface

import (
	"math"
	"testing"
)

func TestDescribe(t *testing.T) {
	tests := []struct {
		name string
		x    any
		want string
	}{
		{
			name: "целое",
			x:    22,
			want: "целое 22",
		},
		{
			name: "строка",
			x:    "привет",
			want: "строка привет длины 6",
		},
		{
			name: "срез",
			x:    []any{2, 3, 4},
			want: "срез из 3 элементов",
		},
		{
			name: "булево",
			x:    true,
			want: "булево true",
		},
		{
			name: "другое",
			x:    22.4,
			want: "что-то типа float64",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Describe(tt.x)

			if got != tt.want {
				t.Errorf("Describe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalArea(t *testing.T) {
	tests := []struct {
		name     string
		shapes   []Shape
		expected float64
	}{
		{
			name: "круг + прямоугольник + треугольник",
			shapes: []Shape{
				Circle{R: 5},
				Rectangle{W: 3, H: 4},
				Triangle{A: 4, H: 3},
			},
			expected: math.Pi*25 + 12 + 6,
		},
		{
			name:     "пустой слайс",
			shapes:   []Shape{},
			expected: 0.0,
		},
		{
			name: "только круги",
			shapes: []Shape{
				Circle{R: 1},
				Circle{R: 2},
			},
			expected: math.Pi + 4*math.Pi,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TotalArea(tt.shapes)
			if got != tt.expected {
				t.Errorf("TotalArea() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestLargest(t *testing.T) {
	tests := []struct {
		name     string
		shapes   []Shape
		expected Shape
	}{
		{
			name: "круг больше всех",
			shapes: []Shape{
				Triangle{A: 4, H: 3},
				Rectangle{W: 3, H: 4},
				Circle{R: 5},
			},
			expected: Circle{R: 5},
		},
		{
			name: "прямоугольник больше всех",
			shapes: []Shape{
				Triangle{A: 1, H: 1},
				Rectangle{W: 10, H: 10},
				Circle{R: 1},
			},
			expected: Rectangle{W: 10, H: 10},
		},
		{
			name:     "пустой слайс",
			shapes:   []Shape{},
			expected: nil,
		},
		{
			name: "один треугольник",
			shapes: []Shape{
				Triangle{A: 4, H: 3},
			},
			expected: Triangle{A: 4, H: 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Largest(tt.shapes)

			if tt.expected == nil {
				if got != nil {
					t.Errorf("Largest() = %v, want nil", got)
				}
				return
			}

			if got == nil || got.Area() != tt.expected.Area() {
				t.Errorf("Largest() area = %v, want area %v", got.Area(), tt.expected.Area())
			}
		})
	}
}
