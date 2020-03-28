package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/jenkins-x/golang-jenkins"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)


var jenkinsUrl = flag.String("url", "http://10.110.180.231:8080/jenkins", "jenkins url.")
var jenkinsJob = flag.String("job", "LeapEdge.agentSign", "jenkins job.")
var localPath = flag.String("local", "agentSign", "download path.")
var user = flag.String("user", "yaohp1", "user name.")
var pwd = flag.String("pwd", "xxxxxx", "password.")
var broker = flag.String("broker", "172.17.170.163:4567", "mqtt server address.")
var buildNumber = flag.String("build", "0", "special build number.")

func saveToFile(msg []byte, fileName string) error {
	var errOpenFile error
	fd, errOpenFile := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if errOpenFile != nil {
		return errOpenFile
	}

	wr := bufio.NewWriter(fd)

	_, errWrite := wr.Write(msg)
	if errWrite != nil {
		return errWrite
	}
	_ = wr.Flush()
	//fmt.Println("write ", wn, " words to ", fileName)
	return nil
}

func downloadCallback(length int64, len int64) () {
	fmt.Println("int: ", length, " len: ", len)
}

func stopProcess(p *os.Process) {
	time.Sleep(20 * time.Second)
	errKill := p.Kill()
	if errKill != nil {
		log.Println("errKill: ", errKill.Error())
	}
	log.Println("kill the test process.")
	return
}

func readStderr(path string, read, write *os.File) {
	defer read.Close()
	defer write.Close()
	var buf = make([]byte, 4*4096)
	for {
		n, err := read.Read(buf)
		if err != nil {
			log.Println("stderr read error: %s", err.Error())
			return
		}
		errSaveToFile := saveToFile(buf[:n], path+"agent.log")
		if errSaveToFile != nil {
			log.Println("errSaveToFile: ", errSaveToFile.Error())
		}
	}
}

func main1() {

}

func main() {
	flag.Parse()

	//result := make(map[string]string)
	//result["linuxarm32"] = "PASS"
	//result["linuxarm64"] = "FAILED"
	//var testStr []string
	//for k, v := range result {
	//	testStr = append(testStr, k + " : " + v)
	//}
	//log.Println("testStr: ", testStr)
	//errSendMail := sendMailToMonitor(testStr)
	//if errSendMail != nil {
	//	log.Println("errSendMail: ", errSendMail.Error())
	//}
	//return

	auth := &gojenkins.Auth{
		Username: *user,
		ApiToken: *pwd,
	}

	jenkins := gojenkins.NewJenkins(auth, *jenkinsUrl)
	job, err := jenkins.GetJob(*jenkinsJob)

	//fmt.Println("job:", job)

	lastBuild, errLastBuild := jenkins.GetLastBuild(job)
	if errLastBuild != nil {
		panic(errLastBuild)
	}

	build, errLastSuccessBuild := jenkins.GetLastSuccessfulBuild(job)
	if errLastSuccessBuild != nil {
		panic(errLastSuccessBuild)
	}

	fmt.Println("lastBuild / lastSuccessBuild : ", lastBuild.Id, "/", build.Id)
	fmt.Println("Download last success build: ", build.Url)

	if *buildNumber != "0" {
		num, _ := strconv.Atoi(*buildNumber)
		var errSpecial error
		build, errSpecial = jenkins.GetBuild(job, num)
		if errSpecial != nil {
			panic(errSpecial)
		}
		log.Println("use special version: ", build.Id)
	}

	var output []byte
	output, err = jenkins.GetBuildConsoleOutput(build)
	if err != nil {
		panic(err)
	}
	if !strings.Contains(string(output), "Finished:") {
		panic(fmt.Errorf("当前job正在运行,build.Number=%d", build.Number))
	}

	errMkdir := os.Mkdir(*localPath, os.ModePerm)
	if errMkdir != nil {
		if !strings.Contains(errMkdir.Error(), "file exists") {
			fmt.Println("errMkdir: ", errMkdir.Error())
		}
	}

	for _, v := range build.Artifacts {
		zip, errArtifact := jenkins.GetArtifact(build, v)
		if errArtifact != nil {
			fmt.Println("errArtifact: ", errArtifact)
		}
		errSaveToFile := saveToFile(zip, *localPath+"/"+v.FileName)
		if errSaveToFile != nil {
			fmt.Println("errSaveToFile: ", errSaveToFile)
		}

		// unzip file
		fileArr := strings.Split(v.FileName, "-")
		if !strings.Contains(fileArr[0], "EdgeAgent_") {
			continue
		}
		dirArr := strings.Split(fileArr[0], "_")
		unZipPath := dirArr[1] + dirArr[2]
		unZipDir := *localPath + "/" + unZipPath
		errUnZip := unzipFile(*localPath+"/"+v.FileName, unZipDir)
		if errUnZip != nil {
			fmt.Println("unzip file err: ", errUnZip)
			return
		}
		// modfiy broker in mqtt config
		errModify := modifyConfig(unZipDir)
		if errModify != nil {
			fmt.Println("errModify: ", errModify.Error())
		}

		// get broker sk file
		errSk := getSk(unZipDir)
		if errSk != nil {
			fmt.Println("errSk: ", errSk.Error())
		}

		// get agent, add exec right, start it.
		files, err := ioutil.ReadDir(unZipDir) //读取目录下文件
		if err != nil {
			return
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			} else {
				if strings.Contains(file.Name(), "EdgeAgent_") {
					fmt.Println("file: ", file.Name())
					// get agent, add exec right, start it.
					files, err := ioutil.ReadDir(unZipDir) //读取目录下文件
					if err != nil {
						return
					}
					for _, file := range files {
						if file.IsDir() {
							continue
						} else {
							if strings.Contains(file.Name(), "EdgeAgent_") {
								fmt.Println("file: ", file.Name())

								errAgent := startAgent(unZipDir, file)
								if errAgent != nil {
									log.Println("errAgent: ", errAgent.Error())
								}
								return
							}
						}
					}
				}
			}
		}
	}

	//
	//errSendMail := sendMailToMonitor("测试结果！！！！")
	//if errSendMail != nil {
	//	log.Println("errSendMail: ", errSendMail.Error())
	//}
}

