package ears

import (
	"fmt"
	"strings"
)

// ManyError contains any number of non-nil errors.
type ManyError struct {
	errs []error
}

// Err pulls the current error from our eventualError. Implements EventualError.
func (e *ManyError) Err() error {
	if len(e.errs) > 0 {
		return e
	}

	return nil
}

// Set adds the error to the list of errors if it is not nil, and returns
// whether the error was set or not.
func (e *ManyError) Set(err error) bool {
	if err != nil {
		e.errs = append(e.errs, err)
		return true
	}
	return false
}

func (e *ManyError) Error() string {
	if len(e.errs) == 1 {
		return e.errs[0].Error()
	}

	var ss []string

	for _, err := range e.errs {
		ss = append(ss, "  "+err.Error())
	}

	return fmt.Sprintf(
		"multiple errors:\n%s",
		strings.Join(ss, "\n"),
	)
}

// NewMany creates an empty ManyError
func NewMany(errs ...error) *ManyError {
	many := &ManyError{}

	for _, err := range errs {
		many.Set(err)
	}

	return many
}
