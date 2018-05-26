package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	v2()
}


func v1() {
	for i:= 0; i< 10; i++ {
		go func(i int) {
			for {
				fmt.Printf("hello from goroutine %d\n", i)  //io 操作有等待过程
			}
		}(i)
	}

	time.Sleep(time.Millisecond)
}

func v2() {
	var a [10]int
	for i:=0; i<1; i++ {
		go func(i int) {
			for{
				a[i]++ //goroutine  不交出控制权
				//runtime.Gosched()  //手动交出控制权
			}
		}(i)
	}

	//控制权未交出， time.Sleep 不执行
	time.Sleep(time.Millisecond)
}

func v3() {
	var a [10]int
	for i:=0; i<10; i++ {
		go func() {
			for{
				a[i]++ //goroutine  不交出控制权
				runtime.Gosched()  //手动交出控制权
			}
		}()
	}

	//控制权未交出， time.Sleep 不执行
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

func v4() {
	var a [10]int
	for i:=0; i<10; i++ {
		go func(i int) {
			for{
				a[i]++ //goroutine  不交出控制权
				runtime.Gosched()  //手动交出控制权
			}
		}(i)
	}

	//控制权未交出， time.Sleep 不执行
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}