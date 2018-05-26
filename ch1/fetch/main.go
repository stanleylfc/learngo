package main

import (
	"time"
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	start := time.Now()

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)
	}

	secs := time.Since(start).Seconds()
	fmt.Printf("%.2f \n", secs)

}
