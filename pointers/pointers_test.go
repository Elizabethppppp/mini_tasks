package pointers

import "testing"

func TestGetOr(t *testing.T) {
	tests := []struct {
		name string
		p    *int
		def  int
		want int
	}{
		{
			name: "*p == nil",
			p:    nil,
			def:  43,
			want: 43,
		},
		{
			name: "*p == 3",
			p:    func() *int { v := 3; return &v }(),
			def:  43,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetOr(tt.p, tt.def)

			if got != tt.want {
				t.Errorf("SafeDiv(%d, %d) = %d, want %d", tt.p, tt.def, got, tt.want)
			}
		})
	}
}

func TestAgeString(t *testing.T) {
	tests := []struct {
		name string
		age  *int
		want string
	}{
		{
			name: "возраст не указан (nil)",
			age:  nil,
			want: "не указан",
		},
		{
			name: "младенец (0 лет)",
			age:  func() *int { v := 0; return &v }(),
			want: "0",
		},
		{
			name: "взрослый (30 лет)",
			age:  func() *int { v := 30; return &v }(),
			want: "30",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AgeString(tt.age)
			if got != tt.want {
				t.Errorf("AgeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParsePoint(t *testing.T) {
	tests := []struct {
		name   string
		coord  string
		wantX  int
		wantY  int
		wantOk bool
	}{
		{
			name:   "3;7",
			coord:  "3;7",
			wantX:  3,
			wantY:  7,
			wantOk: true,
		},
		{
			name:   "0;0",
			coord:  "0;0",
			wantX:  0,
			wantY:  0,
			wantOk: true,
		},
		{
			name:   "неверный формат",
			coord:  "abc",
			wantX:  0,
			wantY:  0,
			wantOk: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY, gotOk := ParsePoint(tt.coord)

			if gotX != tt.wantX || gotY != tt.wantY || gotOk != tt.wantOk {
				t.Errorf("ParsePoint(%q) = (%d, %d, %v), want (%d, %d, %v)",
					tt.coord, gotX, gotY, gotOk, tt.wantX, tt.wantY, tt.wantOk)
			}
		})
	}
}

func TestParsePointInto(t *testing.T) {
	tests := []struct {
		name   string
		coord  string
		x      *int
		y      *int
		wantOk bool
	}{
		{
			name:   "3;7",
			coord:  "3;7",
			x:      func() *int { v := 3; return &v }(),
			y:      func() *int { v := 7; return &v }(),
			wantOk: true,
		},
		{
			name:   "0;0",
			coord:  "0;0",
			x:      func() *int { v := 0; return &v }(),
			y:      func() *int { v := 0; return &v }(),
			wantOk: true,
		},
		{
			name:   "неверный формат",
			coord:  "abc",
			wantOk: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk := ParsePointInto(tt.coord, tt.x, tt.y)

			if gotOk != tt.wantOk {
				t.Errorf("ParsePoint(%q) = (%v), want ( %v)",
					tt.coord, gotOk, tt.wantOk)
			}
		})
	}
}

func BenchmarkSumVal(b *testing.B) {
	data := Big{}
	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumVal(data)
	}
}

func BenchmarkSumPtr(b *testing.B) {
	data := Big{}
	for i := 0; i < len(data); i++ {
		data[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumPtr(&data)
	}
}

func TestResetAndRedirect(t *testing.T) {
	x := 10
	y := 20
	p := &x

	Redirect(&p, &y)
	if *p != 20 {
		t.Fatalf("после Redirect: *p должно быть 20, получено %d", *p)
	}

	Reset(&p)
	if p != nil {
		t.Fatalf("после Reset: p должен быть nil, получен %v", p)
	}

	if x != 10 {
		t.Errorf("x должен быть 10, получено %d", x)
	}
	if y != 20 {
		t.Errorf("y должен быть 20, получено %d", y)
	}
}
