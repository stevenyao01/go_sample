package smoke

import (
	"bytes"
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
	runtime        string
	skStr          string
	jk             *jenkins
	ag             *agent
	config         *config
	ml             *mail
	result         map[string]string
	wk 			   *worker
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

	attachFile := make([]string,100,100)
	var content bytes.Buffer
	i := 0
	for k, v := range s.result{
		content.WriteString(k + " : " + v + "\n")
		target := strings.Split(k, "_")
		atStr := target[0] + "/" + target[1] + target[2]
		attachFile[i] = s.localPathStr + "/" + atStr + "/" + target[0] + target[1] + target[2] + ".log"
		i++
	}
	fmt.Println("content: ", content.String())
	errSendMailToMonitor := s.sendMailToMonitor(content.String(), attachFile)
	if errSendMailToMonitor != nil {
		fmt.Println("errSendMailToMonitor: ", errSendMailToMonitor.Error())
		return errSendMailToMonitor
	}
	if strings.Compare(s.jobStr, "LeapEdge.agentSign") == 0 {
		fmt.Println("所有平台测试完毕，请查收邮件结果。")
	} else {
		fmt.Println("(linux, amd64)平台测试完毕，请查收邮件结果。")
	}


	return nil
}

func (s *smoke) process(build gojenkins.Build) error {
	totalCount := 0
	passCount := 0
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
		unZipPath := dirArr[0] + "/" + dirArr[1] + dirArr[2]
		unZipDir := s.localPathStr + "/" + unZipPath
		// fmt.Println("为您下载: ", dirArr[0], "(os:", dirArr[1], ", arch:", dirArr[2], ") 到您的", unZipPath, "目录下。")
		if strings.Contains(v.FileName, "modbus_linux_amd64") ||
			strings.Contains(v.FileName, "opcua_linux_amd64") ||
			strings.Contains(v.FileName, "opcda_linux_amd64") ||
			strings.Contains(v.FileName, "profinet_linux_amd64") ||
			strings.Contains(v.FileName, "bacnet_linux_amd64") ||
			strings.Contains(v.FileName, "melsec_linux_amd64") ||
			strings.Contains(v.FileName, "iec104_linux_amd64") ||
			strings.Contains(v.FileName, "fins_linux_amd64") ||
			strings.Contains(v.FileName, "mewtocol_linux_amd64") ||
			strings.Contains(v.FileName, "filebeat_linux_amd64") ||
			strings.Contains(v.FileName, "EdgeAgent_") {
			//|| strings.Contains(v.FileName, "modbus_linux_amd32") ||
			//strings.Contains(v.FileName, "opcua_linux_amd32") ||
			//strings.Contains(v.FileName, "opcda_linux_amd32") ||
			//strings.Contains(v.FileName, "profinet_linux_amd32") ||
			//strings.Contains(v.FileName, "bacnet_linux_amd32") ||
			//strings.Contains(v.FileName, "melsec_linux_amd32") ||
			//strings.Contains(v.FileName, "iec104_linux_amd32") ||
			//strings.Contains(v.FileName, "fins_linux_amd32") ||
			//strings.Contains(v.FileName, "mewtocol_linux_amd32") ||
			//strings.Contains(v.FileName, "filebeat_linux_amd32") ||
			//strings.Contains(v.FileName, "EdgeAgent_") {

		} else {
			continue
		}
		errUnZip := utils.UnzipFile(s.localPathStr+"/"+v.FileName, unZipDir)
		if errUnZip != nil {
			fmt.Println("unzip file err: ", errUnZip)
			return errUnZip
		}
		// modfiy broker in mqtt config
		// fmt.Println("配置您的broker到您: ", unZipDir, " 目录下的mqtt.conf文件中。")
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
		// fmt.Println("获取您的broker上的device.sk文件到您: ", unZipDir, " 目录下。")
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

		// update autosmoke config file for bacnet profinet opcda
		if !utils.FileIsExisted("mqtt.conf") {
			_, _ = utils.CopyFile(unZipDir+"/mqtt.conf", "./mqtt.conf")
		}
		if !utils.FileIsExisted("device.sk") {
			_, _ = utils.CopyFile(unZipDir+"/device.sk", "./device.sk")
		}
		if !utils.FileIsExisted("server.crt") {
			_, _ = utils.CopyFile("config/server.crt", "./server.crt")
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
				if s.jobStr == "LeapEdge.agentSign" {
					if strings.Contains(file.Name(), "EdgeAgent_") {
						//if strings.Contains(file.Name(), "windows") {
						//	fmt.Println("暂时不支持windows，待续。。。")
						//	continue
						//}
						errProcessAgent := s.processAgent(file, unZipDir)
						if errProcessAgent != nil {
							return errProcessAgent
						}
					}
				} else if s.jobStr == "LeapEdge.workerSign" { // TODO process worker
					if strings.Compare(file.Name(), "modbus") == 0 ||
						strings.Compare(file.Name(), "opcua") == 0 ||
						strings.Compare(file.Name(), "opcda") == 0 ||
						strings.Compare(file.Name(), "profinet") == 0 ||
						strings.Compare(file.Name(), "bacnet") == 0 ||
						strings.Compare(file.Name(), "melsec") == 0 ||
						strings.Compare(file.Name(), "iec104") == 0 ||
						strings.Compare(file.Name(), "fins") == 0 ||
						strings.Compare(file.Name(), "mewtocol") == 0 ||
						strings.Compare(file.Name(), "filebeat") == 0 {

						//strings.Contains(file.Name(), "opcua_linux_amd64") ||
						//strings.Contains(file.Name(), "opcda_linux_amd64") ||
						//strings.Contains(file.Name(), "profinet_linux_amd64") ||
						//strings.Contains(file.Name(), "bacnet_linux_amd64") ||
						//strings.Contains(file.Name(), "melsec_linux_amd64") ||
						//strings.Contains(file.Name(), "iec104_linux_amd64") ||
						//strings.Contains(file.Name(), "fins_linux_amd64") ||
						//strings.Contains(file.Name(), "mewtocol_linux_amd64") ||
						//strings.Contains(file.Name(), "filebeat_linux_amd64") ||
						//strings.Contains(file.Name(), "modbus_linux_amd32") ||
						//strings.Contains(file.Name(), "opcua_linux_amd32") ||
						//strings.Contains(file.Name(), "opcda_linux_amd32") ||
						//strings.Contains(file.Name(), "profinet_linux_amd32") ||
						//strings.Contains(file.Name(), "bacnet_linux_amd32") ||
						//strings.Contains(file.Name(), "melsec_linux_amd32") ||
						//strings.Contains(file.Name(), "iec104_linux_amd32") ||
						//strings.Contains(file.Name(), "fins_linux_amd32") ||
						//strings.Contains(file.Name(), "mewtocol_linux_amd32") ||
						//strings.Contains(file.Name(), "filebeat_linux_amd32") {
					//if strings.Contains(file.Name(), "modbus"){
						errProcessWorker := s.processWorker(file, unZipDir)
						if errProcessWorker != nil {
							return errProcessWorker
						}
					} else {
						continue
					}
					// else {
					//	errProcessWorker := s.processWorker(file, unZipDir)
					//	if errProcessWorker != nil {
					//		return errProcessWorker
					//	}
					//}
				}

			}
		}
		fmt.Println(dirArr[0], " (os:", dirArr[1]+", arch:"+dirArr[2], ") 处理完毕.")
		fmt.Println("")
		//if dirArr[1] != "windows" {

			if s.jobStr == "LeapEdge.agentSign" {
				data, err := ioutil.ReadFile(unZipDir + "/" + dirArr[0] + dirArr[1] + dirArr[2] + ".log")
				if err != nil {
					//return err
					continue
				}
				if strings.Contains(string(data), "heartbeat") {
					fmt.Println("PASS	可以正常上报心跳。")
					s.result[fileArr[0]] = "PASS"
					passCount += 1
				} else {
					fmt.Println("FAIL	未能正常上报心跳。")
					s.result[fileArr[0]] = "FAIL"
				}
			} else if s.jobStr == "LeapEdge.workerSign" {
				data, err := ioutil.ReadFile(unZipDir + "/stream.log")
				if err != nil {
					//return err
					continue
				}
				if strings.Contains(string(data), "onOut") {
					fmt.Println("PASS	可以正常capture data。")
					s.result[fileArr[0]] = "PASS"
					passCount += 1
				} else {
					fmt.Println("FAIL	未能正常capture data。")
					s.result[fileArr[0]] = "FAIL"
				}
			}

		//} else {
		//	errSaveToFile := utils.SaveToFile([]byte("windows程序无法在ubuntu上运行!!!!"), unZipDir+"/"+unZipPath+".log")
		//	if errSaveToFile != nil {
		//		fmt.Println("errSaveToFile: ", errSaveToFile.Error())
		//	}
		//	fmt.Println("FAIL	不能在此平台运行。")
		//	s.result[fileArr[0]] = "FAIL"
		//}
		totalCount += 1
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
	}
	fmt.Println("totalCount: ", totalCount, " passCount: ", passCount)

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
	fmt.Println("为您运行您的程序: ", file.Name(), " 并在"+s.runtime+"秒后结束其运行。")
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
				var errAgentNew error
				s.ag, errAgentNew = AgentNew(unZipDir, file, s.runtime)
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
			bin := strings.Split(unZipDir, "/")
			//if strings.Contains(file.Name(), bina[1]) {
			if strings.Compare(file.Name(), bin[1]) == 0 {
				var errWorkerNew error
				s.wk, errWorkerNew = WorkerNew(unZipDir, file, s.runtime)
				if errWorkerNew != nil {
					fmt.Println("errWorkerNew: ", errWorkerNew.Error())
					return errWorkerNew
				}

				errStartWorker := s.wk.startWorker()
				if errStartWorker != nil {
					log.Println("errStartWorker: ", errStartWorker.Error())
					return errStartWorker
				}
			}
		}
	}
	return nil
}

