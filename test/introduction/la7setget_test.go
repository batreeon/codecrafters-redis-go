package introduction

import (
	"testing"
	"time"

	"github.com/batreeon/codecrafters-redis-go/test"
	"github.com/batreeon/codecrafters-redis-go/util"
)

func Test_SET_GET(t *testing.T) {
	tests := []test.Test{
		{
			Name:    "0",
			CmdArgs: []string{"program"},
			Args:    []string{string(util.BuildArrays([]string{"SET", "foo", "bar"}))},
			Want:    []string{"+OK\r\n"},
		},
		{
			Name:    "1",
			CmdArgs: []string{"program"},
			Args:    []string{string(util.BuildArrays([]string{"Get", "foo"}))},
			Want:    []string{"$3\r\nbar\r\n"},
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
