package main

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

#include <open62541/client_config_default.h>
#include <open62541/client_highlevel.h>
#include <open62541/client_subscriptions.h>
#include <open62541/plugin/log_stdout.h>

#include <signal.h>
#include <stdlib.h>
#include "opcua.h"

void Subscribe(Opc_Ua_Config *Ua_Config, int len);

extern void OpcCallback(Ua_Sub_Node retval, int length);

UA_Boolean running1 = true;
Opc_Ua_Config *opcUaConfig;
int nodeSize = 0;

static void stopHandler(int sign) {
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "Received Ctrl-C");
    running1 = 0;
}

static void
handler_dataChanged(UA_Client *client, UA_UInt32 subId, void *subContext,
                           UA_UInt32 monId, void *monContext, UA_DataValue *value) {

    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "handler_dataChanged!!!");

    //////////////////////////////////////////
    //int sizeType = 0;
    //int sizeData = 0;
    //UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "currentTime has changed!");
    //Ua_Read_Retval retvalaa;
	//
    ////retvalaa = newOPcUaRetval(2);
	//
    //// typeName
    //sizeType = strlen(value->value.type->typeName);
    ////memcpy(retvalaa.Usn->typeName[0], value->value.type->typeName, sizeType);
    //UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "type size: %d", sizeType);
	//
	//
    //// arrayLength
    ////retvalaa.Usn->arrayLength[0] = value->value.arrayLength;
    ////UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "arrayLength: %d", retvalaa.Usn->arrayLength[0]);
	//
    //// alloc mem for data
    //sizeData = sizeof(bool) * value->value.arrayLength;
    ////memcpy(retvalaa.Usn->data[0], value->value.data, sizeData);
    //UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "data size: %d", sizeData);
    ////retval.Usn->data = (void *)malloc(sizeData);
    ////memset(retval.Usn->data, 0 , sizeData);
    ////memcpy(retval.Usn->data, value->value.data, sizeData);
	//
    ////// call go func
    ////OpcCallback(retval, 1);
	//
    //// release mem
    ////free(retval->typeName);
    ////free(retval.data);
    ////deleteOpcUaRetval(retvalaa, 2);
    /////////////////////////////////

    ///////////////////////////////////////
    int sizeType = 0;
    int sizeData = 0;
    int sizeKey = 0;
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "data has changed!");
    Ua_Sub_Node retval;



    // typeName
    sizeType = strlen(value->value.type->typeName);
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "type size: %d", sizeType);
    retval.typeName = malloc(sizeType);
    memset(retval.typeName, 0, sizeType);
    memcpy(retval.typeName, value->value.type->typeName, sizeType);
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "typeName %s", retval.typeName);


    // arrayLength
    retval.arrayLength = value->value.arrayLength;
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "arrayLength: %d", retval.arrayLength);

    // alloc mem for key
    sizeKey = strlen((char*)(monContext));
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "key size: %d", sizeKey);
    retval.key = malloc(sizeKey);
    memset(retval.key, 0 , sizeKey);
    memcpy(retval.key, monContext, sizeKey);

    // alloc mem for data
    sizeData = sizeof(bool) * value->value.arrayLength;
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "data size: %d", sizeData);
    retval.data = (void *)malloc(sizeData);
    memset(retval.data, 0 , sizeData);
    memcpy(retval.data, value->value.data, sizeData);

    // call go func
    OpcCallback(retval, 1);

    // release mem
    free(retval.typeName);
    free(retval.data);

    ////////////////////////////////////////////



    //int sizeType = 0;
    //int sizeData = 0;
    //UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "currentTime has changed!");
    //UA_Read_Retval retval;
	//
    //// typeName
    //sizeType = strlen(value->value.type->typeName);
    //UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "type size: %d", sizeType);
    //char str[32];
    //memcpy(str, value->value.type->typeName, sizeType);
    //retval.typeName = str;
    //UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "typeName %s", retval.typeName);
	//
	//
    //// arrayLength
    //retval.arrayLength = value->value.arrayLength;
    //UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "arrayLength: %d", retval.arrayLength);
	//
    //// alloc mem for data
    //sizeData = sizeof(bool) * value->value.arrayLength;
    //UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "data size: %d", sizeData);
    //retval.data = (void *)malloc(sizeData);
    //memset(retval.data, 0 , sizeData);
    //memcpy(retval.data, value->value.data, sizeData);
	//
    //// call go func
    //OpcCallback(retval);
	//
    //// release mem
    ////free(retval->typeName);
    //free(retval.data);

    if(UA_Variant_hasScalarType(&value->value, &UA_TYPES[UA_TYPES_DATETIME])) {
        UA_DateTime raw_date = *(UA_DateTime *) value->value.data;
        UA_DateTimeStruct dts = UA_DateTime_toStruct(raw_date);
        UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND,
                    "date is: %02u-%02u-%04u %02u:%02u:%02u.%03u",
                    dts.day, dts.month, dts.year, dts.hour, dts.min, dts.sec, dts.milliSec);
    }
}

