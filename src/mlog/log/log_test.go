package log

import (
	"testing"
)

func Test_Main(t *testing.T) {
	Println("test Println.")
	Info("test Info.")
	Debug("test Debug.")
	Error("test Error.")
}

