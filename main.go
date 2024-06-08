package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// HTTP Parse
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error acceptiong connection: ", err)
	}

	req := string(buf)
	path := parsePath(req)
	res := makeResponseFromPath(path, req)
	conn.Write([]byte(res))
}

func parseUserAgent(req string) (string, error) {
	lines := strings.Split(req, "\r\n")
	for _, value := range lines {
		fmt.Println(value)
		header := strings.Split(value, " ")
		key := header[0]
		val := header[1]
		if key == "User-Agent:" {
			return val, nil
		}
	}

	return "error", errors.New("User-Agent Header Not Found")
}

func parsePath(req string) string {
	lines := strings.Split(req, "\r\n")
	path := strings.Split(lines[0], " ")[1]
	return path
}

func makeResponseFromPath(path string, req string) string {
	var res string
	if path == "/" {
		res = "HTTP/1.1 200 OK\r\n\r\n"
	} else if strings.Split(path, "/")[1] == "echo" {
		message := strings.Split(path, "/")[2]
		res = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(message), message)
	} else if path == "/user-agent" {
		agent, err := parseUserAgent(req)
		if err != nil {
			res = "HTTP/1.1 400 Bad Request\r\n\r\n"
		} else {
			res = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(agent), agent)
		}
	} else {
		res = "HTTP/1.1 404 Not Found\r\n\r\n"
	}
	return res
}
