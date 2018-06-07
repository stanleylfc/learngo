package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func main() {
	request, err := http.NewRequest(
		http.MethodGet,
		"http://www.imooc.com", nil)

	request.Header.Add("User-Agent", "")

	resp, err:= http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	fmt.Printf("%s/n", s)
}
