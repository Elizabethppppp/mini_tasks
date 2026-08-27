package myErrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestValidateForm(t *testing.T) {
	tests := []struct {
		name            string
		inputName       string
		inputEmail      string
		inputAge        int
		wantErr         bool
		wantErrBadName  bool
		wantErrBadEmail bool
		wantErrBadAge   bool
	}{
		{
			name:            "все валидно",
			inputName:       "Alice",
			inputEmail:      "alice@example.com",
			inputAge:        25,
			wantErr:         false,
			wantErrBadName:  false,
			wantErrBadEmail: false,
			wantErrBadAge:   false,
		},
		{
			name:            "три невалидных поля",
			inputName:       "",
			inputEmail:      "not-an-email",
			inputAge:        -5,
			wantErr:         true,
			wantErrBadName:  true,
			wantErrBadEmail: true,
			wantErrBadAge:   true,
		},
		{
			name:            "только имя невалидно",
			inputName:       "",
			inputEmail:      "alice@example.com",
			inputAge:        25,
			wantErr:         true,
			wantErrBadName:  true,
			wantErrBadEmail: false,
			wantErrBadAge:   false,
		},
		{
			name:            "только email невалиден",
			inputName:       "Alice",
			inputEmail:      "invalid",
			inputAge:        25,
			wantErr:         true,
			wantErrBadName:  false,
			wantErrBadEmail: true,
			wantErrBadAge:   false,
		},
		{
			name:            "только возраст невалиден",
			inputName:       "Alice",
			inputEmail:      "alice@example.com",
			inputAge:        150,
			wantErr:         true,
			wantErrBadName:  false,
			wantErrBadEmail: false,
			wantErrBadAge:   true,
		},
		{
			name:            "имя с пробелом и дефисом - валидно",
			inputName:       "Anna-Maria",
			inputEmail:      "anna@example.com",
			inputAge:        30,
			wantErr:         false,
			wantErrBadName:  false,
			wantErrBadEmail: false,
			wantErrBadAge:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateForm(tt.inputName, tt.inputEmail, tt.inputAge)

			if tt.wantErr && err == nil {
				t.Errorf("ValidateForm() expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("ValidateForm() expected nil, got %v", err)
			}

			if err == nil {
				return
			}

			hasBadName := errors.Is(err, errBadName)
			if hasBadName != tt.wantErrBadName {
				t.Errorf("ValidateForm() errors.Is(err, ErrBadName) = %v, want %v",
					hasBadName, tt.wantErrBadName)
			}

			hasBadEmail := errors.Is(err, errBadEmail)
			if hasBadEmail != tt.wantErrBadEmail {
				t.Errorf("ValidateForm() errors.Is(err, ErrBadEmail) = %v, want %v",
					hasBadEmail, tt.wantErrBadEmail)
			}

			hasBadAge := errors.Is(err, errBadAge)
			if hasBadAge != tt.wantErrBadAge {
				t.Errorf("ValidateForm() errors.Is(err, ErrBadAge) = %v, want %v",
					hasBadAge, tt.wantErrBadAge)
			}
		})
	}
}

func TestValidateForm_JoinReturnsNilForNoErrors(t *testing.T) {
	var errs []error
	result := errors.Join(errs...)
	if result != nil {
		t.Errorf("errors.Join(nil...) should return nil, got %v", result)
	}

	err := ValidateForm("John", "john@test.com", 30)
	if err != nil {
		t.Errorf("ValidateForm() should return nil for valid data, got %v", err)
	}

	err2 := ValidateForm("Jane", "jane@test.com", 25)
	if err2 != nil {
		t.Errorf("ValidateForm() with valid data should return nil, got %v", err2)
	}
}

