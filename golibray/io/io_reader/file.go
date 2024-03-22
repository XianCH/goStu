package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
	}
	ReadAndPrintFIle(file)
}

func ReadAndPrintFIle(r io.Reader) {
	buffer := make([]byte, 100)
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
