package main

import "fmt"

type Celsius float64
type Fahrenheit float64

type node struct {
	str   string
	Left  *node
	b     byte
	r     rune
	Right *node
	Value int
	flag  bool
	f     float64
}

func main() {
	var c Celsius
	var f Fahrenheit

	c = 20
	f = 40
	fmt.Println(c, f)
	fmt.Println(c == Celsius(f))
	fmt.Println(Celsius(f))

	var n node

	fmt.Println(n)

	var str string
	fmt.Println(str)

	var s []int
	fmt.Println(s == nil)
}
