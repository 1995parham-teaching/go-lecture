package main

import (
	"errors"
	"fmt"
	"strconv"
)

//
// type error interface {
//   Error() string
// }
//
// when there is no implementation behind an interface, it will be nil.
// so we will check err with nil.

type My3Error struct {
	LastError error
}

func (me My3Error) Unwrap() error {
	return me.LastError
}

type My1Error struct {
	Message string
	Number  int
}

func (me My1Error) Error() string {
	return fmt.Sprintf("%s: %d", me.Message, me.Number)
}

// errors should prefix with Err when they are created with
// errors.New() and they are not structure.
var ErrMy2 = errors.New("i am error 2")

func toNumber(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}

	return int(i)
}

func iReturnError(n int) (int, error) {
	if n%2 == 0 {
		return 0, My1Error{
			Message: "Hello, I am error number 1",
			Number:  n,
		}
	}

	return 1, ErrMy2
}

// iReturnErrorError calls iReturnError and wraps its error.
func iReturnErrorError(n int) error {
	_, err := iReturnError(n)
	if err != nil {
		return fmt.Errorf("from iReturnErrorError %w", err)
	}

	return nil
}

func main() {
	fmt.Println(toNumber("123"))
	fmt.Println(toNumber("abc"))
	fmt.Println(toNumber("123abc"))

	// why MyError2 without {} but MyError1 with {}.
	// also pay attention to the error parameters because they also must be equal.
	// the reason behind this is because of the way that we write Error method.
	fmt.Printf("iReturnErrorError(1) is MyError2? %t\n", errors.Is(iReturnErrorError(1), ErrMy2))
	fmt.Printf("iReturnErrorError(0) is MyError1? %t\n", errors.Is(iReturnErrorError(0), My1Error{
		Message: "Hello, I am error number 1",
		Number:  0,
	}))
	fmt.Printf("iReturnErrorError(2) is MyError1? %t\n", errors.Is(iReturnErrorError(2), My1Error{
		Message: "Hello, I am error number 1",
		Number:  0,
	}))

	// parses the returned error from iReturnErrorError as My1Error
	// so we can get access to its details.
	//
	// errors.AsType[E] (new in Go 1.26) is the generic, type-safe replacement
	// for the older errors.As(err, &target) pattern — no need to allocate a
	// zero value first and no out-parameter.
	if myError1, ok := errors.AsType[My1Error](iReturnErrorError(0)); ok {
		fmt.Printf("we have My1Error from iReturnErrorError (%s, %d)\n", myError1.Message, myError1.Number)
	}
}
