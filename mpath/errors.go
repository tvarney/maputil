package mpath

import (
	"fmt"

	"github.com/tvarney/maputil/consterr"
)

const (
	// ErrBadRange is an error indicating that a range value was invalid.
	ErrBadRange consterr.Error = "invalid range"

	// ErrBadIndex is an error indicating that an index value was invalid.
	ErrBadIndex consterr.Error = "invalid index"

	// ErrUnmatchedOpenBracket is an error indicating that an open bracket in
	// a path was never terminated.
	ErrUnmatchedOpenBracket consterr.Error = "unmatched open bracket '['"

	// ErrUnmatchedCloseBracket is an error indicating that a close bracket in
	// a path was encountered prior to an open bracket.
	ErrUnmatchedCloseBracket consterr.Error = "unmatched closing bracket ']'"

	// ErrInvalidEscape is an error indicating that an escape character was
	// encountered at the end of the string.
	ErrInvalidEscape consterr.Error = "invalid escape"

	// ErrMissingSep is an error indicating that a path separator was not
	// found while parsing.
	ErrMissingSep consterr.Error = "missing separator"
)

// BadRangeStartError is an error indicating that a range start value was
// invalid.
type BadRangeStartError struct {
	Value string
}

func (e BadRangeStartError) Error() string {
	return fmt.Sprintf(string(ErrBadRange)+"; invalid start value %q", e.Value)
}

func (e BadRangeStartError) Unwrap() error {
	return ErrBadRange
}

// BadRangeEndError is an error indicating that a range end value was invalid.
type BadRangeEndError struct {
	Value string
}

func (e BadRangeEndError) Error() string {
	return fmt.Sprintf(string(ErrBadRange)+"; invalid end value %q", e.Value)
}

func (e BadRangeEndError) Unwrap() error {
	return ErrBadRange
}
