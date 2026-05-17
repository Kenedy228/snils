package snils

import "errors"

var (
	ErrInvalidLength   = errors.New("неверная длина СНИЛС")
	ErrInvalidContent  = errors.New("СНИЛС должен содержать только цифры")
	ErrForbiddenSNILS  = errors.New("запрещенное значение СНИЛС")
	ErrInvalidChecksum = errors.New("неверная контрольная сумма СНИЛС")
)
