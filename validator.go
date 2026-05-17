package snils

// Validate выполняет полную проверку нормализованного СНИЛС.
//
// Ожидается, что входная строка уже приведена к нормализованному виду
// (например, с помощью функции Normalize) и содержит только цифры.
//
// Проверки выполняются в следующем порядке:
//   - длина СНИЛС;
//   - допустимость символов (только цифры);
//   - проверка на запрещенное значение;
//   - проверка контрольного числа (если применимо);
//
// Возвращает одну из ошибок:
//   - ErrInvalidLength;
//   - ErrInvalidContent;
//   - ErrForbiddenSNILS;
//   - ErrInvalidChecksum.
//
// При успешной валидации возвращает nil.
func Validate(normalizedSnils string) error {
	if err := validateLength(normalizedSnils); err != nil {
		return err
	}

	if err := validateDigitsOnly(normalizedSnils); err != nil {
		return err
	}

	if err := validateForbiddenSNILS(normalizedSnils); err != nil {
		return err
	}

	if shouldValidateChecksum(normalizedSnils) {
		if err := validateChecksum(normalizedSnils); err != nil {
			return err
		}
	}

	return nil
}
