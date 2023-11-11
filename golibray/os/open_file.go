package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	data := []byte("hello world \n")

	_, err = file.Write(data)
	if err != nil {
		log.Fatal("file writer faild", err)
		return
	}

	file.Close()

	file, err = os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 100)

	n, err := file.Read(buffer)
	if err != nil {
		log.Println("file read false", err)
		return
	}

	fmt.Println("file data from file.txt", string(buffer[:n]))
}
