package snils

func Normalize(snils string) string {
	runes := make([]rune, 0, len(snils))

	for _, r := range snils {
		if isSpace(r) || isDelimiter(r) {
			continue
		}

		runes = append(runes, r)
	}

	return string(runes)
}
