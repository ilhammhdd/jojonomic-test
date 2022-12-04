package utils

import (
	"math"
)

func CheckMaxDecimalPlaces(m int, n float64) bool {
	mult := n * math.Pow10(m)
	roundedDown := float64(int64(mult))

	return mult-roundedDown == 0
}

func AddWithDecimalPlaces(m int, a, b float64) float64 {
	power := math.Pow10(m)
	aInt := uint64(a * power)
	bInt := uint64(b * power)
	res := aInt + bInt
	exp := res / uint64(power)
	mantissa := res % uint64(power)
	return float64(exp) + float64(mantissa)*(1/power)
}
