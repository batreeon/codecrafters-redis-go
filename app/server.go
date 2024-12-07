package main

import (
	"fmt"
	"os"

	"github.com/batreeon/codecrafters-redis-go/app/config"
	"github.com/batreeon/codecrafters-redis-go/app/rdb"
	"github.com/batreeon/codecrafters-redis-go/app/server"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	config.SetConfigs()
	rdb.Load()
	err := server.StartServer()
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}
}
