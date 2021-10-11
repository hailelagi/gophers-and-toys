package rps

import (
	"bufio"
	"fmt"
	"net"
)

func Server() {
	fmt.Println("Starting server..")

	ln, _ := net.Listen("tcp", ":8000")

	// accept connection
	conn, _ := ln.Accept()

	for {
		// get message, output
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))
	}
}
