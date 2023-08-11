package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p1 person
	p1.name = "x14n"
	p1.city = "maoming"
	p1.age = 22
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)

	var user struct {
		Name string
		age  int
	}
	user.Name = "x14n.cn"
	user.age = 22
	fmt.Printf("%#v\n", user)

	p3 := &person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("p3=%#v\n", p3)
	p3.name = "xianchaoahoa"
	p3.age = 30
	p3.city = "zhongshan"
	fmt.Printf("p3=%#v\n", p3)

	fmt.Println("")
	fmt.Println("//////////结构体初始化")
	var p4 person
	fmt.Printf("p4=%#v\n", p4)

	p5 := person{
		name: "x14n.cn",
		city: "guangzhou",
		age:  90,
	}
	fmt.Printf("p5=%#v\n", p5)
}
