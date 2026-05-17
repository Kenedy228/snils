package snils

import (
	"fmt"
	"unicode/utf8"
)

func validateLength(snils string) error {
	if utf8.RuneCountInString(snils) != SNILSLength {
		return fmt.Errorf("")
	}

	return nil
}
