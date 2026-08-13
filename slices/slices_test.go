package slices

import (
	"reflect"
	"testing"
)

func TestHead(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		n    int
		want []int
	}{
		{
			name: "3 элемента среза",
			s:    []int{1, 2, 3, 4, 5, 6},
			n:    3,
			want: []int{1, 2, 3},
		},
		{
			name: "",
			s:    []int{},
			want: []int{},
		},
		{
			name: "all elements",
			s:    []int{1, 2, 3, 4, 5, 6},
			n:    6,
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Head(tt.s, tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Freq(%d) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestCopy1(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{
			name: "3 элемента среза",
			s:    []int{1, 2, 3, 4, 5, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "",
			s:    []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Copy1(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Freq(%d) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestCopy2(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{
			name: "3 элемента среза",
			s:    []int{1, 2, 3, 4, 5, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "",
			s:    []int{},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Copy2(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Freq(%d) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestCopy3(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{
			name: "3 элемента среза",
			s:    []int{1, 2, 3, 4, 5, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "",
			s:    []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Copy3(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Freq(%d) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestDeleteAt1(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		n    int
		want []int
	}{
		{
			name: "3 элемента среза",
			s:    []int{1, 2, 3, 4, 5, 6},
			n:    3,
			want: []int{1, 2, 3, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeleteAt1(tt.s, tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Freq(%d) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestDeleteAt2(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		n    int
		want []int
	}{
		{
			name: "3 элемента среза",
			s:    []int{1, 2, 3, 4, 5, 6},
			n:    3,
			want: []int{1, 2, 3, 6, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeleteAt2(tt.s, tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Freq(%d) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestInsertAt1(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		i    int
		v    int
		want []int
	}{
		{
			name: "в начало",
			s:    []int{1, 2, 4},
			i:    0,
			v:    3,
			want: []int{3, 1, 2, 4},
		},
		{
			name: "в середину",
			s:    []int{1, 2, 4},
			i:    2,
			v:    3,
			want: []int{1, 2, 3, 4},
		},
		{
			name: "в конец",
			s:    []int{1, 2, 3, 4, 5, 6},
			i:    6,
			v:    7,
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name: "в пустой срез",
			s:    []int{},
			i:    0,
			v:    1,
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertAt1(tt.s, tt.i, tt.v)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertAt1(%v, %d, %d) = %v, want %v", tt.s, tt.i, tt.v, got, tt.want)
			}
		})
	}
}
