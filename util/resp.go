package util

import (
	"fmt"

	"github.com/batreeon/codecrafters-redis-go/constant"
)

func BuildSimpleStrings(s string) []byte {
	return []byte(fmt.Sprintf(constant.SimpleStrings, s))
}

func BuildBulkStrings(s string) []byte {
	return []byte(fmt.Sprintf(constant.BulkStrings, len(s), s))
}

func BuildArrays(s []string) []byte {
	head := fmt.Sprintf(constant.Arrays, len(s))
	var elements []byte
	for _, ss := range s {
		elements = append(elements, BuildBulkStrings(ss)...)
	}
	return append([]byte(head), elements...)
}
