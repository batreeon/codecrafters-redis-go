package storage

import (
	"fmt"
	"time"

	"github.com/codecrafters-io/redis-starter-go/constant"
)

var db = map[string]string{}
var expireAt = map[string]time.Time{}

func set(k, v string) {
	db[k] = v
}

func SetWithExpire(k, v string, expire time.Duration) {
	set(k, v)
	if expire > 0 {
		expireAt[k] = time.Now().Add(expire)
		fmt.Println("set k: ", k, " expire: ", expire, " now: ", time.Now(), " expireAt: ", expireAt[k])
	}
}

func Get(k string) string {
	if isExpired(k) {
		delete(db, k)
		delete(expireAt, k)
		return constant.NullBulkStrings
	}
	v, ok := db[k]
	if !ok {
		return constant.NullBulkStrings
	}
	return v
}

func isExpired(k string) bool {
	ea, ok := expireAt[k]
	fmt.Println("get k: ", k, " expireAt: ", ea, " now: ", time.Now())
	if !ok {
		return false
	}
	return time.Now().After(ea)
}
