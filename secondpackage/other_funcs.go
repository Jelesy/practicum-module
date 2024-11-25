package secondpackage

import "math"

func AverageDigit(s []int) any {
	avr := 0
	for _, v := range s {
		avr += v
	}
	return math.Round(float64(avr / len(s)))
}
