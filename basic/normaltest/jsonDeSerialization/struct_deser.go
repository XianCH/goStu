package main

import (
	"encoding/json"
	"fmt"
)

type Peopel struct {
	Name    string
	Age     int16
	Address string
}

func main() {

	p1 := `{"Name":"Jason Yin","Age":18,"Address":"北京"}`

	var s1 Peopel

	fmt.Printf("反序列化之前: \n\ts1 = %v \n\ts1.Name = %s\n\n", s1, s1.Name)

	err := json.Unmarshal([]byte(p1), &s1)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}

	fmt.Printf("反序列化之后: \n\ts1 = %v \n\ts1.Name = %s\n", s1, s1.Name)

}
