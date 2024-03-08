package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LogLevel string

const (
	INFO    LogLevel = "INFO"
	ERROR   LogLevel = "ERROR"
	DEBUG   LogLevel = "DEBUG"
	WARNING LogLevel = "WARNING"
)

func main() {
	//open log file
	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("file Open error")
		return
	}
	//close file
	defer file.Close()

	//init log counter
	logCount := make(map[LogLevel]int)
	logCount[INFO] = 0
	logCount[ERROR] = 0
	logCount[DEBUG] = 0
	logCount[WARNING] = 0

	//read the text file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		//analyze log infomation
		files := strings.Fields(line)
		if len(files) < 3 {
			fmt.Println("Invaild log entry:", line)
			continue
		}

		//Extract log level and updates counters
		logLevel := LogLevel(files[1])
		logCount[logLevel]++
	}

	//print the log level count result
	fmt.Println("Log level counts:")
	fmt.Println("INFO:", logCount[INFO])
	fmt.Println("ERROR:", logCount[ERROR])
	fmt.Println("DEBUG:", logCount[DEBUG])
	fmt.Println("WARNING:", logCount[WARNING])

	// 检查扫描器是否遇到任何错误
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}

}
