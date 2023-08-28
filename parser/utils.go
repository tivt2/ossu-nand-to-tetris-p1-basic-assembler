package parser

import "math"

func decimalToBinary(num int) string {
	out := []byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0', '0'}
	for num > 0{
		if num == 1 {
			out[15] = '1'
			break
		}

		log2 := int(math.Floor(math.Log2(float64(num))))
		out[15 - log2] = '1'
		num -= int(math.Pow(2, float64(log2)))
	}
	return string(out)
}