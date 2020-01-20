/*
 * =====================================================================================
 *
 *       Filename:  opcua.h
 *
 *    Description:  add for opcua 
 *
 *        Version:  1.0
 *        Created:  2019年09月30日 12时06分55秒
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  steven yao
 *        Company:  a lenovo co.ltd
 *
 * =====================================================================================
 */

#ifndef UA_OPCUA_H_
#define UA_OPCUA_H_

#include <stdlib.h>
#include <open62541/types.h>
#include <open62541/types_generated_handling.h>
#include <open62541/plugin/log_stdout.h>

#define UA_ENABLE_ENCRYPTION true
// define for ua config
#define NODEIDS_IDENTIFIER_LENGTH 512
#define NODEIDS_FIELD_LENGTH 32
#define NODEIDS_IDENTIFIERTYPE_LENGTH 32
#define NODEIDS_NAMESPACEINDEX_LENGTH 32
// define for ua return value
#define NODEIDS_RET_TYPENAME_LENGTH 32
#define NODEIDS_RET_KEY_LENGTH 32
#define NODEIDS_RET_DATA_LENGTH 512
#define NODEIDS_RET_ARRAY_LENGTH 4

typedef struct {
    char *typeName;
    char *key;
	int  arrayLength;
	void *data;
} Ua_Sub_Node;

typedef struct {
    char **typeName;
    char **key;
	int   *arrayLength;
	void  **data;
} Ua_Single_Node;

typedef struct {
    Ua_Single_Node *Usn;
} Ua_Read_Retval;

typedef struct{
	int  *NamespaceIndex;
	char **Identifier;
	char **Field;
	char **IdentifierType;
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

void copyStr(char *strDst, char *strSrc);
void* galloc(int length);
Ua_Read_Retval newOpcUaRetval(int len);
void deleteOpcUaRetval(Ua_Read_Retval urr, int length);
Opc_Ua_Config newOpcUaConfig(int len);
void deleteOpcUaConfig(Opc_Ua_Config ouc, int length);

#endif