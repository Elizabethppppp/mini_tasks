package slices

import (
	"fmt"
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

func TestInsertAt2(t *testing.T) {
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
			got := InsertAt2(tt.s, tt.i, tt.v)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertAt1(%v, %d, %d) = %v, want %v", tt.s, tt.i, tt.v, got, tt.want)
			}
		})
	}
}

func TestFilterEven(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{
			name: "чётные маленткие",
			s:    []int{1, 2, 4},
			want: []int{2, 4},
		},
		{
			name: "чётные длинные",
			s:    []int{1, 2, 4, 50, 22},
			want: []int{2, 4, 50, 22},
		},
		{
			name: "ничего",
			s:    []int{1, 3, 5},
			want: []int{},
		},
		{
			name: "в пустой срез",
			s:    []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addrBefore := fmt.Sprintf("%p", tt.s)
			got := FilterEven(tt.s)
			addrAfter := fmt.Sprintf("%p", got)
			if addrBefore != addrAfter {
				t.Errorf("адрес изменился: до %s, после %s", addrBefore, addrAfter)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertAt1(%d) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		i    int
		want []int
	}{
		{
			name: "1",
			s:    []int{1, 2, 4},
			i:    1,
			want: []int{2, 4, 1},
		},
		{
			name: "2",
			s:    []int{1, 2, 3, 4, 5},
			i:    2,
			want: []int{3, 4, 5, 1, 2},
		},
		{
			name: "k равно длине",
			s:    []int{1, 2, 3, 4, 5, 6},
			i:    6,
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "пустой срез",
			s:    []int{},
			want: []int{},
		},
		{
			name: "0",
			s:    []int{1, 2, 3},
			i:    0,
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]int, len(tt.s))
			copy(original, tt.s)
			Rotate(tt.s, tt.i)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("Rotate(%v, %d) = %v, want %v", original, tt.i, tt.s, tt.want)
			}
		})
	}
}

func TestStackTable(t *testing.T) {
	tests := []struct {
		name     string
		pushes   []int
		pops     int
		wantVals []int
		wantOk   []bool
		wantLen  int
	}{
		{
			name:     "Push 1,2,3 и Pop все",
			pushes:   []int{1, 2, 3},
			pops:     3,
			wantVals: []int{3, 2, 1},
			wantOk:   []bool{true, true, true},
			wantLen:  0,
		},
		{
			name:     "Pop из пустого",
			pushes:   []int{},
			pops:     1,
			wantVals: []int{0},
			wantOk:   []bool{false},
			wantLen:  0,
		},
		{
			name:     "Push после Pop",
			pushes:   []int{1, 2, 3},
			pops:     4,
			wantVals: []int{3, 2, 1, 0},
			wantOk:   []bool{true, true, true, false},
			wantLen:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{}
			for _, v := range tt.pushes {
				s.Push(v)
			}
			for i := 0; i < tt.pops; i++ {
				val, ok := s.Pop()
				if val != tt.wantVals[i] || ok != tt.wantOk[i] {
					t.Errorf("Pop() #%d = %d, %v, want %d, %v",
						i, val, ok, tt.wantVals[i], tt.wantOk[i])
				}
			}

			if s.Len() != tt.wantLen {
				t.Errorf("Len() = %d, want %d", s.Len(), tt.wantLen)
			}
		})
	}
}
