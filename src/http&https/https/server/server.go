package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!   liuxin")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8080", "/home/liuxin/Desktop/server.crt",
		"/home/liuxin/Desktop/server.key", nil)
}