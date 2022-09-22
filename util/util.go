package util

import (
	"math"
)

// Math

func Min(min, val int) int {
	if val > min {
		return min
	}

	return val
}

func Max(max, val int) int {
	if val < max {
		return max
	}

	return val
}

func Clamp(min, max, val int) int {
	if val <= min {
		return min
	}

	if val >= max {
		return max
	}

	return val
}

func Abs(val int) int {
	if val > 0 {
		return val
	}

	return int(math.Abs(float64(val)))
}
