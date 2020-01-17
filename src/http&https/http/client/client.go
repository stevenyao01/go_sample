package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func main() {
	resp, err := http.Get("https://10.112.26.79:8080")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
