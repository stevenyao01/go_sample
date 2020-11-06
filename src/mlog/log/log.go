package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Debuga *log.Logger
	Infoa  *log.Logger
	Errora *log.Logger
)

func Println(format string, v... interface{}){
	Infoa.Println(fmt.Sprintf(format, v...))
}

func Info(format string, v... interface{}){
	Infoa.Println(fmt.Sprintf(format, v...))
}

func Debug(format string, v... interface{}){
	Debuga.Println(fmt.Sprintf(format, v...))
}

func Error(format string, v... interface{}){
	Errora.Println(fmt.Sprintf(format, v...))
}

func init() {
	Infoa = log.New(os.Stdout, time.Now().Format(time.RFC3339Nano) + " ", 0)
	Debuga = log.New(os.Stdout, time.Now().Format(time.RFC3339Nano) + " ", 0)
	Errora = log.New(os.Stderr, time.Now().Format(time.RFC3339Nano) + " ", 0)
}