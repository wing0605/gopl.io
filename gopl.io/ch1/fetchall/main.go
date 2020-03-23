package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) //创建通道(channel) ch
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // goroutine是一种函数的并发执行方式，而channel是用来在goroutine之间进行参数传递。main函数本身也运行在一个goroutine中，而go function则表示创建一个新的goroutine，并在这个新的goroutine中执行这个函数。
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //从通道ch中接收数据并打印
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	//返回数据的字节数，io.Copy会把响应的Body内容拷贝到ioutil.Discard输出流中
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("耗时：%.2fs  字节数：%7d   网址：%s", secs, nbytes, url)
}
