package main

import (
	"net/http"
	"github.com/stanleylfc/learngo/errhandling/filelistingserver/filelisting"
	"os"
	"log"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

//统一
//函数式编程
func errWarpper(
	handler appHandler) func(
	http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		err := handler(writer, request) //handler 处理业务逻辑
		if err != nil {
			log.Printf("error occurred %s", err.Error())
			//log.Warn("error handling request: %s",
			//	err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError

			}
			http.Error(writer,
				http.StatusText(code), code)
		}
	}

}

func main() {
	http.HandleFunc("/list/",
		errWarpper(filelisting.HandlerFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
