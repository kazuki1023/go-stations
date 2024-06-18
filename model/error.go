package model

// ErrNotFound represents a not found error.
type ErrNotFound struct{}

func (e *ErrNotFound) Error() string {
	return "not found"
}

// NewErrNotFound returns a new instance of ErrNotFound.
func NewErrNotFound() error {
	return &ErrNotFound{}
}
