package main

import (
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"github.com/cyberdelia/lzo"
	"github.com/go_sample/src/compress/utils"
	"github.com/golang/snappy"
	"io"
	"os"
	"time"
)

type CompressMode int

const (
	Unknow     CompressMode = iota
	FlateMode               // falte
	GzipMode                // gzip
	LzoMode                 // lzo
	SnappyMode              // snappy
	ZlibMode                // zlib
)

type Compresser struct {
	inFile     *os.File
	ouFile     *os.File
	absPath    string
	originSize int64
	writer     io.Writer
	modeStr    string
}

func New(input string) (*Compresser, error) {
	fi, err := os.Stat(input)
	if err != nil {
		return nil, err
	}
	if fi.IsDir() {
		return nil, fmt.Errorf("error read: because  %s is dir", input)
	}
	c := &Compresser{originSize: fi.Size(), absPath: fi.Name()}
	c.inFile, err = os.Open(input)
	return c, err
}
func (c *Compresser) UseGzipWriter(out string) (io.Writer, error) {
	file, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	c.ouFile = file
	w, err := gzip.NewWriterLevel(file, gzip.BestSpeed)
	return w, err
}
func (c *Compresser) UseLzoWriter(out string) (io.Writer, error) {
	file, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	c.ouFile = file
	w, err := lzo.NewWriterLevel(file, lzo.BestSpeed)
	return w, err
}
func (c *Compresser) UseSnappyWriter(out string) (io.Writer, error) {
	file, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	c.ouFile = file
	return snappy.NewBufferedWriter(file), nil
}
func (c *Compresser) UseZlibWriter(out string) (io.Writer, error) {
	file, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	c.ouFile = file
	w, err := zlib.NewWriterLevel(file, zlib.BestSpeed)
	return w, err
}
func (c *Compresser) UseFlateWriter(out string) (io.Writer, error) {
	file, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}
	c.ouFile = file
	w, err := flate.NewWriter(file, flate.BestSpeed)
	return w, err
}

func (c *Compresser) WithMode(mode CompressMode) (*Compresser, error) {
	var writer io.Writer
	var err error
	var modStr string
	switch mode {
	case GzipMode:
		modStr = "GZIP"
		writer, err = c.UseGzipWriter(fmt.Sprintf("%s.gz", c.inFile.Name()))
	case LzoMode:
		modStr = "LZO"
		writer, err = c.UseLzoWriter(fmt.Sprintf("%s.lzo", c.inFile.Name()))
	case SnappyMode:
		modStr = "SNAPPY"
		writer, err = c.UseSnappyWriter(fmt.Sprintf("%s.snap", c.inFile.Name()))
	case ZlibMode:
		modStr = "ZLIB"
		writer, err = c.UseZlibWriter(fmt.Sprintf("%s.zlib", c.inFile.Name()))
	case FlateMode:
		modStr = "FLATE"
		writer, err = c.UseFlateWriter(fmt.Sprintf("%s.flate", c.inFile.Name()))

	}

	c.writer = writer
	c.modeStr = modStr
	return c, err
}

// 启动压缩 并计算时间,压缩前后文件大小
func (c *Compresser) Compress(log bool) {
	if log {
		startTime := time.Now()
		defer c.ouFile.Close()
		defer c.inFile.Close()
		defer func() {
			nano := time.Since(startTime)
			fi, _ := c.ouFile.Stat()
			fmt.Printf("使用%s压缩, 原始文件大小:%s, 压缩后大小:%s, 耗时:%v\n",
				c.modeStr, utils.Bytes(uint64(c.originSize)), utils.Bytes(uint64(fi.Size())), nano)
		}()
	}

	//buf := make([]byte,1024*100)
	//_,err := io.CopyBuffer(c.writer,c.inFile,buf)

	_, err := io.Copy(c.writer, c.inFile)
	if err != nil {
		panic(err)
	}

}
