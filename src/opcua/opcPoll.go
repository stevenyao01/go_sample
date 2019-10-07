package main
/**
 * @Package Name: yao
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-9-2 下午3:56
 * @Description:
 */

import "C"
import (
	"fmt"
	"unsafe"
)

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
#include "opcua.h"

void Polling(Ua_Read_Retval pRet, Opc_Ua_Config *Ua_Config, int len);

//#define UA_ENABLE_ENCRYPTION true
//// define for ua config
//#define NODEIDS_IDENTIFIER_LENGTH 512
//#define NODEIDS_FIELD_LENGTH 32
//#define NODEIDS_IDENTIFIERTYPE_LENGTH 32
//#define NODEIDS_NAMESPACEINDEX_LENGTH 32
//// define for ua return value
//#define NODEIDS_RET_TYPENAME_LENGTH 32
//#define NODEIDS_RET_KEY_LENGTH 32
//#define NODEIDS_RET_DATA_LENGTH 512
//#define NODEIDS_RET_ARRAY_LENGTH 4

UA_Boolean running = true;
char *opcNumeric = "Numeric";
char *opcString = "String";
char *opcUUID = "UUID";
char *opcOpaque = "Opaque";

static void stopHandler(int sign) {
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_CLIENT, "Received Ctrl-C");
    running = 0;
}

#ifdef UA_ENABLE_ENCRYPTION
static UA_ByteString loadFile(const char *const path) {
	UA_ByteString fileContents = UA_BYTESTRING_NULL;
	if(path == NULL)
		return fileContents;

	FILE *fp = fopen(path, "rb");
	if(!fp) {
		errno = 0;
		return fileContents;
	}

	fseek(fp, 0, SEEK_END);
	fileContents.length = (size_t) ftell(fp);
	fileContents.data = (UA_Byte *) UA_malloc(fileContents.length * sizeof(UA_Byte));
	if(fileContents.data) {
		fseek(fp, 0, SEEK_SET);
		size_t read = fread(fileContents.data, sizeof(UA_Byte), fileContents.length, fp);
		if(read != fileContents.length)
			UA_ByteString_clear(&fileContents);
	} else {
		fileContents.length = 0;
	}

	fclose(fp);
	return fileContents;
}
#endif

