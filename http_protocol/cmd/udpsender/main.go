package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	// Setup udp 'server'
	udpServer, err := net.ResolveUDPAddr("udp", "localhost:42069")

	if err != nil {
		fmt.Errorf("Failed to listen to port 42069: %w", err)
	}

	// get udp connection
	conn, err := net.DialUDP("udp", nil, udpServer)

	if err != nil {
		fmt.Errorf("Dial failed: %w", err)
	}

	// Always remember to close your doors :)
	defer conn.Close()

	// New reader
	reader := bufio.NewReader(os.Stdin)

	for {
		// Print > to say 'Connection is listening/ready' type message
		fmt.Printf(">")
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Errorf("Failed to read line: %w", err)
		}

		if _, err := conn.Write(line); err != nil {
			fmt.Errorf("Failed to write line to connection: %w", err)
		}
	}
}
