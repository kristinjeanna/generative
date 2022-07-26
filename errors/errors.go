package errors

const (
	MustBeGreaterThanZero          string = "argument %q must be greater than 0; %s=%d"
	MustBeGreaterThanOrEqualToZero string = "argument %q must be greater than or equal to 0; %s=%d"

	ExpectedError string = "expected an error, but got nil"
)
