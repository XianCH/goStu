package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var s1 map[string]interface{}
	s1 = make(map[string]interface{})

	s2 := make(map[string]interface{})

	s1["name"] = "x14n"
	s1["age"] = 20
	s1["Address"] = "zhongshan"

	s2["name"] = "Jason Yin"
	s2["age"] = 20
	s2["address"] = [2]string{"北京", "陕西"}

	data, err := json.Marshal(s1)
	data2, err := json.Marshal(s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	fmt.Println(string(data2))
}
