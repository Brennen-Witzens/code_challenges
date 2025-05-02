package main

import (
	"fmt"
	"http_protocol/internal/request"
	"net"
)

const chunkSize = 8

func main() {

	ln, err := net.Listen("tcp", ":42069")

	if err != nil {
		fmt.Errorf("Failed to listen to port 42069: %w", err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Errorf("failed to accept a connection: %v", err)
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())

		req, err := request.RequestFromReader(conn)

		if err != nil {
			fmt.Errorf("failed to read from request: %v", err)
		}

		fmt.Printf("Request Line: \n - Method: %s \n - Target: %s \n -Version: %s\n", req.RequestLine.Method, req.RequestLine.RequestTarget, req.RequestLine.HttpVersion)

		fmt.Println("Connect to ", conn.RemoteAddr(), "closed")
	}

}
