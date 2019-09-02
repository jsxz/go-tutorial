package main

import (
	"fmt"
	"strconv"
)

func Read(ch chan int) {
	value := <-ch
	fmt.Println("value:" + strconv.Itoa(value))
}
func Write(ch chan int) {
	ch <- 3335
}
func Add(x, y int, quit chan int) {
	z := x + y
	fmt.Println(z)
	quit <- 1
}

func main() {
	//ch := make(chan int)
	//go Read(ch)
	//go Write(ch)
	//time.Sleep(1)
	//fmt.Println("end ")
	chs := make([]chan int, 11)
	for i := 0; i < 11; i++ {
		chs[i] = make(chan int)
		go Add(i, i, chs[i])
	}
	for _, v := range chs {
		<-v
	}
}
