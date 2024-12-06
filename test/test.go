package test

import (
	"flag"
	"net"
	"reflect"
	"testing"

	"github.com/batreeon/codecrafters-redis-go/app/config"
	"github.com/batreeon/codecrafters-redis-go/app/server"
)

// 在测试前设置模拟的命令行参数
func init() {
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
}

type Test struct {
	Name    string
	CmdArgs []string
	Args    []string
	Want    []string
}

func TestServer(t *testing.T) {
	config.SetConfigs()
	err := server.StartServer()
	if err != nil {
		t.Errorf("failed to start server, err: %v", err)
	}
}

func TestClient(t *testing.T, tt Test) {
	conn, err := net.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		t.Fatalf("failed to connect to server, err: %v", err)
	}
	for i, arg := range tt.Args {
		arg := arg
		want := tt.Want[i]
		_, err := conn.Write([]byte(arg))
		if err != nil {
			t.Errorf("failed to write to server, err: %v", err)
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			t.Errorf("failed to read from server, err: %v", err)
		}
		if !reflect.DeepEqual(string(buf[:n]), want) {
			t.Errorf("got = %v, want = %v", string(buf[:n]), tt.Want[i])
		}
	}

}
