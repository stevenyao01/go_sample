package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/go_sample/src/tsfile/common/log"
	"github.com/json-iterator/go"
	"time"
)

type request struct {
	CompanyId  int    `json:"company_id"`
	CompanySk  string `json:"company_sk"`
	DeviceId   string `json:"device_id"`
	DeviceDesc string `json:"device_desc"`
}

func main2(){
	request1 := request{100000, "1234567890qwertyuiop", "8888", "223.203.201.251:8200"}
	data, _ := json.Marshal(request1)
	log.Info("data: %s", string(data))
	request2 := request{100, "1234567890qwertyuiop", "8888", "223.203.201.251:8200"}
	var jsyao = jsoniter.ConfigCompatibleWithStandardLibrary
	t1 := time.Now()
	_ = jsyao.Unmarshal(data, &request2)
	json.Unmarshal(data, &request2)
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

/* define new struct MyStruct for marshal*/
type MyStruct struct {
	Key1          string `json:"key1"`
	Key2          string `json:"key2"`
	Key3          string `json:"key3"`
	Key4          string `json:"key4"`
	ComputingInfo string `json:"computingInfo"`
}

func main5() {
	ms := new(MyStruct)

	ms.Key1 = "kkey1"
	ms.Key2 = "kkey2"
	ms.Key3 = "kkey3"
	ms.Key4 = "kkey4"
	//ms.ComputingInfo = "ccccccomputingInfo"

	b, _ := json.Marshal(ms)
	cnStr := string(b)
	fmt.Println("cnStr: ", cnStr)
	cnJson, _ := simplejson.NewJson([]byte(cnStr))
	cnBody := cnJson.Get("computingInfo")
	fmt.Println("getStr: ", cnBody)

	cnJson.Del("computingInfo")
	fmt.Println("cnBody: ", cnJson)
	str, _ := cnJson.MarshalJSON()
	fmt.Println("string data", string(str))

}

func main() {
	//拼凑json   body为map数组
	//var rbody []map[string]interface{}
	//t := make(map[string]interface{})
	//t["device_id"] = "dddddd"
	//t["device_hid"] = "ddddddd"
	//
	//rbody = append(rbody, t)
	//t1 := make(map[string]interface{})
	//t1["device_id"] = "aaaaa"
	//t1["device_hid"] = "aaaaa"
	//
	//rbody = append(rbody, t1)

	sli:=make([]int ,0)
	sli = append(sli, 1)
	sli = append(sli, 2)
	sli = append(sli, 3)


	cnnJson := make(map[string]interface{})
	cnnJson["agenttime"] = "2342342442"
	cnnJson["message"] = sli
	cnnJson["description"] = "get all active worker"


	b, _ := json.Marshal(cnnJson)
	var obj interface{}
	if err := json.Unmarshal(b, &obj); err != nil {
		fmt.Println("unmarshal err: ", err.Error())
	}
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


