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

void Browser(UA_Read_Retval *pRet);

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

UA_Boolean running2 = true;

static void stopHandler(int sign) {
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_CLIENT, "Received Ctrl-C");
    running2 = 0;
}

void
browserObjects(UA_Client *client){
    printf("Browsing nodes in objects folder:\n");
    UA_BrowseRequest bReq;
    UA_BrowseRequest_init(&bReq);
    bReq.requestedMaxReferencesPerNode = 0;
    bReq.nodesToBrowse = UA_BrowseDescription_new();
    bReq.nodesToBrowseSize = 1;
    bReq.nodesToBrowse[0].nodeId = UA_NODEID_NUMERIC(2, 2250);
    //bReq.nodesToBrowse[0].nodeId = UA_NODEID_STRING(2, "Demo.Static.Arrays");
	bReq.nodesToBrowse[0].resultMask = UA_BROWSERESULTMASK_ALL;
	UA_BrowseResponse bResp = UA_Client_Service_browse(client, bReq);
	printf("%-9s %-16s %-16s %-16s\n", "NAMESPACE", "NODEID", "BROWSE NAME", "DISPLAY NAME");
	for(size_t i = 0; i < bResp.resultsSize; ++i) {
		for(size_t j = 0; j < bResp.results[i].referencesSize; ++j) {
			UA_ReferenceDescription *ref = &(bResp.results[i].references[j]);
			//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua browser.");
			if(ref->nodeId.nodeId.identifierType == UA_NODEIDTYPE_NUMERIC) {
				printf("%-9d %-16d %-16.*s %-16.*s\n", ref->nodeId.nodeId.namespaceIndex,
				ref->nodeId.nodeId.identifier.numeric, (int)ref->browseName.name.length,
				ref->browseName.name.data, (int)ref->displayName.text.length,
				ref->displayName.text.data);
			} else if(ref->nodeId.nodeId.identifierType == UA_NODEIDTYPE_STRING) {
				printf("%-9d %-16.*s %-16.*s %-16.*s\n", ref->nodeId.nodeId.namespaceIndex,
				(int)ref->nodeId.nodeId.identifier.string.length,
				ref->nodeId.nodeId.identifier.string.data,
				(int)ref->browseName.name.length, ref->browseName.name.data,
				(int)ref->displayName.text.length, ref->displayName.text.data);
			}
		}
	}
	UA_BrowseRequest_clear(&bReq);
	UA_BrowseResponse_clear(&bResp);
}

void
Browser(UA_Read_Retval *pRet) {
    signal(SIGINT, stopHandler);

	UA_Client *client = UA_Client_new();
	UA_ClientConfig *cc = UA_Client_getConfig(client);
	UA_ClientConfig_setDefault(cc);
	cc->timeout = 1000;

	//UA_Variant *value;
	//UA_Variant_init(value);
	UA_Variant *value = UA_Variant_new();
	while(running2) {
		UA_StatusCode retval = UA_Client_connect(client, "opc.tcp://10.111.66.220:48030");
		if(retval != UA_STATUSCODE_GOOD) {
			UA_LOG_ERROR(UA_Log_Stdout, UA_LOGCATEGORY_CLIENT, "Not connected. Retrying to connect in 1 second");

			UA_sleep_ms(1000);
			continue;
		}

		browserObjects(client);

		UA_Variant_clear(value);
		//UA_sleep_ms(1000);
		UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua program loop.");
		break;
	};

	UA_Variant_clear(value);
	UA_Client_delete(client);

	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua program exit success.");
	return;
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type OpcBrowser struct {
	filename string
}

func (b *OpcBrowser) BrowserRead()(){
	fmt.Println("=============================== cgo opcua browser ===========================")
	var pDetectInfo C.UA_Read_Retval
	pDetectInfo.data = unsafe.Pointer((*C.void)(C.malloc(dataCacheSize)))
	if pDetectInfo.data == nil {
		fmt.Println("go malloc data failed.")
	}
	defer C.free(unsafe.Pointer(pDetectInfo.data))

	fmt.Println("pDetectInfo: ", pDetectInfo)
	C.Browser(&pDetectInfo)

	typeName := C.GoString(pDetectInfo.typeName)
	arrayLength := pDetectInfo.arrayLength

	fmt.Println("typeName: ", typeName)
	fmt.Println("arrayLength: ", arrayLength)

	// loop for arrayLength to convert value.
	for i := 0; i < int(arrayLength); i++ {
		fmt.Println("data: ", *(*bool)(unsafe.Pointer(uintptr(pDetectInfo.data) + uintptr(i))))
	}

	fmt.Println("end....")
}


func NewOpcBrowser(filename string)(*OpcBrowser, error){

	return &OpcBrowser{
		filename:configFileName,
	}, nil
}