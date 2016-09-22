package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

/*
 * 这里是尝试写一个监控程序，有如下效果：
 * 1、可以实时看到举报同步的情况
 * 2、可以监测到部分网页是否运转正常，比如是否能够正常服务，不出现502无响应或空白的情况
 * 3、可以监测到部分服务响应的效率，比如请求返回的计时，是否超过了指定时间等
 */

func main() {
	mux := http.NewServeMux() //上例中http.HandleFunc其实就是一个ServeMux//更加底层的操作//自己实现

	mux.HandleFunc("/heartbeat", Heartbeat)

	//一个静态文件服务部分
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(wd))))
	//尝试访问一下以下地址
	//localhost:8080/static/static/1.txt
	//localhost:8080/static/static/
	//localhost:8080/static/

	//err 前面有 则去掉 := 中的 :
	err = http.ListenAndServe(":8081", mux) //两个参数:addr,handler
	if err != nil {
		log.Fatal(err) //打印出错误
	}
}
type HeartbeatItem struct {
    svrsttus string
    svrspeed  string
    jbsync string
}
func Heartbeat(w http.ResponseWriter, r *http.Request)() { //作为要注册路由的hanlder//具体有哪些参数和参数类型等要求看文档
// func Heartbeat(w http.ResponseWriter, r *http.Request)(result []HeartbeatItem, err error) { //作为要注册路由的hanlder//具体有哪些参数和参数类型等要求看文档
	io.WriteString(w, "Heartbeat,this is going!") //输出的字符串
	
    // m:=map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e"}
    // s:=make([]int,len(m))
    // result:=HeartbeatItem{svrsttus:"static result in it",svrspeed:"speed result in it",_jbsync:"jbsync result in it"}
    // result:=
 //   m:=HeartbeatItem{"static result in it","speed result in it","jbsync result in it"}
    
 //   result, err := json.Marshal(m)
	// b == []byte(`{"svrsttus":"static result in it","svrspeed":"speed result in it","jbsync":jbsync result in it}`)

	// return
}
// func Jbsync(w http.ResponseWriter, r *http.Request) {
// 	url="http://4g.womenxing.com/dev/index-test.php?r=site/syncjubao"
// 	resp, err1 := http.Get(url)
// 	if err1 != nil {
// 		return
// 	}
// }
