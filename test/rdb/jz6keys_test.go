package rdb

import (
	"testing"
	"time"

	"github.com/batreeon/codecrafters-redis-go/app/rdb"
	"github.com/batreeon/codecrafters-redis-go/test"
	"github.com/batreeon/codecrafters-redis-go/util"
)

func TestKeys(t *testing.T) {
	tests := []test.Test{
		{
			Name:    "0",
			CmdArgs: []string{"program", "--dir", "../../data", "--dbfilename", "dump.rdb"},
			Args:    []string{string(util.BuildArrays([]string{"KEYS","*"}))},
			Want:    []string{"*1\r\n$5\r\nhello\r\n"},
		},
	}

	go test.TestServer(t)
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			test.SetConfig(tt)
			rdb.Load()

			// waiting server start up
			time.Sleep(1 * time.Second)

			test.TestClient(t, tt)
		})
	}
}
