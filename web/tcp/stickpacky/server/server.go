package main

import (
	"fmt"
	"net"

	"github.com/x14n/goStu/web/tcp/stickpacky"
)

func main() {
	listenr, _ := net.Listen("tcp", ":12345")
	fmt.Println("server start at 12345")

	for {
		conn, err := listenr.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		// go stickpacky.HanderMsgLength(conn)
		go stickpacky.HanderSymbol(conn)
	}
}
