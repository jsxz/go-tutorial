package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	//数组编码
	s, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
	//字典编码
	m := make(map[string]float64)
	m["an"] = 11.4
	s, err = json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
	//对象编码
	student := Student{"xiaoan", 2}
	s, err = json.Marshal(student)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))

	var s3 interface{}
	var s4 Student
	//解码
	json.Unmarshal(s, &s3)
	json.Unmarshal(s, &s4)

	fmt.Println("type:", reflect.TypeOf(s))
	fmt.Println(s3)
	fmt.Println(s4)
}
