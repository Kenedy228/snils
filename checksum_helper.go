package snils

import "strconv"

// calculateWeightedSum вычисляет взвешенную сумму первых девяти цифр СНИЛС.
//
// Каждая цифра умножается на соответствующий вес из snilsChecksumWeights,
// после чего результаты суммируются.
func calculateWeightedSum(numberPart string) int {
	sum := 0

	for i, weight := range snilsChecksumWeights {
		digit := int(numberPart[i] - '0')
		sum += digit * weight
	}

	return sum
}

// calculateControlNumber вычисляет контрольное число СНИЛС
// на основе взвешенной суммы.
//
// Алгоритм:
//   - если сумма < 100, контрольное число равно сумме;
//   - если сумма == 100, контрольное число равно 0;
//   - если сумма > 100:
//   - вычисляется остаток от деления на 101;
//   - если остаток == 100, контрольное число равно 0;
//   - иначе контрольное число равно остатку.
func calculateControlNumber(weightedSum int) int {
	if weightedSum < 100 {
		return weightedSum
	}

	if weightedSum == 100 {
		return 0
	}

	remainder := weightedSum % 101
	if remainder == 100 {
		return 0
	}

	return remainder
}

// parseControlNumber преобразует контрольную часть СНИЛС в число.
//
// Предполагается, что входная строка содержит только цифры.
// В случае некорректного значения функция вызывает panic.
func parseControlNumber(controlPart string) int {
	parsed, err := strconv.Atoi(controlPart)
	if err != nil {
		panic("контрольная часть СНИЛС должна состоять только из цифр")
	}

	return parsed
}
