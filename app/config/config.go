package config

import (
	"flag"
)

var configs = map[string]string{}

func SetConfigs() {
	dir := flag.String("dir", "", "the path to the directory where the RDB file is stored (example: /tmp/redis-data)")
	dbfilename := flag.String("dbfilename", "", "the name of the RDB file (example: rdbfile)")
	flag.Parse()

	configs["dir"] = *dir
	configs["dbfilename"] = *dbfilename
}

func GetConfig(k string) string {
	return configs[k]
}
