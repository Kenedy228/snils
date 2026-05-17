package snils

// isDelimiter проверяет, является ли символ допустимым разделителем СНИЛС.
func isDelimiter(r rune) bool {
	return r == '-'
}

// isSpace проверяет, является ли символ пробелом.
func isSpace(r rune) bool {
	return r == ' '
}

// isDigit проверяет, является ли символ десятичной цифрой ('0'–'9').
func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
