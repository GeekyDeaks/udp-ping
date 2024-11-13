package main

// Run with ./server $PORT

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <port>\n", os.Args[0])
		os.Exit(1)
	}

	port := os.Args[1]
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println("Error listening on UDP port:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("UDP echo server listening on port %s\n", port)

	// Buffer to store received data
	buffer := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		timestamp := time.Now().Format(time.RFC3339)

		fmt.Printf("%s | Received: %s from %s\n", timestamp, string(buffer[:n]), clientAddr.String())
		_, err = conn.WriteToUDP(buffer[:n], clientAddr)
		if err != nil {
			fmt.Println("Error writing to UDP:", err)
			continue
		}
	}
}
