package smoke

import (
	"fmt"
	"github.com/jenkins-x/golang-jenkins"
	"strconv"
)

type jenkins struct {
	urlStr         string
	jobStr         string
	localPathStr   string
	userStr        string
	pwdStr         string
	buildNumberStr string
	jenk             *gojenkins.Jenkins
	jb             gojenkins.Job
	build          gojenkins.Build
}

func (j *jenkins) init() error {
	auth := &gojenkins.Auth{
		Username: j.userStr,
		ApiToken: j.pwdStr,
	}

	j.jenk = gojenkins.NewJenkins(auth, j.urlStr)
	var errJob error
	j.jb, errJob = j.jenk.GetJob(j.jobStr)
	if errJob != nil {
		fmt.Println("errJob: ", errJob.Error())
		return errJob
	}
	return nil
}

func (j *jenkins) getJenkins() *gojenkins.Jenkins {
	return j.jenk
}

func (j *jenkins) getJob(jk *gojenkins.Jenkins) gojenkins.Job {
	return j.jb
}

func (j *jenkins) getLastBuild() (gojenkins.Build, error) {
	lastBuild, errLastBuild := j.jenk.GetLastBuild(j.jb)
	if errLastBuild != nil {
		fmt.Println("errLastBuild: ", errLastBuild.Error())
		return lastBuild, errLastBuild
	}
	return lastBuild, nil
}

func (j *jenkins) getLastSuccessfulBuild() (gojenkins.Build, error) {
	lastSuccessfulBuild, errLastSuccessfulBuild := j.jenk.GetLastSuccessfulBuild(j.jb)
	if errLastSuccessfulBuild != nil {
		fmt.Println("errLastBuild: ", errLastSuccessfulBuild.Error())
		return lastSuccessfulBuild, errLastSuccessfulBuild
	}
	return lastSuccessfulBuild, nil
}

func (j *jenkins) getSpecialBuild() (gojenkins.Build, error) {
	bn, errGetSpecialBuild := strconv.Atoi(j.buildNumberStr)
	if errGetSpecialBuild != nil {
		fmt.Println("errGetSpecialBuild: ", errGetSpecialBuild.Error())
	}
	return j.jenk.GetBuild(j.jb, bn)
}

func (j *jenkins) setCurrentBuild(b gojenkins.Build) {
	j.build = b
	return
}

func (j *jenkins) getCurrentBuild() gojenkins.Build {
	return j.build
}


func (j *jenkins) getConsoleOutput(build gojenkins.Build) (o []byte, err error) {
	output, errConsole := j.jenk.GetBuildConsoleOutput(build)
	if errConsole != nil {
		fmt.Println("errConsole: ", errConsole.Error())
		return output, errConsole
	}
	return output, nil
}

func (j *jenkins) getArtifact(af gojenkins.Artifact) ([]byte, error) {
	return j.jenk.GetArtifact(j.build, af)
}

func JenkinsNew(ur string, j string, l string, us string, p string, b string) (*jenkins, error) {
	jen := &jenkins{
		urlStr:         ur,
		jobStr:         j,
		localPathStr:   l,
		userStr:        us,
		pwdStr:         p,
		buildNumberStr: b,
	}
	errInit := jen.init()
	if errInit != nil {
		fmt.Println("errInit: ", errInit.Error())
		return jen, errInit
	}
	return jen, nil
}
