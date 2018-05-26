package main

import "fmt"

func adder() func(int) int {
	sum := 0 //自由变量

	return func(v int) int {
		x := 0 //局部变量
		x = x + v
		fmt.Println(v, x)
		sum += v
		return sum
	}

}
func main() {
	a := adder()
	for i := 0; i < 4; i++ {
		fmt.Println(i, a(i))
	}
}
