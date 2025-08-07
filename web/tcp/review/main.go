package main

import (
	"time"

	tcpClient "github.com/x14n/goStu/web/tcp/review/client"

	tcpserver "github.com/x14n/goStu/web/tcp/review/server"
)

func main() {
	go tcpserver.TcpServer()
	time.Sleep(1 * time.Second)
	tcpClient.TcpClient()
}
