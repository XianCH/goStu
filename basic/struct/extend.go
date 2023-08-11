package main

import "fmt"

type Dog struct {
	Feet int
	*Animal
}

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪叫~\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "猪猪",
		},
	}
	d1.wang()
	d1.move()
}
