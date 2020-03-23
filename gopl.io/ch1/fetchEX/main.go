// 练习 1.7： 函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中，使用这个函数替代掉例子中的ioutil.ReadAll来拷贝响应结构体到os.Stdout，避免申请一个缓冲区（例子中的b）来存储。记得处理io.Copy返回结果中的错误。

// 练习 1.8： 修改fetch这个范例，如果输入的url参数没有  http://  前缀的话，为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。

// 练习 1.9： 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") { //strings.HasPrefix()函数用来检测字符串是否以指定的前缀开头。
			url = strings.Join([]string{"http://", url}, "")
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v \n", err)
			os.Exit(1)
		}

		b, err := io.Copy(os.Stdout, resp.Body)
		//io.Copy() 方法将副本从 src 复制到 dst ，直到 src 达到文件末尾 ( EOF ) 或发生错误，然后返回复制的字节数和复制时遇到的第一个错误（ 如果有 ）有了这个函数，我们就省去了先把内容读取到内存，然后将内存中的内容写到文件的过程
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}

		s := resp.Status

		fmt.Printf("%v\n %s", b, s)
	}
}
