package validator

import "strings"

func ValidateIBAN(iban string) error {
	iban = strings.ToUpper(strings.ReplaceAll(iban, " ", ""))
	if len(iban) != 26 || iban[:2] != "IR" {
		return ErrInvalidIBANLength
	}
	rearranged := iban[4:] + iban[:4]
	rem := 0
	for _, c := range rearranged {
		var val int
		if c >= '0' && c <= '9' {
			val = int(c - '0')
		} else if c >= 'A' && c <= 'Z' {
			val = int(c - 'A' + 10)
		} else {
			return ErrInvalidIBANCharacter
		}
		rem = (rem*10 + val) % 97
	}

	if rem != 1 {
		return ErrInvalidIBANChecksum
	}

	return nil
}
