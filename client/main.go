package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"time"
)

type PingPacket struct {
	Timestamp time.Time `json:"timestamp"`
	Id        int       `json:"id"`
}

func ping(conn *net.UDPConn) {

	defer conn.Close()

	for i := 0; ; i++ {

		message, err := json.Marshal(PingPacket{
			Timestamp: time.Now(),
			Id:        i,
		})

		if err != nil {
			fmt.Printf("Error creating packet: %v\n", err)
			return
		}

		_, err = conn.Write(message)
		if err != nil {
			fmt.Printf("Error sending packet: %v\n", err)
			return
		}

		time.Sleep(1 * time.Second)

	}
}

func main() {
	server := flag.String("server", "firestone.racedepartment.com", "Server address")
	port := flag.String("port", "9600", "Server port")
	flag.Parse()

	serverAddr := fmt.Sprintf("%s:%s", *server, *port)

	addr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Printf("Error resolving address: %v\n", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Printf("Error dialing UDP: %v\n", err)
		return
	}

	fmt.Printf("Starting UDP ping to %s\n", serverAddr)

	go ping(conn)

	for {

		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		timestamp := time.Now().Format(time.RFC3339)
		if err != nil {
			fmt.Printf("%s | Error receiving response %v\n", timestamp, err)
			return
		}

		pong := PingPacket{}

		err = json.Unmarshal(buffer[:n], &pong)
		if err != nil {
			fmt.Printf("%s | Error decoding response %s, %v\n", timestamp, string(buffer[:n]), err)
			return
		}

		duration := time.Since(pong.Timestamp)
		fmt.Printf("%s | Received: %d | RTT: %v ms\n", timestamp, pong.Id, duration.Milliseconds())

	}
}
