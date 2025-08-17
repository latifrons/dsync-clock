package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Listen on TCP port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP Time Server listening on :8080")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		// Handle each connection in a goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Get current time in nanoseconds
	currentTime := time.Now().UnixNano()
	timeStr := fmt.Sprintf("%d", currentTime)

	// Send the time to the client
	_, err := conn.Write([]byte(timeStr))
	if err != nil {
		fmt.Printf("Error writing to client: %v\n", err)
	}

	fmt.Printf("Sent time %d to client %s\n", currentTime, conn.RemoteAddr())
}
