package smoke

import (
	"errors"
	"fmt"
	"github.com/go_sample/src/autosmoke/utils"
	"github.com/jenkins-x/golang-jenkins"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type smoke struct {
	urlStr         string
	jobStr         string
	localPathStr   string
	userStr        string
	pwdStr         string
	brokerStr      string
	buildNumberStr string
	skStr          string
	jk             *jenkins
	ag             *agent
	config         *config
	ml             *mail
}

func (s *smoke) Start() error {
	errPrepare := s.prepareJenkins()
	if errPrepare != nil {
		fmt.Println("errPrepare: ", errPrepare.Error())
		return errPrepare
	}

	build := s.jk.getCurrentBuild()

	errProcess := s.process(build)
	if errProcess != nil {
		fmt.Println("errProcess: ", errProcess)
		return errProcess
	}

	return nil
}

func (s *smoke) process(build gojenkins.Build) error {
	for _, v := range build.Artifacts {
		errGetArtifact := s.getArtifact(v)
		if errGetArtifact != nil {
			fmt.Println("errGetArtifact: ", errGetArtifact.Error())
			return errGetArtifact
		}

		// unzip file
		fileArr := strings.Split(v.FileName, "-")
		if len(fileArr) < 2 {
			continue
		}

		dirArr := strings.Split(fileArr[0], "_")
		if len(dirArr) != 3 {
			continue
		}
		unZipPath := dirArr[0] + dirArr[1] + dirArr[2]
		unZipDir := s.localPathStr + "/" + unZipPath
		errUnZip := utils.UnzipFile(s.localPathStr+"/"+v.FileName, unZipDir)
		if errUnZip != nil {
			fmt.Println("unzip file err: ", errUnZip)
			return errUnZip
		}
		// modfiy broker in mqtt config
		config, errConfigNew := ConfigNew(s.brokerStr, unZipDir)
		if errConfigNew != nil {
			fmt.Println("errConfigNew: ", errConfigNew.Error())
			return errConfigNew
		}
		errModify := config.modifyConfig()
		if errModify != nil {
			fmt.Println("errModify: ", errModify.Error())
		}

		// get broker sk file
		errSk := utils.GetSk(s.brokerStr, unZipDir)
		if errSk != nil {
			fmt.Println("errSk: ", errSk.Error())
			if s.skStr != "" {
				_, errCopyFile := utils.CopyFile(s.skStr, unZipDir+"/device.sk")
				if errCopyFile != nil {
					fmt.Println("errCopyFile: ", errCopyFile.Error())
				}
			}
			return errors.New("download device.sk failed, please special a device.sk")
		}

		// get all file in download zip
		files, errReadDir := ioutil.ReadDir(unZipDir) //读取目录下文件
		if errReadDir != nil {
			return errReadDir
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			} else {
				if strings.Contains(file.Name(), "EdgeAgent_") {
					errProcessAgent := s.processAgent(file, unZipDir)
					if errProcessAgent != nil {
						return errProcessAgent
					}
				} else {
					errProcessWorker := s.processWorker(file, unZipDir)
					if errProcessWorker != nil {
						return errProcessWorker
					}
				}
			}
		}
	}
	return nil
}

func (s *smoke) getArtifact(v gojenkins.Artifact) error {
	zip, errArtifact := s.jk.getArtifact(v)
	if errArtifact != nil {
		fmt.Println("errArtifact: ", errArtifact.Error())
		return errArtifact
	}
	errSaveToFile := utils.SaveToFile(zip, s.localPathStr+"/"+v.FileName)
	if errSaveToFile != nil {
		fmt.Println("errSaveToFile: ", errSaveToFile)
		return errSaveToFile
	}
	return nil
}

func (s *smoke) processAgent(file os.FileInfo, unZipDir string) error {
	fmt.Println("file: ", file.Name())
	// get agent, add exec right, start it.
	files, errReadFile := ioutil.ReadDir(unZipDir)
	//读取目录下文件
	if errReadFile != nil {
		return errReadFile
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			if strings.Contains(file.Name(), "EdgeAgent_") {
				fmt.Println("file: ", file.Name())
				var errAgentNew error
				s.ag, errAgentNew = AgentNew(unZipDir, file)
				if errAgentNew != nil {
					fmt.Println("errAgentNew: ", errAgentNew.Error())
					return errAgentNew
				}

				errStartAgent := s.ag.startAgent()
				if errStartAgent != nil {
					log.Println("errStartAgent: ", errStartAgent.Error())
					return errStartAgent
				}
			}
		}
	}
	return nil
}

