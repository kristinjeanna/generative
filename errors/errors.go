package errors

import "fmt"

const (
	MustBeGreaterThanZeroDecimal          string = "argument %q must be greater than 0; %s=%d"
	MustBeGreaterThanZeroFloat            string = "argument %q must be greater than 0; %s=%f"
	MustBeGreaterThanOrEqualToZeroDecimal string = "argument %q must be greater than or equal to 0; %s=%d"
	MustBeGreaterThanOrEqualToZeroFloat   string = "argument %q must be greater than or equal to 0; %s=%f"

	mustBeGreaterThanFloat   string = "argument %q must be greater than argument %q; %s=%f, %s=%f"
	mustBeGreaterThanDecimal string = "argument %q must be greater than argument %q; %s=%d, %s=%d"

	ExpectedError string = "expected an error, but got nil"
)

// TODO: consider moving type constraints to their own package if needed elsewhere besides this package

// MustBeFloat provides a type constraint for all float types.
type MustBeFloat interface {
	float64 | float32
}

// MustBeDecimal provides a type constraint for all integer types.
type MustBeDecimal interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// MustBeGreaterThanFloat returns an error message for float values
// where the first value must be greater than the second value.
// This is just a convenience method for creating the error instance.
func MustBeGreaterThanFloat[T MustBeFloat](name1 string, value1 T, name2 string, value2 T) error {
	return fmt.Errorf(mustBeGreaterThanFloat, name1, name2, name1, value1, name2, value2)
}

// MustBeGreaterThanDecimal returns an error message for decimal values
// where the first value must be greater than the second value.
// This is just a convenience method for creating the error instance.
func MustBeGreaterThanDecimal[T MustBeDecimal](name1 string, value1 T, name2 string, value2 T) error {
	return fmt.Errorf(mustBeGreaterThanDecimal, name1, name2, name1, value1, name2, value2)
}
