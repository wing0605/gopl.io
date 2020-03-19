//练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。

package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(strings.Join(os.Args[1:], " "))
	//range产生一对值；索引以及在该索引处的元素值.
	//os.Args变量是一个字符串（string）的切片（slice）
	//通过for循环，利用range取出索引和元素值，然后打印
	for e, d := range os.Args[1:] {
		fmt.Println(e, d)
	}
}