func (s *smoke) processWorker(file os.FileInfo, unZipDir string) error {
	fmt.Println("file: ", file.Name())
	// get agent, add exec right, start it.
	files, errReadFile := ioutil.ReadDir(unZipDir)
	//读取目录下文件
	if errReadFile != nil {
		return errReadFile
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			if strings.Contains(file.Name(), "EdgeAgent_") {
				fmt.Println("file: ", file.Name())
				var errWorkerNew error
				s.ag, errWorkerNew = WorkerNew(unZipDir, file)
				if errWorkerNew != nil {
					fmt.Println("errWorkerNew: ", errWorkerNew.Error())
					return errWorkerNew
				}

				errStartWorker := s.ag.startWorker()
				if errStartWorker != nil {
					log.Println("errStartAgent: ", errStartWorker.Error())
					return errStartWorker
				}
			}
		}
	}
	return nil
}

func (s *smoke) sendMailToMonitor(msg string) error {
	fromUser := "leapiot@126.com"
	toUser := []string{"yaohp1@lenovo.com", "leapiot@126.com"}
	subject := "LeapIot Smoke Test"
	var errMail error
	s.ml, errMail = MailNew(fromUser, toUser, subject)
	if errMail != nil {
		fmt.Println("errMail: ", errMail.Error())
		return errMail
	}
	err := s.ml.SendMail(msg)
	if err != nil {
		log.Println("发送邮件失败, err: ", err.Error())
		return err
	}
	log.Println("发送邮件成功")
	return nil
}

func (s *smoke) prepareJenkins() error {
	s.jk, _ = JenkinsNew(s.urlStr, s.jobStr, s.localPathStr, s.userStr, s.pwdStr, s.buildNumberStr)
	lastBuild, errLastBuild := s.jk.getLastBuild()
	if errLastBuild != nil {
		fmt.Println("errLastBuild: ", errLastBuild.Error())
		return errLastBuild
	}
	build, errLastSuccessfulBuild := s.jk.getLastBuild()
	if errLastSuccessfulBuild != nil {
		fmt.Println("errLastSuccessfulBuild: ", errLastSuccessfulBuild.Error())
		return errLastSuccessfulBuild
	}
	fmt.Println("lastBuild / lastSuccessBuild : ", lastBuild.Id, "/", build.Id)
	fmt.Println("Download last success build: ", build.Url)
	if s.buildNumberStr != "0" {
		var errSpecial error
		build, errSpecial = s.jk.getSpecialBuild()
		if errSpecial != nil {
			fmt.Println("errSpecial: ", errSpecial.Error())
			return errSpecial
		}
		fmt.Println("use special version: ", build.Id)
	}

	s.jk.setCurrentBuild(build)

	output, errConsoleOutput := s.jk.getConsoleOutput(build)
	if errConsoleOutput != nil {
		fmt.Println("errConsoleOutput: ", errConsoleOutput.Error())
		return errConsoleOutput
	}
	//fmt.Println("output: ", string(output))
	if !strings.Contains(string(output), "Finished:") {
		errMsg := fmt.Sprintf("当前job正在运行,build.Number=%d, or build failed", build.Number)
		return errors.New(errMsg)
	}

	errMkdir := os.Mkdir(s.localPathStr, os.ModePerm)
	if errMkdir != nil {
		if !strings.Contains(errMkdir.Error(), "file exists") {
			fmt.Println("errMkdir: ", errMkdir.Error())
		}
	}
	return nil
}

func New(ur string, j string, l string, us string, p string, br string, bn string, sk string) (*smoke, error) {
	return &smoke{
		urlStr:         ur,
		jobStr:         j,
		localPathStr:   l,
		userStr:        us,
		pwdStr:         p,
		brokerStr:      br,
		buildNumberStr: bn,
		skStr:          sk,
	}, nil
}
