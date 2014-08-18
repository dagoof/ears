package ears

// LatestError exposes a non-nil error
type LatestError struct {
	error
}

// Err pulls the current error from our eventualError. Implements EventualError.
func (e *LatestError) Err() error {
	return e.error
}

// Set sets the latest error's latest error to err if it is not nil, and returns
// whether the error was set or not.
func (e *LatestError) Set(err error) bool {
	if err != nil {
		e.error = err
		return true
	}
	return false
}

// NewLatest creates an empty LatestError
func NewLatest() *LatestError {
	return &LatestError{}
}
