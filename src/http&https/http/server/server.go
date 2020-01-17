package main
import (
	"fmt"
	"net/http"
)
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("http server handler.")
}
func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("10.112.26.79:8088", nil)
	if err != nil {
		fmt.Println("server err: ", err.Error())
	}
}