func startAgent(unZipDir string, file os.FileInfo) error {
	// add exec right
	errChmod := os.Chmod(unZipDir+file.Name(), 0755)
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
	binary := unZipDir + file.Name()
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
	go readStderr(unZipDir, read, write)
	go stopProcess(pro)
	ps, errWait := pro.Wait()
	if errWait != nil {
		log.Println("wait worker error:%s", err.Error())
		return errWait
	}
	log.Println("ps: ", ps.String())
	return nil
}

func sendMailToMonitor(msg string) error {
	fromUser := "leapiot@126.com"
	toUser := []string{"yaohp1@lenovo.com", "leapiot@126.com"}
	subject := "LeapIot Smoke Test"
	err := SendMail(fromUser, toUser, subject, msg)
	if err != nil {
		log.Println("发送邮件失败, err: ", err.Error())
		return err
	}
	log.Println("发送邮件成功")
	return nil
}

func modifyConfig(unZipDir string) error {
	// modify local mqtt.conf broker
	conf, err := loadConfigFile(unZipDir + "/mqtt.conf")
	if err != nil {
		panic(err)
	}
	fmt.Println(conf)
	var buf bytes.Buffer
	for k, v := range conf {
		if k == "broker" {
			buf.WriteString(k + "=" + *broker + "\n")
		} else {
			buf.WriteString(k + "=" + v + "\n")
		}
	}
	errWrite := ioutil.WriteFile(unZipDir+"/mqtt.conf", buf.Bytes(), 0644)
	if errWrite != nil {
		fmt.Println("errWrite: ", errWrite)
		return errWrite
	}
	return nil
}

func unzipFile(zipFileName string, targetPath string) error {
	// judge zip file
	if lIsZip(zipFileName) {
		fmt.Println("the target is zip file.")
	} else {
		fmt.Println("the target is not zip file.")
		return errors.New("input file is not zip file")
	}
	// unzip packet
	errUnZip := lUnZip(zipFileName, targetPath)
	if errUnZip != nil {
		fmt.Println("errUnZip: ", errUnZip)
	} else {
		fmt.Println("UnZip file success.")
	}
	return errUnZip
}

func rebuildJob(err error, jenkins *gojenkins.Jenkins, job gojenkins.Job, build gojenkins.Build, output []byte) error {
	params := make(url.Values)
	params.Add("branch", "dev")
	params.Add("buildVersion", "1.6")
	params.Add("goVersion", "go1.9.2")
	err = jenkins.Build(job, params)
	if err != nil {
		panic(err)
	}
	build, err = jenkins.GetLastBuild(job)
	if err != nil {
		panic(err)
	}
	//fmt.Println("build:", build)
	for i := 0; i < 60; i++ {
		output, err = jenkins.GetBuildConsoleOutput(build)
		if err != nil {
			panic(err)
		}
		if strings.Contains(string(output), "Finished:") {
			break
		}
		time.Sleep(time.Second)
	}
	if strings.Contains(string(output), "Finished: FAILURE") {
		outputs := strings.Split(string(output), "\n")
		panic(errors.New("build fail:" + strings.Join(outputs[len(outputs)-5:len(outputs)-2], "\n")))
	}
	fmt.Println("build ok:", string(output))
	return err
}

func getSk(dir string) error {
	strArr := strings.Split(*broker, ":")
	errDownload := downloadFile("http://"+strArr[0]+":8080/services/certificate/download?uid=1", dir+"/device.sk", downloadCallback)
	if errDownload != nil {
		fmt.Println("errDownload: ", errDownload.Error())
		return errDownload
	}
	return nil
}

var DEPTH = 5

func walkDir(dirpath string, depth int) {
	if depth > DEPTH { //大于设定的深度
		return
	}
	files, err := ioutil.ReadDir(dirpath) //读取目录下文件
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			walkDir(dirpath+"/"+file.Name(), depth+1)
			continue
		} else {
			fmt.Println("file: ", file.Name())
		}
	}
}

func walkDirOne(dirpath string) {
	files, err := ioutil.ReadDir(dirpath) //读取目录下文件
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			fmt.Println("file: ", file.Name())
		}
	}
}
