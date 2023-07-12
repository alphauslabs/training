package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer listen.Close()
	for {
		log.Println("waiting for connections...")
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		log.Println("processing...")
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// Incoming request.
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	// Write data to response.
	time := time.Now().Format(time.RFC3339)
	responseStr := fmt.Sprintf("Your message is: %v. Received time: %v", string(buffer[:]), time)
	conn.Write([]byte(responseStr))

	// Close conn.
	conn.Close()
}
