package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	m1 := `{"address":["北京","陕西"],"age":20,"name":"Jason Yin"}`

	/**
	  定义map变量，类型必须与之前序列化的类型完全一致。
	*/
	var s1 map[string]interface{}
	fmt.Println("反序列化之前：s1 =", s1)

	/**
	  温馨提示:
	      不需要使用make函数给m初始化，开辟空间。这是因为在反序列化函数Unmarshal()中会判断传入的参数2，如果是map类型数据，会自动开辟空间。相当于是Unmarshal()函数可以帮助我们做make操作。
	      但传参时需要注意，Unmarshal的第二个参数，是用作传出，返回结果的。因此必须传m的地址值。
	*/
	err := json.Unmarshal([]byte(m1), &s1)
	if err != nil {
		fmt.Println("反序列化失败,错误原因: ", err)
		return
	}

	fmt.Println("反序列化之后：s1 =", s1)
}
