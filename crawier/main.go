package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"bufio"
)

func main() {
	resp, err := http.Get(
		"http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//httputil.DumpResponse()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",
			resp.StatusCode)
		return
	}

	// 代码识别 编码
	e := determineEncoding(resp.Body)
	//gbk 转 utf8 手动写死 GBk
	//utf8Reader := transform.NewReader(resp.Body,
	//	simplifiedchinese.GBK.NewDecoder())

	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", all)
}

func determineEncoding(
	r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
