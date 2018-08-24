package log

import (
	"fmt"
	"log"
	"strings"
	"os"
)

var level = 0

func getTag() string {
	s := strings.Split(os.Args[0], "/")//[2:3]
	l := len(s)
	v := s[l-1:l][0]
	vv := v+"                    "
	return fmt.Sprintf("[ %.15s.. ] ",vv)
}

var tag = getTag()

func SetLevel(l int) {
	level = l;	
}

func Debuga(format string) {
	log.Printf("%8.2s : %s","debug",format)
}

func Debug(s string, v ...interface{}) {
	if (level > 1)  {
		return	
	}
	fmt.Print("\033[32m")
	fmt.Print(tag)
	log.Print(" DEBUG:",fmt.Sprintf(s,v...))
	fmt.Print("\033[0m")
}
func Info(s string,v ...interface{}) {
	if (level > 3)  {
		return	
	}
	fmt.Print("\033[34m")
	//fmt.Print(tag)
	//log.SetFlags(log.Lshortfile)
	log.Print("  INFO: ",fmt.Sprintf(s,v...))
	fmt.Print("\033[0m")
}

func Error(s string,v ...interface{}) {
	if (level > 5)  {
		return	
	}
	fmt.Print("\033[31m")
	fmt.Print(tag)
	log.Print(" ERROR:",fmt.Sprintf(s,v...))
	fmt.Print("\033[0m")
}
