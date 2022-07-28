package mathextra

import (
	"math"
)

const (
	Pi        float64 = math.Pi
	TwoPi             = Pi * 2
	HalfPi            = Pi / 2
	QuarterPi         = Pi / 4
)

// Floor returns the math.Floor when v is positive, and returns the math.Ceil
// when v is negative. Stated differently, Floor returns the largest mathematical
// integer less than v when v is positive and the smallest mathematical integer
// greater than v when v is negative.
func Floor(v float64) float64 {
	switch {
	case v < 0:
		return math.Ceil(v)
	default:
		return math.Floor(v)
	}
}

// Normalize normalizes an angle theta to a value from 0 to π (inclusive) for
// positive angles and 0 to -π (exclusive) for negative angles. -π is always
// returned as π.
func Normalize(theta float64) float64 {
	multiples := Floor(theta / TwoPi)
	theta -= multiples * TwoPi

	var result float64
	if theta > 0 {
		if theta > Pi {
			// if greater than π, subtract 2π
			result = theta - TwoPi
		} else {
			result = theta
		}
	} else {
		if theta < -Pi {
			// if less than -π, add 2π
			result = theta + TwoPi
		} else {
			result = theta
		}
	}

	// favor π over -π
	if result == -Pi {
		return Pi
	}
	return result
}

func ToRadians(degrees float64) float64 {
	return Normalize(math.Mod(degrees, 360.0) / 360.0 * TwoPi)
}
