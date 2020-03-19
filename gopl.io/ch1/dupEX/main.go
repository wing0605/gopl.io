//练习 1.4： 修改dup2，出现重复的行时打印文件名称。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			//os.Open函数返回两个值。第一个值是被打开的文件(*os.File），其后被Scanner读取。
			//os.Open返回的第二个值是内置error类型的值。如果err等于内置值nil（译注：相当于其它语言里的NULL），那么文件被成功打开。
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("file: %s\t%d\t%s\n", files[0], n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
