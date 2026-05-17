package snils

func calculateWeightedSum(numberPart string) int {
	sum := 0

	for i, weight := range snilsChecksumWeights {
		digit := int(numberPart[i] - '0')
		sum += digit * weight
	}

	return sum
}
