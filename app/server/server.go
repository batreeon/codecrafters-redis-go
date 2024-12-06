package server

import (
	"fmt"
	"net"
	"os"

	"github.com/batreeon/codecrafters-redis-go/app/execute"
	"github.com/batreeon/codecrafters-redis-go/app/parse"
)

func StartServer() error {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		return fmt.Errorf("failed to bind to port 6379. %v", err)
	}
	defer l.Close()

	for { // handle multiple client connections
		conn, err := l.Accept()
		if err != nil {
			return fmt.Errorf("error accepting connection: %v", err.Error())
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for { // handle multiple commands in the same connection
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error read request: ", err.Error())
			return // return when one connection handler completed
		}
		fmt.Printf("input %q\n", buf[:n])

		cmds, err := parse.ParserInput(buf[:n])
		if err != nil {
			fmt.Println("Error parse input: ", err.Error())
			os.Exit(1)
		}
		fmt.Printf("parsed input: %q\n", cmds)

		err = execute.ExecuteCmd(conn, cmds)
		if err != nil {
			fmt.Println("Error execute cmds: ", err.Error())
			os.Exit(1)
		}
	}
}
