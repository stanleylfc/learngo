package main

import (
	"net/url"
	"log"
	"fmt"
)

func main() {
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
	fmt.Println(u.Scheme)
	fmt.Println(u.Opaque)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawPath)
	fmt.Println(u.ForceQuery)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Fragment)
}
