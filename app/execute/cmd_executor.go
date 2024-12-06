package execute

import (
	"fmt"
	"strconv"
	"time"

	"github.com/batreeon/codecrafters-redis-go/app/config"
	"github.com/batreeon/codecrafters-redis-go/app/storage"
	"github.com/batreeon/codecrafters-redis-go/constant"
	"github.com/batreeon/codecrafters-redis-go/util"
)

type cmdExecutor func(cmds []string) ([]string, []byte, error)

var cmdExecutorMap = map[string]cmdExecutor{
	"ping":   pingExecutor,
	"echo":   echoExecutor,
	"set":    setExecutor,
	"get":    getExecutor,
	"config": configExecutor,
	"keys":   keysExecutor,
}

func buildBulkStrings(s string) []byte {
	return []byte(fmt.Sprintf(constant.BulkStrings, len(s), s))
}

func buildArrays(s []string) []byte {
	head := fmt.Sprintf(constant.Arrays, len(s))
	var elements []byte
	for _, ss := range s {
		elements = append(elements, buildBulkStrings(ss)...)
	}
	return append([]byte(head), elements...)
}

func pingExecutor(cmds []string) ([]string, []byte, error) {
	cmds = util.RemoveFirstNElements(cmds, 1)
	return cmds, []byte(fmt.Sprintf(constant.SimpleStrings, "PONG")), nil
}

func echoExecutor(cmds []string) ([]string, []byte, error) {
	if len(cmds) < 2 {
		return cmds, nil, constant.ErrParameterMissing
	}

	resp := buildBulkStrings(cmds[1])
	cmds = util.RemoveFirstNElements(cmds, 2)
	return cmds, resp, nil
}

func setExecutor(cmds []string) ([]string, []byte, error) {
	if len(cmds) < 3 {
		return cmds, nil, constant.ErrParameterMissing
	}

	k := cmds[1]
	v := cmds[2]
	var expireTime time.Duration
	if len(cmds) >= 5 && cmds[3] == "px" {
		expire, err := strconv.Atoi(cmds[4])
		if err != nil {
			return cmds, nil, err
		}
		expireTime = time.Duration(expire) * time.Millisecond
		cmds = util.RemoveFirstNElements(cmds, 5)
	} else {
		cmds = util.RemoveFirstNElements(cmds, 3)
	}
	storage.SetWithExpire(k, v, expireTime)

	resp := buildBulkStrings("OK")
	return cmds, resp, nil
}

func getExecutor(cmds []string) ([]string, []byte, error) {
	if len(cmds) < 2 {
		return cmds, nil, constant.ErrParameterMissing
	}

	k := cmds[1]
	v := storage.Get(k)
	var resp []byte
	if v != constant.NullBulkStrings {
		resp = buildBulkStrings(v)
	} else {
		resp = []byte(v)
	}
	cmds = util.RemoveFirstNElements(cmds, 2)
	return cmds, resp, nil
}

func configExecutor(cmds []string) ([]string, []byte, error) {
	if len(cmds) < 3 {
		return cmds, nil, constant.ErrParameterMissing
	}
	if cmds[1] == "get" {
		configKey := cmds[2]
		configValue := config.GetConfig(configKey)
		resp := buildArrays([]string{configKey, configValue})
		cmds = util.RemoveFirstNElements(cmds, 3)
		return cmds, resp, nil
	}
	return cmds, nil, constant.ErrInvaildInput
}

func keysExecutor(cmds []string) ([]string, []byte, error) {
	if len(cmds) < 2 {
		return cmds, nil, constant.ErrParameterMissing
	}

	if cmds[1] == "*" {
		keys := storage.Keys()
		resp := buildArrays(keys)
		cmds = util.RemoveFirstNElements(cmds, 2)
		return cmds, resp, nil
	}

	return cmds, nil, constant.ErrInvaildInput
}
