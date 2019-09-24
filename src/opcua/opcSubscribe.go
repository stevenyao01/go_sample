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

typedef struct {
    char *typeName;
    char *key;
	int arrayLength;
	void *data;
} UA_Read_Retval;

void Subscribe();

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

extern void OpcCallback(UA_Read_Retval pRet);

UA_Boolean running1 = true;

static void stopHandler(int sign) {
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "Received Ctrl-C");
    running1 = 0;
}

static void
handler_currentTimeChanged(UA_Client *client, UA_UInt32 subId, void *subContext,
                           UA_UInt32 monId, void *monContext, UA_DataValue *value) {
    //char str[] = "hello yaohaping...";
    int sizeType = 0;
    int sizeData = 0;
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "currentTime has changed!");
    UA_Read_Retval retval;

    // typeName
    sizeType = strlen(value->value.type->typeName);
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "type size: %d", sizeType);
    char str[32];
    memcpy(str, value->value.type->typeName, sizeType);
    retval.typeName = str;
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "typeName %s", retval.typeName);


    // arrayLength
    retval.arrayLength = value->value.arrayLength;
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "arrayLength: %d", retval.arrayLength);

    // alloc mem for data
    sizeData = sizeof(bool) * value->value.arrayLength;
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "data size: %d", sizeData);
    retval.data = (void *)malloc(sizeData);
    memset(retval.data, 0 , sizeData);
    memcpy(retval.data, value->value.data, sizeData);

    // call go func
    OpcCallback(retval);

    // release mem
    //free(retval->typeName);
    free(retval.data);

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
			UA_MonitoredItemCreateRequest_default(UA_NODEID_STRING(2, "Demo.Dynamic.Arrays.Boolean"));

			UA_MonitoredItemCreateResult monResponse =
			UA_Client_MonitoredItems_createDataChange(client, response.subscriptionId, UA_TIMESTAMPSTORETURN_BOTH,
				monRequest, NULL, handler_currentTimeChanged, NULL);
			if(monResponse.statusCode == UA_STATUSCODE_GOOD)
				UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND,"Monitoring UA_NS0ID_SERVER_SERVERSTATUS_CURRENTTIME', id %u",
					monResponse.monitoredItemId);
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
Subscribe() {
	signal(SIGINT, stopHandler);

	UA_Client *client = UA_Client_new();
	UA_ClientConfig *cc = UA_Client_getConfig(client);
	UA_ClientConfig_setDefault(cc);

	cc->stateCallback = stateCallback;
	cc->subscriptionInactivityCallback = subscriptionInactivityCallback;

	while(running1) {
		UA_StatusCode retval = UA_Client_connect(client, "opc.tcp://10.111.66.220:48030");
		if(retval != UA_STATUSCODE_GOOD) {
			UA_LOG_ERROR(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "Not connected. Retrying to connect in 1 second");
			UA_sleep_ms(1000);
			continue;
		}

		UA_Client_run_iterate(client, 1000);
	};

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

func (s *OpcSubscribe) SubscribeRead()(){
	fmt.Println("=============================== cgo opcua subscribe ===========================")

	var pDetectInfo C.UA_Read_Retval
	pDetectInfo.data = unsafe.Pointer((*C.void)(C.malloc(dataCacheSize)))
	if pDetectInfo.data == nil {
		fmt.Println("go malloc data failed.")
	}
	defer C.free(unsafe.Pointer(pDetectInfo.data))

	fmt.Println("pDetectInfo: ", pDetectInfo)
	//C.Subscribe(&pDetectInfo)
	C.Subscribe()

	fmt.Println("end....")
}

func NewOpcSubscribe(filename string)(*OpcSubscribe, error){

	return &OpcSubscribe{
		filename:configFileName,
	}, nil
}