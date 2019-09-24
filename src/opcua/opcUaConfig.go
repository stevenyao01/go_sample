package main

/**
 * @Package Name: yao
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-9-12 下午2:56
 * @Description:
 */

import (
 	"encoding/json"
	"io/ioutil"
	"fmt"
)

type NodeId struct {
	Identifier     string `json:"identifier"`
	Field          string `json:"field"`
	IdentifierType string `json:"identifierType"`
	NamespaceIndex uint32 `json:"namespaceIndex"`
}

type Security struct {
	Password         string `json:"password"`
	StoreType        string `json:"storeType"`
	KeystoreFilePath string `json:"keystoreFilePath"`
	Alias            string `json:"alias"`
	SecurityPolicy   string `json:"securityPolicy"`
}

type ChannelConfig struct {
	MaxChunkCount   uint32 `json:"maxChunkCount"`
	MaxArrayLength  uint32 `json:"maxArrayLength"`
	MaxMessageSize  uint32 `json:"maxMessageSize"`
	MaxStringLength uint32 `json:"maxStringLength"`
	MaxChunkSize    uint32 `json:"maxChunkSize"`
}

type Config struct {
	ResourceUrl     string `json:"resourceurl"`
	UseCredenials   bool   `json:"useCredenials"`
	PollingInterval uint32 `json:"pollingInterval"`
	ApplicationUrl  string `json:"applicationUrl"`
	SessionTimeOut  uint32 `json:"sessionTimeOut"`
	ProcessingMode  string `json:"processingMode"`
	RequestTimeOut  uint32 `json:"requestTimeOut"`
	ReconnectTime   uint32 `json:"reConnectTime"`
}

//typed Ua_Credenials struct {
//char *userName;
//char *passWord;
//}

type OpcUaConfig struct {
	NodeIds       []NodeId      `json:"nodeIds"`
	Security      Security      `json:"security"`
	ChannelConfig ChannelConfig `json:"channelConfig"`
	Config        Config        `json:"config"`
}

func (o *OpcUaConfig) Fix() () {
	fmt.Println("add default value.")
	return
}

func (o *OpcUaConfig) String() (string) {
	return string(o.Json())
}

func (o *OpcUaConfig) Json() ([]byte) {
	j, err := json.MarshalIndent(o, "", "\t")
	if err != nil {
		panic(err)
	}

	return j
}

func (o *OpcUaConfig) ToFile(filename string) error {
	//o.fix()
	o.Fix()
	return ioutil.WriteFile(filename, o.Json(), 0755)
}

func NewOpcUaConfig()(*OpcUaConfig, error) {
	return &OpcUaConfig{
	}, nil
}