static void
deleteSubscriptionCallback(UA_Client *client, UA_UInt32 subscriptionId, void *subscriptionContext) {
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND,
                "Subscription Id %u was deleted", subscriptionId);
}

static void
subscriptionInactivityCallback (UA_Client *client, UA_UInt32 subId, void *subContext) {
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "Inactivity for subscription %u", subId);
}

static void
stateCallback (UA_Client *client, UA_ClientState clientState) {
    switch(clientState) {
        case UA_CLIENTSTATE_DISCONNECTED:
            UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "The client is disconnected");
        break;
        case UA_CLIENTSTATE_WAITING_FOR_ACK:
            UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "Waiting for ack");
        break;
        case UA_CLIENTSTATE_CONNECTED:
            UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND,
                        "A TCP connection to the server is open");
        break;
        case UA_CLIENTSTATE_SECURECHANNEL:
            UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND,
                        "A SecureChannel to the server is open");
        break;
        case UA_CLIENTSTATE_SESSION:{
            UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "A session with the server is open");

			for (int i = 0; i < nodeSize; i++) {
				UA_CreateSubscriptionRequest request = UA_CreateSubscriptionRequest_default();
				UA_CreateSubscriptionResponse response = UA_Client_Subscriptions_create(client, request,
					NULL, NULL, deleteSubscriptionCallback);

				if(response.responseHeader.serviceResult == UA_STATUSCODE_GOOD)
					UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND,
						"Create subscription succeeded, id %u", response.subscriptionId);
				else
					return;

//            UA_MonitoredItemCreateRequest monRequest =
//                UA_MonitoredItemCreateRequest_default(UA_NODEID_NUMERIC(0, UA_NS0ID_SERVER_SERVERSTATUS_CURRENTTIME));

				UA_MonitoredItemCreateRequest monRequest =
					UA_MonitoredItemCreateRequest_default(UA_NODEID_STRING(opcUaConfig->NodeIds->NamespaceIndex[i], opcUaConfig->NodeIds->Identifier[i]));

				UA_MonitoredItemCreateResult monResponse =
					UA_Client_MonitoredItems_createDataChange(client, response.subscriptionId, UA_TIMESTAMPSTORETURN_BOTH,
					monRequest, opcUaConfig->NodeIds->Field[i], handler_dataChanged, NULL);
				if(monResponse.statusCode == UA_STATUSCODE_GOOD)
					UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND,"Monitoring UA_NS0ID_SERVER_SERVERSTATUS_CURRENTTIME', id %u",
						monResponse.monitoredItemId);
				}

			}
		break;
		case UA_CLIENTSTATE_SESSION_RENEWED:
			UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "A session with the server is open (renewed)");
		break;
		case UA_CLIENTSTATE_SESSION_DISCONNECTED:
			UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "Session disconnected");
		break;
	}
	return;
}

void
Subscribe(Opc_Ua_Config *Ua_Config, int len) {
	signal(SIGINT, stopHandler);

	opcUaConfig = Ua_Config;
	nodeSize = len;

	UA_Client *client = UA_Client_new();
	UA_ClientConfig *cc = UA_Client_getConfig(client);
	UA_ClientConfig_setDefault(cc);

	cc->timeout = Ua_Config->Config->RequestTimeOut;
	// set channelconfig
	UA_ConnectionConfig uc;
	uc.protocolVersion = 0;
	uc.sendBufferSize = Ua_Config->ChannelConfig->MaxChunkSize;
	uc.recvBufferSize = Ua_Config->ChannelConfig->MaxChunkSize;
	uc.maxMessageSize = Ua_Config->ChannelConfig->MaxMessageSize;
	uc.maxChunkCount = Ua_Config->ChannelConfig->MaxChunkCount;
	cc->localConnectionConfig = uc;

	// set config
	cc->clientDescription.applicationUri =  UA_STRING_ALLOC(Ua_Config->Config->ApplicationUrl);

	// set requestedSessionTimeout
	cc->requestedSessionTimeout = Ua_Config->Config->SessionTimeOut;

	// state callback
	cc->stateCallback = stateCallback;
	// subscribe callback
	cc->subscriptionInactivityCallback = subscriptionInactivityCallback;

	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "subscribe loop begin..");
	while(running1) {
		UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "subscribe loop 001..");
		UA_StatusCode retval;
		if (Ua_Config->Config->UseCredenials == 1) {
			retval = UA_Client_connect_username(client, Ua_Config->Config->ResourceUrl, Ua_Config->Credenials->userName, Ua_Config->Credenials->passWord);
		} else {
			retval = UA_Client_connect(client, Ua_Config->Config->ResourceUrl);
		}
		if(retval != UA_STATUSCODE_GOOD) {
			UA_LOG_ERROR(UA_Log_Stdout, UA_LOGCATEGORY_CLIENT, "Not connected. Retrying to connect in 1 second");

			UA_sleep_ms(Ua_Config->Config->ReconnectTime);
			continue;
		}

		UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "subscribe loop 002..");
		UA_Client_run_iterate(client, Ua_Config->Config->ReconnectTime);
	};

	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "subscribe loop end..");
	UA_Client_delete(client);
	return;
}
*/
import "C"

