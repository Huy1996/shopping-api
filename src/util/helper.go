package util

import "math"

func WithinTolerance(num1, num2, tolerance float64) bool {
	return math.Abs(num1-num2) < tolerance
}
