package main

import (
	"fmt"
	"time"
)

var ch chan int

func test_channel() {
	ch <- 1
	fmt.Println("ch 1")
	ch <- 1
	fmt.Println("ch 2")
	ch <- 1
	fmt.Println(" wait")
}

func main() {
	ch = make(chan int, 0)
	ch = make(chan int, 2) //带2个缓存的channel
	go test_channel()
	time.Sleep(2 * time.Second)
	fmt.Println("running end")
	<-ch
	time.Sleep(time.Second)
}
