package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 连接到服务端
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 发送数据到服务端
	message := "Hello, TCP Server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending:", err)
		return
	}

	// 从服务端接收响应
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error receiving:", err)
		return
	}

	fmt.Println("Received from server:", string(buffer[:n]))
}
