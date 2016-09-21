package math

import "math"

// Round Round float (e.g. 0.5 = 1, 0.4 = 0).
// The parameters is a float64.
// It returns the value rounded
func Round(f float64) int {
	if math.Abs(f) < 0.5 {
		return 0
	}
	return int(f + math.Copysign(0.5, f))
}

// MaxFloat calculates the max value of the float.
// The parameters are two floats.
// It returns the max value.
func MaxFloat(a float64, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
