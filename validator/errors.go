package validator

import "errors"

var (
	ErrInvalidLength        = errors.New("کپ ملی باید ۱۰ رقم باشد")
	ErrInvalidCharacter     = errors.New("کد ملی باید فقط شامل ارقام باشد")
	ErrRepeatedDigits       = errors.New("کد ملی نمی‌تواند شامل ارقام تکراری باشد")
	ErrInvalidChecksum      = errors.New("کد ملی نامعتبر است")
	ErrInvalidIBANLength    = errors.New("طول IBAN نامعتبر است")
	ErrInvalidIBANCharacter = errors.New("IBAN شامل کاراکترهای نامعتبر است")
	ErrInvalidIBANChecksum  = errors.New("چک‌سام IBAN نامعتبر است")
)
