package snils

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func validateLength(snils string) error {
	if utf8.RuneCountInString(snils) != SNILSLength {
		return fmt.Errorf("")
	}

	return nil
}

func validateDigitsOnly(snils string) error {
	runes := []rune(snils)

	for _, r := range runes {
		if isDigit(r) {
			continue
		}

		return fmt.Errorf("")
	}

	return nil
}

func validateForbiddenSNILS(snils string) error {
	if snils == ForbiddenSNILS {
		return fmt.Errorf("")
	}

	return nil
}

func shouldValidateChecksum(snils string) bool {
	conv, err := strconv.ParseInt(snils, 10, 64)
	if err != nil {
		panic("")
	}

	if conv > MinNumberForChecksum {
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
		return fmt.Errorf("")
	}

	return nil
}
