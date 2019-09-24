package main

/**
 * @Package Name: config
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-9-13 上午11:27
 * @Description:
 */

import "C"
/*
//#cgo LDFLAGS: -lopen62541
//#cgo CFLAGS: -I F:/Project/Go/src/github.com/worker/opcua-c/open62541/include
//#cgo LDFLAGS: F:/Project/Go/src/github.com/worker/opcua-c/open62541/build/bin/open62541.lib
//#cgo LDFLAGS: -L $GOPATH/src/github.com/worker/opcua-c/open62541/build/bin/Debug/ -lopen62541

#cgo CFLAGS: -I ./open62541/include
#cgo CFLAGS: -I ./open62541/arch
#cgo CFLAGS: -I ./open62541/plugins/include
#cgo CFLAGS: -I ./open62541/build-shared64/src_generated
#cgo LDFLAGS: -L./open62541/build-shared64/bin/Debug -lopen62541

typedef struct {
    char *typeName;
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

void Polling(Opc_Ua_Config pRet);

//#include "open62541/include/open62541/client_config_default.h"
//#include "open62541/include/open62541/client_highlevel.h"
//#include "open62541/include/open62541/client_subscriptions.h"
//#include "open62541/include/open62541/plugin/log_stdout.h"

#include <open62541/client_config_default.h>
#include <open62541/client_highlevel.h>
#include <open62541/client_subscriptions.h>
#include <open62541/plugin/log_stdout.h>

#include <signal.h>
#include <stdlib.h>

UA_Boolean running = true;

static void stopHandler(int sign) {
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_CLIENT, "Received Ctrl-C");
    running = 0;
}

void
Polling(Opc_Ua_Config pRet) {
	printf("hello c language.");
	printf("identifier: %s", pRet.NodeIds.Identifier);
	printf("identifier: %d", pRet.NodeIds.NamespaceIndex);
//  signal(SIGINT, stopHandler);
//
//	UA_Client *client = UA_Client_new();
//	UA_ClientConfig *cc = UA_Client_getConfig(client);
//	UA_ClientConfig_setDefault(cc);
//	cc->timeout = 1000;
//
//	//UA_Variant *value;
//	//UA_Variant_init(value);
//	UA_Variant *value = UA_Variant_new();
//	while(running) {
//		UA_StatusCode retval = UA_Client_connect(client, "opc.tcp://10.111.66.220:48030");
//		if(retval != UA_STATUSCODE_GOOD) {
//			UA_LOG_ERROR(UA_Log_Stdout, UA_LOGCATEGORY_CLIENT, "Not connected. Retrying to connect in 1 second");
//
//			UA_sleep_ms(1000);
//			continue;
//		}
////        const UA_NodeId nodeId =
////                UA_NODEID_NUMERIC(0, UA_NS0ID_SERVER_SERVERSTATUS_CURRENTTIME);
////        retval = UA_Client_readValueAttribute(client, nodeId, &value);
//		retval = UA_Client_readValueAttribute(client, UA_NODEID_STRING(2, "Demo.Static.Arrays.Boolean"), value);
//		if(retval == UA_STATUSCODE_BADCONNECTIONCLOSED) {
//			UA_LOG_ERROR(UA_Log_Stdout, UA_LOGCATEGORY_CLIENT, "Connection was closed. Reconnecting ...");
//			continue;
//		}
//		if(retval == UA_STATUSCODE_GOOD) {
//			pRet->typeName = value->type->typeName;
//			pRet->arrayLength = value->arrayLength;
//			memcpy(pRet->data, value->data, 8);
//		}
//
//		UA_Variant_clear(value);
//		//UA_sleep_ms(1000);
//		UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua program loop.");
//		break;
//	};
//
//	UA_Variant_clear(value);
//	UA_Client_delete(client);
//
//	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua program exit success.");
	return;
}
*/
import "C"

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/go_sample/src/config/OpcUaConfig"
)

// opcua config
const (
	dataCacheSize  = 1024 * 1024
	configFileName = "opcua.conf"
)

func main() {
	opcUaConfig, err := OpcUaConfig.NewOpcUaConfig()
	if err != nil {
		fmt.Println("get opc ua config error.")
	}

	data, err := ioutil.ReadFile(configFileName)
	if err != nil {
		fmt.Println("read config file failed, filename: ", configFileName)
	}

	opcUaConfig.Fix()

	if err = json.Unmarshal(data, opcUaConfig); err != nil {
		fmt.Println("parse json error: ", err.Error())
	}

	fmt.Println("config: ", opcUaConfig.String())

	///////////////////

	//for i := 0; i < len(OpcUaConfig.NodeIds); i++ {
	//	fmt.Println("nodeids: ", opcUaConfig.NodeId[i])
	//}
	for _, nodeId := range opcUaConfig.NodeIds {
		fmt.Println("nodeId: ", nodeId.Identifier)
	}




	var pDetectInfo C.Opc_Ua_Config
	//pDetectInfo.data = unsafe.Pointer((*C.void)(C.malloc(dataCacheSize)))
	//if pDetectInfo.data == nil {
	//	fmt.Println("go malloc data failed.")
	//}
	//defer C.free(unsafe.Pointer(pDetectInfo.data))

	fmt.Println("pDetectInfo: ", pDetectInfo)
	//pDetectInfo.NodeIds->Identifier = "asdfsdf"
	pDetectInfo.NodeIds.NamespaceIndex = 10
	C.Polling(pDetectInfo)
}
