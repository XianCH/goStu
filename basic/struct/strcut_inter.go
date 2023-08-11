package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {

	m := make(map[string]*student)

	stus := []student{
		{"x14n", 22},
		{"xianchaohao", 18},
	}

	for _, stu := range stus {
		m[stu.name] = &student{name: stu.name, age: stu.age}
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.age)
	}

}
