package write

/**
 * @Package Name: main
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-8-22 下午4:34
 * @Description:
 */

import (
	"fmt"
	"os"
	"time")

func benchmarkFileWrite(filename string, n int, index int) (d time.Duration) {
	v := "ni shuo wo shi bu shi tai wu liao le a?"
	os.Remove(filename)
	t0 := time.Now()
	switch index {
	case 0: // open file and write, then close, repeat n times
		for i := 0; i < n; i++ {
			file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(index, i, "open file failed.", err.Error())
				break
			}
			file.WriteString(v)
			file.WriteString("\n")
			file.Close()
		}
	case 1: // open file and write, defer close, repeat n times
		for i := 0; i < n; i++ {
			file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(index, i, "open file failed.", err.Error())
				break
			}
			defer file.Close()
			file.WriteString(v)
			file.WriteString("\n")
		}
	case 2: // open file and write n times, then close file
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(index, "open file failed.", err.Error())
			break
		}
		defer file.Close()
		for i := 0; i < n; i++ {
			file.WriteString(v)
			file.WriteString("\n")
		}
	}
	t1 := time.Now()
	d = t1.Sub(t0)
	fmt.Printf("time of way(%d)=%v\n", index, d)
	return d
}

func main1() {
	const k, n int = 3, 5000
	d := [k]time.Duration{}
	for i := 0; i < k; i++ {
		d[i] = benchmarkFileWrite("benchmarkFile.txt", n, i)
	}
	for i := 0; i < k-1; i++ {
		fmt.Printf("way %d cost time is %6.1f times of way %d\n", i, float32(d[i])/float32(d[k-1]), k-1)
	}
}