void
Polling(Ua_Read_Retval pRet, Opc_Ua_Config *Ua_Config, int len) {
    signal(SIGINT, stopHandler);

	//// print param
	//for (int i = 0; i < len; i++){
	//	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua identifier: %s", Ua_Config->NodeIds->Identifier[i]);
	//	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua field: %s", Ua_Config->NodeIds->Field[i]);
	//	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua identifiertype: %s", Ua_Config->NodeIds->IdentifierType[i]);
	//	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua namespaceindex: %d", Ua_Config->NodeIds->NamespaceIndex[i]);
	//}
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua password: %s", Ua_Config->Security->Password);
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua resource url: %s", Ua_Config->Config->ResourceUrl);
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua maxMessageSize: %d", Ua_Config->ChannelConfig->MaxMessageSize);
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua maxChunkCount: %d", Ua_Config->ChannelConfig->MaxChunkCount);
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua maxChunkSize: %d", Ua_Config->ChannelConfig->MaxChunkSize);
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua SecurityPolicy: %s", Ua_Config->Security->SecurityPolicy);
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua node len: %d", len);

    UA_String securityPolicyUri = UA_STRING_NULL;
    UA_MessageSecurityMode securityMode = UA_MESSAGESECURITYMODE_INVALID;
#ifdef UA_ENABLE_ENCRYPTION
	char *certfile = NULL;
	char *keyfile = NULL;
#endif
	UA_Client *client = UA_Client_new();
	UA_ClientConfig *cc = UA_Client_getConfig(client);
#ifdef UA_ENABLE_ENCRYPTION
    if(certfile) { // todo get certificate and privateKey from pfx file
        UA_ByteString certificate = loadFile(certfile);
        UA_ByteString privateKey  = loadFile(keyfile);
        //UA_ClientConfig_setDefaultEncryption(cc, certificate, privateKey, NULL, 0, NULL, 0);
        UA_ByteString_deleteMembers(&certificate);
        UA_ByteString_deleteMembers(&privateKey);
    } else {
        UA_ClientConfig_setDefault(cc);
    }
#else
    UA_ClientConfig_setDefault(cc);
#endif


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

	while(running) {
		UA_StatusCode retval;
		if (Ua_Config->Config->UseCredenials == 1) {
			retval = UA_Client_connect_username(client, Ua_Config->Config->ResourceUrl, "username", "password");
		} else {
			retval = UA_Client_connect(client, Ua_Config->Config->ResourceUrl);
		}
		if(retval != UA_STATUSCODE_GOOD) {
			UA_LOG_ERROR(UA_Log_Stdout, UA_LOGCATEGORY_CLIENT, "Not connected. Retrying to connect in 1 second");

			UA_sleep_ms(Ua_Config->Config->ReconnectTime);
			continue;
		}

		for(int i = 0; i < len; i++){
			if (strcmp(*Ua_Config->NodeIds->IdentifierType, opcString) == 0){
				UA_Variant *valueStr = UA_Variant_new();
				UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "equal string!!!!!!!!!!!!");
				retval = UA_Client_readValueAttribute(client, UA_NODEID_STRING(Ua_Config->NodeIds->NamespaceIndex[i], Ua_Config->NodeIds->Identifier[i]), valueStr);
				if(retval == UA_STATUSCODE_BADCONNECTIONCLOSED) {
					UA_LOG_ERROR(UA_Log_Stdout, UA_LOGCATEGORY_CLIENT, "Connection was closed. Reconnecting ...");
					UA_sleep_ms(Ua_Config->Config->ReconnectTime);
					continue;
				}
				if(retval == UA_STATUSCODE_GOOD) {
					memcpy(pRet.Usn->typeName[i], valueStr->type->typeName, sizeof(char) * NODEIDS_RET_TYPENAME_LENGTH);
					memcpy(pRet.Usn->key[i], Ua_Config->NodeIds->Field[i], sizeof(char) * NODEIDS_RET_KEY_LENGTH);
					pRet.Usn->arrayLength[i] = (int)(valueStr->arrayLength);
					//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "%p pRet.Usn->arrayLength[%d] = %d", &pRet.Usn->arrayLength[i], i, pRet.Usn->arrayLength[i]);
					memcpy(pRet.Usn->data[i], valueStr->data, sizeof(char) * NODEIDS_RET_DATA_LENGTH);
				}
				UA_Variant_clear(valueStr);
			}
			if (strcmp(*Ua_Config->NodeIds->IdentifierType, opcNumeric) == 0){
				UA_Variant value;
				UA_Variant_init(&value);
				UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "equal Numeric!!!!!!!!!!!!");
				const UA_NodeId nodeId = UA_NODEID_NUMERIC(0, UA_NS0ID_SERVER_SERVERSTATUS_CURRENTTIME);
        		retval = UA_Client_readValueAttribute(client, nodeId, &value);
			//if(retval == UA_STATUSCODE_GOOD &&
			//	UA_Variant_hasScalarType(&value, &UA_TYPES[UA_TYPES_DATETIME])) {
			//	UA_DateTime raw_date = *(UA_DateTime *) value->data;
			//	UA_DateTimeStruct dts = UA_DateTime_toStruct(raw_date);
			//	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND,
			//	"date is: %02u-%02u-%04u %02u:%02u:%02u.%03u",
			//	dts.day, dts.month, dts.year, dts.hour, dts.min, dts.sec, dts.milliSec);
			//}
				UA_Variant_clear(&value);
			}
			if (strcmp(*Ua_Config->NodeIds->IdentifierType, opcUUID) == 0){
				UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "equal UUID!!!!!!!!!!!!");
			}
			if (strcmp(*Ua_Config->NodeIds->IdentifierType, opcOpaque) == 0){
				UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "equal Opaque!!!!!!!!!!!!");
			}
		}

		//UA_sleep_ms(1000);
		//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua program loop.");
		break;
	};

	//UA_Variant_clear(value);
	UA_Client_delete(client);

	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua program exit success.");
	return;
}
*/
import "C"


type OpcPoll struct {
	filename string
}


func (p *OpcPoll) PollRead(opcUaConfig OpcUaConfig)(){
	fmt.Println("=============================== cgo opcua polling ===========================")

	nodeIdLength := (C.int)(len(opcUaConfig.NodeIds))
	urr := C.newOpcUaRetval(nodeIdLength)
	defer C.deleteOpcUaRetval(urr, (C.int)(len(opcUaConfig.NodeIds)))

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

	// call cgo here
	C.Polling(urr, &uaConfig, nodeIdLength)

	addrTypeName := uintptr(unsafe.Pointer(urr.Usn.typeName))
	addrKey := uintptr(unsafe.Pointer(urr.Usn.key))
	addrData := uintptr(unsafe.Pointer(urr.Usn.data))
	addrArrayLength := uintptr(unsafe.Pointer(urr.Usn.arrayLength))
	for i := 0; i < int(nodeIdLength); i++ {
		fmt.Println("arrayLength: ", *(*int32)(unsafe.Pointer(addrArrayLength + uintptr(i * 4))))
		fmt.Println("addrTypeName: ", C.GoString(*(**C.char)(unsafe.Pointer(addrTypeName + uintptr(i * 8)))))
		fmt.Println("addrKey: ", C.GoString(*(**C.char)(unsafe.Pointer(addrKey + uintptr(i * 8)))))
		//fmt.Println("addrData: ", *(**C.bool)(unsafe.Pointer(addrData + uintptr(i * 8))))
		pDataAddr := *(**C.bool)(unsafe.Pointer(addrData + uintptr(i * 8)))
		for j := 0; j < int(*urr.Usn.arrayLength); j++ {
			fmt.Println("data : ", *(*bool)(unsafe.Pointer(uintptr(unsafe.Pointer(pDataAddr)) + uintptr(j))))
		}
	}

	fmt.Println("end....")
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func NewOpcPoll(fileName string)(*OpcPoll, error) {

	return &OpcPoll{
		filename: fileName,
	}, nil
}
