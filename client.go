package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

type DSyncClock struct {
	ServerAddress string
	ServerTime    uint64
	Diff          int64
}

func (d *DSyncClock) doSync() {
	// Connect to the server
	conn, err := net.Dial("tcp", d.ServerAddress)
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to TCP Time Server")

	// Read the response from the server
	buffer := make([]byte, 20)
	n, err := conn.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading from server: %v\n", err)
		return
	}

	// Print the received time
	response := string(buffer[:n])
	now := time.Now().UnixNano()

	newServerTime, err := strconv.ParseUint(response, 10, 64)
	if err != nil {
		fmt.Printf("Error parsing server time: %v\n", err)
		return
	}

	fmt.Printf("Diff: %d fixed to %d\n", d.Diff, int64(newServerTime)-now)
	d.ServerTime = newServerTime
	d.Diff = int64(newServerTime) - now
}

func (d *DSyncClock) LocalTimeToDSyncClockTime() uint64 {
	return uint64(time.Now().UnixNano() + d.Diff)
}
