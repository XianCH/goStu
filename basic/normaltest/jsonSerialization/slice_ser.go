package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m1 := make(map[string]interface{})

	s1 := make([]map[string]interface{}, 0)

	m1["name"] = "kun"
	m1["post"] = "前锋"

	m2 := make(map[string]interface{})
	m2["name"] = "cr7"
	m2["post"] = "前锋"

	s1 = append(s1, m1, m2)

	b, _ := json.Marshal(s1)
	fmt.Println(string(b[:]))
	// fmt.Println(string(b))

}
