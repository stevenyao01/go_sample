package main

import(
	"flag"
	"fmt"
	"github.com/go_sample/src/autosmoke/smoke"
	"github.com/go_sample/src/autosmoke/utils"
)

const (
	coreDumpFilename = "smoke.dump"
)

var jenkinsUrl = flag.String("url", "http://10.110.180.231:8080/jenkins", "jenkins url.")
var jenkinsJob = flag.String("job", "LeapEdge.agentSign", "jenkins job.")
var localPath = flag.String("local", "agentSign", "download path.")
var user = flag.String("user", "yaohp1", "user name.")
var pwd = flag.String("pwd", "xxxxxx", "password.")
var broker = flag.String("broker", "172.17.170.163:4567", "mqtt server address.")
var buildNumber = flag.String("build", "0", "special build number.")
var sk = flag.String("sk", "", "device.sk")

func main() {
	defer func(){
		err := recover()
		errCore := utils.CoreDump(coreDumpFilename, err)
		if errCore != nil {
			fmt.Println("errCore: ", errCore.Error())
		}
	}()

	flag.Parse()

	s, _ := smoke.New(*jenkinsUrl, *jenkinsJob, *localPath, *user, *pwd, *broker, *buildNumber, *sk)
	errSmoke := s.Start()
	if errSmoke != nil {
		fmt.Println("errSmoke: ", errSmoke.Error())
	}
}
