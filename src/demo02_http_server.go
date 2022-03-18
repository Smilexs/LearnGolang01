package src

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
)

var (
	port = flag.Int("port", 8090, "port to listen on")
)


func main(){
	//	var a []int
	//	a = append(a, []int{1,2,3}...) //追加一个切片, 切片需要解包
	//	for i := 0; i < len(a); i++ {
	//
	//		fmt.Println(a[i])
	//	}

	//测试： 在cmd中执行 go run main.go 或者直接IDE中运行
	//浏览器输入：localhost:8090/data
	//解析http
	flag.Parse()
	//监听端口
	ListenHttpPort(*port)
	fmt.Println("listen")
}


func ListenHttpPort(port int){
	http.Handle("/data", http.HandlerFunc(HttpDataHandler))
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func HttpDataHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("http Hello world"))
}
