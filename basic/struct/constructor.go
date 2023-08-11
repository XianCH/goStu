package main

import "fmt"

func newPerson(name string, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

type person struct {
	name string
	city string
	age  int
}

func (p *person) SetAge(newAge int) {
	p.age = newAge
}

func (p person) SetAge2(newAge int) {
	p.age = newAge
}

func (p person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言\n", p.name)
}

func main() {
	p9 := newPerson("X14N", "ZHONGSHAN", 22)
	fmt.Printf("%#v\n", p9)

	p1 := newPerson("x14n1", "guangzhou", 22)
	fmt.Println(p1.age)
	p1.SetAge(30)
	fmt.Println(p1.age)

	p2 := newPerson("冼超豪", "dontknow", 90)
	p2.Dream()
	fmt.Println(p2.name)
	p2.SetAge2(30)
	fmt.Println(p2.age)
}
