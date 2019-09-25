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

typedef struct {
    const char *typeName;
    const char *key;
	int   arrayLength;
	void  *data;
} UA_Read_Retval;

typedef struct{
	char **Identifier;
	char **Field;
	char **IdentifierType;
	int  *NamespaceIndex;
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
	int  UseCredenials;
	int  PollingInterval;
	char *ApplicationUrl;
	int  SessionTimeOut;
	char *ProcessingMode;
	int  RequestTimeOut;
	int  ReconnectTime;
} Ua_Connect_Config;

typedef struct {
	char *userName;
	char *passWord;
} Ua_Credenials;

typedef struct {
	Ua_Node_Id 			*NodeIds;
	Ua_Security 		*Security;
	Ua_Channel_Config 	*ChannelConfig;
	Ua_Connect_Config   *Config;
	Ua_Credenials		*Credenials;
} Opc_Ua_Config;


void Polling(UA_Read_Retval *pRet, Opc_Ua_Config *Ua_Config);

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

#define UA_ENABLE_ENCRYPTION true
#define NODEIDS_IDENTIFIER_LENGTH 512
#define NODEIDS_FIELD_LENGTH 32
#define NODEIDS_IDENTIFIERTYPE_LENGTH 32
#define NODEIDS_NAMESPACEINDEX_LENGTH 32

UA_Boolean running = true;
char *opcNumeric = "Numeric";
char *opcString = "String";
char *opcUUID = "UUID";
char *opcOpaque = "Opaque";

static void stopHandler(int sign) {
    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_CLIENT, "Received Ctrl-C");
    running = 0;
}

Opc_Ua_Config newOpcUaConfig(int len){

UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "hello world!");
	Opc_Ua_Config ouc;
	ouc.NodeIds = malloc(sizeof(Ua_Node_Id) * len);
	//memset(ouc.NodeIds, 0, sizeof(Ua_Node_Id) * len);

	ouc.Security = malloc(sizeof(Ua_Security) * len);
	//memset(ouc.Security, 0, sizeof(Ua_Security) * len);

	ouc.ChannelConfig = malloc(sizeof(Ua_Channel_Config) * len);
	//memset(ouc.ChannelConfig, 0, sizeof(Ua_Channel_Config) * len);

	ouc.Config = malloc(sizeof(Ua_Connect_Config) * len);
	//memset(ouc.Config, 0, sizeof(Ua_Connect_Config) * len);

	ouc.Credenials = malloc(sizeof(Ua_Credenials) * len);
	//memset(ouc.Credenials, 0, sizeof(Ua_Credenials) * len);

	for (int i = 0; i < len; i++) {
		ouc.NodeIds->Identifier[i] = malloc(sizeof(char) * NODEIDS_IDENTIFIER_LENGTH);
		//memset(ouc.NodeIds->Identifier[i], 0, sizeof(char) * NODEIDS_IDENTIFIER_LENGTH);
		UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "identifier: %p", ouc.NodeIds->Identifier[i]);

		ouc.NodeIds->Field[i] = malloc(sizeof(char) * NODEIDS_FIELD_LENGTH);
		//memset(ouc.NodeIds->Field[i], 0, sizeof(char) * NODEIDS_FIELD_LENGTH);
		UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "field: %p", ouc.NodeIds->Field[i]);

		ouc.NodeIds->IdentifierType[i] = malloc(sizeof(char) * NODEIDS_IDENTIFIERTYPE_LENGTH);
		//memset(ouc.NodeIds->IdentifierType[i], 0, sizeof(char) * NODEIDS_IDENTIFIERTYPE_LENGTH);
	}
	ouc.NodeIds->NamespaceIndex = malloc(sizeof(int) * len);
	//memset(ouc.NodeIds->NamespaceIndex[i], 0, sizeof(int) * len);

	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "hello world end!");
	return ouc;
}

