package channel

import (
	"fmt"
)

type info struct {
	ab 	int
	bc 	string
	cd  string
}

func main1() {
	var ch = make(chan info)
	go routine1(ch)
	for true {
		value := <- ch
		if value.ab == 1 {
			fmt.Println("bc:", value.bc)
			fmt.Println("cd:", value.cd)
		}else if value.ab == 5 {
			fmt.Println("bc:", value.bc)
			fmt.Println("cd:", value.cd)
			break
		}

	}
	close(ch)
	fmt.Println("main end!!!")
}

func routine1(ch chan info) {
	fmt.Println("my new routine1")
	var ii info
	ii.ab = 1
	ii.bc = "ehllo"
	ii.cd = "zzq"
	ch <- ii
	var iii info
	iii.ab = 5
	iii.bc = "ehlloasdfsdf"
	iii.cd = "zzqasdfsadf"
	ch <- iii
	return
}
