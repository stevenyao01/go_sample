package smoke

import (
	"fmt"
	"github.com/go_sample/src/autosmoke/utils"
	"log"
	"os"
	"strings"
)

type worker struct {
	workerDir  string
	runtime    string
	binaryFile os.FileInfo
}

func (w *worker) startWorker() error {
	// add exec right
	currenDir, errGetPwd := os.Getwd()
	if errGetPwd != nil {
		fmt.Println("errGetwd: ", errGetPwd.Error())
		return errGetPwd
	}

	bina := strings.Split(w.workerDir, "/")
	errCopyConfig := utils.CopyDir("config/"+bina[1], w.workerDir)
	if errCopyConfig != nil {
		fmt.Println("copy ", bina[1], " config file err: ", errCopyConfig.Error())
	}

	errChdir := os.Chdir(w.workerDir)
	if errChdir != nil {
		fmt.Println("errChdir: ", errChdir.Error())
		return errChdir
	}
	errChmod := os.Chmod(w.binaryFile.Name(), 0755)
	if errChmod != nil {
		fmt.Println("errChmod: ", errChmod.Error())
		return errChmod
	}

	if w.binaryFile.Name() == "filebeat" {
		errChmod := os.Chmod("fb", 0755)
		if errChmod != nil {
			fmt.Println("errChmod: ", errChmod.Error())
			return errChmod
		}
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
	binary := w.binaryFile.Name()
	var pro *os.Process
	var errStart error
	if strings.Contains(w.workerDir, "windows") {
		//pro, errStart = os.StartProcess("/usr/bin/wine", []string{"wine", binary}, attr)
		pro, errStart = os.StartProcess("/usr/bin/wine", []string{"wine", "./" + binary, "-c", "input.conf", "-d", "output.conf", "-o", "other.conf", "-debug"}, attr)
	} else {
		//pro, errStart = os.StartProcess(binary, []string{binary}, attr)
		pro, errStart = os.StartProcess(binary, []string{binary, "-c", "input.conf", "-d", "output.conf", "-o", "other.conf", "-debug"}, attr)
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

	go utils.ReadStderr(w.workerDir, read, write)
	if w.binaryFile.Name() == "filebeat" {
		f, errOpen := os.OpenFile(w.workerDir+"/a.log", os.O_WRONLY|os.O_APPEND, 0666)
		if errOpen != nil {
			fmt.Println("filebeat open a.log err: ", errOpen.Error())
		} else {
			_, errWrite := f.Write([]byte("hello world!!!\n"))
			if errWrite != nil {
				fmt.Println("filebeat write string err: ", errWrite)
			}
			_, errWrite = f.Write([]byte("hello world&&&\n"))
			if errWrite != nil {
				fmt.Println("filebeat write string err: ", errWrite)
			}
		}
	}
	go utils.StopProcess(pro, w.runtime)
	ps, errWait := pro.Wait()
	if errWait != nil {
		log.Println("wait worker error: ", errWait.Error())
		return errWait
	}
	log.Println("wait到信号: ", ps.String())
	return nil
}

func WorkerNew(dir string, f os.FileInfo, rt string) (*worker, error) {
	return &worker{
		workerDir:  dir,
		runtime:    rt,
		binaryFile: f,
	}, nil
}
