package misc

import (
	"fmt"
	"math"

	"github.com/kristinjeanna/generative/errors"
)

// LastNIndexes returns n indexes up to and including the specified index.
// If fromIndex is less than n, only indexes 0 to fromIndex will be returned,
// thus this function may return fewer than n indexes.
//
// For example:
//
// LastNIndexes(97, 5) returns []int{93, 94, 95, 96, 97}
//
// LastNIndexes(2, 5) returns []int{0, 1, 2}
//
// Errors
//
// An error is returned (1) if n is less than or equal to zero, or
// (2) if fromIndex is less than zero.
func LastNIndexes(fromIndex, n int) ([]int, error) {
	if n <= 0 {
		return nil, fmt.Errorf(errors.MustBeGreaterThanZero, "n", "n", n)
	}
	if fromIndex < 0 {
		return nil, fmt.Errorf(errors.MustBeGreaterThanOrEqualToZero, "fromIndex", "fromIndex", fromIndex)
	}

	indexes := make([]int, 0)
	start := fromIndex - n + 1

	for i := start; i <= fromIndex; i++ {
		if i >= 0 {
			indexes = append(indexes, i)
		}
	}

	return indexes, nil
}

// LastNIndexesWrap returns n indexes up to and including the specified index.
// Wraps the indexes if fromIndex is less than n, thus this function always
// returns exactly n indexes.
//
// For example:
//
// LastNIndexesWrap(100, 5, 1000) returns []int{96, 97, 98, 99, 100}
//
// LastNIndexesWrap(1, 5, 100) returns []int{97, 98, 99, 0, 1}
//
// Errors
//
// An error is returned (1) if n is less than or equal to zero, or (2) if fromIndex
// or max is less than zero, or (3) if fromIndex is greater than or equal to max.
func LastNIndexesWrap(fromIndex, n, max int) ([]int, error) {
	if n <= 0 {
		return nil, fmt.Errorf(errors.MustBeGreaterThanZero, "n", "n", n)
	}
	if fromIndex < 0 {
		return nil, fmt.Errorf(errors.MustBeGreaterThanOrEqualToZero, "fromIndex", "fromIndex", fromIndex)
	}
	if max < 0 {
		return nil, fmt.Errorf(errors.MustBeGreaterThanOrEqualToZero, "max", "max", max)
	}
	if fromIndex >= max {
		return nil, fmt.Errorf("argument \"fromIndex\" must be less than argument \"max\"; fromIndex=%d; max=%d", fromIndex, max)
	}

	indexes := make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		idx := int(math.Mod(
			math.Abs(float64(fromIndex-i+max)),
			float64(max)))
		indexes = append(indexes, idx)
	}

	return indexes, nil
}
