package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	//time.After实现超时
	select {
	case <-ch:
		fmt.Println("可读")
	case <-time.After(5 * time.Second):
		fmt.Println("超时")
	}
}
