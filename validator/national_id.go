package validator

func ValidateNationalID(id string) error {

	if len(id) != 10 {
		return ErrInvalidLength
	}

	if charsAreInvalid(id) {
		return ErrInvalidCharacter
	}

	if charsAreRepeated(id) {
		return ErrRepeatedDigits
	}

	if checksumInvalid(id) {
		return ErrInvalidChecksum
	}

	return nil
}

func charsAreInvalid(id string) bool {
	for _, c := range id {
		if c < '0' || c > '9' {
			return true
		}
	}
	return false
}

func charsAreRepeated(id string) bool {
	for i := 1; i < len(id); i++ {
		if id[i] != id[0] {
			return false
		}
	}
	return true
}

func checksumInvalid(id string) bool {
	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(id[i]-'0') * (10 - i)
	}

	remainder := sum % 11
	check := int(id[9] - '0')

	if remainder < 2 {
		return check != remainder
	}

	return check != (11 - remainder)
}
