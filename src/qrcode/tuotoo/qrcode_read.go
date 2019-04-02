package main

import (
	"fmt"
	"os"

	"github.com/tuotoo/qrcode"
)

/**
 * @Package Name: tuotoo
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-3-12 上午10:50
 * @Description:
 */


func main() {

	fi, err := os.Open("qr.png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(qrmatrix.Content)
}