package main

import (
	"fmt"
	"strings"
)

//闭包的使用 代替了递归
func fib() intGen {
	a, b := 0, 1
	return func() int {
		fmt.Println("in")
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func main() {
	f := fib()

	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
