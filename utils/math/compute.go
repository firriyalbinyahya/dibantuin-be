package math

import "math"

func RoundToTwoDecimalPlaces(value float64) float64 {
	// Membulatkan nilai ke dua tempat desimal
	return math.Round(value*100) / 100
}
