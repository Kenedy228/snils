package snils

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func validateLength(snils string) error {
	rc := utf8.RuneCountInString(snils)

	if rc != SNILSLength {
		return fmt.Errorf("%w: длина %d символов, ожидалось %d", ErrInvalidLength, rc, SNILSLength)
	}

	return nil
}

func validateDigitsOnly(snils string) error {
	runes := []rune(snils)

	for _, r := range runes {
		if isDigit(r) {
			continue
		}

		return fmt.Errorf("%w: запрещенный символ %q, разрешены только цифры", ErrInvalidContent, string(r))
	}

	return nil
}

func validateForbiddenSNILS(snils string) error {
	if snils == ForbiddenSNILS {
		return fmt.Errorf("%w", ErrForbiddenSNILS)
	}

	return nil
}

func shouldValidateChecksum(snils string) bool {
	numberPart := snils[:9]
	number, err := strconv.Atoi(numberPart)

	if err != nil {
		panic("числовая часть СНИЛС должна состоять только из цифр")
	}

	if number > MinNumberForChecksum {
		return true
	}

	return false
}

func validateChecksum(snils string) error {
	numberPart := snils[:9]
	controlPart := snils[9:]

	expectedControlNumber := calculateControlNumber(
		calculateWeightedSum(numberPart),
	)

	actualControlNumber := parseControlNumber(controlPart)

	if actualControlNumber != expectedControlNumber {
		return fmt.Errorf("%w: ожидалось %02d, получено %02d", ErrInvalidChecksum, expectedControlNumber, actualControlNumber)
	}

	return nil
}
