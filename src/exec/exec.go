package main

import (
    "fmt"
    "os/exec"
    "os"
    "time"
)

func main1() {
    fmt.Println("enter main, pid : ", os.Getpid())
    binary, err := exec.LookPath("./publish_message")
    //binary, err := exec.LookPath("ls")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(binary)
    cmd := exec.Command(binary)
    out, err := cmd.CombinedOutput()
    if err != nil  {
        fmt.Println(err)
    }
    fmt.Println(string(out))
}

func main() {
    fmt.Println("enter publish_message main, pid : ", os.Getpid())
    time.Sleep(10 * time.Second)
    fmt.Println("leave publish_message main, pid : ", os.Getpid())


    //////////////////////////////
    ////var mm map[string]
	//
    //var mm = make(map[string]string)
    //mm["1"] = "yao"
    //mm["2"] = "hai"
    //mm["5"] = "ping"
    //by,err:=json.Marshal(mm)
    //if err!=nil{
    //    log.Error(err.Error())
    //    return
    //}
    //identifier:=string(by)
    //log.Info("yaohp: %s", identifier)
    ////////////////////////////////
}