package consterr

// Error is a constant error type.
type Error string

func (e Error) Error() string {
	return string(e)
}
