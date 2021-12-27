package tcp

type Msg struct {
	// protocol id
	ProtocolId int `json:"protocolId,omitempty"`
	// protocol md5
	ProtocolMd5 string `json:"protocolMd5,omitempty"`
	// protocol excuter
	ProtocolExcuter string `json:"executer,omitempty"`

	// input conf
	IpConf string `json:"inputConf,omitempty"`
	// output conf
	OpConf string `json:"outputConf,omitempty"`
	// other conf
	OtConf string `json:"otherConf,omitempty"`
}

type Config struct {
	// msg type
	MsgType string `json:"msgType,omitempty"`
	// msg id
	MsgId string `json:"msgId,omitempty"`
	// msg version
	MsgVersion string `json:"msgVersion,omitempty"`
	// msg struct
	Msg Msg `json:"msg,omitempty"`

}