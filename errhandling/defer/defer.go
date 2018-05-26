package _defer

import (
	"fmt"
	"os"
	"bufio"
	"github.com/stanleylfc/learngo/errhandling/defer/fib"
)

func tryDefer() (string) {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred") // 3 2 1
	fmt.Println(4)
	defer fmt.Println(5)
	return "tryDefer function end"
}

// defer 意义： 管理资源
// defer 使用 open/close lock/unlock printHeader/printFooter
func writeFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file) //增加文件写的速度
	defer writer.Flush()

	f := fib.Fibonacci()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//method := tryDefer()
	//fmt.Println(method)

	writeFile("fib.txt")
}
