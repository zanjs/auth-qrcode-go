package utils

import (
	"fmt"
	"os"
)

var (
	// LogFIlePath is
	LogFIlePath = "log"
)

// Mkdir is
func Mkdir(name string) {
	var path string
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		path = "\\"
	} else {
		path = "/"
	}
	fmt.Println(path)
	dir, _ := os.Getwd()                        //当前的目录
	err := os.Mkdir(dir+path+name, os.ModePerm) //在当前目录下生成目录
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("创建目录" + dir + path + name + "成功")
}
