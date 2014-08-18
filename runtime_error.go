/*
Package ears exposes a few ways of wrapping errors to make them easier to
collect and interpret across multiple goroutines.
*/
package ears

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// RuntimeError marks the file and line number that the error was generated on.
// This can be potentially expensive and should be used sparingly.
type RuntimeError struct {
	error
	file string
	line int
}

func (e *RuntimeError) Error() string {
	dir, file := filepath.Split(e.file)
	pkg := filepath.Base(dir)

	return fmt.Sprintf("%s/%s:%d %s", pkg, file, e.line, e.error)
}

// New creates a new RuntimeError from a given error. The source of the error is
// marked when new is called. Because this makes a call to runtime.Caller, it is
// potentially costly and this error type should be used sparingly.
func New(err error) error {
	if err == nil {
		return err
	}

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}

	return &RuntimeError{
		err,
		file,
		line,
	}
}
