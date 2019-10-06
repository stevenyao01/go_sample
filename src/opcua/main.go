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
#include "opcua.h"


//typedef struct {
//	char *typeName;
//	char *key;
//	int arrayLength;
//	void *data;
//} UA_Read_Retval;
//typedef struct{
//	char *Identifier;
//	char *Field;
//	char *IdentifierType;
//	int  NamespaceIndex;
//} Ua_Node_Id;
//
//typedef struct {
//	char *Password;
//	char *StoreType;
//	char *KeystoreFilePath;
//	char *Alias;
//	char *SecurityPolicy;
//} Ua_Security;
//
//typedef struct {
//	int MaxChunkCount;
//	int MaxArrayLength;
//	int MaxMessageSize;
//	int MaxStringLength;
//	int MaxChunkSize;
//} Ua_Channel_Config;
//
//typedef struct {
//	char *ResourceUrl;
//	int UseCredenials;
//	int  PollingInterval;
//	char *ApplicationUrl;
//	int  SessionTimeOut;
//	char *ProcessingMode;
//	int  RequestTimeOut;
//	int  ReconnectTime;
//} Ua_Connect_Config;
//
//typedef struct {
//	Ua_Node_Id 			NodeIds;
//	Ua_Security 		Security;
//	Ua_Channel_Config 	ChannelConfig;
//	Ua_Connect_Config   Config;
//} Opc_Ua_Config;
*/
import "C"
import (
	"fmt"
	"unsafe"
	"io/ioutil"
	"encoding/json"
	//"time"
	"time"
	"strconv"
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
func OpcCallback(urr C.Ua_Sub_Node, length int)(){
	//fmt.Println("typeName: ", pRet.typeName)
	//fmt.Println("arrayLength: ", pRet.arrayLength)
	//for i := 0; i < int(pRet.arrayLength); i++ {
	//	fmt.Println("data: ", *(*bool)(unsafe.Pointer(uintptr(pRet.data) + uintptr(i))))
	//}

	addrTypeName := uintptr(unsafe.Pointer(urr.typeName))
	//addrKey := uintptr(unsafe.Pointer(urr.key))
	addrData := uintptr(unsafe.Pointer(urr.data))

	for i := 0; i < length; i++ {
		fmt.Println("arrayLength: ", int(urr.arrayLength))
		fmt.Println("addrTypeName: ", C.GoString((*C.char)(unsafe.Pointer(addrTypeName))))
		//fmt.Println("addrKey: ", C.GoString((*C.char)(unsafe.Pointer(addrKey))))
		fmt.Println("addrData: ", *(*C.bool)(unsafe.Pointer(addrData)))
		pDataAddr := (*C.bool)(unsafe.Pointer(addrData))
		for j := 0; j < int(urr.arrayLength); j++ {
			fmt.Println("data bbbb: ", *(*bool)(unsafe.Pointer(uintptr(unsafe.Pointer(pDataAddr)) + uintptr(j))))
		}
	}

	//addrTypeName := uintptr(unsafe.Pointer(urr.Usn.typeName))
	//addrKey := uintptr(unsafe.Pointer(urr.Usn.key))
	//addrData := uintptr(unsafe.Pointer(urr.Usn.data))
	//
	//for i := 0; i < length; i++ {
	//	fmt.Println("arrayLength: ", int(*urr.Usn.arrayLength))
	//	fmt.Println("addrTypeName: ", C.GoString(*(**C.char)(unsafe.Pointer(addrTypeName + uintptr(i * 8)))))
	//	fmt.Println("addrKey: ", C.GoString(*(**C.char)(unsafe.Pointer(addrKey + uintptr(i * 8)))))
	//	fmt.Println("addrData: ", *(**C.bool)(unsafe.Pointer(addrData + uintptr(i * 8))))
	//	pDataAddr := *(**C.bool)(unsafe.Pointer(addrData + uintptr(i * 8)))
	//	for j := 0; j < int(*urr.Usn.arrayLength); j++ {
	//		fmt.Println("data bbbb: ", *(*bool)(unsafe.Pointer(uintptr(unsafe.Pointer(pDataAddr)) + uintptr(j))))
	//	}
	//}
}

func checkIdentifier(opcUaConfig *OpcUaConfig, exitFlag bool) bool {
	for _, nodeId := range opcUaConfig.NodeIds {
		if len(nodeId.Identifier) > int(opcUaConfig.ChannelConfig.MaxStringLength) {
			fmt.Println("input identifier: ", nodeId.Identifier, "is too long. maxLength is: ", strconv.Itoa(int(opcUaConfig.ChannelConfig.MaxStringLength)))
			exitFlag = true
		}
	}
	return exitFlag
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

	var exitFlag bool = false
	exitFlag = checkIdentifier(opcUaConfig, exitFlag)
	if exitFlag {
		return
	}

	if opcUaConfig.Config.ProcessingMode == "polling" {
		// use poll interface
		p, err := NewOpcPoll(configFileName)
		if err != nil {
			fmt.Println("init opc poll failed! err: ", err.Error())
		}
		start := time.Now()
		for i := 0; i < 1000; i++ {
			p.PollRead(*opcUaConfig)
			//time.Sleep(time.Duration(opcUaConfig.Config.PollingInterval) * time.Millisecond)
			fmt.Println("i: ", i)
		}
		cost := time.Since(start)
		fmt.Println("cost=[%s]", cost)
	} else if opcUaConfig.Config.ProcessingMode == "subscribe" {
		// use subscribe interface
		s, err := NewOpcSubscribe(configFileName)
		if err == nil {
			fmt.Println("init opc subscribe failed!")
		}
		for i := 0; i < 1; i++ {
			s.SubscribeRead(*opcUaConfig)
			time.Sleep(time.Duration(opcUaConfig.Config.PollingInterval) * time.Millisecond)
		}
	} else if opcUaConfig.Config.ProcessingMode == "browser" {
		// use browser interface
		b, err := NewOpcBrowser(configFileName)
		if err == nil {
			fmt.Println("init opc subscribe failed!")
		}
		for i := 0; i < 2; i++ {
			b.BrowserRead()
			time.Sleep(time.Duration(opcUaConfig.Config.PollingInterval) * time.Millisecond)
		}
	}

}
