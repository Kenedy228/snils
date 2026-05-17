package snils

const (
	SNILSLength          = 11
	ForbiddenSNILS       = "00000000000"
	MinNumberForChecksum = 1001998
)

var (
	snilsChecksumWeights = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
)
