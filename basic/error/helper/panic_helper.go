package helper

import "fmt"
import errortest "github.com/x14n/goStu/basic/error"

func PanicHelper() {
	if r := recover(); r != nil {
		err := errortest.Wrap(fmt.Errorf("%v", r), "Recovered from panic")
		fmt.Println(err)
		if customErr, ok := err.(*errortest.CustomError); ok {
			fmt.Printf("Error: %s\nCode: %d\nCallers:\n", customErr.Error(), customErr.Code())
			for _, caller := range customErr.Caller() {
				fmt.Printf("  %s:%d %s\n", caller.FileName, caller.FileLine, caller.FuncName)
			}
			if len(customErr.Wrapped()) > 0 {
				fmt.Printf("Wrapped errors:\n")
				for _, wrappedErr := range customErr.Wrapped() {
					fmt.Printf("  %s\n", wrappedErr.Error())
				}
			}
		}
	}
}
