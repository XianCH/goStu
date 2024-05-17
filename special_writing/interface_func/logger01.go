package main

import (
	"log"
	"os"
)

type Logger interface {
	Log(message string) error
}
type LoggerFunc func(message string) error

func (l LoggerFunc) Log(message string) error {
	return l(message)
}

func Logger01Test() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}

	logger := LoggerFunc(func(message string) error {
		_, err := file.WriteString(message + "/n")
		return err
	})

	err = logger("hello")
	if err != nil {
		log.Fatal(err)
	}
}
