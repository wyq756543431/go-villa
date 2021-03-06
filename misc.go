package villa

import (
	"fmt"
)

// N is a very long slice of elements with size of zero-bytes.
// It is useful for generating a serial of numbers in a for-range clause. e.g.
//    for i := range villa.N[:200] { ... }
// The above i in the above for clause will range from 0 to 200(exclusive).
var N [int(^uint(0) >> 1)]struct{}

// An variable of zero-size bytes
type Empty struct{}

/* NestedError is an error with current message and nested error */
type NestedError struct {
	// The messsage of this error
	Message string
	// The nested error
	Nested error
}

// Error implements error interface
func (err *NestedError) Error() string {
	if err.Nested == nil {
		return err.Message
	}
	return err.Message + ": " + err.Nested.Error()
}

/*
	Deepest returns the deepest non-*NestedError error, which is the original
  	error.
*/
func (err *NestedError) Deepest() error {
	for {
		if err.Nested == nil {
			return nil
		}

		ne, ok := err.Nested.(*NestedError)

		if !ok {
			return err.Nested
		}

		err = ne
	}
}

/*
	DeepestNested returns the deepest nested error. If err is not *NestedError,
	it is directly returned.
*/
func DeepestNested(err error) error {
	ne, ok := err.(*NestedError)
	if ok {
		err = ne.Deepest()
	}

	return err
}

/* NestErrorf returns a *NestedError error with a message */
func NestErrorf(err error, fmtstr string, args ...interface{}) *NestedError {
	return &NestedError{
		Message: fmt.Sprintf(fmtstr, args...),
		Nested:  err,
	}
}
