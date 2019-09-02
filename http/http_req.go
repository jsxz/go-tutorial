package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Req() {
	//resp, err := http.Get("http://baidu.com")
	resp, err := http.Post("http://baidu.com", "application/x-www-form-urlencoded", strings.NewReader("id=1"))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}
func main() {
	Req()
}
