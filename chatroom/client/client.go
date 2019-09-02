package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal("error:%s", err.Error())
		os.Exit(1)
	}
}
func main() {
	conn, e := net.Dial("tcp", "127.0.0.1:8888")
	CheckError(e)
	defer conn.Close()
	conn.Write([]byte("hello "))
	fmt.Println("this a message")

}
