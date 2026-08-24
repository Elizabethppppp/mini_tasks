package methordStruct

import "testing"

func TestCelsius_ToF(t *testing.T) {
	tests := []struct {
		name string
		cel  Celsius
		want Fahrenheit
	}{
		{
			name: "22.4 cel to 72.32 fahrenheit",
			cel:  22.4,
			want: 72.32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cel.ToF()

			const epsilon = 0.0001
			if diff := float64(got - tt.want); diff < -epsilon || diff > epsilon {
				t.Errorf("Celsius(%v).ToF() = %v, want %v", tt.cel, got, tt.want)
			}
		})
	}
}

func TestPath_String(t *testing.T) {
	tests := []struct {
		name string
		s    Path
		want string
	}{
		{
			name: "есть несколько",
			s:    Path{"usr", "local", "bin"},
			want: "usr/local/bin",
		},
		{
			name: "пустой",
			s:    Path{},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.String()

			if got != tt.want {
				t.Errorf("Path(%v).String() = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestNewAccount(t *testing.T) {
	tests := []struct {
		name      string
		owner     string
		start     int
		wantAcc   bool
		wantOwner string
		wantBal   int
	}{
		{
			name:      "valid account",
			owner:     "Alice",
			start:     1000,
			wantAcc:   true,
			wantOwner: "Alice",
			wantBal:   1000,
		},
		{
			name:      "valid with zero balance",
			owner:     "Bob",
			start:     0,
			wantAcc:   true,
			wantOwner: "Bob",
			wantBal:   0,
		},
		{
			name:      "empty owner",
			owner:     "",
			start:     500,
			wantAcc:   false,
			wantOwner: "",
			wantBal:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := NewAccount(tt.owner, tt.start)

			if ok != tt.wantAcc {
				t.Errorf("NewAccount() ok = %v, want %v", ok, tt.wantAcc)
			}

			if !tt.wantAcc {
				if got != nil {
					t.Errorf("NewAccount() got = %v, want nil", got)
				}
				return
			}

			if got == nil {
				t.Errorf("NewAccount() returned nil, but expected valid account")
				return
			}

			if got.owner != tt.wantOwner {
				t.Errorf("NewAccount() owner = %q, want %q", got.owner, tt.wantOwner)
			}

			if got.start != tt.wantBal {
				t.Errorf("NewAccount() balance = %d, want %d", got.start, tt.wantBal)
			}
		})
	}
}
