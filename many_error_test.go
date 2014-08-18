package ears

import (
	"errors"
	"fmt"
)

func ExampleManyError() {
	C := errors.New
	e := NewMany()

	e.Set(New(C("cant make grouping")))
	e.Set(nil)
	e.Set(New(C("cant process post")))
	e.Set(C("cant store match"))
	fmt.Println(e)

	// Output:
	// multiple errors:
	//   ears/many_error_test.go:12 cant make grouping
	//   ears/many_error_test.go:14 cant process post
	//   cant store match
}
