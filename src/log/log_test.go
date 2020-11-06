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

func Test_New(t *testing.T) {
	New("opcda")
	Println("test Println.")
	Info("test Info....")
	Debug("test Debug...")
	Warning("test Warning.")
	//WarningWithExtension("{\"extension\":\"yaohp1\"}", "test WarningWithExtension.")
	Error("test Error...")
	//ErrorWithExtension("{\"extension\":\"yaohp2\"}", "test ErrorWithExtension.")
	Fatal("test Fatal...")
	//FatalWithExtension("{\"extension\":\"yaohp3\"}", "test FatalWithExtension.")
}

