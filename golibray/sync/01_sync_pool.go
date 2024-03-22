package syncTest

import (
	"fmt"
	"sync"
)

var pool *sync.Pool

type Person struct {
	Name string
}

type Usser struct {
	Name string
	Age  string
}

func initPool() {
	pool = &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new Person")
			return new(Person)
		},
	}
}

// func main() {
// 	initPool()
// 	p := pool.Get().(*Person)
// 	fmt.Println("首次从pool中获取对", p)
//
// 	p.Name = "first"
// 	fmt.Println("person 's name", p.Name)
//
// 	pool.Put(p)
//
// 	fmt.Println("Pool 里已有一个对象：&{first}，调用 Get: ", pool.Get().(*Person))
// 	fmt.Println("Pool 没有对象了，调用 Get: ", pool.Get().(*Person))
// }
