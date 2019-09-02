package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	//go func(ch chan int) {
	//	ch<-1
	//}(ch)
	go func() {
		ch <- 1
	}()
	time.Sleep(1 * time.Second)
	select {
	case <-ch:
		fmt.Println("可读")
	default:
		fmt.Println("default")
	}
}
