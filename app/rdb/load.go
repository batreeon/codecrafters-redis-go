package rdb

import (
	"bytes"
	"errors"
	"os"

	"github.com/batreeon/codecrafters-redis-go/app/config"
)

func Load() error {
	rdbFilePath := config.GetConfig("dir") + "/" + config.GetConfig("dbfilename")
	rdbFile, err := os.ReadFile(rdbFilePath)
	if err != nil {
		return err
	}

	fb := bytes.IndexByte(rdbFile, opcodeDatabaseSizeFB)
	if fb == -1 {
		return errors.New("rdb file format incorrect")
	}

	_, err = pairsDecode(rdbFile[fb+3:])
	if err != nil {
		return err
	}

	return nil
}