func TestChain(t *testing.T) {
	tests := []struct {
		name  string
		input error
		want  []string
	}{
		{
			name:  "nil ошибка",
			input: nil,
			want:  []string{},
		},
		{
			name:  "простая ошибка",
			input: errors.New("simple error"),
			want:  []string{"simple error"},
		},
		{
			name: "трёхэтажная цепочка",
			input: func() error {
				err1 := errors.New("bottom error")
				err2 := fmt.Errorf("middle error: %w", err1)
				err3 := fmt.Errorf("top error: %w", err2)
				return err3
			}(),
			want: []string{
				"top error: middle error: bottom error",
				"middle error: bottom error",
				"bottom error",
			},
		},
		{
			name: "двухэтажная цепочка",
			input: func() error {
				err1 := errors.New("inner error")
				err2 := fmt.Errorf("outer error: %w", err1)
				return err2
			}(),
			want: []string{
				"outer error: inner error",
				"inner error",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Chain(tt.input)

			if len(got) != len(tt.want) {
				t.Errorf("Chain() length = %d, want %d", len(got), len(tt.want))
				t.Errorf("Got: %q", got)
				t.Errorf("Want: %q", tt.want)
				return
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Chain()[%d] = %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestRetry(t *testing.T) {
	baseErr := errors.New("connection failed")

	tests := []struct {
		name      string
		attempts  int
		wantErr   bool
		wantCalls int
		f         func(callCount *int) error
	}{
		{
			name:      "успех на 3-й попытке с 5 попытками",
			attempts:  5,
			wantErr:   false,
			wantCalls: 3,
			f: func(callCount *int) error {
				*callCount++
				if *callCount < 3 {
					return baseErr
				}
				return nil
			},
		},
		{
			name:      "2 попытки недостаточно для успеха на 3-й",
			attempts:  2,
			wantErr:   true,
			wantCalls: 2,
			f: func(callCount *int) error {
				*callCount++
				if *callCount < 3 {
					return baseErr
				}
				return nil
			},
		},
		{
			name:      "всегда возвращает ошибку",
			attempts:  3,
			wantErr:   true,
			wantCalls: 3,
			f: func(callCount *int) error {
				*callCount++
				return baseErr
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callCount := 0

			f := func() error {
				return tt.f(&callCount)
			}

			err := Retry(tt.attempts, f)

			if callCount != tt.wantCalls {
				t.Errorf("Retry() called %d times, want %d", callCount, tt.wantCalls)
			}

			if tt.wantErr && err == nil {
				t.Error("Retry() expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Retry() expected nil, got error: %v", err)
			}

			if tt.wantErr {
				if !errors.Is(err, baseErr) {
					t.Errorf("errors.Is(%v, %v) = false, want true", err, baseErr)
				}
				expectedMsg := fmt.Sprintf("после %d попыток: connection failed", tt.attempts)
				if err.Error() != expectedMsg {
					t.Errorf("error message = %q, want %q", err.Error(), expectedMsg)
				}
			}
		})
	}
}

func TestNewRetry(t *testing.T) {
	var ErrPermanent = errors.New("permanent error")

	tests := []struct {
		name       string
		attempts   int
		f          func(callCount *int) error
		wantErr    bool
		wantCalls  int
		wantErrIs  error
		wantErrMsg string
	}{
		{
			name:     "временная ошибка, затем успех",
			attempts: 3,
			f: func(callCount *int) error {
				*callCount++
				if *callCount < 2 {
					return fmt.Errorf("network timeout: %w", ErrTemporary)
				}
				return nil
			},
			wantErr:   false,
			wantCalls: 2,
		},
		{
			name:     "постоянная ошибка на 1-й попытке → досрочный выход",
			attempts: 5,
			f: func(callCount *int) error {
				*callCount++
				return fmt.Errorf("invalid credentials: %w", ErrPermanent)
			},
			wantErr:    true,
			wantCalls:  1,
			wantErrIs:  ErrPermanent,
			wantErrMsg: "после 1 попыток: invalid credentials: permanent error",
		},
		{
			name:     "постоянная ошибка на 2-й попытке",
			attempts: 5,
			f: func(callCount *int) error {
				*callCount++
				if *callCount < 2 {
					return fmt.Errorf("network timeout: %w", ErrTemporary)
				}
				return fmt.Errorf("database locked: %w", ErrPermanent)
			},
			wantErr:    true,
			wantCalls:  2,
			wantErrIs:  ErrPermanent,
			wantErrMsg: "после 2 попыток: database locked: permanent error",
		},
		{
			name:     "все попытки временные → последняя ошибка",
			attempts: 3,
			f: func(callCount *int) error {
				*callCount++
				return fmt.Errorf("connection refused: %w", ErrTemporary)
			},
			wantErr:    true,
			wantCalls:  3,
			wantErrIs:  ErrTemporary,
			wantErrMsg: "после 3 попыток: connection refused: temporary error",
		},
		{
			name:     "успех с первой попытки",
			attempts: 5,
			f: func(callCount *int) error {
				*callCount++
				return nil
			},
			wantErr:   false,
			wantCalls: 1,
		},
		{
			name:     "временная ошибка обёрнута несколько раз, затем успех",
			attempts: 3,
			f: func(callCount *int) error {
				*callCount++
				if *callCount < 3 {
					err := ErrTemporary
					err = fmt.Errorf("network: %w", err)
					err = fmt.Errorf("timeout: %w", err)
					return err
				}
				return nil
			},
			wantErr:   false,
			wantCalls: 3,
		},
		{
			name:     "все попытки временные, глубоко обёрнутые",
			attempts: 3,
			f: func(callCount *int) error {
				*callCount++
				err := ErrTemporary
				err = fmt.Errorf("network: %w", err)
				err = fmt.Errorf("timeout: %w", err)
				return err
			},
			wantErr:    true,
			wantCalls:  3,
			wantErrIs:  ErrTemporary,
			wantErrMsg: "после 3 попыток: timeout: network: temporary error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callCount := 0

			f := func() error {
				return tt.f(&callCount)
			}

			err := NewRetry(tt.attempts, f)

			if callCount != tt.wantCalls {
				t.Errorf("Retry() called %d times, want %d", callCount, tt.wantCalls)
			}

			if tt.wantErr && err == nil {
				t.Error("Retry() expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Retry() expected nil, got error: %v", err)
			}

			if tt.wantErrIs != nil && err != nil {
				if !errors.Is(err, tt.wantErrIs) {
					t.Errorf("errors.Is(%v, %v) = false, want true", err, tt.wantErrIs)
				}
			}

			if tt.wantErrMsg != "" && err != nil {
				if err.Error() != tt.wantErrMsg {
					t.Errorf("error message = %q, want %q", err.Error(), tt.wantErrMsg)
				}
			}
		})
	}
}
