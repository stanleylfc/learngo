package main

import (
	"regexp"
	"fmt"
)

const text = `
My email is ccmouse@gmail.com@abc.com
eamil1 is abc@def.org
email2 is kkk@qq.com.cn
`

func main() {
	v2()
}

func v1() {
	re := regexp.MustCompile("ccmouse@gmail.com")
	match := re.FindString(text)
	fmt.Println(match)
}

func v2() {
	//re := regexp.MustCompile(`.+@.+\..+`)
	re := regexp.MustCompile(`[a-zA-z0-9]+@[a-zA-z0-9]+\.[a-zA-Z0-9]+`)
	match := re.FindString(text)
	fmt.Println(match)
}

// 匹配所有
func v3() {
	//re := regexp.MustCompile(`.+@.+\..+`)
	re := regexp.MustCompile(`[a-zA-z0-9]+@[a-zA-z0-9]+\.[a-zA-Z0-9.]+`)
	match := re.FindAllString(text, -1)
	fmt.Println(match)
}

// 正则表达式提取
func v4() {
	//re := regexp.MustCompile(`.+@.+\..+`)
	re := regexp.MustCompile(`([a-zA-z0-9]+)@([a-zA-z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}