void deleteOpcUaConfig(Opc_Ua_Config ouc, int len1){
	int len = sizeof(ouc.NodeIds->Identifier);
	for (int i = 0; i < len; i++) {
		free(ouc.NodeIds->Identifier[i]);
		free(ouc.NodeIds->Field[i]);
		free(ouc.NodeIds->IdentifierType[i]);
	}
	free(ouc.NodeIds->NamespaceIndex);

	free(ouc.NodeIds);
	free(ouc.Security);
	free(ouc.ChannelConfig);
	free(ouc.Config);
	free(ouc.Credenials);


	return;
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

void copyStr(char *strDst, char *strSrc) {
	strcpy(strDst, strSrc);
	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "copyStr strDst: %s", strDst);
}

void* galloc(int length){
	char *pmem;
	pmem = malloc(length);
	memset(pmem, 0, length);
	return pmem;
}

void
Polling(UA_Read_Retval *pRet, Opc_Ua_Config *Ua_Config) {
    signal(SIGINT, stopHandler);
//////////////////
	for (int i = 0; i < 2; i++){
		UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua identifier: %s", Ua_Config->NodeIds->Identifier[i]);
		UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua field: %s", Ua_Config->NodeIds->Field[i]);
		UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua identifiertype: %s", Ua_Config->NodeIds->IdentifierType[i]);
		UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua namespaceindex: %d", Ua_Config->NodeIds->NamespaceIndex[i]);
	}
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua NamespaceIndex: %d", *Ua_Config->NodeIds->NamespaceIndex);
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua identifier: %s", Ua_Config->NodeIds->Identifier);
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua identifier: %p", Ua_Config->NodeIds->Identifier);
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua password: %s", Ua_Config->Security->Password);
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua resource url: %s", Ua_Config->Config->ResourceUrl);
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua identifierType: %s", Ua_Config->NodeIds->IdentifierType);
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua maxMessageSize: %d", Ua_Config->ChannelConfig->MaxMessageSize);
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua maxChunkCount: %d", Ua_Config->ChannelConfig->MaxChunkCount);
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua maxChunkSize: %d", Ua_Config->ChannelConfig->MaxChunkSize);
	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua SecurityPolicy: %s", Ua_Config->Security->SecurityPolicy);
////////////
return;
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
		//if (strcmp(Ua_Config->NodeIds->IdentifierType, opcString) == 0){
		//	UA_Variant *valueStr = UA_Variant_new();
		//	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "equal string!!!!!!!!!!!!");
		//	retval = UA_Client_readValueAttribute(client, UA_NODEID_STRING(*Ua_Config->NodeIds->NamespaceIndex, Ua_Config->NodeIds->Identifier), valueStr);
		//	if(retval == UA_STATUSCODE_BADCONNECTIONCLOSED) {
		//		UA_LOG_ERROR(UA_Log_Stdout, UA_LOGCATEGORY_CLIENT, "Connection was closed. Reconnecting ...");
		//		UA_sleep_ms(Ua_Config->Config->ReconnectTime);
		//		continue;
		//	}
		//	if(retval == UA_STATUSCODE_GOOD) {
		//		pRet->typeName = valueStr->type->typeName;
		//		pRet->key = Ua_Config->NodeIds->Field;
		//		pRet->arrayLength = valueStr->arrayLength;
		//		memcpy(pRet->data, valueStr->data, 8);
		//	}
		//	UA_Variant_clear(valueStr);
		//}
		//if (strcmp(Ua_Config->NodeIds->IdentifierType, opcNumeric) == 0){
		//	UA_Variant value;
		//	UA_Variant_init(&value);
		//	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "equal Numeric!!!!!!!!!!!!");
		//	const UA_NodeId nodeId = UA_NODEID_NUMERIC(0, UA_NS0ID_SERVER_SERVERSTATUS_CURRENTTIME);
        	//retval = UA_Client_readValueAttribute(client, nodeId, &value);
		//	//if(retval == UA_STATUSCODE_GOOD &&
		//	//	UA_Variant_hasScalarType(&value, &UA_TYPES[UA_TYPES_DATETIME])) {
		//	//	UA_DateTime raw_date = *(UA_DateTime *) value->data;
		//	//	UA_DateTimeStruct dts = UA_DateTime_toStruct(raw_date);
		//	//	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND,
		//	//	"date is: %02u-%02u-%04u %02u:%02u:%02u.%03u",
		//	//	dts.day, dts.month, dts.year, dts.hour, dts.min, dts.sec, dts.milliSec);
		//	//}
		//	UA_Variant_clear(&value);
		//}
		//if (strcmp(Ua_Config->NodeIds->IdentifierType, opcUUID) == 0){
		//	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "equal UUID!!!!!!!!!!!!");
		//}
		//if (strcmp(Ua_Config->NodeIds->IdentifierType, opcOpaque) == 0){
		//	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "equal Opaque!!!!!!!!!!!!");
		//}




		//UA_sleep_ms(1000);
		UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "opcua program loop.");
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
	var pDetectInfo C.UA_Read_Retval
	//pDetectInfo.data = unsafe.Pointer((*C.void)(C.malloc(dataCacheSize)))
	//if pDetectInfo.data == nil {
	//	fmt.Println("go malloc data failed.")
	//}
	pDetectInfo.data = unsafe.Pointer(C.galloc(dataCacheSize))
	defer C.free(unsafe.Pointer(pDetectInfo.data))
