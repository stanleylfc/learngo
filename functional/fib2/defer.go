package main

import (
	"os"
	"bufio"
	"fmt"
	"github.com/pkg/errors"
)

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i) //参数在defer语句时计算
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	//f, err := os.Create(filename)
	f, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	//err = errors.New("this is a custom error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s",
				pathError.Op,
				pathError.Path,
				pathError.Err,
			)
		}

		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(w, i)
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}
