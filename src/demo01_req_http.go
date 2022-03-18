package src

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main_Demo01() {
	//fmt.Println("Hello world hh") //打印

	//1.获取应用程序输入的参数
	url := os.Args[1]

	//2.根据url获取资源
	res,err := http.Get(url)
	if err != nil { //异常处理
		fmt.Println(os.Stderr, "抓取网页信息异常:%v\n", err)
		os.Exit(1)
	}

	//3.读取资源
	content,err := ioutil.ReadAll(res.Body)
	//4.关闭资源流
	res.Body.Close()
	if err != nil { //异常处理
		fmt.Println(os.Stderr, "读取抓取的信息异常 %s: %v\n", url, err)
		os.Exit(1)
	}

	//打印输出的内容
	fmt.Println("%s", content)
	fmt.Println(string(content))
}
