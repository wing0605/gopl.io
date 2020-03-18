package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	//变量声明 等价于
	//s := ""
	// var s string
	// var s = ""
	// var s string = ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
