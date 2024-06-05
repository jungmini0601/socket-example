package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// Uncomment this block to pass the first stage
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	fmt.Println("Server Started to Port: 4221")
	fmt.Println("Wait for Connection...")
	conn, err := l.Accept()
	fmt.Println("Connected!")
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	fmt.Println("Server Stop")
}
