package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			log.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		log.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Println("net start file")
		os.Exit(1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("server read Error")
			continue
		}
		go process(conn)
	}
}
