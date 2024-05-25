package helper

import (
	"fmt"
	"log"

	errortest "github.com/x14n/goStu/basic/error"
)

func ErrorHelper(err error) {
	if customErr, ok := err.(errortest.Error); ok {
		fmt.Println("Error:", customErr)
		fmt.Println("Error Code:", customErr.Code())
		fmt.Println("Callers:")
		for _, caller := range customErr.Caller() {
			fmt.Printf("  %s:%d %s\n", caller.FileName, caller.FileLine, caller.FuncName)
		}
		log.Println("Wrapped Error:", customErr.Wrapped())
	} else {
		fmt.Println("Error:", err)
	}
}
