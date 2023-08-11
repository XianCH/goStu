package main

import "fmt"

func recv(c chan int) {
	ret := <-c
	fmt.Println("接受成功",ret)
}

func main() {
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	fmt.Println("发送成功")

	//缓冲通道
	ch1 := make(chan int,1)
	ch1 <- 20
	fmt.Println("发送成功")

}