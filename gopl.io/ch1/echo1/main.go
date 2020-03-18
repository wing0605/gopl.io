package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		//赋值语句, 将s的旧值跟sep与os.Args[i]连接后赋值回s
		//s = s + sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
