package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/**
	  定义数组
	*/
	var s1 = [5]int{9, 5, 2, 7, 5200}
	/**
	  将数组使用Marshal函数进行序列化
	*/

	fmt.Println(s1)

	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("序列化错误: ", err)
		return
	}
	/**
	  查看序列化后的json字符串
	*/
	fmt.Println("数组序列化后的数据为: ", string(data))
}
