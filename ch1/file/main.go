package main

import (
	"fmt"
	"path"
	"strings"
)

func main() {
	pathFilename := "/Users/lifeicheng/Workspaces/go/src/github.com/stanleylfc/learngo/ch1/file/main.go"
//	pathFilename := "http://tvax3.sinaimg.cn/crop.0.0.199.199.1024/006XhRHpgy1fjriolc8tjj305k05kmwy.jpg"
	fmt.Println("fullFilename =", pathFilename)

	dir := path.Dir(pathFilename)
	fmt.Println("file dir :", dir)

	filenameWithSuffix := path.Base(pathFilename) //后缀
	fmt.Println("filenameWithSuffix =", filenameWithSuffix)

	fileSuffix := path.Ext(filenameWithSuffix)  //后缀
	fmt.Println("filesuffix =", fileSuffix)

	filename := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	fmt.Println("filenam =", filename)
}