func (s *smoke) sendMailToMonitor(msg string, att []string) error {
	fromUser := "leapiot@126.com"
	toUser := []string{"yaohp1@lenovo.com", "leapiot@126.com"}
	subject := "LeapIot Smoke Test"
	var errMail error
	s.ml, errMail = MailNew(fromUser, toUser, subject, att)
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
	fmt.Println("最新的Build / 最新的成功的Build : <", lastBuild.Id, "/", build.Id, ">")

	if s.buildNumberStr != "0" {
		var errSpecial error
		build, errSpecial = s.jk.getSpecialBuild()
		if errSpecial != nil {
			fmt.Println("errSpecial: ", errSpecial.Error())
			return errSpecial
		}
		fmt.Println("为您使用您指定的build: ", build.Id)
	} else {
		fmt.Println("为您使用最新的成功的Build: ", build.Url)
	}

	fmt.Println("")

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

func New(ur string, j string, l string, us string, p string, br string, bn string, sk string, rt string) (*smoke, error) {
	return &smoke{
		urlStr:         ur,
		jobStr:         j,
		localPathStr:   l,
		userStr:        us,
		pwdStr:         p,
		brokerStr:      br,
		buildNumberStr: bn,
		skStr:          sk,
		runtime:        rt,
		result:         make(map[string]string),
	}, nil
}