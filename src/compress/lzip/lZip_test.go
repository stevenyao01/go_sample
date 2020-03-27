package lzip

import (
	"fmt"
	"testing"
)

const (
	zipFileName = "./20200228.zip"
	fileName = "20200228.txt"
	targetPath = "./cache/"
)

func TestZip(t *testing.T) {
	// zip file to packet test.zip
	errZip := lZip(zipFileName, fileName, "hello world.")
	if errZip != nil {
		fmt.Println("errZip is ", errZip.Error())
	}


}

func TestUnZip(t *testing.T) {
	// judge zip file
	if lIsZip(zipFileName) {
		fmt.Println("the target is zip file.")
	}else{
		fmt.Println("the target is not zip file.")
		return
	}

	// unzip packet
	errUnZip := lUnZip(zipFileName, targetPath)
	if errUnZip != nil {
		fmt.Println("errUnZip: ", errUnZip)
	}else {
		fmt.Println("UnZip file success.")
	}
}
