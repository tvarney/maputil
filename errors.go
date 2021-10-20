package maputil

import (
	"fmt"
	"strings"

	"github.com/tvarney/maputil/consterr"
)

const (
	// ErrInvalidType is the root type error.
	ErrInvalidType consterr.Error = "invalid type"

	// ErrInvalidValue is the root value error.
	ErrInvalidValue consterr.Error = "invalid value"

	// ErrMissingRequiredValue is the root error of a missing required value.
	ErrMissingRequiredValue consterr.Error = "missing required value"
)

// InvalidTypeError is an error indicating that a type did not match the
// expected type.
type InvalidTypeError struct {
	Expected []string
	Actual   string
}

// Error returns the string representation of this invalid type error.
func (e InvalidTypeError) Error() string {
	switch len(e.Expected) {
	case 0:
		return string(ErrInvalidType) + " " + e.Actual
	case 1:
		return string(ErrInvalidType) + " " + e.Actual + "; expected " + e.Expected[0]
	case 2:
		return fmt.Sprintf(string(ErrInvalidType)+" %s; expected %s or %s", e.Actual, e.Expected[0], e.Expected[1])
	}
	// Calculate base capacity
	c := len(string(ErrInvalidType)) + len(e.Actual) + 15 + 2*(len(e.Expected)-1)
	for _, v := range e.Expected {
		c += len(v)
	}
	b := &strings.Builder{}
	b.Grow(c)
	b.WriteString(string(ErrInvalidType) + " ")
	b.WriteString(e.Actual)
	b.WriteString("; expected ")
	for _, v := range e.Expected[:len(e.Expected)-1] {
		b.WriteString(v)
		b.WriteString(", ")
	}
	b.WriteString("or ")
	b.WriteString(e.Expected[len(e.Expected)-1])
	return b.String()
}

// Unwrap returns the parent error of this type error.
func (e InvalidTypeError) Unwrap() error {
	return ErrInvalidType
}

// EnumStringError is an error indicating that a value was not one of the
// accepted enum values.
type EnumStringError struct {
	Value string
	Enum  []string
}

// Error returns the string representation of this enum error.
func (e EnumStringError) Error() string {
	switch len(e.Enum) {
	case 0:
		return fmt.Sprintf(string(ErrInvalidValue)+" %q", e.Value)
	case 1:
		return fmt.Sprintf(string(ErrInvalidValue)+" %q; expected %q", e.Value, e.Enum[0])
	case 2:
		return fmt.Sprintf(string(ErrInvalidValue)+" %q; expected %q or %q", e.Value, e.Enum[0], e.Enum[1])
	}
	b := &strings.Builder{}
	fmt.Fprintf(b, string(ErrInvalidValue)+" %q; expected one of ", e.Value)
	for _, v := range e.Enum[:len(e.Enum)-1] {
		fmt.Fprintf(b, "%q, ", v)
	}
	fmt.Fprintf(b, "or %q", e.Enum[len(e.Enum)-1])
	return b.String()
}

// Unwrap returns the parent error for this enum error.
func (e EnumStringError) Unwrap() error {
	return ErrInvalidValue
}

// MissingRequiredValueError is an error indicating that the given value was
// not present.
type MissingRequiredValueError struct {
	Key string
}

// Error returns the string representation of this missing required value error.
func (e MissingRequiredValueError) Error() string {
	return fmt.Sprintf("%s %q", string(ErrMissingRequiredValue), e.Key)
}

// Unwrap returns the parent error for this missing required value error.
func (e MissingRequiredValueError) Unwrap() error {
	return ErrMissingRequiredValue
}
