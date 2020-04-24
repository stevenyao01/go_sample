package smoke

import (
	"fmt"
	"github.com/go_sample/src/autosmoke/utils"
	"log"
	"os"
	"strings"
)

type agent struct {
	agentDir   string
	runtime    string
	binaryFile os.FileInfo
}

func (a *agent) startAgent() error {
	// add exec right
	currenDir, errGetPwd := os.Getwd()
	if errGetPwd != nil {
		fmt.Println("errGetwd: ", errGetPwd.Error())
		return errGetPwd
	}

	errChdir := os.Chdir(a.agentDir)
	if errChdir != nil {
		fmt.Println("errChdir: ", errChdir.Error())
		return errChdir
	}
	errChmod := os.Chmod(a.binaryFile.Name(), 0755)
	if errChmod != nil {
		fmt.Println("errChmod: ", errChmod.Error())
		return errChmod
	}

	// start
	read, write, err := os.Pipe()
	if err != nil {
		fmt.Println("err: ", err.Error())
		return err
	}
	attr := &os.ProcAttr{
		Files: []*os.File{os.Stdin, write, write},
	}
	//p, err := os.StartProcess(binary, []string{binary, "-c", local, "-d", inout, "-o", other}, attr)
	binary := a.binaryFile.Name()
	var pro *os.Process
	var errStart error
	if strings.Contains(a.agentDir, "windows") {
		pro, errStart = os.StartProcess("/usr/bin/wine", []string{"wine", binary}, attr)
	}else {
		pro, errStart = os.StartProcess(binary, []string{binary}, attr)
	}
	if errStart != nil {
		if err := read.Close(); err != nil {
			log.Println("close pipe read error:", err.Error())
		}
		if err := write.Close(); err != nil {
			log.Println("close pipe write error:", err.Error())
		}
		return err
	}
	errRetChdir := os.Chdir(currenDir)
	if errRetChdir != nil {
		fmt.Println("errChdir: ", errRetChdir.Error())
		return errRetChdir
	}

	log.Println("agentDir: ", a.agentDir)
	go utils.ReadStderr(a.agentDir, read, write)
	go utils.StopProcess(pro, a.runtime)
	ps, errWait := pro.Wait()
	if errWait != nil {
		log.Println("wait worker error: ", errWait.Error())
		return errWait
	}
	log.Println("wait到信号: ", ps.String())
	return nil
}

func AgentNew(dir string, f os.FileInfo, rt string) (*agent, error) {
	return &agent{
		agentDir:   dir,
		runtime:    rt,
		binaryFile: f,
	}, nil
}
