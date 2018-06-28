package main
import (
	"fmt"
	"net"
	"encoding/json"
)
type InstallInfo struct{
	Seq 	   	string		`json:"seq"`
	Cmd 	 	string   	`json:"cmd"`
	WorkerId	string   	`json:"workerid"`
	Url	     	string   	`json:"url"`
	Type	 	string   	`json:"type"`
	MD5			string		`json:"md5"`
}

func doServerStuff1(conn net.Conn){
	remote:=conn.RemoteAddr().String()
	fmt.Println(remote," connected!")
	for {
		buf:=make([]byte,512)
		size,err:=conn.Read(buf)
		if err!=nil{
			fmt.Println("Read Error:",err.Error());
			return
		}
		//fmt.Println("data from client:",string(buf),"size:",size)
		var info InstallInfo
		err=json.Unmarshal(buf[:size],&info)
		if err!=nil{
			fmt.Println("Unmarshal Error:",err.Error());
			return
		}
		fmt.Println("client report after Unmarshal:",info)
		fmt.Println("get workerid : " + info.WorkerId + " report msg.")
		//buf,err=json.Marshal(info)
		//if err!=nil{
		//	fmt.Println("Marshal Error:",err.Error());
		//	return
		//}
		//conn.Write(buf)
		conn.Close()
		break
	}
}
func main(){
	fmt.Println("Starting the server...")
	listener,err:=net.Listen("tcp","0.0.0.0:7777")
	if err!=nil{
		fmt.Println("Listen Error:",err.Error())
		return
	}
	for{
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println("Accept Error:",err.Error())
			return
		}
		go doServerStuff1(conn)
	}
}
