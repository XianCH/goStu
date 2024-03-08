package main

import (
	"errors"
	"fmt"
)

type Operation interface {
	Excute(int, int) int
}

type Add struct {
	name string
}

func (a *Add) Excute(x, y int) int {
	return x + y
}

type Subtract struct {
}

func (s *Subtract) Excute(x, y int) int {
	return x - y
}

type Calculate struct {
	Operation
}

func (c *Calculate) SetOperation(o Operation) {
	c.Operation = o
}

func (o *Calculate) PerfomOperation(x, y int) error {
	if o.Operation != nil {
		fmt.Println(o.Excute(x, y))
		return nil
	} else {
		return errors.New("not opration")
	}
}

// func main() {
// 	add := &Add{}
// 	subtract := &Subtract{}
// 	Calculate := &Calculate{}

// 	err := Calculate.PerfomOperation(1, 1)
// 	fmt.Println(err.Error())

// 	Calculate.SetOperation(add)
// 	Calculate.PerfomOperation(1, 1)

// 	Calculate.SetOperation(subtract)
// 	Calculate.PerfomOperation(1, 1)
// }
