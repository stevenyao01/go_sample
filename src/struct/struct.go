package main

import (
	"fmt"
	"reflect"
)


type EnvKey struct {
	RD 	string
	BR 	string
	PT 	string
	AD 	string
	LE 	string
}

func main() {
	t := EnvKey{"/dev/ttyUSB0", "19200", "8080", "1105", "20"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

//var arraySlice = []string{
//	"hello",
//	"yao",
//	"apple",
//	"google"}
//
//func main(){
//	for k, v := range arraySlice {
//		fmt.Println("k: ", k)
//		fmt.Println("v: ", v)
//	}
//}
