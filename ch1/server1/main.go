package main

import (
	"net/http"
	"github.com/prometheus/common/log"
	"fmt"
)

func main() {
	http.HandleFunc("/", hanlder)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path=%q\n", r.URL.Path)
}
