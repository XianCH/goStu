package main

import (
	"fmt"
	"net"

	"github.com/x14n/goStu/web/tcp/stickpacky"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	defer conn.Close()

	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	// messageClientoA(conn)
	messageClientoB(conn)
}

func padMessage(msg string) []byte {
	if len(msg) > stickpacky.MsgLength {
		return []byte(msg[:stickpacky.MsgLength])
	}
	buf := make([]byte, 0, stickpacky.MsgLength)

	buf = fmt.Appendf(buf, "%-*s", stickpacky.MsgLength, msg)
	return buf
	// return []byte(fmt.Sprintf("%-*s", stickpacky.MsgLength, msg)) // 补空格

}

func messageClientoA(conn net.Conn) {
	messages := []string{"Hello", "World", "Fixed length message", "111111111111111111111111111111222"}

	for _, msg := range messages {
		conn.Write(padMessage(msg))
	}
}

func messageClientoB(conn net.Conn) {
	messages := []string{"Hello", "World", "Fixed length message", "111111111111111111111111111111222"}
	for _, msg := range messages {
		conn.Write([]byte(msg + "\n"))
	}
}
