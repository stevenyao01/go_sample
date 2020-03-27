package main

import (
	//"crypto/md5"
	"errors"
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"os"
	"strconv"
	//"time"
)

func IsFileExist(filename string, filesize int64) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		//fmt.Println(info)
		return false
	}
	//if filesize == info.Size() {
	//        fmt.Println("安装包已存在！", info.Name(), info.Size(), info.ModTime())
	//        return true
	//    }
	del := os.Remove(filename)
	if del != nil {
		fmt.Println(del)
	}
	return false
}
func downloadFile(url string, localPath string, fb func(length, downLen int64)) error {
	var (
		fsize   int64
		buf     = make([]byte, 32*1024)
		written int64
	)
	tmpFilePath := localPath + ".download"
	fmt.Println(tmpFilePath)
	//创建一个http client
	client := new(http.Client)
	//client.Timeout = time.Second * 60 //设置超时时间
	//get方法获取资源
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	//读取服务器返回的文件大小
	fsize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	if IsFileExist(localPath, fsize) {
		return err
	}
	//fmt.Println("fsize", fsize)
	//创建文件
	file, err := os.Create(tmpFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if resp.Body == nil {
		return errors.New("body is null")
	}
	defer resp.Body.Close()
	//下面是 io.copyBuffer() 的简化版本
	for {
		//读取bytes
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			//写入bytes
			nw, ew := file.Write(buf[0:nr])
			//数据长度大于0
			if nw > 0 {
				written += int64(nw)
			}
			//写入出错
			if ew != nil {
				err = ew
				break
			}
			//读取是数据长度不等于写入的数据长度
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}

	}
	//没有错误了快使用 callback
	fb(fsize, written)

	if err == nil {
		file.Close()
		err = os.Rename(tmpFilePath, localPath)
	}
	return err
}
