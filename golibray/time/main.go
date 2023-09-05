package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("输入要执行的程序")
	var fileName string
	flag.StringVar(&fileName, "filename", "datatype", "输入执行的文件")
	flag.Parse()

	switch fileName {
	case "d":
		datatype()
	case "duradd":
		add()
	case "dursub":
		sub()
	case "durtickDemo":
		tickDemo()
	default:
		fmt.Println("hello")
	}
}
