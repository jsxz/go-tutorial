package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func ProcessInfo(conn net.Conn) {
	buf := make([]byte, 111)
	defer conn.Close()
	for {
		numOfBytes, err2 := conn.Read(buf)
		CheckError(err2)
		if numOfBytes != 0 {
			fmt.Printf("recived message:%s\n", string(buf))
		}
	}
}
func CheckError(err error) {
	if err != nil {
		log.Fatal("error:%s", err)
		os.Exit(1)
	}
}
func main() {
	listen_sockt, err := net.Listen("tcp", "127.0.0.1:8888")
	CheckError(err)
	defer listen_sockt.Close()
	for {
		conn, err1 := listen_sockt.Accept()
		CheckError(err1)
		go ProcessInfo(conn)
	}
}
