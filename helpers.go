package snils

import "slices"

func isDelimiter(r rune) bool {
	return slices.Contains(delimiters, r)
}
