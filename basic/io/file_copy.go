package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//打开源文件
	srcFile, err := os.Open("./xxxx.txt")
	if err != nil {
		fmt.Println("open srcFIle error", err)
		panic(err)
	}

	dstFile, err := os.Create("./copy.txt")
	if err != nil {
		fmt.Println("create dstFIle error", err)
		panic(err)
	}

	buf := make([]byte, 1024)

	for {
		n, err2 := srcFile.Read(buf)
		if err2 == io.EOF {
			fmt.Println("读取结束")
			break
		}
		if err2 != nil {
			fmt.Println("read srcFile error", err2)
			return
		}
		_, err3 := dstFile.Write(buf[:n])
		if err3 != nil {
			fmt.Println("copy file error", err3)
		}
	}

	srcFile.Close()
	dstFile.Close()
}
