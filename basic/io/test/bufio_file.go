package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func WriteFile(filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		log.Println("FIle open error", err)
		return
	}

	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		writer.WriteString("hello world \n")
	}
	writer.Flush()
}

func ReadFile(filename string) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		log.Println("open file falus", err)
		return
	}
	reader := bufio.NewReader(file)
	for {
		line, _, err2 := reader.ReadLine()
		if err2 == io.EOF {
			break
		}
		if err2 != nil {
			log.Println("in read file", err2)
			return
		}
		log.Println(string(line))
	}
}

func main() {
	WriteFile("./xxxx.txt")
	ReadFile("./xxxx.txt")
}
