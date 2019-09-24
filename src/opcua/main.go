package main

/**
 * @Package Name: yao
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-9-2 下午3:56
 * @Description:
 */

import "C"

/*
typedef struct {
	char *typeName;
	char *key;
	int arrayLength;
	void *data;
} UA_Read_Retval;
typedef struct{
	char *Identifier;
	char *Field;
	char *IdentifierType;
	int  NamespaceIndex;
} Ua_Node_Id;

typedef struct {
	char *Password;
	char *StoreType;
	char *KeystoreFilePath;
	char *Alias;
	char *SecurityPolicy;
} Ua_Security;

typedef struct {
	int MaxChunkCount;
	int MaxArrayLength;
	int MaxMessageSize;
	int MaxStringLength;
	int MaxChunkSize;
} Ua_Channel_Config;

typedef struct {
	char *ResourceUrl;
	int UseCredenials;
	int  PollingInterval;
	char *ApplicationUrl;
	int  SessionTimeOut;
	char *ProcessingMode;
	int  RequestTimeOut;
	int  ReconnectTime;
} Ua_Connect_Config;

typedef struct {
	Ua_Node_Id 			NodeIds;
	Ua_Security 		Security;
	Ua_Channel_Config 	ChannelConfig;
	Ua_Connect_Config   Config;
} Opc_Ua_Config;
*/
import "C"
import (
	"fmt"
	"unsafe"
	"io/ioutil"
	"encoding/json"
	//"time"
)

// opcua config
const (
	dataCacheSize = 1024 * 1024
	configFileName = "opcua.conf"
	//opcString = "String"
	//opcNumeric = "Numeric"
	//opcUUID = "UUID"
	//opcOpaque = "Opaque"
)

//export OpcCallback
func OpcCallback(pRet C.UA_Read_Retval)(){
	fmt.Println("typeName: ", pRet.typeName)
	fmt.Println("arrayLength: ", pRet.arrayLength)
	for i := 0; i < int(pRet.arrayLength); i++ {
		fmt.Println("data: ", *(*bool)(unsafe.Pointer(uintptr(pRet.data) + uintptr(i))))
	}
}

func main(){

	opcUaConfig, err := NewOpcUaConfig()
	if err != nil {
		fmt.Println("get opc ua config error.")
	}

	data, err := ioutil.ReadFile(configFileName)
	if err != nil {
		fmt.Println("read config file failed, filename: ", configFileName)
	}

	opcUaConfig.Fix()

	if err = json.Unmarshal(data, opcUaConfig); err != nil {
		fmt.Println("parse json error.")
	}
	fmt.Println("")

	p, err := NewOpcPoll(configFileName)
	if err != nil {
		fmt.Println("init opc poll failed! err: ", err.Error())
	}
	for i := 1; i < 2; i++ {
		p.PollRead(*opcUaConfig)
		//time.Sleep(1 * time.Second)
	}



	//// use poll interface
	//p, err := NewOpcPoll(configFileName)
	//if err == nil {
	//	fmt.Println("init opc poll failed!")
	//}
	//for i := 0; i < 2; i++ {
	//	p.PollRead()
	//	time.Sleep(1 * time.Second)
	//}

	//// use subscribe interface
	//s, err := NewOpcSubscribe(configFileName)
	//if err == nil {
	//	fmt.Println("init opc subscribe failed!")
	//}
	//for i := 0; i < 1; i++ {
	//	s.SubscribeRead()
	//	time.Sleep(1 * time.Second)
	//}

	//// use browser interface
	//b, err := NewOpcBrowser(configFileName)
	//if err == nil {
	//	fmt.Println("init opc subscribe failed!")
	//}
	//for i := 0; i < 2; i++ {
	//	b.BrowserRead()
	//	time.Sleep(1 * time.Second)
	//}

}
