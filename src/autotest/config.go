package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const separator = "=: "

func loadConfigFile(name string) (map[string]string, error) {
	dir, _ := os.Getwd()
	fmt.Println("dir: ", dir)
	data,err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var conf = make(map[string]string)
	lines := strings.Split(string(data), "\n")
	for _,line := range lines {
		line = strings.TrimSpace(line)
		if line == "" ||
			strings.HasPrefix(line, "#") ||
			strings.HasPrefix(line, "//") ||
			!strings.ContainsAny(line, separator) {
			continue
		}

		for _,s := range []byte(separator) {
			kv := strings.Split(line, string(s))
			if len(kv) > 1 && kv[0] != "" && kv[1] != "" {
				key := strings.Trim(kv[0], ` "`)
				val := strings.Trim(kv[1], ` "`)
				conf[key] = val
				break
			}
		}
	}

	return conf, nil
}