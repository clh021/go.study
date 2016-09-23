package main

import (
	"io"
    "io/ioutil"
	"log"
	"net/http"
	"reflect"//用来获取对象的类型
	"encoding/json"//
	"os"
	"fmt"//用来终端调试
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
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(wd+"/static"))))
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
    Svrsttus 	string
    Svrspeed  	string
    Svrjbsync 	string
}

type Server struct {
    ServerName string
    ServerIP   string
}
func Show(item HeartbeatItem, w http.ResponseWriter) {
	fmt.Println(reflect.TypeOf(item.Svrsttus))
	io.WriteString(w,item.Svrsttus+"<hr />")
	io.WriteString(w,item.Svrspeed+"<hr />")
	io.WriteString(w,item.Svrjbsync+"<hr />")
}
func Heartbeat(w http.ResponseWriter, r *http.Request)() { //作为要注册路由的hanlder//具体有哪些参数和参数类型等要求看文档
// func Heartbeat(w http.ResponseWriter, r *http.Request)(result []HeartbeatItem, err error) { //作为要注册路由的hanlder//具体有哪些参数和参数类型等要求看文档
    s := HeartbeatItem{"static result in it<br />","speed result in it<br />",Jbsync()}
    res, err := json.Marshal(s)
    if err != nil {
        fmt.Println("json err:", err)
    } else {
    	w.Header().Set("Content-Type", "application/json")
    }
    // fmt.Println(reflect.TypeOf(res))
    // fmt.Println(string(res))
    
	io.WriteString(w, string(res)) //输出的字符串
	// io.WriteString(w, "Heartbeat,this is going!") //输出的字符串
	
    // Show(res,w)
    
    // result, err := json.Marshal(m)
    // if err != nil {
    // 	fmt.Println(err)
    //     // io.WriteString(w,err.Error())
    // }
    // fmt.Println(string(b))
    // fmt.Printf(result);
    // fmt.Println(reflect.TypeOf(m))
    // fmt.Println(m)
    // fmt.Println(reflect.TypeOf(result))
    // fmt.Println(result)
    // fmt.Println(string(result))
    // io.WriteString(w,reflect.TypeOf(result))//输出结果
	// b == []byte(`{"svrsttus":"static result in it","svrspeed":"speed result in it","jbsync":jbsync result in it}`)

	// return
}
func Get(url string) (content string, statusCode int) {
    resp, err1 := http.Get(url)
    if err1 != nil {
        statusCode = -100
        return
    }
    defer resp.Body.Close()
    data, err2 := ioutil.ReadAll(resp.Body)
    if err2 != nil {
        statusCode = -200
        return
    }
    statusCode = resp.StatusCode
    content = string(data)
    return
}
func Jbsync()(resp string) {
	url:="http://4g.womenxing.com/?r=site/syncjubao3"
	resp, statusCode := Get(url)
	// fmt.Println(reflect.TypeOf(resp))
	// fmt.Println(resp)
	// fmt.Println(statusCode)
	if 200 != statusCode {
		resp = ""
	}
	return
}
