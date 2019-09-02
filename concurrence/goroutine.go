package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var ch = make(chan int)
	go func() {
		for i := 1; i < 111; i++ {
			if i == 21 {
				//runtime.Gosched()//主动让出cpu
				<-ch
			}
			fmt.Println("routine 1 " + strconv.Itoa(i))
		}
	}()
	go func() {
		for i := 112; i < 222; i++ {
			fmt.Println("routine 2 " + strconv.Itoa(i))
		}
		ch <- 1
	}()
	time.Sleep(2 * time.Second)
}
