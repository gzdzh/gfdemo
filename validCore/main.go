package main

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
)

func Valid() {
	contents := gfile.GetContents("cert.key")
	if contents != "123456" {
		fmt.Println("不通过")

		//主程序执行静态编译的文件，停止不到主程序
		//panic("not pass")
	}

}

func main() {
	Valid()

}
