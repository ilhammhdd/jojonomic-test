package utils

import "math"

func CheckMaxDecimalPlaces(m int, n float64) bool {
	mult := n * math.Pow10(m)
	roundedDown := float64(int64(mult))

	return mult-roundedDown == 0
}
