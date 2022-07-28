package errors

const (
	MustBeGreaterThanZeroDecimal          string = "argument %q must be greater than 0; %s=%d"
	MustBeGreaterThanZeroFloat            string = "argument %q must be greater than 0; %s=%f"
	MustBeGreaterThanOrEqualToZeroDecimal string = "argument %q must be greater than or equal to 0; %s=%d"
	MustBeGreaterThanOrEqualToZeroFloat   string = "argument %q must be greater than or equal to 0; %s=%f"

	ExpectedError string = "expected an error, but got nil"
)
