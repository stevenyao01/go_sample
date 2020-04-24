package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/json-iterator/go"
	"github.com/go_sample/src/tsfile/common/log"
	"time"
)

type request struct {
	Company_id              int             `json:"company_id"`
	Company_sk              string          `json:"company_sk"`
	Device_id              	string          `json:"device_id"`
	Device_desc             string          `json:"device_desc"`
}

func main(){
	request1 := request{100000, "1234567890qwertyuiop", "8888", "223.203.201.251:8200"}
	data, _ := json.Marshal(request1)
	log.Info("data: %s", string(data))
	request2 := request{100, "1234567890qwertyuiop", "8888", "223.203.201.251:8200"}
	var jsyao = jsoniter.ConfigCompatibleWithStandardLibrary
	t1 := time.Now()
	jsyao.Unmarshal(data, &request2)
	//json.Unmarshal(data, &request2)
	nano := time.Since(t1)
	log.Info("nano: ", nano)
}

func main1() {
	//拼凑json   body为map数组
	var rbody []map[string]interface{}
	t := make(map[string]interface{})
	t["device_id"] = "dddddd"
	t["device_hid"] = "ddddddd"

	rbody = append(rbody, t)
	t1 := make(map[string]interface{})
	t1["device_id"] = "aaaaa"
	t1["device_hid"] = "aaaaa"

	rbody = append(rbody, t1)

	cnnJson := make(map[string]interface{})
	cnnJson["code"] = 0
	cnnJson["request_id"] = 123
	cnnJson["code_msg"] = ""
	cnnJson["body"] = rbody
	cnnJson["page"] = 0
	cnnJson["page_size"] = 0

	b, _ := json.Marshal(cnnJson)
	cnnn := string(b)
	fmt.Println("cnnn:%s", cnnn)
	cn_json, _ := simplejson.NewJson([]byte(cnnn))
	cn_body, _ := cn_json.Get("body").Array()

	for _, di := range cn_body {
		//就在这里对di进行类型判断
		newdi, _ := di.(map[string]interface{})
		device_id := newdi["device_id"]
		device_hid := newdi["device_hid"]
		fmt.Println(device_hid, device_id)
	}

}


