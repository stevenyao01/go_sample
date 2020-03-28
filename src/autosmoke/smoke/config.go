package smoke

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const separator = "=: "

type config struct {
	broker   string
	confPath string
}

func (c *config) loadConfigFile(name string) (map[string]string, error) {
	dir, _ := os.Getwd()
	fmt.Println("dir: ", dir)
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var conf = make(map[string]string)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" ||
			strings.HasPrefix(line, "#") ||
			strings.HasPrefix(line, "//") ||
			!strings.ContainsAny(line, separator) {
			continue
		}

		for _, s := range []byte(separator) {
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

func (c *config) modifyConfig() error {
	// modify local mqtt.conf broker
	conf, errLoadConfigFile := c.loadConfigFile(c.confPath + "/mqtt.conf")
	if errLoadConfigFile != nil {
		fmt.Println("errLoadConfigFile: ", errLoadConfigFile.Error())
	}
	//fmt.Println(conf)
	defaultConfig := map[string]string{
		"broker":             "172.17.203.36:4567",
		"heartbeat-interval": "30",
		"heartbeat-token":    "localid",
		"metrics-port":       "0",
		"ping-timeout":       "1",
		"device-type":        "1",
		"ip-address":         "127.0.0.1",
		"metrics-interval":   "60",
		"keepalive":          "2",
	}

	var buf bytes.Buffer
	for k, v := range defaultConfig{
		if k == "broker" {
			buf.WriteString(k + "=" + c.broker + "\n")
			continue
		}
		if conf[k] != "" {
			buf.WriteString(k + "=" + conf[k] + "\n")
		}else {
			buf.WriteString(k + "=" + v + "\n")
		}
	}
	errWrite := ioutil.WriteFile(c.confPath+"/mqtt.conf", buf.Bytes(), 0644)
	if errWrite != nil {
		fmt.Println("errWrite: ", errWrite)
		return errWrite
	}
	return nil
}

func ConfigNew(b string, cp string) (*config, error) {
	return &config{
		broker:   b,
		confPath: cp,
	}, nil
}
