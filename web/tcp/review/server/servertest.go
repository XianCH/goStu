package server

import (
	"bufio"
	"fmt"
	"net"
)

func TcpServer() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			return
		}
		go tcpRequestHelper(conn)
	}
}

func tcpRequestHelper(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	fmt.Printf("Echoing message: %s", message)

	conn.Write([]byte("Message received\n"))
	conn.Close()
}
