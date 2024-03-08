package main

import "fmt"

type Speaker interface {
	Speaker() string
	Introduce() string
}

type Person struct {
	name    string
	Age     string
	Address string
}

func (p *Person) Speaker() string {
	return p.name
}

func (p *Person) Introduce() string {
	return fmt.Sprintf("hello my name is %s,im %s yeas old,im from %s", p.name, p.Age, p.Address)
}

func DoSpeak(s Speaker) {
	fmt.Println(s.Speaker())
	fmt.Println(s.Introduce())
}

func main() {
	x14n := &Person{
		name:    "x14n",
		Age:     "23",
		Address: "maoming",
	}
	DoSpeak(x14n)
}
