package validator

import "errors"

var (
	ErrInvalidLength        = errors.New("national id length must be 10 digits")
	ErrInvalidCharacter     = errors.New("national id must contain only numeric characters")
	ErrRepeatedDigits       = errors.New("national id cannot be all repeated digits")
	ErrInvalidChecksum      = errors.New("national id checksum is invalid")
	ErrInvalidIBANLength    = errors.New("IBAN length is invalid")
	ErrInvalidIBANCharacter = errors.New("IBAN contains invalid characters")
	ErrInvalidIBANChecksum  = errors.New("IBAN checksum is invalid")
)
