package introduction

import (
	"testing"
	"time"

	"github.com/batreeon/codecrafters-redis-go/test"
)

func Test_PING(t *testing.T) {
	tests := []test.Test{
		{
			Name:    "0",
			CmdArgs: []string{""},
			Args:    []string{"*1\r\n$4\r\nPING\r\n"},
			Want:    []string{"+PONG\r\n"},
		},
	}

	go test.TestServer(t)
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			// waiting server start up
			time.Sleep(1 * time.Second)

			test.TestClient(t, tt)
		})
	}
}
