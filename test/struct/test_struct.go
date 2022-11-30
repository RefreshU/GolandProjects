package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	name string
	age  int
}
type UserInfo struct {
	Name string `json:"name"`
	Age int `json:"Age"`
}

func struct2maps(){
	u1 := UserInfo{Name: "dream", Age: 19}
	b, _ := json.Marshal(u1)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m{
		fmt.Printf("key:%v value:%v",k, v)
	}
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		fmt.Printf("name: %p\n", &stu.name)
		m[stu.name] = &stu
	}
	fmt.Printf("address =%p\n", &stus)
	for k, v := range m {
		fmt.Printf("v=%p\n", v)
		fmt.Println(k, "=>", v.age)
	}
}
