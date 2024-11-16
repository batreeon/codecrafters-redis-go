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
		fmt.Println("3 ", o)
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
			if len(cmds) > 1 {
				cmds = cmds[1:]
			} else if len(cmds) == 1 {
				cmds = []string{}
			}
			return cmds, []byte(fmt.Sprintf(simpleStrings, "PONG")), nil
		}
	case "ECHO":
		if len(cmds) < 2 {
			return cmds, nil, errors.New("Error ECHO parameter missing")
		}

		result := []byte(fmt.Sprintf(bulkStrings, len(cmds[1]), cmds[1]))
		if len(cmds) > 2 {
			cmds = cmds[2:]
		} else if len(cmds) == 2 {
			cmds = []string{}
		}
		return cmds, result, nil
	}
	return cmds, nil, errors.ErrUnsupported
}
