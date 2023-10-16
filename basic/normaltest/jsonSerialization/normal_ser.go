package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/**
	  定义基础数据类型数据
	*/
	var (
		Surname       = '尹'
		Name          = "尹正杰"
		Age           = 18
		Temperature   = 35.6
		HubeiProvince = false
	)

	/**
	  将基础数据类型进行序列化操作
	*/
	surname, _ := json.Marshal(Surname)
	name, _ := json.Marshal(Name)
	age, _ := json.Marshal(Age)
	temperature, _ := json.Marshal(Temperature)
	hubeiProvince, _ := json.Marshal(HubeiProvince)

	/**
	  查看序列化后的json字符串
	*/
	fmt.Println("Surname序列化后的数据为: ", string(surname))
	fmt.Println("Name序列化后的数据为: ", string(name))
	fmt.Println("Age序列化后的数据为: ", string(age))
	fmt.Println("Temperature序列化后的数据为: ", string(temperature))
	fmt.Println("HubeiProvince序列化后的数据为: ", string(hubeiProvince))
}
