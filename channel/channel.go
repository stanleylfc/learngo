package main

import "fmt"

func chanDemo1() {
	//var c chan int  //c == nil 不能使用
	c := make(chan int) //创建channel, 做一个channel
	// channel 是两个协程的通信
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()

	c <- 1
	c <- 2
}

func worker(id int, c chan int) {
	for {
		fmt.Printf("worker %d received %c \n", id, <-c)
	}
}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
}





func main() {
	chanDemo()
}
