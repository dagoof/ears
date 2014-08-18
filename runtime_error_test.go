package ears

import (
	"errors"
	"fmt"
)

func ExampleRuntimeError() {
	fmt.Println(New(errors.New("a very brittle test")))

	// Output:
	// ears/runtime_error_test.go:9 a very brittle test
}
