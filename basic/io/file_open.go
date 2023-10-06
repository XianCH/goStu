package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./file_open.go")
	if err != nil {
		fmt.Println("open file failed!,err:", err)
		return
	}
	file.Close()
}
