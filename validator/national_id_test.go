package validator

import "testing"

func TestValidateNationalID(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect error
	}{
		{
			name:   "valid national id",
			input:  "0081013434",
			expect: nil,
		},
		{
			name:   "invalid checksum",
			input:  "0013520839",
			expect: ErrInvalidChecksum,
		},
		{
			name:   "invalid length",
			input:  "12345",
			expect: ErrInvalidLength,
		},
		{
			name:   "non-numeric",
			input:  "12A4567890",
			expect: ErrInvalidCharacter,
		},
		{
			name:   "repeated digits",
			input:  "1111111111",
			expect: ErrRepeatedDigits,
		},
		{
			name:   "invalid checksum",
			input:  "0013520831",
			expect: ErrInvalidChecksum,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateNationalID(tt.input)
			if got != tt.expect {
				t.Errorf("ValidateNationalID(%s) = %v; expected %v", tt.input, got, tt.expect)
			}
		})
	}
}
