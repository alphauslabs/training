package main

import (
	"net"
	"os"
)

func main() {
	tcpServer, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpServer)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte("hello, server"))
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}

	// Buffer to get data.
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}

	println("received:", string(received))
	conn.Close()
}
