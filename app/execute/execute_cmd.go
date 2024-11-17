package execute

import (
	"errors"
	"fmt"
	"net"

	"github.com/codecrafters-io/redis-starter-go/constant"
)

func ExecuteCmd(conn net.Conn, cmds []string) error {
	if len(cmds) == 0 {
		return nil
	}
	for len(cmds) > 0 {
		var resp []byte
		var err error
		cmds, resp, err = output(cmds)
		if err != nil {
			return err
		}
		fmt.Println("3 ", resp)
		_, err = conn.Write(resp)
		if err != nil {
			return err
		}
	}

	return nil
}

func output(cmds []string) ([]string, []byte, error) {
	executor, ok := cmdExecutorMap[cmds[0]]
	if !ok {
		return cmds, nil, errors.ErrUnsupported
	}
	return executor(cmds)
}
