package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("./xxxx.txt")
	if err != nil {
		fmt.Println("open file err :", err)
		return
	}
	defer file.Close()
	// 定义接收文件读取的字节数组
	var buf [128]byte
	var content []byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))

	f, err2 := os.Create("./copy.go")
	defer f.Close()
	if err2 != nil {
		fmt.Println("create file erro", err2)
	}
	_, err3 := f.WriteString(string(content))
	if err3 != nil {
		fmt.Println("file write file", err3)
		return
	}

}
