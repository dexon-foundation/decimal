package decimal

import "fmt"

// ErrorExponentLimit is returned when the decimal exponent exceed int32 range.
type ErrorExponentLimit struct {
	value string
}

// Error implements error interface.
func (e *ErrorExponentLimit) Error() string {
	return fmt.Sprintf("can't convert %s to decimal: fractional part too long", e.value)
}

// ErrorInvalidFormat is returned when the input string is not valid integer.
type ErrorInvalidFormat struct {
	reason string
}

// Error implements error interface.
func (e *ErrorInvalidFormat) Error() string {
	return e.reason
}

// ErrorInvalidType is returned when the value passed into sql.Scanner is not
// with expected type. (valid types: int64, float64, []byte, string)
type ErrorInvalidType struct {
	reason string
}

// Error implements error interface.
func (e *ErrorInvalidType) Error() string {
	return e.reason
}

func assertErrorInterface() {
	var _ error = (*ErrorExponentLimit)(nil)
	var _ error = (*ErrorInvalidFormat)(nil)
	var _ error = (*ErrorInvalidType)(nil)
}
