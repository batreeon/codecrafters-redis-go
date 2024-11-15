package main

import (
	"fmt"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	// l, err := net.Listen("tcp", "0.0.0.0:6379")
	// if err != nil {
	// 	fmt.Println("Failed to bind to port 6379")
	// 	os.Exit(1)
	// }
	// _, err = l.Accept()
	// if err != nil {
	// 	fmt.Println("Error accepting connection: ", err.Error())
	// 	os.Exit(1)
	// }
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	
	go func() {
		for {
			buf := make([]byte, 128)
			_, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Error read request: ", err.Error())
				os.Exit(1)
			}
			_, err = conn.Write([]byte("+PONG\r\n"))
			if err != nil {
				fmt.Println("Error write response: ", err.Error())
				os.Exit(1)
			}
		}
	}()
}
