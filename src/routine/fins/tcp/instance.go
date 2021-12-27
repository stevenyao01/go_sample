package tcp

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type myInstance struct {
	lock sync.Mutex
	ch   chan int
}
func (m *myInstance) getBaseDir() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	dir = strings.Replace(dir, "\\", "/", -1)

	return dir, nil
}

func (m *myInstance) isPathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (m *myInstance) mkdir(path string) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	if !strings.Contains(path, ":") && !strings.HasPrefix(path, "/") {
		path = strings.Replace(path, "\\", "/", -1)
		dirs := strings.Split(path, "/")
		path = ""
		for _, dir := range dirs {
			path += dir + "/"
			if !m.isPathExists(path) {
				if err := os.Mkdir(path, 0755); err != nil {
					return err
				}
			}
		}

		return nil
	}

	root := "/"
	if strings.Contains(path, ":") {
		root = strings.Split(path, ":")[0] + ":/"
		path = strings.TrimLeft(strings.Split(path, ":")[1], "/")
	}
	if strings.HasPrefix(path, "/") {
		root = "/"
		path = strings.TrimLeft(path, "/")
	}

	dirs := strings.Split(path, "/")
	path = root
	for _, dir := range dirs {
		path += dir + "/"
		if !m.isPathExists(path) {
			if err := os.Mkdir(path, 0755); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *myInstance) StartMyInstance(command string, version string, proId int) {
	base,err := m.getBaseDir()
	if err != nil {
		fmt.Println("get base dir err: ", err)
	}

	//dir := base + "/bin/opcua_" + version + "/" + proId
	dir := fmt.Sprintf("%s/bin/opcua_%s/%d", base, version, proId)
	err = m.mkdir(dir)
	if err != nil {
		fmt.Println("mkdir err: ", err)
	}

	stopFlag := false
	log.Println("yaohiaping")
	for {
		select {
		case <- m.ch:
			stopFlag = true
		default:
			log.Println("loop in StartMyInstance() pId: ", proId)
			time.Sleep(5 * time.Second)
		}
		if stopFlag {
			break
		}
	}
	return
}

func (m *myInstance) StopMyInstance(command string, version string, proId int) {
	m.ch <- 1
	return
}

func MyInstanceNew() *myInstance {
	return &myInstance{
		ch: 	make(chan int),
	}
}
