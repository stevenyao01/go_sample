package main

import (
	"bufio"
	"github.com/go_sample/src/log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

/**
 * @Project: go_sample
 * @Package Name: main
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 2020/7/7 下午1:54
 * @Description:
 */

var mutex sync.Mutex

func readStderr(read, write *os.File, id string) {
	defer read.Close()
	defer write.Close()
	rd := bufio.NewReader(read)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			log.Error("stderr read error: %s", err.Error())
			return
		}
		log.Info("line: %s", line)
	}
}

func startProcess(binary string, id string) (*os.Process, error) {
	mutex.Lock()
	defer mutex.Unlock()

	local := "input.conf"
	inout := "output.conf"
	other := "other.conf"

	read, write, err := os.Pipe()
	if err != nil {
		return nil, err
	}

	//readOut, writeOut, errOut := os.Pipe()
	//if errOut != nil {
	//	return nil, errOut
	//}

	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, write},
		//Files: []*os.File{os.Stdin, writeOut, write},
	}
	//attr := &os.ProcAttr{
	//	Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	//}

	log.Info("start: %s -c %s -d %s -o %s", binary, local, inout, other)

	p, err := os.StartProcess(binary, []string{binary, "-c", local, "-d", inout, "-o", other, "-debug"}, attr)
	if err != nil {
		if err := read.Close(); err != nil {
			log.Error("close pipe read error:", err.Error())
		}
		if err := write.Close(); err != nil {
			log.Error("close pipe write error:", err.Error())
		}
		return nil, err
	}

	go readStderr(read, write, id)
	//go a.readStdout(readOut, writeOut, id)

	return p, nil
	//return os.StartProcess(binary, []string{binary, "-c", local, "-d", inout, "-o", other}, attr)
}

func GetBaseDir() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	dir = strings.Replace(dir, "\\", "/", -1)

	return dir, nil
}

func main() () {
	var pro *os.Process

	base,err := GetBaseDir()
	if err != nil {
		log.Error("GetBaseDir error:%s", err.Error())
		return
	}
	log.Info("GetBaseDir: %s", base)

	binary := base + "/bin/worker/110/modbus"

	for i := 0; i < 3; i++ {
		pro, err = startProcess(binary, "110")
		if err != nil {
			log.Error("start worker error:%s", err.Error())
			continue
		}

		go myWait(pro)
	}
	log.Info("before sleep")
	time.Sleep(time.Second * 100)
	log.Info("after sleep")
}

func myWait(pro *os.Process) {
	if pro != nil {
		ps, err := pro.Wait()
		if err != nil {
			log.Error("wait worker error:%s", err.Error())
		} else {
			log.Error("wait worker ps-string:%s", ps.String())
		}
	}
}
