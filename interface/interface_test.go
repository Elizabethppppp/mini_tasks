package _interface

import "testing"

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
