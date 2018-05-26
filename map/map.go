package main

import "fmt"

func main() {

	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	fmt.Println(m)

	m2 := make(map[string]int) //m2 == empty
	m3 := map[string]int{}     // m3= nil

	m2["abc"] = 2
	m3["cba"] = 3
	fmt.Println(m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Deleting values")

	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")

	name, ok = m["name"]
	fmt.Println(name, ok)

}
