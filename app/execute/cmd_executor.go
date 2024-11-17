package execute

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/app/storage"
	"github.com/codecrafters-io/redis-starter-go/constant"
	"github.com/codecrafters-io/redis-starter-go/util"
)

type cmdExecutor func(cmds []string) ([]string, []byte, error)

var cmdExecutorMap = map[string]cmdExecutor{
	"PING": PINGExecutor,
	"ECHO": ECHOExecutor,
	"SET":  SETEXecutor,
	"GET":  GETEXecutor,
}

func PINGExecutor(cmds []string) ([]string, []byte, error) {
	cmds = util.RemoveFirstNElements(cmds, 1)
	return cmds, []byte(fmt.Sprintf(constant.SimpleStrings, "PONG")), nil
}

func ECHOExecutor(cmds []string) ([]string, []byte, error) {
	if len(cmds) < 2 {
		return cmds, nil, constant.ErrParameterMissing
	}

	resp := buildBulkStrings(cmds[1])
	cmds = util.RemoveFirstNElements(cmds, 2)
	return cmds, resp, nil
}

func SETEXecutor(cmds []string) ([]string, []byte, error) {
	if len(cmds) < 3 {
		return cmds, nil, constant.ErrParameterMissing
	}

	k := cmds[1]
	v := cmds[2]
	storage.M[k] = v

	resp := buildBulkStrings("OK")
	cmds = util.RemoveFirstNElements(cmds, 3)
	return cmds, resp, nil
}

func GETEXecutor(cmds []string) ([]string, []byte, error) {
	if len(cmds) < 2 {
		return cmds, nil, constant.ErrParameterMissing
	}

	k := cmds[1]
	v, ok := storage.M[k]
	var resp []byte
	if !ok {
		resp = []byte(constant.NullBulkStrings)
	} else {
		resp = buildBulkStrings(v)
	}

	cmds = util.RemoveFirstNElements(cmds, 2)
	return cmds, resp, nil
}


func buildBulkStrings(s string) []byte {
	return []byte(fmt.Sprintf(constant.BulkStrings, len(s), s))
}
