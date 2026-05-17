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
	numericPart := snils[:9]
	checksumPart := snils[9:]

	var checksum int64

	for i, w := range snilsChecksumWeights {
		d := int(numericPart[i] - '0')
		checksum += int64(d * w)
	}

	var controlNumber int64

	if checksum < 100 {
		controlNumber = checksum
	} else if checksum == 100 {
		controlNumber = 0
	} else {
		remainder := checksum % 101
		if remainder == 100 {
			controlNumber = 0
		} else {
			controlNumber = remainder
		}
	}

	conv, _ := strconv.ParseInt(checksumPart, 10, 64)

	if conv != controlNumber {
		return fmt.Errorf("")
	}

	return nil
}
