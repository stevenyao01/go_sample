
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

//客户端对服务器校验
func main() {
	//CertPool代表一个证书集合/证书池。
	//创建一个CertPool
	pool := x509.NewCertPool()
	caCertPath := "/home/liuxin/Desktop/ca.crt"
	//调用ca.crt文件
	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	//解析证书
	pool.AppendCertsFromPEM(caCrt)

	tr := &http.Transport{
		////把从服务器传过来的非叶子证书，添加到中间证书的池中，使用设置的根证书和中间证书对叶子证书进行验证。
		TLSClientConfig: &tls.Config{RootCAs: pool},

		//TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //InsecureSkipVerify用来控制客户端是否证书和服务器主机名。如果设置为true,//		//则不会校验证书以及证书中的主机名和服务器主机名是否一致。


	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8080")
	if err != nil {
		fmt.Println("Get error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
