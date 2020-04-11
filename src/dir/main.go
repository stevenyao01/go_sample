package main

import (
	"errors"
	//	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	err := CopyDir("/home/steven/code/go/src/github.com/vjeantet/grok_bak", "/home/steven/code/go/src/github.com/vjeantet/grok_yao")
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
}

func IsDir(name string) bool {
	if info, err := os.Stat(name); err == nil {
		return info.IsDir()
	}
	return false
}

func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func MakeDir(dir string) error {
	if !FileIsExisted(dir) {
		if err := os.MkdirAll(dir, 0777); err != nil { //os.ModePerm
			fmt.Println("MakeDir failed:", err)
			return err
		}
	}
	return nil
}

//复制文件过程中一定要注意将原始文件的权限也要复制过去，否则可能会导致可执行文件不能执行等问题。

//使用io.Copy
func CopyFile(src, des string) (written int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	//获取源文件的权限
	fi, _ := srcFile.Stat()
	perm := fi.Mode()

	//desFile, err := os.Create(des)  //无法复制源文件的所有权限
	desFile, err := os.OpenFile(des, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)  //复制源文件的所有权限
	if err != nil {
		return 0, err
	}
	defer desFile.Close()

	return io.Copy(desFile, srcFile)
}


//使用ioutil.WriteFile()和ioutil.ReadFile()
func CopyFile2(src, des string) (written int64, err error) {
	//获取源文件的权限
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	fi, _ := srcFile.Stat()
	perm := fi.Mode()
	srcFile.Close()

	input, err := ioutil.ReadFile(src)
	if err != nil {
		return 0, err
	}

	err = ioutil.WriteFile(des, input, perm)
	if err != nil {
		return 0, err
	}

	return int64(len(input)), nil
}


//使用os.Read()和os.Write()
func CopyFile3(src, des string, bufSize int) (written int64, err error) {
	if bufSize <= 0 {
		bufSize = 1*1024*1024   //1M
	}
	buf := make([]byte, bufSize)

	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	//获取源文件的权限
	fi, _ := srcFile.Stat()
	perm := fi.Mode()

	desFile, err := os.OpenFile(des, os.O_CREATE|os.O_RDWR|os.O_TRUNC, perm)
	if err != nil {
		return 0, err
	}
	defer desFile.Close()

	count := 0
	for {
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		if n == 0 {
			break
		}

		if wn, err := desFile.Write(buf[:n]); err != nil {
			return 0, err
		} else {
			count += wn
		}
	}

	return int64(count), nil
}

func CopyDir(srcPath, desPath string) error {
	//检查目录是否正确
	if srcInfo, err := os.Stat(srcPath); err != nil {
		return err
	} else {
		if !srcInfo.IsDir() {
			return errors.New("源路径不是一个正确的目录！")
		}
	}

	if desInfo, err := os.Stat(desPath); err != nil {
		errMakeDir := MakeDir(desPath)
		if errMakeDir != nil {
			fmt.Println("MakeDir err: ", errMakeDir.Error())
		}
	} else {
		if !desInfo.IsDir() {
			return errors.New("目标路径不是一个正确的目录！")
		}
	}

	if strings.TrimSpace(srcPath) == strings.TrimSpace(desPath) {
		return errors.New("源路径与目标路径不能相同！")
	}

	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		//复制目录是将源目录中的子目录复制到目标路径中，不包含源目录本身
		if path == srcPath {
			return nil
		}

		//生成新路径
		destNewPath := strings.Replace(path, srcPath, desPath, -1)

		if !f.IsDir() {
			CopyFile(path, destNewPath)
		} else {
			if !FileIsExisted(destNewPath) {
				return MakeDir(destNewPath)
			}
		}

		return nil
	})

	return err
}

/* 获取指定路径下的所有文件，只搜索当前路径，不进入下一级目录，可匹配后缀过滤（suffix为空则不过滤）*/
func ListDir(dir, suffix string) (files []string, err error) {
	files = []string{}

	_dir, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	suffix = strings.ToLower(suffix)  //匹配后缀

	for _, _file := range _dir {
		if _file.IsDir() {
			continue   //忽略目录
		}
		if len(suffix) == 0 || strings.HasSuffix(strings.ToLower(_file.Name()), suffix) {
			//文件后缀匹配
			files = append(files, path.Join(dir, _file.Name()))
		}
	}

	return files, nil
}

/* 获取指定路径下以及所有子目录下的所有文件，可匹配后缀过滤（suffix为空则不过滤）*/
func WalkDir(dir, suffix string) (files []string, err error) {
	files = []string{}

	err = filepath.Walk(dir, func(fname string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			//忽略目录
			return nil
		}

		if len(suffix) == 0 || strings.HasSuffix(strings.ToLower(fi.Name()), suffix) {
			//文件后缀匹配
			files = append(files, fname)
		}

		return nil
	})

	return files, err
}
//8. 删除文件
//os.Remove(filename)
//9. 删除文件夹及其包含的所有子目录和所有文件
//os.RemoveAll(dir)