////////////////////
	// test trans opc ua config
	uaConfig := C.newOpcUaConfig((C.int)(len(opcUaConfig.NodeIds)))
	//defer C.deleteOpcUaConfig(uaConfig, (C.int)(len(opcUaConfig.NodeIds)))

	fmt.Println("lenovo length: ", len(opcUaConfig.NodeIds))

	//uaConfig.NodeIds = (*C.Ua_Node_Id)(C.malloc((C.ulong)(16 * len(opcUaConfig.NodeIds))))
	i := 0
	for _, nodeId := range opcUaConfig.NodeIds {
		//copy Identifier
		addrIdentifier := unsafe.Pointer(uaConfig.NodeIds.Identifier)
		C.strcpy(*(**C.char)(unsafe.Pointer(uintptr(addrIdentifier) + uintptr(i * 8))), C.CString(nodeId.Identifier))
		fmt.Println("identifier: ", *(**C.char)(unsafe.Pointer(uintptr(addrIdentifier) + uintptr(i * 8))))
		//// copy Field
		addrField := unsafe.Pointer(uaConfig.NodeIds.Field)
		fmt.Println("field: ", nodeId.Field, "identifier: ", nodeId.Identifier)
		C.strcpy(*(**C.char)(unsafe.Pointer(uintptr(addrField) + uintptr(i * 8))), C.CString(nodeId.Field))
		fmt.Println("field: ", *(**C.char)(unsafe.Pointer(uintptr(addrField) + uintptr(i * 8))))
		//addrField := C.galloc((C.int)(len(nodeId.Identifier)))
		//uaConfig.NodeIds.Field = &(*(*C.char)(addrField))
		//copyStr(uaConfig.NodeIds.Field, nodeId.Field)
		//defer C.free(unsafe.Pointer(uaConfig.NodeIds.Field))
		//// copy IdentifierType
		addrIdentifierType := unsafe.Pointer(uaConfig.NodeIds.IdentifierType)
		C.strcpy(*(**C.char)(unsafe.Pointer(uintptr(addrIdentifierType) + uintptr(i * 8))), C.CString(nodeId.IdentifierType))
		fmt.Println("IdentifierType: ", *(**C.char)(unsafe.Pointer(uintptr(addrIdentifierType) + uintptr(i * 8))))
		//addrIdentifierType := C.galloc((C.int)(len(nodeId.Identifier)))
		//uaConfig.NodeIds.IdentifierType = &(*(*C.char)(addrIdentifierType))
		//copyStr(uaConfig.NodeIds.IdentifierType, nodeId.IdentifierType)
		//defer C.free(unsafe.Pointer(uaConfig.NodeIds.IdentifierType))
		//// copy NamespaceIndexs
		addrNamespaceIndex := unsafe.Pointer(uaConfig.NodeIds.NamespaceIndex)
		*(*int)(unsafe.Pointer(uintptr(addrNamespaceIndex) + uintptr(i * 4))) = int(nodeId.NamespaceIndex)
		//*(*C.int)(unsafe.Pointer(uintptr(addrNamespaceIndex) + uintptr(i * 8))) = (C.int)(nodeId.NamespaceIndex)
		//*(**C.int)(unsafe.Pointer(uintptr(addrNamespaceIndex) + uintptr(i * 8))) = *(C.int)(nodeId.NamespaceIndex)
		//C.strcpy(*(**C.char)(unsafe.Pointer(uintptr(addrNamespaceIndex) + uintptr(i * 8))), C.CString(nodeId.NamespaceIndex))
		//addrNamespaceIndex := C.galloc((C.int)(len(nodeId.Identifier)))
		//uaConfig.NodeIds.NamespaceIndex = &(*(*C.int)(addrNamespaceIndex))
		//*(*int)(unsafe.Pointer(uintptr(addrNamespaceIndex) + uintptr(i*4))) = int(nodeId.NamespaceIndex)
		i++
	}


	//// init security
	//uaConfig.Security = (*C.Ua_Security)(C.malloc(28))
	//// copy Password
	//uaConfig.Security.Password = C.galloc(len(opcUaConfig.Security.Password))
	//copyStr(uaConfig.Security.Password, opcUaConfig.Security.Password)
	//defer C.free(unsafe.Pointer(uaConfig.Security.Password))
	//// copy StoreType
	//uaConfig.Security.StoreType = C.galloc(len(opcUaConfig.Security.StoreType))
	//copyStr(uaConfig.Security.StoreType, opcUaConfig.Security.StoreType)
	//defer C.free(unsafe.Pointer(uaConfig.Security.StoreType))
	//// copy KeystoreFilePath
	//uaConfig.Security.KeystoreFilePath = C.galloc(len(opcUaConfig.Security.KeystoreFilePath))
	//copyStr(uaConfig.Security.KeystoreFilePath, opcUaConfig.Security.KeystoreFilePath)
	//defer C.free(unsafe.Pointer(uaConfig.Security.KeystoreFilePath))
	//// copy Alias
	//uaConfig.Security.Alias = C.galloc(len(opcUaConfig.Security.Alias))
	//copyStr(uaConfig.Security.Alias, opcUaConfig.Security.Alias)
	//defer C.free(unsafe.Pointer(uaConfig.Security.Alias))
	//// copy SecurityPolicy
	//uaConfig.Security.SecurityPolicy = C.galloc(len(opcUaConfig.Security.SecurityPolicy))
	//copyStr(uaConfig.Security.SecurityPolicy, opcUaConfig.Security.SecurityPolicy)
	//defer C.free(unsafe.Pointer(uaConfig.Security.SecurityPolicy))
	//
	//
	//// init channel
	//uaConfig.ChannelConfig = (*C.Ua_Channel_Config)(C.malloc(20))
	//// copy MaxChunkCount
	//uaConfig.ChannelConfig.MaxChunkCount = (C.int)(opcUaConfig.ChannelConfig.MaxChunkCount)
	//// copy MaxArrayLength
	//uaConfig.ChannelConfig.MaxArrayLength = (C.int)(opcUaConfig.ChannelConfig.MaxArrayLength)
	//// copy MaxMessageSize
	//uaConfig.ChannelConfig.MaxMessageSize = (C.int)(opcUaConfig.ChannelConfig.MaxMessageSize)
	//// copy MaxStringLength
	//uaConfig.ChannelConfig.MaxStringLength = (C.int)(opcUaConfig.ChannelConfig.MaxStringLength)
	//// copy MaxChunkSize
	//uaConfig.ChannelConfig.MaxChunkSize = (C.int)(opcUaConfig.ChannelConfig.MaxChunkSize)
	//
	//
	//// init connect config
	//uaConfig.Config = (*C.Ua_Connect_Config)(C.malloc(44))
	//// copy ResourceUrl
	//uaConfig.Config.ResourceUrl = C.galloc(len(opcUaConfig.Config.ResourceUrl))
	//copyStr(uaConfig.Config.ResourceUrl, opcUaConfig.Config.ResourceUrl)
	//defer C.free(unsafe.Pointer(uaConfig.Config.ResourceUrl))
	//// copy UseCredenials
	//uaConfig.Config.UseCredenials = (C.int)(boolToInt(opcUaConfig.Config.UseCredenials))
	//// copy PollingInterval
	//uaConfig.Config.PollingInterval = (C.int)(opcUaConfig.Config.PollingInterval)
	//// copy ApplicationUrl
	//uaConfig.Config.ApplicationUrl = C.galloc(len(opcUaConfig.Config.ApplicationUrl))
	//copyStr(uaConfig.Config.ApplicationUrl, opcUaConfig.Config.ApplicationUrl)
	//defer C.free(unsafe.Pointer(uaConfig.Config.ApplicationUrl))
	//// copy SessionTimeOut
	//uaConfig.Config.SessionTimeOut = (C.int)(opcUaConfig.Config.SessionTimeOut)
	//// copy ProcessingMode
	//uaConfig.Config.ProcessingMode = C.galloc(len(opcUaConfig.Config.ProcessingMode))
	//copyStr(uaConfig.Config.ProcessingMode, opcUaConfig.Config.ProcessingMode)
	//defer C.free(unsafe.Pointer(uaConfig.Config.ProcessingMode))
	//// copy RequestTimeOut
	//uaConfig.Config.RequestTimeOut = (C.int)(opcUaConfig.Config.RequestTimeOut)
	//// copy ReconnectTime
	//uaConfig.Config.ReconnectTime = (C.int)(opcUaConfig.Config.ReconnectTime)
	//
	//// init credenials
	////uaConfig.Credenials = (*C.Ua_Connect_Config)(C.malloc(44))
	//// copy useName
	////uaConfig.Config.ProcessingMode = galloc(len(opcUaConfig.Config.ProcessingMode))
	////copyStr(uaConfig.Config.ProcessingMode, opcUaConfig.Config.ProcessingMode)
	////defer C.free(unsafe.Pointer(uaConfig.Config.ProcessingMode))
	////
	////// copy passWord
	////uaConfig.Config.ProcessingMode = galloc(len(opcUaConfig.Config.ProcessingMode))
	////copyStr(uaConfig.Config.ProcessingMode, opcUaConfig.Config.ProcessingMode)
	////defer C.free(unsafe.Pointer(uaConfig.Config.ProcessingMode))

	// call cgo here
	C.Polling(&pDetectInfo, &uaConfig)

	// get cgo return
	typeName := C.GoString(pDetectInfo.typeName)
	arrayLength := pDetectInfo.arrayLength
	key := C.GoString(pDetectInfo.key)

	fmt.Println("typeName: ", typeName)
	fmt.Println("key: ", key)
	fmt.Println("arrayLength: ", arrayLength)

	if arrayLength > uaConfig.ChannelConfig.MaxArrayLength {
		fmt.Println("Get Field(", key, ")'s arrayLength is toolarge:", arrayLength, " maxArrayLength is: ", uaConfig.ChannelConfig.MaxArrayLength)
		return
	}

	// loop for arrayLength to convert value.
	for i := 0; i < int(arrayLength); i++ {
		fmt.Println("data: ", *(*bool)(unsafe.Pointer(uintptr(pDetectInfo.data) + uintptr(i))))
	}

	fmt.Println("end....")
}

//func galloc(length int) (*C.char) {
//	pMem := (*C.char)(C.malloc(C.ulong(length)))
//	if pMem == nil {
//		fmt.Println("go malloc failed.")
//	}
//	return pMem
//}

//func copyStr(strDst *C.char, strSrc string) {
//	fmt.Println("strDst:: ", strDst)
//	fmt.Println("strsrc: ", strSrc)
//	C.strcpy(strDst, C.CString(strSrc))
//}

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