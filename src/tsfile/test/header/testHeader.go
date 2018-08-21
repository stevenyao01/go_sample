package main


import (
	"github.com/go_sample/src/tsfile/common/header"
	"github.com/go_sample/src/tsfile/common/log"
)

func main() {
	log.Info("test!!!")
	log.Info("float=", header.FLOAT)
	log.Info("SNAPPY=", header.SNAPPY)
	log.Info("BITMAP=", header.BITMAP)
	log.Info("mqtt config = ", header.MQTT_CONF)
}
