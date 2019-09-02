package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	hash := md5.New()
	hash.Write([]byte("123456"))
	sum := hash.Sum([]byte(""))
	fmt.Printf("%x\n\n", sum)

}
