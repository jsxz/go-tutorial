package main

import (
	"fmt"
	"regexp"
)

// 参考 https://www.cnblogs.com/golove/p/3269099.html
func main() {
	//matched, _ := regexp.Match("[a-zA-Z0-9]{3}", []byte("zh1"))
	//matched, _ := regexp.MatchString("[a-zA-Z0-9]{3}", "zh14")
	compile := regexp.MustCompile("[a-z]([1-9])")
	allString := compile.FindAllString("zaf", 32)
	allString2 := compile.FindAllStringSubmatch("zaf32", 32)
	fmt.Printf("%v", allString)
	fmt.Printf("%v", allString2)
}
