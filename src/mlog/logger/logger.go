package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

/**
 * @Project: worker
 * @Package Name: main
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/7/2 下午12:39
 * @Description: add new logger for all worker
 */

/* define new struct loggerWorker*/
type loggerWorker interface {
	Info(format string, v... interface{})
	Error(format string, v... interface{})
	Warning(format string, v... interface{})
}

/* define new struct logger*/
type logger struct {
	log   loggerWorker
	info  *log.Logger
	debug *log.Logger
	err   *log.Logger
}

/**
 * @Description: interface for info print
 * @Params: format, what user want to print
 * @return: nothing
 * @Date: 2020/7/2 下午1:40
 */
func (l *logger) Info(format string, v... interface{}) {
	if l.log != nil {
		l.log.Info(format, v...)
	} else {
		l.info.Println(fmt.Sprintf(format, v...))
	}
}

/**
 * @Description: interface for error print
 * @Params: format, what user want to print
 * @return: nothing
 * @Date: 2020/7/2 下午1:40
 */
func (l *logger) Error(format string, v... interface{}) {
	if l.log != nil {
		l.log.Error(format, v...)
	} else {
		l.err.Println(fmt.Sprintf(format, v...))
	}
}

/**
 * @Description: interface for debug print
 * @Params: format, what user want to print
 * @return: nothing
 * @Date: 2020/7/2 下午1:40
 */
func (l *logger) Debug(format string, v... interface{}) {
	if l.log != nil {
		l.log.Warning(format, v...)
	} else {
		l.debug.Println(fmt.Sprintf(format, v...))
	}
}

/**
 * @Description: interface for user to add method himself
 * @Params: log, print interface
 * @return: nothing
 * @Date: 2020/7/2 下午1:40
 */
func (l *logger) SetLogger(log loggerWorker) {
	l.log = log
}

/**
 * @Description: init default logger
 * @Params: no
 * @return: no
 * @Date: 2020/7/2 下午1:14
 */
func (l *logger) init(){
	l.info = log.New(os.Stdout, time.Now().Format(time.RFC3339Nano) + " ", 0)
	l.debug = log.New(os.Stdout, time.Now().Format(time.RFC3339Nano) + " ", 0)
	l.err = log.New(os.Stderr, time.Now().Format(time.RFC3339Nano) + " ", 0)
	return
}

/**
 * @Description:
 * @Params:
 * @return:
 * @Date: 2020/7/2 下午1:12
 */
func LoggerNew() (*logger, error) {
	logger := &logger{}
	logger.init()
	return logger, nil
}
