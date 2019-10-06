/*
 * =====================================================================================
 *
 *       Filename:  opcua.c
 *
 *    Description:  add for opcua
 *
 *        Version:  1.0
 *        Created:  2019年09月30日 12时53分28秒
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  steven yao,
 *        Company:  a lenovo co.ltd
 *
 * =====================================================================================
 */

#include <stdlib.h>
#include "opcua.h"


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

Ua_Read_Retval newOpcUaRetval(int len){
	Ua_Read_Retval urr;
	urr.Usn = malloc(sizeof(Ua_Single_Node));
	memset(urr.Usn, 0, sizeof(Ua_Single_Node));
	urr.Usn->arrayLength = malloc(sizeof(int) * len);
	memset(urr.Usn, 0, sizeof(int) * len);
	urr.Usn->typeName = malloc(sizeof(char *) * len);
	memset(urr.Usn->typeName, 0, sizeof(char *) * len);
	urr.Usn->key = malloc(sizeof(char *) * len);
	memset(urr.Usn->key, 0, sizeof(char *) * len);
	urr.Usn->data = malloc(sizeof(char *) * len);
	memset(urr.Usn->data, 0, sizeof(char *) * len);
	for(int i = 0; i < len; i++){
		urr.Usn->typeName[i] = malloc(sizeof(char) * NODEIDS_RET_TYPENAME_LENGTH);
		memset(urr.Usn->typeName[i], 0, sizeof(char) * NODEIDS_RET_TYPENAME_LENGTH);
		urr.Usn->key[i] = malloc(sizeof(char) * NODEIDS_RET_KEY_LENGTH);
		memset(urr.Usn->key[i], 0, sizeof(char) * NODEIDS_RET_KEY_LENGTH);
		urr.Usn->data[i] = malloc(sizeof(void) * NODEIDS_RET_DATA_LENGTH);
		memset(urr.Usn->data[i], 0, sizeof(void) * NODEIDS_RET_DATA_LENGTH);
	}

	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "newOPcUaRetval Success!!");
	return urr;
}

void deleteOpcUaRetval(Ua_Read_Retval urr, int length){
	for (int i = 0; i < length; i++) {
		free(urr.Usn->data[i]);
		free(urr.Usn->key[i]);
		free(urr.Usn->typeName[i]);
	}
	free(urr.Usn->data);
	free(urr.Usn->key);
	free(urr.Usn->typeName);
	free(urr.Usn->arrayLength);
	free(urr.Usn);
	return;
}

Opc_Ua_Config newOpcUaConfig(int len){

	Opc_Ua_Config ouc;
	ouc.NodeIds = malloc(sizeof(Ua_Node_Id));

	ouc.Security = malloc(sizeof(Ua_Security));
	memset(ouc.Security, 0, sizeof(Ua_Security));

	ouc.ChannelConfig = malloc(sizeof(Ua_Channel_Config));
	memset(ouc.ChannelConfig, 0, sizeof(Ua_Channel_Config));

	ouc.Config = malloc(sizeof(Ua_Connect_Config));
	memset(ouc.Config, 0, sizeof(Ua_Connect_Config));

	ouc.Credenials = malloc(sizeof(Ua_Credenials));
	memset(ouc.Credenials, 0, sizeof(Ua_Credenials));

	ouc.NodeIds->Identifier = malloc(sizeof(char*) * len);
	ouc.NodeIds->Field = malloc(sizeof(char*) * len);
	ouc.NodeIds->IdentifierType = malloc(sizeof(char*) * len);

	for (int i = 0; i < len; i++) {

		ouc.NodeIds->Identifier[i] = malloc(sizeof(char) * NODEIDS_IDENTIFIER_LENGTH);
		memset(ouc.NodeIds->Identifier[i], 0, sizeof(char) * NODEIDS_IDENTIFIER_LENGTH);

		ouc.NodeIds->Field[i] = malloc(sizeof(char) * NODEIDS_FIELD_LENGTH);
		memset(ouc.NodeIds->Field[i], 0, sizeof(char) * NODEIDS_FIELD_LENGTH);

		ouc.NodeIds->IdentifierType[i] = malloc(sizeof(char) * NODEIDS_IDENTIFIERTYPE_LENGTH);
		memset(ouc.NodeIds->IdentifierType[i], 0, sizeof(char) * NODEIDS_IDENTIFIERTYPE_LENGTH);
	}
	ouc.NodeIds->NamespaceIndex = malloc(sizeof(int) * len);
	memset(ouc.NodeIds->NamespaceIndex, 0, sizeof(int) * len);

	//UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "len: %d", len);
	//for (int j = 0; j < len; j++){
	//	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "identifier: %p", ouc.NodeIds->Identifier[j]);
	//	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "field: %p", ouc.NodeIds->Field[j]);
	//	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "IdentifierType: %p", ouc.NodeIds->IdentifierType[j]);
	//	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "NamespaceIndex: %p", &(ouc.NodeIds->NamespaceIndex[j]));
	//}

	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "newOpcUaConfig Success!!");

	return ouc;
}

void deleteOpcUaConfig(Opc_Ua_Config ouc, int length){
	//int length = sizeof(ouc.NodeIds->Identifier);
	UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "deleteOpcUaConfig begin.");
	for (int i = 0; i < length; i++) {
		free(ouc.NodeIds->Identifier[i]);
		free(ouc.NodeIds->Field[i]);
		free(ouc.NodeIds->IdentifierType[i]);
	}
	free(ouc.NodeIds->NamespaceIndex);

	free(ouc.NodeIds->Identifier);
	free(ouc.NodeIds->Field);
	free(ouc.NodeIds->IdentifierType);

	free(ouc.NodeIds);
	free(ouc.Security);
	free(ouc.ChannelConfig);
	free(ouc.Config);
	free(ouc.Credenials);

    UA_LOG_INFO(UA_Log_Stdout, UA_LOGCATEGORY_USERLAND, "deleteOpcUaConfig end.");
	return;
}
