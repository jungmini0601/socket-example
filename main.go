package main

import (
	"fmt"
	"strings"

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
	// 커넥션 받을 때 까지 대기
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	// HTTP Parse
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Error acceptiong connection: ", err)
	}

	req := string(buf)
	path := parsePath(req)
	res := makeResponseFromPath(path)
	conn.Write([]byte(res))
	fmt.Println("Server Stop")
}

func parsePath(req string) string {
	lines := strings.Split(req, "\r\n")
	path := strings.Split(lines[0], " ")[1]
	fmt.Println(path)
	return path
}

func makeResponseFromPath(path string) string {
	var res string
	if path == "/" {
		res = "HTTP/1.1 200 OK\r\n\r\n"
	} else if strings.Split(path, "/")[1] == "echo" {
		message := strings.Split(path, "/")[2]
		res = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(message), message)
	} else {
		res = "HTTP/1.1 404 Not Found\r\n\r\n"
	}
	return res
}
