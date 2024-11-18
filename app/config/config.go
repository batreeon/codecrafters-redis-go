package config

import (
	"flag"
	"fmt"
	"os"
)

var Configs = map[string]string{}

func SetConfigs() {
	fmt.Println("cmd augs: ", os.Args)
	dir := flag.String("dir", "", "the path to the directory where the RDB file is stored (example: /tmp/redis-data)")
	dbfilename := flag.String("dbfilename", "", "the name of the RDB file (example: rdbfile)")
	fmt.Printf("configs: dir=%s, dbfilename=%s", *dir, *dbfilename)
	Configs["dir"] = *dir
	Configs["dbfilename"] = *dbfilename
}

func GetConfig(k string) string {
	return Configs[k]
}
