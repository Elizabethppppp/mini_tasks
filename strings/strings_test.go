package strings

import (
	"reflect"
	"testing"
)

func TestMyRuneCount(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "русское слово",
			s:    "привет",
			want: 6,
		},
		{
			name: "английское слово",
			s:    "hello",
			want: 5,
		},
		{
			name: "пустая строка",
			s:    "",
			want: 0,
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MyRuneCount(tt.s)
			if got != tt.want {
				t.Errorf("process(%s) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestUpperASCII(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "русское слово",
			s:    "привет",
			want: "привет",
		},
		{
			name: "английское слово",
			s:    "hello",
			want: "HELLO",
		},
		{
			name: "пустая строка",
			s:    "",
			want: "",
		},
		{
			name: "слово с цифрами",
			s:    "go1.22",
			want: "GO1.22",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UpperASCII(tt.s)
			if got != tt.want {
				t.Errorf("process(%s) = %s, want %s", tt.s, got, tt.want)
			}
		})
	}
}

func TestTruncateBytes(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		count int
		want  string
	}{
		{
			name:  "русское слово часть",
			s:     "привет",
			count: 5,
			want:  "п...",
		},
		{
			name:  "русское слово всё",
			s:     "привет",
			count: 100,
			want:  "привет",
		},
		{
			name:  "английское слово",
			s:     "hello",
			count: 2,
			want:  "",
		},
		{
			name:  "английское слово 3 символа",
			count: 3,
			s:     "ohh",
			want:  "...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TruncateBytes(tt.s, tt.count)
			if got != tt.want {
				t.Errorf("process(%s) = %s, want %s", tt.s, got, tt.want)
			}
		})
	}
}

func TestReverseRunes(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "русское слово",
			s:    "привет",
			want: "тевирп",
		},
		{
			name: "английское слово",
			s:    "hello",
			want: "olleh",
		},
		{
			name: "пустая строка",
			s:    "",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseRunes(tt.s)
			if got != tt.want {
				t.Errorf("process(%s) = %s, want %s", tt.s, got, tt.want)
			}
			gotTwice := ReverseRunes(got)
			if gotTwice != tt.s {
				t.Errorf("ReverseRunes(ReverseRunes(%s)) = %s, want %s", tt.s, gotTwice, tt.s)
			}
		})
	}
}

func TestMyCount(t *testing.T) {
	tests := []struct {
		name string
		s    string
		sub  string
		want int
	}{
		{
			name: "русское слово",
			s:    "привет",
			sub:  "hello",
			want: 0,
		},
		{
			name: "русское слово",
			s:    "привет",
			sub:  "п",
			want: 1,
		},
		{
			name: "английское слово",
			s:    "hello",
			sub:  "ll",
			want: 1,
		},
		{
			name: "английское слово",
			s:    "hello",
			sub:  "l",
			want: 2,
		},
		{
			name: "aaaa",
			s:    "aaaa",
			sub:  "aa",
			want: 2,
		},
		{
			name: "aaaa",
			s:    "aaaa",
			sub:  "aaa",
			want: 1,
		},
		{
			name: "пустая строка",
			s:    "",
			sub:  "hello",
			want: 0,
		},
		{
			name: "пустая подстрока",
			s:    "hello",
			sub:  "",
			want: 0,
		},
		{
			name: "пустые",
			s:    "",
			sub:  "",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MyCount(tt.s, tt.sub)
			if got != tt.want {
				t.Errorf("process(%s) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestMyFields(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "русские слова",
			s:    "  раз\tдва  три\n",
			want: []string{"раз", "два", "три"},
		},
		{
			name: "английские слова",
			s:    "hello\n piece",
			want: []string{"hello", "piece"},
		},
		{
			name: "пустая строка",
			s:    "",
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MyFields(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MyFields(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestCaesar(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		shift int
		want  string
	}{
		{
			name:  "русские слова",
			s:     "Hello, мир!",
			shift: 3,
			want:  "Khoor, мир!",
		},
		{
			name:  "английские слова",
			s:     "Hello, World!",
			shift: 3,
			want:  "Khoor, Zruog!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Caesar(tt.s, tt.shift)
			if got != tt.want {
				t.Errorf("process(%s) = %s, want %s", tt.s, got, tt.want)
			}
		})
	}
}

func TestMyValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "русское слово",
			s:    "привет",
			want: true,
		},
		{
			name: "английское слово",
			s:    "hello",
			want: true,
		},
		{
			name: "symbol",
			s:    "a\xffb",
			want: false,
		},
		{
			name: "symbol 2",
			s:    "\xd0",
			want: false,
		},
		{
			name: "3 символа",
			s:    "\xe0\xa0\x80",
			want: true,
		},
		{
			name: "",
			s:    "",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MyValid(tt.s)
			if got != tt.want {
				t.Errorf("process(%s) = %t, want %t", tt.s, got, tt.want)
			}
		})
	}
}

func TestSanitize(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "русское слово",
			s:    "привет",
			want: "привет",
		},
		{
			name: "2 symbols",
			s:    "\xc2\x80",
			want: "\u0080",
		},
		{
			name: "symbol 2",
			s:    "\xd0",
			want: "�",
		},
		{
			name: "",
			s:    "",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sanitize(tt.s)
			if got != tt.want {
				t.Errorf("process(%s) = %s, want %s", tt.s, got, tt.want)
			}
		})
	}
}

func TestFreq(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want map[rune]int
	}{
		{
			name: "русские слова",
			s:    "абракадабра",
			want: map[rune]int{'а': 5, 'б': 2, 'р': 2, 'к': 1, 'д': 1},
		},
		{
			name: "",
			s:    "",
			want: map[rune]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Freq(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Freq(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
