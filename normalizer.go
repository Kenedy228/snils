package snils

// Normalize приводит СНИЛС к нормализованному виду.
//
// Функция удаляет пробелы и допустимые разделители,
// возвращая строку, состоящую только из символов исходного ввода,
// не являющихся пробелами или разделителями.
//
// Normalize не выполняет валидацию содержимого СНИЛС.
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
