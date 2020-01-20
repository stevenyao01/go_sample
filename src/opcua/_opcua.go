package main

/*
#cgo CFLAGS: -I ./open62541/include
#cgo CFLAGS: -I ./open62541/arch
#cgo CFLAGS: -I ./open62541/plugins/include
#cgo CFLAGS: -I ./open62541/build-shared64/src_generated
#cgo LDFLAGS: -L./open62541/build-shared64/bin/Debug -lopen62541

#include <open62541/client_config_default.h>
#include <open62541/client_highlevel.h>
#include <open62541/client_subscriptions.h>
#include <open62541/plugin/log_stdout.h>
#include <stdio.h>
#include <stdlib.h>

void handler_TheAnswerChanged(UA_Client *client, UA_UInt32 subId, void *subContext, UA_UInt32 monId, void *monContext, UA_DataValue *value);
UA_StatusCode nodeIter(UA_NodeId childId, UA_Boolean isInverse, UA_NodeId referenceTypeId, void *handle);
*/
import "C"
import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

//export handler_TheAnswerChanged
func handler_TheAnswerChanged(client *C.UA_Client, subId C.UA_UInt32, subContext unsafe.Pointer, monId C.UA_UInt32, monContext unsafe.Pointer, value *C.UA_DataValue) {
	//C.printf(C.CString("The Answer has changed!\n"))
}

//export nodeIter
func nodeIter(childId C.UA_NodeId, isInverse C.UA_Boolean, referenceTypeId C.UA_NodeId, handle unsafe.Pointer) C.UA_StatusCode {
	if isInverse {
		return C.UA_STATUSCODE_GOOD
	}
	parent := (*C.UA_NodeId)(handle)
	fmt.Printf("%d, %d --- %d ---> NodeId %d, %d\n",
		int(parent.namespaceIndex),
		binary.LittleEndian.Uint32(referenceTypeId.identifier[:]),
		binary.LittleEndian.Uint32(referenceTypeId.identifier[:]),
		int(childId.namespaceIndex),
		binary.LittleEndian.Uint32(childId.identifier[:]),
	)

	return C.UA_STATUSCODE_GOOD
}

