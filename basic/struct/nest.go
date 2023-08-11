package main

import "fmt"

type User struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	City     string
	Provious string
}

func main() {
	user1 := User{
		Name: "xianchaoaho",
		Age:  18,
		Address: Address{
			City:     "zhongshan",
			Provious: "guangdong",
		},
	}
	fmt.Printf("%#v\n", user1)
}
