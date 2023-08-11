package main

import "fmt"

type Animal interface{
	Say() string
}

type Dog struct{}

type Cat struct{}

func (c Cat) Say() string{
	return "喵喵毛"
}

func (c Dog) Say() string{
	return "汪汪汪"
}

type Mover interface{
	move()
}

func(c Dog) move(){
	fmt.Println("狗会动")
}

func main() {
	var a Animal
	a = Cat{}
	fmt.Println("猫：",a.Say())

	a = Dog{}
	fmt.Println("狗：",a.Say())


	var x Mover
	var wangcai = Dog{} // 旺财是dog类型
	x = wangcai         // x可以接收dog类型
	var fugui = &Dog{}  // 富贵是*dog类型
	x = fugui           // x可以接收*dog类型
	x.move()
}