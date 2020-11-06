package log

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	cLog CLog
	Debugf *log.Logger
	Infof  *log.Logger
	Errorf *log.Logger
)

/* define new struct Extension*/
type Extension struct {
	Time     string `json:"timestamp"`
	WorkerId string `json:"workerid,omitempty"`
	Uuid     string `json:"uuid,omitempty"`
}

/* define new struct CLog*/
type CLog struct {
	Platform   string    `json:"platform,omitempty"`
	Product    string    `json:"product,omitempty"`
	LogType    string    `json:"logType,omitempty"`
	ModuleName string    `json:"module"`
	LogLevel   string    `json:"loglevel"`
	Ip         string    `json:"ip,omitempty"`
	Content    string    `json:"content"`
	Extension  Extension `json:"extension,omitempty"`
}

//func Println(format string, v... interface{}){
//	Infof.Println(fmt.Sprintf(format, v...))
//}
//
//func Info(format string, v... interface{}){
//	Infof.Println(fmt.Sprintf(format, v...))
//}
//
//func Debug(format string, v... interface{}){
//	Debugf.Println(fmt.Sprintf(format, v...))
//}
//
//func Error(format string, v... interface{}){
//	Errorf.Println(fmt.Sprintf(format, v...))
//}

func Println(format string, v... interface{}){
	cLog.LogLevel = "info"
	cLog.Extension.Time = time.Now().Format(time.RFC3339Nano)
	cLog.Content = fmt.Sprintf(format, v...)
	data, err := json.Marshal(cLog)
	if err != nil {
		log.Println("ErrorLog json.Marshal error: ", err.Error())
	}
	Infof.Println(string(data))
}

func Info(format string, v... interface{}){
	cLog.LogLevel = "info"
	cLog.Extension.Time = time.Now().Format(time.RFC3339Nano)
	cLog.Content = fmt.Sprintf(format, v...)
	data, err := json.Marshal(cLog)
	if err != nil {
		log.Println("ErrorLog json.Marshal error: ", err.Error())
	}
	Infof.Println(string(data))
}

func Debug(format string, v... interface{}){
	cLog.LogLevel = "debug"
	cLog.Extension.Time = time.Now().Format(time.RFC3339Nano)
	cLog.Content = fmt.Sprintf(format, v...)
	data, err := json.Marshal(cLog)
	if err != nil {
		log.Println("ErrorLog json.Marshal error: ", err.Error())
	}
	Debugf.Println(string(data))
}

func Warning(format string, v... interface{}){
	cLog.LogLevel = "warning"
	cLog.Extension.Time = time.Now().Format(time.RFC3339Nano)
	cLog.Content = fmt.Sprintf(format, v...)
	data, err := json.Marshal(cLog)
	if err != nil {
		log.Println("ErrorLog json.Marshal error: ", err.Error())
	}
	Errorf.Println(string(data))
}

func Error(format string, v... interface{}){
	cLog.LogLevel = "error"
	cLog.Extension.Time = time.Now().Format(time.RFC3339Nano)
	cLog.Content = fmt.Sprintf(format, v...)
	data, err := json.Marshal(cLog)
	if err != nil {
		log.Println("ErrorLog json.Marshal error: ", err.Error())
	}
	Errorf.Println(string(data))
}

func Fatal(format string, v... interface{}){
	cLog.LogLevel = "fatal"
	cLog.Extension.Time = time.Now().Format(time.RFC3339Nano)
	cLog.Content = fmt.Sprintf(format, v...)
	data, err := json.Marshal(cLog)
	if err != nil {
		log.Println("ErrorLog json.Marshal error: ", err.Error())
	}
	Errorf.Println(string(data))
}

//func WarningWithExtension(extension string, format string, v... interface{}){
//	cLog.LogLevel = "warning"
//	cLog.Extension = extension
//	cLog.Content = fmt.Sprintf(format, v...)
//	data, err := json.Marshal(cLog)
//	if err != nil {
//		log.Println("ErrorLog json.Marshal error: ", err.Error())
//	}
//	Errorf.Println(string(data))
//	cLog.Extension = ""
//}
//
//func ErrorWithExtension(extension string, format string, v... interface{}){
//	cLog.LogLevel = "error"
//	cLog.Extension = extension
//	cLog.Content = fmt.Sprintf(format, v...)
//	data, err := json.Marshal(cLog)
//	if err != nil {
//		log.Println("ErrorLog json.Marshal error: ", err.Error())
//	}
//	Errorf.Println(string(data))
//	cLog.Extension = ""
//}
//
//func FatalWithExtension(extension string, format string, v... interface{}){
//	cLog.LogLevel = "fatal"
//	cLog.Extension = extension
//	cLog.Content = fmt.Sprintf(format, v...)
//	data, err := json.Marshal(cLog)
//	if err != nil {
//		log.Println("ErrorLog json.Marshal error: ", err.Error())
//	}
//	Errorf.Println(string(data))
//	cLog.Extension = ""
//}


func New(module string) {
	cLog.ModuleName = module

	Infof = log.New(os.Stdout, "", 0)
	Debugf = log.New(os.Stdout, "", 0)
	Errorf = log.New(os.Stderr, "", 0)

	//Infof = log.New(os.Stdout, time.Now().Format(time.RFC3339Nano) + " [" + module + "] " + " [" + level + "] ", 0)
	//Debugf = log.New(os.Stdout, time.Now().Format(time.RFC3339Nano) + " [" + module + "] " + " [" + level + "] ", 0)
	//Errorf = log.New(os.Stderr, time.Now().Format(time.RFC3339Nano) + " [" + module + "] " + " [" + level + "] ", 0)
}

//func init() {
//	Infof = log.New(os.Stdout, time.Now().Format(time.RFC3339Nano) + " [WORK] ", 0)
//	Debugf = log.New(os.Stdout, time.Now().Format(time.RFC3339Nano) + " [WORK] ", 0)
//	Errorf = log.New(os.Stderr, time.Now().Format(time.RFC3339Nano) + " [WORK] ", 0)
//}