package utils

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func CheckMaxDecimalPlaces(m int, n float64) bool {
	mult := n * math.Pow10(m)
	roundedDown := float64(int64(mult))

	return mult-roundedDown == 0
}

func parseInt(s string) (int64, bool) {
	var fracXOk bool
	fracX, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		fracXOk = false
		log.Println(err)
	}
	fracXOk = true
	return fracX, fracXOk
}

func parseBytes(s string, n int) []byte {
	var b []byte = make([]byte, n)
	c := []byte(s)
	for i := range c {
		if c[i] >= 48 {
			b[i] = c[i] - 48
		}
	}
	return b
}

func AddFloats(a, b float64) float64 {
	x := strings.Split(strconv.FormatFloat(a, 'f', 3, 64), ".")
	y := strings.Split(strconv.FormatFloat(b, 'f', 3, 64), ".")

	var fracX []byte = parseBytes(x[1], 3)
	var fracY []byte = parseBytes(y[1], 3)

	var fracRes []byte = make([]byte, 3)
	var carry byte = 0
	for i := 2; i >= 0; i-- {
		a := fracX[i] + fracY[i] + carry
		if a > 9 {
			a %= 10
			carry = 1
		} else {
			carry = 0
		}
		fracRes[i] = a
	}
	for i := range fracRes {
		fracRes[i] = fracRes[i] + 48
	}

	expX, _ := parseInt(x[0])
	expY, _ := parseInt(y[0])

	expRes := expX + expY + int64(carry)
	res := fmt.Sprintf("%d.%s", expRes, fracRes)

	resF, err := strconv.ParseFloat(res, 64)
	if err != nil {
		log.Println(err)
	}

	return resF
}

func SubtractFloats(a, b float64) float64 {
	var flipped bool
	if a < b {
		c := a
		a = b
		b = c
		flipped = true
	}
	x := strings.Split(strconv.FormatFloat(a, 'f', 3, 64), ".")
	y := strings.Split(strconv.FormatFloat(b, 'f', 3, 64), ".")

	var fracX []byte = parseBytes(x[1], 3)
	var fracY []byte = parseBytes(y[1], 3)

	var fracRes []byte = make([]byte, 3)
	var carry byte = 0
	for i := 2; i >= 0; i-- {
		var xx byte = 0
		if carry == 1 {
			if fracX[i] > 0 && fracX[i] < fracY[i] {
				xx = 10 + fracX[i] - carry
				carry = 1
			} else if fracX[i] > 0 && fracX[i]-carry >= fracY[i] {
				xx = fracX[i] - carry
				carry = 0
			} else if fracX[i] > 0 && fracX[i]-carry < fracY[i] {
				xx = 10 + fracX[i] - carry
				carry = 1
			} else if fracX[i] == 0 {
				xx = 10 + fracX[i] - carry
				carry = 1
			}
		} else if carry == 0 {
			if fracX[i] < fracY[i] {
				xx = 10 + fracX[i]
				carry = 1
			} else if fracX[i] >= fracY[i] {
				xx = fracX[i]
			}
		}

		a := xx - fracY[i]
		fracRes[i] = a
	}
	for i := range fracRes {
		fracRes[i] = fracRes[i] + 48
	}

	expX, _ := parseInt(x[0])
	expY, _ := parseInt(y[0])

	expRes := expX - expY - int64(carry)
	res := fmt.Sprintf("%d.%s", expRes, fracRes)

	resF, err := strconv.ParseFloat(res, 64)
	if err != nil {
		log.Println(err)
	}

	if flipped {
		return resF * -1
	}
	return resF
}