func test() {
	client := C.UA_Client_new()
	C.UA_ClientConfig_setDefault(C.UA_Client_getConfig(client))

	var endpointArray *C.UA_EndpointDescription
	var endpointArraySize C.size_t
	retval := C.UA_Client_getEndpoints(client, C.CString("opc.tcp://10.111.66.220:4840"), &endpointArraySize, &endpointArray)
	if retval != C.UA_STATUSCODE_GOOD {
		C.UA_Array_delete(unsafe.Pointer(endpointArray), endpointArraySize, &C.UA_TYPES[C.UA_TYPES_ENDPOINTDESCRIPTION])
		C.UA_Client_delete(client)
		panic(C.EXIT_FAILURE)
	}

	//C.printf("%i endpoints found\n", (C.int)(endpointArraySize))
	fmt.Println(endpointArraySize, "endpoints found")

	for i := C.size_t(0); i < endpointArraySize; i++ {
		fmt.Printf("URL of endpoint %d is %.*s\n",
			int(i),
			int((*(*C.UA_EndpointDescription)(unsafe.Pointer(uintptr(unsafe.Pointer(endpointArray)) + uintptr(i)))).endpointUrl.length),
			string(C.GoString((*C.char)(unsafe.Pointer((*(*C.UA_EndpointDescription)(unsafe.Pointer(uintptr(unsafe.Pointer(endpointArray)) + uintptr(i)))).endpointUrl.data)))),
		)
		//C.printf("URL of endpoint %i is %.*s\n", C.int(i),
		//	C.int(endpointArray[i].endpointUrl.length),
		//	endpointArray[i].endpointUrl.data)
	}

	C.UA_Array_delete(unsafe.Pointer(endpointArray), endpointArraySize, &C.UA_TYPES[C.UA_TYPES_ENDPOINTDESCRIPTION])

	retval = C.UA_Client_connect(client, C.CString("opc.tcp://10.111.66.220:4840"))
	if retval != C.UA_STATUSCODE_GOOD {
		C.UA_Client_delete(client)
		panic(C.EXIT_FAILURE)
	}
	fmt.Println("Browsing nodes in objects folder:")
	var bReq C.UA_BrowseRequest

	C.UA_BrowseRequest_init(&bReq)
	bReq.requestedMaxReferencesPerNode = 0
	bReq.nodesToBrowse = C.UA_BrowseDescription_new()
	bReq.nodesToBrowseSize = 1
	(*(*C.UA_BrowseDescription)(unsafe.Pointer(bReq.nodesToBrowse))).nodeId = C.UA_NODEID_NUMERIC(0, C.UA_NS0ID_OBJECTSFOLDER)
	(*(*C.UA_BrowseDescription)(unsafe.Pointer(bReq.nodesToBrowse))).resultMask = C.UA_BROWSERESULTMASK_ALL
	bResp := C.UA_Client_Service_browse(client, bReq)

	fmt.Printf("%-9s %-16s %-16s %-16s\n", "NAMESPACE", "NODEID", "BROWSE NAME", "DISPLAY NAME")
	for i := C.size_t(0); i < bResp.resultsSize; i++ {
		for j := C.size_t(0); j < (*(*C.UA_BrowseResult)(unsafe.Pointer(uintptr(unsafe.Pointer(bResp.results)) + uintptr(i)))).referencesSize; j++ {
			results := *(*C.UA_BrowseResult)(unsafe.Pointer(uintptr(unsafe.Pointer(bResp.results)) + uintptr(i)))
			references := *(*C.UA_ReferenceDescription)(unsafe.Pointer(uintptr(unsafe.Pointer(results)) + uintptr(j)))
			//ref := &(bResp.results[i].references[j])
			ref := &references
			if ref.nodeId.nodeId.identifierType == C.UA_NODEIDTYPE_NUMERIC {
				//C.printf("%-9d %-16d %-16.*s %-16.*s\n", ref.nodeId.nodeId.namespaceIndex,
				//	ref.nodeId.nodeId.identifier.numeric, C.int(ref.browseName.name.length),
				//	ref.browseName.name.data, C.int(ref.displayName.text.length),
				//	ref.displayName.text.data)
				fmt.Printf("%-9d %-16d %-16.*s %-16.*s\n",
					int(ref.nodeId.nodeId.namespaceIndex),
					int(ref.nodeId.nodeId.identifier.numeric, ),
					int(ref.browseName.name.length),
					string(C.GoString(ref.browseName.name.data)),
					string(C.GoString(ref.displayName.text.data)),
				)
			} else if ref.nodeId.nodeId.identifierType == C.UA_NODEIDTYPE_STRING {
				//C.printf("%-9d %-16.*s %-16.*s %-16.*s\n", ref.nodeId.nodeId.namespaceIndex,
				//	C.int(ref.nodeId.nodeId.identifier.string.length),
				//	ref.nodeId.nodeId.identifier.string.data,
				//	C.int(ref.browseName.name.length), ref.browseName.name.data,
				//	C.int(ref.displayName.text.length), ref.displayName.text.data)
				fmt.Printf("%-9d %-16.*s %-16.*s %-16.*s\n",
					int(ref.nodeId.nodeId.namespaceIndex),
					string(C.GoString(ref.nodeId.nodeId.identifier.string.data)),
					string(C.GoString(ref.browseName.name.data)),
					string(C.GoString(ref.displayName.text.data)),
				)
			}
		}
	}

	C.UA_BrowseRequest_clear(&bReq)
	C.UA_BrowseResponse_clear(&bResp)

	parent := C.UA_NodeId_new()
	*parent = C.UA_NODEID_NUMERIC(0, C.UA_NS0ID_OBJECTSFOLDER)
	C.UA_Client_forEachChildNodeCall(client, C.UA_NODEID_NUMERIC(0, C.UA_NS0ID_OBJECTSFOLDER),
		C.nodeIter, unsafe.Pointer(parent))
	C.UA_NodeId_delete(parent)

	request := C.UA_CreateSubscriptionRequest_default()
	response := C.UA_Client_Subscriptions_create(client, request, nil, nil, nil)

	subId := response.subscriptionId
	if response.responseHeader.serviceResult == C.UA_STATUSCODE_GOOD {
		//C.printf("Create subscription succeeded, id %u\n", subId)
		fmt.Println("Create subscription succeeded, id", subId)
	}

	monRequest := C.UA_MonitoredItemCreateRequest_default(C.UA_NODEID_STRING(1, C.CString("the.answer")))
	monResponse := C.UA_Client_MonitoredItems_createDataChange(client, response.subscriptionId, C.UA_TIMESTAMPSTORETURN_BOTH, monRequest, nil, C.handler_TheAnswerChanged, nil)
	if monResponse.statusCode == C.UA_STATUSCODE_GOOD {
		//C.printf("Monitoring 'the.answer', id %u\n", monResponse.monitoredItemId)
		fmt.Println("Monitoring 'the.answer', id", monResponse.monitoredItemId)
	}
	C.UA_Client_run_iterate(client, 1000)
	var value C.UA_Int32 = 0
	//C.printf("\nReading the value of node (1, \"the.answer\"):\n")
	fmt.Println("\nReading the value of node (1, \"the.answer\"):")
	val := C.UA_Variant_new()
	retval = C.UA_Client_readValueAttribute(client, C.UA_NODEID_STRING(1, C.CString("the.answer")), val)
	if retval == C.UA_STATUSCODE_GOOD && C.UA_Variant_isScalar(val) && val._type == &C.UA_TYPES[C.UA_TYPES_INT32] {
		value = *(*C.UA_Int32)(val.data)
		//C.printf("the value is: %i\n", value)
		fmt.Println("the value is:", value)
	}
	C.UA_Variant_delete(val)

	value++
	//C.printf("\nWriting a value of node (1, \"the.answer\"):\n")
	fmt.Println("\nWriting a value of node (1, \"the.answer\"):")
	var wReq C.UA_WriteRequest
	C.UA_WriteRequest_init(&wReq)
	wReq.nodesToWrite = C.UA_WriteValue_new()
	wReq.nodesToWriteSize = 1
	wReq.nodesToWrite[0].nodeId = C.UA_NODEID_STRING_ALLOC(1, C.CString("the.answer"))
	wReq.nodesToWrite[0].attributeId = C.UA_ATTRIBUTEID_VALUE
	wReq.nodesToWrite[0].value.hasValue = C.true
	wReq.nodesToWrite[0].value.value._type = &C.UA_TYPES[C.UA_TYPES_INT32]
	wReq.nodesToWrite[0].value.value.storageType = C.UA_VARIANT_DATA_NODELETE
	wReq.nodesToWrite[0].value.value.data = &value
	wResp := C.UA_Client_Service_write(client, wReq)
	if wResp.responseHeader.serviceResult == C.UA_STATUSCODE_GOOD {
		//C.printf("the new value is: %i\n", value)
		fmt.Println("the new value is:", value)
	}
	C.UA_WriteRequest_clear(&wReq)
	C.UA_WriteResponse_clear(&wResp)

	value++
	myVariant := C.UA_Variant_new()
	C.UA_Variant_setScalarCopy(myVariant, &value, &C.UA_TYPES[C.UA_TYPES_INT32])
	C.UA_Client_writeValueAttribute(client, C.UA_NODEID_STRING(1, "the.answer"), myVariant)
	C.UA_Variant_delete(myVariant)

	C.UA_Client_run_iterate(client, 100)
	if C.UA_Client_Subscriptions_deleteSingle(client, subId) == C.UA_STATUSCODE_GOOD {
		//C.printf("Subscription removed\n")
		fmt.Println("Subscription removed")
	}

	var input C.UA_Variant
	argString := C.UA_STRING("Hello Server")
	C.UA_Variant_init(&input)
	C.UA_Variant_setScalarCopy(&input, &argString, &C.UA_TYPES[C.UA_TYPES_STRING])
	var outputSize C.size_t
	var output *C.UA_Variant
	retval = C.UA_Client_call(client, C.UA_NODEID_NUMERIC(0, C.UA_NS0ID_OBJECTSFOLDER), C.UA_NODEID_NUMERIC(1, 62541), 1, &input, &outputSize, &output)
	if retval == C.UA_STATUSCODE_GOOD {
		//C.printf("Method call was successful, and %lu returned values available.\n", C.ulong(outputSize))
		fmt.Printf("Method call was successful, and %d returned values available.\n", uint(outputSize))
		C.UA_Array_delete(output, outputSize, &C.UA_TYPES[C.UA_TYPES_VARIANT])
	} else {
		//C.printf("Method call was unsuccessful, and %x returned values available.\n", retval)
		fmt.Printf("Method call was unsuccessful, and %x returned values available.\n", int(retval))
	}
	C.UA_Variant_clear(&input)

	var ref_id C.UA_NodeId
	ref_attr := C.UA_ReferenceTypeAttributes_default
	ref_attr.displayName = C.UA_LOCALIZEDTEXT("en-US", "NewReference")
	ref_attr.description = C.UA_LOCALIZEDTEXT("en-US", "References something that might or might not exist")
	ref_attr.inverseName = C.UA_LOCALIZEDTEXT("en-US", "IsNewlyReferencedBy")
	retval = C.UA_Client_addReferenceTypeNode(client,
		C.UA_NODEID_NUMERIC(1, 12133),
		C.UA_NODEID_NUMERIC(0, C.UA_NS0ID_ORGANIZES),
		C.UA_NODEID_NUMERIC(0, C.UA_NS0ID_HASSUBTYPE),
		C.UA_QUALIFIEDNAME(1, "NewReference"),
		ref_attr, &ref_id)
	if retval == C.UA_STATUSCODE_GOOD {
		//C.printf("Created 'NewReference' with numeric NodeID %u\n", ref_id.identifier.numeric)
		fmt.Println("Created 'NewReference' with numeric NodeID", uint(ref_id.identifier.numeric))
	}

	var objt_id C.UA_NodeId
	objt_attr := C.UA_ObjectTypeAttributes_default
	objt_attr.displayName = C.UA_LOCALIZEDTEXT("en-US", "TheNewObjectType")
	objt_attr.description = C.UA_LOCALIZEDTEXT("en-US", "Put innovative description here")
	retval = C.UA_Client_addObjectTypeNode(client,
		C.UA_NODEID_NUMERIC(1, 12134),
		C.UA_NODEID_NUMERIC(0, C.UA_NS0ID_BASEOBJECTTYPE),
		C.UA_NODEID_NUMERIC(0, C.UA_NS0ID_HASSUBTYPE),
		C.UA_QUALIFIEDNAME(1, "NewObjectType"),
		objt_attr, &objt_id)
	if retval == C.UA_STATUSCODE_GOOD {
		//C.printf("Created 'NewObjectType' with numeric NodeID %u\n", objt_id.identifier.numeric)
		fmt.Println("Created 'NewObjectType' with numeric NodeID", uint(objt_id.identifier.numeric))
	}

	var obj_id C.UA_NodeId
	obj_attr := C.UA_ObjectAttributes_default
	obj_attr.displayName = C.UA_LOCALIZEDTEXT("en-US", "TheNewGreatNode")
	obj_attr.description = C.UA_LOCALIZEDTEXT("de-DE", "Hier koennte Ihre Webung stehen!")
	retval = C.UA_Client_addObjectNode(client,
		C.UA_NODEID_NUMERIC(1, 0),
		C.UA_NODEID_NUMERIC(0, C.UA_NS0ID_OBJECTSFOLDER),
		C.UA_NODEID_NUMERIC(0, C.UA_NS0ID_ORGANIZES),
		C.UA_QUALIFIEDNAME(1, "TheGreatNode"),
		C.UA_NODEID_NUMERIC(1, 12134),
		obj_attr, &obj_id)
	if retval == C.UA_STATUSCODE_GOOD {
		//C.printf("Created 'NewObject' with numeric NodeID %u\n", obj_id.identifier.numeric)
		fmt.Println("Created 'NewObject' with numeric NodeID", uint(obj_id.identifier.numeric))
	}

	var var_id C.UA_NodeId
	var_attr := C.UA_VariableAttributes_default
	var_attr.displayName = C.UA_LOCALIZEDTEXT("en-US", "TheNewVariableNode")
	var_attr.description = C.UA_LOCALIZEDTEXT("en-US", "This integer is just amazing - it has digits and everything.")
	var int_value = C.UA_Int32(1234)

	C.UA_Variant_setScalar(&var_attr.value, &int_value, &C.UA_TYPES[C.UA_TYPES_INT32])
	var_attr.dataType = C.UA_TYPES[C.UA_TYPES_INT32].typeId
	retval = C.UA_Client_addVariableNode(client,
		C.UA_NODEID_NUMERIC(1, 0),
		C.UA_NODEID_NUMERIC(0, C.UA_NS0ID_OBJECTSFOLDER),
		C.UA_NODEID_NUMERIC(0, C.UA_NS0ID_ORGANIZES),
		C.UA_QUALIFIEDNAME(0, "VariableNode"),
		C.UA_NODEID_NULL,
		var_attr, &var_id)
	if retval == C.UA_STATUSCODE_GOOD {
		//C.printf("Created 'NewVariable' with numeric NodeID %u\n", var_id.identifier.numeric)
		fmt.Println("Created 'NewVariable' with numeric NodeID", uint(var_id.identifier.numeric))
	}

	C.UA_Client_disconnect(client)
	C.UA_Client_delete(client)
}
