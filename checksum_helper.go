package snils

func calculateWeightedSum(numberPart string) int {
	sum := 0

	for i, weight := range snilsChecksumWeights {
		digit := int(numberPart[i] - '0')
		sum += digit * weight
	}

	return sum
}

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
