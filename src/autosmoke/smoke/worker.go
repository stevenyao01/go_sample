package smoke

import (
	"fmt"
	"github.com/go_sample/src/autosmoke/utils"
	"log"
	"os"
)

type worker struct {
	agentDir   string
	binaryFile os.FileInfo
}

func (a *agent) startWorker() error {
	// add exec right
	errChmod := os.Chmod(a.agentDir + a.binaryFile.Name(), 0755)
	if errChmod != nil {
		fmt.Println("errChmod: ", errChmod.Error())
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
	binary := a.agentDir + a.binaryFile.Name()
	pro, err := os.StartProcess(binary, []string{binary}, attr)
	if err != nil {
		if err := read.Close(); err != nil {
			log.Println("close pipe read error:", err.Error())
		}
		if err := write.Close(); err != nil {
			log.Println("close pipe write error:", err.Error())
		}
		return err
	}
	go utils.ReadStderr(a.agentDir, read, write)
	go utils.StopProcess(pro, a.runtime)
	ps, errWait := pro.Wait()
	if errWait != nil {
		log.Println("wait worker error: ", errWait.Error())
		return errWait
	}
	log.Println("ps: ", ps.String())
	return nil
}

func WorkerNew(dir string, f os.FileInfo) (*agent, error) {
	return &agent{
		agentDir:   dir,
		binaryFile: f,
	}, nil
}
