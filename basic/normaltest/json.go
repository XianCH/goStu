package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	Name    string
	ID      uint
	Age     uint
	Address string
}

func main() {
	s1 := &Teacher{
		Name:    "x14n",
		ID:      13,
		Age:     40,
		Address: "ZHONG SHAN",
	}

	b, err := json.Marshal(&s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	fmt.Println(string(b))
}
