package main

import (
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Buffer for reading
	buffer := make([]byte, 1024)

	// Read data from the client
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	// Convert received data to string and split into username and password
	data := strings.TrimSpace(string(buffer[:n]))
	credentials := strings.Split(data, "\n")
	if len(credentials) != 2 {
		fmt.Println("Invalid credentials format")
		conn.Write([]byte("Invalid credentials\n"))
		return
	}

	username := strings.TrimSpace(strings.Split(credentials[0], ":")[1])
	password := strings.TrimSpace(strings.Split(credentials[1], ":")[1])

	// Check username and password
	if username == "std1" && password == "p@ssw0rd" {
		// Send a response back to the client
		response := "Hello\n"
		conn.Write([]byte(response))
		fmt.Println("Authentication successful for user:", username)
	} else {
		// Send a response back to the client
		response := "Invalid credentials\n"
		conn.Write([]byte(response))
		fmt.Println("Invalid credentials for user:", username)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 5000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("New connection")

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}
