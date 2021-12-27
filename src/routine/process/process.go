package process

import (
	"fmt"
	"log"
	"os"
	"time"
)

type process struct {
	agentDir   string
	binaryFile string
}

func (p *process) stratfins() {
	//
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}, //其他变量如果不清楚可以不设定
	}
	pro, err := os.StartProcess("./fins/fins", []string{"./fins/fins", "-c /home/steven/code/code_sync/go/src/github.com/agent/bin/worker/6/worker6_input.conf -d /home/steven/code/code_sync/go/src/github.com/agent/bin/worker/6/worker6_output.conf"}, attr) //vim 打开tmp.txt文件
	if err != nil {
		fmt.Println(err)
	}
	//go func() {
	//	p.Signal(os.Kill) //kill process
	//}()

	pstat, err := pro.Wait()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pstat)

	return
}

func (p *process) StartFins() error {

	//// start Metric server
	//var addr = ":"
	//addr += strconv.Itoa(6555)
	//listener,err := net.Listen("tcp", addr)
	////listener,err := netutil.Listen("tcp", addr)
	//if err != nil {
	//	return err
	//}
	//fmt.Println("listener port: ", listener.Addr().String())


	go p.stratfins()
	time.Sleep(10 * time.Second)


	tcp := MTcpNew()
	tcp.Receiving()

	log.Println("StartFins() in process.go")
	return nil
}

func ProcessNew(dir string, f string) (*process, error) {
	return &process{
		agentDir:   dir,
		binaryFile: f,
	}, nil
}