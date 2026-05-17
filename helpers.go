package snils

import "slices"

func isDelimiter(r rune) bool {
	return slices.Contains(delimiters, r)
}

func isSpace(r rune) bool {
	return r == ' '
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
