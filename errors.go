package snils

import "errors"

var (
	// ErrInvalidLength возвращается, если длина СНИЛС не равна SNILSLength.
	ErrInvalidLength = errors.New("неверная длина СНИЛС")

	// ErrInvalidContent возвращается, если СНИЛС содержит символы,
	// отличные от цифр.
	ErrInvalidContent = errors.New("СНИЛС должен содержать только цифры")

	// ErrForbiddenSNILS возвращается, если СНИЛС равен запрещённому значению.
	ErrForbiddenSNILS = errors.New("запрещенное значение СНИЛС")

	// ErrInvalidChecksum возвращается, если контрольное число СНИЛС
	// не соответствует вычисленному значению.
	ErrInvalidChecksum = errors.New("неверная контрольная сумма СНИЛС")
)
