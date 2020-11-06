package logger

import (
	"testing"
)

func Test_Main(t *testing.T) {
	logger, _ := LoggerNew()

	logger.Info("info")
	logger.Debug("debug")
	logger.Error("error")
}

