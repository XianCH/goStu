package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func TcpClient() {
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Send data to server
	fmt.Println("Enter message to send:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	conn.Write([]byte(input))

	// Receive data from server
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	fmt.Println("Message from server:", response)
}
