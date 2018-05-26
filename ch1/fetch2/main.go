package main

import (
	"time"
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"io"
	"bytes"
)

func main() {
	start := time.Now()
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch2 %s", err)
			os.Exit(1)
		}

		nbytes, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch2 reading %s", err)
		}

		path := strings.Split(url, "/")
		name := path[len(path) - 1]
		out, err := os.Create(name)

		n, err := io.Copy(out, bytes.NewReader(nbytes))

		fmt.Printf("%d %s \n", n, nbytes)
	}

	secs := time.Since(start).Seconds()
	fmt.Printf("%.2f \n", secs)
}
