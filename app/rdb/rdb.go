package rdb

import (
	"github.com/codecrafters-io/redis-starter-go/app/storage"
	"github.com/codecrafters-io/redis-starter-go/constant"
)

const (
	opcodeMetadataFA byte = 250 // 0xFA, the start of a meatdata subsection

	opcodeDatabaseFE     byte = 254 // 0xFE, the start a database subsection
	opcodeDatabaseSizeFB byte = 251 // 0xFB, hash table size information follows
	opcodeMsExpireKeyFC  byte = 252 // 0xFC
	opcodeSExpireFD      byte = 253 // 0xFD

	dataTypeString00 byte = 00
)

func pairsDecode(rdbFile []byte) ([]byte, error) {
	// dataType := rdbFile[0]
	rdbFile = rdbFile[1:]

	keyLen := rdbFile[0]
	rdbFile = rdbFile[1:]
	key := rdbFile[:keyLen]
	rdbFile = rdbFile[keyLen:]

	valueLen := rdbFile[0]
	rdbFile = rdbFile[1:]
	value := rdbFile[:valueLen]
	rdbFile = rdbFile[valueLen:]

	storage.SetWithExpire(string(key), string(value), constant.NoExpire)

	return rdbFile, nil
}

func sizeDecode([]byte) {

}
