package main

import (
	"flag"
	"fmt"
	"github.com/go_sample/src/autosmoke/smoke"
)

const (
	coreDumpFilename = "smoke.dump"
)

var jenkinsUrl = flag.String("url", "http://10.110.180.231:8080/jenkins", "jenkins url.")
var jenkinsJob = flag.String("job", "LeapEdge.agentSign", "jenkins job.")
var localPath = flag.String("local", "agentSign", "download path.")
var user = flag.String("user", "yaohp1", "user name.")
var pwd = flag.String("pwd", "xxxxxx", "password.")
var broker = flag.String("broker", "172.17.203.36:4567", "mqtt server address.")
var buildNumber = flag.String("build", "0", "special build number.")
var sk = flag.String("sk", "", "device.sk")
var runTime = flag.String("runtime", "20", "program run time.")

func main() {
	//defer func(){
	//	err := recover()
	//	errCore := utils.CoreDump(coreDumpFilename, err)
	//	if errCore != nil {
	//		fmt.Println("errCore: ", errCore.Error())
	//	}
	//}()

	flag.Parse()

	s, _ := smoke.New(*jenkinsUrl, *jenkinsJob, *localPath, *user, *pwd, *broker, *buildNumber, *sk, *runTime)
	errSmoke := s.Start()
	if errSmoke != nil {
		fmt.Println("errSmoke: ", errSmoke.Error())
	}
}
