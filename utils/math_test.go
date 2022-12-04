package utils_test

import (
	"testing"

	"github.com/ilhammhdd/jojonomic_test/utils"
)

var addWithDecimalPlacesCases [][]float64 = [][]float64{
	{0.005, 0.1, 0.105},
	{0.1, 0.005, 0.105},
	{0.002, 0.1, 0.102},
	{0.1, 0.002, 0.102},
	{6, 0.007, 6.007},
	{0.007, 6, 6.007},
	{0.02, 0.1, 0.12},
	{0.9, 0.05, 0.95},
	{0.007, 0.004, 0.011},
	{0.099, 0.067, 0.166},
	{0.99, 0.067, 1.057},
	{2.542, 0.42, 2.962},
	{8.397, 0.604, 9.001},
	{2.542, 0.4, 2.942},
	{0.3, 4.3, 4.6},
	{1, 1, 2},
	{6, 9, 15},
	{2314, 542, 2856},
}

func TestAddWithDecimalPlaces(t *testing.T) {
	for i := range addWithDecimalPlacesCases {
		res := utils.AddFloats(addWithDecimalPlacesCases[i][0], addWithDecimalPlacesCases[i][1])
		if res != addWithDecimalPlacesCases[i][2] {
			t.Errorf("test case %d, expected: %v, got: %v", i, addWithDecimalPlacesCases[i][2], res)
		}
	}
}

var subtractFloatsCases [][]float64 = [][]float64{
	{1.001, 0.199, 0.802},
	{0.861, 0.719, 0.142},
	{0.876, 0.984, -0.108},
	{0.001, 0.999, -0.998},
	{31.143, 25.856, 5.287},
	{76.871, 892.984, -816.113},
	{13.726, 0.727, 12.999},
}

func TestSubtractFloats(t *testing.T) {
	for i := range subtractFloatsCases {
		res := utils.SubtractFloats(subtractFloatsCases[i][0], subtractFloatsCases[i][1])
		if res != subtractFloatsCases[i][2] {
			t.Errorf("test case %d, expected: %v, got: %v", i, subtractFloatsCases[i][2], res)
		}
	}
}
