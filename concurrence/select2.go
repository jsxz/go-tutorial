package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// select 超时经典实现
	timeout := make(chan int, 0)
	go func() {
		time.Sleep(5 * time.Second)
		timeout <- 1
	}()

	select {
	case <-ch:
		fmt.Println("可读")
	case <-timeout:
		fmt.Println("超时")
	}
}
