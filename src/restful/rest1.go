package main

import "fmt"

/**
 * @Package Name: restful
 * @Author: steven yao
 * @Email:  yhp.linux@gmail.com
 * @Create Date: 18-10-29 下午2:58
 * @Description:
 */

import (
	//"net/http"
	////"fmt"
	//"html"
	//"log"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

//func main() {
//	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
//		fmt.Fprintln(writer, "Hello, ", html.EscapeString(request.URL.Path))
//	})
//	log.Fatal(http.ListenAndServe(":8080",nil))
//}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}
func TodoShow(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	todoId := vars["todoId"]
	fmt.Fprintln(writer, "Todo show:", todoId)
}
func TodoIndex(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Todo Index!",request.URL.Path)
}
func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome!")
}