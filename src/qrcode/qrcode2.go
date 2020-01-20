
package main

import qrcode "github.com/skip2/go-qrcode"
import "fmt"

/**
 * @Package Name: main
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 19-3-12 上午10:46
 * @Description:
 */


func main() {
	err := qrcode.WriteFile("https://www.baidu.com", qrcode.Medium, 256, "qr.png")
	if err != nil {
		fmt.Println("write error")
	}
}