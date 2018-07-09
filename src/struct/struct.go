package main

import (
	"fmt"
)

var arraySlice = []string{
	"hello",
	"yao",
	"apple",
	"google"}

func main(){
	for k, v := range arraySlice {
		fmt.Println("k: ", k)
		fmt.Println("v: ", v)
	}
}
