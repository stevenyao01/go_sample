package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

func GetSk(broker string, dir string) error {
	strArr := strings.Split(broker, ":")
	errDownload := downloadFile("http://"+strArr[0]+":8080/services/certificate/download?uid=1", dir+"/device.sk", downloadCallback)
	if errDownload != nil {
		fmt.Println("errDownload: ", errDownload.Error())
		return errDownload
	}
	return nil
}

func UnzipFile(zipFileName string, targetPath string) error {
	// judge zip file
	if lIsZip(zipFileName) {
		//fmt.Println("the target is zip file.")
	} else {
		fmt.Println("the target is not zip file.")
		return errors.New("input file is not zip file")
	}
	// unzip packet
	errUnZip := lUnZip(zipFileName, targetPath)
	if errUnZip != nil {
		fmt.Println("errUnZip: ", errUnZip)
		return errUnZip
	}
	return nil
}

func SaveToFile(msg []byte, fileName string) error {
	var errOpenFile error
	fd, errOpenFile := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if errOpenFile != nil {
		return errOpenFile
	}

	wr := bufio.NewWriter(fd)

	_, errWrite := wr.Write(msg)
	if errWrite != nil {
		return errWrite
	}
	_ = wr.Flush()
	//fmt.Println("write ", wn, " words to ", fileName)
	return nil
}

func downloadCallback(length int64, len int64) () {
	//fmt.Println("下载的device.sk的文件大小: ", len)
	//fmt.Println("int: ", length, " len: ", len)
	return
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func StopProcess(p *os.Process, rt string) {
	runTime, errStr := strconv.Atoi(rt)
	if errStr != nil {
		fmt.Println("string to int error: ", errStr.Error())
	}
	time.Sleep(time.Duration(runTime) * time.Second)
	errKill := p.Kill()
	if errKill != nil {
		log.Println("errKill: ", errKill.Error())
	}
	log.Println("kill the test process.")
	return
}

func ReadStderr(path string, read, write *os.File) {
	defer read.Close()
	defer write.Close()
	var buf = make([]byte, 4*4096)
	for {
		n, err := read.Read(buf)
		if err != nil {
			log.Println("stderr read error: %s", err.Error())
			return
		}
		logArr := strings.Split(path, "/")
		errSaveToFile := SaveToFile(buf[:n], path+"/"+ logArr[1]+".log")
		if errSaveToFile != nil {
			log.Println("errSaveToFile: ", errSaveToFile.Error())
		}
	}
}


func CoreDump(filename string, err interface{}) error {
	stack := debug.Stack()

	filename = filename + "." + fmt.Sprintf("%d", time.Now().Unix())
	//fmt.Println(filename)

	var content []byte
	if err != nil {
		content = []byte(fmt.Sprintf("%s\n%s", err, stack))
	} else {
		content = []byte(fmt.Sprintf("%s", stack))
	}

	return ioutil.WriteFile(filename, content, 0644)
}

func isFileExist(filename string, filesize int64) bool {
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
	//fmt.Println(tmpFilePath)
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
	if isFileExist(localPath, fsize) {
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