import (
	"fmt"
	"unsafe"
	//"time"
)

/**
 * @Package Name: yao
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-9-2 下午3:56
 * @Description:
 */

type OpcSubscribe struct {
	filename string
}

func (s *OpcSubscribe) SubscribeRead(opcUaConfig OpcUaConfig)(){
	fmt.Println("=============================== cgo opcua subscribe ===========================")
	nodeIdLength := (C.int)(len(opcUaConfig.NodeIds))
	//urr := C.newOpcUaRetval(nodeIdLength)
	//defer C.deleteOpcUaRetval(urr, (C.int)(len(opcUaConfig.NodeIds)))

	// test trans opc ua config
	uaConfig := C.newOpcUaConfig(nodeIdLength)
	defer C.deleteOpcUaConfig(uaConfig, (C.int)(len(opcUaConfig.NodeIds)))

	i := 0
	for _, nodeId := range opcUaConfig.NodeIds {
		addrIdentifier := uintptr(unsafe.Pointer(uaConfig.NodeIds.Identifier))
		C.strcpy(*(**C.char)(unsafe.Pointer(addrIdentifier + uintptr(8 * i))), C.CString(nodeId.Identifier))
		//fmt.Println("Identifier: ", *(**C.char)(unsafe.Pointer(addrIdentifier + uintptr(i * 8))))

		addrField := uintptr(unsafe.Pointer(uaConfig.NodeIds.Field))
		C.strcpy(*(**C.char)(unsafe.Pointer(addrField + uintptr(8 * i))), C.CString(nodeId.Field))
		//fmt.Println("Field: ", *(**C.char)(unsafe.Pointer(addrField + uintptr(i * 8))))

		addrIdentifierType := uintptr(unsafe.Pointer(uaConfig.NodeIds.IdentifierType))
		C.strcpy(*(**C.char)(unsafe.Pointer(addrIdentifierType + uintptr(8 * i))), C.CString(nodeId.IdentifierType))
		//fmt.Println("IdentifierType: ", *(**C.char)(unsafe.Pointer(addrIdentifierType + uintptr(8 * i))))

		// NamespaceIndex
		addrNamespaceIndex := uintptr(unsafe.Pointer(uaConfig.NodeIds.NamespaceIndex))
		*(*C.int)(unsafe.Pointer(addrNamespaceIndex + uintptr(i *4))) = (C.int)(nodeId.NamespaceIndex)

		i++
	}


	// init security
	// copy Password
	uaConfig.Security.Password = (*C.char)(C.galloc((C.int)(len(opcUaConfig.Security.Password))))
	C.strcpy(uaConfig.Security.Password, C.CString(opcUaConfig.Security.Password))
	defer C.free(unsafe.Pointer(uaConfig.Security.Password))
	// copy StoreType
	uaConfig.Security.StoreType = (*C.char)(C.galloc((C.int)((C.int)(len(opcUaConfig.Security.StoreType)))))
	C.strcpy(uaConfig.Security.StoreType, C.CString(opcUaConfig.Security.StoreType))
	defer C.free(unsafe.Pointer(uaConfig.Security.StoreType))
	// copy KeystoreFilePath
	uaConfig.Security.KeystoreFilePath = (*C.char)(C.galloc((C.int)(len(opcUaConfig.Security.KeystoreFilePath))))
	C.strcpy(uaConfig.Security.KeystoreFilePath, C.CString(opcUaConfig.Security.KeystoreFilePath))
	defer C.free(unsafe.Pointer(uaConfig.Security.KeystoreFilePath))
	// copy Alias
	uaConfig.Security.Alias = (*C.char)(C.galloc((C.int)(len(opcUaConfig.Security.Alias))))
	C.strcpy(uaConfig.Security.Alias, C.CString(opcUaConfig.Security.Alias))
	defer C.free(unsafe.Pointer(uaConfig.Security.Alias))
	// copy SecurityPolicy
	uaConfig.Security.SecurityPolicy = (*C.char)(C.galloc((C.int)(len(opcUaConfig.Security.SecurityPolicy))))
	C.strcpy(uaConfig.Security.SecurityPolicy, C.CString(opcUaConfig.Security.SecurityPolicy))
	defer C.free(unsafe.Pointer(uaConfig.Security.SecurityPolicy))


	// init channel
	// copy MaxChunkCount
	uaConfig.ChannelConfig.MaxChunkCount = (C.int)(opcUaConfig.ChannelConfig.MaxChunkCount)
	// copy MaxArrayLength
	uaConfig.ChannelConfig.MaxArrayLength = (C.int)(opcUaConfig.ChannelConfig.MaxArrayLength)
	// copy MaxMessageSize
	uaConfig.ChannelConfig.MaxMessageSize = (C.int)(opcUaConfig.ChannelConfig.MaxMessageSize)
	// copy MaxStringLength
	uaConfig.ChannelConfig.MaxStringLength = (C.int)(opcUaConfig.ChannelConfig.MaxStringLength)
	// copy MaxChunkSize
	uaConfig.ChannelConfig.MaxChunkSize = (C.int)(opcUaConfig.ChannelConfig.MaxChunkSize)


	// init connect config
	// copy ResourceUrl
	uaConfig.Config.ResourceUrl = (*C.char)(C.galloc((C.int)(len(opcUaConfig.Config.ResourceUrl))))
	C.strcpy(uaConfig.Config.ResourceUrl, C.CString(opcUaConfig.Config.ResourceUrl))
	defer C.free(unsafe.Pointer(uaConfig.Config.ResourceUrl))
	// copy UseCredenials
	uaConfig.Config.UseCredenials = (C.int)(boolToInt(opcUaConfig.Config.UseCredenials))
	// copy PollingInterval
	uaConfig.Config.PollingInterval = (C.int)(opcUaConfig.Config.PollingInterval)
	// copy ApplicationUrl
	uaConfig.Config.ApplicationUrl = (*C.char)(C.galloc((C.int)(len(opcUaConfig.Config.ApplicationUrl))))
	C.strcpy(uaConfig.Config.ApplicationUrl, C.CString(opcUaConfig.Config.ApplicationUrl))
	defer C.free(unsafe.Pointer(uaConfig.Config.ApplicationUrl))
	// copy SessionTimeOut
	uaConfig.Config.SessionTimeOut = (C.int)(opcUaConfig.Config.SessionTimeOut)
	// copy ProcessingMode
	uaConfig.Config.ProcessingMode = (*C.char)(C.galloc((C.int)(len(opcUaConfig.Config.ProcessingMode))))
	C.strcpy(uaConfig.Config.ProcessingMode, C.CString(opcUaConfig.Config.ProcessingMode))
	defer C.free(unsafe.Pointer(uaConfig.Config.ProcessingMode))
	// copy RequestTimeOut
	uaConfig.Config.RequestTimeOut = (C.int)(opcUaConfig.Config.RequestTimeOut)
	// copy ReconnectTime
	uaConfig.Config.ReconnectTime = (C.int)(opcUaConfig.Config.ReconnectTime)

	// init credenials
	// copy useName
	uaConfig.Credenials.userName = (*C.char)(C.galloc((C.int)(len(opcUaConfig.Credenials.userName))))
	C.strcpy(uaConfig.Credenials.userName, C.CString(opcUaConfig.Credenials.userName))
	defer C.free(unsafe.Pointer(uaConfig.Credenials.userName))
	//fmt.Println("Credenials userName: ", *uaConfig.Credenials.userName)
	// copy passWord
	uaConfig.Credenials.passWord = (*C.char)(C.galloc((C.int)(len(opcUaConfig.Credenials.passWord))))
	C.strcpy(uaConfig.Credenials.passWord, C.CString(opcUaConfig.Credenials.passWord))
	defer C.free(unsafe.Pointer(uaConfig.Credenials.passWord))
	//fmt.Println("Credenials passWord: ", *uaConfig.Credenials.passWord)
	//
	//var pDetectInfo C.UA_Read_Retval
	//pDetectInfo.data = unsafe.Pointer((*C.void)(C.malloc(dataCacheSize)))
	//if pDetectInfo.data == nil {
	//	fmt.Println("go malloc data failed.")
	//}
	//defer C.free(unsafe.Pointer(pDetectInfo.data))
	//
	//fmt.Println("pDetectInfo: ", pDetectInfo)
	////C.Subscribe(&pDetectInfo)
	C.Subscribe(&uaConfig, nodeIdLength)

	fmt.Println("end....")
}

func NewOpcSubscribe(filename string)(*OpcSubscribe, error){

	return &OpcSubscribe{
		filename:configFileName,
	}, nil
}