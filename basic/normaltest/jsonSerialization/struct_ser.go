package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	ID      int16
	Name    string
	Age     int16
	Address string
}

func main() {
	s1 := &Teacher{
		ID:      10,
		Name:    "x14n",
		Age:     30,
		Address: "zhong shan",
	}

	data, err := json.Marshal(&s1)
	if err != nil {
		fmt.Println("输出错误")
	}
	s := string(data)
	fmt.Println(s)
}
