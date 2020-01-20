package main


import (
	"github.com/go_sample/src/tsfile/common/header"
	"github.com/go_sample/src/tsfile/common/log"
)

func main() {
	log.Info("test!!!")
	log.Info("float=%d", header.FLOAT)
	log.Info("SNAPPY=%d", header.SNAPPY)
	log.Info("BITMAP=%d", header.BITMAP)
	log.Info("mqtt config = %s", header.MQTT_CONF)
}
