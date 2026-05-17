package snils

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
