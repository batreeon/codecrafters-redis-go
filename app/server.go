package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
	}
	defer l.Close()

	for { // handle multiple client connections
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go func() {
			defer conn.Close()
			for { // handle multiple commands in the same connection
				buf := make([]byte, 128)
				_, err := conn.Read(buf)
				if err != nil {
					fmt.Println("Error read request: ", err.Error())
					return // return when one connection handler completed
				}

				fmt.Println("1",string(buf))
				cmds, err := ParserInput(buf)
				if err != nil {
					fmt.Println("Error parse input: ", err.Error())
					os.Exit(1)
				}

				fmt.Println("2", cmds)
				err = executeCmd(conn, cmds)
				if err != nil {
					fmt.Println("Error execute cmds: ", err.Error())
					os.Exit(1)
				}
			}
		}()
	}
}
