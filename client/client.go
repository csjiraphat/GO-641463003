package main

import (
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close() // close connection before exit

	// buffer for reading
	buffer := make([]byte, 1024)
	for {
		// Read data from the client
		n, err := conn.Read(buffer) // Read() blocks until it reads some data from the network and n is the number of bytes read
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		// Print the number of bytes read
		fmt.Printf("Received %d bytes\n", n)

		// Convert received data to string
		receivedData := string(buffer[:n])

		// Split received data into username and password
		credentials := strings.Fields(receivedData)
		if len(credentials) != 2 {
			fmt.Println("Invalid input format. Please provide username and password.")
			continue
		}

		username := credentials[0]
		password := credentials[1]

		// Validate username and password (you can replace this with your own validation logic)
		if username == "your_username" && password == "your_password" {
			// Send a response back to the client
			response := "Hello\n"
			conn.Write([]byte(response))
		} else {
			// Send an error response back to the client
			response := "Invalid username or password\n"
			conn.Write([]byte(response))
		}
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

		fmt.Println("New connection established")

		// handle the connection in a new goroutine
		go handleConnection(conn)
	}
}
