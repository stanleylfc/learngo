package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for i, url := range os.Args[1:] {
		go fetch(i, url, ch) //不确定性
	}

	for range os.Args[1:] {
		//fmt.Printf("main %d %s ", i, url)
		fmt.Println(<-ch)  //阻塞
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(i int, url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s:%v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	fmt.Printf("fetch %s", url)
	ch <- fmt.Sprintf(" %d %0.2fs %d %s", i, secs, nbytes, url)
}
