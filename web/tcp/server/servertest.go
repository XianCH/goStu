package main

import (
	"fmt"
	"net"
)

func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected")

	buffer := make([]byte, 1024)
	for {
		// 从客户端读取数据
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		// 将收到的数据回传给客户端
		_, err = conn.Write(buffer[:n])
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}
}

// func main() {
// 	// 创建监听器
// 	listener, err := net.Listen("tcp", "localhost:8080")
// 	if err != nil {
// 		fmt.Println("Error listening:", err)
// 		os.Exit(1)
// 	}
// 	defer listener.Close()
// 	fmt.Println("Server listening on localhost:8080")

// 	for {
// 		// 等待客户端连接
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("Error accepting: ", err)
// 			continue
// 		}
// 		// 启动一个协程处理客户端连接
// 		go handleClient(conn)
// 	}
// }
