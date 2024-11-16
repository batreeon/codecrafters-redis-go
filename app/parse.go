package main

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/constant"
)

const (
	simpleStrings    = "+%s\r\n"
	bulkStrings      = "$%d\r\n%s\r\n"
	emptyBulkStrings = "$%d\r\n\r\n"
	arrays           = "*%d\r\n"
)

func ParserInput(bs []byte) ([]string, error) {
	result := []string{}
	for len(bs) > 0 {
		var eles []string
		var err error
		bs, eles, err = parseRESPDataType(bs)
		if err != nil {
			return nil, err
		}

		result = append(result, eles...)
	}

	return result, nil
}

func parseRESPDataType(bs []byte) ([]byte, []string, error) {
	if len(bs) == 0 {
		return bs, nil, constant.ErrInvaildInput
	}
	var result []string
	switch bs[0] {
	case '+':
		{
			var ele string
			_, err := fmt.Sscanf(string(bs), simpleStrings, &ele)
			if err != nil {
				return bs, nil, err
			}

			pl := parsedLength(simpleStrings, ele)
			if len(bs) > pl {
				bs = bs[pl:]
			} else {
				bs = []byte{}
			}

			result = append(result, ele)
			return bs, result, nil
		}
	case '$':
		{
			// parse the lenth, first
			var elelen int
			_, err := fmt.Sscanf(string(bs), bulkStrings[:3], &elelen)
			if err != nil {
				return bs, nil,err
			}

			var ele string
			var pattern string
			if elelen == 0 {
				_, err = fmt.Sscanf(string(bs), emptyBulkStrings, &elelen)
				pattern = emptyBulkStrings
			} else if elelen > 0 {
				_, err = fmt.Sscanf(string(bs), bulkStrings, &elelen, &ele)
				pattern = bulkStrings
			} else {
				return bs, nil, constant.ErrInvaildInput
			}

			if err != nil {
				return bs, nil, err
			}
			if elelen != len(ele) {
				return bs, nil, constant.ErrInvaildInput
			}

			pl := parsedLength(pattern, elelen, ele)
			if len(bs) > pl {
				bs = bs[pl:]
			} else {
				bs = []byte{}
			}

			result = append(result, ele)
			return bs, result, nil
		}
	case '*':
		{
			var n int
			_, err := fmt.Sscanf(string(bs), arrays, &n)
			if err != nil {
				return bs, nil, err
			}
			pl := parsedLength(arrays, n)
			if len(bs) > pl {
				bs = bs[pl:]
			} else {
				bs = []byte{}
			}

			for i := 0; i < n; i++ {
				var ele []string
				var err error
				bs, ele, err = parseRESPDataType(bs)
				if err != nil {
					return bs, nil, err
				}
				result = append(result, ele...)
			}
			return bs, result, nil
		}
	}
	return bs, nil, constant.ErrInvaildInput
}

func parsedLength(pattern string, v ...any) int {
	return len(fmt.Sprintf(pattern, v...))
}
