package rdb

import (
	"os"
	"testing"
	"time"

	"github.com/batreeon/codecrafters-redis-go/test"
)

func TestConfig(t *testing.T) {
	tests := []test.Test{
		{
			Name:    "0",
			CmdArgs: []string{"program", "--dir", "/tmp/redis-files", "--dbfilename", "dump.rdb"},
			Args:    []string{"*3\r\n$6\r\nCONFIG\r\n$3\r\nGET\r\n$3\r\ndir\r\n"},
			Want:    []string{"*2\r\n$3\r\ndir\r\n$16\r\n/tmp/redis-files\r\n"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			go test.TestServer(t)

			// 使用 mockArgs 来替换 os.Args
			oldArgs := os.Args
			os.Args = tt.CmdArgs
			defer func() { os.Args = oldArgs }()

			// waiting server start up
			time.Sleep(1 * time.Second)

			test.TestClient(t, tt)
		})
	}
}
