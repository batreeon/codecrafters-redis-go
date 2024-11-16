package main

import (
	"errors"
	"fmt"
	"net"
)

func executeCmd(conn net.Conn, cmds []string) error {
	if len(cmds) == 0 {
		return nil
	}
	for len(cmds) > 0 {
		var o []byte
		var err error
		cmds, o, err = output(cmds)
		if err != nil {
			return err
		}
		_, err = conn.Write(o)
		if err != nil {
			return err
		}
	}

	return nil
}

func output(cmds []string) ([]string, []byte, error) {
	switch cmds[0] {
	case "PING":
		{
			return cmds, []byte(fmt.Sprintf(simpleStrings, "PONG")), nil
		}
	case "ECHO":
		if len(cmds) < 2 {
			return cmds, nil, errors.New("Error ECHO parameter missing")
		}
		return cmds, []byte(fmt.Sprintf(bulkStrings, len(cmds), cmds[1])), nil
	}
	return cmds, nil, errors.ErrUnsupported
}
