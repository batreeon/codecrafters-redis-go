package rdb

import (
	"testing"
	"time"

	"github.com/batreeon/codecrafters-redis-go/test"
)

func TestConfig(t *testing.T) {
	tests := []test.Test{
		{
			Name:    "0",
			CmdArgs: []string{"program", "--dir", "/tmp/redis-files", "--dbfilename", "dump.rdb"},
			Args:    []string{"*3\r\n$6\r\nCONFIG\r\n$3\r\nGET\r\n$3\r\ndir\r\n", "*3\r\n$6\r\nCONFIG\r\n$3\r\nGET\r\n$10\r\ndbfilename\r\n"},
			Want:    []string{"*2\r\n$3\r\ndir\r\n$16\r\n/tmp/redis-files\r\n", "*2\r\n$10\r\ndbfilename\r\n$8\r\ndump.rdb\r\n"},
		},
	}

	go test.TestServer(t)
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			test.SetConfig(tt)

			// waiting server start up
			time.Sleep(1 * time.Second)

			test.TestClient(t, tt)
		})
	}
}
