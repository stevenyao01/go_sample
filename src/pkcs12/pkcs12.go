package main

import (
	"encoding/base64"
	"crypto/rsa"
	"golang.org/x/crypto/pkcs12"
	"golang.org/x/crypto/tls"
	"fmt"
	"io/ioutil"
	"encoding/pem"
)

/**
 * @Package Name: pkcs12
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-10-6 下午4:54
 * @Description:
 */

const (
	pdfFile = "a.pfx"
)

func main() {
	p12, _ := base64.StdEncoding.DecodeString(`MIIJzgIBAzCCCZQGCS ... CA+gwggPk==`)

	blocks, err := pkcs12.ToPEM(p12, "password")
	if err != nil {
		panic(err)
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	// then use PEM data for tls to construct tls certificate:
	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		panic(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	_ = config


	//testdata, err := ioutil.ReadFile(pdfFile)
	//if err != nil {
	//	fmt.Println("read config file failed, filename: ", pdfFile)
	//}
	////for commonName, base64P12 := range testdata {
	//aa := string(testdata)
	//	p12, _ := base64.StdEncoding.DecodeString(aa)
	//
	//	priv, cert, err := pkcs12.Decode(p12, "")
	//	if err != nil {
	//		fmt.Println("err: ", err)
	//	}
	//
	//	if err := priv.(*rsa.PrivateKey).Validate(); err != nil {
	//		fmt.Println("error while validating private key: ", err)
	//	}
	//
	//	//if cert.Subject.CommonName != commonName {
	//	//	fmt.Println("expected common name to be %q, but found %q", commonName, cert.Subject.CommonName)
	//	//}
	//	fmt.Println("cert.Subject.CommonName: ", cert.Subject.CommonName)
	////}
}