package main

import (
	"fmt"

	"github.com/x14n/goStu/basic/interface/mypage"
)

func main() {
	dog := &mypage.Dog{}
	s := dog.Speak()
	c := &mypage.Cat{}
	fmt.Println(s)
	a := c.Speak()
	fmt.Println(a)
}
