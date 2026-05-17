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
