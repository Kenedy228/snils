package snils

func isDelimiter(r rune) bool {
	return r == '-'
}

func isSpace(r rune) bool {
	return r == ' '
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
