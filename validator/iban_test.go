package validator

import "testing"

func TestValidateIBAN(t *testing.T) {
	tests := []struct {
		iban string
		want error
	}{
		{"IR620540105180021273113007", nil},
		{"IR1234567890123456789012", ErrInvalidIBANLength},
	}
	for _, tt := range tests {
		if got := ValidateIBAN(tt.iban); got != tt.want {
			t.Errorf("ValidateIBAN(%q) = %v, want %v", tt.iban, got, tt.want)
		}
	}
}
