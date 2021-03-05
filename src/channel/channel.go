package main

import (
	"fmt"
)

type info struct {
	ab 	int
	bc 	string
	cd  string
}

func main() {
	var chi = make(chan int)
	//go routine001(chi)
	//go routine002(chi)
	//go routine003(chi)
	//go routine004(chi)
	//go routine005(chi)
	//go routine006(chi)
	//go routine007(chi)
	//go routine008(chi)
	//go routine009(chi)




    go func(chi chan int) {
		for true {
			value := <- chi
			fmt.Println("value: ", value)
			if value == 9 {
				break
			}
		}
	}(chi)

	routine001(chi)
	routine002(chi)
	routine003(chi)
	routine004(chi)
	routine005(chi)
	routine006(chi)
	routine007(chi)
	routine008(chi)
	routine009(chi)

	close(chi)
	fmt.Println("main end!!!")
}

func routine001(chi chan int) {
	fmt.Println("my new routine001")
	chi <- 1
	return
}

func routine002(chi chan int) {
	fmt.Println("my new routine002")
	chi <- 2
	return
}

func routine003(chi chan int) {
	fmt.Println("my new routine003")
	chi <- 3
	return
}

func routine004(chi chan int) {
	fmt.Println("my new routine004")
	chi <- 4
	return
}

func routine005(chi chan int) {
	fmt.Println("my new routine005")
	chi <- 5
	return
}

func routine006(chi chan int) {
	fmt.Println("my new routine006")
	chi <- 6
	return
}

func routine007(chi chan int) {
	fmt.Println("my new routine007")
	chi <- 7
	return
}

func routine008(chi chan int) {
	fmt.Println("my new routine008")
	chi <- 8
	return
}

func routine009(chi chan int) {
	fmt.Println("my new routine009")
	chi <- 9
	return
}

func routine010(chi chan int) {
	fmt.Println("my new routine010")
	chi <- 10
	return
}

//func routine2(chi chan int) {
//	fmt.Println("my new routine1")
//
//	for i := 0; i < 10 ; i++ {
//		chi <- i
//	}
//	return
//}


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
