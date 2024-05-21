package errortest

import (
	"errors"
	"fmt"
)

func TestErrorJoin() {
	err1 := errors.New("error1")
	err2 := errors.New("error2")
	err3 := errors.New("error3")

	err := errors.Join(err1, err2, err3)
	fmt.Println(err)
	errs, ok := err.(interface{ Unwrap() []error })
	if !ok {
		fmt.Println("not imple Unwrap []error")
		return
	}
	fmt.Println(errs.Unwrap())
}
