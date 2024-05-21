package errortest

import (
	"errors"
	"fmt"
)

func TestErrorChain() {
	err1 := errors.New("error1")

	err2 := errors.New("error2")

	err3 := errors.New("error3")

	err := fmt.Errorf("warp multiple error: %w,%w,%w", err1, err2, err3)
	fmt.Println(err)
	e, ok := err.(interface{ Unwrap() []error })
	if !ok {
		fmt.Println("not imple Unwarp []error")
		return
	}
	fmt.Println(e.Unwrap())
}
