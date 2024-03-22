package main

import (
	"fmt"
	"io"
)

// func main() {
// 	data := "你好世界!"
// 	reader := strings.NewReader(data)

// 	readerAndPrint(reader)

// }

func readerAndPrintStirng(r io.Reader) {
	buffer := make([]byte, 8)
	for {
		n, err := r.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}
