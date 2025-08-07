package stickpacky

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func HanderMsgLength(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, MsgLength)
	for {
		n, err := io.ReadFull(conn, buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by client")
			} else {
				fmt.Printf("Error reading from connection: %v\n", err)
			}
			return
		}

		msg := string(buf[:n])
		fmt.Printf("Received message: %s\n", msg)
	}
}

func HanderSymbol(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading from connection: %v\n", err)
			return
		}
		fmt.Printf("Received message: %s", msg)
	}
}
